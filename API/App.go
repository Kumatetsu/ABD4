/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 4:44:19 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package main

import (
	"ABD4/API/context"
	boltM "ABD4/API/database/boltdatabase/manager"
	"ABD4/API/database/mongo"
	mongoM "ABD4/API/database/mongo/manager"
	"ABD4/API/model"
	"ABD4/API/server"
	"ABD4/API/utils"
	"fmt"
	"net/http"
)

type App struct {
	Ctx *context.AppContext
}

func testUser(a *App) {
	a.Ctx.UserManager.Create(&model.User{
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

// Initialize launch the server, making the match between server/model/handler and context/database/logger
func (a *App) Initialize(opts *server.Option) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.Initialize))

	// define dao access (database/manager package)
	err := a.setDAO(opts.GetDatabaseType())
	if err != nil {
		a.Ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(a.Initialize), err.Error())
	}
	if a.Ctx.Opts.GetDatabaseType() == "mongo" {
		defer a.Ctx.Mongo.Close()
	}
	// Insert, retrieve and print a user
	// testUser(a)

	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	return http.ListenAndServe(opts.GetAddress(), nil)
}

// InitializeAndWaitForSignal launch the server, making the match between server/model/handler and context/database/logger
// this function need to be used with a channel assuring that Ctx is instanciated (see main_test.go)
func (a *App) InitializeAndWaitForSignal(opts *server.Option, signal chan bool) error {
	a.Ctx = &context.AppContext{}
	a.Ctx.Instanciate(opts)
	a.Ctx.Log.Info.Printf("%s API starting...", utils.Use().GetStack(a.InitializeAndWaitForSignal))
	// define dao access (database/manager package)
	err := a.setDAO(a.Ctx.Opts.GetDatabaseType())
	if err != nil {
		a.Ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(a.InitializeAndWaitForSignal), err.Error())
	}
	router := server.Routing(a.Ctx)
	http.Handle("/", router)
	signal <- true
	return http.ListenAndServe(opts.GetAddress(), nil)
}

func (a *App) setDAO(kind string) error {
	switch kind {
	case "mongo":
		mongoAddr := a.Ctx.Opts.GetMongoIP() + ":" + a.Ctx.Opts.GetMongoPort()
		mongo, err := mongo.GetMongo(mongoAddr)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(a.setDAO), err.Error())
		}
		a.Ctx.UserManager = &mongoM.UserManager{}
		err = a.Ctx.UserManager.Init(map[string]string{
			"dbName": "abd4",
			"entity": "user",
		})
		err = a.Ctx.UserManager.SetDB(mongo)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(a.setDAO), err.Error())
		}
	case "bolt":
		userManager := &boltM.UserManager{}
		userManager.Init(map[string]string{
			"name":     "abd4",
			"fullpath": a.Ctx.DataPath,
			"secret":   context.SECRET,
		})
		a.Ctx.UserManager = userManager
	}
	return nil
}
