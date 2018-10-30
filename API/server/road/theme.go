/*
 * File: theme.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:58:01 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 12:00:12 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetThemeRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /theme [WIP]",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            GET + " /theme/loadfile [WIP]",
			Method:          GET,
			Pattern:         "/loadfile",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            OPTIONS + " /theme [WIP]",
			Method:          OPTIONS,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /theme [WIP]",
			Method:          POST,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            DELETE + " /theme [WIP]",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
	}
}
