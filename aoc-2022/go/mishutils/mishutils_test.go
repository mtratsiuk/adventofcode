package mishutils

import (
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		in  []int
		max int
		ok  bool
	}{
		{[]int{}, 0, false},
		{[]int{1, 2, 3}, 3, true},
		{[]int{3, 2, 1}, 3, true},
	}

	for _, c := range cases {
		if max, ok := Max(c.in); max != c.max || ok != c.ok {
			t.Errorf("Max(%v) != %v, %v", c.in, c.max, c.ok)
		}
	}
}

func TestSum(t *testing.T) {
	cases := []struct {
		in  []int
		sum int
	}{
		{[]int{}, 0},
		{[]int{1, 2, 3}, 6},
		{[]int{1, -2, 3}, 2},
	}

	for _, c := range cases {
		if sum := Sum(c.in); sum != c.sum {
			t.Errorf("Sum(%v) != %v", c.in, c.sum)
		}
	}
}
