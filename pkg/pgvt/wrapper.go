package pgvt

const wrapperConstTpl = `
{{- renderConstants (unwrap .) }}
`

const wrapperTpl = `
{{ .IfHasBegin }}
	{{ render (unwrap .) }}
{{ .IfHasEnd }}

{{- $maybeEmpty := .IfHasBegin -}}

{{ if and $maybeEmpty .MessageRules.GetRequired }} else {
	throw {{ .PgdeFile.As }}.RequiredError({{ .Err3Args }});
} {{ end }}
`
