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

func (nty *Entity) Zeros() *Zeros           { return nty.File.Zeros }
func (nty *Entity) Translator() *Translator { return nty.File.Translator }
func (nty *Entity) Validators() *Validators { return nty.File.Validators }

type Message struct {
	Entity

	Pgs pgs.Message

	Names dart.MessageNames

	Zero *ZeroMessage

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

	Zero *ZeroField

	L10n *L10nField

	Format *FormatField

	Validate *ValidateField
}

type OneOf struct {
	Entity

	Pgs pgs.OneOf

	Names dart.OneOfNames

	Message *Message

	Zero *ZeroOneOf

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

func (f *Field) ElementOrFieldProtoType() pgs.ProtoType {
	elem := f.Pgs.Type().Element()
	if elem != nil {
		return elem.ProtoType()
	}
	return f.Pgs.Type().ProtoType()
}

func (f *Field) ElementOrEmbed() pgs.Message {
	pgsType := f.Pgs.Type()
	target := pgsType.Embed()
	if target == nil {
		target = pgsType.Element().Embed()
	}
	return target
}

func (f *Field) ElementOrEnum() pgs.Enum {
	pgsType := f.Pgs.Type()
	target := pgsType.Enum()
	if target == nil {
		target = pgsType.Element().Enum()
	}
	return target
}

func (f *Field) Number() int32 { return f.Pgs.Descriptor().GetNumber() }

func (f *Field) parseEnumValue(v int32) (pgs.EnumValue, error) {
	for _, ev := range f.ElementOrEnum().Values() {
		if ev.Value() == v {
			return ev, nil
		}
	}
	return nil, fmt.Errorf("%s value not found: %d", f.Pgs.Type().Enum().FullyQualifiedName(), v)
}
