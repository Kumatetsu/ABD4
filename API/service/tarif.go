/*
 * File: tarif.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 9:47:02 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:20:18 am
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
	"time"
)

type tarifService struct {
	t          *model.Tarif
	tarifs     []*model.Tarif
	TimeFormat string
	ctx        *context.AppContext
}

// Tarif instancie un tarifService ()
// Chaque appel à Tarif crée une instance unique pour éviter tout chevauchement de data
// Seul accesseur aux méthodes exportées de tarifService, cette méthode permet
// de s'assurer qu'on utilise un tarifService correctement instancié
// l'appel sera de la forme service.Tarif().Method()
func Tarif(ctx *context.AppContext) *tarifService {
	t := &model.Tarif{}
	tarifs := []*model.Tarif{}
	return &tarifService{
		// on pousse le format depuis le model
		// c'est le model qui dicte sa conduite au service
		// le model n'a pas connaissance du service
		TimeFormat: model.TARIF_TIME_FORMAT,
		t:          t,
		tarifs:     tarifs,
		ctx:        ctx,
	}
}

// CalculateTotal
// - extract current tarifs
// - compare description (can be done in request)
// - get the last by date
func (ts *tarifService) CalculateTotal(t *model.Transaction) error {
	total := 0.0
	// On load tout les tarifs et on gère le traitement en interne,
	// ca limite les requêtes
	err := ts.loadTarifs()
	if err != nil {
		return err
	}
	for _, reservation := range t.Reservation {
		extractedTarifs := ts.extractByDescription(reservation.Tarif)
		if len(extractedTarifs) == 0 {
			return fmt.Errorf("%s no tarif with this description: %s", utils.Use().GetStack(ts.CalculateTotal), reservation.Tarif)
		}
		current, err := ts.getLast(extractedTarifs)
		if err != nil {
			return err
		}
		total = total + current.Prix
	}
	t.Total = total
	return nil
}

// Load fullfill Tarif.tarifs with the content of tarifs/tarifs.json file
func (ts *tarifService) LoadFromFile() ([]*model.Tarif, error) {
	tarifFilePath := filepath.Join(ts.ctx.Opts.GetExeFolder(), "tarifs/tarifs.json")
	fmt.Printf("%s tarifFilePath : %s", utils.Use().GetStack(ts.LoadFromFile), tarifFilePath)
	file, err := ioutil.ReadFile(tarifFilePath)
	if err != nil {
		return nil, fmt.Errorf("%s %s", utils.Use().GetStack(ts.LoadFromFile), err.Error())
	}
	tarifs := []*model.Tarif{}
	err = json.Unmarshal(file, &tarifs)
	if err != nil {
		return nil, err
	}
	return tarifs, nil
}

func (ts *tarifService) loadTarifs() error {
	var err error
	ts.tarifs, err = ts.ctx.TarifManager.FindAll()
	if err != nil {
		return err
	}
	return nil
}

// expect ts.Tarifs to be loaded
func (ts *tarifService) extractByDescription(desc string) []*model.Tarif {
	tarifs := []*model.Tarif{}
	for _, tarif := range ts.tarifs {
		if tarif.Description == desc {
			tarifs = append(tarifs, tarif)
		}
	}
	return tarifs
}

func (ts *tarifService) getLast(extractedTarifs []*model.Tarif) (*model.Tarif, error) {
	if len(extractedTarifs) == 0 {
		return nil, fmt.Errorf("%s extractedTarifs slice is empty", utils.Use().GetStack(ts.getLast))
	}
	current := extractedTarifs[0]
	for _, extractedTarif := range extractedTarifs {
		timeCurrent, err := time.Parse(ts.TimeFormat, current.Date)
		if err != nil {
			return nil, err
		}
		timeExtract, err := time.Parse(ts.TimeFormat, extractedTarif.Date)
		if err != nil {
			return nil, err
		}
		if timeCurrent.Before(timeExtract) {
			current = extractedTarif
		}
	}
	return current, nil
}
