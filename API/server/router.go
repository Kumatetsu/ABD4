/*
 * File: router.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:13:26 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import (
	"ABD4/API/context"
	"ABD4/API/server/middleware"
	"ABD4/API/server/road"
	"ABD4/API/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// RoadGetter type for router definition
type RoadGetter func() []*road.Road

// Routing function is the entry point to build the API routing
func Routing(ctx *context.AppContext) *mux.Router {
	// the base router is a mux entity
	router := &mux.Router{}
	// We set the Response object memory space
	ctx.Rw = &Response{}
	for url, roadGetter := range map[string]RoadGetter{
		// exemple for localhost:8000/user/* set of road:
		"/user":        road.GetUserRouting,
		"/auth":        road.GetAuthRouting,
		"/backup":      road.GetBackupRouting,
		"/elastic":     road.GetElasticRouting,
		"/transaction": road.GetTransactionRouting,
		"/tarif":       road.GetTarifRouting,
		"/theme":       road.GetThemeRouting,
	} {
		routing := router.PathPrefix(url).Subrouter()
		for _, road := range roadGetter() {
			// call prepareHandler on each context.HandlerWrapper
			road.AppendTo(ctx, ApplyMiddlewares(ctx, road), routing)
		}
	}
	AngularRouting(ctx, router)
	return router
}

// applyMiddlewares prepare the stack of methods called for each road
// last registered is first to be called
func ApplyMiddlewares(ctx *context.AppContext, r *road.Road) *context.HandlerWrapper {
	wrapper := &context.HandlerWrapper{
		Ctx: ctx,
		H:   r.HandlerFunc,
	}
	wrapper = middleware.SetHeaders(ctx, wrapper)
	wrapper = middleware.Authenticate(ctx, wrapper, r.StatusProtected)
	// logger must be the last to be the first...
	wrapper = middleware.Logger(ctx, wrapper, r.Name)
	return wrapper
}

func AngularRouting(ctx *context.AppContext, router *mux.Router) {
	home := road.GetHome()
	routes := road.GetWebAppRouting(ctx)
	for _, r := range routes {
		ctx.Log.Info.Printf("%s __ Load... __ %v", utils.Use().GetStack(AngularRouting), r.Name)
		router.PathPrefix(r.Pattern).Handler(http.StripPrefix(r.Pattern, ApplyMiddlewares(ctx, home)))
	}
	ctx.Log.Info.Printf("%s __ Load... __ %v", utils.Use().GetStack(AngularRouting), home.Name)
	router.PathPrefix(home.Pattern).Handler(http.FileServer(http.Dir(ctx.Opts.GetWebDir())))
}
