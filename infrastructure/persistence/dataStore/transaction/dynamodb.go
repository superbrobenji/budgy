package datastore

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrCreatingTransactionRepository = errors.New("error creating the transaction repository")
	ErrParsingUuid                   = errors.New("error parsing uuid in Transaction repository")
	ErrNoDynamoObject                = errors.New("error parsing dynamo object in Transaction repository")
	ErrParsingTime                   = errors.New("error parsing time in Transaction repository")
)

type DynamoTransactionRepository struct {
	db *dynamodb.Client
}

type dynamoTransaction struct {
	Id     string  `dynamodbav:"Id"`
	Name   string  `dynamodbav:"name"`
	Amount float64 `dynamodbav:"amount"`
	Date   string  `dynamodbav:"date"`
	ItemID string  `dynamodbav:"item_id"`
}

func NewDynamoTransaction(transaction *aggregate.Transaction) (dynamoTransaction, error) {
	if transaction == nil {
		return dynamoTransaction{}, aggregate.ErrUnInitialised
	}
	return dynamoTransaction{
		Id:     transaction.GetID().String(),
		Name:   transaction.GetName(),
		Amount: transaction.GetAmount(),
		Date:   transaction.GetDate().String(),
		ItemID: transaction.GetItemID().String(),
	}, nil
}

func NewAggregateTransaction(transaction *dynamoTransaction) (aggregate.Transaction, error) {
    if transaction == nil {
        return aggregate.Transaction{}, ErrNoDynamoObject
    }

	timeCreated, errorParsingTime := time.Parse("2021-11-22", transaction.Date)
	if errorParsingTime != nil {
		return aggregate.Transaction{}, ErrParsingTime
	}

	transactionID, errorParsingUuid := uuid.Parse(transaction.Id)
	if errorParsingUuid != nil {
		return aggregate.Transaction{}, ErrParsingUuid
	}

	itemID, errorParsingUuid := uuid.Parse(transaction.ItemID)
	if errorParsingUuid != nil {
		return aggregate.Transaction{}, ErrParsingUuid
	}

	newTransaction, errorCreatingTransactionAggregate :=
		aggregate.NewTransaction(transaction.Name, timeCreated, transaction.Amount, itemID)
	if errorCreatingTransactionAggregate != nil {
		return aggregate.Transaction{}, errorCreatingTransactionAggregate
	}

	errorSettingID := newTransaction.SetID(transactionID)
	if errorSettingID != nil {
		//TODO create error handler
		return aggregate.Transaction{}, errorSettingID
	}

	return newTransaction, nil
}

//TODO functions:
// - GetTransactionsByItemID
// - GetTransactionsByDate (range)
// - GetTransactionByID
// - PutTransaction
// - DeleteTransaction
