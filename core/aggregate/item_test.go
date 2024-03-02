package aggregate_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/model/entity"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

func TestItem_NewItem(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		budget      float64
		expectedErr error
	}
	category := &entity.Category{
		ID:   uuid.New(),
		Name: "John",
		Budget: &valueobject.Budget{
			Total:     100,
			Spent:     0,
			Remaining: 100,
		},
	}
	testCases := []testCase{
		{
			test:        "should return an error when name is empty",
			name:        "",
			budget:      100,
			expectedErr: aggregate.ErrInvalidName,
		},
		{
			test:        "should return an item when name is not empty",
			name:        "John",
			budget:      100,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewItem(tc.name, tc.budget, category)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestItem_SetName(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	category := &entity.Category{
		ID:   uuid.New(),
		Name: "John",
		Budget: &valueobject.Budget{
			Total:     100,
			Spent:     0,
			Remaining: 100,
		},
	}
	item, error := aggregate.NewItem("John", 100, category)
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	testCases := []testCase{
		{
			test:        "should return an error when name is empty",
			name:        "",
			expectedErr: aggregate.ErrInvalidName,
		},
		{
			test:        "should run successfully when name is not empty",
			name:        "John",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := item.SetName(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestItem_AddTransaction(t *testing.T) {
	type testCase struct {
		test          string
		item          *valueobject.Transaction
		expectedValue []*valueobject.Transaction
		expectedErr   error
	}
	category := &entity.Category{
		ID:   uuid.New(),
		Name: "John",
		Budget: &valueobject.Budget{
			Total:     100,
			Spent:     0,
			Remaining: 100,
		},
	}
	var startVal float64 = 100
	var updateVal float64 = 100

	item, error := aggregate.NewItem("John", startVal, category)
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	transaction := &valueobject.Transaction{
		Name:   "cards",
		ID:     uuid.New(),
		Amount: updateVal,
		Date:   time.Now(),
	}

	testCases := []testCase{
		{
			test:          "should return an error when no transaction is supplied",
			item:          nil,
			expectedValue: nil,
			expectedErr:   aggregate.ErrInvalidTransaction,
		},
		{
			test:          "should remove supplied transaction from items",
			item:          transaction,
			expectedValue: []*valueobject.Transaction{transaction},
			expectedErr:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := item.AddTransaction(tc.item)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			transactions := item.GetTransactions()
			isEqual := equalTransactionSlices(transactions, tc.expectedValue)
			if !isEqual {
				t.Errorf("expected value %v, got %v", tc.expectedValue, transactions)
			}
			budget := item.GetBudget()
			if err == nil && budget.Remaining != startVal-updateVal {
				t.Errorf("expected value %v, got %v", startVal-updateVal, budget.Remaining)
			}
			if err == nil && budget.Spent != updateVal {
				t.Errorf("expected value %v, got %v", updateVal, budget.Spent)
			}
		})
	}
}

func TestItem_RemoveTransaction(t *testing.T) {
	type testCase struct {
		test          string
		item          *valueobject.Transaction
		expectedValue []*valueobject.Transaction
		expectedErr   error
	}
	var startVal float64 = 200
	var val1 float64 = 100
	var val2 float64 = 50
	category := &entity.Category{
		ID:   uuid.New(),
		Name: "John",
		Budget: &valueobject.Budget{
			Total:     100,
			Spent:     0,
			Remaining: 100,
		},
	}

	items, error := aggregate.NewItem("John", startVal, category)
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	transaction1 := &valueobject.Transaction{
		Name:   "John",
		ID:     uuid.New(),
		Amount: val1,
	}
	transaction2 := &valueobject.Transaction{
		Name:   "Sam",
		ID:     uuid.New(),
		Amount: val2,
	}
	err := items.AddTransaction(transaction1)
	if err != nil {
		t.Fatalf("unexpected error %v", error)
	}
	err2 := items.AddTransaction(transaction2)
	if err2 != nil {
		t.Fatalf("unexpected error %v", error)
	}

	testCases := []testCase{
		{
			test:          "should return an error when no transaction is supplied",
			item:          nil,
			expectedValue: []*valueobject.Transaction{transaction1, transaction2},
			expectedErr:   aggregate.ErrInvalidTransaction,
		},
		{
			test:          "should remove supplied transaction from items",
			item:          transaction1,
			expectedValue: []*valueobject.Transaction{transaction2},
			expectedErr:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := items.RemoveTransation(tc.item)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			transactions := items.GetTransactions()
			isEqual := equalTransactionSlices(transactions, tc.expectedValue)
			if !isEqual {
				t.Errorf("expected value %v, got %v", tc.expectedValue, items)
			}
			budget := items.GetBudget()
			if err == nil && budget.Remaining != startVal-val2 {
				t.Errorf("expected value %v, got %v", startVal-val2, budget.Remaining)
			}
			if err == nil && budget.Spent != val2 {
				t.Errorf("expected value %v, got %v", val2, budget.Spent)
			}

		})
	}
}

func TestItem_SetTotalBudget(t *testing.T) {
	type testCase struct {
		test        string
		budget      float64
		expectedErr error
	}
	var startVal float64 = 100
	var updateVal float64 = 200
	category := &entity.Category{
		ID:   uuid.New(),
		Name: "John",
		Budget: &valueobject.Budget{
			Total:     100,
			Spent:     0,
			Remaining: 100,
		},
	}

	item, error := aggregate.NewItem("John", startVal, category)
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	testCases := []testCase{
		{
			test:        "should return an error when budget is negative",
			budget:      -100,
			expectedErr: aggregate.ErrInvalidAmount,
		},
		{
			test:        "should update budget when budget is positive",
			budget:      updateVal,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := item.SetTotalBudget(tc.budget)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			budget := item.GetBudget()
			if err == nil && budget.Total != tc.budget {
				t.Errorf("expected value %v, got %v", tc.budget, budget.Total)
			}
			if err == nil && budget.Remaining != tc.budget-budget.Spent {
				t.Errorf("expected value %v, got %v", tc.budget-budget.Spent, budget.Remaining)
			}

		})
	}
}

func equalTransactionSlices(a, b []*valueobject.Transaction) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
