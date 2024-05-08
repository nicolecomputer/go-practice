package lisp

import (
	"errors"
	"fmt"
	"strings"
)

type Expression interface {
	String() string
}

type (
	List struct {
		Children []Expression
	}

	NumberAtom struct {
		Value float64
	}

	SymbolicAtom struct {
		Value string
	}

	OperatorAtom struct {
		Value string
	}
)

func (list List) String() string {
	children := []string{}

	for _, child := range list.Children {
		children = append(children, child.String())
	}

	return fmt.Sprintf("(%s)", strings.Join(children, ", "))
}

func (p NumberAtom) String() string {
	return fmt.Sprintf("[Number Atom %v]", p.Value)
}

func (p SymbolicAtom) String() string {
	return fmt.Sprintf("[Symbolic Atom %s]", p.Value)
}

func (p OperatorAtom) String() string {
	return fmt.Sprintf("[Operator Atom %s]", p.Value)
}

// TODO: This is a good place for go generics
func pop(alist *[]Token) Token {
	rv := (*alist)[0]
	*alist = (*alist)[1:]
	return rv
}

func parseList(remainingTokens *[]Token) (error, Expression) {
	var children []Expression = []Expression{}

	var token Token
	for len(*remainingTokens) > 0 {
		token = pop(remainingTokens)
		switch token.(type) {
		case LParenToken:
			err, childList := parseList(remainingTokens)
			if err != nil {
				return err, nil
			}
			children = append(children, childList)
		case RParenToken:
			return nil, List{Children: children}
		default:
			err, child := parseAtom(token)
			if err != nil {
				return err, nil
			}
			children = append(children, child)
		}
	}

	// This is an error, we got to the end of the tokens and didn't find a closing paren
	return errors.New("Unterminated list"), nil
}

func parseAtom(token Token) (error, Expression) {
	switch t := token.(type) {
	case StringToken:
		return nil, SymbolicAtom{Value: t.Value}
	case NumericToken:
		return nil, NumberAtom{Value: t.Value}
	case PlusOperatorToken:
		return nil, OperatorAtom{Value: "+"}
	default:
		return errors.New("Unknown token type"), nil
	}
}

func parseTokens(tokens []Token) (error, Expression) {
	remainingTokens := &tokens

	if len(*remainingTokens) == 1 {
		return parseAtom(pop(remainingTokens))
	}

	for len(*remainingTokens) > 0 {
		token := pop(remainingTokens)

		switch token.(type) {
		case LParenToken:
			return parseList(remainingTokens)
		}
	}

	return errors.New("Parse error"), nil
}

func Parse(input string) (error, Expression) {
	return parseTokens(Tokenize(Lex(input)))
}
