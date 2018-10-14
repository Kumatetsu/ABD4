/*
 * File: iserveroption.go
 * Project: ABD4/VMD Escape Game
 * File Created: Thursday, 11th October 2018 4:28:51 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Saturday, 13th October 2018 7:19:29 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package context

// IServerOption :
// SetEnv setter for environnement ("prod", "dev", "test")
// GetEnv getter for environnement
// GetExeFolder getter for .exe folder
// GetLogpath getter for log folder
// GetDataPath getter for bolt data folder
// GetAddress return full host name (IP:Port)
// GetPort return port
// GetIp return ip
type IServerOption interface {
	SetEnv(string)
	SetLogpath(string)
	SetDatabaseType(string)
	SetDatapath(string)
	SetMongoIP(ip string)
	SetMongoPort(port string)
	GetExeFolder() string
	GetEnv() string
	GetLogpath() string
	GetDatabaseType() string
	GetDatapath() string
	GetAddress() string
	GetPort() string
	GetIP() string
	GetMongoIP() string
	GetMongoPort() string
}
