package aggregate

import (
	"time"

	"github.com/google/uuid"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Transaction struct {
	transaction *valueobject.Transaction
}

func NewTransaction(name string, amount float64) (Transaction, error) {
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
		Date:   time.Now(),
	}
	return Transaction{
		transaction: transaction,
	}, nil
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
