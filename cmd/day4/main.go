package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func main() {
	fullIntersections := 0
	partialIntersections := 0
	for _, pairs := range ASSIGNMENTS {

		assignmentOne := structures.Sequence(pairs[0][0], pairs[0][1])
		assignmentTwo := structures.Sequence(pairs[1][0], pairs[1][1])
		intersect := structures.Intersection(assignmentOne, assignmentTwo)

		if len(intersect) == len(assignmentOne) ||
			len(intersect) == len(assignmentTwo) {
			fullIntersections += 1
		}

		if len(intersect) > 0 {
			partialIntersections += 1
		}
	}

	log.Printf("Part One: %d", fullIntersections)
	log.Printf("Part Two: %d", partialIntersections)
}
