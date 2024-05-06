package lisp

import (
	"reflect"
	"testing"
)

func TestLex(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		want := []string{}
		got := Lex("")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("123", func(t *testing.T) {
		want := []string{"123"}
		got := Lex("123")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("()", func(t *testing.T) {
		want := []string{"(", ")"}
		got := Lex("()")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(1)", func(t *testing.T) {
		want := []string{"(", "1", ")"}
		got := Lex("(1)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(123)", func(t *testing.T) {
		want := []string{"(", "123", ")"}
		got := Lex("(123)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+)", func(t *testing.T) {
		want := []string{"(", "+", ")"}
		got := Lex("(+)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 200000 30)", func(t *testing.T) {
		want := []string{"(", "+", "10", "200000", "30", ")"}
		got := Lex("(+ 10 200000 30)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 (- 30 6))", func(t *testing.T) {
		want := []string{"(", "+", "10", "(", "-", "30", "6", ")", ")"}
		got := Lex("(+ 10 (- 30 6))")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	// Strings with spaces
}

func TestParseTokens(t *testing.T) {
	t.Run("()", func(t *testing.T) {
		want := List{Children: []Expression{}}

		input := "()"
		_, got := parseTokens(Tokenize(Lex(input)))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("123", func(t *testing.T) {
		want := NumberAtom{Value: 123}

		input := "123"
		_, got := parseTokens(Tokenize(Lex(input)))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("abc", func(t *testing.T) {
		want := SymbolicAtom{Value: "abc"}

		input := "abc"
		_, got := parseTokens(Tokenize(Lex(input)))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(abc 123 99 zz)", func(t *testing.T) {
		want := List{Children: []Expression{
			SymbolicAtom{"abc"},
			NumberAtom{123},
			NumberAtom{99},
			SymbolicAtom{"zz"},
		}}

		input := "(abc 123 99 zz)"
		_, got := parseTokens(Tokenize(Lex(input)))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("lists with lists (abc (1 2 3))", func(t *testing.T) {
		want := List{Children: []Expression{
			SymbolicAtom{"abc"},
			List{Children: []Expression{
				NumberAtom{1},
				NumberAtom{2},
				NumberAtom{3},
			}},
		}}

		input := "(abc (1 2 3))"
		_, got := parseTokens(Tokenize(Lex(input)))

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	// TODO:
	// - Error handling for unterminated lists
	// - Error handling for top levels tokens like "abc 123 99 a"
}
