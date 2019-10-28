package pgvt

const anyConstTpl = anyTplDef + `
{{- if $r.In }}
	static const Map<String, Null> {{ $kIn }} = {
		{{- range $r.In }}
		"{{ . }}": null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static const Map<String, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		"{{ . }}": null,
		{{- end }}
	};
{{- end -}}`

const anyTplDef = `{{ $r := .Rules }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const anyTpl = anyTplDef + `
{{- template "required" . -}}

{{ .IfHasBegin }}
	{{- if $r.In }}
		if (!{{ $kIn }}.containsKey(_v.typeUrl))
			throw {{ .PgdeFile.As }}.InError(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kIn }}.keys);
	{{- end -}}
	{{- if $r.NotIn }}
		if ({{ $kNotIn }}.containsKey(_v.typeUrl))
			throw {{ .PgdeFile.As }}.InError.not(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kNotIn }}.keys);
	{{- end -}}
{{ .IfHasEnd }}
`
