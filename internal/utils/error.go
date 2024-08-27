package utils

import "fmt"

type ErrorResponse struct {
	Error string `json:"error"`
}

func (err *ErrorResponse) ErrorMessage() string {
	return fmt.Sprintf("Error: %s", err.Error)
}
