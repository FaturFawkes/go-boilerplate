package repository

import (
	"context"
	"golang_boilerplate/domain/repository"
	"golang_boilerplate/pkg/common"

	"go.uber.org/zap"
)

type (
	Repository struct {
		ctx    context.Context
		base   *common.BaseRepository
		logger *zap.Logger
	}

	Constructor struct {
		Ctx      context.Context
		BaseRepo *common.BaseRepository
		Logger   *zap.Logger
	}
)

func New(cs Constructor) repository.IRepository {
	return &Repository{
		ctx:    cs.Ctx,
		base:   cs.BaseRepo,
		logger: cs.Logger,
	}
}
