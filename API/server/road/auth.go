/*
 * File: auth.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:55:19 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:55:48 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetAuthRouting() []*Road {
	return []*Road{
		{
			Name:            OPTIONS + " /auth/login",
			Method:          OPTIONS,
			Pattern:         "/login",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            OPTIONS + " /auth/register",
			Method:          OPTIONS,
			Pattern:         "/register",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /auth/login",
			Method:          POST,
			Pattern:         "/login",
			StatusProtected: false,
			HandlerFunc:     handler.Login,
		},
		{
			Name:            POST + " /auth/register",
			Method:          POST,
			Pattern:         "/register",
			StatusProtected: false,
			HandlerFunc:     handler.Register,
		},
	}
}
