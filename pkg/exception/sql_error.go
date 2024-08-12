package exception

type SQLError struct {
	message string
}

func (e SQLError) Error() string {
	return e.message
}

func NewSQLError() error {
	return &SQLError{message: "SQL Error"}
}
