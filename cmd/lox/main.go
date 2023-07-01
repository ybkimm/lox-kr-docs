package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/dcaiafa/lox/internal/ast"
	"github.com/dcaiafa/lox/internal/errs"
	"github.com/dcaiafa/lox/internal/parser"
	"github.com/dcaiafa/lox/internal/parsergen"
)

func main() {
	flag.Parse()

	grammarFile := flag.Arg(0)
	err := run(grammarFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed: %v\n", err)
		os.Exit(1)
	}
}

func run(grammarFile string) error {
	grammarData, err := os.ReadFile(grammarFile)
	if err != nil {
		return err
	}
	errs := errs.New()
	spec := parser.Parse(grammarFile, grammarData, errs)
	if errs.HasErrors() {
		errs.Dump(os.Stderr)
		return fmt.Errorf("errors ocurred")
	}

	grammar := toParserGrammar(spec)

	fmt.Println("Before:")
	grammar.Print(os.Stdout)

	err = grammar.Analyze()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return fmt.Errorf("errors ocurred")
	}

	fmt.Println("After:")
	grammar.Print(os.Stdout)

	return nil
}

func toParserGrammar(spec *ast.Spec) *parsergen.Grammar {
	grammar := new(parsergen.Grammar)

	for _, section := range spec.Sections {
		switch section := section.(type) {
		case *ast.Lexer:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.CustomTokenDecl:
					for _, token := range decl.CustomTokens {
						terminal := &parsergen.Terminal{
							Name: customTokenName(token),
						}
						grammar.Terminals = append(grammar.Terminals, terminal)
					}
				default:
					panic("not-reached")
				}
			}
		case *ast.Parser:
			for _, decl := range section.Decls {
				switch decl := decl.(type) {
				case *ast.Rule:
					rule := &parsergen.Rule{
						Name: decl.Name,
					}
					for _, astProd := range decl.Prods {
						prod := &parsergen.Prod{}
						for _, astTerm := range astProd.Terms {
							term := &parsergen.Term{
								Name: termName(astTerm),
							}
							switch astTerm.Qualifier {
							case ast.NoQualifier:
								term.Qualifier = parsergen.NoQualifier
							case ast.ZeroOrMore:
								term.Qualifier = parsergen.ZeroOrMore
							case ast.OneOrMore:
								term.Qualifier = parsergen.OneOrMore
							case ast.ZeroOrOne:
								term.Qualifier = parsergen.ZeroOrOne
							default:
								panic("not-reached")
							}
							prod.Terms = append(prod.Terms, term)
						}
						rule.Prods = append(rule.Prods, prod)
					}
					grammar.Rules = append(grammar.Rules, rule)

				default:
					panic("not-reached")
				}
			}
		default:
			panic("not-reached")
		}
	}

	return grammar
}

func customTokenName(token *ast.CustomToken) string {
	if token.Name != "" {
		return token.Name
	}
	return tokenLiteralName(token.Literal)
}

func tokenLiteralName(l string) string {
	return "TOKEN__" + literalName(l)
}

func termName(term *ast.Term) string {
	if term.Name != "" {
		return term.Name
	}
	return tokenLiteralName(term.Literal)
}

func literalName(l string) string {
	var str strings.Builder
	str.Grow(len(l))
	for i := 0; i < len(l); i++ {
		c := l[i]
		isAlphaNum :=
			(c >= 'a' && c <= 'z') ||
				(c >= 'A' && c <= 'Z') ||
				(c >= '0' && c <= '9')
		if isAlphaNum {
			str.WriteByte(c)
		} else {
			fmt.Fprintf(&str, "%02x", c)
		}
	}
	return str.String()
}
