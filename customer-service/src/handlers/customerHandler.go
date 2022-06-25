package handlers

import (
	request "customer-service/src/infrastructure/models/request"
	"customer-service/src/services"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var model request.CreateCustomerRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err = services.CreateCustomer(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, http.StatusCreated, nil)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	var model request.UpdateCustonerRequestModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode((&model))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	//Todo go services valid
	err = services.UpdateCustomer(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	JSON(w, http.StatusOK, nil)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := uuid.Parse(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	//Todo go services valid
	services.DeleteCustomer(id)

	JSON(w, http.StatusOK, nil)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := uuid.Parse(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	//Todo go services valid
	responseModel, err := services.GetCustomer(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	JSON(w, http.StatusOK, responseModel)
}

func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	responseModel, err := services.GetAllCustomer()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, http.StatusOK, responseModel)
}

func ValidationCheckCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := uuid.Parse(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	//Todo go services valid
	err = services.ValidationCheckCustomer(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, http.StatusOK, nil)
}
