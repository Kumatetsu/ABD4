/*
 * File: logger.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Sunday, 30th September 2018 4:54:16 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package logger

import (
	"io"
	"log"
)

// AppLogger simply embed three log.Logger struct
type AppLogger struct {
	Debug *log.Logger
	Info  *log.Logger
	Error *log.Logger
}

// Instanciate receive three io.Writer to build an AppLogger structure:
//
// type AppLogger struct {
//	Debug *log.Logger
//	Info  *log.Logger
//	Error *log.Logger
// }
//
// Typical call:
//
// ctx.Log.Info.Printf("%s Setting headers...", utils.Use().GetStack(SetHeaders))
//
// Typical log entry:
//
// INFO:2018/09/28 14:28:56 header.go:31: [ABD4/API/server/middleware.SetHeaders]  Setting headers...
func Instanciate(debugOut io.Writer, infoOut io.Writer, errorOut io.Writer) *AppLogger {
	return &AppLogger{
		Debug: log.New(infoOut, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile),
		Info:  log.New(debugOut, "INFO:", log.Ldate|log.Ltime|log.Lshortfile),
		Error: log.New(errorOut, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile),
	}
}
