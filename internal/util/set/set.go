package set

import (
	"github.com/dcaiafa/lox/internal/util/stablemap"
)

type Set[T comparable] struct {
	set stablemap.Map[T, bool]
}

func New[T comparable](xs ...T) Set[T] {
	var set Set[T]
	set.AddSlice(xs)
	return set
}

func (s *Set[T]) Clear() {
	s.set.Clear()
}

func (s *Set[T]) Add(x T) bool {
	if s.set.Has(x) {
		return false
	}
	s.set.Put(x, true)
	return true
}

func (s *Set[T]) AddSlice(xs []T) bool {
	changed := false
	for _, x := range xs {
		changed = s.Add(x) || changed
	}
	return changed
}

func (s *Set[T]) AddSet(o Set[T]) bool {
	changed := false
	o.set.ForEach(func(k T, v bool) {
		changed = s.Add(k) || changed
	})
	return changed
}

func (s *Set[T]) Remove(x T) {
	s.set.Remove(x)
}

func (s *Set[T]) Has(v T) bool {
	return s.set.Has(v)
}

func (s *Set[T]) Empty() bool {
	return s.set.Len() == 0
}

func (s *Set[T]) Len() int {
	return s.set.Len()
}

func (s *Set[T]) Elements() []T {
	return s.set.Keys()
}

func (s *Set[T]) Equal(o Set[T]) bool {
	if s.Len() != o.Len() {
		return false
	}
	isEqual := true
	s.ForEach(func(x T) {
		isEqual = isEqual && o.Has(x)
	})
	return true
}

func (s *Set[T]) ForEach(fn func(e T)) {
	s.set.ForEach(func(k T, _ bool) {
		fn(k)
	})
}

func (s *Set[T]) Clone() Set[T] {
	var c Set[T]
	s.ForEach(func(e T) {
		c.Add(e)
	})
	return c
}

func Sub[T comparable](a, b *Set[T]) *Set[T] {
	var c Set[T]
	a.ForEach(func(e T) {
		if !b.Has(e) {
			c.Add(e)
		}
	})
	return &c
}
