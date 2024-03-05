package sdk

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamodbClient struct {
    db *dynamodb.Client
}
type KeyBasedStruct struct {
	Id          string `dynamodbav:"Id"`
}

func NewDynamodbClient() *DynamodbClient {
    return &DynamodbClient{
        db: GetDynamodbClient(),
    }
}

func (ddbClient *DynamodbClient) DynamodbPutWrapper(item interface{}, conditionExp *string, table string) (*dynamodb.PutItemOutput, error) {
    fmt.Println(item)
	av, marshalErr := attributevalue.MarshalMap(item)
    fmt.Println(av)
	if marshalErr != nil {
		return &dynamodb.PutItemOutput{}, marshalErr
	}

	putItemRes, putItemErr := ddbClient.db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: conditionExp,
	})
	if putItemErr != nil {
		return &dynamodb.PutItemOutput{}, putItemErr
	}

	return putItemRes, nil
}

func (ddbClient *DynamodbClient) DynamodbGetWrapper(key interface{}, resultItem interface{}, table string) (*dynamodb.GetItemOutput, error) {
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.GetItemOutput{}, marshalErr
	}

	getItemRes, getItemErr := ddbClient.db.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key:       av,
	})
	if getItemErr != nil {
		return &dynamodb.GetItemOutput{}, getItemErr
	}
	unmarshalErr := attributevalue.UnmarshalMap(getItemRes.Item, resultItem)
	if unmarshalErr != nil {
		return &dynamodb.GetItemOutput{}, unmarshalErr
	}

	return getItemRes, nil
}
func (ddbClient *DynamodbClient) DynamodbQueryWrapper(query expression.KeyConditionBuilder, resultItem interface{}, table string) (*dynamodb.QueryOutput, error) {
    expr, builderErr := expression.NewBuilder().WithKeyCondition(query).Build()
    if builderErr != nil {
        return &dynamodb.QueryOutput{}, builderErr
    }

	getItemRes, getItemErr := ddbClient.db.Query(context.TODO(), &dynamodb.QueryInput{
		TableName: aws.String(table),
		KeyConditionExpression: expr.KeyCondition(),
	})

	if getItemErr != nil {
		return &dynamodb.QueryOutput{}, getItemErr
	}
	unmarshalErr := attributevalue.UnmarshalListOfMaps(getItemRes.Items, resultItem)
	if unmarshalErr != nil {
		return &dynamodb.QueryOutput{}, unmarshalErr
	}

	return getItemRes, nil
}

func (ddbClient *DynamodbClient) DynamodbUpdateWrapper(key interface{}, update expression.UpdateBuilder, table string) (*dynamodb.UpdateItemOutput, error) {
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.UpdateItemOutput{}, marshalErr
	}

	expr, builderErr := expression.NewBuilder().WithUpdate(update).Build()
	if builderErr != nil {
		return &dynamodb.UpdateItemOutput{}, builderErr
	}

	updateItemRes, updateItemErr := ddbClient.db.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName:                 aws.String(table),
		Key:                       av,
		UpdateExpression:          expr.Update(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
	})
	if updateItemErr != nil {
		return &dynamodb.UpdateItemOutput{}, updateItemErr
	}

	return updateItemRes, nil
}

func (ddbClient *DynamodbClient) DynamodbDeleteWrapper(key interface{}, table string) (*dynamodb.DeleteItemOutput, error) {
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.DeleteItemOutput{}, marshalErr
	}

	deleteItemRes, deleteItemErr := ddbClient.db.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       av,
	})
	if deleteItemErr != nil {
		return &dynamodb.DeleteItemOutput{}, deleteItemErr
	}

	return deleteItemRes, nil
}
