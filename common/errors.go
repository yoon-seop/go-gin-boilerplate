package common

import "net/http"

type APIError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	RequiredFieldMissingError = &APIError{
		Code:    http.StatusBadRequest,
		Message: "required fields missing",
	}
	EmailAlreadyExistsError = &APIError{Code: http.StatusConflict, Message: "email already exists"}
	UserNotFoundError       = &APIError{Code: http.StatusNotFound, Message: "user not found"}
)
