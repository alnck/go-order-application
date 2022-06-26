package handlers

import (
	"encoding/json"
	"net/http"
	response "order-service/src/infrastructure/models/response"
)

func Error(w http.ResponseWriter, code int, err error, msg string) {
	e := &response.ErrorResponseModel{
		Message: msg,
		Error:   err,
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, e)
}

func JSON(w http.ResponseWriter, success bool, src interface{}) {
	if success {
		JSONHttpOK(w, src)
		return
	}

	JSONHttpBadRequest(w, src)
}

func JSONHttpOK(w http.ResponseWriter, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, http.StatusOK, src)
}

func JSONHttpBadRequest(w http.ResponseWriter, src interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, http.StatusBadRequest, src)
}

func Respond(w http.ResponseWriter, code int, src interface{}) {
	var body []byte
	var err error

	if src == nil {
		w.WriteHeader(code)
		w.Write(body)
		return
	}

	switch s := src.(type) {
	case []byte:
		if !json.Valid(s) {
			Error(w, http.StatusInternalServerError, err, "invalid json")
			return
		}
		body = s
	case string:
		body = []byte(s)
	case *response.ErrorResponseModel, response.ErrorResponseModel:
		// avoid infinite loop
		if body, err = json.Marshal(src); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("{\"reason\":\"failed to parse json\"}"))
			return
		}
	default:
		if body, err = json.Marshal(src); err != nil {
			Error(w, http.StatusInternalServerError, err, "failed to parse json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}
