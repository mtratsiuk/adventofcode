package mishutils

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	cases := []struct {
		in  []float64
		max float64
	}{
		{[]float64{1, 2, 3}, 3},
		{[]float64{3, 2, 1}, 3},
	}

	for _, c := range cases {
		if max := Max(c.in); max != c.max {
			t.Errorf("Max(%v) != %v", c.in, c.max)
		}
	}
}

func TestMax_whenSliceIsEmpty(t *testing.T) {
	out := Max([]float64{})

	if !math.IsNaN(out) {
		t.Errorf("Max([]) != NaN")
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
