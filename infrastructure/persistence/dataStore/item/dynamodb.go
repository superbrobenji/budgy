package datastore

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

type DynamoItemRepository struct {
	db *dynamodb.Client
}

type dynamoItem struct {
	ID           string   `dynamodbav:"id"`
	Name         string   `dynamodbav:"name"`
	Transactions []string `dynamodbav:"transactions"`
	Total        float64  `dynamodbav:"total"`
	Spent        float64  `dynamodbav:"spent"`
	CategoryID   string   `dynamodbav:"category_id"`
}

func NewDynamoItem(item *aggregate.Item) dynamoItem {
	stringTransactions := make([]string, 0)
	for _, transactionID := range item.GetTransactionIDs() {
		stringTransactions = append(stringTransactions, transactionID.String())
	}
	return dynamoItem{
		ID:           item.GetID().String(),
		Name:         item.GetName(),
		Transactions: stringTransactions,
		Total:        item.GetBudget().Total,
		Spent:        item.GetBudget().Spent,
		CategoryID:   item.GetCategoryID().String(),
	}
}

// TOOD create logger to track errors
// TODO create proper error handling
func NewAggregateItem(item *dynamoItem) (aggregate.Item, error) {
	var error error = nil
	categoryID, _ := uuid.Parse(item.CategoryID)
	itemID, _ := uuid.Parse(item.ID)
	uuidTransactions := make([]uuid.UUID, 0)
	for _, transactionID := range item.Transactions {
		uuidTransaction, _ := uuid.Parse(transactionID)
		uuidTransactions = append(uuidTransactions, uuidTransaction)
	}
	newItem, error := aggregate.NewItem(item.Name, item.Total, categoryID)
	error = newItem.SetID(itemID)
	error = newItem.SetBudgetSpent(item.Spent)
	error = newItem.SetTransactionIDs(uuidTransactions)
    if error != nil {
        return aggregate.Item{}, error
    }
	return newItem, nil
}
//TODO functions:
// - GetItemsByCategoryID
// - GetItemByID
// - PutItem
// - DeleteItem
// - UpdateItem
// - GetItemsByDate (range)
