package exception

type ServiceUnavailable struct {
	message string
}

func (e ServiceUnavailable) Error() string {
	return e.message
}

func NewServiceUnavailable(message string) error {
	return &ServiceUnavailable{message: message}
}
