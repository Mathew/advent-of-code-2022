package structures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixElements(t *testing.T) {
	m := NewMatrix([][]string{{"a", "b", "c"}, {"d", "e", "f"}})

	assert.Equal(t, NewElement(0, 0, "a"), m.Elements()[0])
	assert.Equal(t, NewElement(0, 1, "b"), m.Elements()[1])
}

func TestMatrixElementsRange(t *testing.T) {
	m := NewMatrix([][]string{{"a", "b", "c"}, {"d", "e", "f"}})

	assert.Equal(t, 4, len(m.ElementsRange(0, 2, 0, 2)))
}
