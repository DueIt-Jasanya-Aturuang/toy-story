package domain

import (
	"context"
)

type Icon struct {
	ID    string
	Title string
	Icon  string
	AuditInfo
}

type RequestCreateIcon struct {
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type RequestUpdateIcon struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type ResponseIcon struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Icon  string `json:"icon"`
}

type IconRepository interface {
	Create(ctx context.Context, payment *Icon) error
	Update(ctx context.Context, payment *Icon) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) (*[]Icon, error)
	GetByID(ctx context.Context, id string) (*Icon, error)
	GetByTitle(ctx context.Context, title string) (*Icon, error)
}

type IconUsecase interface {
	Create(ctx context.Context, req *RequestCreateIcon) (*ResponseIcon, error)
	Update(ctx context.Context, req *RequestUpdateIcon) (*ResponseIcon, error)
	GetAll(ctx context.Context) (*[]ResponseIcon, error)
	Delete(ctx context.Context, id string) error
}
