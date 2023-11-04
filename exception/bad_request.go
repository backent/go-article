package exception

type BadRequest struct {
	Message string
}

func (e *BadRequest) Error() string {
	return e.Message
}

func NewBadRequest(message string) BadRequest {
	return BadRequest{Message: message}
}
