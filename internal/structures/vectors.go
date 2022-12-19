package structures

import "github.com/mathew/advent-of-code-2022/internal/maths"

type Vector struct {
	x int
	y int
}

func NewVector(x, y int) Vector {
	return Vector{x, y}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{
		v.x + v2.x,
		v.y + v2.y,
	}
}

func (v Vector) Subtract(v2 Vector) Vector {
	return Vector{
		v.x - v2.x,
		v.y - v2.y,
	}
}

func (v Vector) IsTouching(v2 Vector) bool {
	diff := v.Subtract(v2)

	return (diff.x == 0 || diff.x == 1 || diff.x == -1) &&
		(diff.y == 1 || diff.y == -1 || diff.y == 0)
}

func (v Vector) IsParallel(v2 Vector) bool {
	return v.x-v2.x == 0 || v.y-v2.y == 0
}

func (v Vector) MoveDirection(v2 Vector, velocity int) Vector {
	return v.Add(NewVector(maths.CapSignum(v2.x, 1), maths.CapSignum(v2.y, 1)))

}

func (v Vector) Coords() []int {
	return []int{v.x, v.y}
}

func VectorsToCoords(vs ...Vector) [][]int {
	nv := [][]int{}
	for _, v := range vs {
		nv = append(nv, v.Coords())
	}

	return nv
}
