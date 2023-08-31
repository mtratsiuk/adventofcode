package main

import (
	"fmt"
	"strings"

	"github.com/mtratsiuk/adventofcode/aoc-2022/go/mishutils"
)

const mRock = "X"
const mPaper = "Y"
const mScissors = "Z"

const rock = "A"
const paper = "B"
const scissors = "C"

const win = 6
const draw = 3
const loose = 0

const mLose = "X"
const mDraw = "Y"
const mWin = "Z"

var rules = map[string]map[string]int{
	mRock: {
		rock:     draw,
		paper:    loose,
		scissors: win,
	},
	mPaper: {
		rock:     win,
		paper:    draw,
		scissors: loose,
	},
	mScissors: {
		rock:     loose,
		paper:    win,
		scissors: draw,
	},
}

var newRules = map[string]map[string]string{
	rock: {
		mLose: mScissors,
		mDraw: mRock,
		mWin:  mPaper,
	},
	paper: {
		mLose: mRock,
		mDraw: mPaper,
		mWin:  mScissors,
	},
	scissors: {
		mLose: mPaper,
		mDraw: mScissors,
		mWin:  mRock,
	},
}

var signScores = map[string]int{
	mRock:     1,
	mPaper:    2,
	mScissors: 3,
}

var outcomeScores = map[string]int{
	mLose: 0,
	mDraw: 3,
	mWin:  6,
}

func main() {
	in := mishutils.ReadInput("02_rock_paper_scissors")

	fmt.Println(solve_1(in))
	fmt.Println(solve_2(in))
}

func solve_1(input string) (int, error) {
	score := 0

	for _, line := range strings.Split(input, "\n") {
		actions := strings.Split(line, " ")
		them := actions[0]
		us := actions[1]

		score += rules[us][them]
		score += signScores[us]
	}

	return score, nil
}

func solve_2(input string) (int, error) {
	score := 0

	for _, line := range strings.Split(input, "\n") {
		actions := strings.Split(line, " ")
		them := actions[0]
		outcome := actions[1]
		us := newRules[them][outcome]

		score += outcomeScores[outcome]
		score += signScores[us]
	}

	return score, nil
}
