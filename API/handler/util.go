/*
 * File: util.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 9:36:54 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 15th October 2018 9:38:00 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/utils"
	"net/http"
)

func Option(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	ctx.Log.Info.Printf("%s %s", utils.Use().GetStack(Option), "Option handler")
	return
}
