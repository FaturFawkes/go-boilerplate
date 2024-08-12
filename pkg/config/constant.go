package config

const (
	EnvHost           = "HOST"
	EnvPort           = "PORT"
	EnvMode           = "ENVIRONMENT_MODE"
	EnvDefaultTimeout = "DEFAULT_TIMEOUT"

	EnvJwtExpiry = "JWT_EXPIRY"
	EnvJwtSecret = "JWT_SECRET"
	EnvHeaderAuthorization = "Authorization"

	// DefaultTimeout will be the default timeout value (in seconds)
	DefaultTimeout = 10
)
