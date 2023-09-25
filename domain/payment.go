package domain

import (
	"context"
)

type Payment struct {
	AuditInfo
}

type RequestCreatePayment struct {
	ID string
}

type RequestUpdatePayment struct {
	ID string
}

type ResponsePayment struct {
	ID string
}

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
}

type PaymentUsecase interface {
	Create(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
}
