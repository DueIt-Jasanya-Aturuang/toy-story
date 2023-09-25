package _repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type IconRepositoryImpl struct {
	db *sql.DB
}

func NewIconRepositoryImpl(db *sql.DB) domain.IconRepository {
	return &IconRepositoryImpl{
		db: db,
	}
}

func (i *IconRepositoryImpl) Create(ctx context.Context, payment *domain.Icon) error {
	// TODO implement me
	panic("implement me")
}
