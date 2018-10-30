/*
 * File: acheteur.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:29:27 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
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
	"time"

	"gopkg.in/mgo.v2/bson"
)

var (
	TARIF_TIME_FORMAT = "2006-01-02T15:04:05Z"
)

// Acheteur is composed:
type Tarif struct {
	ObjectID    bson.ObjectId `bson:"_id,omitempty"`
	ID          string        `json:"ID"`
	Description string        `json:"Description"`
	Prix        float64       `json:"Prix,float64"`
	Date        string        `json:"Date"`
	mapped      map[string]interface{}
}

// ToString return string conversion of marshal user
// absorb error...
func (t *Tarif) ToString() string {
	ret, _ := t.MarshalJSON()
	return string(ret)
}

func (td *Tarif) GetID() string {
	return td.ID
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (t *Tarif) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(t.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, t)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(t.UnmarshalFromRequest), err.Error())
	}
	t.Date = time.Now().Format(TARIF_TIME_FORMAT)
	return nil
}

// Marshal implement ISerial
func (t Tarif) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.toMap())
}

func (t *Tarif) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(t).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	t.mapped = mapped
	return mapped
}

func (t Tarif) GetMapped() map[string]interface{} {
	if len(t.mapped) == 0 {
		t.toMap()
	}
	return t.mapped
}
