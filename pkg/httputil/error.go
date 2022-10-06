package httputil

import (
	"errors"
	"net/http"
)

const Message404 = "not found"
const Message500 = "internal server error"

type Error struct {
	StatusCode int    `json:"statusCode"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
}

func NewError(
	statucCode int,
	errorCode string,
	err error,
) Error {
	return Error{
		StatusCode: statucCode,
		ErrorCode:  errorCode,
		Message:    err.Error(),
	}
}

func NewError400(errorCode string, msg string) (int, Error) {
	return 400, NewError(400, errorCode, errors.New(msg))
}

func NewError404(errorCode string, msg string) (int, Error) {
	if msg == "" {
		msg = Message404
	}
	return http.StatusNotFound,
		NewError(http.StatusNotFound, errorCode, errors.New(msg))
}

func NewError500(errorCode string, msg string) (int, Error) {
	if msg == "" {
		msg = Message500
	}
	return http.StatusInternalServerError,
		NewError(http.StatusInternalServerError, errorCode, errors.New(msg))
}
