package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day9/rope"
)

func main() {
	rpe := rope.NewRope()
	for _, m := range MOVEMENTS {
		rpe = rpe.MoveRope(m.direction, m.distance)
	}

	log.Printf("Day one: %d", rpe.TailPositionsVisitedCount())
}
