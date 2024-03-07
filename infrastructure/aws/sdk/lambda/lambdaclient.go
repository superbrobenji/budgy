package sdk

import (
	"bytes"
	"context"
	"errors"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

// FunctionWrapper encapsulates function actions used in the examples.
// It contains an AWS Lambda service client that is used to perform user actions.
type FunctionWrapper struct {
	LambdaClient *lambda.Client
}

// CreateFunction creates a new Lambda function from code contained in the zipPackage
// buffer. The specified handlerName must match the name of the file and function
// contained in the uploaded code. The role specified by iamRoleArn is assumed by
// Lambda and grants specific permissions.
// When the function already exists, types.StateActive is returned.
// When the function is created, a lambda.FunctionActiveV2Waiter is used to wait until the
// function is active.
func (wrapper FunctionWrapper) CreateFunction(functionName string, handlerName string,
	iamRoleArn *string, zipPackage *bytes.Buffer) types.State {
	var state types.State
	_, err := wrapper.LambdaClient.CreateFunction(context.TODO(), &lambda.CreateFunctionInput{
		Code:         &types.FunctionCode{ZipFile: zipPackage.Bytes()},
		FunctionName: aws.String(functionName),
		Role:         iamRoleArn,
		Handler:      aws.String(handlerName),
		Publish:      true,
		Runtime:      types.RuntimePython38,
	})
	if err != nil {
		var resConflict *types.ResourceConflictException
		if errors.As(err, &resConflict) {
			log.Printf("Function %v already exists.\n", functionName)
			state = types.StateActive
		} else {
			log.Panicf("Couldn't create function %v. Here's why: %v\n", functionName, err)
		}
	} else {
		waiter := lambda.NewFunctionActiveV2Waiter(wrapper.LambdaClient)
		funcOutput, err := waiter.WaitForOutput(context.TODO(), &lambda.GetFunctionInput{
			FunctionName: aws.String(functionName)}, 1*time.Minute)
		if err != nil {
			log.Panicf("Couldn't wait for function %v to be active. Here's why: %v\n", functionName, err)
		} else {
			state = funcOutput.Configuration.State
		}
	}
	return state
}
