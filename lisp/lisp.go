package lisp

// type Atom interface {
// 	IsEqual(a2 Atom) bool
// }

// type NumericAtom = float64

// func (a NumericAtom) IsEqual(a2 Atom) bool {
// 	switch v := a2.(type) {
// 	case NumericAtom:
// 		return a == a2
// 	default:
// 		return false
// 	}
// }

// // type List struct {
// // 	Children []Atom
// // }

// // type Item interface {
// // 	isLispItem() bool
// // }

// // func Parse(input string) List {
// // 	return List{}
// // }

// // func (l *List) Equal(l2 *List) bool {
// // 	if len(l.Children) != len(l2.Children) {
// // 		return false
// // 	}

// // 	return true
// // }

// // func (a Atom) Equal(a2 Atom) bool {
// // 	return a.IsEqual(a2)
// // }

// // func parseAtom(input string) Atom {
// // 	// If open-quote go until closed quote
// // 	// Else Go until space or EOF
// // 	return Atom()
// // }
