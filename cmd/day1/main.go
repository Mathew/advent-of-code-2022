package main

import (
	"log"
	"sort"

	"github.com/mathew/advent-of-code-2022/internal/maths"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func main() {
	elfToCalories := map[int]int{}

	for i, inventory := range ELVES_CALORIES {
		elfToCalories[i] = maths.Sum(inventory...)
	}

	calorieValues := structures.Values(elfToCalories)
	log.Printf("Part One: %v", maths.Max(calorieValues...))

	sort.Sort(sort.Reverse(sort.IntSlice(calorieValues)))
	log.Printf("Part Two: %v", maths.Sum(calorieValues[:3]...))
}
