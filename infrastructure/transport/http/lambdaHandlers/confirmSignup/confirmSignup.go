package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	services "github.com/superbrobenji/budgy/core/service/auth"
)

func confirmSignUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authService := services.NewAuthService()
	err := authService.ConfirmSignUp(event.QueryStringParameters["username"], event.QueryStringParameters["confirmationCode"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 202,
	}, nil
}
func main() {
	lambda.Start(confirmSignUp)
}
