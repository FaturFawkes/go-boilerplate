package exception

type Internal struct {
	message string
}

func (e Internal) Error() string {
	return e.message
}

func NewInternal(message string) error {
	return &Internal{message: message}
}
