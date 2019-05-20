package pgvt

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

type Context struct {
	shared.RuleContext
	dart *dart.Dart
	vr   *dart.ValidatorResolver
	// add imports resolver to generate validate l10n
}

func (ctx *Context) L10nEnumValue(v int32) string {
	return "_l10n."
}
func (ctx *Context) L10nEnumValues(vs []int32) string {
	return "const [_l10n., _l10n.]"
}
func (ctx *Context) L10nField() string {
	return "_l10n."
}
func (ctx *Context) L10nBoolField(kConst bool) string {
	return "_l10n."
}

func (ctx *Context) MessageL10n() (string, error) {
	ctx.Field.FullyQualifiedName()
	return "$i0.XXX()"
}

func (ctx *Context) MessageValidator() (string, error) {
	ctx.Field.FullyQualifiedName()
	return "$v0.XXX()"
}

func (ctx *Context) Number() int32 { return ctx.Field.Descriptor().GetNumber() }

func (ctx *Context) Accessor() string {
	if ctx.AccessorOverride != "" {
		return ctx.AccessorOverride
	}
	return fieldAccessor(ctx.Field)
}

func (ctx *Context) IfHasBegin() string {
	if ctx.AccessorOverride != "" {
		return ""
	}
	return fmt.Sprintf("if (proto.has%s()) {", ctx.Field.Name().UpperCamelCase())
}

func (ctx *Context) IfHasEnd() string {
	if ctx.AccessorOverride != "" {
		return ""
	}
	return "}"
}

func (ctx *Context) IsNull() string    { return ctx.Accessor() + " == null" }
func (ctx *Context) IsNotNull() string { return ctx.Accessor() + " != null" }

func fieldAccessor(f pgs.Field) string { return "proto." + fieldName(f) }
func fieldName(f pgs.Field) string     { return f.Name().LowerCamelCase().String() }

func (ctx *Context) EnumLiteral(v int32) (pgs.Name, error) {
	ev, err := ctx.parseEnumValue(v)
	if err != nil {
		return "", err
	}
	return ctx.dart.NameOf(ev), nil
}

func (ctx *Context) parseEnumValue(v int32) (pgs.EnumValue, error) {
	for _, ev := range ctx.Field.Type().Enum().Values() {
		if ev.Value() == v {
			return ev, nil
		}
	}
	return nil, fmt.Errorf("%s value not found: %d", ctx.Field.Type().Enum().FullyQualifiedName(), v)
}

// only can be used in constRef
func (ctx *Context) FieldType() pgs.Name {
	t := ctx.Field.Type()

	// Map key and value types
	if t.IsMap() {
		switch ctx.AccessorOverride {
		case "key":
			return ctx.typeForProtoType(t.Key().ProtoType())
		case "value":
			return ctx.typeForProtoType(t.Element().ProtoType())
		}
	}

	if t.IsRepeated() {
		if t.ProtoType() == pgs.MessageT {
			return ctx.dart.NameOf(t.Element().Embed())
		} else if t.ProtoType() == pgs.EnumT {
			return ctx.dart.NameOf(t.Element().Enum())
		}
		return ctx.typeForProtoType(t.Element().ProtoType())
	}

	return ctx.typeForProtoType(t.ProtoType())
}

func (ctx *Context) typeForProtoType(t pgs.ProtoType) pgs.Name {
	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32,
		pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return "int"
	case pgs.DoubleT, pgs.FloatT:
		return "double"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "String"
	case pgs.BytesT:
		return "List<int>"
	}

	if t.IsEmbed() {
		if m := t.Embed(); m.IsWellKnown() {
			switch m.WellKnownType() {
			case pgs.AnyWKT:
				return "String"
			case pgs.DurationWKT:
				return "Duration"
			case pgs.TimestampWKT:
				return "DateTime"
			case pgs.Int32ValueWKT, pgs.UInt32ValueWKT, pgs.Int64ValueWKT, pgs.UInt64ValueWKT:
				return "int"
			case pgs.DoubleValueWKT, pgs.FloatValueWKT:
				return "double"
			}
		}
	}

	if t.IsEnum() {
		return ctx.dart.NameOf(t.Enum())
	}

	return "dynamic"
}
