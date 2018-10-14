/*
 * File: iresponsewriter.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:30:13 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Thursday, 11th October 2018 4:39:21 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/iserial"
	"net/http"
)

// IResponseWriter interface define the required methods to
// use the AppContext.Rw variable into the API
type IResponseWriter interface {
	Send(*AppContext, http.ResponseWriter, int, iserial.Serializable, string, string)
	SendError(*AppContext, http.ResponseWriter, int, string, string)
	SendItSelf(*AppContext, http.ResponseWriter)
	NewResponse(int, string, string, string) IResponseWriter
}
