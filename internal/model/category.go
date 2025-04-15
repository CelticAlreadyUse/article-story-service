package model

import (
	"context"
	"time"
)

type CategoriesRepository interface {
	Create(ctx context.Context, body Category) (Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, Body Category) (*Category, error)
	GetAll(ctx context.Context, params CategoryParams) ([]*Category, error)
	GetAllCategoriesByIds(ctx context.Context, id []int64) ([]*Category, error)
}
type CategoriesUsecases interface {
	Create(ctx context.Context, body Category) (Category, error)
	GetByID(ctx context.Context, id int64) (*Category, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, Body Category) (*Category, error)
	GetAll(ctx context.Context, params CategoryParams) ([]*Category, error)
	GetAllCategoriesByIds(ctx context.Context, id []int64) ([]*Category, error)
}

type Category struct {
	ID        int64     `json:"id" gorm:"id"`
	Slug      string    `json:"slug" gorm:"slug"`
	Name      string    `json:"name" validate:"required"  gorm:"name"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"updaetd_at"`
}
