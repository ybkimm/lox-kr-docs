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

%token<tok> EOF LEXERR

%token<tok> ID LITERAL
%token<tok> '=' '.' '|' '*' '+' '?' '#'

%type<prod> syntax
%type<prod> decls
%type<prod> decls_opt
%type<prod> decl
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

S: syntax EOF;

syntax: decls_opt
        {
          $$ = &grammar.Syntax{
            Decls: cast[[]grammar.Decl]($1),
          }
        }
      ;

decls: decls decl
       {
         $$ = append(
           cast[[]grammar.Decl]($1),
           cast[grammar.Decl]($2),
         )
       }
     | decl
       {
         $$ = []grammar.Decl{
           cast[grammar.Decl]($1),
         }
       }
     ;

decls_opt: decls
         | { $$ =nil }
         ;

decl: rule;

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
                 cast[*grammar.Prod]($2),
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
