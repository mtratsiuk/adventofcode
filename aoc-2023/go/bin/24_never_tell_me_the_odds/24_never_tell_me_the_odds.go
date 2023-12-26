package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("24_never_tell_me_the_odds")

	fmt.Println(solve1(in, Range{200000000000000, 400000000000000}))
	fmt.Println(solve2(in))
}

func solve1(in string, area Range) int {
	stones := ParseStones(in)
	pairs := gotils.Pairs(stones)
	sum := 0

	for _, p := range pairs {
		i, ok := p[0].Intersection(p[1])

		if !ok {
			continue
		}

		if i.x >= area.left && i.x <= area.right &&
			i.y >= area.left && i.y <= area.right {
			sum += 1
		}
	}

	return sum
}

func solve2(in string) int {
	sum := 0

	return sum
}

type Range struct {
	left, right float64
}

type Pos3d struct {
	x, y, z float64
}

type Vel3d struct {
	x, y, z float64
}

type Stone struct {
	id  int
	pos Pos3d
	v   Vel3d
}

func (s Stone) Slope() float64 {
	return s.v.y / s.v.x
}

func (s Stone) Inter() float64 {
	return s.pos.y - s.Slope()*s.pos.x
}

func (s Stone) Intersection(o Stone) (Pos3d, bool) {
	if s.Slope()-o.Slope() == 0 {
		return Pos3d{}, false
	}

	x := (o.Inter() - s.Inter()) / (s.Slope() - o.Slope())
	y := s.Slope()*x + s.Inter()

	isInFutureS := gotils.Abs(x-(s.pos.x+s.v.x)) < gotils.Abs(x-s.pos.x)
	isInFutureO := gotils.Abs(x-(o.pos.x+o.v.x)) < gotils.Abs(x-o.pos.x)

	return Pos3d{x, y, -1}, isInFutureS && isInFutureO
}

func ParseStones(in string) []Stone {
	stones := make([]Stone, 0)

	for i, line := range strings.Split(in, "\n") {
		line = strings.ReplaceAll(line, " ", "")
		posStr, vStr, _ := strings.Cut(line, "@")
		pos := strings.Split(posStr, ",")
		v := strings.Split(vStr, ",")

		stone := Stone{}
		stone.id = i
		stone.pos.x = float64(gotils.MustParseInt(pos[0]))
		stone.pos.y = float64(gotils.MustParseInt(pos[1]))
		stone.pos.z = float64(gotils.MustParseInt(pos[2]))
		stone.v.x = float64(gotils.MustParseInt(v[0]))
		stone.v.y = float64(gotils.MustParseInt(v[1]))
		stone.v.z = float64(gotils.MustParseInt(v[2]))

		stones = append(stones, stone)
	}

	return stones
}
