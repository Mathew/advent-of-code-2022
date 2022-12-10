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

func calculateSightline(height int, trees []int) int {
	viewDistance := 0

	for _, tree := range trees {
		viewDistance += 1

		if tree >= height {
			return viewDistance
		}
	}

	return viewDistance
}

func (f Forest) BestTree() int {
	bestSightlines := 0
	for _, tree := range f.trees.ElementsRange(1, f.xSize-1, 1, f.ySize-1) {
		rowTrees := f.trees.GetRow(tree.X)
		columnTrees := f.trees.GetColumn(tree.Y)

		sightlines := calculateSightline(tree.Item, columnTrees[tree.X+1:]) *
			calculateSightline(tree.Item, structures.Reverse(columnTrees[:tree.X])) *
			calculateSightline(tree.Item, rowTrees[tree.Y+1:]) *
			calculateSightline(tree.Item, structures.Reverse(rowTrees[:tree.Y]))

		if sightlines > bestSightlines {
			bestSightlines = sightlines
		}

	}
	return bestSightlines
}
