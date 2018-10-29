/*
 * File: test.go
 * Project: ABD4/VMD Escape Game
 * File Created: Wednesday, 10th October 2018 4:21:20 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 3:46:25 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package mongo

import (
	"ABD4/API/utils"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

func GetMongo(serverAddr string) (*mgo.Session, error) {
	session, err := mgo.Dial(serverAddr)
	if err != nil {
		return nil, fmt.Errorf("%s mgo.Dial: %s", utils.Use().GetStack(GetMongo), err.Error())
	}
	/*
		Strong consistency uses a unique connection with the master
		so that all reads and writes are as up-to-date as possible and consistent with each other.

		Monotonic consistency will start reading from a slave if possible,
		so that the load is better distributed, and once the first write happens the connection is switched to the master.
		This offers consistent reads and writes, but may not show the most up-to-date data on reads which precede the first write.

		Eventual consistency offers the best resource usage,
		distributing reads across multiple slaves and writes across multiple connections to the master,
		but consistency isn't guaranteed.
	*/
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
