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
%type<prod> productions
%type<prod> productions_opt
%type<prod> production
%type<prod> rules
%type<prod> rule
%type<prod> qfactors
%type<prod> qfactor
%type<prod> factor
%type<prod> qualifier
%type<prod> qualifier_opt
%type<prod> label
%type<prod> label_opt

%start S

%%

S: syntax EOF;

syntax: productions_opt
        {
          $$ = &grammar.Syntax{
            Productions: cast[[]*grammar.Production]($1),
          }
        }
      ;

productions: productions production
             {
               $$ = append(
                 cast[[]*grammar.Production]($1),
                 cast[*grammar.Production]($2),
               )
             }
           | production
             {
               $$ = []*grammar.Production{
                 cast[*grammar.Production]($1),
               }
             }
           ;

productions_opt: productions
               | { $$ =nil }
               ;


production: ID '=' rules '.'
            {
              $$ = &grammar.Production{
                Name: $1.Str,
                rules: cast[[]*grammar.Rule]($2),
              }
            }
          ;

rules: rules '|' rule
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

rule: qfactors label_opt
      {
        $$ = &grammar.Rule{
          Factors: cast[[]*grammar.Factor]($1),
          Label: cast[*grammar.Label]($2),
        } 
      }
    ;

qfactors: qfactors qfactor
          {
            $$ = append(
              cast[[]*grammar.Factor]($1),
              cast[*grammar.Factor]($2),
            )
          }
        | qfactor
          {
            $$ = []*grammar.Factor{
              cast[*grammar.Factor]($1),
            }
          }
        ;

qfactor: factor qualifier_opt
         {
           factor := cast[*grammar.Factor]($1)
           qualifier := cast[grammar.Qualifier]($2)
           factor.Qualifier = qualifier
           $$ = factor
         }
       ;

factor: ID       { $$ = &grammar.Factor{ Name: $1.Str } }
      | LITERAL  { $$ = &grammar.Factor{ Literal: $1.Str} }
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
