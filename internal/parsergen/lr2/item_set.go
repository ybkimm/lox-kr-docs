package lr2

import (
	"encoding/binary"
	"strings"

	"github.com/dcaiafa/lox/internal/util/set"
)

type ItemSet struct {
	set         set.Set[Item]
	cachedItems []Item
}

func (s *ItemSet) Add(item Item) bool {
	if s.set.Has(item) {
		return false
	}
	s.set.Add(item)
	s.cachedItems = nil
	return true
}

func (s *ItemSet) AddSet(o *ItemSet) bool {
	changed := false
	o.set.ForEach(func(i Item) {
		if s.set.Has(i) {
			return
		}
		changed = true
		s.set.Add(i)
	})
	return changed
}

func (s *ItemSet) ForEach(f func(i Item)) {
	s.set.ForEach(f)
}

func (s *ItemSet) Empty() bool {
	return s.set.Empty()
}

func (s *ItemSet) Clear() {
	s.set.Clear()
	s.cachedItems = nil
}

func (s *ItemSet) Items() []Item {
	if s.cachedItems == nil {
		s.cachedItems = s.set.Elements()
		SortItems(s.cachedItems)
	}
	return s.cachedItems
}

func (s *ItemSet) ToString(g *Grammar) string {
	var str strings.Builder
	for i, item := range s.Items() {
		if i != 0 {
			str.WriteRune('\n')
		}
		str.WriteString(item.ToString(g))
	}
	return str.String()
}

func (s *ItemSet) LR0Key() string {
	type lr0Item struct {
		Prod uint32
		Dot  uint32
	}
	var seen set.Set[lr0Item]
	key := make([]byte, 0, s.set.Len())
	for _, i := range s.Items() {
		if !i.IsKernel() {
			continue
		}
		lr0Key := lr0Item{Prod: uint32(i.Prod), Dot: uint32(i.Dot)}
		if seen.Has(lr0Key) {
			continue
		}
		seen.Add(lr0Key)
		key = binary.BigEndian.AppendUint32(key, lr0Key.Prod)
		key = binary.BigEndian.AppendUint32(key, lr0Key.Dot)
	}
	return string(key)
}
