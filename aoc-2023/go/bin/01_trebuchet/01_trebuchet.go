package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("01_trebuchet")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(input string) int {
	sum := 0

	for _, line := range strings.Fields(input) {
		digits := make([]rune, 0)

		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, c)
			}
		}

		first := string(digits[0])
		last := string(digits[len(digits)-1])

		sum += gotils.MustParseInt(first + last)
	}

	return sum
}

var digitsMap = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func solve2(input string) int {
	sum := 0

	for _, line := range strings.Fields(input) {
		fmt.Println(line)

		for i := 0; i < len(line); i += 1 {
			for k, v := range digitsMap {
				if strings.HasPrefix(line[i:], k) {
					line = strings.Replace(line, k, v, 1)
					break
				}
			}
		}

		fmt.Println(line)

		digits := make([]rune, 0)

		for _, c := range line {
			if unicode.IsDigit(c) {
				digits = append(digits, c)
			}
		}

		first := string(digits[0])
		last := string(digits[len(digits)-1])
		fmt.Println(first + last)
		sum += gotils.MustParseInt(first + last)
	}

	return sum
}
