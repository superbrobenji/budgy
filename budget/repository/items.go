package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/aggregate"
)

var (
	ErrItemNotFound    = errors.New("item not found")
	ErrFailedToAddItem = errors.New("failed to add item")
	ErrUpdateItem      = errors.New("failed to update item")
	ErrDeleteItem      = errors.New("failed to delete item")
)

type ItemRepository interface {
	Get(uuid.UUID) (aggregate.Category, error)
	Add(aggregate.Category) error
	Update(aggregate.Category) error
	Delete(uuid.UUID) error
}
