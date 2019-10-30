package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type Validator struct {
	Package               *Package
	RootFilePath          string
	ClassName             dart.Qualifier
	BuildContextAccessor  dart.Qualifier
	InfoAccessor          dart.Qualifier
	L10nAccessor          dart.Qualifier
	FieldAccessor         dart.Qualifier
	EnumFieldL10nAccessor dart.Qualifier

	FullL10nClass dart.Qualifier

	ImportManager      *dart.ImportManager
	DartConvertLib     *dart.ImportFile
	MaterialFile       *dart.ImportFile
	FixnumFile         *dart.ImportFile
	EmailValidatorFile *dart.ImportFile
	PgdeFile           *dart.ImportFile
	L10nFile           *dart.ImportFile
}

func (v *Validator) Int64Type() dart.Qualifier {
	return v.FixnumFile.As.Append("Int64")
}

func (v *Validator) ImportValidatorAs(i *Validator) (dart.Qualifier, error) {
	return v.ImportRootFileAs(i.RootFilePath)
}

func (v *Validator) ImportRootFileAs(target string) (dart.Qualifier, error) {
	resolvedImport, err := ResolveImport(v.RootFilePath, target)
	if err != nil {
		return "", err
	}
	return v.ImportManager.GetAs(resolvedImport), nil
}

func (v *Validator) FullPgsPbClass(pgsNty pgs.Message) (dart.Qualifier, error) {
	nty := v.Package.Group.PgsToMsg[pgsNty]
	return v.FullPbClass(nty)
}

func (v *Validator) FullPgsPbEnum(pgsNty pgs.Enum) (dart.Qualifier, error) {
	nty := v.Package.Group.PgsToEnum[pgsNty]
	return v.FullPbEnum(nty)
}

func (v *Validator) FullPbClass(nty *Message) (dart.Qualifier, error) {
	as, err := v.ImportRootFileAs(nty.PbRootFilePath)
	if err != nil {
		return "", err
	}
	return as.Append(nty.DartName), nil
}

func (v *Validator) FullPbEnum(nty *Enum) (dart.Qualifier, error) {
	as, err := v.ImportRootFileAs(nty.PbRootFilePath)
	if err != nil {
		return "", err
	}
	return as.Append(nty.DartName), nil
}

type ValidateMessage struct {
	*Message
	Disabled bool
}

func (nty *ValidateMessage) ImportSelfPbAs() (dart.Qualifier, error) {
	return nty.Validator().ImportRootFileAs(nty.PbRootFilePath)
}

func (nty *ValidateMessage) FullPbClass() (dart.Qualifier, error) {
	return nty.Validator().FullPbClass(nty.Message)
}

type ValidateOneOf struct {
	*OneOf
	Required bool
}

func (nty *ValidateOneOf) Names() dart.OneOfNames {
	return nty.Package.Group.OneOfNames(nty.Pgs)
}

func (nty *ValidateOneOf) FullPbWhichEnum() (dart.Qualifier, error) {
	as, err := nty.Message.Validate.ImportSelfPbAs()
	if err != nil {
		return "", err
	}
	return as.Append(nty.Names().OneofEnumName()), nil
}

type ValidateField struct {
	*Field
	Pgv *shared.RuleContext
}

func (vf *ValidateField) Key(name, idx string) (*ValidateField, error) {
	outPgv, err := vf.Pgv.Key(name, idx)
	if err != nil {
		return nil, err
	}
	return &ValidateField{
		Field: vf.Field,
		Pgv:   &outPgv,
	}, nil
}

func (vf *ValidateField) Elem(name, idx string) (*ValidateField, error) {
	outPgv, err := vf.Pgv.Elem(name, idx)
	if err != nil {
		return nil, err
	}
	return &ValidateField{
		Field: vf.Field,
		Pgv:   &outPgv,
	}, nil
}

func (vf *ValidateField) Unwrap(name string) (out *ValidateField, err error) {
	outPgv, err := vf.Pgv.Unwrap(name)
	if err != nil {
		return nil, err
	}

	return &ValidateField{
		Field: vf.Field,
		Pgv:   &outPgv,
	}, nil
}

func (vf *ValidateField) BuildContextAccessor() dart.Qualifier {
	return vf.Validator().BuildContextAccessor
}
func (vf *ValidateField) InfoAccessor() dart.Qualifier  { return vf.Validator().InfoAccessor }
func (vf *ValidateField) L10nAccessor() dart.Qualifier  { return vf.Validator().L10nAccessor }
func (vf *ValidateField) FieldAccessor() dart.Qualifier { return vf.Validator().FieldAccessor }
func (vf *ValidateField) EnumFieldL10nAccessor() dart.Qualifier {
	return vf.Validator().EnumFieldL10nAccessor
}
func (vf *ValidateField) Int64Type() dart.Qualifier { return vf.Validator().Int64Type() }

