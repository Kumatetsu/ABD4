/*
 * File: game.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 23rd October 2018 5:48:06 pm
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

// Game is composed:
type Game struct {
	Nom     string `json:"Nom"`
	Jour    string `json:"Jour"`
	Horaire string `json:"Horaire"`
	VR      string `json:"VR"`
	mapped  map[string]interface{}
}

// ToString return string conversion of marshal user
// absorb error...
func (g *Game) ToString() string {
	ret, _ := g.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (g *Game) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(g.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, g)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(g.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (g Game) Marshal() ([]byte, error) {
	return json.Marshal(g)
}

func (g Game) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(g).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	g.mapped = mapped
	return mapped
}

func (g Game) GetMapped() map[string]interface{} {
	if len(g.mapped) == 0 {
		g.toMap()
	}
	return g.mapped
}
