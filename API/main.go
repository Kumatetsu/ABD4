/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 2:07:56 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package main

import (
	"ABD4/API/server"
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var app *App

// launchApp set context with opts
// and start the server
func launchApp(opts *server.Option) {
	app = &App{}
	err := app.Initialize(opts)
	if err != nil {
		log.Fatal(err)
	}
}

// main: manage flag and lauch app.Initialize
func main() {
	// flag management
	var ip = flag.String("ip", "127.0.0.1", "define ip address")
	var env = flag.String("env", "dev", "define environnement context, [prod|dev|test]")
	var es = flag.String("es", "127.0.0.1", "define the es server")
	var port = flag.Int("p", 8000, "define port")
	var debug = flag.Bool("d", false, "active debug logging")
	var embedES = flag.Bool("elastic", true, "Active elastic search")
	var index = flag.Bool("index", false, "indexation for ES")
	var reindex = flag.Bool("reindex", false, "indexation with mapping reloading")
	var rmindex = flag.Bool("rmindex", false, "delete the actual mapping")
	var logpath = flag.String("logpath", "", "define log folder path from exe folder")
	var databasetype = flag.String("db", "mongo", "define the database to use: mongo or bolt")
	var datapath = flag.String("datapath", "", "define data folder path from exe folder, for bolt db")
	var mongoIP = flag.String("mip", "127.0.0.1", "define the mongo server ip")
	var mongoPort = flag.Int("mp", 27017, "define the mongo server ip")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	flag.Parse()
	if err != nil {
		log.Fatal(err.Error())
	}
	portStr := strconv.Itoa(*port)
	mongoPortStr := strconv.Itoa(*mongoPort)
	opts := &server.Option{}
	opts.Hydrate(
		*env,          // server environnement
		dir,           // exe folder absolute path
		*ip,           // server ip
		portStr,       // server port
		*logpath,      // log folder path from dir
		*databasetype, // kind of database defined (mongo or bolt)
		*mongoIP,      // mongo server instance address
		mongoPortStr,  // mongo server instane port
		*datapath,     // for bolt database, this is the path for .dat files from dir
		*es,           // elasticsearch server instance address
		*embedES,      // if true we set elastic search, this is default, set to false to deactivate elastic search
		*index,        // true: users and transactions index will be set in elasticsearch
		*reindex,      // true: indexes are removed and create, data are pushed in elastic search
		*rmindex,      // true: destroy indexes
		*debug)        // useles, suppose to define a debug mode
	launchApp(opts)
}
