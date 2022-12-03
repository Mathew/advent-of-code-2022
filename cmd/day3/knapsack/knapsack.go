package knapsack

import (
	"unicode"

	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type Knapsack struct {
	CompartmentOne []string
	CompartmentTwo []string
}

func NewKnapsack(items []string) Knapsack {
	mid := len(items) / 2

	return Knapsack{
		CompartmentOne: items[:mid],
		CompartmentTwo: items[mid:],
	}
}

func (k Knapsack) Overlap() string {
	return structures.Intersection(k.CompartmentOne, k.CompartmentTwo)[0]
}

func ItemPriority(item string) int {
	ascii_value_upper := 32
	ascii_diff_lower := 65

	asciiValue := []rune(item)[0]
	if unicode.IsUpper(asciiValue) {
		// 26 letters in the alphabet, 6 spaces are used for other chars to offset lower + upper by 32
		return int(asciiValue) - ascii_value_upper + 1 - 7
	}

	return int(asciiValue) - ascii_value_upper - ascii_diff_lower + 1
}
