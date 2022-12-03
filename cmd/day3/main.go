package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day3/knapsack"
)

func main() {
	score := 0
	for _, kp := range KNAPSACKS {
		score += knapsack.ItemPriority(knapsack.NewKnapsack(kp).Overlap())
	}

	log.Printf("Part One: %d", score)
}
