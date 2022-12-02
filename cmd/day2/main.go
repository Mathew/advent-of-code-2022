package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day2/rps"
)

func main() {
	sg := rps.NewStrategyGuide(
		Instructions,
		map[string]string{
			"A": rps.Rock,
			"B": rps.Paper,
			"C": rps.Scissor,
			"X": rps.Rock,
			"Y": rps.Paper,
			"Z": rps.Scissor,
		}, map[string]int{
			rps.Rock:    1,
			rps.Paper:   2,
			rps.Scissor: 3,
		})

	log.Printf("Part One: %d", sg.TotalScore())
}
