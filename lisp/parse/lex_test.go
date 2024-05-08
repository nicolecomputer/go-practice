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
