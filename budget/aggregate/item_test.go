package aggregate_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/budget/aggregate"
	valueobject "github.com/superbrobenji/budgy/budget/model/valueObject"
)

func TestItem_NewItem(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "should return an error when name is empty",
			name:        "",
			expectedErr: aggregate.ErrInvalidName,
		},
		{
			test:        "should return a customer when name is not empty",
			name:        "John",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewItem(tc.name)
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
	item, error := aggregate.NewItem("John")
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
			test:        "should return a customer when name is not empty",
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

	item, error := aggregate.NewItem("John")
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	transaction := &valueobject.Transaction{
		Name: "John",
		ID:   uuid.New(),
        Amount: 100,
        Date: time.Now(),
	}

	testCases := []testCase{
		{
			test:          "should return an error when no item is supplied",
			item:          nil,
			expectedValue: nil,
			expectedErr:   aggregate.ErrInvalidTransaction,
		},
		{
			test:          "should remove supplied item from items",
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
		})
	}
}

func TestCustomer_RemoveTransaction(t *testing.T) {
	type testCase struct {
		test          string
		item          *valueobject.Transaction
		expectedValue []*valueobject.Transaction
		expectedErr   error
	}

	items, error := aggregate.NewItem("John")
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	transaction1 := &valueobject.Transaction{
		Name: "John",
		ID:   uuid.New(),
	}
	transaction2 := &valueobject.Transaction{
		Name: "Sam",
		ID:   uuid.New(),
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
			test:          "should return an error when no item is supplied",
			item:          nil,
			expectedValue: []*valueobject.Transaction{transaction1, transaction2},
			expectedErr:   aggregate.ErrInvalidTransaction,
		},
		{
			test:          "should remove supplied item from items",
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
