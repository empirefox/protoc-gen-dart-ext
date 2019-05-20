package pgvt

const enumConstTpl = enumTplDef + `
{{- if $r.In }}
	static const Map<{{ .FieldType }}, Null> {{ $kIn }} = {
		{{- range $r.In }}
		{{ $.EnumLiteral . }}: null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static const Map<{{ .FieldType }}, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		{{ $.EnumLiteral . }}: null,
		{{- end }}
	};
{{- end -}}`

const enumTplDef = `{{ $r := .Rules }}
{{- $kConst := constRef . "kConst" }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const enumTpl = `
{{ if $r.Const }}
	if (_v != {{ .EnumLiteral $r.GetConst }})
		throw ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ .L10nEnumValue $r.GetConst }});
{{ end }}

{{ if $r.GetDefinedOnly }}
	// {{ $fieldType }} instance cannot exceed the definition
	{{ template "required" . }}
{{ end }}

{{- if $r.In }}
	if (!{{ $kIn }}.containsKey(_v))
		throw $pgde.InError(info, {{ $.Number }}, {{ $.L10nField }}, {{ .L10nEnumValues $r.GetIn }});
{{ else if $r.NotIn }}
	if ({{ $kNotIn }}.containsKey(_v))
		throw $pgde.InError.not(info, {{ $.Number }}, {{ $.L10nField }}, {{ .L10nEnumValues $r.GetNotIn }});
{{ end }}
`
