package exception

type BadRequest struct {
	message string
}

func (e BadRequest) Error() string {
	return e.message
}

func NewBadRequest(message string) error {
	return &BadRequest{message: message}
}
