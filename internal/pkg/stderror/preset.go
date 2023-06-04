package stderror

import "net/http"

var (
	ErrBadRequest          = New(http.StatusBadRequest, "Not valid request")
	ErrInternalServerError = New(http.StatusBadRequest, "Internal server error")
	ErrServiceUnavailable  = New(http.StatusServiceUnavailable, "Server is unavailable")
)
