package exception

type Unauthorized struct {
	message string
}

func (e Unauthorized) Error() string {
	return e.message
}

func NewUnauthorized(message string) error {
	return &Unauthorized{message: message}
}
