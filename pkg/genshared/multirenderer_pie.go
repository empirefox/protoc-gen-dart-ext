package genshared

// Append will return a new slice with the elements appended to the end.
//
// It is acceptable to provide zero arguments.
func (ss MultiRenderer) Append(elements ...Renderer) MultiRenderer {
	// Copy ss, to make sure no memory is overlapping between input and
	// output. See issue #97.
	result := append(MultiRenderer{}, ss...)

	result = append(result, elements...)
	return result
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
