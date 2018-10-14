/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
<<<<<<< 49c39e848c51bf035cfe2bc69205781907917321
 * Last Modified: Sunday, 14th October 2018 10:01:33 pm
=======
 * Last Modified: Sunday, 14th October 2018 10:01:33 pm
>>>>>>> feat(main/main_test): new database wrokflow
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

func launchApp(opts *server.Option) {
	app = &App{}
	err := app.Initialize(opts)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var ip = flag.String("ip", "0.0.0.0", "define ip address")
	var env = flag.String("env", "dev", "define environnement context, [prod|dev|test]")
	var es = flag.String("es", "127.0.0.1", "define the es server")
	var port = flag.Int("p", 8000, "define port")
	var debug = flag.Bool("d", false, "active debug logging")
	var index = flag.Bool("index", false, "indexation for ES")
	var reindex = flag.Bool("reindex", false, "indexation with mapping reloading")
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
	opts.Hydrate(portStr, *ip, *env, *es, dir, *logpath, *databasetype, *datapath, *mongoIP, mongoPortStr, *index, *reindex, *debug)
	launchApp(opts)
}
