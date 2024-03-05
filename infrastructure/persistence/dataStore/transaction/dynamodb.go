package datastore

import (
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	sdk "github.com/superbrobenji/budgy/infrastructure/aws/sdk/dynamodb"
)

var (
	ErrCreatingTransactionRepository = errors.New("error creating the transaction repository")
	ErrParsingUuid                   = errors.New("error parsing uuid in Transaction repository")
	ErrNoDynamoObject                = errors.New("error parsing dynamo object in Transaction repository")
	ErrParsingTime                   = errors.New("error parsing time in Transaction repository")
)

type DynamoTransactionRepository struct {
	db *sdk.DynamodbClient
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
		Date:   strconv.FormatInt(transaction.GetDate().Unix(), 10),
		ItemID: transaction.GetItemID().String(),
	}, nil
}

func NewAggregateTransaction(transaction *dynamoTransaction) (aggregate.Transaction, error) {
	if transaction == nil {
		return aggregate.Transaction{}, ErrNoDynamoObject
	}
    intTime, err := strconv.ParseInt(transaction.Date, 10, 64)
    if err != nil {
        return aggregate.Transaction{}, err
    }
    timeCreated := time.Unix(intTime, 0)

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

func NewDynamoTransactionRepository() *DynamoTransactionRepository {
	return &DynamoTransactionRepository{
		db: sdk.NewDynamodbClient(),
	}
}

func (dtr *DynamoTransactionRepository) CreateTransaction(transaction *aggregate.Transaction) error {
	dynamoTransaction, errorCreatingDynamoTransaction := NewDynamoTransaction(transaction)
	if errorCreatingDynamoTransaction != nil {
		return errorCreatingDynamoTransaction
	}

	_, errorPuttingTransaction := dtr.db.DynamodbPutWrapper(dynamoTransaction, nil, "transactions")
	if errorPuttingTransaction != nil {
		return errorPuttingTransaction
	}

	return nil
}

func (dtr *DynamoTransactionRepository) GetTransactionByID(id uuid.UUID) (*aggregate.Transaction, error) {
	key := &sdk.KeyBasedStruct{
		Id: id.String(),
	}
	result := &dynamoTransaction{}
	_, errorGettingTransaction := dtr.db.DynamodbGetWrapper(key, result, "transactions")
	if errorGettingTransaction != nil {
		return &aggregate.Transaction{}, errorGettingTransaction
	}

	transaction, errorCreatingTransactionAggregate := NewAggregateTransaction(result)
	if errorCreatingTransactionAggregate != nil {
		return &aggregate.Transaction{}, errorCreatingTransactionAggregate
	}

	return &transaction, nil
}

func (dtr *DynamoTransactionRepository) GetTransactionsByItemID(id uuid.UUID) (*[]*aggregate.Transaction, error) {
	var query expression.KeyConditionBuilder
	query = expression.Key("item_id").Equal(expression.Value(id.String()))
	dbItems := []dynamoTransaction{}
	_, err := dtr.db.DynamodbQueryWrapper(query, &dbItems, "items")
	if err != nil {
		return &[]*aggregate.Transaction{}, err
	}
	transactions := make([]*aggregate.Transaction, len(dbItems), 0)
	for index, dbItem := range dbItems {
		transaction, err := NewAggregateTransaction(&dbItem)
		if err != nil {
			return &[]*aggregate.Transaction{}, err
		}
		transactions[index] = &transaction
	}
	return &transactions, nil
}

// TODO update to only update the fields that have changed
func (dtr *DynamoTransactionRepository) UpdateTransaction(transaction *aggregate.Transaction) error {
    dynamoTransaction, errorCreatingDynamoTransaction := NewDynamoTransaction(transaction)
    if errorCreatingDynamoTransaction != nil {
        return errorCreatingDynamoTransaction
    }

    var update expression.UpdateBuilder
    values := reflect.ValueOf(dynamoTransaction)
    types := reflect.TypeOf(dynamoTransaction)
    for i := 0; i < values.NumField(); i++ {
        update = update.Set(expression.Name(types.Field(i).Name), expression.Value(values.Field(i)))
    }
    _, err := dtr.db.DynamodbUpdateWrapper(&sdk.KeyBasedStruct{Id: transaction.GetID().String()}, update, "transactions")
    if err != nil {
        return err
    }
    return nil
}

func (dtr *DynamoTransactionRepository) DeleteTransaction(id uuid.UUID) error {
    key := &sdk.KeyBasedStruct{
        Id: id.String(),
    }
    _, err := dtr.db.DynamodbDeleteWrapper(key, "transactions")
    if err != nil {
        return err
    }
    return nil
}

func (dtr *DynamoTransactionRepository) GetTransactionsByDate(startDate time.Time, endDate time.Time) (*[]*aggregate.Transaction, error) {
    unixStart := startDate.Unix()
    unixEnd := endDate.Unix()
    startValExp := expression.Value(unixStart)
    endValExp := expression.Value(unixEnd)
	var query expression.KeyConditionBuilder
	query = expression.Key("date").Between(startValExp, endValExp)
	dbItems := []dynamoTransaction{}
	_, err := dtr.db.DynamodbQueryWrapper(query, &dbItems, "items")
	if err != nil {
		return &[]*aggregate.Transaction{}, err
	}
	transactions := make([]*aggregate.Transaction, len(dbItems), 0)
	for index, dbItem := range dbItems {
		transaction, err := NewAggregateTransaction(&dbItem)
		if err != nil {
			return &[]*aggregate.Transaction{}, err
		}
		transactions[index] = &transaction
	}
	return &transactions, nil
}
