package datastore

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrCreatingItemRepository = errors.New("error creating the item repository")
	ErrParsingUuid            = errors.New("error parsing uuid in Item repository")
	ErrNoDynamoObject         = errors.New("error parsing dynamo object in Item repository")
)

type DynamoItemRepository struct {
	db *dynamodb.Client
}

type dynamoItem struct {
	Id           string   `dynamodbav:"Id"`
	Name         string   `dynamodbav:"name"`
	Transactions []string `dynamodbav:"transactions"`
	Total        float64  `dynamodbav:"total"`
	Spent        float64  `dynamodbav:"spent"`
	CategoryID   string   `dynamodbav:"category_id"`
}

func NewDynamoItem(item *aggregate.Item) (dynamoItem, error) {
	if item == nil {
		return dynamoItem{}, aggregate.ErrUnInitialised
	}
	stringTransactions := make([]string, 0)
	for _, transactionID := range item.GetTransactionIDs() {
		stringTransactions = append(stringTransactions, transactionID.String())
	}
	return dynamoItem{
		Id:           item.GetID().String(),
		Name:         item.GetName(),
		Transactions: stringTransactions,
		Total:        item.GetBudget().Total,
		Spent:        item.GetBudget().Spent,
		CategoryID:   item.GetCategoryID().String(),
	}, nil
}

// TOOD create logger to track errors
// TODO create proper error handling
func NewAggregateItem(item *dynamoItem) (aggregate.Item, error) {
	if item == nil {
		return aggregate.Item{}, ErrNoDynamoObject
	}

	categoryID, errParsingUuid := uuid.Parse(item.CategoryID)
	if errParsingUuid != nil {
		return aggregate.Item{}, ErrParsingUuid
	}

	itemID, errorParsingUuid := uuid.Parse(item.Id)
	if errorParsingUuid != nil {
		return aggregate.Item{}, ErrParsingUuid
	}

	uuidTransactions := make([]uuid.UUID, 0)
	for _, transactionID := range item.Transactions {
		uuidTransaction, errorParsingUuid := uuid.Parse(transactionID)
		if errorParsingUuid != nil {
			return aggregate.Item{}, ErrParsingUuid
		}
		uuidTransactions = append(uuidTransactions, uuidTransaction)
	}

	newItem, errorCreatingItem := aggregate.NewItem(item.Name, item.Total, categoryID)
	if errorCreatingItem != nil {
		return aggregate.Item{}, errorCreatingItem
	}

	errorSettingID := newItem.SetID(itemID)
	if errorSettingID != nil {
		return aggregate.Item{}, errorSettingID
	}

	errorSettingBudgetSpent := newItem.SetBudgetSpent(item.Spent)
	if errorSettingBudgetSpent != nil {
		return aggregate.Item{}, errorSettingBudgetSpent
	}

	errorSettingTransactionIDs := newItem.SetTransactionIDs(uuidTransactions)
	if errorSettingTransactionIDs != nil {
		return aggregate.Item{}, errorSettingTransactionIDs
	}

	return newItem, nil
}

//TODO functions:
// - GetItemsByCategoryID
// - GetItemByID
// - PutItem
// - DeleteItem
// - UpdateItem
