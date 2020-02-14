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

const anyTplDef = `{{ $r := .Pgv.Rules }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const anyTpl = anyTplDef + `
{{- template "required" . -}}

{{ .IfHasValueBegin }}
	{{- if $r.In }}
		if (!{{ $kIn }}.containsKey({{ .Accessor }}.typeUrl))
			throw {{ .PgdeFile.AsDot "InError" }}({{ .Err4Args }}, {{ $kIn }}.keys.toList());
	{{- end -}}
	{{- if $r.NotIn }}
		if ({{ $kNotIn }}.containsKey({{ .Accessor }}.typeUrl))
			throw {{ .PgdeFile.AsDot "InError" }}.not ({{ .Err3Args }}, {{ $kNotIn }}.keys.toList());
	{{- end -}}
{{ .IfHasValueEnd }}
`
