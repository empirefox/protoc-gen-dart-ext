package dartpb

import (
	"fmt"
	"path/filepath"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
)

func ResolveImport(source, target string) (string, error) {
	if source == target {
		return "", nil
	}
	return filepath.Rel(filepath.Dir(source), target)
}

type Entity struct {
	DartName dart.Qualifier
	Package  *Package
}

func (nty *Entity) L10n() *L10n           { return nty.Package.L10n }
func (nty *Entity) Validator() *Validator { return nty.Package.Validator }

type Message struct {
	Entity

	Pgs pgs.Message

	Arb *MsgOrEnumArb

	PbRootFilePath string

	Fields []*Field

	NonOneOfFields []*Field

	OneOfs []*OneOf

	Validate *ValidateMessage
}

type Field struct {
	Entity

	Pgs pgs.Field

	Arb *FieldArb

	Message *Message

	Validate *ValidateField
}

type OneOf struct {
	Entity

	Pgs pgs.OneOf

	Arb *OneOfArb

	Message *Message

	Fields []*Field

	Validate *ValidateOneOf
}

type Enum struct {
	Entity

	Pgs pgs.Enum

	Arb *MsgOrEnumArb

	PbRootFilePath string

	Values []*EnumValue
}

type EnumValue struct {
	Entity

	Pgs pgs.EnumValue

	Arb *EnumValueArb

	Enum *Enum
}

func (f *Field) Number() int32 { return f.Pgs.Descriptor().GetNumber() }

func (f *Field) parseEnumValue(v int32) (*EnumValue, error) {
	for _, ev := range f.Pgs.Type().Enum().Values() {
		if ev.Value() == v {
			return f.Package.Group.PgsToValue[ev], nil
		}
	}
	return nil, fmt.Errorf("%s value not found: %d", f.Pgs.Type().Enum().FullyQualifiedName(), v)
}
