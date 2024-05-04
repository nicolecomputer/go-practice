package lisp

import "unicode"

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
	LParenToken  struct{}
	RParentToken struct{}
)

func tokenize(tokens []string) []Token {
	result := []Token{}

	return result
}
