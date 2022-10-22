package httputil

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	error1 := NewError(400, "INVALID_EMAIL", errors.New("email is invalid"))
	assert.Equal(t, Error{
		StatusCode: 400,
		ErrorCode:  "INVALID_EMAIL",
		Message:    "email is invalid",
	}, error1)
}

func TestNewError400(t *testing.T) {
	statusCode1, error1 := NewError400("INVALID_PASSWORD", "password is invalid")
	assert.Equal(t, 400, statusCode1)
	assert.Equal(t, Error{
		StatusCode: 400,
		ErrorCode:  "INVALID_PASSWORD",
		Message:    "password is invalid",
	}, error1)
}

func TestNewError404(t *testing.T) {
	statusCode1, error1 := NewError404("USER_NOT_FOUND", "test not found")
	assert.Equal(t, 404, statusCode1)
	assert.Equal(t, Error{
		StatusCode: 404,
		ErrorCode:  "USER_NOT_FOUND",
		Message:    "test not found",
	}, error1)

	statusCode2, error2 := NewError404("USER_NOT_FOUND", "")
	assert.Equal(t, 404, statusCode2)
	assert.Equal(t, Error{
		StatusCode: 404,
		ErrorCode:  "USER_NOT_FOUND",
		Message:    "not found",
	}, error2)
}

func TestNewError500(t *testing.T) {
	statusCode1, error1 := NewError500("INTERNAL_SERVER_ERROR", "internal test error")
	assert.Equal(t, 500, statusCode1)
	assert.Equal(t, Error{
		StatusCode: 500,
		ErrorCode:  "INTERNAL_SERVER_ERROR",
		Message:    "internal test error",
	}, error1)

	statusCode2, error2 := NewError500("INTERNAL_SERVER_ERROR", "")
	assert.Equal(t, 500, statusCode2)
	assert.Equal(t, Error{
		StatusCode: 500,
		ErrorCode:  "INTERNAL_SERVER_ERROR",
		Message:    "internal server error",
	}, error2)
}
