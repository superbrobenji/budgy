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
	ID     string  `dynamodbav:"id"`
	Name   string  `dynamodbav:"name"`
	Amount float64 `dynamodbav:"amount"`
	Date   string  `dynamodbav:"date"`
    ItemID string  `dynamodbav:"item_id"`
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

func NewAggregateTransaction(transaction *dynamoTransaction) (aggregate.Transaction, error) {
	var error error = nil
	timeCreated, _ := time.Parse("2021-11-22", transaction.Date)
	transactionID, _ := uuid.Parse(transaction.ID)
	itemID, _ := uuid.Parse(transaction.ItemID)
	newTransaction, error :=
		aggregate.NewTransaction(transaction.Name, timeCreated, transaction.Amount, itemID)
	error = newTransaction.SetID(transactionID)
	if error != nil {
		//TODO create error handler
		return aggregate.Transaction{}, error
	}
	return newTransaction, nil
}

//TODO functions:
// - GetTransactionsByItemID
// - GetTransactionsByDate (range)
// - GetTransactionByID
// - PutTransaction
// - DeleteTransaction

