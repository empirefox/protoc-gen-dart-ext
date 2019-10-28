package pgvt

const ltgtTpl = `{{ $r := .Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if (_v <= {{ $r.Gt }} || _v >= {{ $r.Lt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetGt }}, {{ $r.GetLt }}));
			{{ else }}
				if (_v >= {{ $r.Lt }} && _v <= {{ $r.Gt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetLt }}, {{ $r.GetGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if (_v < {{ $r.Gte }} || _v >= {{ $r.Lt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetGte }}, {{ $r.GetLt }}));
			{{ else }}
				if (_v >= {{ $r.Lt }} && _v < {{ $r.Gte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetLt }}, {{ $r.GetGte }}));
			{{ end }}
		{{ else }}
			if (_v >= {{ $r.Lt }})
				throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLt, {{ $r.GetLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if (_v <= {{ $r.Gt }} || _v > {{ $r.Lte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inNE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetGt }}, {{ $r.GetLte }}));
			{{ else }}
				if (_v > {{ $r.Lte }} && _v <= {{ $r.Gt }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEN(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetLte }}, {{ $r.GetGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if (_v < {{ $r.Gte }} || _v > {{ $r.Lte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.inEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetGte }}, {{ $r.GetLte }}));
			{{ else }}
				if (_v > {{ $r.Lte }} && _v < {{ $r.Gte }})
					throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetLte }}, {{ $r.GetGte }}));
			{{ end }}
		{{ else }}
			if (_v > {{ $r.Lte }})
				throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if (_v <= {{ $r.Gt }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGt, {{ $r.GetGt }});
	{{ else if $r.Gte }}
		if (_v < {{ $r.Gte }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetGte }});
	{{ end }}
`
