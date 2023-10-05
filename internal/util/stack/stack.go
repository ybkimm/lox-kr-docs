package stack

type Stack[T any] struct {
	elems []T
}

func (s *Stack[T]) SetCapacity(c int) {
	if cap(s.elems) < c {
		newElems := make([]T, len(s.elems), c)
		copy(newElems, s.elems)
		s.elems = newElems
	}
}

func (s *Stack[T]) Elements() []T {
	if s == nil {
		return nil
	}
	return s.elems
}

func (s *Stack[T]) Len() int {
	return len(s.elems)
}

func (s *Stack[T]) Empty() bool {
	return len(s.elems) == 0
}

func (s *Stack[T]) Push(e T) {
	s.elems = append(s.elems, e)
}

func (s *Stack[T]) Pop() T {
	l := len(s.elems)
	elem := s.elems[l-1]
	s.elems = s.elems[:l-1]
	return elem
}

func (s *Stack[T]) Peek() T {
	return s.elems[len(s.elems)-1]
}
