package pgvt

const inConstTpl = inTplDef + `
{{- if $r.In }}
	static const Map<{{ .FieldType }}, Null> {{ $kIn }} = {
		{{- range $r.In }}
		{{ lit . }}: null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static const Map<{{ .FieldType }}, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		{{ lit . }}: null,
		{{- end }}
	};
{{- end -}}`

const inTplDef = `{{ $r := .Rules }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const inTpl = inTplDef + `
{{- if $r.In }}
	if (!{{ $kIn }}.containsKey(_v))
		throw $pgde.InError({{ $kIn }}.keys);
{{- end -}}
{{- if $r.NotIn }}
	if ({{ $kNotIn }}.containsKey(_v))
		throw $pgde.InError.not({{ $kNotIn }}.keys);
{{- end -}}
`
