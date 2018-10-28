/*
 * File: road.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 8:13:27 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import (
	"ABD4/API/context"
	"ABD4/API/handler"
	"ABD4/API/utils"

	"github.com/gorilla/mux"
)

const (
	GET     = "GET"
	POST    = "POST"
	OPTIONS = "OPTIONS"
	PUT     = "PUT"
	DELETE  = "DELETE"
)

// RoadGetter type for router definition
type RoadGetter func() []*Road

// Road define what a router entry must embed
type Road struct {
	Name            string
	Method          string
	Pattern         string
	StatusProtected bool
	HandlerFunc     context.CustomHandler
}

// appendTo log the happened road and add it to the router section with a prepared handler
func (r *Road) appendTo(ctx *context.AppContext, preparedHandler *context.HandlerWrapper, router *mux.Router) {
	ctx.Log.Info.Printf("%s __ Load... __ %v", utils.Use().GetStack(r.appendTo), r.Name)
	router.StrictSlash(true).Handle(r.Pattern, preparedHandler).Methods(r.Method).Name(r.Name)
}

// GetRoadKit must return a map wih a key for the
// mux.Router router.Pathprefix(key).Subrouter() method as key (see appendRoadKit in router.go)
// and a roadGetter function as value
func GetRoadKit() map[string]RoadGetter {
	return map[string]RoadGetter{
		// exemple for localhost:8000/user/* set of road:
		"/user":        getUserRouting,
		"/auth":        getAuthRouting,
		"/backup":      getBackupRouting,
		"/elastic":     getElasticRouting,
		"/transaction": getTransactionRouting,
	}
}

// getUserRouting return the /user routing
func getUserRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /user",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.GetUsers,
		},
		{
			Name:            DELETE + " /user",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveAllUsers,
		},
	}
}

func getAuthRouting() []*Road {
	return []*Road{
		{
			Name:            OPTIONS + " /auth/login",
			Method:          OPTIONS,
			Pattern:         "/login",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            OPTIONS + " /auth/register",
			Method:          OPTIONS,
			Pattern:         "/register",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /auth/login",
			Method:          POST,
			Pattern:         "/login",
			StatusProtected: false,
			HandlerFunc:     handler.Login,
		},
		{
			Name:            POST + " /auth/register",
			Method:          POST,
			Pattern:         "/register",
			StatusProtected: false,
			HandlerFunc:     handler.Register,
		},
	}
}

func getBackupRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /backup",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.BackupBoltDatabase,
		},
	}
}

func getElasticRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /elastic/index/all",
			Method:          GET,
			Pattern:         "/index/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetCreateIndexation,
		},
		{
			Name:            GET + " /elastic/index/{index}",
			Method:          GET,
			Pattern:         "/index/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetCreateIndex,
		},
		{
			Name:            GET + " /elastic/rmindex/all",
			Method:          GET,
			Pattern:         "/rmindex/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetRemoveIndexation,
		},
		{
			Name:            GET + " /elastic/rmindex/{index}",
			Method:          GET,
			Pattern:         "/rmindex/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetRemoveIndex,
		},
		{
			Name:            GET + " /elastic/reindex/all",
			Method:          GET,
			Pattern:         "/reindex/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetReindexation,
		},
		{
			Name:            GET + " /elastic/reindex/{index}",
			Method:          GET,
			Pattern:         "/reindex/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetReindex,
		},
		{
			Name:            GET + " /elastic/indexdata",
			Method:          GET,
			Pattern:         "/indexdata",
			StatusProtected: false,
			HandlerFunc:     handler.GetIndexationData,
		},
		{
			Name:            GET + " /elastic/indexdata/{index}",
			Method:          GET,
			Pattern:         "/indexdata/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetIndexData,
		},
	}
}

func getTransactionRouting() []*Road {
	return []*Road{
		{
			Name:            OPTIONS + " /transaction",
			Method:          OPTIONS,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /transaction",
			Method:          POST,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.AddTransaction,
		},
		{
			Name:            GET + " /transaction",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.GetTransaction,
		},
		{
			Name:            DELETE + " /transaction",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveAllTX,
		},
	}
}
