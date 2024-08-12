package exception

type Services struct {
	message string
}

func (e Services) Error() string {
	return e.message
}

func NewServices(message string) error {
	return &Services{message: message}
}
