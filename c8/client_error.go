package c8

import (
	"fmt"
)

// APIError return the api error
type APIError struct {
	Code         int    `json:"code"`
	IsError      bool   `json:"error"`
	ErrorMessage string `json:"errorMessage"`
	ErrorNum     int    `json:"errorNum"`
}

// Error return the error message
func (e APIError) Error() string {
	return fmt.Sprintf("APIError: Code=%v, Number=%v,  Error Message=%v", e.Code, e.ErrorNum, e.ErrorMessage)
}
