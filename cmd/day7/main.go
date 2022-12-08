package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day7/directory"
	"github.com/mathew/advent-of-code-2022/internal/iter"
	"github.com/mathew/advent-of-code-2022/internal/maths"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func main() {
	r := maths.Sum(
		iter.Filter(
			func(i int) bool { return i < 100001 },
			structures.MapValues(directory.Evaluate(COMMANDS))...,
		)...,
	)

	log.Printf("Part one: %d", r)
}
