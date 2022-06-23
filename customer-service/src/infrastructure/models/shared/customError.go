package models

type CustomError struct {
	Message string
}

func (err CustomError) NewCustomErr(message string) CustomError {
	return CustomError{Message: message}
}
