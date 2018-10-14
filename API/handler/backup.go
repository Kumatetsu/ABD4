/*
 * File: backup.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 9:31:29 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Thursday, 11th October 2018 7:14:09 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/database/boltdatabase"
	"ABD4/API/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/boltdb/bolt"
)

// BackupBoltDatabase CustomHandler to allow to download the database content
func BackupBoltDatabase(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	if _, ok := ctx.UserManager.GetDB().(boltdatabase.DBManager); !ok {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(BackupBoltDatabase), "The embedded database don't allow backup")
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Backup data failed", msg)
		return
	}
	dbManager := ctx.UserManager.GetDB().(*boltdatabase.DBManager)
	if err := dbManager.OpenDB(); err != nil {
		return
	}
	defer dbManager.CloseDB()
	err := dbManager.GetDB().DB.View(func(tx *bolt.Tx) error {
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Disposition", `attachment; filename="user.dat"`)
		w.Header().Set("Content-Length", strconv.Itoa(int(tx.Size())))
		_, err := tx.WriteTo(w)
		return err
	})
	if err != nil {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(BackupBoltDatabase), err.Error())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Backup data failed", msg)
	}
}
