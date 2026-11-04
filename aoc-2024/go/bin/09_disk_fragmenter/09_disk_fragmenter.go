package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("09_disk_fragmenter")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	return Checksum(Compact(DiskMapToBlocks(in)), false)
}

func solve2(in string) int {
	blocks := DiskMapToBlocks(in)
	right := len(blocks) - 1

	for right != 0 {
		span, nextRight := FindBlocksSpanToCompactFromTheRight(blocks, right)
		right = nextRight

		if len(span) != 0 {
			emptySpan := FindEmptySpanMoreOrEqualSizeFromTheLeft(blocks, len(span), right+1)

			if len(emptySpan) != 0 {
				copy(emptySpan, span)
				copy(span, slices.Repeat([]int{-1}, len(span)))
			}
		}
	}

	return Checksum(blocks, true)
}

func DiskMapToBlocks(in string) []int {
	blocks := make([]int, 0, len(in))

	id := 0
	for i, s := range strings.Split(in, "") {
		n := gotils.MustParseInt(s)

		var idToAppend int

		if i%2 == 0 {
			idToAppend = id
			id += 1
		} else {
			idToAppend = -1
		}

		blocks = append(blocks, slices.Repeat([]int{idToAppend}, n)...)
	}

	return blocks
}

func Compact(blocks []int) []int {
	left, right := 0, len(blocks)-1

	for left < right {
		if blocks[left] != -1 {
			left += 1
			continue
		}

		if blocks[right] == -1 {
			right -= 1
			continue
		}

		blocks[left] = blocks[right]
		blocks[right] = -1
	}

	return blocks
}

func Checksum(blocks []int, fragmented bool) int {
	sum := 0

	for i, block := range blocks {
		if block == -1 {
			if fragmented {
				continue
			} else {
				break
			}
		}

		sum += i * block
	}

	return sum
}

func FindBlocksSpanToCompactFromTheRight(blocks []int, right int) ([]int, int) {
	for blocks[right] == -1 && right > 0 {
		right -= 1
	}

	fileId := blocks[right]

	if fileId == -1 {
		return []int{}, 0
	}

	for i := right - 1; i >= 0; i -= 1 {
		if i == 0 {
			return blocks[i : right+1], i
		}

		if blocks[i] != fileId {
			return blocks[i+1 : right+1], i
		}
	}

	panic("shouldn't be reachable")
}

func FindEmptySpanMoreOrEqualSizeFromTheLeft(blocks []int, size int, until int) []int {
	gotils.Assert(until < len(blocks), func() string { return "unexpected `until` out of bounds" })

	left := 0
	right := left + 1

	for ; left < until; left += 1 {
		if blocks[left] != -1 {
			continue
		}

		for right = left + 1; blocks[right] == -1; right += 1 {
		}

		if right-left >= size {
			return blocks[left : left+size]
		}

		left = right
	}

	return []int{}
}
