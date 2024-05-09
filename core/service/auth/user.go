package services

import (
	"time"

	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/repository"
	datastore "github.com/superbrobenji/budgy/infrastructure/persistence/dataStore/dynamo/user"
)

type userConfiguration func(is *userService) error

type userService struct {
	users repository.UserRepositoryReadWrite
}

// NewUserService creates a new UserService takes in a variadic number of ItemConfiguration functions
func NewUserService(cfgs ...userConfiguration) (*userService, error) {
	us := &userService{}
	for _, cfg := range cfgs {
		err := cfg(us)
		if err != nil {
			return nil, err
		}
	}
	return us, nil
}

// example of a configuration function
func withUserRepository(cr repository.UserRepositoryReadWrite) userConfiguration {
	return func(is *userService) error {
		is.users = cr
		return nil
	}
}

func WithDynamoUserRepository() userConfiguration {
	cr := datastore.NewDynamoUserRepository()
	return withUserRepository(cr)
}

func (u *userService) CreateNewUser(userId string, email string, username string) error {
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
