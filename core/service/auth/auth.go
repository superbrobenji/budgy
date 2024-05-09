package services

import (
	"errors"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	sdk "github.com/superbrobenji/budgy/infrastructure/aws/sdk/cognito"
)

type authService struct {
	cognito *sdk.CognitoActions
}

var (
	ErrInvalidCridentials = errors.New("all cridentials must be passed")
	ErrUserNotConfirmed   = errors.New("there was an error confirming the user")
)

func NewAuthService() *authService {
	return &authService{
		cognito: sdk.NewCognitoClient(),
	}
}

func (c *authService) SignUp(username string, password string, email string) (bool, error) {
	if username == "" || password == "" || email == "" {
		return false, ErrInvalidCridentials
	}
	userConfirmed, err := c.cognito.SignUp(username, password, email)
	return userConfirmed, err
}

func (c *authService) ConfirmSignUp(username string, confirmationCode string) error {
	if username == "" || confirmationCode == "" {
		return ErrInvalidCridentials
	}
	_, err := c.cognito.ConfirmSignUp(username, confirmationCode)
	if err != nil {
		return err
	}

	// auth, err := c.cognito.SignIn(username, password)
	// if err != nil {
	// 	return nil, err
	// }

	return nil
}

func (c *authService) Login(username string, password string) (*types.AuthenticationResultType, error) {
	if username == "" || password == "" {
		return nil, ErrInvalidCridentials
	}
	auth, err := c.cognito.SignIn(username, password)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
