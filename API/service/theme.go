/*
 * File: theme.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 1st November 2018 6:46:06 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Saturday, 3rd November 2018 3:12:51 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package service

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type themeService struct {
	t      *model.Theme
	themes []*model.Theme
	ctx    *context.AppContext
}

var themeByGame = map[string][]string{
	"Impot sur le revenu":                  []string{"Braquage", "Stress"},
	"Greve de la SNCF":                     []string{"Rapidité", "Mythologique"},
	"Interminable attente chez le medecin": []string{"Stratégie", "Psychologique"},
	"Soutenance finale":                    []string{"Stress", "Rapidité"},
	"Mon compte en banque en fin du mois":  []string{"Mythologique", "Braquage"},
	"Mariage sans alcool":                  []string{"Santé", "Amour"},
	"Diner de famille insoutenable":        []string{"Psychologique", "Stratégie"},
	"Plus de PQ dans les toilettes":        []string{"Horreur", "Santé"},
	"En plein dans la Friendzone":          []string{"Amour", "Horreur"},
}

//Theme instanciate a themeService
func Theme(ctx *context.AppContext) *themeService {
	t := &model.Theme{}
	themes := []*model.Theme{}
	return &themeService{
		// on pousse le format depuis le model
		// c'est le model qui dicte sa conduite au service
		// le model n'a pas connaissance du service
		t:      t,
		themes: themes,
		ctx:    ctx,
	}
}

// Load fullfill Tarif.tarifs with the content of tarifs/tarifs.json file
func (ts *themeService) LoadFromFile() ([]*model.Theme, error) {
	buff := []string{}
	themeFilePath := filepath.Join(ts.ctx.Opts.GetExeFolder(), "themes/themes.json")
	fmt.Printf("%s themeFilePath : %s", utils.Use().GetStack(ts.LoadFromFile), themeFilePath)
	file, err := ioutil.ReadFile(themeFilePath)
	if err != nil {
		return nil, fmt.Errorf("%s %s", utils.Use().GetStack(ts.LoadFromFile), err.Error())
	}
	themes := []*model.Theme{}
	err = json.Unmarshal(file, &buff)
	if err != nil {
		return nil, err
	}
	for _, theme := range buff {
		themes = append(themes, &model.Theme{Theme: theme})
	}
	return themes, nil
}

// DefineTheme Define a main and secondary theme for a transaction
func (ts *themeService) DefineTheme(t *model.Transaction) {
	t.Game.Themes = themeByGame[t.Game.Nom]
}
