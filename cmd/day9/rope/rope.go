package rope

import (
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type Rope struct {
	tailPositionsVisited map[structures.Vector]int
	headPosition         structures.Vector
	tailPosition         structures.Vector
}

func NewRope() Rope {
	return Rope{
		map[structures.Vector]int{structures.NewVector(0, 0): 1},
		structures.NewVector(0, 0),
		structures.NewVector(0, 0),
	}
}

func (r Rope) MoveRope(direction string, distance int) Rope {
	vectors := map[string]structures.Vector{
		"U": structures.NewVector(1, 0),
		"D": structures.NewVector(-1, 0),
		"R": structures.NewVector(0, 1),
		"L": structures.NewVector(0, -1),
	}
	vec := vectors[direction]

	for i := 0; i < distance; i++ {
		r.headPosition = r.headPosition.Add(vec)

		if r.headPosition.IsTouching(r.tailPosition) {
			continue
		}

		if !r.headPosition.IsTouching(r.tailPosition) &&
			r.headPosition.IsParallel(r.tailPosition) {

			r.tailPosition = r.tailPosition.Add(vec)
			r.tailPositionsVisited[r.tailPosition] += 1
		} else {
			diff := r.headPosition.Subtract(r.tailPosition)
			r.tailPosition = r.tailPosition.MoveDirection(diff, 1)
			r.tailPositionsVisited[r.tailPosition] += 1
		}
	}

	return Rope{r.tailPositionsVisited, r.headPosition, r.tailPosition}
}

func (r Rope) TailPositionsVisitedCount() int {
	return len(structures.MapValues(r.tailPositionsVisited))
}
