/*
 * File: tarif.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:45:21 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:51:25 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetTarifRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /tarif",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.GetTarifs,
		},
		{
			Name:            GET + " /tarif/loadfile",
			Method:          GET,
			Pattern:         "/loadfile",
			StatusProtected: false,
			HandlerFunc:     handler.LoadTarifsFromFile,
		},
		{
			Name:            OPTIONS + " /tarif",
			Method:          OPTIONS,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /tarif",
			Method:          POST,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.AddTarif,
		},
		{
			Name:            DELETE + " /tarif",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveAllTarif,
		},
	}
}
