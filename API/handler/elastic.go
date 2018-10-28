/*
 * File: elasticHandler.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 7:11:57 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/elasticsearch"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

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
	msg := fmt.Sprintf("%s Indexes: %s sucessfully created", utils.Use().GetStack(GetCreateIndexation), strings.Join(context.INDEXES, ","))
	indexes, _ := json.Marshal(context.INDEXES)
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
		msg := fmt.Sprintf("%s index %s creation failed", utils.Use().GetStack(GetCreateIndexation), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully created", utils.Use().GetStack(GetCreateIndexation), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}

// GetRemoveIndexation remove all the indexes based on filename in elasticsearch/esmapping
// handler for GET /elastic/rmindex/all
func GetRemoveIndexation(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	for _, index := range context.INDEXES {
		err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
		if err != nil {
			msg := fmt.Sprintf("%s indexes: %s removal failed", utils.Use().GetStack(GetCreateIndexation), strings.Join(context.INDEXES, ","))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s Indexes: %s sucessfully created", utils.Use().GetStack(GetCreateIndexation), strings.Join(context.INDEXES, ","))
	indexes, _ := json.Marshal(context.INDEXES)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, string(indexes), "")
	return
}

// GetRemoveIndex handler for GET /elastic/rmindex/{index}
func GetRemoveIndex(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetCreateIndex))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetCreateIndex))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s index %s removing failed", utils.Use().GetStack(GetCreateIndexation), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully removed", utils.Use().GetStack(GetCreateIndexation), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}

// GetReindexation handler for GET /elastic/reindex/all
func GetReindexation(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	for _, index := range context.INDEXES {
		err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
		if err != nil {
			msg := fmt.Sprintf("%s indexes: %s removal failed", utils.Use().GetStack(GetCreateIndexation), strings.Join(context.INDEXES, ","))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	err := elasticsearch.CreateIndexation(ctx.ElasticClient, false)
	if err != nil {
		msg := fmt.Sprintf("%s indexes creation failed", utils.Use().GetStack(GetCreateIndexation))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Indexes: %s sucessfully reindexed", utils.Use().GetStack(GetCreateIndexation), strings.Join(context.INDEXES, ","))
	indexes, _ := json.Marshal(context.INDEXES)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, string(indexes), "")
	return
}

// GetReindexation handler for GET /elastic/reindex/all
func GetReindex(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	index, ok := mux.Vars(r)["index"]
	if !ok {
		msg := fmt.Sprintf("%s missing index parameter", utils.Use().GetStack(GetCreateIndex))
		err := fmt.Errorf("%s no index parameter provide in url request", utils.Use().GetStack(GetCreateIndex))
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, msg, err.Error())
		return
	}
	err := elasticsearch.RemoveIndex(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s Index: %s removal failed", utils.Use().GetStack(GetCreateIndexation), index)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	err = elasticsearch.Index(ctx.ElasticClient, index)
	if err != nil {
		msg := fmt.Sprintf("%s Index creation failed", utils.Use().GetStack(GetCreateIndexation))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s Index: %s sucessfully reindexed", utils.Use().GetStack(GetCreateIndexation), index)
	ctx.Rw.SendString(ctx, w, http.StatusOK, msg, index, "")
	return
}
