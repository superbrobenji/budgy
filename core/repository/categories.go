package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrCategoryNotFound    = errors.New("category not found")
	ErrFailedToAddCategory = errors.New("failed to add category")
	ErrUpdateCategory      = errors.New("failed to update category")
	ErrDeleteCategory      = errors.New("failed to delete category")
)

type CategoryRepositoryWrite interface {
	CreateCategory(aggregate.Category) error
	UpdateCategory(aggregate.Category) error
	DeleteCategory(uuid.UUID) error
}

type CategoryRepositoryRead interface {
	GetCategoryByID(uuid.UUID) (aggregate.Category, error)
	GetCategoryByItemID(uuid.UUID) (aggregate.Category, error)
	GetCategoriesByUserID(string) (*[]*aggregate.Category, error)
}
