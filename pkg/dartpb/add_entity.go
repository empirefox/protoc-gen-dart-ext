package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/format"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/l10n"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/zero"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

func (f *File) addMessage(pgsNty pgs.Message) error {
	rpcNty := new(RPCMessage)
	isRPC, err := pgsNty.Extension(form.E_Node, &rpcNty.Extension)
	if err != nil {
		return err
	}
	if !isRPC {
		rpcNty = nil
	}

	zeroDisabled, err := ZeroDisabled(pgsNty)
	if err != nil {
		return err
	}

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
		Pgs:   pgsNty,
		Names: names,
		RPC:   rpcNty,
		Zero: &ZeroMessage{
			ImportManagerCommonFiles: f.Zeros.ImportManager,
			Disabled:                 zeroDisabled,
		},
		L10n: &l10nNty,
		Validate: &ValidateMessage{
			ImportManagerCommonFiles: f.Validators.ImportManager,
			Disabled:                 validateDisabled,
		},
	}

	if isRPC {
		rpcNty.Message = nty
		if rpcNty.IsLeaf() {
		}
		if rpcNty.IsView() {
			rpcNty.ViewGroup, _ = f.RPCS.findOneof(pgsNty, "group")
			if rpcNty.ViewGroup != nil {
				hasElement, err := f.RPCS.HasElement(pgsNty)
				if err != nil {
					return err
				}
				if hasElement {
					rpcNty.IdMessage = f.RPCS.findEmbedId(pgsNty)
					if rpcNty.IdMessage == nil {
						return fmt.Errorf("Group view with Element must have Id: %s",
							pgsNty.FullyQualifiedName())
					}
				}
			}
		}
	}
	nty.Zero.Message = nty
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

func (msg *Message) addField(pgsToOneOf map[pgs.OneOf]*OneOf, pgsNty pgs.Field) (err error) {
	rpcNty := new(RPCField)
	var isRPC bool
	if msg.RPC != nil {
		if msg.RPC.IsEntry() {
			isRPC, err = pgsNty.Extension(form.E_EntryField, &rpcNty.EntryExt)
			if err != nil {
				return err
			}
			if rpcNty.EntryExt.Is == nil {
				isRPC, err = msg.File.RPCS.checkFieldType(pgsNty, form.Node_typeName)
			}
		} else if msg.RPC.IsLeaf() {
			isRPC, err = pgsNty.Extension(form.E_Field, &rpcNty.LeafExt)
		}
		if err != nil {
			return err
		}
	}
	if !isRPC {
		rpcNty = nil
	}

	if (msg.RPC.IsLeaf() || msg.RPC.IsView()) && msg.RPC.IdField == nil {
		ok, err := msg.File.RPCS.IsId(pgsNty)
		if err != nil {
			return err
		}
		if ok {
			msg.RPC.IdField = pgsNty
		}
	}

	zeroNty := ZeroField{ImportManagerCommonFiles: msg.File.Zeros.ImportManager}
	_, err = pgsNty.Extension(zero.E_To, &zeroNty.Extension)
	if err != nil {
		return err
	}

	pgvCtx, err := rulesContext(pgsNty)
	if err != nil {
		return err
	}

	var l10nNty L10nField
	_, err = pgsNty.Extension(l10n.E_Field, &l10nNty.Extension)
	if err != nil {
		return err
	}

	var fmtNty FormatField
	_, err = pgsNty.Extension(format.E_To, &fmtNty.Extension)
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

	if l10nNty.RefMessage != nil {
		l10nNty.IsRefWKT = l10nNty.RefMessage.IsWellKnown()
	}

	names := msg.File.Dart.FieldNames(pgsNty)

	nty := &Field{
		Entity: Entity{
			DartName: names.Name(),
			File:     msg.File,
		},
		Pgs:     pgsNty,
		Names:   names,
		Message: msg,
		RPC:     rpcNty,
		Zero:    &zeroNty,
		L10n:    &l10nNty,
		Format:  &fmtNty,
		Validate: &ValidateField{
			ImportManagerCommonFiles: msg.File.Validators.ImportManager,
			Pgv:                      &pgvCtx,
		},
	}

	if isRPC {
		rpcNty.Field = nty
	}
	zeroNty.Field = nty
	l10nNty.Entity = nty
	fmtNty.Field = nty
	nty.Validate.Field = nty
	msg.Fields = append(msg.Fields, nty)

	if pgsNty.InOneOf() {
		oo := pgsToOneOf[pgsNty.OneOf()]
		oo.Fields = append(oo.Fields, nty)
		if pgsNty.Descriptor().GetNumber() == oo.Zero.defaultNumber {
			oo.Zero.Default = nty.Zero
		}
	} else {
		msg.NonOneOfFields = append(msg.NonOneOfFields, nty)
	}

	return nil
}

func (msg *Message) addOneof(pgsNty pgs.OneOf) (*OneOf, error) {
	var zeroNty ZeroOneOf
	_, err := pgsNty.Extension(zero.E_Default, &zeroNty.defaultNumber)
	if err != nil {
		return nil, err
	}

	if zeroNty.defaultNumber != 0 {
		var found bool
		for _, fnty := range pgsNty.Fields() {
			if fnty.Descriptor().GetNumber() == zeroNty.defaultNumber {
				found = true
			}
		}
		if !found {
			return nil, fmt.Errorf("default not found for oneof: %s",
				pgsNty.FullyQualifiedName())
		}
	}

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
		Pgs:     pgsNty,
		Names:   names,
		Message: msg,
		Zero:    &zeroNty,
		L10n:    &l10nNty,
		Validate: &ValidateOneOf{
			ImportManagerCommonFiles: msg.File.Validators.ImportManager,
			Required:                 validateRequired,
		},
	}

	zeroNty.OneOf = nty
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

	_, err = pgsNty.Extension(l10n.E_Error, &l10nNty.IsError)
	if err != nil {
		return err
	}

	names := f.Dart.EnumNames(pgsNty)

	if l10nNty.IsError {
		if names.Name() != "ErrorCode" {
			return fmt.Errorf("Enum name must be `ErrorCode`, when pgde.l10n.error=true: %s",
				pgsNty.FullyQualifiedName())
		}
		f.isL10nErrorCoder = true
	}

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
