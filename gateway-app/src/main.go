package main

import (
	"api-gateway/src/config"
	"api-gateway/src/customer"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	r := gin.Default()

	r.Run(":5000")

	router := mux.NewRouter().StrictSlash(true)

	customer.RegisterRoutes(r, &c)

	log.Fatal(http.ListenAndServe(":5000", router))
}
