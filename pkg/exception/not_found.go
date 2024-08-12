package exception

type NotFound struct {
	message string
}

func (e NotFound) Error() string {
	return e.message
}

func NewNotFound(message string) error {
	return &NotFound{message: message}
}
