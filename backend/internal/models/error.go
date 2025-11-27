package models

type Error struct {
	HttpStatus int
	ErrorType  string `json:"error_type"`
	Error      error  `json:"error"`
}

func NewError(status int, errtype string, err error) *Error {
	return &Error{HttpStatus: status, ErrorType: errtype, Error: err}
}
