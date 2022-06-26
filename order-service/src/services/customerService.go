package services

import (
	"encoding/json"
	"net/http"
	model "order-service/src/infrastructure/models/response"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CustomerIsValid(id primitive.ObjectID) bool {
	resp, err := http.Get("http://localhost:5001/validate/" + id.Hex())

	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}

	var response model.ResponseModelwithData

	json.NewDecoder(resp.Body).Decode(&response)

	return response.Success
}
