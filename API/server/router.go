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
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Routing function is the entry point to build the API routing
func Routing(ctx *context.AppContext) *mux.Router {
	// the base router is a mux entity
	router := &mux.Router{}
	// We set the Response object memory space
	ctx.Rw = &Response{}
	appendRootRoad(ctx, router)
	appendRoadKit(ctx, router)
	return router
}

// prepareHandler prepare the stack of methods called for each road
// last registered is first to be called
func prepareHandler(ctx *context.AppContext, r *Road) *context.HandlerWrapper {
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

// separate function to append a handler on localhost/
func appendRootRoad(ctx *context.AppContext, router *mux.Router) {
	rootRoads := []*Road{
		&Road{
			Name:    "/",
			Method:  GET,
			Pattern: "/",
			HandlerFunc: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
				// w.Write take a []byte, fmt.Sprintf return a string
				// we cast with []byte(string)
				w.Write([]byte(fmt.Sprintf("Root road: %s on %s", GET, ctx.Opts.GetAddress())))
			},
		},
	}
	for _, road := range rootRoads {
		road.appendTo(ctx, prepareHandler(ctx, road), router)
	}
}

// appendRoadKit use GetRoadKit function to build the routing
// with sections defined in road.go
func appendRoadKit(ctx *context.AppContext, router *mux.Router) {
	for url, roadGetter := range GetRoadKit() {
		routing := router.PathPrefix(url).Subrouter()
		for _, road := range roadGetter() {
			// call prepareHandler on each context.HandlerWrapper
			road.appendTo(ctx, prepareHandler(ctx, road), routing)
		}
	}
}
