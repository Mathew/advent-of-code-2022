package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPop(t *testing.T) {
	var s = Stack[string]{[]string{"a", "b", "c"}}
	assert.Equal(t, "c", s.Pop())
}

func TestStackPush(t *testing.T) {
	var s = Stack[string]{[]string{"a", "b", "c"}}
	s.Push("E")
	assert.Equal(t, []string{"a", "b", "c", "E"}, s.items)
}
