package visualisation

import (
	"fmt"

	"github.com/mathew/advent-of-code-2022/internal/maths"
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

func VisualMatrix(name string, coords [][]int, names []string) string {
	xs := []int{}
	ys := []int{}
	for _, cs := range coords {
		xs = append(xs, cs[1])
		ys = append(ys, cs[0])
	}

	xMaxSize, yMaxSize := maths.Max(xs...), maths.Max(ys...)
	xMinSize, yMinSize := maths.Min(append(xs, 0)...), maths.Min(append(ys, 0)...)

	adjustedCoords := [][]int{}
	for _, c := range coords {
		adjustedCoords = append(adjustedCoords, []int{c[0] - yMinSize, c[1] - xMinSize})
	}

	xLength := xMaxSize + 2 - xMinSize
	yLength := yMaxSize + 2 - yMinSize

	grid := make([][]string, yLength)
	for y := 0; y < yLength; y++ {
		grid[y] = make([]string, xLength)
	}

	for i, c := range adjustedCoords {
		grid[c[0]][c[1]] = names[i]
	}

	s := fmt.Sprintf("============ Visual: %s ============\n", name)
	for _, row := range structures.Reverse(grid) {
		for _, col := range row {
			if col == "" {
				col = "."
			}
			s += fmt.Sprintf("%v", col)
		}

		s += fmt.Sprintf("\n")
	}
	s += fmt.Sprintf("============  End: %s   ============\n", name)

	return s
}
