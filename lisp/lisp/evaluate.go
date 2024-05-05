package lisp

type (
	Environment struct{}
	Result      struct{}
)

func evaluate(exp Expression, env Environment) (error, Result, Environment) {
	return nil, Result{}, Environment{}
}

// It would be nice to have
// save environment
// load environment
