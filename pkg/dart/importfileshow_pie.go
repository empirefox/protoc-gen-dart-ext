package dart

// Keys returns the keys in the map. All of the items will be unique.
//
// Due to Go's randomization of iterating maps the order is not deterministic.
func (m ImportFileShow) Keys() []string {
	// Avoid allocation
	l := len(m)
	if l == 0 {
		return nil
	}

	i := 0
	keys := make([]string, len(m))
	for key := range m {
		keys[i] = key
		i++
	}

	return keys
}
