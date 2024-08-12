package exception

type InvalidMandatoryParameter struct {
	message string
}

func (e InvalidMandatoryParameter) Error() string {
	return e.message
}

func NewInvalidMandatoryParameter() error {
	return &InvalidMandatoryParameter{message: "Invalid Mandatory Parameter"}
}
