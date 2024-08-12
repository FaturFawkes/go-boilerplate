package exception

type ServiceProviderResponseTimeout struct {
	message string
}

func (e ServiceProviderResponseTimeout) Error() string {
	return e.message
}

func NewServiceProviderResponseTimeout() error {
	return &ServiceProviderResponseTimeout{message: "Service Provider Response Timeout"}
}
