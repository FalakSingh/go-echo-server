package appError

import "fmt"

type AppError struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Type    string `json:"type"`
	Err     error  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("Status: %d, Message: %s, Error: %s", e.Status, e.Message, e.Err.Error())
	}
	return fmt.Sprintf("Status: %d, Message: %s", e.Status, e.Message)
}

func New(status int, message string, err error) *AppError {
	return &AppError{
		Success: false,
		Status:  status,
		Message: message,
		Err:     err,
	}
}

func NewBadRequest(message string) *AppError {
	return &AppError{
		Success: false,
		Status:  400,
		Message: message,
		Err:     nil,
	}
}

func NewNotFound(message string) *AppError {
	return &AppError{
		Success: false,
		Status:  404,
		Message: message,
		Err:     nil,
	}
}

func NewInternalServerError(message string, err error) *AppError {
	return &AppError{
		Success: false,
		Status:  500,
		Message: message,
		Err:     err,
	}
}

func NewUnauthorized(message string) *AppError {
	return &AppError{
		Success: false,
		Status:  401,
		Message: message,
		Err:     nil,
	}
}
