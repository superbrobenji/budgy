package aggregate_test

import (
	"testing"

	"github.com/superbrobenji/budgy/core/aggregate"
)

func TestTransaction_NewTransaction(t *testing.T) {
	type testCase struct {
		test        string
		name        string
		amount      float64
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "should return an error when name is empty",
			name:        "",
			amount:      100,
			expectedErr: aggregate.ErrInvalidName,
		},
		{
			test:        "should return an error when amount is negative",
			name:        "John",
			amount:      -100,
			expectedErr: aggregate.ErrInvalidAmount,
		},
		{
			test:        "should return a Transaction when name and is not empty",
			name:        "John",
			amount:      100,
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewTransaction(tc.name, tc.amount)
			if err != tc.expectedErr {
				t.Errorf("expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
