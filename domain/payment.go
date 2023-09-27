package domain

import (
	"context"
	"database/sql"
	"mime/multipart"
)

type Payment struct {
	ID          string
	Name        string
	Description sql.NullString
	Image       string
	AuditInfo
}

type RequestCreatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
}

type RequestUpdatePayment struct {
	Name        string                `form:"name"`
	Description string                `form:"description"`
	Image       *multipart.FileHeader `form:"image"`
	ID          string
}

type ResponsePayment struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Image       string  `json:"image"`
}

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) error
	Update(ctx context.Context, payment *Payment) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]Payment, error)
	GetByID(ctx context.Context, id string) (*Payment, error)
	GetByName(ctx context.Context, name string) (*Payment, error)
}

type PaymentUsecase interface {
	Create(ctx context.Context, req *RequestCreatePayment) (*ResponsePayment, error)
	Update(ctx context.Context, req *RequestUpdatePayment) (*ResponsePayment, error)
	GetAll(ctx context.Context) (*[]ResponsePayment, error)
	Delete(ctx context.Context, id string) error
}
