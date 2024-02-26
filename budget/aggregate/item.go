package aggregate

import (
	"github.com/superbrobenji/budgy/budget/model/entity"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

type Item struct {
	Item         *entity.Item
	Transactions []*valueobject.Transaction
}
