package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day8/trees"
)

func main() {
	ts := trees.NewForest(FOREST)
	log.Printf("Part one: %d", ts.VisibleTreesCount())
	log.Printf("Part Two: %d", ts.BestTree())
}
