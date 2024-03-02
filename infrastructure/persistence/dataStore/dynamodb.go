//remove this file
package datastore

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBRepository struct {
    db *dynamodb.Client
}
type dynamoCategory struct {
}
type dynamoItem struct {
}
type dynamoTransaction struct {
}

//func NewDynamoDB() {
//    // Load the Shared AWS Configuration (~/.aws/config)
//    cfg, err := config.LoadDefaultConfig(context.TODO())
//    if err != nil {
//        log.Fatal(err)
//    }
//    // Create an Amazon DynamoDB service client
//    client := dynamodb.NewFromConfig(cfg)
//
//}
