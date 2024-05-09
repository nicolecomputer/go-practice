# Lisp

This is a toy lisp interpreter.

## Next steps

- Checking if a list is balanced during parsing (e.g. `")))))))()"`)
- Handing strings that contain spaces (e.g. `"berries for all!"`)
- Operators other than +
- Evaluating
  - Evaluating +
- Up arrow in REPL to see last entered expression
- "!!" in REPL for last entered expression

## Guide to reading this code

- The entry point for the REPL (the interactive program that evaluates lisp
  code) is in cli/repl.go
- Parsing
  - Parsing happens in 3 steps
  - First the string that's passed is lexed (`parse/lex.go`) the goal of this
    step is to just break up string into so `(+ 1 2)` will become the list `(`,
    `+`, `1` ,`2`, `)`
  - Next these strings are passed off to be tokenized (`parse/tokenize.go`). In
    this step the strings are interpreted into higher level tokens. Like `(`
    becomes an LParenToken
  - Finally the tokenized tokens are evaluated for meaning in the parser
    (`parse/parse.go`). In this step the tokens are checked if they make sense
    in the order that they were passed in and an abstract syntax tree is built
    up
- Evaluation
  - There currently isn't any code that does evaluation but when there is will
    live in `evaluate/evaluate.go`

There are tests for most everything which helps make sure that each level of
producing a result is doing what it's supposed to and that refactoring is safe.

## Decisions

One decision that I'm unsure about is that this interpreter currently does not
use s-expressions. It builds up full arrays instead. This is different than most
lisps and I'm unsure how it'll play out in the long term.
