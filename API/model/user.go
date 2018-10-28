/*
 * File: user.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 11:49:17 am
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
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User is composed:
// Claim are used for jwt token signature
// Permission
type User struct {
	ObjectID   bson.ObjectId `bson:"_id,omitempty"`
	ID         string        `json:"id,omitempty"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Password   string        `json:"password"`
	Permission string        `json:"permission,omitempty"`
	Claim      string        `json:"claim,omitempty"`
	createdAt  time.Time
	updatedAt  time.Time
	mapped     map[string]interface{}
}

var USER = "user"

// ToString return string conversion of marshal user
// absorb error...
func (u *User) ToString() string {
	ret, _ := u.Marshal()
	return string(ret)
}

func (u *User) SetCreatedAt(now time.Time) {
	u.createdAt = now
}

func (u *User) SetUpdatedAt(now time.Time) {
	u.updatedAt = now
}

func (u User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

// UnmarshalFromRequest take a request object supposed to contain
// a user data object, extract it, convert to User and send back
// the string representation of the content
func (u *User) UnmarshalFromRequest(r *http.Request) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(u.UnmarshalFromRequest), err.Error())
	}
	err = json.Unmarshal(body, u)
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(u.UnmarshalFromRequest), err.Error())
	}
	return nil
}

// GetId implement ISessionUser
func (u *User) GetID() string {
	return u.ID
}

// Marshal implement ISerial
func (u User) Marshal() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) OrderById(results []*User, i, j int) bool {
	return strings.Compare(results[i].GetID(), results[j].GetID()) == 1
}

func (u *User) OrderByEmail(results []*User, i, j int) bool {
	return strings.Compare(results[i].Email, results[j].Email) == 1
}

func (u *User) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(u).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	u.mapped = mapped
	return mapped
}

func (u User) GetMapped() map[string]interface{} {
	if len(u.mapped) == 0 {
		u.toMap()
	}
	return u.mapped
}

func (u *User) ToES() *User {
	uToES := u
	uToES.ID = ""
	return uToES
}
