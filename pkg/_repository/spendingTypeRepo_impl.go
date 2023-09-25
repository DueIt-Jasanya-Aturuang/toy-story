package _repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type SpendingTypeRepositoryImpl struct {
	db *sql.DB
}

func NewSpendingTypeRepository(db *sql.DB) domain.SpendingTypeRepository {
	return &SpendingTypeRepositoryImpl{
		db: db,
	}
}

func (s *SpendingTypeRepositoryImpl) Create(ctx context.Context, req *domain.SpendingType) error {
	// TODO implement me
	panic("implement me")
}
