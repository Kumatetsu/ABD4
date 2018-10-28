/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 7:42:41 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 6:32:26 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/model"
	"ABD4/API/utils"
	"fmt"
	"net/http"
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
	ctx.Rw.SendArraySerializable(ctx, w, http.StatusOK, toSerialize, "", "")
}

// AddTransaction add a transaction
func AddTransaction(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	transaction := &model.Transaction{}
	err := transaction.UnmarshalFromRequest(r)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Decode request data failed", err.Error())
		return
	}
	transaction, err = ctx.TransactionManager.Create(transaction)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "Insert transaction in mongo failed", err.Error())
		return
	}
	// si on utilise elastic search on index la nouvelle transaction
	if ctx.Opts.GetEmbedES() {
		err = ctx.IndexData(transaction, context.TXs, context.TX)
		if err != nil {
			msg := fmt.Sprintf("%s failed to index transaction in elasticsearch", utils.Use().GetStack(AddTransaction))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		ctx.Log.Info.Printf("%s successfull indexation of new %s", utils.Use().GetStack(AddTransaction), context.TX)
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
	}
	err = ctx.RemoveIndex(context.TXs)
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s index", utils.Use().GetStack(RemoveAllTX), context.TXs)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveAllTX), deleted, context.TX)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}
