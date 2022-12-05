package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day5/crates"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func convertStacks(map[int][]string) map[int]*structures.Stack[string] {
	stacks := map[int]*structures.Stack[string]{}

	for no, crates := range STACKS() {
		stacks[no] = structures.NewStack(crates)
	}

	return stacks
}

func main() {

	crane := crates.NewCrateMover9000(convertStacks(STACKS()))
	crane.RunInstructions(MOVES)
	log.Printf("Part one: %v", crane.TopCrates())

	craneTwo := crates.NewCrateMover9001(convertStacks(STACKS()))
	craneTwo.RunInstructions(MOVES)
	log.Printf("Part Two: %v", craneTwo.TopCrates())
}
