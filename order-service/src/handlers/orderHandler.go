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
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	result, err := service.Create(model)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, true, response.IdResponseModel{Id: result}, "")
}

func Update(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	var model request.UpdateOrderRequestModel
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode((&model))
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	result, err := service.Update(model)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, result, nil, "")
}

func Delete(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	result, err := service.Delete(id)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, result, nil, "")
}

func Get(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	responseModel, err := service.GetById(id)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, true, responseModel, "")
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
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, true, responseModel, "")
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

	vars := mux.Vars(r)
	customerid := vars["id"]

	id, err := primitive.ObjectIDFromHex(customerid)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	responseModel, err := service.GetByCustomerId(page, limit, id)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, true, responseModel, "")
}

func ChangeStatus(w http.ResponseWriter, r *http.Request, service interfaces.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
	}

	status := r.URL.Query().Get("status")
	if len(status) < 1 {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	result, err := service.ChangeStatus(id, status)
	if err != nil {
		JSON(w, http.StatusBadRequest, false, nil, err.Error())
		return
	}

	JSON(w, http.StatusOK, result, nil, "")
}
