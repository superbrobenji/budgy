package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func signUp(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf(event.Body)
	return events.APIGatewayProxyResponse{}, nil
}
func main() {
	lambda.Start(signUp)
}
