package restface

import "net/http"

type ApiError struct {
	StatusCode int      `json:"status"`
	Message    string   `json:"message"`
	Errors     []string `json:"cause,omitempty"`
}

func (e *ApiError) Error() string {
	return e.Message
}

func (e *ApiError) Code() int {
	return e.StatusCode
}

func (e *ApiError) Cause() []string {
	return e.Errors
}

func InternalError(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Error",
		Errors:     err,
	}
}

func NotFound(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusNotFound,
		Message:    "Bad Request",
		Errors:     err,
	}
}

func BadRequest(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusBadRequest,
		Message:    "Bad request",
		Errors:     err,
	}
}

func Forbidden(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusForbidden,
		Message:    "Forbidden",
		Errors:     err,
	}
}

func Unauthorized(err ...string) *ApiError {
	return &ApiError{
		StatusCode: http.StatusUnauthorized,
		Message:    "Unauthorized",
		Errors:     err,
	}
}
