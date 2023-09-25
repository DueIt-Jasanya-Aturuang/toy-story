package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type SpendingTypeUsecaseImpl struct {
	spendingTypeRepo domain.SpendingTypeRepository
}

func NewSpendingTypeUsecaseImpl(
	spendingTypeRepo domain.SpendingTypeRepository,
) domain.SpendingTypeUsecase {
	return &SpendingTypeUsecaseImpl{
		spendingTypeRepo: spendingTypeRepo,
	}
}

func (s *SpendingTypeUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateSpendingType) (*domain.ResponseSpendingType, error) {
	// TODO implement me
	panic("implement me")
}
