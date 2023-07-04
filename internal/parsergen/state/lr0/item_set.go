package lr1

import (
	"encoding/binary"
	"fmt"
	"sort"
	"strings"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/set"
)

type ItemSet struct {
	g     *grammar.AugmentedGrammar
	items map[Item]struct{}
}

func NewItemSet(g *grammar.AugmentedGrammar) *ItemSet {
	return &ItemSet{
		g:     g,
		items: make(map[Item]struct{}),
	}
}

func (b *ItemSet) Add(item Item) bool {
	if _, ok := b.items[item]; ok {
		return false
	}
	b.items[item] = struct{}{}
	return true
}

// Closure computes the closure of lr(0) items as outlined in
// the Dragon Book section 4.7 Fig. 4.33.
func (b *ItemSet) Closure() {
	changed := true
	for changed {
		changed = false
		// For each item [A -> α.Bβ]:
		for item := range b.items {
			prod := b.g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				continue
			}
			B, ok := b.g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Rule)
			if !ok {
				continue
			}
			// B is a rule; add all [B -> .γ] to the closure.
			for _, prod := range B.Prods {
				changed = b.Add(NewItem(b.g, prod, 0)) || changed
			}
		}
	}
}

func (s *ItemSet) FollowingSymbols() []grammar.Symbol {
	symSet := new(set.Set[grammar.Symbol])
	for item := range s.items {
		prod := s.g.Prods[item.Prod]
		if item.Dot >= uint32(len(prod.Terms)) {
			continue
		}
		symSet.Add(s.g.TermSymbol(prod.Terms[item.Dot]))
	}
	syms := symSet.Elements()

	// Symbol order determines state creation order.
	// Make the analysis deterministic by sorting.
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})

	return syms
}

func (b *ItemSet) ForEach(fn func(item Item)) {
	items := make([]Item, 0, len(b.items))
	for item := range b.items {
		items = append(items, item)
	}
	sortItems(items)
	for _, item := range items {
		fn(item)
	}
}

func (b *ItemSet) State() *State {
	items := make([]Item, 0, len(b.items))
	for item := range b.items {
		if item.IsKernel() {
			items = append(items, item)
		}
	}
	sortItems(items)

	itemKey := func(i Item) []byte {
		key := make([]byte, 0, binary.MaxVarintLen32*3)
		key = binary.AppendUvarint(key, uint64(i.Prod))
		key = binary.AppendUvarint(key, uint64(i.Dot))
		return key
	}

	keyLen := 0
	itemKeys := make([][]byte, len(items))
	for i, item := range items {
		itemKeys[i] = itemKey(item)
		keyLen += len(itemKeys[i])
	}

	key := make([]byte, 0, keyLen)
	for _, itemKey := range itemKeys {
		key = append(key, itemKey...)
	}

	return &State{
		Items: items,
		Key:   string(key),
	}
}

func (s *ItemSet) String() string {
	str := strings.Builder{}
	s.ForEach(func(item Item) {
		fmt.Fprintln(&str, item.ToString(s.g))
	})
	return str.String()
}
