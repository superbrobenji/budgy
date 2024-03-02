package datastore

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

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
