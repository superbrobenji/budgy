package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	authService "github.com/superbrobenji/budgy/core/service/auth"
	services "github.com/superbrobenji/budgy/core/service/auth"
)

func confirmResetPassword(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authService, err := services.NewAuthService(authService.WithDynamoUserRepository())
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	err = authService.ConfirmResetPassword(event.QueryStringParameters["username"], event.QueryStringParameters["confirmationCode"], event.QueryStringParameters["password"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 201,
	}, nil
}
func main() {
	lambda.Start(confirmResetPassword)
}