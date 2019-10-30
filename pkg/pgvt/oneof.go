package pgvt

const oneOfConstTpl = `
{{ range .Fields }}{{ renderConstants .Validate }}{{ end }}
`

const oneOfTpl = `
switch ({{ .InfoAccessor }}.proto.{{ .Names.WhichOneofMethodName }}()) {
	{{ range .Fields -}}
		case {{ $.FullPbWhichEnum }}.{{ .DartName }}:
			assertField_{{ .DartName }}();
			break;
	{{ end -}}
	{{ if .Required }}
		default: throw {{ .PgdeFile.As }}.OneofRequiredError({{ .InfoAccessor }}, {{ .L10nAccessor }}.{{ .L10n.ResourceId }});
	{{ end }}
}
`
