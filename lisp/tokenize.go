package lisp

import "unicode"

type token string

func shouldAppendToken(token string) bool {
	return len(token) > 0
}

func lexer(str string) []token {
	result := []token{}

	var currentToken string = ""

	for _, rune := range str {
		t := token(rune)
		if unicode.IsSpace(rune) {
			if shouldAppendToken(currentToken) {
				result = append(result, token(currentToken))
				currentToken = ""
			}
		} else if t == "(" || t == ")" {
			if shouldAppendToken(currentToken) {
				result = append(result, token(currentToken))
				currentToken = ""
			}

			result = append(result, t)
		} else {
			currentToken += string(rune)
		}
	}

	return result
}
