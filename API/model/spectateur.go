/*
 * File: spectateur.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 23rd October 2018 5:51:01 pm
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
)

// Spectateur is composed:
type Spectateur struct {
	Civilite string `json:"Civilite"`
	Nom      string `json:"Nom"`
	Prenom   string `json:"Prenom"`
	Age      int    `json:"Age, float64"`
	mapped   map[string]interface{}
}

// ToString return string conversion of marshal user
// absorb error...
func (s *Spectateur) ToString() string {
	ret, _ := s.Marshal()
	return string(ret)
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (s *Spectateur) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(s.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(s.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// Marshal implement ISerial
func (s Spectateur) Marshal() ([]byte, error) {
	return json.Marshal(s)
}

func (s Spectateur) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(s).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	s.mapped = mapped
	return mapped
}

func (s Spectateur) GetMapped() map[string]interface{} {
	if len(s.mapped) == 0 {
		s.toMap()
	}
	return s.mapped
}
