/*
 * File: userHandler.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 4:35:32 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetUsers handler for Get /users
func GetUsers(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	users, err := ctx.UserManager.FindAll()
	if err != nil {
		msg := fmt.Sprintf("%s Seek users failed: %s", utils.Use().GetStack(GetUsers), err.Error())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, "")
		return
	}
	usersJSON, err := json.Marshal(users)
	if err != nil {
		msg := fmt.Sprintf("%s marshal failed: %s", utils.Use().GetStack(GetUsers), err.Error())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "encoding users failed", msg)
		return
	}
	response := ctx.Rw.NewResponse(200, string(usersJSON), "", "")
	response.SendItSelf(ctx, w)
	return
}
