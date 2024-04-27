package services

import (
	"log"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/repository"
	datastore "github.com/superbrobenji/budgy/infrastructure/persistence/dataStore/item"
)

type IncomeConfiguration func(is *IncomeService) error

type IncomeService struct {
	items repository.ItemRepositoryWrite
}

// NewItemService creates a new IncomeService takes in a variadic number of IncomeConfiguration functions
func NewIncomeService(cfgs ...IncomeConfiguration) (*IncomeService, error) {
	is := &IncomeService{}
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
func withItemRepositoryWrite(cr repository.ItemRepositoryWrite) IncomeConfiguration {
	return func(is *IncomeService) error {
		is.items = cr
		return nil
	}
}

func withDynamoItemRepository() IncomeConfiguration {
	cr := datastore.NewDynamoItemRepository()
	return withItemRepositoryWrite(cr)
}

func (i *IncomeService) CreateItem(itemID uuid.UUID) (*aggregate.Item, error) {
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
