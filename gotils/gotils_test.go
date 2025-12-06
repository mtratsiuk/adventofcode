package gotils

import (
	"fmt"
	"testing"
)

func TestDigits(t *testing.T) {
	testCases := []struct {
		val      int
		expected int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{999, 3},
		{1000, 4},
		{12345, 5},
		{999999999, 9},
		{-123, 3},
		{1000000, 7},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("DigitsCount(%d)", tc.val), func(t *testing.T) {
			actual := DigitsCount(tc.val)
			if actual != tc.expected {
				t.Errorf("DigitsCount(%d) = %d; expected %d", tc.val, actual, tc.expected)
			}
		})
	}
}
