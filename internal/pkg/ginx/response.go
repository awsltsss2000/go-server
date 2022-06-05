package ginx

import "go-server/pkg/errors"

func UnWrapResponseError(err error) *errors.InnerError {
	if v, ok := err.(*errors.InnerError); ok {
		return v
	}
	return nil
}

func WrapResponseError(code int, msg string) error {
	return &errors.InnerError{
		Code:    code,
		Message: msg,
	}
}

func Wrap400ResponseError(msg ...string) error {
	var m string
	if msg == nil {
		m = "Bad Request"
	} else {
		m = msg[0]
	}
	return WrapResponseError(400, m)
}

func Wrap500ResponseError() error {
	return WrapResponseError(500, "Internal Error")
}
