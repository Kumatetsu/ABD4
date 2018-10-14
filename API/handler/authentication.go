/*
 * File: authentication.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 4:36:26 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Login handler, allow to get a token if authentication goes well
// Must receive: {"email": string, "password": string}
func Login(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := user.UnmarshalFromRequest(r)
	if nil != err {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "can't interpret users data", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
		return
	}
	// if no user is found, user is erased
	user, err = ctx.UserManager.FindOneBy(map[string]string{
		"email":    user.Email,
		"password": user.Password,
	})
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "retrieving user failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
		return
	}
	ctx.Log.Info.Printf("%s %s want to login! %v", utils.Use().GetStack(Login), user.Email, user)
	if user.ID != "" {
		ctx.Log.Info.Printf("%s %s is in login process!", utils.Use().GetStack(Login), user.Email)
		claim := &model.Claim{}
		if user.Claim != "" {
			// L'utilisateur est connu des services et a des claims
			err = json.Unmarshal([]byte(user.Claim), claim)
			if err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "decoding existing informations failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
				return
			}
		} else {
			if err := claim.NewFromUser(user); err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "generate claim failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
				return
			}
		}
		token, err := claim.GetToken()
		if err != nil {
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "generate token failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Login), err.Error()))
			return
		}
		ctx.SessionUser = user
		ctx.Log.Info.Printf("%s login process ok, token: %s", utils.Use().GetStack(Login), token)
		response := ctx.Rw.NewResponse(200, token, "", "")
		response.SendItSelf(ctx, w)
		return
	}
	msg := fmt.Sprintf("%s invalid credentials", utils.Use().GetStack(Login))
	ctx.Rw.SendError(ctx, w, 401, msg, "")
	return
}

// Register manage the inscription of a user
// Must receive: {"email": string, "password": string}
func Register(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	user := &model.User{}
	err := user.UnmarshalFromRequest(r)
	if nil != err {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "can't interpret user data", fmt.Sprintf("%s %s", utils.Use().GetStack(Register), err.Error()))
		return
	}
	user, err = ctx.UserManager.Create(user)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "error saving user", fmt.Sprintf("%s %s", utils.Use().GetStack(Register), err.Error()))
		return
	}
	ctx.Rw.Send(ctx, w, 200, user, "", "")
	return
}
