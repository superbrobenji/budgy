package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func saveUserToDB(event events.CognitoEventUserPoolsPostAuthentication) (events.CognitoEventUserPoolsPostAuthentication, error) {
	fmt.Println(event)
	return event, nil
}

func main() {
	lambda.Start(saveUserToDB)
}
