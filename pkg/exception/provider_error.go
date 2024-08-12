package exception

type ProviderError struct {
	message string
}

// Error will return error message
func (e ProviderError) Error() string {
	return e.message
}

// NewProviderError will return error but with "ProviderError" type
// it should be used when you've encountered error with 3rd party provider
func NewProviderError(message string) error {
	return &ProviderError{message: message}
}
