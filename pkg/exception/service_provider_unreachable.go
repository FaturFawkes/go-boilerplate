package exception

type ServiceProviderUnreachable struct {
	message string
}

func (e ServiceProviderUnreachable) Error() string {
	return e.message
}

func NewServiceProviderUnreachable() error {
	return &ServiceProviderUnreachable{message: "Service Provider Unreachable"}
}
