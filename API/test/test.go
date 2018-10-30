/*
 * File: test.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 28th October 2018 2:12:15 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package test

import (
	"ABD4/API/context"
	"ABD4/API/handler"
	"ABD4/API/server"
	"fmt"
	"net/http"
	"strings"
)

// Suite general structure to embed
// your own Test set distribution
type Suite struct{}

// Test define content used to launch a test
type Test struct {
	Description        string
	URL                string
	Method             string
	Body               string
	Handler            context.CustomHandler
	ExpectedStatusCode int
	ExpectedBody       string
}

// GetTests return tests definitions
func (ts *Suite) GetTests(mock *Mock) []*Test {
	fmt.Printf("test on users: %v, and post user: %v", mock.Users, mock.PostUser)
	allUsersResponse := &server.StringResponse{
		Status: 200,
		Data:   `[` + strings.Join(mock.Users, ",") + `]`,
	}
	registerUserResponse := &server.StringResponse{
		Status: 200,
		Data:   mock.PostUser.ToString(),
	}
	loginUserResponse := &server.StringResponse{
		Status: 200,
		Data:   "token",
	}
	return []*Test{
		{
			Description: "Get on root",
			URL:         "/",
			Method:      "GET",
			Handler: func(ctx *context.AppContext, w http.ResponseWriter, r *http.Request) {
				// w.Write take a []byte, fmt.Sprintf return a string
				// we cast with []byte(string)
				w.Write([]byte(fmt.Sprintf("Root road")))
			},
			ExpectedBody:       fmt.Sprintf("Root road"),
			ExpectedStatusCode: 200,
		},
		{
			Description:        "GET on /user",
			URL:                "/user",
			Method:             "GET",
			Handler:            handler.GetUsers,
			ExpectedBody:       allUsersResponse.ToString(),
			ExpectedStatusCode: 200,
		},
		{
			Description:        "POST on /auth/register",
			URL:                "/auth/register",
			Method:             "POST",
			Handler:            handler.Register,
			ExpectedBody:       registerUserResponse.ToString(),
			ExpectedStatusCode: 200,
		},
		{
			Description:        "POST on /auth/login",
			URL:                "/auth/login",
			Method:             "POST",
			Handler:            handler.Login,
			ExpectedBody:       loginUserResponse.ToString(),
			ExpectedStatusCode: 200,
		},
	}
}
