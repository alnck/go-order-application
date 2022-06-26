package main

import (
	"customer-service/src/handlers"
	"customer-service/src/infrastructure"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var customerService = infrastructure.NewCustomerServiceResolve()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAll(w, r, customerService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Get(w, r, customerService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Create(w, r, customerService)
	}).Methods(http.MethodPost)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Update(w, r, customerService)
	}).Methods(http.MethodPut)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Delete(w, r, customerService)
	}).Methods(http.MethodDelete)

	router.HandleFunc("/validate/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Validate(w, r, customerService)
	}).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":5001", router))
}
