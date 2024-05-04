package lisp

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	t.Run("Empty string", func(t *testing.T) {
		want := []token{}
		got := lexer("")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("()", func(t *testing.T) {
		want := []token{"(", ")"}
		got := lexer("()")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(1)", func(t *testing.T) {
		want := []token{"(", "1", ")"}
		got := lexer("(1)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(123)", func(t *testing.T) {
		want := []token{"(", "123", ")"}
		got := lexer("(123)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+)", func(t *testing.T) {
		want := []token{"(", "+", ")"}
		got := lexer("(+)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 200000 30)", func(t *testing.T) {
		want := []token{"(", "+", "10", "200000", "30", ")"}
		got := lexer("(+ 10 200000 30)")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	t.Run("(+ 10 (- 30 6))", func(t *testing.T) {
		want := []token{"(", "+", "10", "(", "-", "30", "6", ")", ")"}
		got := lexer("(+ 10 (- 30 6))")

		if !reflect.DeepEqual(want, got) {
			t.Errorf("expected %s got %s", want, got)
		}
	})

	// Strings with spaces
}
