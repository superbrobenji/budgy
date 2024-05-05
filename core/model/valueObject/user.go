package valueobject

import (
	"time"
)

type User struct {
	UserID     string
	Email      string
	Username   string
	DateJoined time.Time
}
