package main

import (
	"fmt"
	"iter"
	"math"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("02_gift_shop")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0
	jobs := 0
	ids := make(chan int)

	for l, r := range ranges(in) {
		jobs += 1
		go findInvalidIds(ids, l, r, 2)
	}

	for jobs > 0 {
		r := <-ids

		if r == -1 {
			jobs -= 1
		} else {
			sum += r
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0
	jobs := 0
	ids := make(chan int)
	seen := gotils.NewSet[int]()

	for l, r := range ranges(in) {
		for i := gotils.DigitsCount(r); i >= 2; i -= 1 {
			jobs += 1
			go findInvalidIds(ids, l, r, i)
		}
	}

	for jobs > 0 {
		r := <-ids

		if r == -1 {
			jobs -= 1
		} else {
			if !seen.Has(r) {
				sum += r
			}

			seen.Add(r)
		}
	}

	return sum
}

func findInvalidIds(ids chan<- int, l, r, n int) {
	defer func() {
		ids <- -1
	}()

	for l <= r {
		digits := gotils.DigitsCount(l)

		if digits%n != 0 {
			l = int(math.Pow10(digits))
			continue
		}

		div := int(math.Pow10(digits / n))
		cur := l / div
		group := l % div
		passed := true

		for cur > 0 {
			nextGroup := cur % div

			if nextGroup != group {
				passed = false
				break
			}

			cur /= div
		}

		if passed {
			ids <- l
		}

		l += 1
	}
}

func ranges(in string) iter.Seq2[int, int] {
	return func(yield func(int, int) bool) {
		for rangeStr := range strings.SplitSeq(in, ",") {
			l, r, _ := strings.Cut(rangeStr, "-")
			if !yield(gotils.MustParseInt(l), gotils.MustParseInt(r)) {
				return
			}
		}
	}
}
