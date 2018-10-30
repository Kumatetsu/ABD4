/*
 * File: context.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 5:30:40 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	boltM "ABD4/API/database/boltdatabase/manager"
	"ABD4/API/database/mongo"
	mongoM "ABD4/API/database/mongo/manager"
	"ABD4/API/elasticsearch"
	"ABD4/API/iserial"
	"ABD4/API/logger"
	"ABD4/API/utils"
	goctx "context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var (
	PROJECT = "abd4"
	MONGO   = "mongo"
	BOLT    = "bolt"
	USER    = "user"
	USERS   = "users"
	TX      = "transaction"
	TXs     = "transactions"
	SECRET  = "==+VDMEG@ABD4"
	INDEXES = map[string]string{USER: USERS, TX: TXs}
)

// logState log final parameters at the launch of API
func (ctx *AppContext) logState() {
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
}

// setDao define database access relying on GetDatabaseType return
// accept mongo or bolt database
func (ctx *AppContext) setDAO(kind string) error {
	switch kind {
	case MONGO:
		mongoAddr := ctx.Opts.GetMongoIP() + ":" + ctx.Opts.GetMongoPort()
		mongo, err := mongo.GetMongo(mongoAddr)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(ctx.setDAO), err.Error())
		}
		ctx.UserManager = &mongoM.UserManager{}
		err = ctx.UserManager.Init(map[string]string{
			"dbName": PROJECT,
			"entity": USER,
		})
		err = ctx.UserManager.SetDB(mongo)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(ctx.setDAO), err.Error())
		}
		ctx.TransactionManager = &mongoM.TransactionManager{}
		err = ctx.TransactionManager.Init(map[string]string{
			"dbName": PROJECT,
			"entity": TXs,
		})
		err = ctx.TransactionManager.SetDB(mongo)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(ctx.setDAO), err.Error())
		}
	case BOLT:
		userManager := &boltM.UserManager{}
		userManager.Init(map[string]string{
			"name":     PROJECT,
			"fullpath": ctx.DataPath,
			"secret":   SECRET,
		})
		ctx.UserManager = userManager
	}
	return nil
}

// indexData take data actually in database and send it to elastic correspondant index
// there is nothing to prevent the reindexation of data
// its easy to fix a bad indexation using rmindex/index/reindex options
func (ctx *AppContext) indexData() error {
	if ctx.UserManager.GetDBName() == "" || ctx.TransactionManager.GetDBName() == "" {
		return fmt.Errorf("%s At list one data manager is missing", utils.Use().GetStack(ctx.indexData))
	}
	users, err := ctx.UserManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	var serialUsers []iserial.Serializable
	for _, u := range users {
		serialUsers = append(serialUsers, u.ToES())
	}
	err = ctx.IndexArrayData(serialUsers, USERS, USER)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	tx, err := ctx.TransactionManager.FindAll()
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	var serialTx []iserial.Serializable
	for _, transaction := range tx {
		serialTx = append(serialTx, transaction.ToES())
	}
	err = ctx.IndexArrayData(serialTx, TXs, TX)
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	return nil
}

// embedElasticSearch is the entry point to settle an elasticsearch context linked with this api
// User and Transaction entities are indexed
// This can be deactivated using flag embedES=false
func (ctx *AppContext) embedElasticSearch() {
	var err error

	ctx.ElasticClient, err = elasticsearch.Instanciate(ctx.Opts.GetEs())
	if err != nil {
		ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
	}
	// if flag rmindex is passed to .exe, indexes 'users' and 'transactions' are removed
	if ctx.Opts.GetRmindex() {
		err = elasticsearch.RemoveIndex(ctx.ElasticClient, USERS)
		if err != nil {
			ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
		}
		err = elasticsearch.RemoveIndex(ctx.ElasticClient, TXs)
		if err != nil {
			ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
		}
	}
	// if index or reindex flag is present
	// index create indexes if don't exist
	// reindex remove and create indexes
	// if reindex, data are pulled from database and send to elasticsearch
	// there is nothing to prevent duplicate in elastic
	if ctx.Opts.GetIndex() || ctx.Opts.GetReindex() {
		err = elasticsearch.CreateIndexation(ctx.ElasticClient, ctx.Opts.GetReindex())
		if err != nil {
			ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
		}
		if ctx.Opts.GetReindex() {
			err = ctx.indexData()
			if err != nil {
				ctx.Log.Error.Fatalf("%s %s", utils.Use().GetStack(ctx.Instanciate), err.Error())
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
	// open a file for log, for now, just one file is defined
	// but we can move it easy to three log files (debug, info, error)
	// for now, debug mode is useless
	logFile, err := os.OpenFile(
		filepath.Join(opts.GetLogpath(), PROJECT+".log"),
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

	// define dao access (database/manager package)
	err = ctx.setDAO(opts.GetDatabaseType())
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

// IndexData send a iserial.Serializable entity to elasticsearch
// index and t are elasticsearch target index and entity type:
// 'transactions' and 'transaction' for exemple
func (ctx *AppContext) IndexData(i iserial.Serializable, index, t string) error {
	op, err := ctx.ElasticClient.Index().
		Index(index).
		Type(t).
		BodyJson(i).
		Refresh("true").
		Do(goctx.Background())
	if err != nil {
		ctx.Log.Error.Printf("%s %s", utils.Use().GetStack(ctx.IndexData), err.Error())
		return err
	}
	ctx.Log.Info.Printf("%s insert unit %s to index %s, entity type: %s", utils.Use().GetStack(ctx.IndexData), op.Id, op.Index, op.Type)
	return nil
}

// IndexArrayData, call IndexData on an array of iserial.Serializable
func (ctx *AppContext) IndexArrayData(serializables []iserial.Serializable, index, t string) error {
	var err error
	for _, i := range serializables {
		err = ctx.IndexData(i, index, t)
		if err != nil {
			break
		}
	}
	return err
}

// RemoveIndex destroy completely an index
func (ctx *AppContext) RemoveIndex(index string) error {
	return elasticsearch.RemoveIndex(ctx.ElasticClient, index)
}
