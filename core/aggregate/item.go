package aggregate

import (
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/model/entity"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Item struct {
	//TODO possibly make the transations aggregates
	item         *entity.Item
	transactions []*valueobject.Transaction
}

func NewItem(name string, amount float64) (Item, error) {
	if name == "" {
		return Item{}, ErrInvalidName
	}
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
		item:         item,
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}

func (i *Item) GetID() uuid.UUID {
	return i.item.ID
}

func (i *Item) GetName() string {
	return i.item.Name
}

func (i *Item) SetName(name string) error {
	if i.item == nil || i.item.Budget == nil {
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

func (i *Item) SetTotalBudget(amount float64) error {
	if i.item == nil || i.item.Budget == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if amount < 0 {
		return ErrInvalidAmount
	}
	i.item.Budget.Total = amount
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	return nil
}

func (i *Item) AddTransaction(transaction *valueobject.Transaction) error {
	if i.item == nil || i.item.Budget == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transaction == nil {
		return ErrInvalidTransaction
	}
	i.item.Budget.Spent += transaction.Amount
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	i.transactions = append(i.transactions, transaction)
	return nil
}

func (i *Item) GetTransactions() []*valueobject.Transaction {
	return i.transactions
}

func (i *Item) RemoveTransation(transactionToRemove *valueobject.Transaction) error {
	if i.item == nil || i.item.Budget == nil {
		//lazy initialise if item does not exist
		// i.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transactionToRemove == nil {
		return ErrInvalidTransaction
	}

	indexToRemove := -1
	for i, item := range i.transactions {
		if item == transactionToRemove {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return nil
	}
	i.item.Budget.Spent -= transactionToRemove.Amount
	i.item.Budget.Remaining = i.item.Budget.Total - i.item.Budget.Spent
	i.transactions = append(i.transactions[:indexToRemove], i.transactions[indexToRemove+1:]...)
	return nil
}
