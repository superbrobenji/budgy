package entity

import (
	"github.com/google/uuid"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

type Category struct {
	ID     uuid.UUID
	Name   string
	Budget *valueobject.Budget
}
