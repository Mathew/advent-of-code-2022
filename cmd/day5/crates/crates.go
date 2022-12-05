package crates

import (
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type CrateMover9000 struct {
	crates     map[int]*structures.Stack[string]
	MoveCrates func(*CrateMover9000, int, int, int)
}

func NewCrateMover9000(stacks map[int]*structures.Stack[string]) CrateMover9000 {
	return CrateMover9000{stacks, MoveCratesStack}
}

func NewCrateMover9001(stacks map[int]*structures.Stack[string]) CrateMover9000 {
	return CrateMover9000{stacks, MoveCratesInPlace}
}

func (c *CrateMover9000) RunInstructions(instructions [][]int) {
	for _, instruction := range instructions {
		number, fromStack, toStack := instruction[0], instruction[1], instruction[2]
		c.MoveCrates(c, number, fromStack, toStack)
	}
}

func (c *CrateMover9000) TopCrates() []string {
	crates := []string{}

	for i := 1; i != len(c.crates)+1; i++ {
		crates = append(crates, c.crates[i].Glimpse())
	}

	return crates
}

func MoveCratesStack(crane *CrateMover9000, number, fromStack, toStack int) {
	for i := 0; i < number; i++ {
		item := crane.crates[fromStack].Pop()
		crane.crates[toStack].Push(item)
	}
}

func MoveCratesInPlace(crane *CrateMover9000, number, fromStack, toStack int) {
	items := structures.NewStack([]string{})

	for i := 0; i < number; i++ {
		items.Push(crane.crates[fromStack].Pop())
	}

	for ok := true; ok != items.IsEmpty(); {
		crane.crates[toStack].Push(items.Pop())
	}
}
