/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Thursday, 1st November 2018 11:09:25 pm
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

// Acheteur is composed:
type Transaction struct {
	ObjectID    bson.ObjectId `bson:"_id,omitempty"`
	ID          string        `json:"id"`
	Acheteur    Acheteur      `json:"Acheteur"`
	Game        Game          `json:"Game"`
	Themes      []string      `json:"Theme"`
	Reservation []Reservation `json:"Reservation"`
	Total       float64       `json:"Total"`
	createdAt   time.Time
	updatedAt   time.Time
	mapped      map[string]interface{}
}

var TRANSACTION = "transaction"

func (t Transaction) GetID() string {
	return t.ID
}

// ToString return string conversion of marshal user
// absorb error...
func (t *Transaction) ToString() string {
	ret, _ := t.MarshalJSON()
	return string(ret)
}

func (t *Transaction) SetCreatedAt(now time.Time) {
	t.createdAt = now
}

func (t *Transaction) SetUpdatedAt(now time.Time) {
	t.updatedAt = now
}

func (t Transaction) GetCreatedAt() time.Time {
	return t.createdAt
}

func (t Transaction) GetUpdatedAt() time.Time {
	return t.updatedAt
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (t *Transaction) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(t.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, t)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(t.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// MarshalJSON implement json.MarshalJSON
func (t Transaction) MarshalJSON() ([]byte, error) {
	// we parse t in map[string]interface{}
	return json.Marshal(t.toMap())
}

// use reflect to define a generic map[string]interface{}
func (t *Transaction) toMap() map[string]interface{} {
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

// GetMapped allow us to access t.mapped, making it a singleton in Transaction
func (t Transaction) GetMapped() map[string]interface{} {
	if len(t.mapped) == 0 {
		t.toMap()
	}
	return t.mapped
}

// ToES is a ack to avoid parsing-mapping error in elastic search
// it seems that ES can't parse string representation of hex mongo ObjecID
func (t *Transaction) ToES() *Transaction {
	tToES := t
	tToES.ID = ""
	return tToES
}
