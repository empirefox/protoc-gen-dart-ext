package pgvt

const msgTpl = `
{{- $file := fileBase . }}
{{- $ShortType := name . -}}
{{- $Type := $Type | printf "$0.%s" -}}

/// Validates [{{ $Type }}] protobuf objects.
class {{ $ShortType }}Validator extends $pgde.GeneratedValidator<{{ $Type }}> {
	static {{ $ShortType }}Validator _instance = {{ $ShortType }}Validator._();

	factory {{ $ShortType }}Validator() => _instance;

	{{ $ShortType }}Validator._();

	{{- range .NonOneOfFields }}
		{{ renderConstants (context .) }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOfConst" . }}
	{{ end }}

	void assertProto($pgde.ValidateInfo info, $l10n.{{ $file }}Localizations _l10n) {
		{{ if disabled . }}
			// Validate is disabled for {{ $Type }}
		{{- else -}}
			{{ range .NonOneOfFields -}}
				{{ $ctx := . | context | dartContext }}
				assertField_{{ .Name }}({{ $ctx.Accessor }});
			{{ end -}}
			{{ range .OneOfs }}
				{{ $ctx := . | oneofContext }}
				assertOneof_{{ .Name.UpperCamelCase }}(proto, {{ $ctx.Rules.GetRequired }});
			{{- end -}}
		{{- end }}
	}

	{{ range .Fields -}}
		{{ $ctx := . | context | dartContext }}
		void assertField_{{ .Name }}({{ $ctx.FieldOriginType }} _v) {
			{{ render $ctx }}
		}
	{{ end -}}

	{{ range .OneOfs }}
		{{ $ctx := . | oneofContext }}
		{{- $OneofType := dartNameOf . | printf "$0.%s_%s" -}}
		{{- $oname := .Name.UpperCamelCase -}}
		void assertOneofRequired_{{ .Name.UpperCamelCase }}({{ $OneofType }}_{{ $oname }} which) {
			{{ if $ctx.Rules.GetRequired }}
				if (which == null || which == {{ $OneofType }}_{{ $oname }}.notSet)
					throw $pgde.OneofRequiredError(info, {{ $ctx.L10nOneofName }});
			{{ end }}
		}

		void assertOneof_{{ .Name.UpperCamelCase }}({{ $Type }} proto, bool required) {
			{{ template "oneOf" . }}
		}
	{{- end -}}
}
`
