package main

import (
	"customer-service/src/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.CreateCustomer).Methods(http.MethodPost)
	router.HandleFunc("/", handlers.UpdateCustomer).Methods(http.MethodPut)

	router.HandleFunc("/", handlers.GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/{id}", handlers.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/validate/{id", handlers.ValidationCheckCustomer).Methods(http.MethodGet)

	router.HandleFunc("/{id}", handlers.DeleteCustomer).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
