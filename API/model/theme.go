/*
 * File: theme.go
 * Project: ABD4/VMD Escape Game
 * File Created: Tuesday, 30th October 2018 8:19:50 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Wednesday, 31st October 2018 12:12:35 am
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

	"gopkg.in/mgo.v2/bson"
)

// Theme is composed:
type Theme struct {
	ObjectID bson.ObjectId `bson:"_id,omitempty"`
	ID       string        `json:"ID"`
	Theme    string        `json:"Theme"`
	mapped   map[string]interface{}
}

// ToString return string conversion of marshal user
// absorb error...
func (t *Theme) ToString() string {
	ret, _ := t.MarshalJSON()
	return string(ret)
}

func (td *Theme) GetID() string {
	return td.ID
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (t *Theme) UnmarshalFromRequest(r *http.Request) error {
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

// MarshalJSON implement ISerial and json.MarshalJSON
func (t Theme) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.toMap())
}

func (t *Theme) toMap() map[string]interface{} {
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

func (t Theme) GetMapped() map[string]interface{} {
	if len(t.mapped) == 0 {
		t.toMap()
	}
	return t.mapped
}

func (t Theme) ToArrayString(themes []*Theme) []string {
	ret := []string{}
	for _, theme := range themes {
		ret = append(ret, theme.Theme)
	}
	return ret
}

// ToES is a ack to avoid parsing-mapping error in elastic search
// it seems that ES can't parse string representation of hex mongo ObjecID
func (t *Theme) ToES() *Theme {
	tToES := t
	tToES.ID = ""
	return tToES
}
