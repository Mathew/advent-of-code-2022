package knapsack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKnapsack(t *testing.T) {
	assert.Equal(t,
		Knapsack{[]string{"a", "b", "c"}, []string{"A", "B", "C"}},
		NewKnapsack([]string{"a", "b", "c", "A", "B", "C"}),
	)
}

func TestItemPriority(t *testing.T) {
	assert.Equal(t, 1, ItemPriority("a"))
	assert.Equal(t, 26, ItemPriority("z"))
	assert.Equal(t, 27, ItemPriority("A"))
	assert.Equal(t, 52, ItemPriority("Z"))
}
