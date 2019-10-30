package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/l10n"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

func (p *Package) addMessage(g *PackageGroup, pgsNty pgs.Message) error {
	validateDisabled, err := shared.Disabled(pgsNty)
	if err != nil {
		return err
	}

	var a MsgOrEnumArb
	_, err = pgsNty.Extension(l10n.E_MsgArb, &a.Extension)
	if err != nil {
		return err
	}

	nty := &Message{
		Entity: Entity{
			DartName: g.NameOf(pgsNty),
			Package:  p,
		},
		Pgs:            pgsNty,
		Arb:            &a,
		PbRootFilePath: pgsNty.File().InputPath().SetExt(".pb.dart").String(),

		Validate: &ValidateMessage{Disabled: validateDisabled},
	}

	g.PgsToMsg[pgsNty] = nty
	a.Message = nty
	nty.Validate.Message = nty
	p.Messages = append(p.Messages, nty)

	for _, child := range pgsNty.OneOfs() {
		err = nty.addOneof(g, child)
		if err != nil {
			return err
		}
	}
	for _, child := range pgsNty.Fields() {
		err = nty.addField(g, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (msg *Message) addField(g *PackageGroup, pgsNty pgs.Field) error {
	pgvCtx, err := rulesContext(pgsNty)
	if err != nil {
		return err
	}

	var a FieldArb
	_, err = pgsNty.Extension(l10n.E_FieldArb, &a.Extension)
	if err != nil {
		return err
	}

	nty := &Field{
		Entity: Entity{
			DartName: g.NameOf(pgsNty),
			Package:  msg.Package,
		},
		Pgs:     pgsNty,
		Arb:     &a,
		Message: msg,

		Validate: &ValidateField{Pgv: &pgvCtx},
	}

	g.PgsToField[pgsNty] = nty
	a.Entity = nty
	nty.Validate.Field = nty
	msg.Fields = append(msg.Fields, nty)
	return nil
}

func (msg *Message) addOneof(g *PackageGroup, pgsNty pgs.OneOf) error {
	validateRequired, err := shared.RequiredOneOf(pgsNty)
	if err != nil {
		return err
	}

	var a OneOfArb
	_, err = pgsNty.Extension(l10n.E_OneofArb, &a.Extension)
	if err != nil {
		return err
	}

	nty := &OneOf{
		Entity: Entity{
			DartName: g.NameOf(pgsNty),
			Package:  msg.Package,
		},
		Pgs:     pgsNty,
		Arb:     &a,
		Message: msg,

		Validate: &ValidateOneOf{Required: validateRequired},
	}

	g.PgsToOneOf[pgsNty] = nty
	a.Entity = nty
	nty.Validate.OneOf = nty
	msg.OneOfs = append(msg.OneOfs, nty)
	return nil
}

func (p *Package) addEnum(g *PackageGroup, pgsNty pgs.Enum) error {
	a := MsgOrEnumArb{
		IsEnum: true,
	}
	_, err := pgsNty.Extension(l10n.E_EnumArb, &a.Extension)
	if err != nil {
		return err
	}

	nty := &Enum{
		Entity: Entity{
			DartName: g.NameOf(pgsNty),
			Package:  p,
		},
		Pgs:            pgsNty,
		Arb:            &a,
		PbRootFilePath: pgsNty.File().InputPath().SetExt(".pb.dart").String(),
	}

	g.PgsToEnum[pgsNty] = nty
	a.Enum = nty
	p.Enums = append(p.Enums, nty)

	for _, child := range pgsNty.Values() {
		err = nty.addValue(g, child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (enum *Enum) addValue(g *PackageGroup, pgsNty pgs.EnumValue) error {
	var a EnumValueArb
	_, err := pgsNty.Extension(l10n.E_ValueArb, &a.Extension)
	if err != nil {
		return err
	}

	nty := &EnumValue{
		Entity: Entity{
			DartName: g.NameOf(pgsNty),
			Package:  enum.Package,
		},
		Pgs:  pgsNty,
		Arb:  &a,
		Enum: enum,
	}

	g.PgsToValue[pgsNty] = nty
	a.Entity = nty
	enum.Values = append(enum.Values, nty)
	return nil
}
