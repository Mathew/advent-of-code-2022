package crates

import (
	"testing"

	"github.com/mathew/advent-of-code-2022/internal/structures"
	"github.com/stretchr/testify/assert"
)

func CraneStub() Crane {
	return NewCrane(map[int]*structures.Stack[string]{
		1: structures.NewStack([]string{"Z", "N"}),
		2: structures.NewStack([]string{"M", "C", "D"}),
		3: structures.NewStack([]string{"P"}),
	})
}

func TestTopCrates(t *testing.T) {
	crane := CraneStub()
	assert.Equal(t, []string{"N", "D", "P"}, crane.TopCrates())
}

func TestMoveCrates(t *testing.T) {
	crane := CraneStub()
	crane.MoveCrates(1, 2, 1)

	assert.Equal(t, "D", crane.crates[1].Pop())
}

func TestRunInstructions(t *testing.T) {
	crane := CraneStub()
	crane.MoveCrates(1, 2, 1)
	crane.MoveCrates(3, 1, 3)
	crane.MoveCrates(2, 2, 1)
	crane.MoveCrates(1, 1, 2)

	assert.Equal(t, []string{"C", "M", "Z"}, crane.TopCrates())
}
