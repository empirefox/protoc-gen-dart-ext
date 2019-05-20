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

const timestampTplDef = `{{ $r := .Rules }}
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
	final _ts = _v.seconds*Duration.microsecondsPerSecond + _v.nanos~/1000;

	{{ if or $r.LtNow $r.GtNow $r.Within }}
		final _now = DateTime.now().microsecondsSinceEpoch;
	{{ end }}

	{{ if $r.Const }}
		if (_ts != {{ $kConst }}.microsecondsSinceEpoch)
			throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $kConst }});
	{{ end }}

	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if tsGt $r.GetLt $r.GetGt }}
				if (_ts <= {{ $kGt }}.microsecondsSinceEpoch || _ts >= {{ $kLt }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.inNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGt }}, {{ $kLt }}));
			{{ else }}
				if (_ts >= {{ $kLt }}.microsecondsSinceEpoch && _ts <= {{ $kGt }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.outNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLt }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if tsGt $r.GetLt $r.GetGte }}
				if (_ts < {{ $kGte }}.microsecondsSinceEpoch || _ts >= {{ $kLt }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.inEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGte }}, {{ $kLt }}));
			{{ else }}
				if (_ts >= {{ $kLt }}.microsecondsSinceEpoch && _ts < {{ $kGte }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.outNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLt }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_ts >= {{ $kLt }}.microsecondsSinceEpoch)
				throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLt, {{ $kLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if tsGt $r.GetLte $r.GetGt }}
				if (_ts <= {{ $kGt }}.microsecondsSinceEpoch || _ts > {{ $kLte }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.inNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGt }}, {{ $kLte }}));
			{{ else }}
				if (_ts > {{ $kLte }}.microsecondsSinceEpoch && _ts <= {{ $kGt }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.outEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLte }}, {{ $kGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if tsGt $r.GetLte $r.GetGte }}
				if (_ts < {{ $kGte }}.microsecondsSinceEpoch || _ts > {{ $kLte }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kGte }}, {{ $kLte }}));
			{{ else }}
				if (_ts > {{ $kLte }}.microsecondsSinceEpoch && _ts < {{ $kGte }}.microsecondsSinceEpoch)
					throw $pgde.RangeError(ErrorRange.outEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kLte }}, {{ $kGte }}));
			{{ end }}
		{{ else }}
			if (_ts > {{ $kLte }}.microsecondsSinceEpoch)
				throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $kLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if (_ts <= {{ $kGt }}.microsecondsSinceEpoch)
			throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGt, {{ $kGt }});
	{{ else if $r.Gte }}
		if (_ts < {{ $kGte }}.microsecondsSinceEpoch)
			throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $kGte }});
	{{ else if $r.LtNow }}
		{{ if $r.Within }}
			if (_ts >= _now || _now - _ts > {{ $kWithin }}.inMicroseconds)
				throw $pgde.WithinError.ltNow(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kWithin }});
		{{ else }}
			if (_ts >= _now)
				throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLt, l10n.now);
		{{ end }}
	{{ else if $r.GtNow }}
		{{ if $r.Within }}
			if (_ts >= _now || ts - _now < {{ $kWithin }}.inMicroseconds)
				throw $pgde.WithinError.gtNow(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kWithin }});
		{{ else }}
			if (_ts <= _now)
				throw $pgde.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGt, l10n.now);
		{{ end }}
	{{ else if $r.Within }}
		if ((_ts - _now).abs() <= {{ $kWithin }}.inMicroseconds)
			throw $pgde.WithinError.now(info, {{ $.Number }}, {{ $.L10nField }}, {{ $kWithin }});
	{{ end }}
{{ end }}
`
