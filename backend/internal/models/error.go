package models

type Error struct {
	HttpStatus int
	ErrorType  string `json:"error_type"`
	Error      string `json:"error"`
}

func NewError(status int, errtype, err string) *Error {
	return &Error{HttpStatus: status, ErrorType: errtype, Error: err}
}
