/*
 * File: context.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 4:38:45 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/logger"
	"ABD4/API/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"gopkg.in/mgo.v2"
	elastic "gopkg.in/olivere/elastic.v5"
)

var (
	PROJECT = "abd4"
	MONGO   = "mongo"
	BOLT    = "bolt"
	USER    = "user"
	USERS   = PROJECT + "-" + "users"
	TX      = "transaction"
	TXs     = PROJECT + "-" + "transactions"
	TARIF   = "tarif"
	TARIFS  = PROJECT + "-" + "tarifs"
	THEME   = "theme"
	THEMES  = PROJECT + "-" + "themes"
	SECRET  = "==+VDMEG@ABD4"
	INDEXES = map[string]string{USER: USERS, TX: TXs, TARIF: TARIFS, THEME: THEMES}
)

// AppContext define globals tools and variable usefull in the API
// It embed the dao's objects (XxxManager *manager.XxxManager),
// a ResponseWriter which offer shorthand to send uniformised Response
type AppContext struct {
	Rw                 IResponseWriter
	SessionUser        ISessionUser
	Opts               IServerOption
	UserManager        IUserManager
	TransactionManager ITransactionManager
	TarifManager       ITarifManager
	ThemeManager       IThemeManager
	Mongo              *mgo.Session
	ElasticClient      *elastic.Client
	Log                *logger.AppLogger
	Time               time.Time
	SavedTime          time.Time
	time               time.Time
	Exe                string
	Logpath            string
	logpath            map[string]string
	DataPath           string
	mutex              *sync.Mutex
	gorout             int
}

// logState log final parameters at the launch of API
func (ctx *AppContext) logState() {
	ctx.Log.Info.Printf("%s %s SERVER PARAMETERS:", utils.Use().GetStack(ctx.Instanciate), PROJECT)
	ctx.Log.Info.Printf("%s IP: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetIP())
	ctx.Log.Info.Printf("%s Port: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetPort())
	ctx.Log.Info.Printf("%s Address: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetAddress())
	ctx.Log.Info.Printf("%s %s FILE SYSTEM:", utils.Use().GetStack(ctx.Instanciate), PROJECT)
	ctx.Log.Info.Printf("%s RootDir: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Exe)
	ctx.Log.Info.Printf("%s LogPaths: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Logpath)
	ctx.Log.Info.Printf("%s\t info: %s", utils.Use().GetStack(ctx.Instanciate), ctx.logpath["info"])
	ctx.Log.Info.Printf("%s\t debug: %s", utils.Use().GetStack(ctx.Instanciate), ctx.logpath["debug"])
	ctx.Log.Info.Printf("%s\t error: %s", utils.Use().GetStack(ctx.Instanciate), ctx.logpath["error"])
	ctx.Log.Info.Printf("%s\t benchmark: %s", utils.Use().GetStack(ctx.Instanciate), ctx.logpath["benchmark"])
	ctx.Log.Info.Printf("%s %s DATABASE CONTEXT:", utils.Use().GetStack(ctx.Instanciate), PROJECT)
	ctx.Log.Info.Printf("%s Database asked: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetDatabaseType())
	if ctx.Opts.GetDatabaseType() == "mongo" {
		ctx.Log.Info.Printf("%s Mongo server address: %s:%s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetMongoIP(), ctx.Opts.GetMongoPort())
		if ctx.Opts.GetMongoReplicatSet() {
			ctx.Log.Info.Printf("%s Mongo replica address: %s:%s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetReplicatIP(), ctx.Opts.GetReplicatPort())
		}
	} else {
		ctx.Log.Info.Printf("%s Bolt datapath: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetDatapath())
	}
	if ctx.Opts.GetEmbedES() {
		ctx.Log.Info.Printf("%s %s ELASTICSEARCH CONTEXT:", utils.Use().GetStack(ctx.Instanciate), PROJECT)
		ctx.Log.Info.Printf("%s ElasticSearch server address: %s", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetEs())
		if ctx.Opts.GetAllowAsync() {
			ctx.Log.Info.Printf("%s %s API allow asynchronous indexation on this instance", utils.Use().GetStack(ctx.Instanciate), PROJECT)
			ctx.Log.Info.Printf("%s %s API allow %d simultaneous goroutines used by server context", utils.Use().GetStack(ctx.Instanciate), PROJECT, ctx.Opts.GetGorout())
			ctx.Log.Info.Printf("%s asynchronous indexation expected batch size: %d", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetBatch())
			ctx.Log.Info.Printf("%s maximum simultaneous pools of batch is hardcoded to 200", utils.Use().GetStack(ctx.Instanciate))
		}
		if ctx.Opts.GetIndex() {
			ctx.Log.Info.Printf("%s You are asking a creation of all indexes in esmapping folder", utils.Use().GetStack(ctx.Instanciate))
		}
		if ctx.Opts.GetRmindex() {
			ctx.Log.Info.Printf("%s You are asking the deletion of all indexes in esmapping folder", utils.Use().GetStack(ctx.Instanciate))
		}
		if ctx.Opts.GetReindex() {
			ctx.Log.Info.Printf("%s You are asking the deletion and creation of all indexes in esmapping folder", utils.Use().GetStack(ctx.Instanciate))
			if ctx.Opts.GetAllowAsync() {
				ctx.Log.Info.Printf("%s data will be pushed asynchronously by batch of %d in a limit of 200 pools", utils.Use().GetStack(ctx.Instanciate), ctx.Opts.GetBatch())

			} else {
				ctx.Log.Info.Printf("%s data will be pushed synchronously, elasticsearch indexation is 100x slower", utils.Use().GetStack(ctx.Instanciate))

			}
		}
	}
}

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
	ctx.mutex = &sync.Mutex{}
	// create required folders and files
	ctx.createFileSystem(opts)
	// define output streams for logger
	info, errorLog, debug, benchmark := ctx.defineLogSystem(opts)
	// instanciate the ctx to return
	ctx.Opts = opts
	//SessionUser: &model.User{},
	ctx.Log = logger.Instanciate(debug, info, errorLog, benchmark)
	ctx.Exe = opts.GetExeFolder()
	ctx.Logpath = opts.GetLogpath()
	ctx.DataPath = opts.GetDatapath()

	// define dao access (database/manager package)
	err := ctx.setDAO(opts.GetDatabaseType())
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}

	//Define elastic serv and index if needed
	if opts.GetEmbedES() {
		ctx.embedElasticSearch()
	}
	ctx.logState()
	return ctx
}

// createFileSystem critical function called before any instanciation
func (ctx *AppContext) createFileSystem(opts IServerOption) {
	if opts.GetExeFolder() == "" {
		log.Fatalf("No exe folder defined, %s unable to provide defaults values", PROJECT)
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
	err = os.MkdirAll(filepath.Join(opts.GetExeFolder(), "benchmark"), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

// defineLogSystem critical function called before any instanciation
// info, err, debug, benchmark
func (ctx *AppContext) defineLogSystem(opts IServerOption) (io.Writer, io.Writer, io.Writer, io.Writer) {
	// open files for log and benchmarking
	ctx.logpath = make(map[string]string)
	ctx.logpath["info"] = PROJECT + "_info.log"
	ctx.logpath["debug"] = PROJECT + "_debug.log"
	ctx.logpath["error"] = PROJECT + "_error.log"
	ctx.logpath["benchmark"] = PROJECT + "_benchmark.log"
	infoFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), ctx.logpath["info"]),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	debugFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), ctx.logpath["debug"]),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	errorFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), ctx.logpath["error"]),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	benchmarkFile, err := os.OpenFile(
		filepath.Join(opts.GetExeFolder(), "benchmark", ctx.logpath["benchmark"]),
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	// we define an output able to log on file and standart output
	info := io.MultiWriter(infoFile, os.Stdout)
	errOut := io.MultiWriter(errorFile, infoFile, os.Stdout)
	debug := io.MultiWriter(os.Stdout)
	if opts.GetDebug() {
		debug = io.MultiWriter(debugFile, infoFile, os.Stdout)
	}
	benchmark := io.MultiWriter(benchmarkFile, os.Stdout)
	return info, errOut, debug, benchmark
}

func (ctx *AppContext) benchmarkStart() {
	ctx.time = time.Now()
}

func (ctx *AppContext) benchmarkStop(target string) {
	now := time.Now()
	diff := now.Sub(ctx.Time)
	ctx.Log.Benchmark.Printf("Target: %s in %s", target, diff.String())
}
