package lisp

import (
	"fmt"
	"strconv"
)

type Token interface {
	String() string
}

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

func (t LParenToken) String() string {
	return "<LParen>"
}

func (t RParenToken) String() string {
	return "<RParen>"
}

func (t StringToken) String() string {
	return fmt.Sprintf("<String: %s>", t.Value)
}

func (t NumericToken) String() string {
	return fmt.Sprintf("<Numeric: %v>", t.Value)
}

func (t PlusOperatorToken) String() string {
	return "<PlusOperator>"
}

func parseToken(token string) Token {
	if token == "(" {
		return LParenToken{}
	} else if token == ")" {
		return RParenToken{}
	} else if token == "+" {
		return PlusOperatorToken{}
	} else {
		num, err := strconv.ParseFloat(token, 64)
		if err != nil {
			return StringToken{Value: token}
		} else {
			return NumericToken{Value: num}
		}
	}
}

func Tokenize(tokens []string) []Token {
	result := []Token{}

	for _, token := range tokens {
		result = append(result, parseToken(token))
	}

	return result
}
