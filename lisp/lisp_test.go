package lisp

/*
(first (list 1 (+ 2 3) 9))

()
*/

/*
Lisp is made up of

Atoms
Lists

For parsing I'm going to assume all atoms are strings :\

Parse will always return a list

*/

// func TestParse(t *testing.T) {
// 	t.Run("An Empty List", func(t *testing.T) {
// 		want := List{}

// 		input := "()"
// 		got := Parse(input)

// 		if got != want {
// 			t.Errorf("Expected %s got %s", want, got)
// 		}
// 	})
// }

// func TestAtomEquality(t *testing.T) {
// 	t.Run("a != b", func(t *testing.T) {
// 		atom1 := Atom("a")
// 		atom2 := Atom("b")

// 		if atom1.Equal(atom2) {
// 			t.Errorf("%s should not equal %s", atom1, atom2)
// 		}
// 	})

// 	t.Run("123 == 123", func(t *testing.T) {
// 		atom1 := Atom("123")
// 		atom2 := Atom("123")

// 		if !atom1.Equal(atom2) {
// 			t.Errorf("%s should equal %s", atom1, atom2)
// 		}
// 	})
// }

// func TestParseAtom(t *testing.T) {
// 	t.Run("123 is a valid Atom", func(t *testing.T) {})
// 	t.Run("abc is a valid Atom", func(t *testing.T) {})

// 	t.Run("+ is a valid Atom", func(t *testing.T) {})

// 	t.Run("\"Hello World\" is a valid Atom", func(t *testing.T) {})
// }
