package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"strconv"

	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/joho/godotenv"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func ConvertInt(env string) int {
	v, _ := strconv.Atoi(os.Getenv(env))
	return v
}

func Logger() *zap.Logger {
	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	// logger = logger.With(zap.String("app", "scm.salt.id/salt-and-pepper/telkomsel-digihub/telkomsel-api-marketplace-microservice")).With(zap.String("environment", "prod"))

	return logger
}

func GetZapObserver(level zapcore.LevelEnabler) (*zap.Logger, *observer.ObservedLogs) {
	observedZapCore, observedLogs := observer.New(level)
	zpLogger := zap.New(observedZapCore)

	return zpLogger, observedLogs
}

func Environment() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load()
		if err != nil {
			Logger().Error("error while loading .env file", zap.Error(err))
		}
	} else {
		Logger().Warn("running service without configuration from .env")
	}
}

func CleanEnvironment[T any](path string) (*T, error) {
	var target *T
	if err := cleanenv.ReadConfig(path, &target); err != nil {
		return nil, err
	}
	return target, nil
}
