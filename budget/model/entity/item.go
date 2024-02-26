package entity

import (
	"github.com/google/uuid"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

type Item struct {
	ID   uuid.UUID
	Name string
	valueobject.Budget
}
