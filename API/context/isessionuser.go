/*
 * File: isessionuser.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:29:46 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Thursday, 11th October 2018 4:39:33 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

// ISessionUser abstract the user from model
type ISessionUser interface {
	GetID() string
}
