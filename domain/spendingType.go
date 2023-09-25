package domain

import (
	"context"
)

type SpendingType struct {
	AuditInfo
}

type RequestCreateSpendingType struct {
	ID string
}

type RequestUpdateSpendingType struct {
	ID string
}

type ResponseSpendingType struct {
	ID string
}

type SpendingTypeRepository interface {
	Create(ctx context.Context, payment *SpendingType) error
}

type SpendingTypeUsecase interface {
	Create(ctx context.Context, req *RequestCreateSpendingType) (*ResponseSpendingType, error)
}
