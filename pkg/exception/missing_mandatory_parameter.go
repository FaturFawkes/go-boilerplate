package exception

type MissingMandatoryParameter struct {
	message string
}

func (e MissingMandatoryParameter) Error() string {
	return e.message
}

func NewMissingMandatoryParameter() error {
	return &MissingMandatoryParameter{message: "Missing Mandatory Parameter"}
}
