/*
 * File: elasticHandler.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 10:41:01 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/elasticsearch"
	"ABD4/API/iserial"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetCreateIndexation create all the indexes based on filename in elasticsearch/esmapping
func GetCreateIndexation(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	err := elasticsearch.CreateIndexation(ctx.ElasticClient, false)
	if err != nil {
		msg := fmt.Sprintf("%s indexes creation failed", utils.Use().GetStack(GetCreateIndexation))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	indexes, _ := json.Marshal(context.INDEXES)
	msg := fmt.Sprintf("%s Indexes: %s sucessfully created", utils.Use().GetStack(GetCreateIndexation), string(indexes))
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, string(indexes), "")
	return
}

// GetCreateIndex handler for Get /elastic/index/{index}
func GetCreateIndex(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetCreateIndex))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetCreateIndex))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	err := elasticsearch.Index(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s index %s creation failed", utils.Use().GetStack(GetCreateIndex), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully created", utils.Use().GetStack(GetCreateIndex), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}

// GetRemoveIndexation remove all the indexes based on filename in elasticsearch/esmapping
// handler for GET /elastic/rmindex/all
func GetRemoveIndexation(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	indexes, _ := json.Marshal(context.INDEXES)
	for _, index := range context.INDEXES {
		err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
		if err != nil {
			msg := fmt.Sprintf("%s indexes: %s removal failed", utils.Use().GetStack(GetRemoveIndexation), string(indexes))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s Indexes: %s sucessfully created", utils.Use().GetStack(GetRemoveIndexation), string(indexes))
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, string(indexes), "")
	return
}

// GetRemoveIndex handler for GET /elastic/rmindex/{index}
func GetRemoveIndex(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetRemoveIndex))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetRemoveIndex))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s index %s removing failed", utils.Use().GetStack(GetRemoveIndex), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully removed", utils.Use().GetStack(GetRemoveIndex), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}

// GetReindexation handler for GET /elastic/reindex/all
func GetReindexation(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	indexes, _ := json.Marshal(context.INDEXES)
	err := elasticsearch.CreateIndexation(ctx.ElasticClient, true)
	if err != nil {
		msg := fmt.Sprintf("%s indexes creation failed", utils.Use().GetStack(GetReindexation))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Indexes: %s sucessfully reindexed", utils.Use().GetStack(GetReindexation), string(indexes))
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, string(indexes), "")
	return
}

// GetReindexation handler for GET /elastic/reindex/all
func GetReindex(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetReindex))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetReindex))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s Index: %s removal failed", utils.Use().GetStack(GetReindex), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	err = elasticsearch.Index(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s Index creation failed", utils.Use().GetStack(GetReindex))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully reindexed", utils.Use().GetStack(GetReindex), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}

func GetIndexationData(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	// On itère sur les indexes possibles
	for entity, index := range context.INDEXES {
		// variable servant à la sérialisation
		var toSerialize []iserial.Serializable
		// en fonction de l'entité concernée par l'index
		// on récupère les données à partir du manager correspondant
		switch entity {
		case context.USER:
			users, err := ctx.UserManager.FindAll()
			if err != nil {
				msg := fmt.Sprintf("%s failed to retrieve users", utils.Use().GetStack(GetIndexationData))
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
				return
			}
			for _, u := range users {
				toSerialize = append(toSerialize, u.ToES())
			}
		case context.TX:
			txs, err := ctx.TransactionManager.FindAll()
			if err != nil {
				msg := fmt.Sprintf("%s failed to retrieve transactions", utils.Use().GetStack(GetIndexationData))
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
				return
			}
			for _, tx := range txs {
				toSerialize = append(toSerialize, tx.ToES())
			}
		}
		// on indexe les données correspondantes, sur un lot important, cela peut prendre du temps
		err := ctx.IndexArrayData(toSerialize, index, entity)
		if err != nil {
			msg := fmt.Sprintf("%s failed to index %s", utils.Use().GetStack(GetIndexationData), index)
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s data successfully indexed", utils.Use().GetStack(GetIndexationData))
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
	return
}

func GetIndexData(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetIndexData))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetIndexData))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	var entity string
	// variable servant à la sérialisation
	var toSerialize []iserial.Serializable
	// en fonction de l'entité concernée par l'index
	// on récupère les données à partir du manager correspondant
	switch index {
	case context.USERS:
		entity = context.USER
		users, err := ctx.UserManager.FindAll()
		if err != nil {
			msg := fmt.Sprintf("%s failed to retrieve users", utils.Use().GetStack(GetIndexData))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		for _, u := range users {
			toSerialize = append(toSerialize, u.ToES())
		}
	case context.TXs:
		entity = context.TX
		txs, err := ctx.TransactionManager.FindAll()
		if err != nil {
			msg := fmt.Sprintf("%s failed to retrieve transactions", utils.Use().GetStack(GetIndexData))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		for _, tx := range txs {
			toSerialize = append(toSerialize, tx.ToES())
		}
	}
	// on indexe les données correspondantes, sur un lot important, cela peut prendre du temps
	err := ctx.IndexArrayData(toSerialize, index, entity)
	if err != nil {
		msg := fmt.Sprintf("%s failed to index %s", utils.Use().GetStack(GetIndexData), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s data successfully indexed", utils.Use().GetStack(GetIndexData))
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
	return
}
