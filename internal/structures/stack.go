package structures

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable](items []T) *Stack[T] {
	return &Stack[T]{items}
}

func (s *Stack[T]) Pop() T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}
