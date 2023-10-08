package codegen2

import (
	gotypes "go/types"
	"strings"

	"github.com/dcaiafa/lox/internal/assert"
	"github.com/dcaiafa/lox/internal/parsergen/lr2"
)

func (c *context) AssignActions() bool {
	c.ReduceGoTypes = make(map[*lr2.Rule]gotypes.Type)

	methods := c.getActionMethods()
	if methods == nil {
		return false
	}

	// Map of name => Rule.
	rules := make(map[string]*lr2.Rule, len(c.ParserGrammar.Rules))
	for _, rule := range c.ParserGrammar.Rules {
		rules[rule.Name] = rule
	}

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
		c.ReduceGoTypes[rule] = firstMethod.Return
	}
	if c.Errs.HasError() {
		return false
	}

	return true
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
