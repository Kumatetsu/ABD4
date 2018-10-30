/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 1:34:22 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package main

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"ABD4/API/server"
	"ABD4/API/utils"
	"net/http"
)

type App struct {
	Ctx *context.AppContext
}

var (
	PROJECT = "abd4"
	MONGO   = "mongo"
	BOLT    = "bolt"
)

func testUser(a *App) {
	_, err := a.Ctx.UserManager.Create(&model.User{
		Name:     "test",
		Email:    "test",
		Password: "test",
	})
	usr, err := a.Ctx.UserManager.FindOneBy(map[string]string{
		"name": "test",
	})
	if err != nil {
		a.Ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(a.Initialize), err.Error())
	}
	a.Ctx.Log.Info.Printf("%s user created in Initialize: %v", utils.Use().GetStack(a.Initialize), usr)
	a.Ctx.Log.Info.Printf("%s try to get createdAt: %s and Updated at: %s", utils.Use().GetStack(a.Initialize), usr.GetCreatedAt(), usr.GetUpdatedAt())
}

// Initialize define ctx variable
// In ctx, we regroup database access, elasticsearch tools, logs
// Ctx will be used in all handlers and middlewares
func (a *App) Initialize(opts *server.Option) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.Initialize))
	if a.Ctx.Opts.GetDatabaseType() == MONGO {
		// Cleanup des data si on veut...
		// defer a.Ctx.TransactionManager.RemoveAll()
		// defer a.Ctx.UserManager.RemoveAll()
		defer a.Ctx.Mongo.Close()
	}
	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	// process wait on this call
	return http.ListenAndServe(opts.GetAddress(), nil)
}

// InitializeAndWaitForSignal launch the server, making the match between server/model/handler and context/database/logger
// this function need to be used with a channel assuring that Ctx is instanciated (see main_test.go)
func (a *App) InitializeAndWaitForSignal(opts *server.Option, signal chan bool) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.InitializeAndWaitForSignal))
	// define dao access (database/manager package)
	if a.Ctx.Opts.GetDatabaseType() == MONGO {
		defer a.Ctx.Mongo.Close()
	}
	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	signal <- true
	return http.ListenAndServe(opts.GetAddress(), nil)
}
