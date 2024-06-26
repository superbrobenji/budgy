package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	authService "github.com/superbrobenji/budgy/core/service/auth"
	services "github.com/superbrobenji/budgy/core/service/auth"
)

func signUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authService, err := services.NewAuthService(authService.WithDynamoUserRepository())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	userConfirmed, err := authService.SignUp(event.QueryStringParameters["username"], event.QueryStringParameters["password"], event.QueryStringParameters["email"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	if userConfirmed != false {
		auth, err := authService.Login(event.QueryStringParameters["username"], event.QueryStringParameters["password"])
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       err.Error(),
			}, nil
		}
		marshalledAuth, err := json.Marshal(auth)
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       err.Error(),
			}, nil
		}
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       string(marshalledAuth),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}
func main() {
	lambda.Start(signUp)
}
