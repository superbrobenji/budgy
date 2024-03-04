package sdk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type KeyBasedStruct struct {
	Id string `dynamodbav:"id"`
}

func DynamodbPutWrapper(item interface{}, conditionExp *string, table string) (*dynamodb.PutItemOutput, error) {
	ddbClient := GetDynamodbClient()
	av, marshalErr := attributevalue.MarshalMap(item)
	if marshalErr != nil {
		return &dynamodb.PutItemOutput{}, marshalErr
	}

	putItemRes, putItemErr := ddbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName:           aws.String(table),
		Item:                av,
		ConditionExpression: conditionExp,
	})
	if putItemErr != nil {
		return &dynamodb.PutItemOutput{}, putItemErr
	}

	return putItemRes, nil
}

func DynamodbGetWrapper(key interface{}, resultItem interface{}, table string) (*dynamodb.GetItemOutput, error) {
	ddbClient := GetDynamodbClient()
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.GetItemOutput{}, marshalErr
	}

	getItemRes, getItemErr := ddbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
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

func DynamodbUpdateWrapper(key interface{}, update expression.UpdateBuilder, table string) (*dynamodb.UpdateItemOutput, error) {
	ddbClient := GetDynamodbClient()
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.UpdateItemOutput{}, marshalErr
	}

	expr, builderErr := expression.NewBuilder().WithUpdate(update).Build()
	if builderErr != nil {
		return &dynamodb.UpdateItemOutput{}, builderErr
	}

	updateItemRes, updateItemErr := ddbClient.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
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

func DynamodbDeleteWrapper(key interface{}, table string) (*dynamodb.DeleteItemOutput, error) {
	ddbClient := GetDynamodbClient()
	av, marshalErr := attributevalue.MarshalMap(key)
	if marshalErr != nil {
		return &dynamodb.DeleteItemOutput{}, marshalErr
	}

	deleteItemRes, deleteItemErr := ddbClient.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key:       av,
	})
	if deleteItemErr != nil {
		return &dynamodb.DeleteItemOutput{}, deleteItemErr
	}

	return deleteItemRes, nil
}
