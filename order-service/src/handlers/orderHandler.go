package handlers

import (
	"encoding/json"
	"net/http"
	request "order-service/src/infrastructure/models/request"
	"order-service/src/services"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
	var model request.CreateOrderRequestModel
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

func Update(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
	var model request.UpdateOrderRequestModel
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

func Delete(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
	vars := mux.Vars(r)
	orderId := vars["id"]

	id, err := primitive.ObjectIDFromHex(orderId)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err = service.Delete(id)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}

	JSON(w, http.StatusOK, nil)
}

func Get(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
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
	JSON(w, http.StatusOK, responseModel)
}

func GetAll(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
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

func ChangeStatus(w http.ResponseWriter, r *http.Request, service services.IOrderService) {
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

	err = service.ChangeStatus(id, status)
	if err != nil {
		Error(w, http.StatusInternalServerError, err, err.Error())
	}
	JSON(w, http.StatusOK, nil)
}
