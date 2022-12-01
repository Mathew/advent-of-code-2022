package main

import (
	"log"
	"sort"

	"github.com/mathew/advent-of-code-2022/internal/maths"
)

func main() {
	elfToCalories := map[int]int{}

	for i, inventory := range ELVES_CALORIES {
		elfToCalories[i] = maths.Sum(inventory...)
	}

	calorieValues := []int{}
	for _, c := range elfToCalories {
		calorieValues = append(calorieValues, c)
	}

	log.Printf("Part One: %v", maths.Max(calorieValues...))

	sort.Slice(calorieValues, func(i, j int) bool {
		return calorieValues[i] > calorieValues[j]
	})

	log.Printf("Part Two: %v", maths.Sum(calorieValues[:3]...))
}
