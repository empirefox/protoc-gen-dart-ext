package pgvt

const requiredTpl = `
{{ if .Rules.GetRequired }}
	if ({{ .Accessor }} == null) throw {{ .PgdeFile.AsDot "RequiredError" }}({{ .Err3Args }});
{{ end }}
`
