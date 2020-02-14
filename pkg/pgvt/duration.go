package pgvt

const durationConstTpl = durationTplDef + `
{{- if $r.Const }}
	static final Duration {{ $kConst }} = {{ coreDur $r.Const }};
{{- end -}}
{{- if $r.Lt }}
	static final Duration {{ $kLt }} = {{ coreDur $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
	static final Duration {{ $kLte }} = {{ coreDur $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
	static final Duration {{ $kGt }} = {{ coreDur $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
	static final Duration {{ $kGte }} = {{ coreDur $r.GetGte }};
{{- end -}}
{{- if $r.In }}
	static final Map<Duration, Null> {{ $kIn }} = {
		{{- range $r.In }}
		{{ coreDur . }}: null,
		{{- end }}
	};
{{- end -}}
{{- if $r.NotIn }}
	static final Map<Duration, Null> {{ $kNotIn }} = {
		{{- range $r.NotIn }}
		{{ coreDur . }}: null,
		{{- end }}
	};
{{- end -}}
`

const durationTplDef = `{{ $r := .Pgv.Rules }}
{{- $kConst := constRef . "kConst" }}
{{- $kLt := constRef . "kLt" }}
{{- $kLte := constRef . "kLte" }}
{{- $kGt := constRef . "kGt" }}
{{- $kGte := constRef . "kGte" }}
{{- $kIn := constRef . "kIn" }}
{{- $kNotIn := constRef . "kNotIn" }}
`

const durationTpl = durationTplDef + `
{{ template "required" . }}

{{ if or $r.In $r.NotIn $r.Lt $r.Lte $r.Gt $r.Gte $r.Const }}
{{ .IfHasValueBegin }}
	final _dur = Duration(seconds: {{ .Accessor }}.seconds, microseconds: {{ .Accessor }}.nanos~/1000);

	{{ if $r.Const }}
		if (_dur != {{ $kConst }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $kConst }});
	{{ end }}

	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if durGt $r.GetLt $r.GetGt }}
				if (_dur <= {{ $kGt }} || _dur >= {{ $kLt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inNN ({{ .Err4Args }}, {{ $kGt }}, {{ $kLt }}));
			{{ else }}
				if (_dur >= {{ $kLt }} && _dur <= {{ $kGt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outNN ({{ .Err4Args }}, {{ $kLt }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if durGt $r.GetLt $r.GetGte }}
				if (_dur < {{ $kGte }} || _dur >= {{ $kLt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inEN ({{ .Err4Args }}, {{ $kGte }}, {{ $kLt }}));
			{{ else }}
				if (_dur >= {{ $kLt }} && _dur < {{ $kGte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outNE ({{ .Err4Args }}, {{ $kLt }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_dur >= {{ $kLt }})
				throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateLt, {{ $kLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if durGt $r.GetLte $r.GetGt }}
				if (_dur <= {{ $kGt }} || _dur > {{ $kLte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inNE ({{ .Err4Args }}, {{ $kGt }}, {{ $kLte }}));
			{{ else }}
				if (_dur > {{ $kLte }} && _dur <= {{ $kGt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outEN ({{ .Err4Args }}, {{ $kLte }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if durGt $r.GetLte $r.GetGte }}
				if (_dur < {{ $kGte }} || _dur > {{ $kLte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inEE ({{ .Err4Args }}, {{ $kGte }}, {{ $kLte }}));
			{{ else }}
				if (_dur > {{ $kLte }} && _dur < {{ $kGte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outEE ({{ .Err4Args }}, {{ $kLte }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_dur > {{ $kLte }})
				throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $kLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if (_dur <= {{ $kGt }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateGt, {{ $kGt }});
	{{ else if $r.Gte }}
		if (_dur < {{ $kGte }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $kGte }});
	{{ end }}

	{{ if $r.In }}
		if (!{{ $kIn }}.containsKey(_dur))
			throw {{ .PgdeFile.AsDot "InError" }}({{ .Err4Args }}, {{ $kIn }}.keys.toList());
	{{ else if $r.NotIn }}
		if ({{ $kNotIn }}.containsKey(_dur))
			throw {{ .PgdeFile.AsDot "InError" }}.not ({{ .Err4Args }}, {{ $kNotIn }}.keys.toList());
	{{ end }}
{{ .IfHasValueEnd }}
{{ end }}
`
