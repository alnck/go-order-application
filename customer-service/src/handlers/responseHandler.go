package handlers

import (
	response "customer-service/src/infrastructure/models/response"
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, success bool, src interface{}, errmsg string) {
	var responseModel = &response.ResponseModelwithData{
		Data:    src,
		Success: success,
		Message: errmsg,
	}

	if success {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		Respond(w, code, responseModel)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, responseModel)
}

func JSONWithoutData(w http.ResponseWriter, code int, success bool, errmsg string) {
	var responseModel = &response.BasicResponseModel{
		Success: success,
		Message: errmsg,
	}

	if success {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		Respond(w, code, responseModel)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	Respond(w, code, responseModel)
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
			JSON(w, http.StatusInternalServerError, false, nil, "invalid json")
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
			JSON(w, http.StatusInternalServerError, false, nil, "invalid json")
			return
		}
	}
	w.WriteHeader(code)
	w.Write(body)
}
