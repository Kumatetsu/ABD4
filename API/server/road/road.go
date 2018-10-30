/*
 * File: const.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:50:11 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:52:50 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import (
	"ABD4/API/context"
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

// Road define what a router entry must embed
type Road struct {
	Name            string
	Method          string
	Pattern         string
	StatusProtected bool
	HandlerFunc     context.CustomHandler
}

// appendTo log the happened road and add it to the router section with a prepared handler
func (r *Road) AppendTo(ctx *context.AppContext, preparedHandler *context.HandlerWrapper, router *mux.Router) {
	ctx.Log.Info.Printf("%s __ Load... __ %v", utils.Use().GetStack(r.AppendTo), r.Name)
	router.StrictSlash(true).Handle(r.Pattern, preparedHandler).Methods(r.Method).Name(r.Name)
}
