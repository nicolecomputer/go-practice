package lisp

import (
	"errors"
	"fmt"
	"strconv"
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

type Expression interface{}

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

func (p NumberAtom) String() string {
	return fmt.Sprintf("[Number Atom %v]", p.Value)
}

// TODO: This is a good place for go generics
func pop(alist *[]Token) Token {
	f := len(*alist)
	rv := (*alist)[f-1]
	*alist = (*alist)[:f-1]
	return rv
}

func parseList(tokenStack *[]Token) {
	// Create new list
	// for each token
	//    if token is rparen
	//        return list
	//    if token is lparen
	//        parseList
	//    else
	//        appendtoken
	// Keep adding childen to list until RParen
	// return List
}

func parseAtom(token Token) (error Expression) {
	switch t := token.(type) {
	case StringToken:
		return SymbolicAtom{Value: t.Value}
	case NumericToken:
		return NumberAtom{Value: t.Value}
	default:
		return errors.New("Unknown token type")
	}
}

func Parse(tokens []Token) Expression {
	remainingTokens := &tokens

	if len(*remainingTokens) == 1 {
		return parseAtom(pop(remainingTokens))
	}

	// for len(*remainingTokens) > 0 {
	// token := pop(remainingTokens)
	// }

	return List{}
}

// IsValid
// - NOPE: Check if left parens count matches right parens count
// ()))))()
// - Check if parens are balanced
