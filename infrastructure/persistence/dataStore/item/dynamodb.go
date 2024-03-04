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
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Transactions []string `json:"transactions"`
	Total        float64  `json:"total"`
	Spent        float64  `json:"spent"`
	CategoryID   string   `json:"category_id"`
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
	categoryID, _ := uuid.Parse(item.CategoryID)
	itemID, _ := uuid.Parse(item.ID)
	uuidTransactions := make([]uuid.UUID, 0)
	for _, transactionID := range item.Transactions {
		uuidTransaction, _ := uuid.Parse(transactionID)
		uuidTransactions = append(uuidTransactions, uuidTransaction)
	}
	newItem, _ := aggregate.NewItem(item.Name, item.Total, categoryID)
	newItem.SetID(itemID)
	newItem.SetBudgetSpent(item.Spent)
	newItem.SetTransactionIDs(uuidTransactions)
	return newItem, nil
}
