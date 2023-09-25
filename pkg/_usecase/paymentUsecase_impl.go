package _usecase

import (
	"context"

	"github.com/DueIt-Jasanya-Aturuang/toy-story/domain"
)

type PaymentUsecaseImpl struct {
	paymentRepo domain.PaymentRepository
	minioRepo   domain.MinioRepo
}

func NewPaymentUsecaseImpl(
	paymentRepo domain.PaymentRepository,
	minioRepo domain.MinioRepo,
) domain.PaymentUsecase {
	return &PaymentUsecaseImpl{
		paymentRepo: paymentRepo,
		minioRepo:   minioRepo,
	}
}

func (p *PaymentUsecaseImpl) Create(ctx context.Context, req *domain.RequestCreatePayment) (*domain.ResponsePayment, error) {
	// TODO implement me
	panic("implement me")
}
