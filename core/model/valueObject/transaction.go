package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID     uuid.UUID
	Name   string
	Amount float64
	Date   time.Time
}
