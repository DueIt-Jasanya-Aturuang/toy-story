package domain

import (
	"context"
)

type Icon struct {
	AuditInfo
}

type RequestCreateIcon struct {
	ID string
}

type RequestUpdateIcon struct {
	ID string
}

type ResponseIcon struct {
	ID string
}

type IconRepository interface {
	Create(ctx context.Context, payment *Icon) error
}

type IconUsecase interface {
	Create(ctx context.Context, req *RequestCreateIcon) (*ResponseIcon, error)
}
