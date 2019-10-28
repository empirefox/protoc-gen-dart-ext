package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type Validator struct {
	Package              *Package
	RootFilePath         string
	ClassName            string
	BuildContextAccessor string
	InfoAccessor         string
	L10nAccessor         string
	FieldAccessor        string

	FullL10nClass string

	ImportManager      *dart.ImportManager
	MaterialFile       *dart.ImportFile
	FixnumFile         *dart.ImportFile
	EmailValidatorFile *dart.ImportFile
	PgdeFile           *dart.ImportFile
	L10nFile           *dart.ImportFile
}

func (v *Validator) ImportValidatorAs(i *Validator) (string, error) {
	return v.ImportRootFileAs(v.RootFilePath, i.RootFilePath)
}

func (v *Validator) ImportRootFileAs(target string) (string, error) {
	resolvedImport, err := ResolveImport(v.RootFilePath, target)
	if err != nil {
		return "", err
	}
	return v.ImportManager.GetAs(resolvedImport), nil
}

func (v *Validator) FullPgsPbClass(pgsNty pgs.Message) (string, error) {
	nty := v.Package.Group.PgsToMsg[pgsNty]
	return v.FullPbClass(nty)
}

func (v *Validator) FullPgsPbEnum(pgsNty pgs.Enum) (string, error) {
	nty := v.Package.Group.PgsToEnum[pgsNty]
	return v.FullPbEnum(nty)
}

func (v *Validator) FullPbClass(nty *Message) (string, error) {
	as, err := v.ImportRootFileAs(nty.PbRootFilePath)
	if err != nil {
		return "", err
	}
	if as == "" {
		return nty.DartName.String(), nil
	}
	return as + "." + nty.DartName.String(), nil
}

func (v *Validator) FullPbEnum(nty *Enum) (string, error) {
	as, err := v.ImportRootFileAs(nty.PbRootFilePath)
	if err != nil {
		return "", err
	}
	if as == "" {
		return nty.DartName.String(), nil
	}
	return as + "." + nty.DartName.String(), nil
}

type ValidateMessage struct {
	*Message
	Disabled bool
}

func (nty *ValidateMessage) ImportSelfPbAs() (string, error) {
	return nty.Validator().ImportRootFileAs(nty.PbRootFilePath)
}

func (nty *ValidateMessage) FullPbClass() (string, error) {
	return nty.Validator().FullPbClass(nty.Message)
}

type ValidateOneOf struct {
	*OneOf
	Required bool
}

func (nty *ValidateOneOf) Names() dart.OneOfNames {
	return nty.Package.Group.OneOfNames(nty.Pgs)
}

func (nty *ValidateOneOf) FullPbWhichEnum() pgs.Name {
	return nty.Message.Validate.ImportSelfPbAs() + "." + nty.Names().OneofEnumName()
}

type ValidateField struct {
	*Field
	Pgv *shared.RuleContext
}

func (vf *ValidateField) Accessor() string {
	if vf.Pgv.AccessorOverride != "" {
		return vf.Pgv.AccessorOverride
	}
	return vf.Validator().InfoAccessor + ".proto." + vf.DartName.String()
}

func (vf *ValidateField) FullFieldTypeMessage() (string, error) {
	targetMessage := vf.Package.Group.PgsToMsg[vf.Pgs.Type().Embed()]
	typeMessageClass := targetMessage.DartName
	as, err := vf.Validator().ImportValidatorAs(targetMessage.Validator())
	if err != nil {
		return "", nil
	}

	if as != "" {
		as = as + "." + typeMessageClass
	}
	return as, nil
}

type ValidateEnum struct {
	*Enum
}

type ValidateEnumValue struct {
	*EnumValue
}

func (f *ValidateField) DartType() (pgs.Name, error) {
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
		return fmt.Sprintf("Map<%s, %s>", key, value), nil
	}

	if t.IsRepeated() {
		value, err := f.dartTypeForFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("List<%s>", value), nil
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

func (f *ValidateField) dartTypeForFieldType(ft embedOrEnumOrCoreFieldType) (pgs.Name, error) {
	if ft.IsEmbed() {
		return f.Validator().FullPgsPbClass(ft.Embed())
	}
	if ft.IsEnum() {
		return f.Validator().FullPgsPbEnum(ft.Enum())
	}
	// it will never return dynamic
	return f.dartTypeForCoreProtoType(ft.ProtoType()), nil
}

// only can be used in constRef
func (f *ValidateField) DartConstRefDefineType() (pgs.Name, error) {
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
				return pgs.Name(f.Validator().FixnumFile.As + ".Int64"), nil
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

func (f *Field) dartTypeForCoreProtoType(t pgs.ProtoType) pgs.Name {
	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		// import 'package:fixnum/fixnum.dart' show Int64;
		return pgs.Name(f.Validator().FixnumFile.As + ".Int64")
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
