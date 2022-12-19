package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day9/rope"
)

func main() {
	rpe := rope.NewRope(2)
	for _, m := range MOVEMENTS {
		rpe = rpe.MoveRope(m.direction, m.distance)
	}

	log.Printf("Day one: %d", rpe.TailPositionsVisitedCount())

	rpeTen := rope.NewRope(10)
	for _, m := range MOVEMENTS {
		rpe = rpeTen.MoveRope(m.direction, m.distance)
	}
	log.Printf("Day two: %d", rpeTen.TailPositionsVisitedCount())
}
