package dart

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

type Qualifier string

func (q Qualifier) String() string { return string(q) }

func (q Qualifier) Init() Qualifier {
	return q + "()"
}

func (q Qualifier) InitWith(simpleName ...Qualifier) Qualifier {
	s := make([]string, len(simpleName))
	for i, n := range simpleName {
		s[i] = string(n)
	}
	return q.InitWithString(s...)
}

func (q Qualifier) InitWithStringer(simpleName ...fmt.Stringer) Qualifier {
	s := make([]string, len(simpleName))
	for i, n := range simpleName {
		s[i] = n.String()
	}
	return q.InitWithString(s...)
}

func (q Qualifier) InitWithString(simpleName ...string) Qualifier {
	return q + "(" + Qualifier(strings.Join(simpleName, ", ")) + ")"
}

func (q Qualifier) DotFromString(owner string) Qualifier {
	return Qualifier(owner).Dot(q)
}

func (q Qualifier) DotFrom(owner Qualifier) Qualifier {
	return owner.Dot(q)
}

func (q Qualifier) DotStringer(simpleName fmt.Stringer) Qualifier {
	return q.DotString(simpleName.String())
}

func (q Qualifier) DotString(simpleName string) Qualifier {
	if q == "" {
		return Qualifier(simpleName)
	}
	return q + "." + Qualifier(simpleName)
}

func (q Qualifier) Dot(simpleName Qualifier) Qualifier {
	if q == "" {
		return simpleName
	}
	return q + "." + simpleName
}

func (q Qualifier) IndexStringer(simpleName fmt.Stringer) Qualifier {
	return q.IndexString(simpleName.String())
}

func (q Qualifier) IndexString(simpleName string) Qualifier {
	if q == "" {
		return Qualifier(simpleName)
	}
	return q + "[" + Qualifier(simpleName) + "]"
}

func (q Qualifier) Index(simpleName Qualifier) Qualifier {
	if q == "" {
		return simpleName
	}
	return q + "[" + simpleName + "]"
}

func (q Qualifier) PayloadSuffix() Qualifier { return q + "Payload" }

// any_kind_of_string
func (q Qualifier) ToSnake() Qualifier {
	return Qualifier(strcase.ToSnake(q.String()))
}

// ANY_KIND_OF_STRING
func (q Qualifier) ToScreamingSnake() Qualifier {
	return Qualifier(strcase.ToScreamingSnake(q.String()))
}

// any-kind-of-string
func (q Qualifier) ToKebab() Qualifier {
	return Qualifier(strcase.ToKebab(q.String()))
}

// ANY-KIND-OF-STRING
func (q Qualifier) ToScreamingKebab() Qualifier {
	return Qualifier(strcase.ToScreamingKebab(q.String()))
}

// any.kind.of.string
func (q Qualifier) ToDelimited(del uint8) Qualifier {
	return Qualifier(strcase.ToDelimited(q.String(), del))
}

// ANY.KIND.OF.STRING
func (q Qualifier) ToScreamingDelimited(del uint8) Qualifier {
	return Qualifier(strcase.ToScreamingDelimited(q.String(), del, 0, true))
}

// AnyKindOfString
func (q Qualifier) ToCamel() Qualifier {
	return Qualifier(strcase.ToCamel(strings.ReplaceAll(string(q), ".", "_")))
}

// anyKindOfString
func (q Qualifier) ToLowerCamel() Qualifier {
	return Qualifier(strcase.ToLowerCamel(strings.ReplaceAll(string(q), ".", "_")))
}
