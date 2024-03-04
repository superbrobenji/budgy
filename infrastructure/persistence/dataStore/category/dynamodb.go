package datastore

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

type DynamoCategoryRepository struct {
	db *dynamodb.Client
}

type dynamoCategory struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Items []string `json:"items"`
	Total float64  `json:"total"`
	Spent float64  `json:"spent"`
}

func NewDynamoCategory(category *aggregate.Category) dynamoCategory {
	stringItemIDs := make([]string, 0)
	for _, itemID := range category.GetItemIDs() {
		stringItemIDs = append(stringItemIDs, itemID.String())
	}
	return dynamoCategory{
		ID:    category.GetID().String(),
		Name:  category.GetName(),
		Items: stringItemIDs,
		Total: category.GetBudget().Total,
		Spent: category.GetBudget().Spent,
	}
}

// TOOD create logger to track errors
// TODO create proper error handling
func NewAggregateCategory(dynamoCategory *dynamoCategory) (aggregate.Category, error) {
	newCategory, error := aggregate.NewCategory(dynamoCategory.Name)
	if error != nil {
		//TODO create error handler
		return aggregate.Category{}, error
	}
	categoryID, _ := uuid.Parse(dynamoCategory.ID)
	itemIDs := make([]uuid.UUID, 0)
	for _, itemID := range dynamoCategory.Items {
		uuidItemID, _ := uuid.Parse(itemID)
		itemIDs = append(itemIDs, uuidItemID)
	}
	//TODO handle errors for all of these
	newCategory.SetID(categoryID)
	newCategory.SetBudgetTotal(dynamoCategory.Total)
	newCategory.SetBudgetSpent(dynamoCategory.Spent)
	newCategory.SetItemIDs(itemIDs)
	return newCategory, nil
}
