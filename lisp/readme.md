## Resources

- https://www.reddit.com/r/golang/comments/1z5hle/writing_a_lisp_interpreter_in_go/
- https://www.jerf.org/iri/post/2917/

## Things that currently do not work

- Checking if a list is balanced during parsing (e.g. `")))))))()"`)
- Handing strings that contain spaces (e.g. `"berries for all!"`)
- Operators other than +
- Evaluating
- Up arrow in REPL to see last entered expression
- "!!" in REPL for last entered expression
