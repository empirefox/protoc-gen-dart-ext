package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
)

type Entity struct {
	DartName dart.Qualifier
	File     *File
}

func (nty *Entity) Translator() *Translator { return nty.File.Translator }
func (nty *Entity) Validators() *Validators { return nty.File.Validators }

type Message struct {
	Entity

	Pgs pgs.Message

	Names dart.MessageNames

	L10n *L10nMsgOrEnum

	Fields []*Field

	NonOneOfFields []*Field

	OneOfs []*OneOf

	Validate *ValidateMessage
}

type Field struct {
	Entity

	Pgs pgs.Field

	Names dart.FieldNames

	Message *Message

	L10n *L10nField

	Validate *ValidateField
}

type OneOf struct {
	Entity

	Pgs pgs.OneOf

	Names dart.OneOfNames

	Message *Message

	L10n *L10nOneOf

	Fields []*Field

	Validate *ValidateOneOf
}

type Enum struct {
	Entity

	Pgs pgs.Enum

	Names dart.EnumNames

	L10n *L10nMsgOrEnum

	Values []*EnumValue
}

type EnumValue struct {
	Entity

	Pgs pgs.EnumValue

	Names dart.ValueNames

	Enum *Enum

	L10n *L10nEnumValue
}

func (f *Field) Number() int32 { return f.Pgs.Descriptor().GetNumber() }

func (f *Field) parseEnumValue(v int32) (pgs.EnumValue, error) {
	for _, ev := range f.Pgs.Type().Enum().Values() {
		if ev.Value() == v {
			return ev, nil
		}
	}
	return nil, fmt.Errorf("%s value not found: %d", f.Pgs.Type().Enum().FullyQualifiedName(), v)
}
