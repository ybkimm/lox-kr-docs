package codegen

import (
	"fmt"
	gotypes "go/types"

	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func (s *State) MapReduceActions() error {
	s.ReduceMap = make(map[*grammar.Prod]*ReduceMethod)
	s.ReduceTypes = make(map[*grammar.Rule]gotypes.Type)

	// Determine the Go type of the reduce-artifact of each rule.
	// Non-generated (user-specified) rules first.
	for ruleName, methods := range s.ReduceMethods {
		var reduceMethod *ReduceMethod
		for _, method := range methods {
			if reduceMethod == nil {
				reduceMethod = method
				continue
			}
			if !gotypes.Identical(method.ReturnType, reduceMethod.ReturnType) {
				return fmt.Errorf(
					"reduce methods %v and %v differ return types",
					method.MethodName, reduceMethod.MethodName)
			}
		}
		assert(reduceMethod != nil && reduceMethod.ReturnType != nil)
		rule, ok := s.Grammar.GetSymbol(ruleName).(*grammar.Rule)
		if !ok {
			return fmt.Errorf(
				"method %v has no corresponding rule",
				reduceMethod.MethodName)
		}
		s.ReduceTypes[rule] = reduceMethod.ReturnType
		fmt.Println(rule.Name, reduceMethod.ReturnType)
	}

	// Determine the Go type of the reduce-artifact of each rule.
	// Process generated rules this time.
	changed := true
	for changed {
		changed = false
		for _, prod := range s.Grammar.Prods {
			rule := s.Grammar.ProdRule(prod)
			typ := s.getReduceTypeForGeneratedRule(rule, prod)
			if typ == nil {
				continue
			}
			existing := s.ReduceTypes[rule]
			if existing != nil {
				assert(gotypes.Identical(existing, typ))
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
			return fmt.Errorf(
				"rule %v does not have a ruduce method",
				rule.Name)
		}
	}

	// Assign each method to a production.
	// Only non-generated methods at this time.
	for prodIndex, prod := range s.Grammar.Prods {
		rule := s.Grammar.ProdRule(prod)
		if rule.Generated != grammar.NotGenerated {
			continue
		}
		method := s.findMethodForProd(prod, s.ReduceMethods[rule.Name])
		if method == nil {
			return fmt.Errorf(
				"there is no reduce method for %v prod #%v",
				rule.Name, prodIndex+1)
		}
		reduceType := method.ReturnType
		if existing := s.ReduceTypes[rule]; existing == nil {
			s.ReduceTypes[rule] = reduceType
		} else if !gotypes.Identical(existing, reduceType) {
			return fmt.Errorf(
				"conflicting reduce types for %v: %v and %v",
				rule.Name, existing, reduceType)
		}
		s.ReduceMap[prod] = method
	}

	for prod, method := range s.ReduceMap {
		if len(method.Params) != len(prod.Terms) {
			return fmt.Errorf(
				"%v: prod has %v terms but reduce method has %v parameters",
				method.MethodName,
				len(prod.Terms),
				len(method.Params))
		}
		for i, param := range method.Params {
			termSym := s.Grammar.TermSymbol(prod.Terms[i])
			termReduceType := s.Token.Type()
			if cRule, ok := termSym.(*grammar.Rule); ok {
				termReduceType = s.ReduceTypes[cRule]
			}
			if !gotypes.AssignableTo(termReduceType, param.Type) {
				return fmt.Errorf(
					"%v: param %v has type %v but term symbol %v has reduce type %v",
					method.MethodName,
					i,
					param.Type,
					termSym.SymName(),
					termReduceType.String())
			}
		}
	}

	return nil
}

func (s *State) findMethodForProd(
	prod *grammar.Prod,
	methods []*ReduceMethod,
) *ReduceMethod {

	isMatch := func(method *ReduceMethod) bool {
		if len(method.Params) != len(prod.Terms) {
			return false
		}
		for i, param := range method.Params {
			termSym := s.Grammar.TermSymbol(prod.Terms[i])
			termReduceType := s.Token.Type()
			if cRule, ok := termSym.(*grammar.Rule); ok {
				termReduceType = s.ReduceTypes[cRule]
			}
			if !gotypes.AssignableTo(termReduceType, param.Type) {
				return false
			}
		}
		return true
	}

	for _, method := range methods {
		if isMatch(method) {
			return method
		}
	}

	return nil
}

func (s *State) getReduceTypeForGeneratedRule(
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
