package authService

import (
	"errors"
	"time"

	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/repository"
	sdk "github.com/superbrobenji/budgy/infrastructure/aws/sdk/cognito"
	datastore "github.com/superbrobenji/budgy/infrastructure/persistence/dataStore/dynamo/user"
)

type authService struct {
	cognito *sdk.CognitoActions
	users   repository.UserRepositoryReadWrite
}
type AuthResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type authConfiguration func(is *authService) error

var (
	ErrInvalidCridentials = errors.New("all cridentials must be passed")
	ErrUserNotConfirmed   = errors.New("there was an error confirming the user")
)

func NewAuthService(cfgs ...authConfiguration) (*authService, error) {
	as := &authService{cognito: sdk.NewCognitoClient()}
	for _, cfg := range cfgs {
		err := cfg(as)
		if err != nil {
			return nil, err
		}
	}
	return as, nil
}

func withUserRepository(cr repository.UserRepositoryReadWrite) authConfiguration {
	return func(is *authService) error {
		is.users = cr
		return nil
	}
}

func WithDynamoUserRepository() authConfiguration {
	cr := datastore.NewDynamoUserRepository()
	return withUserRepository(cr)
}

func (c *authService) SignUp(username string, password string, email string) (bool, error) {
	if username == "" || password == "" || email == "" {
		return false, ErrInvalidCridentials
	}

	output, err := c.cognito.SignUp(username, password, email)
	if err != nil {
		return false, err
	}

	err = c.createNewUser(*output.UserSub, email, username)
	if err != nil && err != repository.ErrUserExists {
		return false, err
	}

	return output.UserConfirmed, err
}

func (c *authService) ConfirmSignUp(username string, confirmationCode string) error {
	if username == "" || confirmationCode == "" {
		return ErrInvalidCridentials
	}
	_, err := c.cognito.ConfirmSignUp(username, confirmationCode)
	if err != nil {
		return err
	}

	return nil
}

func (c *authService) ResendConfirmationCode(username string) error {
	if username == "" {
		return ErrInvalidCridentials
	}
	_, err := c.cognito.ResendConfirmationCode(username)
	if err != nil {
		return err
	}
	return nil
}

func (c *authService) Login(username string, password string) (*AuthResponse, error) {
	if username == "" || password == "" {
		return nil, ErrInvalidCridentials
	}
	auth, err := c.cognito.SignIn(username, password)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{AccessToken: *auth.AccessToken, RefreshToken: *auth.RefreshToken}, nil
}

func (u *authService) createNewUser(userId string, email string, username string) error {
	dateJoined := time.Now()
	existingUser, getUserErr := u.users.GetUserByID(userId)
	if getUserErr != nil {
		return getUserErr
	}

	if *existingUser != (aggregate.User{}) {
		return repository.ErrUserExists
	}

	newUser, getUserErr := aggregate.NewUser(userId, username, dateJoined, email)
	if getUserErr != nil {
		return getUserErr
	}
	createUserErr := u.users.CreateUser(&newUser)
	if createUserErr != nil {
		return createUserErr
	}
	return nil
}

func (u *authService) ResetPassword(username string) error {
	if username == "" {
		return ErrInvalidCridentials
	}
	_, error := u.cognito.ForgotPassword(username)
	if error != nil {
		return error
	}

	return nil
}
func (u *authService) ConfirmResetPassword(username string, code string, password string) error {
	if username == "" || code == "" || password == "" {
		return ErrInvalidCridentials
	}
	error := u.cognito.ConfirmForgotPassword(code, username, password)
	if error != nil {
		return error
	}
	return nil
}
