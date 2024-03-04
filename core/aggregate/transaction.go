package aggregate

import (
	"time"

	"github.com/google/uuid"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Transaction struct {
	transaction *valueobject.Transaction
	itemID      uuid.UUID
}

func NewTransaction(name string, date time.Time, amount float64, itemID uuid.UUID) (Transaction, error) {
	if name == "" {
		return Transaction{}, ErrInvalidName
	}
	if amount < 0 {
		return Transaction{}, ErrInvalidAmount
	}

	transaction := &valueobject.Transaction{
		Name:   name,
		Amount: amount,
		ID:     uuid.New(),
		Date:   date,
	}
	return Transaction{
		transaction: transaction,
		itemID:      itemID,
	}, nil
}
func (t *Transaction) SetID(id uuid.UUID) error {
	if t.transaction == nil {
		//lazy initialise if transaction does not exist
		// t.transaction = &valueobject.Transaction{}
		return ErrUnInitialised
	}
	t.transaction.ID = id
	return nil
}
func (t *Transaction) GetParentItemID() uuid.UUID {
	return t.itemID
}

func (t *Transaction) GetID() uuid.UUID {
	return t.transaction.ID
}

func (t *Transaction) GetName() string {
	return t.transaction.Name
}
func (t *Transaction) GetAmount() float64 {
	return t.transaction.Amount
}
func (t *Transaction) GetDate() time.Time {
	return t.transaction.Date
}
func (t *Transaction) GetItemID() uuid.UUID {
	return t.itemID
}
