package pgvt

const boolTpl = `{{ $r := .Rules -}}
{{- if $r.Const }}
	if ({{ if $r.GetConst }}!{{ end }}_v)
		throw $pgde.BoolError(info, {{ $.Number }}, {{ $.L10nField }}, {{ $.L10nBoolField $r.GetConst }});
{{- end }}
`