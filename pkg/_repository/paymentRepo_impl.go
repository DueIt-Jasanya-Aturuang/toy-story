package _repository

import (
	"context"
	"database/sql"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type PaymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepositoryImpl(db *sql.DB) domain.PaymentRepository {
	return &PaymentRepositoryImpl{
		db: db,
	}
}

func (p *PaymentRepositoryImpl) Create(ctx context.Context, payment *domain.Payment) error {
	// TODO implement me
	panic("implement me")
}
