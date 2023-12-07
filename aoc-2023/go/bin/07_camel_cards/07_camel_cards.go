package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/mtratsiuk/adventofcode/gotils"
)

func main() {
	in := gotils.ReadInput("07_camel_cards")

	fmt.Println(solve1(in))
	fmt.Println(solve2(in))
}

func solve1(in string) int {
	sum := 0
	hands := make([]Hand, 0)

	for _, line := range strings.Split(in, "\n") {
		hands = append(hands, ParseHand(line, false))
	}

	slices.SortFunc(hands, HandSorter(false))

	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

func solve2(in string) int {
	sum := 0
	hands := make([]Hand, 0)

	for _, line := range strings.Split(in, "\n") {
		hands = append(hands, ParseHand(line, true))
	}

	slices.SortFunc(hands, HandSorter(true))

	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

type HandType int

const (
	HandTypeHighCard = iota
	HandTypeOnePair
	HandTypeTwoPair
	HandTypeThreeOfAKind
	HandTypeFullHouse
	HandTypeFourOfAKind
	HandTypeFiveOfAKind
)

const Joker = rune('J')

var CardStrengthMap = map[rune]int{
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
	'T': 10,
	'J': 11,
	'Q': 12,
	'K': 13,
	'A': 14,
}

func CardStrength(c rune, joker bool) int {
	if joker && c == Joker {
		return 1
	}

	return CardStrengthMap[c]
}

type HandGroups map[rune]int

type Hand struct {
	repr    string
	bid     int
	joker   bool
	grouped HandGroups
	types   []HandType
}

func HandSorter(joker bool) func(Hand, Hand) int {
	return func(lh, rh Hand) int {
		bestL := gotils.Max(lh.types)
		bestR := gotils.Max(rh.types)

		if bestL != bestR {
			return int(bestL - bestR)
		}

		for i, lc := range lh.repr {
			rc := rune(rh.repr[i])

			ls := CardStrength(lc, joker)
			rs := CardStrength(rc, joker)

			if ls != rs {
				return ls - rs
			}
		}

		return 0
	}
}

func ParseHand(line string, joker bool) Hand {
	hand := Hand{}
	hand.joker = joker
	hand.grouped = make(HandGroups, 0)
	hand.types = make([]HandType, 0)

	repr, bidStr, _ := strings.Cut(line, " ")
	hand.repr = repr
	hand.bid = gotils.MustParseInt(bidStr)

	for _, c := range repr {
		hand.grouped[c] += 1
	}

	// high card
	hand.types = append(hand.types, HandTypeHighCard)

	// one pair
	if hand.CountOf(2) >= 1 ||
		hand.Jokers() >= 1 {
		hand.types = append(hand.types, HandTypeOnePair)
	}

	// two pair
	if hand.CountOf(2) >= 2 ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 1) ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 1) {
		hand.types = append(hand.types, HandTypeTwoPair)
	}

	// three of a kind
	if hand.CountOf(3) >= 1 ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 1) ||
		hand.Jokers() >= 2 {
		hand.types = append(hand.types, HandTypeThreeOfAKind)
	}

	// full house
	if hand.CountOf(2) >= 1 && hand.CountOf(3) >= 1 ||
		(hand.CountOf(2) >= 2 && hand.Jokers() >= 1) ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 2) {
		hand.types = append(hand.types, HandTypeFullHouse)
	}

	// four of a kind
	if hand.CountOf(4) >= 1 ||
		(hand.CountOf(3) >= 1 && hand.Jokers() >= 1) ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 2) ||
		hand.Jokers() >= 3 {
		hand.types = append(hand.types, HandTypeFourOfAKind)
	}

	// five of a kind
	if hand.CountOf(5) >= 1 ||
		(hand.CountOf(4) >= 1 && hand.Jokers() >= 1) ||
		(hand.CountOf(3) >= 1 && hand.Jokers() >= 2) ||
		(hand.CountOf(2) >= 1 && hand.Jokers() >= 3) ||
		(hand.Jokers() >= 4) {
		hand.types = append(hand.types, HandTypeFiveOfAKind)
	}

	return hand
}

func (h Hand) CountOf(c int) int {
	matches := 0

	for k, count := range h.grouped {
		if h.joker && k == Joker {
			continue
		}

		if count == c {
			matches += 1
		}
	}

	return matches
}

func (h Hand) Jokers() int {
	if !h.joker {
		return 0
	}

	return h.grouped[Joker]
}
