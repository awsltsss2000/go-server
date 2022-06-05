package errors

type InnerError struct {
	Code    int
	Message string
}

func (e *InnerError) Error() string {
	return e.Message
}

func New(code int, msg string) *InnerError {
	return &InnerError{Code: code, Message: msg}
}
