package handlers

import (
	"encoding/json"
	"net/http"
	"order-service/src/infrastructure/interfaces"
	request "order-service/src/infrastructure/models/request"
	response "order-service/src/infrastructure/models/response"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	var model request.CreateOrderRequestModel
	err := json.NewDecoder(r.Body).Decode(&model)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	result, err := service.Create(model)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSONHttpOK(w, response.IdResponseModel{Id: result})
}

func Update(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	var model request.UpdateOrderRequestModel
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

func Delete(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	result, err := service.Delete(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, result, nil)
}

func Get(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
	}

	responseModel, err := service.GetById(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSONHttpOK(w, responseModel)
}

func GetAll(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
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

func GetByCustomerId(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	page, errPage := strconv.Atoi(r.URL.Query().Get("page"))
	if errPage != nil || page < 1 {
		page = 1
	}
	limit, errLimit := strconv.Atoi(r.URL.Query().Get("limit"))
	if errLimit != nil || limit < 1 {
		limit = 10
	}

	customerid, err := primitive.ObjectIDFromHex(r.URL.Query().Get("customerid"))
	if err != nil {
		Error(w, http.StatusBadRequest, nil, "customerid is required")
	}

	responseModel, err := service.GetByCustomerId(page, limit, customerid)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSONHttpOK(w, responseModel)
}

func ChangeStatus(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		Error(w, http.StatusBadRequest, err, err.Error())
	}

	status := r.URL.Query().Get("status")
	if len(status) < 1 {
		Error(w, http.StatusBadRequest, nil, "status is required")
	}

	result, err := service.ChangeStatus(id, status)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, result, nil)
}
