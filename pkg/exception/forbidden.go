package exception

type Forbidden struct {
	message string
}

func (e Forbidden) Error() string {
	return e.message
}

func NewForbidden(message string) error {
	return &Forbidden{message: message}
}
