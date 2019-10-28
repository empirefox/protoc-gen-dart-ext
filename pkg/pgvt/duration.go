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

const durationTplDef = `{{ $r := .Rules }}
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
	final _dur = Duration(seconds: _v.seconds, microseconds: _v.nanos~/1000);

	{{ if $r.Const }}
		if (_dur != {{ $kConst }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $kConst }});
	{{ end }}

	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if durGt $r.GetLt $r.GetGt }}
				if (_dur <= {{ $kGt }} || _dur >= {{ $kLt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGt }}, {{ $kLt }}));
			{{ else }}
				if (_dur >= {{ $kLt }} && _dur <= {{ $kGt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLt }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if durGt $r.GetLt $r.GetGte }}
				if (_dur < {{ $kGte }} || _dur >= {{ $kLt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGte }}, {{ $kLt }}));
			{{ else }}
				if (_dur >= {{ $kLt }} && _dur < {{ $kGte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLt }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_dur >= {{ $kLt }})
				throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLt, {{ $kLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if durGt $r.GetLte $r.GetGt }}
				if (_dur <= {{ $kGt }} || _dur > {{ $kLte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGt }}, {{ $kLte }}));
			{{ else }}
				if (_dur > {{ $kLte }} && _dur <= {{ $kGt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLte }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if durGt $r.GetLte $r.GetGte }}
				if (_dur < {{ $kGte }} || _dur > {{ $kLte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGte }}, {{ $kLte }}));
			{{ else }}
				if (_dur > {{ $kLte }} && _dur < {{ $kGte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLte }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_dur > {{ $kLte }})
				throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $kLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if (_dur <= {{ $kGt }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGt, {{ $kGt }});
	{{ else if $r.Gte }}
		if (_dur < {{ $kGte }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $kGte }});
	{{ end }}

	{{ if $r.In }}
		if (!{{ $kIn }}.containsKey(_dur))
			throw {{ .PgdeFile.As }}.InError(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kIn }}.keys);
	{{ else if $r.NotIn }}
		if ({{ $kNotIn }}.containsKey(_dur))
			throw {{ .PgdeFile.As }}.InError.not(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kNotIn }}.keys);
	{{ end }}
{{ end }}
`
