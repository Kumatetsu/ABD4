/*
 * File: transaction.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:57:20 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:57:44 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetTransactionRouting() []*Road {
	return []*Road{
		{
			Name:            OPTIONS + " /transaction",
			Method:          OPTIONS,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.Option,
		},
		{
			Name:            POST + " /transaction",
			Method:          POST,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.AddTransaction,
		},
		{
			Name:            GET + " /transaction",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.GetTransaction,
		},
		{
			Name:            DELETE + " /transaction",
			Method:          DELETE,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.RemoveAllTX,
		},
	}
}
