package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("04_scratchcards")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(input string) int {
	sum := 0

	for _, line := range strings.Split(input, "\n") {
		card := ParseCard(line)
		won := card.WonCount()

		if won == 0 {
			continue
		}

		points := 1

		for i := won - 1; i > 0; i -= 1 {
			points *= 2
		}

		sum += points
	}

	return sum
}

func solve2(input string) int {
	sum := 0

	cards := make([]CardsPile, 0)

	for _, line := range strings.Split(input, "\n") {
		cards = append(cards, CardsPile{ParseCard(line)})
	}

	for i, copies := range cards {
		wc := copies[0].WonCount()

		for range copies {
			for j, won := i+1, wc; j < len(cards) && won > 0; j, won = j+1, won-1 {
				cards[j].AddCopy()
			}
		}

		sum += len(copies)
	}

	return sum
}

type Card struct {
	num     int
	winning []int
	actual  []int
}

type CardsPile []Card

func (cp *CardsPile) AddCopy() {
	*cp = append(*cp, (*cp)[0])
}

func (card *Card) WonCount() int {
	won := 0

	for _, a := range card.actual {
		for _, w := range card.winning {
			if a == w {
				won += 1
				break
			}
		}
	}

	return won
}

func ParseCard(line string) Card {
	card := Card{}

	cardStr, rest, _ := strings.Cut(line, ": ")
	cardNum := strings.Fields(cardStr)[1]

	card.num = gotils.MustParseInt(cardNum)

	winningStr, actualStr, _ := strings.Cut(rest, " | ")

	for _, w := range strings.Fields(winningStr) {
		card.winning = append(card.winning, gotils.MustParseInt(w))
	}

	for _, a := range strings.Fields(actualStr) {
		card.actual = append(card.actual, gotils.MustParseInt(a))
	}

	return card
}
