package config

import (
	"os"
	"strconv"
	"time"
)

type (
	Timeout struct {
		duration int
	}
)

// NewTimeout will load default timeout from `DEFAULT_TIMEOUT`
// environment, as default if it error or not set it will
// use "10" as default value
func NewTimeout() *Timeout {
	timeout := &Timeout{DefaultTimeout}
	if value, exists := os.LookupEnv(EnvDefaultTimeout); exists && value != "" {
		strTimeout, _ := strconv.Atoi(value)
		timeout.duration = strTimeout
	}

	return timeout
}

// Second will return timeout duration in seconds
func (tm *Timeout) Second() time.Duration {
	return time.Duration(tm.duration) * time.Second
}
