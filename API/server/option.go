/*
 * File: option.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 14th October 2018 1:15:35 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

type Option struct {
	exe       string
	env       string
	debug     bool
	logpath   string
	dbType    string
	datapath  string
	address   string
	port      string
	ip        string
	mongoIP   string
	mongoPort string
}

var (
	PROD = "production"
	DEV  = "développement"
	TEST = "test"
)

func (o *Option) Hydrate(port, ip, env, dir, logpath, dbType, datapath, mongoIP, mongoPort string, debug bool) {
	o.port = port
	o.ip = ip
	o.datapath = datapath
	o.dbType = dbType
	o.logpath = logpath
	o.env = env
	o.debug = debug
	o.exe = dir
	o.mongoIP = mongoIP
	o.mongoPort = mongoPort
}

// GetAddress concat ip and port and affect to address if needed
// else default address is define to 127.0.0.1:80
func (o *Option) GetAddress() string {
	return o.ip + ":" + o.port
}

/*
context.serverOption interface:

type IServerOption interface {
		GetExeFolder() string
		SetEnv(string)
		SetLogpath(string)
		SetDatapath(string)
		GetEnv() string
		GetLogpath() string
		GetDatapath() string
		GetAddress() string
		GetPort() string
		GetIp() string
	}
*/

// GetIP return ip
func (o *Option) GetIP() string {
	return o.ip
}

// GetPort return port
func (o *Option) GetPort() string {
	return o.port
}

// GetExeFolder return ./app.exe folder
func (o *Option) GetExeFolder() string {
	return o.exe
}

// GetLogpath return the path to logs folder
func (o *Option) GetLogpath() string {
	return o.logpath
}

func (o *Option) GetDatabaseType() string {
	return o.dbType
}

// GetDatapath return the path to data folder
func (o *Option) GetDatapath() string {
	return o.datapath
}

// GetEnv return environnement description default on DEV: "développement"
func (o *Option) GetEnv() string {
	if o.env == "" {
		return DEV
	}
	return o.env
}

func (o *Option) GetMongoIP() string {
	return o.mongoIP
}

func (o *Option) GetMongoPort() string {
	return o.mongoPort
}

func (o *Option) SetMongoPort(port string) {
	o.mongoPort = port
}

func (o *Option) SetDatabaseType(dbType string) {
	o.dbType = dbType
}

// SetEnv allow a IServerOption to change the environnement
// "prod" => "production"
// "test" => "test"
// (default) "dev" => développement
func (o *Option) SetEnv(env string) {
	if env == "prod" {
		o.env = PROD
	} else if env == "test" {
		o.env = TEST
	} else {
		o.env = DEV
	}
}

// SetLogpath allow IServerOption to change the path to logs
func (o *Option) SetLogpath(logpath string) {
	o.logpath = logpath
}

// SetDatapath allow IServerOption to change the path to data
func (o *Option) SetDatapath(datapath string) {
	o.datapath = datapath
}

func (o *Option) SetMongoIP(ip string) {
	o.mongoIP = ip
}
