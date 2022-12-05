package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day5/crates"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func main() {
	stacks := map[int]*structures.Stack[string]{}
	for no, crates := range STACKS {
		stacks[no] = structures.NewStack(crates)
	}

	crane := crates.NewCrane(stacks)
	crane.RunInstructions(MOVES)
	log.Printf("Part one: %v", crane.TopCrates())
}
