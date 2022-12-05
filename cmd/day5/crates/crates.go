package crates

import "github.com/mathew/advent-of-code-2022/internal/structures"

type Crane struct {
	crates map[int]*structures.Stack[string]
}

func NewCrane(stacks map[int]*structures.Stack[string]) Crane {
	return Crane{stacks}
}

func (c *Crane) MoveCrates(number, fromStack, toStack int) {
	for i := 0; i < number; i++ {
		item := c.crates[fromStack].Pop()
		c.crates[toStack].Push(item)
	}
}

func (c *Crane) RunInstructions(instructions [][]int) {
	for _, instruction := range instructions {
		number, fromStack, toStack := instruction[0], instruction[1], instruction[2]
		c.MoveCrates(number, fromStack, toStack)
	}
}

func (c *Crane) TopCrates() []string {
	crates := []string{}

	for i := 1; i != len(c.crates)+1; i++ {
		crates = append(crates, c.crates[i].Pop())
	}

	return crates
}
