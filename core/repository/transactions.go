package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrFailedToAddTransaction = errors.New("failed to add transaction")
	ErrDeleteTransaction      = errors.New("failed to delete transaction")
)

type TransactionRepository interface {
	Get(uuid.UUID) (aggregate.Category, error)
	Add(aggregate.Category) error
	Delete(uuid.UUID) error
}
