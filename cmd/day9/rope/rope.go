package rope

import (
	"github.com/mathew/advent-of-code-2022/internal/structures"
)

type Rope struct {
	tailPositionsVisited map[structures.Vector]int
	knotPositions        []structures.Vector
}

func NewRope(numKnots int) Rope {
	knots := make([]structures.Vector, numKnots)

	return Rope{
		map[structures.Vector]int{structures.NewVector(0, 0): 1},
		knots,
	}
}

func NewRopeFromState(knots [][]int) Rope {
	vs := []structures.Vector{}
	for _, k := range knots {
		vs = append(vs, structures.NewVector(k[0], k[1]))
	}

	return Rope{
		map[structures.Vector]int{},
		vs,
	}
}

func (r Rope) moveKnot(prevKnot, knot, move structures.Vector) structures.Vector {
	if prevKnot.IsTouching(knot) {
		return knot
	}
	diff := prevKnot.Subtract(knot)
	knot = knot.MoveDirection(diff, 1)

	return knot
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
		r.knotPositions[0] = r.knotPositions[0].Add(vec)
		prevKnot := r.knotPositions[0]

		for kPos := 1; kPos < len(r.knotPositions); kPos++ {
			knot := r.knotPositions[kPos]
			newCalc := r.moveKnot(prevKnot, knot, vec)

			if newCalc != knot && kPos == len(r.knotPositions)-1 {
				r.tailPositionsVisited[newCalc] += 1
			}

			r.knotPositions[kPos] = newCalc
			prevKnot = newCalc
		}
	}

	return Rope{r.tailPositionsVisited, r.knotPositions}
}

func visualSnake() {}

func (r Rope) TailPositionsVisitedCount() int {
	return len(structures.MapValues(r.tailPositionsVisited))
}
