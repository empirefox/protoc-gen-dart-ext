package genshared

// Append will return a new slice with the elements appended to the end. It is a
// wrapper for the internal append(). It is offered as a function so that it can
// more easily chained.
//
// It is acceptable to provide zero arguments.
func (ss MultiRenderer) Append(elements ...Renderer) MultiRenderer {
	return append(ss, elements...)
}

// Extend will return a new slice with the slices of elements appended to the
// end.
//
// It is acceptable to provide zero arguments.
func (ss MultiRenderer) Extend(slices ...MultiRenderer) (ss2 MultiRenderer) {
	ss2 = ss

	for _, slice := range slices {
		ss2 = ss2.Append(slice...)
	}

	return ss2
}
