package pgvt

// num and string
const inConstTpl = inTplDef + `
{{- if $r.In }}
	static const Map<{{ .DartConstRefDefineType }}, Null> {{ $kIn }} = {
		{{- range $r.In }}
		{{ lit $ . }}: null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static const Map<{{ .DartConstRefDefineType }}, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		{{ lit $ . }}: null,
		{{- end }}
	};
{{- end -}}`

const inTplDef = `{{ $r := .Pgv.Rules }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const inTpl = inTplDef + `
{{- if $r.In }}
	if (!{{ $kIn }}.containsKey({{ .Accessor }}))
		throw {{ .PgdeFile.AsDot "InError" }}({{ .Err4Args }}, {{ $kIn }}.keys.toList());
{{- end -}}
{{- if $r.NotIn }}
	if ({{ $kNotIn }}.containsKey({{ .Accessor }}))
		throw {{ .PgdeFile.AsDot "InError" }}.not({{ .Err3Args }}, {{ $kNotIn }}.keys.toList());
{{- end -}}
`
