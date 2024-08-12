package config

import (
	"os"
)

type JwtConfig struct {
	JwtExpiry string
	JwtSecret string
}

func Jwt() JwtConfig {
	return JwtConfig{
		JwtExpiry: os.Getenv(EnvJwtExpiry),
		JwtSecret: os.Getenv(EnvJwtSecret),
	}
}
