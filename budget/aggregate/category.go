package aggregate

import (
	"strings"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/model/entity"
)

type Category struct {
    //TODO possibly make the items aggregates
	category *entity.Category
	items    []*entity.Item
}

// TODO add all the budget functions
func NewCategory(name string) (Category, error) {
	nameValidationErr := nameValidation(name)
	if nameValidationErr != nil {
		return Category{}, nameValidationErr
	}

	category := &entity.Category{
		Name: name,
		ID:   uuid.New(),
	}
	return Category{
		category: category,
		items:    make([]*entity.Item, 0),
	}, nil
}

func (c *Category) GetID() uuid.UUID {
	return c.category.ID
}

func (c *Category) GetName() string {
	return c.category.Name
}

func (c *Category) SetName(name string) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	nameValidationErr := nameValidation(name)
	if nameValidationErr != nil {
		return nameValidationErr
	}
	c.category.Name = name
	return nil
}

func (c *Category) AddItem(item *entity.Item) error {
	//TODO update the budget based on amount on transaction
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if item == nil {
		return ErrInvalidItem
	}
	c.items = append(c.items, item)
	return nil
}

func (c *Category) GetItems() []*entity.Item {
	return c.items
}

func (c *Category) RemoveItem(itemToRemove *entity.Item) error {
	//TODO update the budget based on amount on transaction
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if itemToRemove == nil {
		return ErrInvalidItem
	}

	indexToRemove := -1
	for i, item := range c.items {
		if item == itemToRemove {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return nil
	}
	c.items = append(c.items[:indexToRemove], c.items[indexToRemove+1:]...)
	return nil
}

func nameValidation(name string) error {
	if name == "" {
		return ErrInvalidName
	}
	lowercaseName := strings.ToLower(name)
	if lowercaseName == "income" || lowercaseName == "pre-income deductions" {
		return ErrReservedCategoryNameSpace
	}
	return nil
}
