/*
 * File: time.go
 * Project: ABD4/VMD Escape Game
 * File Created: Saturday, 3rd November 2018 11:45:59 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Saturday, 3rd November 2018 11:49:11 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package middleware

import (
	"ABD4/API/context"
	"ABD4/API/utils"
	"net/http"
	"time"
)

func SetTime(ctx *context.AppContext, process *context.HandlerWrapper) *context.HandlerWrapper {
	return &context.HandlerWrapper{
		Ctx: ctx,
		H: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
			ctx.Time = time.Now()
			ctx.Log.Info.Printf("%s %s", utils.Use().GetStack(SetTime), ctx.Time.String())
			process.ServeHTTP(w, r)
		},
	}
}
