/*
 * File: theme.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:58:01 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 5th November 2018 6:13:50 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetThemeRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /theme",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.GetTheme,
		},
		{
			Name:            GET + " /theme/loadfile",
			Method:          GET,
			Pattern:         "/loadfile",
			StatusProtected: false,
			HandlerFunc:     handler.LoadThemesFromFile,
		},
		{
			Name:            OPTIONS + " /theme",
			Method:          OPTIONS,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /theme",
			Method:          POST,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.AddTheme,
		},
		{
			Name:            DELETE + " /theme",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveAllThemes,
		},
		{
			Name:            DELETE + " /theme/id/{id}",
			Method:          DELETE,
			Pattern:         "/id/{id}",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveThemeByID,
		},
		{
			Name:            DELETE + " /theme/theme/{theme}",
			Method:          DELETE,
			Pattern:         "/theme/{theme}",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveThemeByName,
		},
	}
}
