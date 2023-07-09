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
		fmt.Println(reduceName)
	}
	return nil
}
