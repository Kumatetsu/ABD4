/*
 * File: mock.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Friday, 12th October 2018 11:34:26 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package test

import (
	"ABD4/API/context"
	"ABD4/API/model"
	"fmt"
	"log"
	"sort"
	"strconv"
)

// Mock embed model for tests
// Stock it in App global variable in test
// Retrieve and stock all persisted data in
// temporary database.
type Mock struct {
	// exemple:
	Ctx      *context.AppContext
	Users    []string
	PostUser *model.User
}

func (m *Mock) MockUsers(n int) {
	var it int
	var err error
	manager := m.Ctx.UserManager
	users := []*model.User{}
	for it = 0; it < n; it++ {
		name := "test_" + strconv.Itoa(it)
		email := name + "@etna-alternance.net"
		password := "test"
		u := &model.User{
			Name:       name,
			Email:      email,
			Password:   password,
			Permission: "permission",
			Claim:      "",
		}
		u, err = manager.Create(u)
		if err != nil {
			log.Fatalf("mocking users failed: %s", err.Error())
		}
		users = append(users, u)
		// m.Users = append([]string{u.ToString()}, m.Users...)
	}
	fmt.Printf("Users: %v", users)
	sort.Slice(users, func(i, j int) bool {
		return users[i].OrderByEmail(users, i, j)
	})
	for _, u := range users {
		m.Users = append(m.Users, u.ToString())
	}
	m.PostUser = &model.User{
		Name:       "register",
		Email:      "regist_t@etna-alternance.net",
		Password:   "register",
		Permission: "permission",
	}
}
