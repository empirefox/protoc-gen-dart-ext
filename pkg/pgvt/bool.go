package pgvt

const boolTpl = `{{ $r := .Pgv.Rules -}}
{{- if $r.Const }}
	if ({{ if $r.GetConst }}!{{ end }}{{ .Accessor }})
		throw {{ .PgdeFile.As }}.BoolError({{ .Err3Args }}, {{ .BoolL10nValue $r.GetConst }});
{{- end }}
`