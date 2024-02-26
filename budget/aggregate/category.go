package aggregate

import (
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/model/entity"
)

var (
	ErrInvalidCategory       = errors.New("a category must have a valid name. 'Income' and 'Pre-Income Deductions' are reserved names")
	ErrReservedNameSpace     = errors.New("'Income' and 'Pre-Income Deductions' are reserved names")
	ErrUnInitialisedCategory = errors.New("a category must first be initialised")
)

type Category struct {
	category *entity.Category
	items    []*entity.Item
}

func NewCategory(name string) (Category, error) {
	if name == "" || name == "Income" || name == "Pre-Income Deductions" {
		lowercaseName := strings.ToLower(name)
		if lowercaseName == "income" || lowercaseName == "pre-income deductions" {
			return Category{}, ErrReservedNameSpace
		}
		return Category{}, ErrInvalidCategory
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
		return ErrUnInitialisedCategory
	}
	c.category.Name = name
	return nil
}

func (c *Category) AddItem(item *entity.Item) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialisedCategory
	}
	c.items = append(c.items, item)
	return nil
}

func (c *Category) GetItems() []*entity.Item {
	return c.items
}

func (c *Category) DeleteItem(itemToRemove *entity.Item) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialisedCategory
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
