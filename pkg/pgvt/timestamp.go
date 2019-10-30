package pgvt

const timestampConstTpl = timestampTplDef + `
{{- if $r.Const }}
	static final DateTime {{ $kConst }} = {{ coreTs $r.Const }};
{{- end -}}
{{- if $r.Lt }}
	static final DateTime {{ $kLt }} = {{ coreTs $r.GetLt }};
{{- end -}}
{{- if $r.Lte }}
	static final DateTime {{ $kLte }} = {{ coreTs $r.GetLte }};
{{- end -}}
{{- if $r.Gt }}
	static final DateTime {{ $kGt }} = {{ coreTs $r.GetGt }};
{{- end -}}
{{- if $r.Gte }}
	static final DateTime {{ $kGte }} = {{ coreTs $r.GetGte }};
{{- end -}}
{{- if $r.Within }}
	static final Duration {{ $kWithin }} = {{ coreDur $r.Within }};
{{- end -}}
`

const timestampTplDef = `{{ $r := .Pgv.Rules }}
{{- $kConst := constRef . "kConst" }}
{{- $kLt := constRef . "kLt" }}
{{- $kLte := constRef . "kLte" }}
{{- $kGt := constRef . "kGt" }}
{{- $kGte := constRef . "kGte" }}
{{- $kWithin := constRef . "kWithin" }}
`

const timestampTpl = timestampTplDef + `
{{ template "required" . }}

{{ if or $r.Lt $r.Lte $r.Gt $r.Gte $r.LtNow $r.GtNow $r.Within $r.Const }}
{{ .IfHasValueBegin }}
	final _ts = {{ .Accessor }}.seconds*Duration.microsecondsPerSecond + {{ .Accessor }}.nanos~/1000;

	{{ if or $r.LtNow $r.GtNow $r.Within }}
		final _now = DateTime.now().microsecondsSinceEpoch;
	{{ end }}

	{{ if $r.Const }}
		if (_ts != {{ $kConst }}.microsecondsSinceEpoch)
			throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $kConst }});
	{{ end }}

	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if tsGt $r.GetLt $r.GetGt }}
				if (_ts <= {{ $kGt }}.microsecondsSinceEpoch || _ts >= {{ $kLt }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNN ({{ .Err3Args }}, {{ $kGt }}, {{ $kLt }}));
			{{ else }}
				if (_ts >= {{ $kLt }}.microsecondsSinceEpoch && _ts <= {{ $kGt }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNN ({{ .Err3Args }}, {{ $kLt }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if tsGt $r.GetLt $r.GetGte }}
				if (_ts < {{ $kGte }}.microsecondsSinceEpoch || _ts >= {{ $kLt }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEN ({{ .Err3Args }}, {{ $kGte }}, {{ $kLt }}));
			{{ else }}
				if (_ts >= {{ $kLt }}.microsecondsSinceEpoch && _ts < {{ $kGte }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNE ({{ .Err3Args }}, {{ $kLt }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_ts >= {{ $kLt }}.microsecondsSinceEpoch)
				throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLt, {{ $kLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if tsGt $r.GetLte $r.GetGt }}
				if (_ts <= {{ $kGt }}.microsecondsSinceEpoch || _ts > {{ $kLte }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNE ({{ .Err3Args }}, {{ $kGt }}, {{ $kLte }}));
			{{ else }}
				if (_ts > {{ $kLte }}.microsecondsSinceEpoch && _ts <= {{ $kGt }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEN ({{ .Err3Args }}, {{ $kLte }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if tsGt $r.GetLte $r.GetGte }}
				if (_ts < {{ $kGte }}.microsecondsSinceEpoch || _ts > {{ $kLte }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEE ({{ .Err3Args }}, {{ $kGte }}, {{ $kLte }}));
			{{ else }}
				if (_ts > {{ $kLte }}.microsecondsSinceEpoch && _ts < {{ $kGte }}.microsecondsSinceEpoch)
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE ({{ .Err3Args }}, {{ $kLte }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_ts > {{ $kLte }}.microsecondsSinceEpoch)
				throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $kLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if (_ts <= {{ $kGt }}.microsecondsSinceEpoch)
			throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGt, {{ $kGt }});
	{{ else if $r.Gte }}
		if (_ts < {{ $kGte }}.microsecondsSinceEpoch)
			throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $kGte }});
	{{ else if $r.LtNow }}
		{{ if $r.Within }}
			if (_ts >= _now || _now - _ts > {{ $kWithin }}.inMicroseconds)
				throw {{ .PgdeFile.As }}.WithinError.ltNow ({{ .Err3Args }}, {{ $kWithin }});
		{{ else }}
			if (_ts >= _now)
				throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLt, l10n.now);
		{{ end }}
	{{ else if $r.GtNow }}
		{{ if $r.Within }}
			if (_ts >= _now || ts - _now < {{ $kWithin }}.inMicroseconds)
				throw {{ .PgdeFile.As }}.WithinError.gtNow ({{ .Err3Args }}, {{ $kWithin }});
		{{ else }}
			if (_ts <= _now)
				throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGt, l10n.now);
		{{ end }}
	{{ else if $r.Within }}
		if ((_ts - _now).abs() <= {{ $kWithin }}.inMicroseconds)
			throw {{ .PgdeFile.As }}.WithinError.now ({{ .Err3Args }}, {{ $kWithin }});
	{{ end }}
{{ .IfHasValueEnd }}
{{ end }}
`
