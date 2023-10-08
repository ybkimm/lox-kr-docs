package lr2

import "github.com/dcaiafa/lox/internal/util/set"

// First returns the set of Terminals that could be derived first from a set of
// symbols composed of Rules or Terminals. The Dragon Book section 4.4 has a
// rigorous albeit hard-to-follow definition for the FIRST function. It is
// easier to understand with examples:
//
// Given:
//
//		A = B C '%' E | D | '+'
//		B = '-' | ε
//		C = '/' | ε
//		D = '*' | ε
//	  E = '$'
//
//	First(B) = { '-', ε }
//
// This should be intuitive.
//
//	First(B, '*') = { '-', '*' }
//
// Because First(B) includes ε, the result must also include First('*') but not
// ε since First('*') does not include ε.
//
//	First(A) = { '-', '/', '%', '*', '+', ε }
//
// Notice that First(B), and First(C) are included because both include ε, but
// First(E) is not included as First('%') does not include ε. '*' is included by
// First(D), and '+' by First('+'). Finally ε is in the final result only
// because First(D) includes it.
func First(g *Grammar, syms []int) set.Set[int] {
	visited := new(set.Set[int])
	if len(syms) == 1 {
		return first(g, visited, syms[0])
	}
	var firstSet set.Set[int]
	for _, sym := range syms {
		partialFirst := first(g, visited, sym)
		firstSet.AddSet(partialFirst)

		// If sym[i] includes ε, include FIRST(sym[i+1]) in FIRST(syms).
		// Otherwise, stop now.
		if !partialFirst.Has(Epsilon) {
			firstSet.Remove(Epsilon)
			break
		}
	}
	return firstSet
}

func first(g *Grammar, visited *set.Set[int], s int) set.Set[int] {
	if IsTerminal(s) {
		return set.New[int](s)
	}

	// Productions can contain recursion.
	// E.g.: xs = xs x | x
	if visited.Has(s) {
		return set.Set[int]{}
	}
	visited.Add(s)

	rule := g.GetRule(s)
	firstSet := set.Set[int]{}
	for _, prodIndex := range rule.Prods {
		prod := g.Prods[prodIndex]
		if len(prod.Terms) == 0 {
			firstSet.Add(Epsilon)
			continue
		}

		addEpsilon := true
		for _, term := range prod.Terms {
			termFirst := first(g, visited, term)
			hasEpsilon := false
			termFirst.ForEach(func(s int) {
				if s == Epsilon {
					hasEpsilon = true
					return
				}
				firstSet.Add(s)
			})
			if !hasEpsilon {
				addEpsilon = false
				break
			}
		}
		if addEpsilon {
			firstSet.Add(Epsilon)
		}
	}
	return firstSet
}
