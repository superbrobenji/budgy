package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/superbrobenji/budgy/core/repository"
	services "github.com/superbrobenji/budgy/core/service"
)

func saveUserToDB(event events.CognitoEventUserPoolsPostAuthentication) (events.CognitoEventUserPoolsPostAuthentication, error) {
	userService, err := services.NewUserService(services.WithDynamoUserRepository())
	if err != nil {
		return event, err
	}
	err = userService.CreateNewUser(event.Request.UserAttributes["sub"], event.Request.UserAttributes["email"], event.UserName)
	if err != nil && err != repository.ErrUserExists {
		return event, err
	}

	return event, nil
}

func main() {
	lambda.Start(saveUserToDB)
}
