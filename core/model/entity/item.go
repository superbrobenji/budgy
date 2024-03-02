package entity

import (
	"github.com/google/uuid"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

type Item struct {
	ID     uuid.UUID
	Name   string
	Budget *valueobject.Budget
}
