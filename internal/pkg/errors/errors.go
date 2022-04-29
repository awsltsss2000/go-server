package errors

type InnerError struct {
	Code    int
	Message string
}

func (e *InnerError) Error() string {
	return e.Message
}
