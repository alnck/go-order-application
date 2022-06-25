package handlers

import (
	request "customer-service/src/infrastructure/models/request"
	"customer-service/src/services"
	"encoding/json"
	"net/http"
	"strconv"

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

	err = service.Delete(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, http.StatusOK, nil)
}

func Get(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
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
	JSON(w, http.StatusOK, responseModel)
}

func GetAll(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
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

	JSON(w, http.StatusOK, responseModel)
}

func Validate(w http.ResponseWriter, r *http.Request, service services.ICustomerService) {
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
	if isValid {
		JSON(w, http.StatusOK, nil)
	}
	JSON(w, http.StatusBadRequest, nil)
}
