package datastore

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type DynamoTransactionRepository struct {
	db *dynamodb.Client
}

type dynamoTransaction struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Amount float64 `json:"amount"`
	Date   string  `json:"date"`
	ItemID string  `json:"item_id"`
}
