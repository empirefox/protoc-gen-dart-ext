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
	emailValidatorType dart.Qualifier
	pbFile             *dart.ImportFile
	l10nType           dart.Qualifier
}

func (s *Validators) EmailValidatorType() dart.Qualifier {
	if s.emailValidatorType == "" {
		s.emailValidatorType = s.ImportManager.
			Import("package:email_validator/email_validator.dart").
			AsDot("EmailValidator")
	}
	return s.emailValidatorType
}
func (s *Validators) PbFile() *dart.ImportFile {
	if s.pbFile == nil {
		s.pbFile = s.ImportManager.ImportPbFile(s.File.Pgs)
	}
	return s.pbFile
}
func (s *Validators) L10nType() dart.Qualifier {
	if s.l10nType == "" {
		s.l10nType = s.ImportManager.ImportL10nFile(s.File.Pgs).
			AsDot(s.File.Names.L10nName())
	}
	return s.l10nType
}

func (s *Validators) FullPbClass(pgsNty pgs.Message) (dart.Qualifier, error) {
	return s.ImportManager.FullPbClassOrEnum(pgsNty)
}

func (s *Validators) FullPbEnum(pgsNty pgs.Enum) (dart.Qualifier, error) {
	return s.ImportManager.FullPbClassOrEnum(pgsNty)
}

type ValidateField struct {
	*Field
	dart.ImportManagerCommonFiles
	Pgv *shared.RuleContext
}

func (im *ValidateField) EmailValidatorType() dart.Qualifier {
	return im.Validators().EmailValidatorType()
}
func (im *ValidateField) PbFile() *dart.ImportFile { return im.Validators().PbFile() }

func (vf *ValidateField) TemplateName() string      { return vf.Pgv.Typ }
func (vf *ValidateField) ConstTemplateName() string { return vf.Pgv.Typ + "Const" }

func (vf *ValidateField) Key(name, idx string) (*ValidateField, error) {
	return vf.withPgv(vf.Pgv.Key(name, idx))
}
func (vf *ValidateField) Elem(name, idx string) (*ValidateField, error) {
	return vf.withPgv(vf.Pgv.Elem(name, idx))
}
func (vf *ValidateField) Unwrap(name string) (out *ValidateField, err error) {
	return vf.withPgv(vf.Pgv.Unwrap(name))
}
func (vf *ValidateField) withPgv(outPgv shared.RuleContext, err error) (*ValidateField, error) {
	if err != nil {
		return nil, err
	}
	out := *vf
	out.Pgv = &outPgv
	return &out, nil
}

func (vf *ValidateField) IfHasBegin() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return fmt.Sprintf("if (%s.proto.%s()) {",
		vf.Validators().InfoAccessor,
		vf.Names.HasMethodName())
}

func (vf *ValidateField) IfHasEnd() string {
	if vf.Pgv.AccessorOverride != "" {
		return ""
	}
	return "}"
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

func (vf *ValidateField) Err4Args() (string, error) {
	kfmt, err := vf.Format.Render(vf.ImportManager())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s, %s", vf.Err3Args(), kfmt), nil
}

func (vf *ValidateField) DartType() (dart.Qualifier, error) {
	return vf.ImportManager().TypeForFieldType(vf.Pgs.Type())
}

// only can be used in constRef
func (vf *ValidateField) DartConstRefDefineType() (dart.Qualifier, error) {
	t := vf.Pgs.Type()
	m := vf.ImportManager()

	// Map key and value types
	if t.IsMap() {
		switch vf.Pgv.AccessorOverride {
		case "key":
			return m.TypeForCoreFieldType(t.Key().ProtoType()), nil
		case "value":
			return m.TypeForCoreFieldType(t.Element().ProtoType()), nil
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

	return m.TypeForCoreFieldType(t.ProtoType()), nil
}
