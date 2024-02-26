package aggregate

import (
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/model/entity"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

type Item struct {
    //TODO possibly make the transations aggregates
	item         *entity.Item
	transactions []*valueobject.Transaction
}

// TODO add all the budget functions
func NewItem(name string) (Item, error) {
	if name == "" {
		return Item{}, ErrInvalidName
	}

	item := &entity.Item{
		Name: name,
		ID:   uuid.New(),
	}
	return Item{
		item:         item,
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}

func (c *Item) GetID() uuid.UUID {
	return c.item.ID
}

func (c *Item) GetName() string {
	return c.item.Name
}

func (c *Item) SetName(name string) error {
	if c.item == nil {
		//lazy initialise if item does not exist
		// c.item = &entity.Item{}
		return ErrUnInitialised
	}
	if name == "" {
		return ErrInvalidName
	}
	c.item.Name = name
	return nil
}

func (c *Item) AddTransaction(transaction *valueobject.Transaction) error {
	//TODO update the budget based on amount on transaction
	if c.item == nil {
		//lazy initialise if item does not exist
		// c.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transaction == nil {
		return ErrInvalidTransaction
	}
	c.transactions = append(c.transactions, transaction)
	return nil
}

func (c *Item) GetTransactions() []*valueobject.Transaction {
	return c.transactions
}

func (c *Item) RemoveTransation(transactionToRemove *valueobject.Transaction) error {
	//TODO update the budget based on amount on transaction
	if c.item == nil {
		//lazy initialise if item does not exist
		// c.item = &entity.Item{}
		return ErrUnInitialised
	}
	if transactionToRemove == nil {
		return ErrInvalidTransaction
	}

	indexToRemove := -1
	for i, item := range c.transactions {
		if item == transactionToRemove {
			indexToRemove = i
			break
		}
	}
	if indexToRemove == -1 {
		return nil
	}
	c.transactions = append(c.transactions[:indexToRemove], c.transactions[indexToRemove+1:]...)
	return nil
}
