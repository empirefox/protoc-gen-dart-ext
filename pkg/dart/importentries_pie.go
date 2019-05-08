package dart

import (
	"github.com/elliotchance/pie/pie"
)

// ToStrings transforms each element to a string.
func (ss ImportEntries) ToStrings(transform func(*ImportEntry) string) pie.Strings {
	l := len(ss)

	// Avoid the allocation.
	if l == 0 {
		return nil
	}

	result := make(pie.Strings, l)
	for i := 0; i < l; i++ {
		result[i] = transform(ss[i])
	}

	return result
}

// Unselect works the same as Select, with a negated condition. That is, it will
// return a new slice only containing the elements that returned false from the
// condition. The returned slice may contain zero elements (nil).
func (ss ImportEntries) Unselect(condition func(*ImportEntry) bool) (ss2 ImportEntries) {
	for _, s := range ss {
		if !condition(s) {
			ss2 = append(ss2, s)
		}
	}

	return
}
