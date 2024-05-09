package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	services "github.com/superbrobenji/budgy/core/service/auth"
)

func signUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	authService := services.NewAuthService()
	fmt.Print(event.QueryStringParameters)
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
		return events.APIGatewayProxyResponse{
			StatusCode: 200,
			Body:       *auth.AccessToken,
		}, nil
	}
	return events.APIGatewayProxyResponse{
		StatusCode: 202,
	}, nil
}
func main() {
	lambda.Start(signUp)
}
