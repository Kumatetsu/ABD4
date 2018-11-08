/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 7:42:41 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 1:53:52 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/model"
	"ABD4/API/service"
	"ABD4/API/utils"
	"fmt"
	"net/http"
	"runtime"
	"time"
)

// GetTransaction return all transaction in database
func GetTransaction(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	var toSerialize []iserial.Serializable
	ctx.Log.Info.Printf("%s %s ", utils.Use().GetStack(GetTransaction), "Getting Transactions")
	tx, err := ctx.TransactionManager.FindAll()
	if err != nil {
		msg := fmt.Sprintf("%s FindAll failed", utils.Use().GetStack(GetTransaction))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	for _, t := range tx {
		toSerialize = append(toSerialize, t)
	}
	msg := fmt.Sprintf("%s %d transactions returned", utils.Use().GetStack(GetTransaction), len(toSerialize))
	ctx.Rw.SendArraySerializable(ctx, w, http.StatusOK, toSerialize, msg, "")
}

// synchronousIndexation : elastic search indexation, a non critical operation
// as far as data are existing in mongo database.
func synchronousIndexation(ctx *context.AppContext, transaction *model.Transaction) {
	// on se prépare à benchmarker l'indexation
	ctx.SavedTime = ctx.Time
	ctx.Time = time.Now()
	// on index
	err := ctx.IndexData(transaction.ToES(), context.TXs, context.TX)
	if err != nil {
		msg := fmt.Sprintf("%s failed to index transaction in elasticsearch", utils.Use().GetStack(AddTransaction))
		ctx.Log.Error.Printf(msg)
	}
	ctx.Log.Info.Printf("%s successfull indexation of new %s", utils.Use().GetStack(AddTransaction), context.TX)
	// on log le temps de l'indexation dans le writer de Benchmark
	service.Benchmark(ctx, "ElasticSearch indexation")
	ctx.Time = ctx.SavedTime
}

// asynchronousIndexation : elastic search indexation with mutex locking on ctx
// and internal goroutine counter decrementing
func asynchronousIndexation(ctx *context.AppContext, transaction *model.Transaction) {
	// code asynchrone, on bloque l'accès à l'espace mémoire de ctx le temps de la requête
	ctx.Lock()
	defer ctx.Unlock()
	defer ctx.RmGorout()

	// on se prépare à benchmarker l'indexation
	ctx.SavedTime = ctx.Time
	ctx.Time = time.Now()
	// on index
	err := ctx.IndexData(transaction.ToES(), context.TXs, context.TX)
	if err != nil {
		msg := fmt.Sprintf("%s failed to index transaction in elasticsearch", utils.Use().GetStack(AddTransaction))
		ctx.Log.Error.Printf(msg)
	}
	ctx.Log.Info.Printf("%s successfull indexation of new %s", utils.Use().GetStack(AddTransaction), context.TX)
	// on log le temps de l'indexation dans le writer de Benchmark
	service.Benchmark(ctx, "ElasticSearch indexation")
	ctx.Time = ctx.SavedTime
}

func esTransactionIndexation(ctx *context.AppContext, transaction *model.Transaction) {
	// Si on est pour l'asynchrone
	if ctx.Opts.GetAllowAsync() {
		// On autorise le parallélisme sur 2 processeur internes
		runtime.GOMAXPROCS(2)
		ctx.AddGorout()
		// Après avoir incrémenté le compteur interne
		// on vérifie qu'on est pas au dessus de la limite définie
		if ctx.GoroutOverflow() {
			for ctx.GoroutOverflow() {
				// tant qu'on a trop de goroutines on patiente
				time.Sleep(time.Millisecond * 50)
			}
		}
		// chaque indexation est effectuée dans une goroutine, ce code n'est pas bloquant
		// la requête sera considérée valide une fois les data en base.
		go asynchronousIndexation(ctx, transaction)
	} else {
		synchronousIndexation(ctx, transaction)
	}
}

// AddTransaction add a transaction
func AddTransaction(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	transaction := &model.Transaction{}
	err := transaction.UnmarshalFromRequest(r)
	service.Theme(ctx).DefineTheme(transaction)
	ctx.Log.Info.Printf("%s ", transaction.ToString())
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Decode request data failed", err.Error())
		return
	}
	err = service.Tarif(ctx).CalculateTotal(transaction)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "defining price failed", err.Error())
		return
	}
	ctx.SavedTime = ctx.Time
	ctx.Time = time.Now()
	transaction, err = ctx.TransactionManager.Create(transaction)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "Insert transaction in mongo failed", err.Error())
		return
	}
	service.Benchmark(ctx, "Mongo insert")
	ctx.Time = ctx.SavedTime
	// si on utilise elastic search on index la nouvelle transaction
	if ctx.Opts.GetEmbedES() {
		esTransactionIndexation(ctx, transaction)
	}
	ctx.Rw.SendSerializable(ctx, w, http.StatusCreated, transaction, "", "")
	return
}

// RemoveAllTX destoy everything
func RemoveAllTX(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	deleted, err := ctx.TransactionManager.RemoveAll()
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s from %s", utils.Use().GetStack(RemoveAllTX), context.TXs, ctx.Opts.GetDatabaseType())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	err = ctx.RemoveIndex(context.TXs)
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s index", utils.Use().GetStack(RemoveAllTX), context.TXs)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveAllTX), deleted, context.TX)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}
