package dartpb

import (
	"path/filepath"

	pgs "github.com/lyft/protoc-gen-star"
)

func ResolveImport(source, target string) (string, error) {
	if source == target {
		return "", nil
	}
	return filepath.Rel(filepath.Dir(source), target)
}

type Entity struct {
	DartName pgs.Name
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

	Validate *ValidateEnum
}

type EnumValue struct {
	Entity

	Pgs pgs.EnumValue

	Arb *EnumValueArb

	Enum *Enum

	Validate *ValidateEnumValue
}
