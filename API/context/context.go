/*
 * File: context.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 1:12:19 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/logger"
	"ABD4/API/utils"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	appName = "abd4"
	SECRET  = "==+VDMEG@ABD4"
)

// Instanciate the global ctx variable
// AppContext know: manager, iserial, utils and logger packages
// AppContext don't need to know: server, model and handler packages
// AppContext rely on three interfaces implementing usefull function for API
// ISerial to serializable model object (model abstraction)
// ISessionUser to typical api user (model abstraction)
// IResponseWriter working with ISerial, it should provide basic shorthand
// to write responses logging and managing errors in handling process
// The context allow us to embed usefull data as datapath or ip and port (Opts)
// in handlers implementing CustomHandler and HandlerWrapper.
// It also centralise common functionnalities as database access, logging process
// and response/error formatting
// the app must define how routing will work and must implement some interfaces:
// - IResponseWriter: Response process where serialisation apply
// - ISerial: On each model entity we want to return
// - ISessionUser: On the User model entity
// - IServerOption: On the structure embedding configuration from flags and harcoded values
func (ctx *AppContext) Instanciate(opts IServerOption) *AppContext {
	if opts.GetExeFolder() == "" {
		log.Fatalf("No exe folder defined, %s unable to provide defaults values", appName)
	}
	// check opts content and define default values if required
	if opts.GetLogpath() == "" {
		opts.SetLogpath(filepath.Join(opts.GetExeFolder(), "log/"))
	}
	if opts.GetDatapath() == "" {
		opts.SetDatapath(filepath.Join(opts.GetExeFolder(), "data/"))
	}
	// create path if don't exist
	err := os.MkdirAll(opts.GetLogpath(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.MkdirAll(opts.GetDatapath(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// open a file for log, for now, just one file is defined
	// but we can move it easy to three log files (debug, info, error)
	// for now, debug mode is useless
	logFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), appName+".log"),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// we define an output able to log on file and standart output
	output := io.MultiWriter(logFile, os.Stdout)
	loggers := logger.Instanciate(output, output, output)
	// instanciate the ctx to return
	ctx.Opts = opts
	//SessionUser: &model.User{},
	ctx.Log = loggers
	ctx.Exe = opts.GetExeFolder()
	ctx.Logpath = opts.GetLogpath()
	ctx.DataPath = opts.GetDatapath()
	ctx.Log.Info.Printf("%s RootDir: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Exe)
	ctx.Log.Info.Printf("%s LogPath: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Logpath)
	ctx.Log.Info.Printf("%s Database asked: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetDatabaseType())
	if ctx.Opts.GetDatabaseType() == "mongo" {
		ctx.Log.Info.Printf("%s Mongo server address: %s:%s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetMongoIP(), ctx.Opts.GetMongoPort())
	} else {
		ctx.Log.Info.Printf("%s Bolt datapath: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetDatapath())
	}
	ctx.Log.Info.Printf("%s IP: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetIP())
	ctx.Log.Info.Printf("%s Port: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetPort())
	ctx.Log.Info.Printf("%s Address: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetAddress())

	return ctx
}
