package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrTransactionNotFound    = errors.New("transaction not found")
	ErrFailedToAddTransaction = errors.New("failed to add transaction")
	ErrDeleteTransaction      = errors.New("failed to delete transaction")
)

type TransactionRepository interface {
	GetTransactionsByItemID(uuid.UUID) (aggregate.Category, error)
	GetTransactionsByDate(time.Time, time.Time) error
	GetTransactionByID(uuid.UUID) error
	PutTransaction(aggregate.Transaction) error
	DeleteTransaction(uuid.UUID) error
}
