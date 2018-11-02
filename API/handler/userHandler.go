/*
 * File: userHandler.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Wednesday, 31st October 2018 9:34:07 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/utils"
	"fmt"
	"net/http"
)

// GetUsers handler for Get /users
func GetUsers(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	var toSerialize []iserial.Serializable
	users, err := ctx.UserManager.FindAll()
	if err != nil {
		msg := fmt.Sprintf("%s Seek users failed: %s", utils.Use().GetStack(GetUsers), err.Error())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, "")
		return
	}
	for _, u := range users {
		toSerialize = append(toSerialize, u)
	}
	ctx.Rw.SendArraySerializable(ctx, w, http.StatusOK, toSerialize, "", "")
	return
}

// RemoveAllUsers handler for DELETE /users
func RemoveAllUsers(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	deleted, err := ctx.UserManager.RemoveAll()
	if err != nil {
		msg := fmt.Sprintf("%s Removing users failed", utils.Use().GetStack(RemoveAllUsers))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	err = ctx.RemoveIndex(context.USERS)
	if err != nil {
		msg := fmt.Sprintf("%s Failed to remove %s index", utils.Use().GetStack(RemoveAllUsers), context.USERS)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
	}
	msg := fmt.Sprintf("%s %d successfully removed", utils.Use().GetStack(RemoveAllUsers), deleted)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
	return
}
