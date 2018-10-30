/*
 * File: angular.go
 * Project: ABD4/VMD Escape Game
 * File Created: Tuesday, 30th October 2018 12:07:30 am
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:13:03 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"fmt"
	"net/http"
	"path/filepath"
)

func Angular(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serving Angular routing...")
	http.ServeFile(w, r, filepath.Join(ctx.Opts.GetWebDir(), "index.html"))
	return
}
