package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrItemNotFound    = errors.New("item not found")
	ErrFailedToAddItem = errors.New("failed to add item")
	ErrUpdateItem      = errors.New("failed to update item")
	ErrDeleteItem      = errors.New("failed to delete item")
)

type ItemRepository interface {
	GetItemsByCategoryID(uuid.UUID) (*[]*aggregate.Item, error)
	GetItemByID(uuid.UUID) (*aggregate.Item, error)
	CreateItem(*aggregate.Item) error
	DeleteItem(uuid.UUID) error
	UpdateItem(*aggregate.Item) error
}
