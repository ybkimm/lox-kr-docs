package array

type Array[T any] struct {
	elems []T
}

func (s *Array[T]) SetCapacity(c int) {
	if cap(s.elems) < c {
		newElems := make([]T, len(s.elems), c)
		copy(newElems, s.elems)
		s.elems = newElems
	}
}

func (s *Array[T]) Elements() []T {
	if s == nil {
		return nil
	}
	return s.elems
}

func (s *Array[T]) Len() int {
	return len(s.elems)
}

func (s *Array[T]) Empty() bool {
	return len(s.elems) == 0
}

func (s *Array[T]) Add(e T) {
	s.elems = append(s.elems, e)
}

func (s *Array[T]) Get(n int) T {
	return s.elems[n]
}
