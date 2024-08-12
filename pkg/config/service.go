package config

import (
	"os"
	"time"
)

type ServiceConfig struct {
	Host           string
	Port           string
	Mode           string
	DefaultTimeout time.Duration
}

func Service() ServiceConfig {
	return ServiceConfig{
		Host:           os.Getenv(EnvHost),
		Port:           os.Getenv(EnvPort),
		Mode:           os.Getenv(EnvMode),
		DefaultTimeout: NewTimeout().Second(),
	}
}
