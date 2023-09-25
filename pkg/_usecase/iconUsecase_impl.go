package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type IconUsecaseImpl struct {
	iconRepo  domain.IconRepository
	minioRepo domain.MinioRepo
}

func NewIconUsecaseImpl(
	iconRepo domain.IconRepository,
	minioRepo domain.MinioRepo,
) domain.IconUsecase {
	return &IconUsecaseImpl{
		iconRepo:  iconRepo,
		minioRepo: minioRepo,
	}
}

func (i *IconUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreateIcon) (*domain.ResponseIcon, error) {
	// TODO implement me
	panic("implement me")
}
