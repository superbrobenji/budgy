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

func (c *authService) SignUp(username string, password string, email string) (*types.AuthenticationResultType, error) {
	if username == "" || password == "" || email == "" {
		return nil, ErrInvalidCridentials
	}
	userConfirmed, err := c.cognito.SignUp(username, password, email)
	if err != nil {
		return nil, err
	}
	if userConfirmed == false {
		return nil, ErrUserNotConfirmed
	}
	auth, err := c.cognito.SignIn(username, password)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
