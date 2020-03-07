package pgvt

const msgTpl = `
/// Validates [{{ .FullPbClass }}] protobuf objects.
class {{ .Names.ValidatorName }} implements {{ .PgdeFile.AsDot "GeneratedValidator" }}<{{ .FullPbClass }}> {
	
	{{- range .NonOneOfFields }}
		{{ renderConstants .Validate }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOfConst" .Validate }}
	{{ end }}

	@override
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

	@override
	void assertField(int tag){
		switch (tag) {
		{{ range .Fields -}}
			case {{ .Pgs.Descriptor.Number }}: return assertField_{{ .DartName }}();
		{{ end }}
		default:
			assert(false, 'tag number($tag) for ${ {{ .Validators.InfoAccessor }}.bi.messageName} not found');
		}
	}

	{{ range .Fields -}}
		void assertField_{{ .DartName }}() {
			{{ if ne .Validate.Pgv.Typ "none" }}
				final {{ $.Validators.FieldAccessor }} = {{ $.Validators.InfoAccessor }}.proto.{{ .DartName }};
			{{ end }}
			{{ render .Validate }}
		}
	{{ end -}}

	@override
	void assertOneof(Type oneof){
		switch (oneof) {
		{{ range .OneOfs }}
			case {{ .Validate.FullPbWhichEnum }}: return assertOneof_{{ .DartName }}();
		{{ end }}
		default:
			assert(false, 'oneof type($oneof) for ${ {{ .Validators.InfoAccessor }}.bi.messageName} not found');
		}
	}

	{{ range .OneOfs }}
		void assertOneof_{{ .DartName }}() {
			{{ template "oneOf" .Validate }}
		}
	{{- end -}}

	{{ .Material.BuildContext }} {{ .Validators.BuildContextAccessor }};

	{{ .PgdeFile.AsDot "ValidateInfo" }}<{{ .FullPbClass }}> {{ .Validators.InfoAccessor }};

	{{ .L10nType }} {{ .Validators.L10nAccessor }};

	{{ range .Validators.ImportManager.InstanceClasses }}
		{{ .FullName }} {{ .Instance }};
	{{ end }}

	static {{ .Names.ValidatorName }} create({{ .Material.BuildContext }} context, {{ .PgdeFile.AsDot "ValidateInfo" }}<{{ .FullPbClass }}> info)
		=> {{ .Names.ValidatorName }}(context, info);

	{{ .Names.ValidatorName }}({{ .Material.BuildContext }} context, {{ .PgdeFile.AsDot "ValidateInfo" }}<{{ .FullPbClass }}> info)
		: {{ .Validators.BuildContextAccessor }} = context,
		  {{ .Validators.InfoAccessor }} = info,
		  {{ .Validators.L10nAccessor }} = {{ .Material.Localizations }}.of<{{ .FullL10nClass }}>(context, {{ .FullL10nClass }})
		  {{- range .Validators.ImportManager.InstanceClasses }}
		  	,{{ .Instance }} = {{ .FullName }}.of(context)
	  	  {{ end -}}
	  ;
}
`
