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
	ErrCreatingItemRepository = errors.New("error creating the item repository")
	ErrParsingUuid            = errors.New("error parsing uuid in Item repository")
	ErrNoDynamoObject         = errors.New("error parsing dynamo object in Item repository")
)

type DynamoItemRepository struct {
	db *sdk.DynamodbClient
}

type dynamoItem struct {
	Id           string   `dynamodbav:"Id"`
	Name         string   `dynamodbav:"name"`
	Transactions []string `dynamodbav:"transactions"`
	Total        float64  `dynamodbav:"total"`
	Spent        float64  `dynamodbav:"spent"`
	CategoryID   string   `dynamodbav:"category_id"`
}
type CategoryKeyBasedStruct struct {
	CategoryID string `dynamodbav:"category_id"`
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

func NewDynamoItemRepository() *DynamoItemRepository {
	return &DynamoItemRepository{
		db: sdk.NewDynamodbClient(),
	}
}

func (dir *DynamoItemRepository) CreateItem(item *aggregate.Item) error {
	dynamoItem, errorCreatingDynamoItem := NewDynamoItem(item)
	if errorCreatingDynamoItem != nil {
		return errorCreatingDynamoItem
	}

	_, errorPuttingItem := dir.db.DynamodbPutWrapper(dynamoItem, nil, "items")
	if errorPuttingItem != nil {
		return errorPuttingItem
	}

	return nil
}

func (dir *DynamoItemRepository) GetItemByID(id uuid.UUID) (*aggregate.Item, error) {
	key := &sdk.KeyBasedStruct{
		Id: id.String(),
	}
	result := &dynamoItem{}
	_, err := dir.db.DynamodbGetWrapper(key, result, "items")
	if err != nil {
		return &aggregate.Item{}, err
	}
	aggregateItem, err := NewAggregateItem(result)
	if err != nil {
		return &aggregate.Item{}, err
	}

	return &aggregateItem, nil
}
func (dir *DynamoItemRepository) UpdateItem(item *aggregate.Item) error {
	dynamoItem, err := NewDynamoItem(item)
	if err != nil {
		return err
	}
	var update expression.UpdateBuilder
	values := reflect.ValueOf(dynamoItem)
	types := reflect.TypeOf(dynamoItem)
	for i := 0; i < values.NumField(); i++ {
		update = update.Set(expression.Name(types.Field(i).Name), expression.Value(values.Field(i)))
	}
	_, err = dir.db.DynamodbUpdateWrapper(&sdk.KeyBasedStruct{Id: item.GetID().String()}, update, "items")
	if err != nil {
		return err
	}
	return nil
}
func (dir *DynamoItemRepository) DeleteItem(id uuid.UUID) error {
	key := &sdk.KeyBasedStruct{
		Id: id.String(),
	}
	_, err := dir.db.DynamodbDeleteWrapper(key, "items")
	if err != nil {
		return err
	}
	return nil
}
func (dir *DynamoItemRepository) GetItemsByCategoryID(categoryID uuid.UUID) (*[]*aggregate.Item, error) {
	var query expression.KeyConditionBuilder
	query = expression.Key("category_id").Equal(expression.Value(categoryID.String()))
	dbItems := []dynamoItem{}
	_, err := dir.db.DynamodbQueryWrapper(query, &dbItems, "items")
	if err != nil {
		return &[]*aggregate.Item{}, err
	}
	items := make([]*aggregate.Item, len(dbItems), 0)
	for index, dbItem := range dbItems {
		item, err := NewAggregateItem(&dbItem)
		if err != nil {
			return &[]*aggregate.Item{}, err
		}
		items[index] = &item
	}
	return &items, nil
}
