%{
package parser

import (
  "reflect"

  "github.com/dcaiafa/lox/internal/token"
  "github.com/dcaiafa/lox/internal/ast"
)

func isNil(i interface{}) bool {
   if i == nil {
      return true
   }
   switch reflect.TypeOf(i).Kind() {
   case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
    //use of IsNil method
    return reflect.ValueOf(i).IsNil()
   }
   return false
}

func cast[T any](v any) T {
  if isNil(v) {
    var zero T
    return zero
  }
	return v.(T)
}

func listAppend[T any](xs any, x any) []T {
  return append(
    cast[[]T](xs),
    cast[T](x),
  )
}

func listOne[T any](x any) []T {
  return []T{
    cast[T](x),
  }
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
%type<prod> parser_decls
%type<prod> parser_decls_opt
%type<prod> parser_decl
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
%type<prod> lexer
%type<prod> lexer_decls
%type<prod> lexer_decls_opt
%type<prod> lexer_decl
%type<prod> custom_token_decl
%type<prod> custom_tokens
%type<prod> custom_token

%start S

%%

S: spec  { yylex.(*lex).Spec = cast[*ast.Spec]($1) };

spec: sections
      {
        $$ = &ast.Spec{
          Sections: cast[[]ast.Section]($1),
        }
      }
    ;

sections: sections section  
          { 
            $$ = append(
              cast[[]ast.Section]($1),
              cast[ast.Section]($2),
            )
          }
        | section
          {
            $$ = []ast.Section{
              cast[ast.Section]($1),
            }
          }
        ;

section: parser | lexer;

parser: kPARSER parser_decls_opt
        {
          $$ = &ast.Parser{
            Decls: cast[[]ast.ParserDecl]($2),
          }
        }
      ;

parser_decls: parser_decls parser_decl
       {
         $$ = listAppend[ast.ParserDecl]($1, $2)
       }
     | parser_decl
       {
         $$ = listOne[ast.ParserDecl]($1)
       }
     ;

parser_decls_opt: parser_decls
         | { $$ = nil }
         ;

parser_decl: rule
           ;

rule: ID '=' productions '.'
      {
        $$ = &ast.Rule{
          Name: $1.Str,
          Prods: cast[[]*ast.Prod]($3),
        }
      }
    ;

productions: productions '|' production
             {
               $$ = append(
                 cast[[]*ast.Prod]($1), 
                 cast[*ast.Prod]($3),
               )
             }
           | production
             {
               $$ = []*ast.Prod{
                 cast[*ast.Prod]($1),
               }
             }
           ;

production: qterms label_opt
            {
              $$ = &ast.Prod{
                Terms: cast[[]*ast.Term]($1),
                Label: cast[*ast.Label]($2),
              } 
            }
          ;

qterms: qterms qterm
        {
          $$ = append(
            cast[[]*ast.Term]($1),
            cast[*ast.Term]($2),
          )
        }
      | qterm
        {
          $$ = []*ast.Term{
            cast[*ast.Term]($1),
          }
        }
      ;

qterm: term qualifier_opt
       {
         term := cast[*ast.Term]($1)
         qualifier := cast[ast.Qualifier]($2)
         term.Qualifier = qualifier
         $$ = term
       }
     ;

term: ID       { $$ = &ast.Term{ Name: $1.Str } }
    | LITERAL  { $$ = &ast.Term{ Literal: $1.Str} }
    ;

qualifier: '*'  { $$ = ast.ZeroOrMore }
         | '+'  { $$ = ast.OneOrMore }
         | '?'  { $$ = ast.ZeroOrOne }
         ;

qualifier_opt: qualifier
             | { $$ = ast.NoQualifier }
             ;

label: '#' ID 
       {
         $$ = &ast.Label{
           Label: $2.Str,
         }
       }
     ;

label_opt: label
         | { $$ = nil }
         ;

lexer: kLEXER lexer_decls_opt
       {
         $$ = &ast.Lexer{
           Decls: cast[[]ast.LexerDecl]($2),
         }
       }
       ;

lexer_decls: lexer_decls lexer_decl
             { $$ = listAppend[ast.LexerDecl]($1, $2) }
           | lexer_decl
             { $$ = listOne[ast.LexerDecl]($1) }
           ;

lexer_decls_opt: lexer_decls
               | /* empty */
                 { $$ = nil }
               ;

lexer_decl: custom_token_decl
          ;

custom_token_decl: kCUSTOM custom_tokens
                   {
                     $$ = &ast.CustomTokenDecl{
                       CustomTokens: cast[[]*ast.CustomToken]($2),
                     }
                   }
                 ;

custom_tokens: custom_tokens custom_token
               {
                 $$ = listAppend[*ast.CustomToken]($1, $2)
               }
             | custom_token
               {
                 $$ = listOne[*ast.CustomToken]($1)
               }
             ;

custom_token: ID      { $$ = &ast.CustomToken{Name: $1.Str} }
            | LITERAL { $$ = &ast.CustomToken{Literal: $1.Str} }
              
