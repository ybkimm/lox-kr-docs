package lr1

import (
	"encoding/binary"
	"sort"
	"strings"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
	"github.com/dcaiafa/lox/internal/util/set"
)

type ItemSet struct {
	itemMap map[Item]struct{}
	items   []Item

	Index int
}

func NewItemSet() *ItemSet {
	return &ItemSet{
		itemMap: make(map[Item]struct{}),
	}
}

func (b *ItemSet) Add(item Item) bool {
	if _, ok := b.itemMap[item]; ok {
		return false
	}
	b.items = nil
	b.itemMap[item] = struct{}{}
	return true
}

// Closure computes the closure of the ItemSet.
// The algorithm is summarized thusly:
//
//	For each item [A -> α.Bβ, a]:
//	 If there is a B -> γ:
//	   For each x in FIRST(β):
//	     Add [B -> .γ, x]
func (b *ItemSet) Closure(g *grammar.AugmentedGrammar) {
	changed := true
	for changed {
		changed = false
		for item := range b.itemMap {
			prod := g.Prods[item.Prod]
			if item.Dot == uint32(len(prod.Terms)) {
				continue
			}
			B, ok := g.TermSymbol(prod.Terms[item.Dot]).(*grammar.Rule)
			if !ok {
				continue
			}
			beta := g.TermSymbols(prod.Terms[item.Dot+1:])
			a := g.Terminals[item.Lookahead]
			firstSet := g.First(append(beta, a))
			for _, prodB := range B.Prods {
				firstSet.ForEach(func(terminal *grammar.Terminal) {
					changed = b.Add(NewItem(g, prodB, 0, terminal)) || changed
				})
			}
		}
	}
}

func (s *ItemSet) Follow(g *grammar.AugmentedGrammar) []grammar.Symbol {
	symSet := new(set.Set[grammar.Symbol])
	for item := range s.itemMap {
		prod := g.Prods[item.Prod]
		if item.Dot >= uint32(len(prod.Terms)) {
			continue
		}
		symSet.Add(g.TermSymbol(prod.Terms[item.Dot]))
	}
	syms := symSet.Elements()

	// Symbol order determines state creation order.
	// Make the analysis deterministic by sorting.
	sort.Slice(syms, func(i, j int) bool {
		return syms[i].SymName() < syms[j].SymName()
	})

	return syms
}

func (from *ItemSet) Goto(g *grammar.AugmentedGrammar, sym grammar.Symbol) *ItemSet {
	toState := NewItemSet()
	for _, item := range from.GetItems() {
		prod := g.Prods[item.Prod]
		if item.Dot == uint32(len(prod.Terms)) {
			continue
		}
		term := g.TermSymbol(prod.Terms[item.Dot])
		if term != sym {
			continue
		}
		toItem := item
		toItem.Dot++
		toState.Add(toItem)
	}
	toState.Closure(g)
	return toState
}

func (b *ItemSet) GetItems() []Item {
	if b.items != nil {
		return b.items
	}
	b.items = make([]Item, 0, len(b.itemMap))
	for item := range b.itemMap {
		b.items = append(b.items, item)
	}
	sortItems(b.items)
	return b.items
}

func (b *ItemSet) Len() int {
	return len(b.itemMap)
}

func (s *ItemSet) ToString(g *grammar.AugmentedGrammar) string {
	var str strings.Builder
	for i, item := range s.GetItems() {
		if i != 0 {
			str.WriteString("\n")
		}
		str.WriteString(item.ToString(g))
	}
	return str.String()
}

func (b *ItemSet) KeyWithLookahead() string {
	return b.key(func(i Item) []byte {
		var keyArr [3 * binary.MaxVarintLen32]byte
		itemKey := keyArr[:0]
		itemKey = binary.AppendUvarint(itemKey, uint64(i.Prod))
		itemKey = binary.AppendUvarint(itemKey, uint64(i.Dot))
		itemKey = binary.AppendUvarint(itemKey, uint64(i.Lookahead))
		return itemKey
	})
}

func (b *ItemSet) KeyWithoutLookahead() string {
	return b.key(func(i Item) []byte {
		var keyArr [3 * binary.MaxVarintLen32]byte
		itemKey := keyArr[:0]
		itemKey = binary.AppendUvarint(itemKey, uint64(i.Prod))
		itemKey = binary.AppendUvarint(itemKey, uint64(i.Dot))
		return itemKey
	})
}

func (b *ItemSet) key(itemKeyFunc func(i Item) []byte) string {
	itemKeys := make([][]byte, b.Len())
	keyLen := 0
	for i, item := range b.GetItems() {
		itemKeys[i] = itemKeyFunc(item)
		keyLen += len(itemKeys[i])
	}

	key := make([]byte, 0, keyLen)
	for _, itemKey := range itemKeys {
		key = append(key, itemKey...)
	}

	return string(key)
}
