package usecase

import (
	"context"
	"golang_boilerplate/domain/repository"
	"golang_boilerplate/domain/usecase"

	"go.uber.org/zap"
)

type (
	UseCase struct {
		ctx    context.Context
		logger *zap.Logger

		repo repository.IRepository
	}

	Constructor struct {
		Ctx    context.Context
		Logger *zap.Logger

		Repo repository.IRepository
	}
)

func New(cs Constructor) usecase.IUseCase {
	return &UseCase{
		ctx:    cs.Ctx,
		logger: cs.Logger,

		repo: cs.Repo,
	}
}
