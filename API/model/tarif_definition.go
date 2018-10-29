/*
 * File: acheteur.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 23rd October 2018 5:52:38 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package model

import (
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

// Acheteur is composed:
type TarifDefinition struct {
	ID        string `json:"ID"`
	TarifType string `json:"TarifType"`
	Prix      int    `json:"Prix, float64"`
	Date      string `json:"Date"`
	mapped    map[string]interface{}
}

// ToString return string conversion of marshal user
// absorb error...
func (tarifDef *TarifDefinition) ToString() string {
	ret, _ := tarifDef.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (tarifDef *TarifDefinition) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(tarifDef.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, tarifDef)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(tarifDef.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (tarifDef TarifDefinition) Marshal() ([]byte, error) {
	return json.Marshal(tarifDef)
}

func (tarifDef TarifDefinition) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(tarifDef).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	tarifDef.mapped = mapped
	return mapped
}

func (tarifDef TarifDefinition) GetMapped() map[string]interface{} {
	if len(tarifDef.mapped) == 0 {
		tarifDef.toMap()
	}
	return tarifDef.mapped
}
