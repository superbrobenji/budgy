package aggregate

import (
	"time"

	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type User struct {
	user *valueobject.User
}

func NewUser(userId string, username string, dateJoined time.Time, email string) (User, error) {
	if username == "" {
		return User{}, ErrInvalidName
	}

	if email == "" {
		return User{}, ErrInvalidName
	}

	user := &valueobject.User{
		Username:   username,
		Email:      email,
		UserID:     userId,
		DateJoined: dateJoined,
	}

	return User{
		user,
	}, nil
}

func (u *User) GetID() string {
	return u.user.UserID
}

func (u *User) GetUsername() string {
	return u.user.Username
}
func (u *User) GetDateJoined() time.Time {
	return u.user.DateJoined
}
func (u *User) GetEmail() string {
	return u.user.Email
}
