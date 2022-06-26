package customer

import (
	"api-gateway/src/config"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, c *config.Config) {

	r.HandleFunc("/customer/", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAll(w, r, customerService)
	}).Methods(http.MethodGet)

	r.HandleFunc("/customer/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Get(w, r, customerService)
	}).Methods(http.MethodGet)

	r.HandleFunc("/customer/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Create(w, r, customerService)
	}).Methods(http.MethodPost)

	r.HandleFunc("/customer/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Update(w, r, customerService)
	}).Methods(http.MethodPut)

	r.HandleFunc("/customer/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Delete(w, r, customerService)
	}).Methods(http.MethodDelete)

	r.HandleFunc("/customer/validate/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.Validate(w, r, customerService)
	}).Methods(http.MethodGet)
}
