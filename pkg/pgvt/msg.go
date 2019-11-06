package pgvt

const msgTpl = `
/// Validates [{{ .FullPbClass }}] protobuf objects.
class {{ .Names.ValidatorName }} extends {{ .PgdeFile.AsDot "GeneratedValidator" }}<{{ .FullPbClass }}> {
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
			final {{ .FieldAccessor }} = {{ .InfoAccessor }}.proto.{{ .DartName }};
			{{ render .Validate }}
		}
	{{ end -}}

	{{ range .OneOfs }}
		void assertOneof_{{ .DartName }}() {
			{{ template "oneOf" .Validate }}
		}
	{{- end -}}

	{{ .MaterialFile.AsDot "BuildContext" }} {{ .Validator.BuildContextAccessor }};

	{{ .PgdeFile.AsDot "ValidateInfo" }}<{{ .FullPbClass }}> {{ .Validator.InfoAccessor }};

	{{ .L10nFile.AsDot .File.Names.L10nName }} {{ .Validator.L10nAccessor }};

	{{ range .Validator.ImportManager.InstanceClasses }}
		{{ .FullName }} {{ .Instance }};
	{{ end }}

	{{ .Names.ValidatorName }}({{ .MaterialFile.AsDot "BuildContext" }} context, {{ .PgdeFile.AsDot "ValidateInfo" }}<{{ .FullPbClass }}> info)
		: {{ .Validator.BuildContextAccessor }} = context,
		  {{ .Validator.InfoAccessor }} = info,
		  {{ .Validator.L10nAccessor }} = {{ .MaterialFile.AsDot "Localizations" }}.of<{{ .FullL10nClass }}>(context, {{ .FullL10nClass }})
		  {{- range .Validator.ImportManager.InstanceClasses }}
		  	,{{ .Instance }} = {{ .FullName }}.of(context)
	  	  {{ end -}}
	  ;
}
`
