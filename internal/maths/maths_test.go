package maths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	assert.Equal(t, 30, Sum(20, 10))
	assert.Equal(t, 30, Sum(20, -10, -5, 15, 10))
	assert.Equal(t, 30, Sum(10, 10, 10))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 5, Max(1, 2, 3, 4, 5))
}
