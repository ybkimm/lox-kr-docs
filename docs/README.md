```
  /\ \       /\  __ \   /\_\_\_\
  \ \ \____  \ \ \/\ \  \/_/\_\/_
   \ \_____\  \ \_____\   /\_\/\_\
    \/_____/   \/_____/   \/_/\/_/
```

Lox is a parser/lexer generator for Go.

## Why use Lox?

Like other parser generators (e.g. Yacc), Lox generates a parser implementation
from a grammar. Unlike most parser generators, Lox was designed specifically for
the Go language. The grammar is separate from action code, and actions are
matched to rules by method signature. The grammar is kept concise, and actions
are type-checked by the Go compiler. Lox is also a lexer generator (like Lex):
the parser and the lexer are defined in a common grammar specification.

## Getting started

## Features

* LALR(1) parsing algorithm (similar to Yacc)
* DFA lexer algorithm (similar to Lex)
* Grammar separate from code
* Action-rule matching by method signature
* Inferred action artifact type
* Artifact bounds annotation
* Common parser generator features, including:
  * Precedence
  * Error recovery
  * Etc.
