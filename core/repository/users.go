package repository

import (
	"errors"
	"time"

	"github.com/superbrobenji/budgy/core/aggregate"
)

var (
	ErrUserNotFound    = errors.New("User not found")
	ErrFailedToAddUser = errors.New("failed to add User")
	ErrDeleteUser      = errors.New("failed to delete User")
	ErrUserExists      = errors.New("User already exists")
)

type UserRepositoryWrite interface {
	CreateUser(*aggregate.User) error
	DeleteUser(string) error
}
type UserRepositoryRead interface {
	GetUserByID(string) (*aggregate.User, error)
	GetUserByUsername(string) (*aggregate.User, error)
	GetUserByEmail(string) (*aggregate.User, error)
	GetUsersByDateJoined(time.Time, time.Time) (*[]*aggregate.User, error)
}

type UserRepositoryReadWrite interface {
	CreateUser(*aggregate.User) error
	DeleteUser(string) error
	GetUserByID(string) (*aggregate.User, error)
	GetUserByUsername(string) (*aggregate.User, error)
	GetUserByEmail(string) (*aggregate.User, error)
	GetUsersByDateJoined(time.Time, time.Time) (*[]*aggregate.User, error)
}
