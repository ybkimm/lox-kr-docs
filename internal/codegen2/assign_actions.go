package codegen2

import (
	gotypes "go/types"
	"strings"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

func (c *context) AssignActions() bool {
	c.RuleGoTypes = make(map[*lr2.Rule]gotypes.Type)

	methods := c.getActionMethods()
	if methods == nil {
		return false
	}

	// Map of name => Rule.
	rules := make(map[string]*lr2.Rule, len(c.ParserGrammar.Rules))
	for _, rule := range c.ParserGrammar.Rules {
		rules[rule.Name] = rule
	}

	// Determine the Go-type of reduce artifacts by matching action method names
	// to rules. All action methods matching the same rule name must have the same
	// return type.
	for ruleName, ruleMethods := range methods {
		var firstMethod *actionMethod
		for _, method := range ruleMethods {
			if firstMethod == nil {
				firstMethod = method
				continue
			}
			if !gotypes.Identical(method.Return, firstMethod.Return) {
				c.Errs.Errorf(
					c.Fset.Position(method.Method.Pos()),
					"action return type conflict: %v returns %v but %v returns %v",
					method.Name(), method.Return,
					firstMethod.Name(), firstMethod.Return)
				c.Errs.Infof(
					c.Fset.Position(firstMethod.Method.Pos()),
					"%v is defined here",
					firstMethod.Name())
			}
		}
		assert.True(firstMethod != nil && firstMethod.Return != nil)
		rule := rules[ruleName]
		if rule == nil {
			c.Errs.Errorf(
				c.Fset.Position(firstMethod.Method.Pos()),
				"action method %v: no rule named %v",
				firstMethod.Name(), ruleName)
			continue
		}
		c.RuleGoTypes[rule] = firstMethod.Return
	}
	if c.Errs.HasError() {
		return false
	}

	// Determine the Go type of reduce artifacts for generated rules, which are
	// derived from the Go-types for user-provided rules determined above. For
	// example, if the Go type for rule 'expr' is 'int', then the Go type for the
	// generated rule that replaced a 'expr+' term is '[]int'.
	changed := true
	for changed {
		changed = false
		for _, prod := range c.ParserGrammar.Prods {
			rule := c.ParserGrammar.GetRule(prod.Rule)
			typ := c.getReduceTypeForGeneratedRule(rule, prod)
			if typ == nil {
				// Rule was not generated, or we can't determine the Go-type for the
				// rule based on this specific Prod.
				continue
			}
			existing := c.RuleGoTypes[rule]
			if existing != nil {
				assert.True(gotypes.Identical(existing, typ))
				continue
			}
			c.RuleGoTypes[rule] = typ
			changed = true
		}
	}

	// Check that every rule has been assigned a Go-type.
	for _, rule := range c.ParserGrammar.Rules {
		if rule.Generated == lr2.GeneratedSPrime {
			// Except for S', which is never reduced.
			continue
		}
		if c.RuleGoTypes[rule] == nil {
			c.Errs.Errorf(
				rule.Position, "rule missing action method: %v",
				rule.Name)
		}
	}

	return !c.Errs.HasError()
}

func (c *context) getActionMethods() map[string][]*actionMethod {
	actionMethods := make(map[string][]*actionMethod)
	for i := 0; i < c.ParserType.NumMethods(); i++ {
		goMethod := c.ParserType.Method(i)
		if goMethod.Name() == onReduce {
			// The parser implements onReduce.
			// The generated parser should call it.
			c.HasOnReduce = true
			continue
		}
		rule := ruleFromMethod(goMethod.Name())
		if rule == "" {
			continue
		}
		sig := goMethod.Type().(*gotypes.Signature)
		if sig.Results().Len() != 1 {
			c.Errs.Errorf(
				c.Fset.Position(goMethod.Pos()),
				"%v: action method must return a single value",
				goMethod.Name())
			continue
		}

		method := &actionMethod{
			Method: goMethod,
			Return: sig.Results().At(0).Type(),
		}
		method.Params = make([]gotypes.Type, sig.Params().Len())
		for i := 0; i < sig.Params().Len(); i++ {
			method.Params[i] = sig.Params().At(i).Type()
		}
		actionMethods[rule] = append(actionMethods[rule], method)
	}
	if c.Errs.HasError() {
		return nil
	}
	return actionMethods
}

func (c *context) getReduceTypeForGeneratedRule(
	rule *lr2.Rule,
	prod *lr2.Prod,
) gotypes.Type {
	switch rule.Generated {
	case lr2.NotGenerated, lr2.GeneratedSPrime:
		// S' is never reduced.
		return nil
	case lr2.GeneratedZeroOrOne:
		// a = b c?
		//  =>
		// a = b a'
		// a' = c | e
		if prod != c.ParserGrammar.GetProd(rule.Prods[0]) {
			return nil
		}
		termC := prod.Terms[0]
		if lr2.IsRule(termC) {
			return c.RuleGoTypes[c.ParserGrammar.GetRule(termC)]
		} else {
			return c.TokenType
		}

	case lr2.GeneratedOneOrMore, lr2.GeneratedList:
		// a = b c+
		//  =>
		// a = b a'
		// a' = a' c | c      (OneOrMore)
		// a' = a' sep c | c  (List)
		if prod != c.ParserGrammar.GetProd(rule.Prods[1]) {
			return nil
		}
		termC := prod.Terms[0]
		typeC := c.TokenType
		if lr2.IsRule(termC) {
			typeC = c.RuleGoTypes[c.ParserGrammar.GetRule(termC)]
		}
		return gotypes.NewSlice(typeC)

	default:
		panic("unreachable")
	}
}

// ruleFromMethod returns the name of the rule corresponding to an action
// method. Returns an empty string if the method is not an action.
//
// Action methods are like the following:
//
//	               prefix         suffix
//	                 +-+         +------+
//	func (p *parser) on_expr_term__simple(...)
//	                    +-------+
//	                       rule
//
// prefix: identifies an action method. If the method starts with "on_", then it
// is an action method.
//
// rule: anything after the prefix, and before an optional "__" is the name of
// the rule, and must match a corresponding grammar rule.
//
// suffix: the suffix is optional and is the remaining of the method name
// starting from "__" (double underscore). The suffix is completely ignored
// during action <=> production matching, but is necessary when two or more
// productions from a same rule require different actions.
//
// For example, given the following parser rules:
//
//	 expr = expr '+' expr @left(1)
//		    | expr '-' expr @left(1)
//		    | NUM ;
//
// Could have the following corresponding actions:
//
//	// matches:
//	// expr '+' expr @left(1)
//	// expr '-' expr @left(1)
//	func (p *parser) on_expr__binary(l any, op Token, r any) any {...}
//
//	// matches:
//	// NUM ;
//	func (p *parser) on_expr__num(num Token) any {...}
func ruleFromMethod(method string) string {
	const prefix = "on_"
	const sep = "__"
	if !strings.HasPrefix(method, prefix) {
		return ""
	}
	rule := method[len(prefix):]
	sepIdx := strings.Index(rule, sep)
	if sepIdx != -1 {
		rule = rule[:sepIdx]
	}
	return rule
}
