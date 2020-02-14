package main

import (
	"strings"

	"github.com/iancoleman/strcase"
)

type Qualifier string

// AnyKindOfString
func (q Qualifier) ToCamel() Qualifier {
	return Qualifier(strcase.ToCamel(strings.ReplaceAll(string(q), ".", "_")))
}

type UsedNames []map[Qualifier]bool

func (un UsedNames) Add(name Qualifier) {
	if un[0] == nil {
		un[0] = make(map[Qualifier]bool)
	}
	un[0][name] = true
}
func (un *UsedNames) AddMap(m map[Qualifier]bool) {
	if len(*un) == 0 {
		return
	}
	*un = append(*un, m)
}
func (un *UsedNames) AddSlice(names []Qualifier) {
	m := make(map[Qualifier]bool, len(names))
	for _, n := range names {
		m[n] = true
	}
	un.AddMap(m)
}
func (un UsedNames) Used(name Qualifier) bool {
	for _, m := range un {
		if m[name] {
			return true
		}
	}
	return false
}
