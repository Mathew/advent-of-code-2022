package trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVisibileTrees(t *testing.T) {
	trees := [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}

	assert.Equal(t, 21, NewForest(trees).VisibleTreesCount())
}
