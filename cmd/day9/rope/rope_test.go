package rope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTailVisited(t *testing.T) {
	movements := []struct {
		direction string
		distance  int
	}{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}

	rpe := NewRope(2)
	for _, movement := range movements {
		rpe = rpe.MoveRope(movement.direction, movement.distance)
	}

	assert.Equal(t, 13, rpe.TailPositionsVisitedCount())
}

func TestRopeMovementParallel(t *testing.T) {
	rope := NewRope(2).MoveRope("U", 4).MoveRope("D", 3)
	assert.Equal(t, 4, rope.TailPositionsVisitedCount())

	rope = rope.MoveRope("L", 4).MoveRope("R", 3)
	assert.Equal(t, 7, rope.TailPositionsVisitedCount())
}

func TestRopeMovementDiagnoal(t *testing.T) {
	rope := NewRope(2).MoveRope("U", 1).MoveRope("R", 2).MoveRope("R", 1).MoveRope("L", 3)
	assert.Equal(t, 3, rope.TailPositionsVisitedCount())
}

func TestMultipleKnotsNoMovement(t *testing.T) {
	movements := []struct {
		direction string
		distance  int
	}{
		{"R", 4},
		{"U", 4},
		{"L", 3},
		{"D", 1},
		{"R", 4},
		{"D", 1},
		{"L", 5},
		{"R", 2},
	}

	rpe := NewRope(10)
	for _, movement := range movements {
		rpe = rpe.MoveRope(movement.direction, movement.distance)
	}

	assert.Equal(t, 1, rpe.TailPositionsVisitedCount())
}

func TestMultipleKnots(t *testing.T) {
	movements := []struct {
		direction string
		distance  int
	}{
		{"R", 5},
		{"U", 8},
		{"L", 8},
		{"D", 3},
		{"R", 17},
		{"D", 10},
		{"L", 25},
		{"U", 20},
	}

	rpe := NewRope(10)
	for _, movement := range movements {
		rpe = rpe.MoveRope(movement.direction, movement.distance)
	}

	assert.Equal(t, 36, rpe.TailPositionsVisitedCount())
}
