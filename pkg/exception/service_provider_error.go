package exception

type ServiceProviderError struct {
	message string
}

func (e ServiceProviderError) Error() string {
	return e.message
}

func NewServiceProviderError() error {
	return &ServiceProviderError{message: "Service Provider Error"}
}
