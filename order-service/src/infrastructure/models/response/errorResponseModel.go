package model

type ErrorResponseModel struct {
	Message string `json:"reason"`
	Error   error  `json:"-"`
}
