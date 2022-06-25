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
		handlers.CreateCustomer(w, r, customerService)
	}).Methods(http.MethodPost)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateCustomer(w, r, customerService)
	}).Methods(http.MethodPut)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllCustomer(w, r, customerService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetCustomer(w, r, customerService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/validate/{id", func(w http.ResponseWriter, r *http.Request) {
		handlers.ValidationCheckCustomer(w, r, customerService)
	}).Methods(http.MethodGet)

	router.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteCustomer(w, r, customerService)
	}).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
