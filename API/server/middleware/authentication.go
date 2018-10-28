/*
 * File: authentication.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 4:40:15 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package middleware

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Authenticate is a middleware to apply to all roads
// It's supposed to receive a boolean defining if the road
// need to be protected by authentication or not
// If the road need to be protected, the request must contain a 'Basic [token]' authorization header
func Authenticate(ctx *context.AppContext, handler *context.HandlerWrapper, protected bool) *context.HandlerWrapper {
	return &context.HandlerWrapper{
		Ctx: ctx,
		H: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
			user := &model.User{}
			claim := &model.Claim{}
			// check if road is protected
			if !protected {
				handler.ServeHTTP(w, r)
				return
			}
			ctx.Log.Info.Printf("%s authentication required", utils.Use().GetStack(Authenticate))
			// retrieve cookie
			tokenStr, err := parseBasicAuthorization(r)
			if err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusUnauthorized, "bad credentials", fmt.Sprintf("%s %s", utils.Use().GetStack(Authenticate), err.Error()))
				return
			}
			// parse token with claims
			token, err := jwt.ParseWithClaims(tokenStr, claim, parseWithClaims)
			if err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "bad authorization format", fmt.Sprintf("%s %s", utils.Use().GetStack(Authenticate), err.Error()))
				return
			}
			// check token's validity and expiration
			claim, ok := token.Claims.(*model.Claim)
			if !ok {
				ctx.Rw.SendError(ctx, w, http.StatusUnauthorized, "readed token is not valid", fmt.Sprintf("%s", utils.Use().GetStack(Authenticate)))
				return
			}
			expiration := time.Unix(claim.StandardClaims.ExpiresAt, 0)
			if time.Now().After(expiration) {
				ctx.Rw.SendError(ctx, w, http.StatusUnauthorized, "expired token", fmt.Sprintf("%s", utils.Use().GetStack(Authenticate)))
				return
			}
			// retrieving user from claim
			err = json.Unmarshal([]byte(claim.User), &user)
			if err != nil {
				ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "encoding session user failed", fmt.Sprintf("%s %s", utils.Use().GetStack(Authenticate), err.Error()))
				return
			}
			// setting user in handler wrapper to reuse it if we need
			handler.Ctx.SessionUser = user
			handler.ServeHTTP(w, r)
		}}
}

/*
 * private methods
 */

// parseAuthorization decode authorization header content
// Return the second part containing the token.
func parseBasicAuthorization(r *http.Request) (string, error) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		return "", fmt.Errorf("\nAuthentication: No valid Basic authorization token provided")
	}
	return auth[1], nil
}

func parseWithClaims(t *jwt.Token) (interface{}, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("<<<< %s unexpected signing method", utils.Use().GetStack(Authenticate))
	}
	return []byte(model.SECRET), nil
}
