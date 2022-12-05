package crates

import (
	"testing"

	"github.com/mathew/advent-of-code-2022/internal/structures"
	"github.com/stretchr/testify/assert"
)

func Crane9000Stub() CrateMover9000 {
	return NewCrateMover9000(map[int]*structures.Stack[string]{
		1: structures.NewStack([]string{"Z", "N"}),
		2: structures.NewStack([]string{"M", "C", "D"}),
		3: structures.NewStack([]string{"P"}),
	})
}

func Crane9001Stub() CrateMover9000 {
	return NewCrateMover9001(map[int]*structures.Stack[string]{
		1: structures.NewStack([]string{"Z", "N"}),
		2: structures.NewStack([]string{"M", "C", "D"}),
		3: structures.NewStack([]string{"P"}),
	})
}

func TestTopCrates(t *testing.T) {
	crane := Crane9000Stub()
	assert.Equal(t, []string{"N", "D", "P"}, crane.TopCrates())
}

func TestRunInstructions(t *testing.T) {
	crane := Crane9000Stub()
	crane.RunInstructions([][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	})

	assert.Equal(t, []string{"C", "M", "Z"}, crane.TopCrates())
}

func TestMoveCrates9001(t *testing.T) {
	crane := Crane9001Stub()
	crane.RunInstructions([][]int{{1, 2, 1}})

	assert.Equal(t, []string{"D", "C", "P"}, crane.TopCrates())

	crane.RunInstructions([][]int{{3, 1, 3}})
	assert.Equal(t, []string{"", "C", "D"}, crane.TopCrates())
}
