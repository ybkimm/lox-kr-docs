```

   __         ______     __  __
  /\ \       /\  __ \   /\_\_\_\
  \ \ \____  \ \ \/\ \  \/_/\_\/_
   \ \_____\  \ \_____\   /\_\/\_\
    \/_____/   \/_____/   \/_/\/_/

```

Lox is a parser/lexer generator for Go.

## Why use Lox?

Like other parser generators (e.g. yacc), Lox generates a parser implementation
from a grammar. Unlike other parser generators, Lox was designed specifically
for the Go language.

## Getting started

## Features

* LALR(1) parsing algorithm (similar to yacc)
* DFA lexer algorithm (similar to lex)
* Grammar separate from code
* Inferred action artifact types
* Artifact bounds annotation
* Common parser generator features, including:
  * Precedence
  * Error recovery
  * Etc.
