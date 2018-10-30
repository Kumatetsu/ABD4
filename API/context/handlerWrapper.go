/*
 * File: handlerWrapper.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 4:45:50 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import "net/http"

// CustomHandler used as app Handler struct
// It embed an AppContext variable
type CustomHandler func(*AppContext, http.ResponseWriter, *http.Request)

// HandlerWrapper implement http.Handler
// Can transmit an AppContext and a CustomHandler
// Use these struct to share AppContext throught an API
type HandlerWrapper struct {
	Ctx *AppContext
	H   CustomHandler
}

// ServeHTTP implement http.ServeHTTP method, handlerWrapper become a http.Handler
// ServeHTTP call the last registered HandlerWrapper in HandlerWrapper.H property
func (hw HandlerWrapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	hw.H(hw.Ctx, w, r)
}
