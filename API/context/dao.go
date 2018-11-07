/*
 * File: dao.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 4th November 2018 10:56:57 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 6:28:38 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	boltM "ABD4/API/database/boltdatabase/manager"
	mongoDB "ABD4/API/database/mongo"
	mongoM "ABD4/API/database/mongo/manager"
	"ABD4/API/utils"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// setDao define database access relying on GetDatabaseType return
// accept mongo or bolt database
func (ctx *AppContext) setDAO(kind string) error {
	var err error

	switch kind {
	case MONGO:
		var mongo *mgo.Session
		mongoAddr := ctx.Opts.GetMongoIP() + ":" + ctx.Opts.GetMongoPort()

		if ctx.Opts.GetMongoReplicatSet() {
			ctx.Log.Info.Printf("%s Set mongo with replicat set", utils.Use().GetStack(ctx.setDAO))
			mongoHosts := []string{
				mongoAddr,
				ctx.Opts.GetReplicatIP() + ":" + ctx.Opts.GetReplicatPort(),
			}
			for _, host := range mongoHosts {
				ctx.Log.Info.Printf("%s %s", utils.Use().GetStack(ctx.setDAO), host)
			}
			mongo, err = mongoDB.GetMongoReplicatSet(mongoHosts, PROJECT)
			ctx.Log.Info.Printf("%s mongo instanciation done", utils.Use().GetStack(ctx.setDAO))
		} else {
			ctx.Log.Info.Printf("%s Set mongo with direct dial", utils.Use().GetStack(ctx.setDAO))
			mongo, err = mongoDB.GetMongo(mongoAddr)
		}
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
			"entity": TX,
		})
		err = ctx.TransactionManager.SetDB(mongo)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(ctx.setDAO), err.Error())
		}
		ctx.TarifManager = &mongoM.TarifManager{}
		err = ctx.TarifManager.Init(map[string]string{
			"dbName": PROJECT,
			"entity": TARIF,
		})
		err = ctx.TarifManager.SetDB(mongo)
		if err != nil {
			return fmt.Errorf("%s %s", utils.Use().GetStack(ctx.setDAO), err.Error())
		}
		ctx.ThemeManager = &mongoM.ThemeManager{}
		err = ctx.ThemeManager.Init(map[string]string{
			"dbName": PROJECT,
			"entity": THEME,
		})
		err = ctx.ThemeManager.SetDB(mongo)
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
