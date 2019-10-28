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
