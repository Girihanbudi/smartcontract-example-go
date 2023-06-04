package stderror

import (
	"errors"
	"fmt"
)

type StdInvalidParams struct {
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type StdError struct {
	Error error `json:"-"`

	Status int    `json:"-"`                // return a http status
	Type   string `json:"type,omitempty"`   // return a documentation link of error explanation
	Title  string `json:"title,omitempty"`  // return a main description of error
	Errors []any  `json:"errors,omitempty"` // return a detail error param if any
}

func New(code int, msg ...any) StdError {
	var format string
	for i := range msg {
		if i == 0 {
			format = format + "%s"
		} else {
			format = format + " %s"
		}
	}

	return StdError{
		Status: code,
		Title:  fmt.Sprintf(format, msg...),
	}
}

func (e *StdError) New(link, title string, err ...any) *StdError {
	newError := *e

	newError.Type = link
	if title != "" {
		newError.Title = title
	}
	newError.AddErrors(err...)
	return &newError
}

func (e *StdError) SetError(err error) *StdError {
	newError := *e
	newError.Error = err
	return &newError
}

func (e *StdError) SetErrorFromString(message string) *StdError {
	newError := *e
	newError.Error = errors.New(message)
	return &newError
}

func (e *StdError) AddErrors(err ...any) *StdError {
	newError := *e
	newError.Errors = append(e.Errors, err...)
	return &newError
}
