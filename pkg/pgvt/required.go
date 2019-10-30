package pgvt

const requiredTpl = `
{{ if .Rules.GetRequired }}
	if ({{ .Accessor }} == null) throw {{ .PgdeFile.As }}.RequiredError({{ .Err3Args }});
{{ end }}
`
