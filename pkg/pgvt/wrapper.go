package pgvt

const wrapperConstTpl = `
{{- renderConstants (unwrap .) }}
`

const wrapperTpl = `
{{ .IfHasBegin }}
	{{ render (unwrap .) }}
{{ .IfHasEnd }}

{{- $maybeEmpty := .IfHasBegin -}}

{{ if and $maybeEmpty .Pgv.MessageRules.GetRequired }} else {
	throw {{ .PgdeFile.AsDot "RequiredError" }}({{ .Err3Args }});
} {{ end }}
`
