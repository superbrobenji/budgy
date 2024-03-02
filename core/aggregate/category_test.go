package aggregate_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/superbrobenji/budgy/core/aggregate"
	"github.com/superbrobenji/budgy/core/model/entity"
	valueobject "github.com/superbrobenji/budgy/core/model/valueObject"
)

func TestCategory_NewCategory(t *testing.T) {
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
			test:        "should return an error when name is income",
			name:        "income",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is Income",
			name:        "Income",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is Pre-Income Deductions",
			name:        "Pre-Income Deductions",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is pre-income deductions",
			name:        "pre-income deductions",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return a customer when name is not empty",
			name:        "John",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCategory(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCategory_SetName(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	category, error := aggregate.NewCategory("John")
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
			test:        "should return an error when name is income",
			name:        "income",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is Income",
			name:        "Income",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is Pre-Income Deductions",
			name:        "Pre-Income Deductions",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should return an error when name is pre-income deductions",
			name:        "pre-income deductions",
			expectedErr: aggregate.ErrReservedCategoryNameSpace,
		},
		{
			test:        "should run successfully when name is not empty",
			name:        "John",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := category.SetName(tc.name)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestCategory_AddItem(t *testing.T) {
	type testCase struct {
		test          string
		item          *entity.Item
		expectedValue []*entity.Item
		expectedErr   error
	}
	var total float64 = 100
	var spent float64 = 0
	var remaining float64 = 100

	category, error := aggregate.NewCategory("John")
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	item1 := &entity.Item{
		Name: "John",
		ID:   uuid.New(),
		Budget: &valueobject.Budget{
			Total:     total,
			Spent:     spent,
			Remaining: remaining,
		},
	}

	testCases := []testCase{
		{
			test:          "should return an error when no item is supplied",
			item:          nil,
			expectedValue: nil,
			expectedErr:   aggregate.ErrInvalidItem,
		},
		{
			test:          "should remove supplied item from items",
			item:          item1,
			expectedValue: []*entity.Item{item1},
			expectedErr:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := category.AddItem(tc.item)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			items := category.GetItems()
			isEqual := equalItemSlices(items, tc.expectedValue)
			if !isEqual {
				t.Errorf("expected value %v, got %v", tc.expectedValue, items)
			}
			budget := category.GetBudget()
			if err == nil && budget.Total != total {
				t.Errorf("expected value %v, got %v", total, budget.Total)
			}
			if err == nil && budget.Spent != spent {
				t.Errorf("expected value %v, got %v", spent, budget.Spent)
			}
			if err == nil && budget.Remaining != remaining {
				t.Errorf("expected value %v, got %v", remaining, budget.Remaining)
			}
		})
	}
}

func TestCategory_RemoveItem(t *testing.T) {
	type testCase struct {
		test          string
		item          *entity.Item
		expectedValue []*entity.Item
		expectedErr   error
	}
	budget1 := &valueobject.Budget{
		Total:     100,
		Spent:     50,
		Remaining: 50,
	}
	budget2 := &valueobject.Budget{
		Total:     200,
		Spent:     0,
		Remaining: 200,
	}

	category, error := aggregate.NewCategory("John")
	if error != nil {
		t.Fatalf("unexpected error %v", error)
	}
	item1 := &entity.Item{
		Name:   "John",
		ID:     uuid.New(),
		Budget: budget1,
	}
	item2 := &entity.Item{
		Name:   "Sam",
		ID:     uuid.New(),
		Budget: budget2,
	}
	err := category.AddItem(item1)
	if err != nil {
		t.Fatalf("unexpected error %v", error)
	}
	err2 := category.AddItem(item2)
	if err2 != nil {
		t.Fatalf("unexpected error %v", error)
	}

	testCases := []testCase{
		{
			test:          "should return an error when no item is supplied",
			item:          nil,
			expectedValue: []*entity.Item{item1, item2},
			expectedErr:   aggregate.ErrInvalidItem,
		},
		{
			test:          "should remove supplied item from items",
			item:          item1,
			expectedValue: []*entity.Item{item2},
			expectedErr:   nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			err := category.RemoveItem(tc.item)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
			items := category.GetItems()
			isEqual := equalItemSlices(items, tc.expectedValue)
			if !isEqual {
				t.Errorf("expected value %v, got %v", tc.expectedValue, items)
			}
			if err == nil && category.GetBudget().Total != 200 {
				t.Errorf("expected value %v, got %v", 200, category.GetBudget().Total)
			}
			if err == nil && category.GetBudget().Spent != 0 {
				t.Errorf("expected value %v, got %v", 0, category.GetBudget().Spent)
			}
			if err == nil && category.GetBudget().Remaining != 200 {
				t.Errorf("expected value %v, got %v", 200, category.GetBudget().Remaining)
			}
		})
	}
}

func equalItemSlices(a, b []*entity.Item) bool {
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
