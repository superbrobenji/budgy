package aggregate

import (
	"strings"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/model/entity"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

type Category struct {
	//TODO possibly make the items aggregates
	category *entity.Category
	items    []*entity.Item
}

func NewCategory(name string) (Category, error) {
	nameValidationErr := nameValidation(name)
	if nameValidationErr != nil {
		return Category{}, nameValidationErr
	}
	budget := &valueobject.Budget{
		Total:     0,
		Spent:     0,
		Remaining: 0,
	}

	category := &entity.Category{
		Name:   name,
		ID:     uuid.New(),
		Budget: budget,
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

func (c *Category) GetBudget() *valueobject.Budget {
	return c.category.Budget
}

func (c *Category) SetName(name string) error {
	if c.category == nil || c.category.Budget == nil {
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
	if c.category == nil || c.category.Budget == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if item == nil {
		return ErrInvalidItem
	}
	c.category.Budget.Total += item.Budget.Total
	c.category.Budget.Spent += item.Budget.Spent
	c.category.Budget.Remaining = c.category.Budget.Total - c.category.Budget.Spent
	c.items = append(c.items, item)
	return nil
}

func (c *Category) GetItems() []*entity.Item {
	return c.items
}

func (c *Category) RemoveItem(itemToRemove *entity.Item) error {
	if c.category == nil || c.category.Budget == nil {
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
	c.category.Budget.Total -= itemToRemove.Budget.Total
	c.category.Budget.Spent -= itemToRemove.Budget.Spent
	c.category.Budget.Remaining = c.category.Budget.Total - c.category.Budget.Spent
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
