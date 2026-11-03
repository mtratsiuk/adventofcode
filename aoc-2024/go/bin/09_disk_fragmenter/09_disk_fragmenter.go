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
	return Checksum(Compact(DiskMapToBlocks(in)))
}

func solve2(in string) int {
	sum := 0

	return sum
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

func Checksum(blocks []int) int {
	sum := 0

	for i, block := range blocks {
		if block == -1 {
			break
		}

		sum += i * block
	}

	return sum
}
