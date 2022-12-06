package main

import (
	"log"

	"github.com/mathew/advent-of-code-2022/cmd/day6/signal"
)

func main() {
	log.Printf("Part one: %d", signal.NewSignal(SIGNAL).StartMarkerLocation(4))
	log.Printf("Part two: %d", signal.NewSignal(SIGNAL).StartMarkerLocation(14))

}
