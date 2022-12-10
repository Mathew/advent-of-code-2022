package structures

type Element[T comparable] struct {
	X    int
	Y    int
	Item T
}

func NewElement[T comparable](x, y int, item T) Element[T] {
	return Element[T]{X: x, Y: y, Item: item}
}

type Matrix[T comparable] struct {
	data  [][]T
	tData [][]T
}

func NewMatrix[T comparable](data [][]T) Matrix[T] {
	return Matrix[T]{
		data,
		Transpose(data),
	}
}

func (m Matrix[T]) GetColumn(column int) []T {
	return m.tData[column]
}

func (m Matrix[T]) GetRow(row int) []T {
	return m.data[row]
}

func (m Matrix[T]) Elements() []Element[T] {
	elements := []Element[T]{}

	for x := 0; x < len(m.data); x++ {
		for y := 0; y < len(m.data[0]); y++ {
			elements = append(elements, NewElement(x, y, m.data[x][y]))
		}
	}

	return elements
}

func (m Matrix[T]) ElementsRange(xStart, xEnd, yStart, yEnd int) []Element[T] {
	elements := []Element[T]{}

	for x := xStart; x < xEnd; x++ {
		for y := yStart; y < yEnd; y++ {
			elements = append(elements, NewElement(x, y, m.data[x][y]))
		}
	}

	return elements

}
