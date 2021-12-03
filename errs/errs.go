package errs

import (
	"net/http"
)

type AppError struct {
	Code    int    `json:"code ,omitempty"`
	Message string `json:"message"`
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusNotFound,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewUserNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusOK,
	}
}

func NewRelationshipNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusOK,
	}
}

func (a AppError) AsMessage() *AppError {
	return &AppError{
		Message: a.Message,
	}
}
