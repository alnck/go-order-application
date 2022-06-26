package main

import (
	"log"
	"net/http"
	"order-service/src/handlers"
	"order-service/src/infrastructure"

	"github.com/gorilla/mux"
)

func main() {
	var orderService = infrastructure.NewOrderServiceResolve()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAll(w, r, orderService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Get(w, r, orderService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/GetByCustomerId/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetByCustomerId(w, r, orderService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Create(w, r, orderService)
	}).Methods(http.MethodPost)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Update(w, r, orderService)
	}).Methods(http.MethodPut)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.ChangeStatus(w, r, orderService)
	}).Methods(http.MethodPut)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Delete(w, r, orderService)
	}).Methods(http.MethodDelete)

	log.Fatal(http.ListenAndServe(":8082", router))
}
