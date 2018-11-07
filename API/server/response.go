/*
 * File: response.go
 * Project: ABD4/VMD Escape Game
 * File Created: Sunday, 30th September 2018 2:24:43 pm
 * Author: ayad_y billaud_j castel_a masera_m
 * Contact: (ayad_y@etna-alternance.net billaud_j@etna-alternance.net castel_a@etna-alternance.net masera_m@etna-alternance.net)
 * -----
 * Last Modified: Saturday, 3rd November 2018 11:55:02 pm
 * Modified By: Aurélien Castellarnau
 * -----
 * Copyright © 2018 - 2018 ayad_y billaud_j castel_a masera_m, ETNA - VDM EscapeGame API
 */

package server

import (
	"ABD4/API/context"
	"ABD4/API/iserial"
	"ABD4/API/service"
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
	Status  int                    `json:"status"`
	Data    map[string]interface{} `json:"data"`
	Message string                 `json:"message,omitempty"`
	Detail  string                 `json:"detail,omitempty"`
}

type ArrayResponse struct {
	Status  int                      `json:"status"`
	Data    []map[string]interface{} `json:"data"`
	Message string                   `json:"message,omitempty"`
	Detail  string                   `json:"detail,omitempty"`
}

type StringResponse struct {
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

func (r StringResponse) ToString() string {
	response, _ := json.Marshal(r)
	return string(response)
}

func (r Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// NewResponse is to use with SendItSelf function to send a non iserial.Serializable response
// used in IResponseWriter implementation
func (r Response) NewResponse(st int, msg, dt string, i iserial.Serializable) context.IResponseWriter {
	d := i.GetMapped()
	return Response{
		Status:  st,
		Data:    d,
		Message: msg,
		Detail:  dt,
	}
}

// Send write and log a successfull answer returning a iserial.Serializable entity
// used in IResponseWriter implementation
func (r Response) SendSerializable(ctx *context.AppContext, w http.ResponseWriter, status int, i iserial.Serializable, msg, detail string) {
	r.Status = status
	r.Data = i.GetMapped()
	r.Message = msg
	r.Detail = detail
	r.SendItSelf(ctx, w)
}

// SendArraySerializable attached to Response struct, use a ArrayResponse struct to send a []Serializable
func (r Response) SendArraySerializable(ctx *context.AppContext, w http.ResponseWriter, status int, i []iserial.Serializable, msg, detail string) {
	data := []map[string]interface{}{}
	arrayResponse := &ArrayResponse{}
	arrayResponse.Status = status
	for _, entity := range i {
		data = append(data, entity.GetMapped())
	}
	arrayResponse.Data = data
	arrayResponse.Message = msg
	arrayResponse.Detail = detail
	arrayResponse.sendItSelf(ctx, w)
}

// SendArraySerializable attached to Response struct, use a StringResponse struct to send a []Serializable
func (r Response) SendString(ctx *context.AppContext, w http.ResponseWriter, status int, data, msg, detail string) {
	stringResponse := &StringResponse{}
	stringResponse.Status = status
	stringResponse.Data = data
	stringResponse.Message = msg
	stringResponse.Detail = detail
	stringResponse.sendItSelf(ctx, w)
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
	service.Benchmark(ctx, "full request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	w.Write(ret)
}

// sendItSelf for ArrayResponse struct
func (ar ArrayResponse) sendItSelf(ctx *context.AppContext, w http.ResponseWriter) {
	r := &Response{}
	ret, err := json.Marshal(ar)
	if err != nil {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(ar.sendItSelf), err.Error())
		r.SendError(ctx, w, ar.Status, msg, marshalError)
		return
	}
	ctx.Log.Info.Print(string(ret))
	service.Benchmark(ctx, "full request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ar.Status)
	w.Write(ret)
}

// sendItSelf for StringResponse struct
func (sr StringResponse) sendItSelf(ctx *context.AppContext, w http.ResponseWriter) {
	r := &Response{}
	ret, err := json.Marshal(sr)
	if err != nil {
		msg := fmt.Sprintf("%s %s", utils.Use().GetStack(sr.sendItSelf), err.Error())
		r.SendError(ctx, w, sr.Status, msg, err.Error())
		return
	}
	ctx.Log.Info.Print(string(ret))
	service.Benchmark(ctx, "full request")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sr.Status)
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
