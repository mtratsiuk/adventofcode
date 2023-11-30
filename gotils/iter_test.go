package gotils

import (
	"reflect"
	"strings"
	"testing"
)

func TestMap(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := []int{2, 3, 4, 5, 6}

	mapped := Collect(
		Map(
			Iter(data),
			func(x int) int { return x + 1 }))

	if !reflect.DeepEqual(mapped, expected) {
		t.Errorf("TestMap: expected %v, got %v", expected, mapped)
	}
}

func TestFold(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	expected := 20

	mapped := Map(Iter(data), func(x int) int { return x + 1 })
	folded := Fold(mapped, 0, func(r int, x int) int { return r + x })

	if !reflect.DeepEqual(folded, expected) {
		t.Errorf("TestFold: expected %v, got %v", expected, folded)
	}
}

func TestFlatMap(t *testing.T) {
	data := "1,2|3,4|5,6"
	expected := []int{1, 2, 3, 4, 5, 6}

	pairs := Iter(strings.Split(data, "|"))
	numbers := FlatMap(pairs, func(s string) []string { return strings.Split(s, ",") })
	parsed := Collect(Map(numbers, MustParseInt))

	if !reflect.DeepEqual(parsed, expected) {
		t.Errorf("TestFlatMap: expected %v, got %v", expected, parsed)
	}
}

func TestChunks(t *testing.T) {
	data := "1,2|3,4|5,6"
	expected := []int{3, 7, 11}

	pairs := Iter(strings.Split(data, "|"))
	numbers := FlatMap(pairs, func(s string) []string { return strings.Split(s, ",") })
	parsed := Map(numbers, MustParseInt)
	parsedPairs := Chunks(parsed, 2)
	sums := Map(parsedPairs, func(p []int) int { return p[0] + p[1] })
	res := Collect(sums)

	if !reflect.DeepEqual(res, expected) {
		t.Errorf("TestFlatMap: expected %v, got %v", expected, res)
	}
}
