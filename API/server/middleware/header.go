/*
 * File: header.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 2:37:52 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package middleware

import (
	"ABD4/API/context"
	"ABD4/API/utils"
	"net/http"
)

// SetHeaders is implement:
// Cors request authorization
// application/json Content-Type
// see http://stackoverflow.com/questions/12830095/setting-http-headers-in-golang
func SetHeaders(ctx *context.AppContext, process *context.HandlerWrapper) *context.HandlerWrapper {
	return &context.HandlerWrapper{
		Ctx: ctx,
		H: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
			if origin := r.Header.Get("Origin"); origin != "" {
				ctx.Log.Info.Printf("%s Setting headers...", utils.Use().GetStack(SetHeaders))
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers",
					"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
				w.Header().Set("Access-Control-Allow-Credentials", "true")
				w.Header().Set("Content-Type", "application/json")
			}
			if r.Method == "OPTIONS" {
				return
			}
			process.ServeHTTP(w, r)
		}}
}
