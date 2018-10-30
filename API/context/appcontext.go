/*
 * File: appcontext.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:32:53 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Monday, 29th October 2018 5:29:25 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

import (
	"ABD4/API/logger"

	mgo "gopkg.in/mgo.v2"
	elastic "gopkg.in/olivere/elastic.v5"
)

// AppContext define globals tools and variable usefull in the API
// It embed the dao's objects (XxxManager *manager.XxxManager),
// a ResponseWriter which offer shorthand to send uniformised Response
type AppContext struct {
	Rw                 IResponseWriter
	SessionUser        ISessionUser
	Opts               IServerOption
	UserManager        IUserManager
	TransactionManager ITransactionManager
	TarifManager       ITarifManager
	Mongo              *mgo.Session
	ElasticClient      *elastic.Client
	Log                *logger.AppLogger
	Exe                string
	Logpath            string
	DataPath           string
}
