package set

type Set[T comparable] struct {
	set   map[T]bool
	keyFn func(x T) string
}

func (s *Set[T]) Add(x T) bool {
	if s.set == nil {
		s.set = make(map[T]bool)
	}
	if s.set[x] {
		return false
	}
	s.set[x] = true
	return true
}

func (s *Set[T]) AddSlice(xs []T) bool {
	changed := false
	for _, x := range xs {
		changed = s.Add(x) || changed
	}
	return changed
}

func (s *Set[T]) AddSet(o *Set[T]) bool {
	changed := false
	for x := range o.set {
		changed = s.Add(x) || changed
	}
	return changed
}

func (s *Set[T]) Remove(x T) {
	delete(s.set, x)
}

func (s *Set[T]) Has(v T) bool {
	if s.set == nil {
		return false
	}
	return s.set[v]
}

func (s *Set[T]) Elements() []T {
	r := make([]T, 0, len(s.set))
	for x := range s.set {
		r = append(r, x)
	}
	return r
}

func (s *Set[T]) Equal(o *Set[T]) bool {
	if len(s.set) != len(o.set) {
		return false
	}
	for e := range s.set {
		if !o.set[e] {
			return false
		}
	}
	return true
}

func (s *Set[T]) ForEach(fn func(e T)) {
	for e := range s.set {
		fn(e)
	}
}

func (s *Set[T]) Clone() *Set[T] {
	c := new(Set[T])
	c.set = make(map[T]bool, len(s.set))
	for e := range s.set {
		c.set[e] = true
	}
	return c
}
