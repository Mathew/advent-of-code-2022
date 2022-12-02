package rps

const Rock string = "ROCK"
const Paper string = "Paper"
const Scissor string = "Scissor"

func scoreRPS(p1, p2 string) int {
	if p1 == p2 {
		return 3
	}

	beats := map[string]string{
		Rock:    Scissor,
		Paper:   Rock,
		Scissor: Paper,
	}

	if p2 == beats[p1] {
		return 6
	}

	return 0
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
