package imports

import (
	"fmt"
)

type Assets []Field_Asset

func NewAssets(in []string) (Assets, error) {
	x := make(Assets, 0, len(in))
	invalid := make([]string, 0, len(in))
	for _, s := range in {
		a, ok := Field_Asset_value[s]
		if ok {
			x = append(x, Field_Asset(a))
		} else {
			invalid = append(invalid, s)
		}
	}

	if len(invalid) != 0 {
		return nil, fmt.Errorf("assets not found: %v", invalid)
	}
	return x, nil
}

func (s Assets) Len() int           { return len(s) }
func (s Assets) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s Assets) Less(i, j int) bool { return s[i] < s[j] }

func (x Assets) Strings() (s []string) {
	if x == nil {
		return nil
	}
	s = make([]string, len(x))
	for i, a := range x {
		s[i] = a.String()
	}
	return
}

func (x Field_Asset) Valid() bool {
	_, ok := Field_Asset_name[int32(x)]
	return ok
}

func (x Field_Asset) Suffix() string {
	if x == Field_Value {
		return ""
	}
	return x.String()
}
