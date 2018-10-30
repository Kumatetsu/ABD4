/*
 * File: backup.go
 * Project: ABD4/VMD Escape Game
 * File Created: Monday, 29th October 2018 11:56:13 pm
 * Author: ayad_y@etna-alternance.net billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 11:56:36 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package road

import "ABD4/API/handler"

func GetBackupRouting() []*Road {
	return []*Road{
		{
			Name:            GET + " /backup",
			Method:          GET,
			Pattern:         "",
			StatusProtected: false,
			HandlerFunc:     handler.BackupBoltDatabase,
		},
	}
}
