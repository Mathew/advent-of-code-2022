package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day10/cpu"
	"github.com/mathew/advent-of-code-2022/internal/maths"
)

func main() {
	pu := cpu.CycleCounter([]int{20, 60, 100, 140, 180, 220}, cpu.NewCPU(getInput()))

	log.Printf("Day One: %d", maths.Sum(pu(), pu(), pu(), pu(), pu(), pu()))
	log.Printf("Day Two: ")
	log.Printf("\n %v", cpu.Monitor(cpu.NewCPU(getInput())))
}
