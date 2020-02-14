package pgvt

const enumConstTpl = enumTplDef + `
{{- if $r.In }}
	static const Map<{{ .EnumType }}, Null> {{ $kIn }} = {
		{{- range $r.In }}
		{{ $.EnumLiteralValue . }}: null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static const Map<{{ .EnumType }}, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		{{ $.EnumLiteralValue . }}: null,
		{{- end }}
	};
{{- end -}}`

const enumTplDef = `{{ $r := .Pgv.Rules }}
{{- $kConst := constRef . "kConst" }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const enumTpl = enumTplDef + `
{{ if $r.Const }}
	if ({{ .Accessor }} != {{ .EnumLiteralValue $r.GetConst }})
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ .L10nEnumValue $r.GetConst }});
{{ end }}

{{ if $r.GetDefinedOnly }}
	// {{ .EnumType }} instance cannot exceed the definition, may be null
	{{ template "required" . }}
{{ end }}

{{- $enumL10nClass := .EnumFullL10nClass }}

{{- if $r.In }}
	if (!{{ $kIn }}.containsKey({{ .Accessor }})) {
		{{ if $enumL10nClass }}
			final {{ .EnumFieldL10nAccessor }} = {{ $enumL10nClass }}.of({{ .BuildContextAccessor }});
		{{ end }}
		throw {{ .PgdeFile.AsDot "InError" }}({{ .Err4Args }}, {{ .EnumL10nValues $r.GetIn }});
	}
{{ else if $r.NotIn }}
	if ({{ $kNotIn }}.containsKey({{ .Accessor }})) {
		{{ if $enumL10nClass }}
			final {{ .EnumFieldL10nAccessor }} = {{ $enumL10nClass }}.of({{ .BuildContextAccessor }});
		{{ end }}
		throw {{ .PgdeFile.AsDot "InError" }}.not({{ .Err3Args }}, {{ .EnumL10nValues $r.GetNotIn }});
	}
{{ end }}
`
