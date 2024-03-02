package aggregate

import "errors"

var (
	ErrInvalidName               = errors.New("a valid name must be supplied")
	ErrReservedCategoryNameSpace = errors.New("'Income' and 'Pre-Income Deductions' are reserved names")
	ErrUnInitialised             = errors.New("aggregate not initialised")
	ErrInvalidItem               = errors.New("a valid item must be supplied")
	ErrInvalidTransaction        = errors.New("a valid transaction must be supplied")
	ErrInvalidAmount             = errors.New("a valid positive amount must be supplied")
)
