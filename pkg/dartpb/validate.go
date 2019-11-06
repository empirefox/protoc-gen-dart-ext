package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type Validators struct {
	File                  *File
	BuildContextAccessor  dart.Qualifier
	InfoAccessor          dart.Qualifier
	L10nAccessor          dart.Qualifier
	FieldAccessor         dart.Qualifier
	EnumFieldL10nAccessor dart.Qualifier

	ImportManager      *dart.ImportManager
	DartConvertLib     *dart.ImportFile
	FoundationFile     *dart.ImportFile
	MaterialFile       *dart.ImportFile
	FixnumFile         *dart.ImportFile
	EmailValidatorFile *dart.ImportFile
	PgdeFile           *dart.ImportFile
	PbFile             *dart.ImportFile
	TranslatorFile     *dart.ImportFile
}

func (v *Validators) fullPbClass(pgsNty pgs.Entity) (dart.Qualifier, error) {
	return v.ImportManager.PbFileDot(pgsNty, v.File.Dart.NameOf(pgsNty))
}

func (v *Validators) FullPbClass(pgsNty pgs.Message) (dart.Qualifier, error) {
	return v.fullPbClass(pgsNty)
}

func (v *Validators) FullPbEnum(pgsNty pgs.Enum) (dart.Qualifier, error) {
	return v.fullPbClass(pgsNty)
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
	return vf.Validators().BuildContextAccessor
}
func (vf *ValidateField) InfoAccessor() dart.Qualifier { return vf.Validators().InfoAccessor }
func (vf *ValidateField) L10nAccessor() dart.Qualifier {
	return vf.Validators().L10nAccessor
}
func (vf *ValidateField) FieldAccessor() dart.Qualifier { return vf.Validators().FieldAccessor }
func (vf *ValidateField) EnumFieldL10nAccessor() dart.Qualifier {
	return vf.Validators().EnumFieldL10nAccessor
}
func (vf *ValidateField) Int64Type() dart.Qualifier {
	return vf.Validators().FixnumFile.AsDot("Int64")
}

func (vf *ValidateField) Accessor() dart.Qualifier {
	if vf.Pgv.AccessorOverride != "" {
		return dart.Qualifier(vf.Pgv.AccessorOverride)
	}
	return vf.Validators().FieldAccessor
}

func (vf *ValidateField) IfHasValueBegin() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return fmt.Sprintf("if (%s.proto.%s()) {",
		vf.InfoAccessor(),
		vf.File.Dart.FieldNames(vf.Pgs).HasMethodName())
}

func (vf *ValidateField) IfHasValueEnd() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return "}"
}

func (vf *ValidateField) L10nField() dart.Qualifier {
	return vf.Validators().L10nAccessor.Dot(vf.L10n.ResourceId())
}

func (vf *ValidateField) ErrCauseTranslatorField() dart.Qualifier {
	if vf.Pgv.Index != "" {
		return dart.Qualifier(fmt.Sprintf("'`${%s}[${%s}]`'", vf.L10nField(), vf.Pgv.Index))
	}
	return vf.L10nField()
}

func (vf *ValidateField) Err3Args() string {
	return fmt.Sprintf("%s, %d, %s",
		vf.InfoAccessor(),
		vf.Number(),
		vf.ErrCauseTranslatorField())
}

func (vf *ValidateField) DartType() (dart.Qualifier, error) {
	t := vf.Pgs.Type()

	if t.IsMap() {
		key, err := vf.dartTypeForFieldType(t.Key())
		if err != nil {
			return "", err
		}
		value, err := vf.dartTypeForFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return dart.Qualifier(fmt.Sprintf("Map<%s, %s>", key, value)), nil
	}

	if t.IsRepeated() {
		value, err := vf.dartTypeForFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return dart.Qualifier(fmt.Sprintf("List<%s>", value)), nil
	}

	return vf.dartTypeForFieldType(t)
}

type embedOrEnumOrCoreFieldType interface {
	IsEmbed() bool
	Embed() pgs.Message
	IsEnum() bool
	Enum() pgs.Enum
	ProtoType() pgs.ProtoType
}

func (vf *ValidateField) dartTypeForFieldType(ft embedOrEnumOrCoreFieldType) (dart.Qualifier, error) {
	if ft.IsEmbed() {
		return vf.Validators().FullPbClass(ft.Embed())
	}
	if ft.IsEnum() {
		return vf.Validators().FullPbEnum(ft.Enum())
	}
	// make sure it will never return dynamic
	return vf.dartTypeForCoreProtoType(ft.ProtoType()), nil
}

// only can be used in constRef
func (vf *ValidateField) DartConstRefDefineType() (dart.Qualifier, error) {
	t := vf.Pgs.Type()

	// Map key and value types
	if t.IsMap() {
		switch vf.Pgv.AccessorOverride {
		case "key":
			return vf.dartTypeForCoreProtoType(t.Key().ProtoType()), nil
		case "value":
			return vf.dartTypeForCoreProtoType(t.Element().ProtoType()), nil
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
				return vf.Int64Type(), nil
			case pgs.DoubleValueWKT, pgs.FloatValueWKT:
				return "double", nil
			}
		}
	}

	if t.IsRepeated() {
		switch t.ProtoType() {
		case pgs.MessageT:
			return vf.Validators().FullPbClass(t.Element().Embed())
		case pgs.EnumT:
			return vf.Validators().FullPbEnum(t.Element().Enum())
		}
	}

	if t.IsEnum() {
		return vf.Validators().FullPbEnum(t.Enum())
	}

	return vf.dartTypeForCoreProtoType(t.ProtoType()), nil
}

func (vf *ValidateField) dartTypeForCoreProtoType(t pgs.ProtoType) dart.Qualifier {
	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return vf.Int64Type()
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
