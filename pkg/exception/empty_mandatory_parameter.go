package exception

type EmptyMandatoryParameter struct {
	message string
}

func (e EmptyMandatoryParameter) Error() string {
	return e.message
}

func NewEmptyMandatoryParameter() error {
	return &EmptyMandatoryParameter{message: "Empty Mandatory Parameter"}
}
