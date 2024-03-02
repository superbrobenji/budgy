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

type CategoryRepository interface {
	Get(uuid.UUID) (aggregate.Category, error)
	Add(aggregate.Category) error
	Update(aggregate.Category) error
	Delete(uuid.UUID) error
}
