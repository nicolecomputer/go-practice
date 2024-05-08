package lisp

import (
	parser "lisp/parse"
)

type (
	Environment struct{}
	Result      struct{}
)

func evaluate(exp parser.Expression, env Environment) (error, Result, Environment) {
	return nil, Result{}, Environment{}
}

// It would be nice to have
// save environment
// load environment
