package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day3/knapsack"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func main() {
	score := 0
	for _, kp := range KNAPSACKS {
		score += knapsack.ItemPriority(knapsack.NewKnapsack(kp).Overlap())
	}

	log.Printf("Part One: %d", score)

	groupScore := 0
	loops := len(KNAPSACKS) / 3

	for i := 0; i < loops; i++ {
		groupScore += knapsack.ItemPriority(
			structures.Intersection(
				KNAPSACKS[i*3+2],
				structures.Intersection(KNAPSACKS[i*3], KNAPSACKS[i*3+1]),
			)[0],
		)
	}

	log.Printf("Part Two: %d", groupScore)
}
