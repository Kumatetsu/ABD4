/*
 * File: logger.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 2:37:58 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package middleware

import (
	"ABD4/API/context"
	"net/http"
	"time"
)

// Logger kind of dependancy injection into the routing/server process
func Logger(ctx *context.AppContext, process *context.HandlerWrapper, name string) *context.HandlerWrapper {

	return &context.HandlerWrapper{
		Ctx: ctx,
		H: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ctx.Log.Info.Printf(
				"Method: %s\tURI: %s\tName: %s\tSince: %s\n",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
			process.ServeHTTP(w, r)
		}}
}
