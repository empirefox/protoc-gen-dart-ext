package pgvt

const msgTpl = `
/// Validates [{{ .FullPbClass }}] protobuf objects.
class {{ .ClassName }} extends {{ .PgdeFile.As }}.GeneratedValidator<{{ .FullPbClass }}> {
	{{ .MaterialFile.As }}.BuildContext {{ .Validator.BuildContextAccessor }};

	{{ .PgdeFile.As }}.ValidateInfo<{{ .FullPbClass }}> {{ .Validator.InfoAccessor }};

	{{ .L10nFile.As }}.{{ .L10n.ClassName }} {{ .Validator.L10nAccessor }};

	{{ .ClassName }}({{ .MaterialFile.As }}.BuildContext context, {{ .PgdeFile.As }}.ValidateInfo<{{ .FullPbClass }}> info)
		: {{ .Validator.BuildContextAccessor }} = context,
		  {{ .Validator.InfoAccessor }} = info,
		  {{ .Validator.L10nAccessor }} = {{ .MaterialFile.As }}.Localizations.of<{{ .Validator.FullL10nClass }}>(context, {{ .Validator.FullL10nClass }});

	{{- range .NonOneOfFields }}
		{{ renderConstants .Validate }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOfConst" .Validate }}
	{{ end }}

	void assertProto() {
		{{ if .Disabled }}
			// Validate is disabled for {{ .FullPbClass }}
		{{- else -}}
			{{ range .NonOneOfFields -}}
				assertField_{{ .DartName }}();
			{{ end -}}
			{{ range .OneOfs }}
				assertOneof_{{ .DartName }}();
			{{- end -}}
		{{- end }}
	}

	{{ range .Fields -}}
		void assertField_{{ .DartName }}() {
			{{ render .Validate }}
		}
	{{ end -}}

	{{ range .OneOfs }}
		void assertOneof_{{ .DartName }}() {
			{{ template "oneOf" .Validate }}
		}
	{{- end -}}
}
`
