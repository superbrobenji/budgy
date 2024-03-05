package aggregate

import (
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/model/entity"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Item struct {
	//TODO possibly make the transations aggregates
	item           *entity.Item
	transactionIDs []uuid.UUID
	categoryID     uuid.UUID
}

func NewItem(name string, amount float64, categoryID uuid.UUID) (Item, error) {
	if amount < 0 {
		return Item{}, ErrInvalidAmount
	}

	item := &entity.Item{
		Name: name,
		ID:   uuid.New(),
		Budget: &valueobject.Budget{
			Total:     amount,
			Spent:     0,
			Remaining: amount,
		},
	}
	return Item{
		item:           item,
		transactionIDs: make([]uuid.UUID, 0),
		categoryID:     categoryID,
	}, nil
}

func (i *Item) GetID() uuid.UUID {
	return i.item.ID
}
func (i *Item) SetID(id uuid.UUID) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	i.item.ID = id
	return nil
}
func (i *Item) GetCategoryID() uuid.UUID {
	return i.categoryID
}
func (i *Item) SetCategoryID(category *Category) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	i.categoryID = category.GetID()
	return nil
}

func (i *Item) GetName() string {
	return i.item.Name
}

func (i *Item) SetName(name string) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if name == "" {
		return ErrInvalidName
	}
	i.item.Name = name
	return nil
}

func (i *Item) GetBudget() *valueobject.Budget {
	return i.item.Budget
}

func (i *Item) SetBudgetTotal(amount float64) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if i.item.Budget == nil {
		i.item.Budget = &valueobject.Budget{}
	}
	if amount < 0 {
		return ErrInvalidAmount
	}
	i.item.Budget.Total = amount
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	return nil
}
func (i *Item) SetBudgetSpent(amount float64) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if i.item.Budget == nil {
		i.item.Budget = &valueobject.Budget{}
	}
	if amount < 0 {
		return ErrInvalidAmount
	}
	i.item.Budget.Spent = amount
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	return nil
}

func (i *Item) GetTransactionIDs() []uuid.UUID {
	return i.transactionIDs
}
func (i *Item) SetTransactionIDs(transactionIDs []uuid.UUID) error {
	if i.item == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	i.transactionIDs = transactionIDs
	return nil
}
func (i *Item) AddTransaction(transaction *Transaction) error {
	if i.item == nil || i.item.Budget == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transaction == nil {
		return ErrInvalidTransaction
	}
	for _, transactionID := range i.transactionIDs {
		if transactionID == transaction.GetID() {
			return nil
		}
	}
	i.item.Budget.Spent += transaction.GetAmount()
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	i.transactionIDs = append(i.transactionIDs, transaction.GetID())
	return nil
}

func (i *Item) RemoveTransation(transactionToRemove *Transaction) error {
	if i.item == nil || i.item.Budget == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transactionToRemove == nil {
		return ErrInvalidTransaction
	}

	indexToRemove := -1
	for i, item := range i.transactionIDs {
		if item == transactionToRemove.GetID() {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return nil
	}
	i.item.Budget.Spent -= transactionToRemove.GetAmount()
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	i.transactionIDs = append(i.transactionIDs[:indexToRemove], i.transactionIDs[indexToRemove+1:]...)
	return nil
}
