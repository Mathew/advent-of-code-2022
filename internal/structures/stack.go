package structures

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable](items []T) *Stack[T] {
	c := make([]T, len(items))
	copy(c, items)
	return &Stack[T]{c}
}

func (s *Stack[T]) Pop() T {
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Glimpse() T {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}

	var result T
	return result
}
