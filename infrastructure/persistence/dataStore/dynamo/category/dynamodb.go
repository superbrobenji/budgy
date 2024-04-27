package datastore

import (
	"errors"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	sdk "github.com/superbrobenji/budgy/infrastructure/aws/sdk/dynamodb"
)

var (
	ErrCreatingCategoryRepository = errors.New("error creating the categorty repository")
	ErrParsingUuid                = errors.New("error parsing uuid in Category repository")
	ErrNoDynamoObject             = errors.New("error parsing dynamo object in Category repository")
)

type DynamoCategoryRepository struct {
	db *sdk.DynamodbClient
}

type dynamoCategory struct {
	Id    string   `dynamodbav:"Id"`
	Name  string   `dynamodbav:"name"`
	Items []string `dynamodbav:"items"`
	Total float64  `dynamodbav:"total"`
	Spent float64  `dynamodbav:"spent"`
}

func NewDynamoCategory(category *aggregate.Category) (dynamoCategory, error) {
	if category == nil {
		return dynamoCategory{}, aggregate.ErrUnInitialised
	}
	stringItemIDs := make([]string, 0)
	for _, itemID := range category.GetItemIDs() {
		stringItemIDs = append(stringItemIDs, itemID.String())
	}
	return dynamoCategory{
		Id:    category.GetID().String(),
		Name:  category.GetName(),
		Items: stringItemIDs,
		Total: category.GetBudget().Total,
		Spent: category.GetBudget().Spent,
	}, nil
}

// TOOD create logger to track errors
// TODO create proper error handling
func NewAggregateCategory(dynamoCategory *dynamoCategory) (*aggregate.Category, error) {
	if dynamoCategory == nil {
		return &aggregate.Category{}, ErrNoDynamoObject
	}

	newCategory, errorCreatingCategory := aggregate.NewCategory(dynamoCategory.Name)
	if errorCreatingCategory != nil {
		return &aggregate.Category{}, errorCreatingCategory
	}

	categoryID, errorParsingUuid := uuid.Parse(dynamoCategory.Id)
	if errorParsingUuid != nil {
		return &aggregate.Category{}, ErrParsingUuid
	}

	itemIDs := make([]uuid.UUID, 0)
	for _, itemID := range dynamoCategory.Items {
		uuidItemID, errorParsingUuid := uuid.Parse(itemID)
		if errorParsingUuid != nil {
			return &aggregate.Category{}, ErrParsingUuid
		}
		itemIDs = append(itemIDs, uuidItemID)
	}

	//TODO handle errors for all of these
	errorSettingId := newCategory.SetID(categoryID)
	if errorSettingId != nil {
		return &aggregate.Category{}, errorSettingId
	}

	errorSettingBudgetTotal := newCategory.SetBudgetTotal(dynamoCategory.Total)
	if errorSettingBudgetTotal != nil {
		return &aggregate.Category{}, errorSettingBudgetTotal
	}

	errorSettingBudgetSpent := newCategory.SetBudgetSpent(dynamoCategory.Spent)
	if errorSettingBudgetSpent != nil {
		return &aggregate.Category{}, errorSettingBudgetSpent
	}

	errorSettingItemIDs := newCategory.SetItemIDs(itemIDs)
	if errorSettingItemIDs != nil {
		return &aggregate.Category{}, errorSettingItemIDs
	}

	return &newCategory, nil
}

func NewDynamoCategoryRepository() *DynamoCategoryRepository {
	return &DynamoCategoryRepository{
		db: sdk.NewDynamodbClient(),
	}
}

func (d *DynamoCategoryRepository) GetCategoryByID(id uuid.UUID) (*aggregate.Category, error) {
	key := &sdk.KeyBasedStruct{
		Id: id.String(),
	}
	result := &dynamoCategory{}
	_, err := d.db.DynamodbGetWrapper(key, result, "categories")
	if err != nil {
		return &aggregate.Category{}, err
	}
	return NewAggregateCategory(result)
}

func (d *DynamoCategoryRepository) CreateCategory(category *aggregate.Category) error {
	dbCategory, err := NewDynamoCategory(category)
	if err != nil {
		return err
	}
	_, err = d.db.DynamodbPutWrapper(dbCategory, nil, "categories")
	if err != nil {
		return err
	}
	return nil
}
func (d *DynamoCategoryRepository) DeleteCategory(id uuid.UUID) error {
	key := &sdk.KeyBasedStruct{
		Id: id.String(),
	}
	_, err := d.db.DynamodbDeleteWrapper(key, "categories")
	if err != nil {
		return err
	}
	return nil
}

// TODO update only the fields that have changed
func (d *DynamoCategoryRepository) UpdateCategory(category *aggregate.Category) error {
	dbCategory, err := NewDynamoCategory(category)
	if err != nil {
		return err
	}
	var update expression.UpdateBuilder
	values := reflect.ValueOf(dbCategory)
	types := reflect.TypeOf(dbCategory)
	for i := 0; i < values.NumField(); i++ {
		update = update.Set(expression.Name(types.Field(i).Name), expression.Value(values.Field(i)))
	}
	_, err = d.db.DynamodbUpdateWrapper(&sdk.KeyBasedStruct{Id: category.GetID().String()}, update, "categories")
	if err != nil {
		return err
	}
	return nil
}
