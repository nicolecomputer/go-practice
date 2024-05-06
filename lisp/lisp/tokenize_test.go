package lisp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Run("()", func(t *testing.T) {
		want := []Token{LParenToken{}, RParenToken{}}

		input := "()"
		got := Tokenize(Lex(input))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(123)", func(t *testing.T) {
		want := []Token{LParenToken{}, NumericToken{Value: 123}, RParenToken{}}

		input := "(123)"
		got := Tokenize(Lex(input))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 12 90)", func(t *testing.T) {
		want := []Token{
			LParenToken{},
			PlusOperatorToken{},
			NumericToken{Value: 12},
			NumericToken{Value: 90},
			RParenToken{},
		}

		input := "(+ 12 90)"
		got := Tokenize(Lex(input))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})
}
