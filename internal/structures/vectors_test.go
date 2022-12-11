package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVectorAdd(t *testing.T) {
	v := NewVector(5, 4)
	v2 := NewVector(1, -1)

	assert.Equal(t, NewVector(6, 3), v.Add(v2))
}

func TestVectorSubtract(t *testing.T) {
	v := NewVector(5, 4)
	v2 := NewVector(4, 5)

	assert.Equal(t, NewVector(1, -1), v.Subtract(v2))
}

func TestVectorTouching(t *testing.T) {
	tests := []struct {
		input    []Vector
		expected bool
	}{
		{
			[]Vector{NewVector(1, 1), NewVector(2, 1)}, true,
		},
		{
			[]Vector{NewVector(2, 1), NewVector(2, 1)}, true,
		},
		{
			[]Vector{NewVector(5, 1), NewVector(2, 1)}, false,
		},
		{
			[]Vector{NewVector(5, 5), NewVector(8, 8)}, false,
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.expected, tc.input[0].IsTouching(tc.input[1]))
	}
}

func TestVectorParallel(t *testing.T) {
	v, v2 := NewVector(1, 1), NewVector(3, 1)

	assert.True(t, v.IsParallel(v2))
}

func TestMoveDirection(t *testing.T) {
	v, v2 := NewVector(1, 1), NewVector(3, 4)

	assert.Equal(t, NewVector(2, 2), v.MoveDirection(v2, 1))
}
