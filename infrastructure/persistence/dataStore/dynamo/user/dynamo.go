package datastore

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/superbrobenji/budgy/core/aggregate"
	sdk "github.com/superbrobenji/budgy/infrastructure/aws/sdk/dynamodb"
)

var (
	ErrCreatingUserRepository = errors.New("error creating the User repository")
	ErrParsingUuid            = errors.New("error parsing uuid in User repository")
	ErrNoDynamoObject         = errors.New("error parsing dynamo object in User repository")
	ErrParsingTime            = errors.New("error parsing time in User repository")
)

type DynamoUserRepository struct {
	db *sdk.DynamodbClient
}

type dynamoUser struct {
	UserID     string `dynamodbav:"Id"`
	Username   string `dynamodbav:"username"`
	DateJoined string `dynamodbav:"date"`
	Email      string `dynamodbav:"email"`
}

func NewDynamoUser(user *aggregate.User) (dynamoUser, error) {
	if user == nil {
		return dynamoUser{}, aggregate.ErrUnInitialised
	}
	return dynamoUser{
		UserID:     user.GetID(),
		Username:   user.GetUsername(),
		DateJoined: user.GetDateJoined().UTC().Format(time.RFC1123),
		Email:      user.GetEmail(),
	}, nil
}

func NewAggregateUser(user *dynamoUser) (aggregate.User, error) {
	if user == nil {
		return aggregate.User{}, ErrNoDynamoObject
	}
	timeCreated, err := time.Parse(time.RFC1123, user.DateJoined)
	if err != nil {
		return aggregate.User{}, err
	}

	newUser, errorCreatingUserAggregate :=
		aggregate.NewUser(user.UserID, user.Username, timeCreated, user.Email)
	if errorCreatingUserAggregate != nil {
		return aggregate.User{}, errorCreatingUserAggregate
	}

	return newUser, nil
}

func NewDynamoUserRepository() *DynamoUserRepository {
	return &DynamoUserRepository{
		db: sdk.NewDynamodbClient(),
	}
}

func (dtr *DynamoUserRepository) CreateUser(user *aggregate.User) error {
	dynamoUser, errorCreatingDynamoUser := NewDynamoUser(user)
	if errorCreatingDynamoUser != nil {
		return errorCreatingDynamoUser
	}

	_, errorPuttingUser := dtr.db.DynamodbPutWrapper(dynamoUser, nil, "users")
	if errorPuttingUser != nil {
		return errorPuttingUser
	}

	return nil
}

func (dtr *DynamoUserRepository) GetUserByID(id string) (*aggregate.User, error) {
	key := &sdk.KeyBasedStruct{
		Id: id,
	}
	result := &dynamoUser{}
	_, errorGettingUser := dtr.db.DynamodbGetWrapper(key, result, "users")
	if errorGettingUser != nil {
		return &aggregate.User{}, errorGettingUser
	}

	if result.UserID == "" {
		return &aggregate.User{}, nil
	}

	user, errorCreatingUserAggregate := NewAggregateUser(result)
	if errorCreatingUserAggregate != nil {
		return &aggregate.User{}, errorCreatingUserAggregate
	}

	return &user, nil
}

func (dtr *DynamoUserRepository) GetUserByUsername(username string) (*aggregate.User, error) {
	type usernameStruct struct {
		username string `dynamodbav:"username"`
	}
	key := &usernameStruct{
		username,
	}
	result := &dynamoUser{}
	_, errorGettingUser := dtr.db.DynamodbGetWrapper(key, result, "users")
	if errorGettingUser != nil {
		return &aggregate.User{}, errorGettingUser
	}

	user, errorCreatingUserAggregate := NewAggregateUser(result)
	if errorCreatingUserAggregate != nil {
		return &aggregate.User{}, errorCreatingUserAggregate
	}

	return &user, nil
}

func (dtr *DynamoUserRepository) GetUserByEmail(email string) (*aggregate.User, error) {
	type usernameStruct struct {
		email string `dynamodbav:"email"`
	}
	key := &usernameStruct{
		email,
	}
	result := &dynamoUser{}
	_, errorGettingUser := dtr.db.DynamodbGetWrapper(key, result, "users")
	if errorGettingUser != nil {
		return &aggregate.User{}, errorGettingUser
	}

	user, errorCreatingUserAggregate := NewAggregateUser(result)
	if errorCreatingUserAggregate != nil {
		return &aggregate.User{}, errorCreatingUserAggregate
	}

	return &user, nil
}

func (dtr *DynamoUserRepository) DeleteUser(id string) error {
	key := &sdk.KeyBasedStruct{
		Id: id,
	}
	_, err := dtr.db.DynamodbDeleteWrapper(key, "users")
	if err != nil {
		return err
	}
	return nil
}

func (dtr *DynamoUserRepository) GetUsersByDateJoined(startDate time.Time, endDate time.Time) (*[]*aggregate.User, error) {
	unixStart := startDate.Unix()
	unixEnd := endDate.Unix()
	startValExp := expression.Value(unixStart)
	endValExp := expression.Value(unixEnd)
	var query expression.KeyConditionBuilder
	query = expression.Key("date").Between(startValExp, endValExp)
	dbUsers := []dynamoUser{}
	_, err := dtr.db.DynamodbQueryWrapper(query, &dbUsers, "items")
	if err != nil {
		return &[]*aggregate.User{}, err
	}
	users := make([]*aggregate.User, len(dbUsers), 0)
	for index, dbUser := range dbUsers {
		user, err := NewAggregateUser(&dbUser)
		if err != nil {
			return &[]*aggregate.User{}, err
		}
		users[index] = &user
	}
	return &users, nil
}
