package codegen

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func ParseGrammar(dir string) (*grammar.AugmentedGrammar, error) {
	loxFiles, err := filepath.Glob(filepath.Join(dir, "*.lox"))
	if err != nil {
		return nil, err
	}

	if len(loxFiles) == 0 {
		return nil, fmt.Errorf("%v contains no .lox files", dir)
	}

	grammar := new(grammar.Grammar)
	for _, loxFile := range loxFiles {
		loxFileData, err := os.ReadFile(loxFile)
		if err != nil {
			return nil, err
		}
		errs := errs.New()
		spec := parser.Parse(loxFile, loxFileData, errs)
		if errs.HasErrors() {
			errs.Dump(os.Stderr)
			return nil, fmt.Errorf("parsing lox files")
		}
		addSpecToGrammar(spec, grammar)
	}

	agrammar, err := grammar.ToAugmentedGrammar()
	if err != nil {
		return nil, err
	}

	return agrammar, nil
}

func addSpecToGrammar(spec *ast.Spec, g *grammar.Grammar) {
	for _, section := range spec.Sections {
		switch section := section.(type) {
		case *ast.Lexer:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.CustomTokenDecl:
					for _, token := range decl.CustomTokens {
						terminal := &grammar.Terminal{
							Name: token.Name,
						}
						g.Terminals = append(g.Terminals, terminal)
					}
				default:
					panic("not-reached")
				}
			}
		case *ast.Parser:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.Rule:
					rule := &grammar.Rule{
						Name: decl.Name,
					}
					for _, astProd := range decl.Prods {
						prod := &grammar.Prod{}
						for _, astTerm := range astProd.Terms {
							term := &grammar.Term{
								Name: astTerm.Name,
							}
							switch astTerm.Qualifier {
							case ast.NoQualifier:
								term.Cardinality = grammar.One
							case ast.ZeroOrMore:
								term.Cardinality = grammar.ZeroOrMore
							case ast.OneOrMore:
								term.Cardinality = grammar.OneOrMore
							case ast.ZeroOrOne:
								term.Cardinality = grammar.ZeroOrOne
							default:
								panic("not-reached")
							}
							prod.Terms = append(prod.Terms, term)
						}
						rule.Prods = append(rule.Prods, prod)
					}
					g.Rules = append(g.Rules, rule)

				default:
					panic("not-reached")
				}
			}
		default:
			panic("not-reached")
		}
	}
}
