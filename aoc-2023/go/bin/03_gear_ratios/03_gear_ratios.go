package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("03_gear_ratios")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(input string) int {
	sum := 0

	scheme := strings.Split(input, "\n")
	width := len(scheme[0])
	height := len(scheme)

	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			cur := scheme[y][x]

			if !isDigit(cur) {
				continue
			}

			start := x
			end := x + 1

			for {
				if end >= width || !isDigit(scheme[y][end]) {
					break
				}

				end += 1
			}

			x = end - 1

			if hasAdjSymbol(scheme, y, start, end-1) {
				sum += gotils.MustParseInt(scheme[y][start:end])
			}
		}
	}

	return sum
}

func solve2(input string) int {
	sum := 0

	scheme := strings.Split(input, "\n")
	width := len(scheme[0])
	height := len(scheme)

	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			cur := scheme[y][x]

			if !isStar(cur) {
				continue
			}

			parts := getAdjNumberParts(scheme, x, y)

			if len(parts) == 2 {
				sum += parts[0] * parts[1]
			}
		}
	}

	return sum
}

const dot = byte('.')
const star = byte('*')

func isSkip(c byte) bool {
	return c == dot
}

func isSymbol(c byte) bool {
	return !isSkip(c) && !unicode.IsDigit(rune(c))
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func isStar(c byte) bool {
	return c == star
}

func hasAdjSymbol(scheme []string, line, start, end int) bool {
	width := len(scheme[0])
	height := len(scheme)

	for y := max(0, line-1); y <= min(height-1, line+1); y += 1 {
		for x := max(0, start-1); x <= min(width-1, end+1); x += 1 {
			if isSymbol(scheme[y][x]) {
				return true
			}
		}
	}

	return false
}

func readNumberFrom(scheme []string, line, start int) (string, int, int) {
	width := len(scheme[0])
	left := start
	right := start + 1

	if !isDigit(scheme[line][start]) {
		return "", start, start
	}

	for ; left >= 0 && isDigit(scheme[line][left]); left -= 1 {
	}
	for ; right < width && isDigit(scheme[line][right]); right += 1 {
	}

	return scheme[line][left+1 : right], left + 1, right - 1
}

func getAdjNumberParts(scheme []string, x, y int) []int {
	height := len(scheme)
	parts := make([]int, 0)

	if y > 0 {
		topLeft, _, r := readNumberFrom(scheme, y-1, x-1)

		if topLeft != "" {
			parts = append(parts, gotils.MustParseInt(topLeft))
		}

		if r < x {
			topRight, _, _ := readNumberFrom(scheme, y-1, x+1)

			if topRight != "" {
				parts = append(parts, gotils.MustParseInt(topRight))
			}
		}
	}

	if y < height {
		bottomLeft, _, r := readNumberFrom(scheme, y+1, x-1)

		if bottomLeft != "" {
			parts = append(parts, gotils.MustParseInt(bottomLeft))
		}

		if r < x {
			bottomRight, _, _ := readNumberFrom(scheme, y+1, x+1)

			if bottomRight != "" {
				parts = append(parts, gotils.MustParseInt(bottomRight))
			}
		}
	}

	left, _, _ := readNumberFrom(scheme, y, x-1)

	if left != "" {
		parts = append(parts, gotils.MustParseInt(left))
	}

	right, _, _ := readNumberFrom(scheme, y, x+1)

	if right != "" {
		parts = append(parts, gotils.MustParseInt(right))
	}

	return parts
}
