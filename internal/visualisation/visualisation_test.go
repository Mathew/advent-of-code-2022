package visualisation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleCoords(t *testing.T) {

	exp := `============ Visual: test ============
.......
.....B.
.......
.......
.......
.A.....
.......
============  End: test   ============
`

	result := VisualMatrix("test", [][]int{{1, 1}, {5, 5}}, []string{"A", "B"})

	assert.Equal(t, exp, result)
}

func TestNegativeCoords(t *testing.T) {
	exp := `============ Visual: test ============
.......
.....B.
.......
.......
.......
.......
.......
.A.....
============  End: test   ============
`
	result := VisualMatrix("test", [][]int{{-1, 1}, {5, 5}}, []string{"A", "B"})

	assert.Equal(t, exp, result)
}
