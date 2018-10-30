/*
 * File: angularIoRouting.go
 * Project: ABD4/VMD Escape Game
 * File Created: Tuesday, 30th October 2018 12:05:46 am
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Tuesday, 30th October 2018 1:38:31 am
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import (
	"ABD4/API/context"
	"ABD4/API/handler"
)

func getAngularRouting() []*Road {
	return []*Road{
	/* HOME
	{
		Name:        GET + "/",
		Method:      GET,
		Pattern:     "",
		HandlerFunc: handler.Angular,
	},
	*/
	}
}

func GetWebAppRouting(ctx *context.AppContext) []*Road {
	var routes []*Road
	routes = append(routes)
	for _, road := range getAngularRouting() {
		routes = append(routes, road)
	}
	return routes
}

func GetHome() *Road {
	return &Road{
		Name:        GET + "/",
		Method:      GET,
		Pattern:     "/",
		HandlerFunc: handler.Angular,
	}
}
