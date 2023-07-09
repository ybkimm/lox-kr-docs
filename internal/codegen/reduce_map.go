package codegen

import (
	"fmt"
	gotypes "go/types"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func (s *State) MapReduceActions() error {
	s.ReduceMap = make(map[*grammar.Prod]*ReduceMethod)
	s.ReduceTypes = make(map[*grammar.Rule]gotypes.Type)
	for _, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		if rule.Generated != grammar.NotGenerated {
			continue
		}
		reduceName := rule.Name + s.ProdLabels[prod]
		method := s.ReduceMethods[reduceName]
		if method == nil {
			fmt.Println("missing reduce method ", reduceName)
			continue
		}
		reduceType := method.ReturnType
		if existing := s.ReduceTypes[rule]; existing == nil {
			s.ReduceTypes[rule] = reduceType
		} else if existing != reduceType {
			return fmt.Errorf(
				"conflicting reduce types for %v: %v and %v",
				rule.Name, existing, reduceType)
		}
	}

	getReduceTypeForGeneratedRule := func(
		rule *grammar.Rule,
		prod *grammar.Prod,
	) gotypes.Type {
		switch rule.Generated {
		case grammar.NotGenerated,
			grammar.GeneratedSPrime:
			// S' is never reduced. Ignore.
			return nil
		case grammar.GeneratedZeroOrOne:
			// a = b c?
			//  =>
			// a = b a'
			// a' = c | e
			if prod != rule.Prods[0] {
				return nil
			}
			cSym := s.Grammar.TermSymbol(prod.Terms[0])
			if cRule, ok := cSym.(*grammar.Rule); ok {
				return s.ReduceTypes[cRule]
			} else {
				return s.Token.Type()
			}

		case grammar.GeneratedOneOrMore:
			// a = b c+
			//  =>
			// a = b a'
			// a' = a' c
			//    | c
			if prod != rule.Prods[1] {
				return nil
			}
			cSym := s.Grammar.TermSymbol(prod.Terms[0])
			cType := s.Token.Type()
			if cRule, ok := cSym.(*grammar.Rule); ok {
				cType = s.ReduceTypes[cRule]
			}
			return gotypes.NewSlice(cType)

		default:
			panic("unreachable")
		}
	}

	changed := true
	for changed {
		changed = false
		for _, prod := range s.Grammar.Prods {
			rule := s.Grammar.ProdRule(prod)
			typ := getReduceTypeForGeneratedRule(rule, prod)
			if typ == nil {
				continue
			}
			existing := s.ReduceTypes[rule]
			if existing != nil {
				if !gotypes.Identical(existing, typ) {
					panic("mismatched types")
				}
				continue
			}
			s.ReduceTypes[rule] = typ
			changed = true
		}
	}

	for _, rule := range s.Grammar.Rules {
		if rule.Generated == grammar.GeneratedSPrime {
			continue
		}
		ruleReduceType := s.ReduceTypes[rule]
		if ruleReduceType == nil {
			panic("unreachable")
		}
		fmt.Println(rule.Name, ruleReduceType)
	}

	return nil
}
