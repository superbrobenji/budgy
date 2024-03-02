package datastore

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

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
