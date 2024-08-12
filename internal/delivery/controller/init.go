package controller

import (
	"context"
	"golang_boilerplate/domain/delivery"
	"golang_boilerplate/domain/usecase"

	"go.uber.org/zap"
)

type (
	DeliveryHandler struct {
		ctx    context.Context
		logger *zap.Logger

		useCase usecase.IUseCase
	}

	Constructor struct {
		Ctx    context.Context
		Logger *zap.Logger

		UseCase usecase.IUseCase
	}
)

func New(cs Constructor) delivery.IDelivery {
	return &DeliveryHandler{
		ctx:     cs.Ctx,
		logger:  cs.Logger,
		useCase: cs.UseCase,
	}
}
