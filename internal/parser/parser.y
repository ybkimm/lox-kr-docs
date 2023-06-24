%{
package parser

import (
  "github.com/dcaiafa/lox/internal/token"
)

/*

syntax = production* .

production = ID '=' expression '.' .

expression = term                  # e1
           | expression '|' term   # e2
           .

term = qfactor+ label? .

qfactor = factor qualifier? .

factor = ID | LITERAL .

qualifier = '+' | '*' | '?' .

label = '#' ID .

---

func syntax(t0 []any) any {
  return nil
}

func production(id0 Token, t1 Token, expression2 any, t3 token) any {
  return nil
}

func expression_e1(t0 any) any {
}

func expression_e2(expression0 any, t1 Token, term2 any) any {
}

func term(qfactor []any) any {
}

func qfactor(factor any, qualifier any) any {
}

func factor(t0 Token) any {
}

func qualifier(t0 Token) any {
}

func label(t0 Token, id Token) any {
}

*/

%}

%union {


}

%token EOF
%token LEXERR

%token ID
%token LITERAL
%token '=' '.' '|' '*' '+' '?' '#'

%start syntax

%%

syntax: productions_opt EOF
      ;

productions_opt: productions
               | /*empty*/
               ;

productions: productions production
           | production
           ;

production: ID '=' terms '.'
          ;

terms: term '|' term
     | term
     ;

term: qfactors label_opt
    ;

qfactors: qfactors qfactor
        | qfactor
        ;

qfactor: factor qualifier_opt
       ;

factor: ID
      | LITERAL
      ;

qualifier: '*' | '+' | '?' ;

qualifier_opt: qualifier
             | /*empty*/
             ;

label: '#' ID ;

label_opt: label
         | /*empty*/
         ;
