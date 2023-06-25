%{
package parser

import (
  "github.com/dcaiafa/lox/internal/token"
  "github.com/dcaiafa/lox/internal/grammar"
)

func cast[T any](v any) T {
	cv, _ := v.(T)
	return cv
}

%}

%union {
  tok token.Token
  prod interface{}
}

%token<tok> LEXERR

%token<tok> ID LITERAL
%token<tok> '=' '.' '|' '*' '+' '?' '#'
%token<tok> kLEXER kPARSER kCUSTOM

%type<prod> spec
%type<prod> sections
%type<prod> section
%type<prod> parser
%type<prod> rules
%type<prod> rules_opt
%type<prod> rule
%type<prod> productions
%type<prod> production
%type<prod> qterms
%type<prod> qterm
%type<prod> term
%type<prod> qualifier
%type<prod> qualifier_opt
%type<prod> label
%type<prod> label_opt

%start S

%%

S: spec  { yylex.(*lex).Spec = cast[*grammar.Spec]($1) };

spec: sections;

sections: sections section  { $$ = cast[*grammar.Spec]($1).AddSection($2) }
        | section           { $$ = new(grammar.Spec).AddSection($1) }
        ;

section: parser;

parser: kPARSER rules_opt
        {
          $$ = &grammar.Parser{
            Rules: cast[[]*grammar.Rule]($2),
          }
        }
      ;

rules: rules rule
       {
         $$ = append(
           cast[[]*grammar.Rule]($1),
           cast[*grammar.Rule]($2),
         )
       }
     | rule
       {
         $$ = []*grammar.Rule{
           cast[*grammar.Rule]($1),
         }
       }
     ;

rules_opt: rules
         | { $$ = nil }
         ;

rule: ID '=' productions '.'
      {
        $$ = &grammar.Rule{
          Name: $1.Str,
          Prods: cast[[]*grammar.Prod]($3),
        }
      }
    ;

productions: productions '|' production
             {
               $$ = append(
                 cast[[]*grammar.Prod]($1), 
                 cast[*grammar.Prod]($3),
               )
             }
           | production
             {
               $$ = []*grammar.Prod{
                 cast[*grammar.Prod]($1),
               }
             }
           ;

production: qterms label_opt
            {
              $$ = &grammar.Prod{
                Terms: cast[[]*grammar.Term]($1),
                Label: cast[*grammar.Label]($2),
              } 
            }
          ;

qterms: qterms qterm
        {
          $$ = append(
            cast[[]*grammar.Term]($1),
            cast[*grammar.Term]($2),
          )
        }
      | qterm
        {
          $$ = []*grammar.Term{
            cast[*grammar.Term]($1),
          }
        }
      ;

qterm: term qualifier_opt
       {
         term := cast[*grammar.Term]($1)
         qualifier := cast[grammar.Qualifier]($2)
         term.Qualifier = qualifier
         $$ = term
       }
     ;

term: ID       { $$ = &grammar.Term{ Name: $1.Str } }
    | LITERAL  { $$ = &grammar.Term{ Literal: $1.Str} }
    ;

qualifier: '*'  { $$ = grammar.ZeroOrMore }
         | '+'  { $$ = grammar.OneOrMore }
         | '?'  { $$ = grammar.ZeroOrOne }
         ;

qualifier_opt: qualifier
             | { $$ = grammar.NoQualifier }
             ;

label: '#' ID 
       {
         $$ = &grammar.Label{
           Label: $2.Str,
         }
       }
     ;

label_opt: label
         | { $$ = nil }
         ;
