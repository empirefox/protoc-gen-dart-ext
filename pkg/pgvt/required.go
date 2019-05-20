package pgvt

const requiredTpl = `
{{ if .Rules.GetRequired }}
	if ({{ .Accessor }} == null) throw $pgde.RequiredError(info);
{{ end }}
`
