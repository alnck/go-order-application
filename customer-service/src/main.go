package main

import (
	"customer-service/src/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/customerservice", handlers.ResponseHandler(handlers.CreateCustomer)).Methods("POST")
	router.Handle("/customerservice/update", handlers.ResponseHandler(handlers.UpdateCustomer)).Methods("POST")

	router.Handle("/customerservice", handlers.ResponseHandler(handlers.GetAllCustomer)).Methods("GET")
	router.Handle("/customerservice/{customer_id}", handlers.ResponseHandler(handlers.GetCustomer)).Methods("GET")
	router.Handle("/customerservice/validate/:customer_id", handlers.ResponseHandler(handlers.ValidationCheckCustomer)).Methods("GET")

	router.Handle("/customerservice", handlers.ResponseHandler(handlers.DeleteCustomer)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", router))
}
