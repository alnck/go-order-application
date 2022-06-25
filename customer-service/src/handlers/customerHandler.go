package handlers

import (
	request "customer-service/src/infrastructure/models/request"
	"customer-service/src/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	var model request.CreateCustomerRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	err = service.Create(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, http.StatusOK, nil)
}

func Update(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	var model request.UpdateCustonerRequestModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode((&model))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	//Todo go services valid
	err = service.Update(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	JSON(w, http.StatusOK, nil)
}

func Delete(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	//Todo go services valid
	service.Delete(id)

	JSON(w, http.StatusOK, nil)
}

func Get(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	//Todo go services valid
	responseModel, err := service.GetById(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	JSON(w, http.StatusOK, responseModel)
}

func GetAll(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	responseModel, err := service.GetAll()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, http.StatusOK, responseModel)
}

func Validate(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	//Todo go services valid
	isValid, err := service.IsValid(id)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}
	if isValid {
		JSON(w, http.StatusOK, nil)
	}
	JSON(w, http.StatusBadRequest, nil)
}
