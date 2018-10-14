/*
 * File: response.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: billaud_j castel_a masera_m
 * Contact: (billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Friday, 12th October 2018 6:37:05 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

// Response define a unique response format for API
// Status: code http
// Data: json encoded data content
// Message: user friendly string
// Detail: technical detail from API (error detail)
// It should implement context.IResponseWriter interface
type Response struct {
	Status  int    `json:"status"`
	Data    string `json:"data"`
	Message string `json:"message,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

var (
	marshalError = "Erreur d'encodage"
)

func (r *Response) ToString() string {
	response, _ := r.Marshal()
	return string(response)
}

func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// NewResponse is to use with SendItSelf function to send a non iserial.Serializable response
// used in IResponseWriter implementation
func (r Response) NewResponse(st int, d, msg, dt string) context.IResponseWriter {
	return Response{
		Status:  st,
		Data:    d,
		Message: msg,
		Detail:  dt,
	}
}

// Send write and log a successfull answer returning a iserial.Serializable entity
// used in IResponseWriter implementation
func (r Response) Send(ctx *context.AppContext, w http.ResponseWriter, status int, i iserial.Serializable, msg, detail string) {
	r.Status = status
	data, err := i.Marshal()
	if err != nil {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(r.Send), err.Error())
		r.SendError(ctx, w, r.Status, msg, marshalError)
		return
	}
	r.Data = string(data)
	r.Message = msg
	r.Detail = detail
	r.SendItSelf(ctx, w)
}

// SendItSelf is used to send an already builded Response object
// used in IResponseWriter implementation
func (r Response) SendItSelf(ctx *context.AppContext, w http.ResponseWriter) {
	ret, err := json.Marshal(r)
	if err != nil {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(r.SendItSelf), err.Error())
		r.SendError(ctx, w, r.Status, msg, marshalError)
		return
	}
	ctx.Log.Info.Print(string(ret))
	w.WriteHeader(r.Status)
	w.Write(ret)
}

// SendError write and log an error answer
// used in IResponseWriter implementation
func (r Response) SendError(ctx *context.AppContext, w http.ResponseWriter, status int, msg, detail string) {
	r.Status = status
	r.Message = msg
	r.Detail = detail
	ret, err := json.Marshal(r)
	if err != nil {
		msg = fmt.Sprintf("%s %s", utils.Use().GetStack(r.SendError), marshalError)
		ctx.Log.Error.Print(msg)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
		return
	}
	ctx.Log.Error.Print(msg + " " + detail)
	w.WriteHeader(r.Status)
	w.Write(ret)
	return
}
