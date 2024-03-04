package datastore

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
)

type DynamoCategoryRepository struct {
	db *dynamodb.Client
}

type dynamoCategory struct {
	ID    string   `dynamodbav:"id"`
	Name  string   `dynamodbav:"name"`
	Items []string `dynamodbav:"items"`
	Total float64  `dynamodbav:"total"`
	Spent float64  `dynamodbav:"spent"`
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
	var error error = nil
	newCategory, error := aggregate.NewCategory(dynamoCategory.Name)
	categoryID, _ := uuid.Parse(dynamoCategory.ID)
	itemIDs := make([]uuid.UUID, 0)
	for _, itemID := range dynamoCategory.Items {
		uuidItemID, _ := uuid.Parse(itemID)
		itemIDs = append(itemIDs, uuidItemID)
	}
	//TODO handle errors for all of these
	error = newCategory.SetID(categoryID)
	error = newCategory.SetBudgetTotal(dynamoCategory.Total)
	error = newCategory.SetBudgetSpent(dynamoCategory.Spent)
	error = newCategory.SetItemIDs(itemIDs)
	if error != nil {
		return aggregate.Category{}, error
	}
	return newCategory, nil
}

func New() (*DynamoCategoryRepository, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoCategoryRepository{db: client}, nil
}
//TODO functions:
// - GetCategoryByID
// - GetCategoryByItemID
// - PutCategory
// - DeleteCategory
// - UpdateCategory
