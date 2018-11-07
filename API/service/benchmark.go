/*
 * File: benchmark.go
 * Project: ABD4/VMD Escape Game
 * File Created: Saturday, 3rd November 2018 10:52:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 4:08:59 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package service

import (
	"ABD4/API/context"
	"ABD4/API/utils"
	"net/http"
	"time"
)

// SetTime define the Time property in ctx at now for benchmarking
func SetTime(ctx *context.AppContext, process *context.HandlerWrapper) *context.HandlerWrapper {
	return &context.HandlerWrapper{
		Ctx: ctx,
		H: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
			ctx.Time = time.Now()
			ctx.Log.Info.Printf("%s %s", utils.Use().GetStack(SetTime), ctx.Time.String())
			process.ServeHTTP(w, r)
		}}
}

func Benchmark(ctx *context.AppContext, target string) {
	now := time.Now()
	diff := now.Sub(ctx.Time)
	ctx.Log.Benchmark.Printf("Target: %s in %s ms", target, diff.String())
}
