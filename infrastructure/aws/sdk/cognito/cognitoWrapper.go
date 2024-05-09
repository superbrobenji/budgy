package sdk

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
}

func NewCognitoClient() *CognitoActions {
	return &CognitoActions{
		CognitoClient: GetCognitoClient(),
	}
}

// SignUp signs up a user with Amazon Cognito.
func (actor CognitoActions) SignUp(userName string, password string, userEmail string) (*cognitoidentityprovider.SignUpOutput, error) {
	output, err := actor.CognitoClient.SignUp(context.TODO(), &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID")),
		Password: aws.String(password),
		Username: aws.String(userName),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("email"), Value: aws.String(userEmail)},
		},
	})
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't sign up user %v. Here's why: %v\n", userName, err)
		}
	}
	return output, err
}

func (actor CognitoActions) SignIn(userName string, password string) (*types.AuthenticationResultType, error) {
	var authResult *types.AuthenticationResultType
	output, err := actor.CognitoClient.InitiateAuth(context.TODO(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID")),
		AuthParameters: map[string]string{"USERNAME": userName, "PASSWORD": password},
	})
	if err != nil {
		var resetRequired *types.PasswordResetRequiredException
		if errors.As(err, &resetRequired) {
			log.Println(*resetRequired.Message)
		} else {
			log.Printf("Couldn't sign in user %v. Here's why: %v\n", userName, err)
		}
	} else {
		authResult = output.AuthenticationResult
	}
	return authResult, err
}

// ForgotPassword starts a password recovery flow for a user. This flow typically sends a confirmation code
// to the user's configured notification destination, such as email.
func (actor CognitoActions) ForgotPassword(userName string) (*types.CodeDeliveryDetailsType, error) {
	output, err := actor.CognitoClient.ForgotPassword(context.TODO(), &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID")),
		Username: aws.String(userName),
	})
	if err != nil {
		log.Printf("Couldn't start password reset for user '%v'. Here;s why: %v\n", userName, err)
	}
	return output.CodeDeliveryDetails, err
}

// ConfirmForgotPassword confirms a user with a confirmation code and a new password.
func (actor CognitoActions) ConfirmForgotPassword(code string, userName string, password string) error {
	_, err := actor.CognitoClient.ConfirmForgotPassword(context.TODO(), &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(os.Getenv("AWS_COGNITO_CLIENT_ID")),
		ConfirmationCode: aws.String(code),
		Password:         aws.String(password),
		Username:         aws.String(userName),
	})
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't confirm user %v. Here's why: %v", userName, err)
		}
	}
	return err
}
