package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/l10n"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

func (f *File) addMessage(pgsNty pgs.Message) error {
	validateDisabled, err := shared.Disabled(pgsNty)
	if err != nil {
		return err
	}

	var l10nNty L10nMsgOrEnum
	_, err = pgsNty.Extension(l10n.E_Message, &l10nNty.Extension)
	if err != nil {
		return err
	}

	names := f.Dart.MessageNames(pgsNty)

	nty := &Message{
		Entity: Entity{
			DartName: names.Name(),
			File:     f,
		},
		Pgs:      pgsNty,
		Names:    names,
		L10n:     &l10nNty,
		Validate: &ValidateMessage{Disabled: validateDisabled},
	}

	l10nNty.Message = nty
	nty.Validate.Message = nty
	f.Messages = append(f.Messages, nty)

	pgsToOneOf := make(map[pgs.OneOf]*OneOf, len(pgsNty.OneOfs()))
	for _, child := range pgsNty.OneOfs() {
		oo, err := nty.addOneof(child)
		if err != nil {
			return err
		}
		pgsToOneOf[child] = oo
	}
	for _, child := range pgsNty.Fields() {
		err = nty.addField(pgsToOneOf, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (msg *Message) addField(pgsToOneOf map[pgs.OneOf]*OneOf, pgsNty pgs.Field) error {
	pgvCtx, err := rulesContext(pgsNty)
	if err != nil {
		return err
	}

	var l10nNty L10nField
	_, err = pgsNty.Extension(l10n.E_Field, &l10nNty.Extension)
	if err != nil {
		return err
	}

	if pgsNty.Type().IsRepeated() || pgsNty.Type().IsMap() {
		elem := pgsNty.Type().Element()
		if elem.IsEnum() {
			l10nNty.IsRefEnum = true
			l10nNty.RefEnum = elem.Enum()
		} else if elem.IsEmbed() {
			l10nNty.RefMessage = elem.Embed()
		}
	} else if pgsNty.Type().IsEnum() {
		l10nNty.IsRefEnum = true
		l10nNty.RefEnum = pgsNty.Type().Enum()
	} else if pgsNty.Type().IsEmbed() {
		l10nNty.RefMessage = pgsNty.Type().Embed()
	}

	names := msg.File.Dart.FieldNames(pgsNty)

	nty := &Field{
		Entity: Entity{
			DartName: names.Name(),
			File:     msg.File,
		},
		Pgs:      pgsNty,
		Names:    names,
		Message:  msg,
		L10n:     &l10nNty,
		Validate: &ValidateField{Pgv: &pgvCtx},
	}

	l10nNty.Entity = nty
	nty.Validate.Field = nty
	msg.Fields = append(msg.Fields, nty)

	if pgsNty.InOneOf() {
		oo := pgsToOneOf[pgsNty.OneOf()]
		oo.Fields = append(oo.Fields, nty)
	} else {
		msg.NonOneOfFields = append(msg.NonOneOfFields, nty)
	}

	return nil
}

func (msg *Message) addOneof(pgsNty pgs.OneOf) (*OneOf, error) {
	validateRequired, err := shared.RequiredOneOf(pgsNty)
	if err != nil {
		return nil, err
	}

	var l10nNty L10nOneOf
	_, err = pgsNty.Extension(l10n.E_Oneof, &l10nNty.Extension)
	if err != nil {
		return nil, err
	}

	names := msg.File.Dart.OneOfNames(pgsNty)

	nty := &OneOf{
		Entity: Entity{
			DartName: names.Name(),
			File:     msg.File,
		},
		Pgs:      pgsNty,
		Names:    names,
		Message:  msg,
		L10n:     &l10nNty,
		Validate: &ValidateOneOf{Required: validateRequired},
	}

	l10nNty.Entity = nty
	nty.Validate.OneOf = nty
	msg.OneOfs = append(msg.OneOfs, nty)
	return nty, nil
}

func (f *File) addEnum(pgsNty pgs.Enum) error {
	l10nNty := L10nMsgOrEnum{IsEnum: true}
	_, err := pgsNty.Extension(l10n.E_Enum, &l10nNty.Extension)
	if err != nil {
		return err
	}

	names := f.Dart.EnumNames(pgsNty)

	nty := &Enum{
		Entity: Entity{
			DartName: names.Name(),
			File:     f,
		},
		Pgs:   pgsNty,
		Names: names,
		L10n:  &l10nNty,
	}

	l10nNty.Enum = nty
	f.Enums = append(f.Enums, nty)

	for _, child := range pgsNty.Values() {
		err = nty.addValue(child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (enum *Enum) addValue(pgsNty pgs.EnumValue) error {
	var l10nNty L10nEnumValue
	_, err := pgsNty.Extension(l10n.E_Value, &l10nNty.Extension)
	if err != nil {
		return err
	}

	names := enum.File.Dart.ValueNames(pgsNty)

	nty := &EnumValue{
		Entity: Entity{
			DartName: names.Name(),
			File:     enum.File,
		},
		Pgs:   pgsNty,
		Names: names,
		Enum:  enum,
		L10n:  &l10nNty,
	}

	l10nNty.Entity = nty
	enum.Values = append(enum.Values, nty)
	return nil
}
