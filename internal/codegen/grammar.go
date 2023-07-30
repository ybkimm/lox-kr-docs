package codegen

import (
	"fmt"
	gotoken "go/token"
	"os"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func ParseGrammar(fset *gotoken.FileSet, dir string) (*grammar.AugmentedGrammar, error) {
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
		file := fset.AddFile(loxFile, -1, len(loxFileData))
		spec, err := parser.Parse(file, loxFileData)
		if err != nil {
			return nil, err
		}
		err = addSpecToGrammar(spec, grammar)
		if err != nil {
			return nil, err
		}
	}

	agrammar, err := grammar.ToAugmentedGrammar()
	if err != nil {
		return nil, err
	}

	return agrammar, nil
}

func addSpecToGrammar(spec *ast.Spec, g *grammar.Grammar) error {
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
						if astProd.Qualifier != nil {
							switch astProd.Qualifier.Associativity {
							case ast.Left:
								prod.Associativity = grammar.Left
							case ast.Right:
								prod.Associativity = grammar.Right
							default:
								panic("not-reached")
							}
							prod.Precence = astProd.Qualifier.Precedence
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
	return nil
}
