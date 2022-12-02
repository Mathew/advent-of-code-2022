package rps

import (
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

const Rock string = "Rock"
const Paper string = "Paper"
const Scissor string = "Scissor"

const win = "Z"
const lose = "X"
const draw = "Y"

var beats = map[string]string{
	Rock:    Scissor,
	Paper:   Rock,
	Scissor: Paper,
}

func scoreRPS(p1, p2 string) int {
	if p1 == p2 {
		return 3
	}

	if p2 == beats[p1] {
		return 6
	}

	return 0
}

func desiredResultChoice(elfChoice string, result string) string {
	if result == draw {
		return elfChoice
	}

	if result == win {
		return structures.MapInvert(beats)[elfChoice]
	}

	return beats[elfChoice]
}

type StrategyGuide struct {
	choices        [][]string
	choiceToObject map[string]string
	scoring        map[string]int
}

func NewStrategyGuide(
	choices [][]string,
	choiceToObject map[string]string,
	scoring map[string]int) StrategyGuide {

	return StrategyGuide{
		choices:        choices,
		choiceToObject: choiceToObject,
		scoring:        scoring,
	}
}

func (sg StrategyGuide) TotalScore() int {
	score := 0

	for _, hands := range sg.choices {
		p2, p1 := sg.choiceToObject[hands[0]], sg.choiceToObject[hands[1]]

		score += scoreRPS(p1, p2) + sg.scoring[p1]
	}

	return score
}

func (sg StrategyGuide) PredictionScore() int {
	score := 0
	for _, hands := range sg.choices {
		elfChoice := sg.choiceToObject[hands[0]]
		yourChoice := desiredResultChoice(elfChoice, hands[1])

		score += scoreRPS(yourChoice, elfChoice) + sg.scoring[yourChoice]
	}

	return score
}
