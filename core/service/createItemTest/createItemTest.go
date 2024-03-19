package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/repository"
	datastore "github.com/superbrobenji/budgy/infrastructure/persistence/dataStore/item"
)

type ItemConfiguration func(is *ItemService) error

type ItemService struct {
	items repository.ItemRepository
}

// NewItemService creates a new ItemService takes in a variadic number of ItemConfiguration functions
func NewItemService(cfgs ...ItemConfiguration) (*ItemService, error) {
	is := &ItemService{}
    for _, cfg := range cfgs {
        err := cfg(is)
        if err != nil {
            return nil, err
        }
    }
	withDynamoItemRepository()(is)
	return is, nil
}

// example of a configuration function
func withItemRepository(cr repository.ItemRepository) ItemConfiguration {
	return func(is *ItemService) error {
		is.items = cr
		return nil
	}
}

func withDynamoItemRepository() ItemConfiguration {
	cr := datastore.NewDynamoItemRepository()
	return withItemRepository(cr)
}

func (i *ItemService) CreateItem(itemID uuid.UUID) (*aggregate.Item, error) {
	item, err := aggregate.NewItem("test", 100, itemID)

	if err != nil {
		return nil, err
	}
	error := i.items.CreateItem(&item)
	if error != nil {
		return nil, error
	}
	log.Printf("customer: %v", item)
	return &item, nil
}
