#!/usr/bin/env bash

set -euo pipefail

. ./bin/env.sh

cd "$(dirname "$0")"/..

AOC_YEAR=${AOC_YEAR:-2022}
AOC_TOKEN=${AOC_TOKEN}

DAY_NUMBER=$1
DAY_NUMBER_PADDED=$(printf "%02d" $DAY_NUMBER)
DAY_NAME=$(echo "$2" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')

INPUT_URL="https://adventofcode.com/${AOC_YEAR}/day/${DAY_NUMBER}/input"
FILE_NAME="${DAY_NUMBER_PADDED}_${DAY_NAME}"
INPUT_PATH="./aoc-${AOC_YEAR}/data/${FILE_NAME}.txt"
SOLUTION_DIR_PATH="./aoc-${AOC_YEAR}/go/bin/${FILE_NAME}"
SOLUTION_PATH="${SOLUTION_DIR_PATH}/${FILE_NAME}.go"
TEST_PATH="${SOLUTION_DIR_PATH}/${FILE_NAME}_test.go"


curl --cookie "session=${AOC_TOKEN}" $INPUT_URL > $INPUT_PATH

mkdir -p $SOLUTION_DIR_PATH

echo 'package main

import (
	"fmt"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput('"\"${FILE_NAME}\""')

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0

	return sum
}

func solve2(in string) int {
	sum := 0

	return sum
}
' > $SOLUTION_PATH

echo 'package main

import (
	"testing"
)

func Test1(t *testing.T) {
	in := `
data
`

	expected := 0

	if res := solve1(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve1() failed!\nexpected: %v\nactual: %v", expected, res)
	}
}

func Test2(t *testing.T) {
	in := `
data
`

	expected := 0

	if res := solve2(strings.TrimSpace(in)); res != expected {
		t.Errorf("\nsolve2() failed!\nexpected: %v\nactual:%v", expected, res)
	}
}
' > $TEST_PATH

go fmt $SOLUTION_DIR_PATH
