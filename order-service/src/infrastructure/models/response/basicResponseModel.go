package model

type BasicResponseModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
