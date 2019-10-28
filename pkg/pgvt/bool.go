package pgvt

const boolTpl = `{{ $r := .Rules -}}
{{- if $r.Const }}
	if ({{ if $r.GetConst }}!{{ end }}_v)
		throw {{ .PgdeFile.As }}.BoolError(info, {{ $.Number }}, {{ $.L10nField }}, {{ $.L10nBoolField $r.GetConst }});
{{- end }}
`