package lisp

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		want := []string{}
		got := lex("")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("()", func(t *testing.T) {
		want := []string{"(", ")"}
		got := lex("()")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(1)", func(t *testing.T) {
		want := []string{"(", "1", ")"}
		got := lex("(1)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(123)", func(t *testing.T) {
		want := []string{"(", "123", ")"}
		got := lex("(123)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+)", func(t *testing.T) {
		want := []string{"(", "+", ")"}
		got := lex("(+)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 200000 30)", func(t *testing.T) {
		want := []string{"(", "+", "10", "200000", "30", ")"}
		got := lex("(+ 10 200000 30)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 (- 30 6))", func(t *testing.T) {
		want := []string{"(", "+", "10", "(", "-", "30", "6", ")", ")"}
		got := lex("(+ 10 (- 30 6))")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	// Strings with spaces
}

func TestTokenize(t *testing.T) {
	t.Run("()", func(t *testing.T) {
		want := []Token{LParenToken{}, RParenToken{}}

		input := "()"
		got := tokenize(lex(input))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(123)", func(t *testing.T) {
		want := []Token{LParenToken{}, NumericToken{Value: 123}, RParenToken{}}

		input := "(123)"
		got := tokenize(lex(input))

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
		got := tokenize(lex(input))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})
}
