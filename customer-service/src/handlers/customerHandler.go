package handlers

import (
	"customer-service/src/infrastructure/interfaces"
	request "customer-service/src/infrastructure/models/request"
	response "customer-service/src/infrastructure/models/response"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	var model request.CreateCustomerRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	result, err := service.Create(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSONHttpOK(w, response.IdResponseModel{Id: result})
}

func Update(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	var model request.UpdateCustomerRequestModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode((&model))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	result, err := service.Update(model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	JSON(w, result, nil)
}

func Delete(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	result, err := service.Delete(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, result, nil)
}

func Get(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
	}

	responseModel, err := service.GetById(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSONHttpOK(w, responseModel)
}

func GetAll(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	page, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	if errPage != nil || page < 1 {
		page = 1
	}
	limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))
	if errLimit != nil || limit < 1 {
		limit = 10
	}

	responseModel, err := service.GetAll(page, limit)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSONHttpOK(w, responseModel)
}

func Validate(w http.ResponseWriter, r *http.Request, service interfaces.ICustomerService) {
	vars := mux.Vars(r)
	customerId := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	isValid, err := service.IsValid(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, isValid, nil)
}
