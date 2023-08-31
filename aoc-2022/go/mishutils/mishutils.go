package mishutils

import (
	"math"
	"os"
	"path"
	"strings"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Integer | constraints.Float
}

func Max[T Numeric](s []T) T {
	if len(s) == 0 {
		return T(math.NaN())
	}

	max := s[0]

	for _, cur := range s[1:] {
		if cur > max {
			max = cur
		}
	}

	return max
}

func Sum[T Numeric](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}

	return
}

func ReadInput(name string) string {
	data := os.Getenv("AOC_DATA_PATH")

	input, err := os.ReadFile(path.Join(data, name+".txt"))

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(input))
}
