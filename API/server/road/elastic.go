/*
 * File: elastic.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:56:55 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:57:07 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetElasticRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /elastic/index/all",
			Method:          GET,
			Pattern:         "/index/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetCreateIndexation,
		},
		{
			Name:            GET + " /elastic/index/{index}",
			Method:          GET,
			Pattern:         "/index/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetCreateIndex,
		},
		{
			Name:            GET + " /elastic/rmindex/all",
			Method:          GET,
			Pattern:         "/rmindex/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetRemoveIndexation,
		},
		{
			Name:            GET + " /elastic/rmindex/{index}",
			Method:          GET,
			Pattern:         "/rmindex/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetRemoveIndex,
		},
		{
			Name:            GET + " /elastic/reindex/all",
			Method:          GET,
			Pattern:         "/reindex/all",
			StatusProtected: false,
			HandlerFunc:     handler.GetReindexation,
		},
		{
			Name:            GET + " /elastic/reindex/{index}",
			Method:          GET,
			Pattern:         "/reindex/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetReindex,
		},
		{
			Name:            GET + " /elastic/indexdata",
			Method:          GET,
			Pattern:         "/indexdata",
			StatusProtected: false,
			HandlerFunc:     handler.GetIndexationData,
		},
		{
			Name:            GET + " /elastic/indexdata/{index}",
			Method:          GET,
			Pattern:         "/indexdata/{index}",
			StatusProtected: false,
			HandlerFunc:     handler.GetIndexData,
		},
	}
}
