package aggregate

import (
	"strings"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/model/entity"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Category struct {
	//TODO possibly make the items aggregates
	category *entity.Category
	itemIDs  []uuid.UUID
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
		itemIDs:  make([]uuid.UUID, 0),
	}, nil
}

func (c *Category) GetID() uuid.UUID {
	return c.category.ID
}
func (c *Category) SetID(id uuid.UUID) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	c.category.ID = id
	return nil
}

func (c *Category) GetName() string {
	return c.category.Name
}

func (c *Category) GetBudget() *valueobject.Budget {
	return c.category.Budget
}
func (c *Category) SetBudgetTotal(total float64) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if c.category.Budget == nil {
		//lazy initialise if budget does not exist
		c.category.Budget = &valueobject.Budget{}
	}
	c.category.Budget.Total = total
	c.category.Budget.Remaining = total - c.category.Budget.Spent
	return nil
}

func (c *Category) SetBudgetSpent(spent float64) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if c.category.Budget == nil {
		//lazy initialise if budget does not exist
		c.category.Budget = &valueobject.Budget{}
	}
	c.category.Budget.Spent = spent
	c.category.Budget.Remaining = c.category.Budget.Total - spent
	return nil
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

func (c *Category) AddItem(item *Item) error {
	if c.category == nil || c.category.Budget == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if item == nil {
		return ErrInvalidItem
	}
	for _, itemID := range c.itemIDs {
		if itemID == item.GetID() {
			return nil
		}
	}
	budget := item.GetBudget()
	c.category.Budget.Total += budget.Total
	c.category.Budget.Spent += budget.Spent
	c.category.Budget.Remaining = c.category.Budget.Total - c.category.Budget.Spent
	c.itemIDs = append(c.itemIDs, item.GetID())
	return nil
}

func (c *Category) GetItemIDs() []uuid.UUID {
	return c.itemIDs
}
func (c *Category) SetItemIDs(itemIDs []uuid.UUID) error {
	if c.category == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	c.itemIDs = itemIDs
	return nil
}

func (c *Category) RemoveItem(itemToRemove *Item) error {
	if c.category == nil || c.category.Budget == nil {
		//lazy initialise if category does not exist
		// c.category = &entity.Category{}
		return ErrUnInitialised
	}
	if itemToRemove == nil {
		return ErrInvalidItem
	}

	indexToRemove := -1
	for i, itemID := range c.itemIDs {
		if itemID == itemToRemove.GetID() {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return nil
	}
	budget := itemToRemove.GetBudget()
	c.category.Budget.Total -= budget.Total
	c.category.Budget.Spent -= budget.Spent
	c.category.Budget.Remaining = c.category.Budget.Total - c.category.Budget.Spent
	c.itemIDs = append(c.itemIDs[:indexToRemove], c.itemIDs[indexToRemove+1:]...)
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
