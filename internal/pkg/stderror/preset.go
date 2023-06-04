package stderror

import "net/http"

func ErrBadRequest() *StdError {
	err := New(http.StatusBadRequest, "Not valid request")
	return &err
}

func ErrInternalServerError() *StdError {
	err := New(http.StatusInternalServerError, "Internal server error")
	return &err
}

func ErrServiceUnavailable() *StdError {
	err := New(http.StatusServiceUnavailable, "Server is unavailable")
	return &err
}
