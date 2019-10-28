package pgvt

const requiredTpl = `
{{ if .Rules.GetRequired }}
	if ({{ .Accessor }} == null) throw {{ .PgdeFile.As }}.RequiredError(info, {{ $.Number }}, {{ $.L10nField }});
{{ end }}
`
