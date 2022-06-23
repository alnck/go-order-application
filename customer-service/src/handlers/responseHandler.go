package handlers

import (
	models "customer-service/src/infrastructure/models/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResponseHandler func(http.ResponseWriter, *http.Request) (interface{}, *models.CustomError)

func (fn ResponseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := fn(w, r)

	if err != nil {
		responseJson(w, err, http.StatusBadRequest)
		return
	}

	responseJson(w, response, http.StatusOK)
}

func responseJson(w http.ResponseWriter, response interface{}, httpStatus int) {
	jsonResponse, jsonError := json.Marshal(response)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)
	w.Write(jsonResponse)
}
