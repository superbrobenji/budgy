package datastore

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

type DynamoTransactionRepository struct {
	db *dynamodb.Client
}

type dynamoTransaction struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	ItemID string  `json:"item_id"`
}

func NewDynamoTransaction(transaction *aggregate.Transaction) dynamoTransaction {
	return dynamoTransaction{
		ID:     transaction.GetID().String(),
		Name:   transaction.GetName(),
		Amount: transaction.GetAmount(),
		Date:   transaction.GetDate().String(),
		ItemID: transaction.GetItemID().String(),
	}
}

// TOOD create logger to track errors
// TODO create proper error handling
func NewAggregateTransaction(transaction *dynamoTransaction) (aggregate.Transaction, error) {
	timeCreated, _ := time.Parse("2021-11-22", transaction.Date)
    transactionID, _ := uuid.Parse(transaction.ID)
    itemID, _ := uuid.Parse(transaction.ItemID)
	newTransaction, error :=
		aggregate.NewTransaction(transaction.Name, timeCreated, transaction.Amount, itemID)
	if error != nil {
		//TODO create error handler
		return aggregate.Transaction{}, error
	}
	//TODO handle errors for all of these
	newTransaction.SetID(transactionID)
	return newTransaction, nil
}
