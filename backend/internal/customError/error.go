package customError

import (
	"fmt"
	"net/http"
)

type Error struct {
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
	Msg        string `json:"msg"`
}

func NewError(code int, t string, msg string) Error {
	return Error{code, t, msg}
}

func (e Error) NewError(err error) Error {
	e.Msg = err.Error()
	return e
}

func (e Error) NewMessage(msg string) Error {
	e.Msg = msg
	return e
}

func (e Error) NewHttpStatusCode(code int) Error {
	e.StatusCode = code
	return e
}

func (e Error) NewType(t string) Error {
	e.Type = t
	return e
}

func (e Error) GetErrorValue() error {
	return fmt.Errorf("%s", e.Msg)
}

var (
	UNAUTHORIZED = Error{
		http.StatusUnauthorized,
		"UNATHORIZED",
		"INVALID USERNAME OR PASSWORD",
	}

	LDAP_ERROR = Error{
		http.StatusInternalServerError,
		"LDAP_ERROR",
		"SOMETHING WENT WRONG WITH LDAP",
	}

	INVALID_BODY = Error{
		http.StatusBadRequest,
		"INVALID_BODY",
		"REQUEST BODY INVALID",
	}

	NOT_FOUND = Error{
		http.StatusNotFound,
		"NOT_FOUND",
		"RESUTLS WERE NOT FOUND FOR INPUT",
	}

	FILE_UNREACHABLE = Error{
		http.StatusInternalServerError,
		"FILE_NOT_FOUND",
		"COULD NOT REACH OR FIND FILE",
	}

	NOT_LOGGED_IN = Error{
		http.StatusUnauthorized,
		"NOT_LOGGED_IN",
		"AGENT NOT LOGGED IN",
	}

	REQUEST_ERROR = Error{
		http.StatusInternalServerError,
		"REQUEST_ERROR",
		"UNABLE TO REACH WEBSITE",
	}

	READ_FILE_ERROR = Error{
		http.StatusInternalServerError,
		"FILE_READ_ERROR",
		"FAILED TO READ FILE NEEDED TO COMPLETE REQUEST",
	}
)
