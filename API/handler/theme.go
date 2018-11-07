/*
 * File: theme.go
 * Project: ABD4/VMD Escape Game
 * File Created: Tuesday, 30th October 2018 9:10:01 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 6:14:34 am
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

	"github.com/gorilla/mux"
)

// GetTheme return all theme in database
func GetTheme(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	var toSerialize []iserial.Serializable
	ctx.Log.Info.Printf("%s %s ", utils.Use().GetStack(GetTheme), "Getting Themes")
	tx, err := ctx.ThemeManager.FindAll()
	if err != nil {
		msg := fmt.Sprintf("%s FindAll failed", utils.Use().GetStack(GetTheme))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	for _, t := range tx {
		toSerialize = append(toSerialize, t)
	}
	ctx.Rw.SendArraySerializable(ctx, w, http.StatusOK, toSerialize, "", "")
}

// AddTheme add a theme
func AddTheme(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	theme := &model.Theme{}
	err := theme.UnmarshalFromRequest(r)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "Decode request data failed", err.Error())
		return
	}
	theme, err = ctx.ThemeManager.Create(theme)
	if err != nil {
		ctx.Rw.SendError(ctx, w, http.StatusBadRequest, "Insert theme in mongo failed", err.Error())
		return
	}
	// si on utilise elastic search on index la nouvelle theme
	if ctx.Opts.GetEmbedES() {
		err = ctx.IndexData(theme.ToES(), context.THEMES, context.THEME)
		if err != nil {
			msg := fmt.Sprintf("%s failed to index theme in elasticsearch", utils.Use().GetStack(AddTheme))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		ctx.Log.Info.Printf("%s successfull indexation of new %s", utils.Use().GetStack(AddTheme), context.THEME)
	}
	ctx.Rw.SendSerializable(ctx, w, http.StatusCreated, theme, "", "")
	return
}

// Remove a theme by his description
func RemoveThemeByName(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	theme, ok := mux.Vars(r)["theme"]
	if !ok {
		err := fmt.Errorf("%s no theme parameter in url", utils.Use().GetStack(RemoveThemeByName))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "parameter is missing", err.Error())
		return
	}
	deleted, err := ctx.ThemeManager.RemoveBy(map[string]string{"Theme": theme})
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s from %s", utils.Use().GetStack(RemoveThemeByName), context.THEMES, ctx.Opts.GetDatabaseType())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	if ctx.Opts.GetEmbedES() {
		err = ctx.RemoveIndex(context.THEMES)
		if err != nil {
			msg := fmt.Sprintf("%s failed to remove %s index", utils.Use().GetStack(RemoveThemeByName), context.THEMES)
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveThemeByName), deleted, context.THEME)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}

// Remove a theme by his id
func RemoveThemeByID(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	theme, ok := mux.Vars(r)["id"]
	if !ok {
		err := fmt.Errorf("%s no theme parameter in url", utils.Use().GetStack(RemoveThemeByID))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, "parameter is missing", err.Error())
		return
	}
	deleted, err := ctx.ThemeManager.RemoveBy(map[string]string{"id": theme})
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s from %s", utils.Use().GetStack(RemoveThemeByID), context.THEMES, ctx.Opts.GetDatabaseType())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveThemeByID), deleted, context.THEME)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}

// RemoveAllTHEME destoy everything
func RemoveAllThemes(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	deleted, err := ctx.ThemeManager.RemoveAll()
	if err != nil {
		msg := fmt.Sprintf("%s failed to remove %s from %s", utils.Use().GetStack(RemoveAllThemes), context.THEMES, ctx.Opts.GetDatabaseType())
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	if ctx.Opts.GetEmbedES() {
		err = ctx.RemoveIndex(context.THEMES)
		if err != nil {
			msg := fmt.Sprintf("%s failed to remove %s index", utils.Use().GetStack(RemoveAllThemes), context.THEMES)
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s %d %s successfully deleted", utils.Use().GetStack(RemoveAllThemes), deleted, context.THEME)
	ctx.Rw.SendString(ctx, w, http.StatusAccepted, msg, "", "")
}

// LoadTarifsFromFile add themes from /themes/themes.json into database
// identical themes will be duplicated
func LoadThemesFromFile(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
	var toSerialize []iserial.Serializable
	// On charge les themes renseignés dans le fichier exe_folder/themes/theme.json
	themesFromFile, err := service.Theme(ctx).LoadFromFile()
	if err != nil {
		msg := fmt.Sprintf("%s failed to load theme from file", utils.Use().GetStack(LoadThemesFromFile))
		ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
		return
	}
	for _, theme := range themesFromFile {
		t := &model.Theme{
			Theme: theme.Theme,
		}
		// le nouveau tarif va être ajouté
		// avec now comme date, ainsi, la description
		// et le prix peuvent être identiques, c'est la date qui fera foi
		_, err = ctx.ThemeManager.Create(t)
		if err != nil {
			msg := fmt.Sprintf("%s failed to create theme: %v", utils.Use().GetStack(LoadThemesFromFile), t)
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
		if ctx.Opts.GetEmbedES() {
			toSerialize = append(toSerialize, t.ToES())
		}
	}
	if ctx.Opts.GetEmbedES() {
		err = ctx.IndexArrayData(toSerialize, context.THEMES, context.THEME)
		if err != nil {
			msg := fmt.Sprintf("%s failed to index themes", utils.Use().GetStack(LoadThemesFromFile))
			ctx.Rw.SendError(ctx, w, http.StatusInternalServerError, msg, err.Error())
			return
		}
	}
	msg := fmt.Sprintf("%s themes successfully updated from theme file", utils.Use().GetStack(LoadThemesFromFile))
	ctx.Rw.SendString(ctx, w, http.StatusCreated, msg, "", "")
}
