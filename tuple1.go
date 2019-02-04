package fuego

// Tuple1 is a tuple with 1 element.
type Tuple1 struct {
	E1 Entry
}

// Hash returns the hash of this tuple.
func (t Tuple1) Hash() uint32 {
	if t.E1 == nil {
		return 0
	}
	return t.E1.Hash()
}

// Equal returns true if 'o' and 't' are equal.
func (t Tuple1) Equal(o Entry) bool {
	oT, ok := o.(Tuple1)
	return t == o ||
		(ok &&
			(t.E1 != nil && t.E1.Equal(oT.E1)))
}

// Arity is the number of elements in this tuple.
func (t Tuple1) Arity() int {
	return 1
}

// ToSlice returns the elements of this tuple as a Go slice.
func (t Tuple1) ToSlice() []Entry {
	return []Entry{t.E1}
}
