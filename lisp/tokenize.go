package lisp

import (
	"strconv"
	"unicode"
)

func shouldAppendToken(token string) bool {
	return len(token) > 0
}

func lexer(str string) []string {
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

// IsValid
// - NOPE: Check if left parens count matches right parens count
// ()))))()
// - Check if parens are balanced
