package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("02_cube_conundrum")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

const MaxRed = 12
const MaxGreen = 13
const MaxBlue = 14

func solve1(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		possible := true

		for _, color := range game.colors {
			if color.r > MaxRed ||
				color.g > MaxGreen ||
				color.b > MaxBlue {
				possible = false
			}
		}

		if possible {
			sum += game.id
		}
	}

	return sum
}

func solve2(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		game := parseGame(line)
		m := Rgb{}

		for _, color := range game.colors {
			m.r = max(m.r, color.r)
			m.g = max(m.g, color.g)
			m.b = max(m.b, color.b)
		}

		sum += m.r * m.g * m.b
	}

	return sum
}

type Rgb struct {
	r int
	g int
	b int
}

type Game struct {
	id     int
	colors []Rgb
}

func parseGame(line string) Game {
	game := Game{}
	game.colors = make([]Rgb, 0)

	gameStr, sets, _ := strings.Cut(line, ": ")
	_, gameId, _ := strings.Cut(gameStr, " ")

	game.id = gotils.MustParseInt(gameId)

	for _, set := range strings.Split(sets, "; ") {
		rgb := Rgb{}

		for _, cube := range strings.Split(set, ", ") {
			nStr, color, _ := strings.Cut(cube, " ")
			n := gotils.MustParseInt(nStr)

			switch color {
			case "red":
				rgb.r = n
			case "green":
				rgb.g = n
			case "blue":
				rgb.b = n
			}

			game.colors = append(game.colors, rgb)
		}

	}

	return game
}
