package models

type CustomError struct {
	Message string
}

func NewCustomErr(message string) *CustomError {
	return &CustomError{Message: message}
}