func (vf *ValidateField) Accessor() dart.Qualifier {
	if vf.Pgv.AccessorOverride != "" {
		return dart.Qualifier(vf.Pgv.AccessorOverride)
	}
	return vf.Validator().FieldAccessor
}

func (vf *ValidateField) IfHasValueBegin() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return fmt.Sprintf("if (%s.proto.%s()) {",
		vf.InfoAccessor(),
		vf.Package.Group.FieldNames(vf.Pgs).HasMethodName())
}

func (vf *ValidateField) IfHasValueEnd() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return "}"
}

func (vf *ValidateField) L10nField() dart.Qualifier {
	return vf.Validator().L10nAccessor.Append(vf.Arb.ResourceId())
}

func (vf *ValidateField) ErrCauseL10nField() dart.Qualifier {
	if vf.Pgv.Index != "" {
		return dart.Qualifier(fmt.Sprintf("'`${%s}[${%s}]`'", vf.L10nField(), vf.Pgv.Index))
	}
	return vf.L10nField()
}

func (vf *ValidateField) Err3Args() string {
	return fmt.Sprintf("%s, %d, %s",
		vf.InfoAccessor(),
		vf.Number(),
		vf.ErrCauseL10nField())
}

func (f *ValidateField) DartType() (dart.Qualifier, error) {
	t := f.Pgs.Type()

	if t.IsMap() {
		key, err := f.dartTypeForFieldType(t.Key())
		if err != nil {
			return "", err
		}
		value, err := f.dartTypeForFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return dart.Qualifier(fmt.Sprintf("Map<%s, %s>", key, value)), nil
	}

	if t.IsRepeated() {
		value, err := f.dartTypeForFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return dart.Qualifier(fmt.Sprintf("List<%s>", value)), nil
	}

	return f.dartTypeForFieldType(t)
}

type embedOrEnumOrCoreFieldType interface {
	IsEmbed() bool
	Embed() pgs.Message
	IsEnum() bool
	Enum() pgs.Enum
	ProtoType() pgs.ProtoType
}

func (f *ValidateField) dartTypeForFieldType(ft embedOrEnumOrCoreFieldType) (dart.Qualifier, error) {
	if ft.IsEmbed() {
		return f.Validator().FullPgsPbClass(ft.Embed())
	}
	if ft.IsEnum() {
		return f.Validator().FullPgsPbEnum(ft.Enum())
	}
	// make sure it will never return dynamic
	return f.dartTypeForCoreProtoType(ft.ProtoType()), nil
}

// only can be used in constRef
func (f *ValidateField) DartConstRefDefineType() (dart.Qualifier, error) {
	t := f.Pgs.Type()

	// Map key and value types
	if t.IsMap() {
		switch f.Pgv.AccessorOverride {
		case "key":
			return f.dartTypeForCoreProtoType(t.Key().ProtoType()), nil
		case "value":
			return f.dartTypeForCoreProtoType(t.Element().ProtoType()), nil
		}
	}

	if t.IsEmbed() {
		if m := t.Embed(); m.IsWellKnown() {
			switch m.WellKnownType() {
			case pgs.AnyWKT:
				return "String", nil
			case pgs.DurationWKT:
				return "Duration", nil
			case pgs.TimestampWKT:
				return "DateTime", nil
			case pgs.Int32ValueWKT, pgs.UInt32ValueWKT:
				return "int", nil
			case pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return f.Int64Type(), nil
			case pgs.DoubleValueWKT, pgs.FloatValueWKT:
				return "double", nil
			}
		}
	}

	if t.IsRepeated() {
		switch t.ProtoType() {
		case pgs.MessageT:
			nty := f.Package.Group.PgsToMsg[t.Element().Embed()]
			return f.Validator().FullPbClass(nty)
		case pgs.EnumT:
			nty := f.Package.Group.PgsToEnum[t.Element().Enum()]
			return f.Validator().FullPbEnum(nty)
		}
	}

	if t.IsEnum() {
		nty := f.Package.Group.PgsToEnum[t.Enum()]
		return f.Validator().FullPbEnum(nty)
	}

	return f.dartTypeForCoreProtoType(t.ProtoType()), nil
}

func (f *ValidateField) dartTypeForCoreProtoType(t pgs.ProtoType) dart.Qualifier {
	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return f.Int64Type()
	case pgs.DoubleT, pgs.FloatT:
		return "double"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "String"
	case pgs.BytesT:
		return "List<int>"
	}
	return "dynamic"
}
