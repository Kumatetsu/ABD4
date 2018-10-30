/*
 * File: claim.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 23rd October 2018 5:46:56 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package model

import (
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	appName = "abd4"
	// SECRET used in jwt signature
	SECRET = "==+ABD4@VDMEscapeGame13683"
)

// Claim is the type used to parse
// a unique Json Web Token, we can stock what we want,
// here the user id to retrieve it during authentication process
type Claim struct {
	User string `json:"user,string"`
	jwt.StandardClaims
	mapped map[string]interface{}
}

func (c *Claim) toString() (string, error) {
	json, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("<<<< <<<< %s %s", utils.Use().GetStack(c.toString), err.Error())
	}
	return string(json), nil
}

// NewFromUser fullfill the Claim object with
// the user information, set the token validity to one week
func (c *Claim) NewFromUser(user *User) error {
	var err error

	exp := time.Now()
	c.User = user.ToString()
	if err != nil {
		return fmt.Errorf("<<<< %s %s", utils.Use().GetStack(c.NewFromUser), err.Error())
	}
	exp = exp.Add(time.Minute * 60 * 24 * 7) // token is set for a week
	c.ExpiresAt = exp.Unix()
	c.Issuer = fmt.Sprintf("%s for %s", appName, user.Email)
	user.Claim, err = c.toString()
	return err
}

// GetToken generate a new token based on relying Claim
func (c *Claim) GetToken() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenStr, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return "", fmt.Errorf("<<<< %s Error generating jwt: %s", utils.Use().GetStack(c.GetToken), err.Error())
	}
	return tokenStr, nil
}

// Marshal to implement ISerial interface
func (c Claim) Marshal() ([]byte, error) {
	return json.Marshal(c)
}

func (c Claim) toMap() map[string]interface{} {
	mapped := make(map[string]interface{})
	structure := reflect.ValueOf(c).Elem()
	typeOfStructure := structure.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if field.CanInterface() {
			mapped[strings.ToLower(typeOfStructure.Field(i).Name)] = field.Interface()
		}
	}
	c.mapped = mapped
	return mapped
}

func (c Claim) GetMapped() map[string]interface{} {
	if len(c.mapped) == 0 {
		c.toMap()
	}
	return c.mapped
}
