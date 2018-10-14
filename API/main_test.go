/*
 * File: README.md
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 3:55:33 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package main

import (
	"ABD4/API/model"
	"ABD4/API/server"
	"ABD4/API/test"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testApp *App
var mock *test.Mock = &test.Mock{}

// For tests we need to launch the server and continue the execution
// To do this and be sure that everything is initialized, we launch
// the serve in a goroutine and we use a channel to trigger
// a bool telling us that everything is initialized
func launchTestApp(opts *server.Option, signal chan bool) {
	testApp = &App{}
	err := testApp.InitializeAndWaitForSignal(opts, signal)
	if err != nil {
		log.Fatal(err)
	}
}
func TestMain(m *testing.M) {
	dir := os.Getenv("GOPATH") + "/src/ABD4/API/test"
	opts := &server.Option{}
	port := strconv.Itoa(8001)
	opts.Hydrate(port, "127.0.0.1", "test", dir, dir+"/log", "mongo", dir+"/data", "127.0.0.1", "27017", true)
	signal := make(chan bool)
	os.Remove(opts.GetDatapath() + "/users.dat")
	go launchTestApp(opts, signal)
	// the signal receive a value when context.AppContext is initialized
	if <-signal {
		testApp.Ctx.Log.Info.Print("App initiliazed in go routine\n\n\n")
	}
	mock.Ctx = testApp.Ctx
	testApp.Ctx.UserManager.RemoveAll()
	mock.MockUsers(2)
	code := m.Run()
	// On delete les data pour repartir sur un context neat la prochaine fois
	os.Remove(opts.GetDatapath() + "/users.dat")
	os.Exit(code)
}

func TestUser(t *testing.T) {
	var req *http.Request
	var err error
	assert := assert.New(t)
	testSuite := &test.Suite{}
	for _, test := range testSuite.GetTests(mock) {
		body := &strings.Reader{}
		response := &server.Response{}
		user := &model.User{}
		w := httptest.NewRecorder()

		testApp.Ctx.Log.Info.Print("Lauching TestUser")
		testApp.Ctx.Log.Info.Printf("Lauching test: %s", test.Description)
		if test.Method == "POST" {
			body = strings.NewReader(mock.PostUser.ToString())
		}
		req, err = http.NewRequest(test.Method, test.URL, body)
		assert.NoError(err)
		test.Handler(testApp.Ctx, w, req)
		assert.Equal(test.ExpectedStatusCode, w.Code, test.Description)
		if test.Method == "GET" {
			assert.Equal(test.ExpectedBody, w.Body.String(), test.Description)
		} else if test.URL == "/auth/register" {
			err := json.Unmarshal(w.Body.Bytes(), response)
			assert.NoError(err)
			err = json.Unmarshal([]byte(response.Data), user)
			assert.NoError(err)
			assert.NotNil(user.ID, test.Description)
			assert.Equal(mock.PostUser.Name, user.Name, test.Description)
			assert.Equal(mock.PostUser.Email, user.Email, test.Description)
			assert.Equal(mock.PostUser.Password, user.Password, test.Description)
			assert.Equal(mock.PostUser.Permission, user.Permission, test.Description)
			assert.Equal(mock.PostUser.Claim, user.Claim, test.Description)
		} else if test.URL == "/auth/login" {
			err := json.Unmarshal(w.Body.Bytes(), response)
			assert.NoError(err)
			assert.NotNil(response.Data, test.Description)
		}
	}
	testApp.Ctx.Log.Info.Print("\n\n\nTest User finished")
}
