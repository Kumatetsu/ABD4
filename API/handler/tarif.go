/*
 * File: tarif.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 15th October 2018 7:42:41 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Friday, 2nd November 2018 7:53:28 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package handler

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/model"
	"ABD4/API/service"
	"ABD4/API/utils"
	"fmt"
	"net/http"
	"time"
)

// GetTarifs return all tarif in database
func GetTarif(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	var toSerialize []iserial.Serializable
	ctx.Log.Info.Printf("%s %s ", utils.Use().GetStack(GetTarif), "Getting Tarifs")
	tx, err := ctx.TarifManager.FindAll()
	if err != nil {
		msg := fmt.Sprintf("%s FindAll failed", utils.Use().GetStack(GetTarif))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	for _, t := range tx {
		toSerialize = append(toSerialize, t)
	}
	ctx.Rw.SendArraySerializable(ctx, w, http.StatusOK, toSerialize, "", "")
}

// AddTarif add a tarif, expect body { "description": string, "prix": float64 }
func AddTarif(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	tarif := &model.Tarif{}
	err := tarif.UnmarshalFromRequest(r)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Decode request data failed", err.Error())
		return
	}
	tarif, err = ctx.TarifManager.Create(tarif)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "Insert tarif in mongo failed", err.Error())
		return
	}
	// si on utilise elastic search on index le nouveau tarif
	if ctx.Opts.GetEmbedES() {
		err = ctx.IndexData(tarif.ToES(), context.TARIFS, context.TARIF)
		if err != nil {
			msg := fmt.Sprintf("%s failed to index tarif in elasticsearch", utils.Use().GetStack(AddTarif))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		ctx.Log.Info.Printf("%s successfull indexation of new %s", utils.Use().GetStack(AddTarif), context.TARIF)
	}
	ctx.Rw.SendSerializable(ctx, w, http.StatusCreated, tarif, "", "")
	return
}

// RemoveAllTarif destoy everything
func RemoveAllTarif(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	deleted, err := ctx.TarifManager.RemoveAll()
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s from %s", utils.Use().GetStack(RemoveAllTarif), context.TARIFS, ctx.Opts.GetDatabaseType())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	err = ctx.RemoveIndex(context.TARIFS)
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s index", utils.Use().GetStack(RemoveAllTarif), context.TARIFS)
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveAllTarif), deleted, context.TARIFS)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}

// LoadTarifsFromFile add tarifs from /tarifs/tarifs.json into database
// it set date to current datetime, so identical tarifs will be override
func LoadTarifsFromFile(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	// On charge les tarifs renseignés dans le fichier exe_folder/tarifs/tarifs.json
	tarifsFromFile, err := service.Tarif(ctx).LoadFromFile()
	if err != nil {
		msg := fmt.Sprintf("%s failed to load tarif from file", utils.Use().GetStack(LoadTarifsFromFile))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	for _, tarif := range tarifsFromFile {
		t := &model.Tarif{
			Description: tarif.Description,
			Prix:        tarif.Prix,
			Date:        time.Now().Format(model.TARIF_TIME_FORMAT),
		}
		// le nouveau tarif va être ajouté
		// avec now comme date, ainsi, la description
		// et le prix peuvent être identiques, c'est la date qui fera foi
		_, err = ctx.TarifManager.Create(t)
		if err != nil {
			msg := fmt.Sprintf("%s failed to create tarif: %v", utils.Use().GetStack(LoadTarifsFromFile), t)
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s tarifs successfully updated from tarif file", utils.Use().GetStack(LoadTarifsFromFile))
	ctx.Rw.SendString(ctx, w, http.StatusCreated, msg, "", "")
}
