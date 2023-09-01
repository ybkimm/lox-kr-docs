package codegen

import (
	gotoken "go/token"
	"os"
	"path/filepath"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errlogger"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen/grammar"
)

func ParseGrammar(fset *gotoken.FileSet, dir string, errs *errlogger.ErrLogger) *grammar.AugmentedGrammar {
	loxFiles, err := filepath.Glob(filepath.Join(dir, "*.lox"))
	if err != nil {
		errs.Errorf(gotoken.Position{}, "%v", err)
		return nil
	}

	if len(loxFiles) == 0 {
		errs.Errorf(gotoken.Position{}, "%v contains no .lox files", dir)
		return nil
	}

	grammar := new(grammar.Grammar)
	for _, loxFile := range loxFiles {
		loxFileData, err := os.ReadFile(loxFile)
		if err != nil {
			errs.Errorf(gotoken.Position{}, "%v", err)
			return nil
		}
		file := fset.AddFile(loxFile, -1, len(loxFileData))
		spec, ok := parser.Parse(file, loxFileData, errs)
		if !ok {
			return nil
		}

		addParserToGrammar(fset, spec, grammar)
	}

	agrammar := grammar.ToAugmentedGrammar(errs)
	if errs.HasError() {
		return nil
	}

	return agrammar
}

func addParserToGrammar(fset *gotoken.FileSet, parser *ast.Parser, g *grammar.Grammar) {
	for _, decl := range parser.Decls {
		switch decl := decl.(type) {
		case *ast.CustomTokenDecl:
			for _, token := range decl.CustomTokens {
				terminal := &grammar.Terminal{
					Name:  token.Name,
					Alias: token.Alias,
					Pos:   fset.Position(token.Bounds().Begin),
				}
				g.Terminals = append(g.Terminals, terminal)
			}
		case *ast.Rule:
			rule := &grammar.Rule{
				Name: decl.Name,
				Pos:  fset.Position(decl.Bounds().Begin),
			}
			for _, astProd := range decl.Prods {
				prod := &grammar.Prod{
					Pos: fset.Position(astProd.Bounds().Begin),
				}
				for _, astTerm := range astProd.Terms {
					term := termASTToGrammar(astTerm)
					term.Pos = fset.Position(astTerm.Bounds().Begin)
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
}

func termASTToGrammar(astTerm *ast.Term) *grammar.Term {
	switch astTerm.Type {
	case ast.Simple:
		return &grammar.Term{
			Type:  grammar.Simple,
			Name:  astTerm.Name,
			Alias: astTerm.Alias,
		}
	case ast.ZeroOrMore, ast.OneOrMore, ast.ZeroOrOne:
		return &grammar.Term{
			Type:  grammar.TermType(astTerm.Type),
			Child: termASTToGrammar(astTerm.Child),
		}
	case ast.List:
		return &grammar.Term{
			Type:  grammar.List,
			Child: termASTToGrammar(astTerm.Child),
			Sep:   termASTToGrammar(astTerm.Sep),
		}
	default:
		panic("not-reached")
	}
}
