package lisp

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func shouldAppendToken(token string) bool {
	return len(token) > 0
}

func lex(str string) []string {
	result := []string{}

	var currentToken string = ""

	for _, rune := range str {
		t := string(rune)
		if unicode.IsSpace(rune) {
			if shouldAppendToken(currentToken) {
				result = append(result, string(currentToken))
				currentToken = ""
			}
		} else if t == "(" || t == ")" {
			if shouldAppendToken(currentToken) {
				result = append(result, string(currentToken))
				currentToken = ""
			}

			result = append(result, t)
		} else {
			currentToken += string(rune)
		}
	}

	if shouldAppendToken(currentToken) {
		result = append(result, string(currentToken))
	}

	return result
}

type Token interface{}

type (
	LParenToken struct{}
	RParenToken struct{}
	StringToken struct {
		Value string
	}
	NumericToken struct {
		Value float64
	}
	PlusOperatorToken struct{}
)

func tokenize(tokens []string) []Token {
	result := []Token{}

	for _, token := range tokens {
		if token == "(" {
			result = append(result, LParenToken{})
		} else if token == ")" {
			result = append(result, RParenToken{})
		} else if token == "+" {
			result = append(result, PlusOperatorToken{})
		} else {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				result = append(result, StringToken{Value: token})
			} else {
				result = append(result, NumericToken{Value: num})
			}
		}
	}

	return result
}

// PARSE

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
	return parseTokens(tokenize(lex(input)))
}
