package errs

// This allows for me to write a custom error and send it back with correct response it should have
// I then do not need to figure out if something happened out of my control or was in my control

import (
	"errors"
	"fmt"
	"net/http"
)

type ApiError struct {
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
	Msg        string `json:"msg"`
}

func NewError(code int, t string, msg string) ApiError {
	return ApiError{code, t, msg}
}

func (e ApiError) NewError(err error) ApiError {
	e.Msg = err.Error()
	return e
}

func (e ApiError) NewMessage(msg string) ApiError {
	e.Msg = msg
	return e
}

func (e ApiError) NewHttpStatusCode(code int) ApiError {
	e.StatusCode = code
	return e
}

func (e ApiError) NewType(t string) ApiError {
	e.Type = t
	return e
}

func (e ApiError) Error() string {
	return fmt.Sprintf("%s", e.Msg)
}

func IsApiError(err error) *ApiError {
	var cApiError *ApiError

	if errors.As(err, &cApiError) {
		return cApiError
	}
	return nil
}

var (
	UNAUTHORIZED = ApiError{
		http.StatusUnauthorized,
		"UNATHORIZED",
		"INVALID USERNAME OR PASSWORD",
	}

	LDAP_ERROR = ApiError{
		http.StatusInternalServerError,
		"LDAP_ERROR",
		"SOMETHING WENT WRONG WITH LDAP",
	}

	INVALID_BODY = ApiError{
		http.StatusBadRequest,
		"INVALID_BODY",
		"REQUEST BODY INVALID",
	}

	NOT_FOUND = ApiError{
		http.StatusNotFound,
		"NOT_FOUND",
		"RESUTLS WERE NOT FOUND FOR INPUT",
	}

	FILE_UNREACHABLE = ApiError{
		http.StatusInternalServerError,
		"FILE_NOT_FOUND",
		"COULD NOT REACH OR FIND FILE",
	}

	NOT_LOGGED_IN = ApiError{
		http.StatusUnauthorized,
		"NOT_LOGGED_IN",
		"AGENT NOT LOGGED IN",
	}

	REQUEST_ERROR = ApiError{
		http.StatusInternalServerError,
		"REQUEST_ERROR",
		"UNABLE TO REACH WEBSITE",
	}

	READ_FILE_ERROR = ApiError{
		http.StatusInternalServerError,
		"FILE_READ_ERROR",
		"FAILED TO READ FILE NEEDED TO COMPLETE REQUEST",
	}
)
