package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day8/trees"
)

func main() {
	log.Printf("Part one: %d", trees.NewForest(FOREST).VisibleTreesCount())
}
