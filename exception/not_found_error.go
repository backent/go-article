package exception

type ErrorNotFound struct {
	Message string
}

func (e *ErrorNotFound) Error() string {
	return e.Message
}

func NewErrorNotFound(message string) *ErrorNotFound {
	return &ErrorNotFound{Message: message}
}
