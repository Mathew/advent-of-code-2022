package trees

import (
	"github.com/mathew/advent-of-code-2022/internal/iter"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type Forest struct {
	trees structures.Matrix[int]
	xSize int
	ySize int
}

func NewForest(trees [][]int) Forest {
	return Forest{structures.NewMatrix(trees), len(trees[0]), len(trees)}
}

func (f Forest) isTreeVisible(colPos, rowPos, height int) bool {
	rowTrees := f.trees.GetRow(rowPos)
	columnTrees := f.trees.GetColumn(colPos)

	return !iter.Any(func(i int) bool { return height <= i }, columnTrees[rowPos+1:]...) ||
		!iter.Any(func(i int) bool { return height <= i }, columnTrees[:rowPos]...) ||
		!iter.Any(func(i int) bool { return height <= i }, rowTrees[:colPos]...) ||
		!iter.Any(func(i int) bool { return height <= i }, rowTrees[colPos+1:]...)
}

func (f Forest) VisibleTreesCount() int {
	visibleTrees := 0
	for _, tree := range f.trees.ElementsRange(1, f.xSize-1, 1, f.ySize-1) {

		if f.isTreeVisible(tree.Y, tree.X, tree.Item) {
			visibleTrees += 1
		}
	}

	return visibleTrees + ((f.xSize + f.ySize - 2) * 2)
}
