package main

import (
	"context"
	"golang_boilerplate/internal/delivery"
	"golang_boilerplate/internal/delivery/controller"
	"golang_boilerplate/internal/delivery/middleware"
	"golang_boilerplate/internal/repository"
	"golang_boilerplate/internal/usecase"
	"golang_boilerplate/pkg/common"
	"golang_boilerplate/pkg/common/gin_engine"
	"golang_boilerplate/pkg/config"

	"go.uber.org/zap"
)

func main() {
	// initialize service configuration
	config.Environment()
	logger := config.Logger()
	ctx := context.Background()

	engine, err := gin_engine.New(gin_engine.Constructor{
		Ctx:                       ctx,
		Logger:                    logger,
		Host:                      config.Service().Host,
		Port:                      config.Service().Port,
		Environment:               config.Service().Mode,
		Timeout:                   config.Service().DefaultTimeout,
		DisableDefaultHealthCheck: false,
	})

	if err != nil {
		logger.Warn("failed to initialize gin engine", zap.Error(err))
	}

	controller := controller.New(controller.Constructor{
		Ctx:    ctx,
		Logger: logger,
		UseCase: usecase.New(usecase.Constructor{
			Ctx:    ctx,
			Logger: logger,
			Repo: repository.New(repository.Constructor{
				Ctx: ctx,
				BaseRepo: common.NewBaseRepository(
					nil,
					nil,
					nil,
				),
				Logger: logger,
			}),
		}),
	})

	mw := middleware.New()

	delivery.RegisterRouting(engine, controller, mw)

	engine.Start()
}
