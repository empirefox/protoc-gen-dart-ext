package pgvt

const ltgtTpl = `{{ $r := .Pgv.Rules }}
	{{ if $r.Lt }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLt $r.GetGt }}
				if ({{ .Accessor }} <= {{ $r.Gt }} || {{ .Accessor }} >= {{ $r.Lt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inNN ({{ .Err3Args }}, {{ $r.GetGt }}, {{ $r.GetLt }}));
			{{ else }}
				if ({{ .Accessor }} >= {{ $r.Lt }} && {{ .Accessor }} <= {{ $r.Gt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outNN ({{ .Err3Args }}, {{ $r.GetLt }}, {{ $r.GetGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{  if gt $r.GetLt $r.GetGte }}
				if ({{ .Accessor }} < {{ $r.Gte }} || {{ .Accessor }} >= {{ $r.Lt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inEN ({{ .Err3Args }}, {{ $r.GetGte }}, {{ $r.GetLt }}));
			{{ else }}
				if ({{ .Accessor }} >= {{ $r.Lt }} && {{ .Accessor }} < {{ $r.Gte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outNE ({{ .Err3Args }}, {{ $r.GetLt }}, {{ $r.GetGte }}));
			{{ end }}
		{{ else }}
			if ({{ .Accessor }} >= {{ $r.Lt }})
				throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLt, {{ $r.GetLt }});
		{{ end }}
	{{ else if $r.Lte }}
		{{ if $r.Gt }}
			{{  if gt $r.GetLte $r.GetGt }}
				if ({{ .Accessor }} <= {{ $r.Gt }} || {{ .Accessor }} > {{ $r.Lte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inNE ({{ .Err3Args }}, {{ $r.GetGt }}, {{ $r.GetLte }}));
			{{ else }}
				if ({{ .Accessor }} > {{ $r.Lte }} && {{ .Accessor }} <= {{ $r.Gt }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outEN ({{ .Err3Args }}, {{ $r.GetLte }}, {{ $r.GetGt }}));
			{{ end }}
		{{ else if $r.Gte }}
			{{ if gt $r.GetLte $r.GetGte }}
				if ({{ .Accessor }} < {{ $r.Gte }} || {{ .Accessor }} > {{ $r.Lte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.inEE ({{ .Err3Args }}, {{ $r.GetGte }}, {{ $r.GetLte }}));
			{{ else }}
				if ({{ .Accessor }} > {{ $r.Lte }} && {{ .Accessor }} < {{ $r.Gte }})
					throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outEE ({{ .Err3Args }}, {{ $r.GetLte }}, {{ $r.GetGte }}));
			{{ end }}
		{{ else }}
			if ({{ .Accessor }} > {{ $r.Lte }})
				throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetLte }});
		{{ end }}
	{{ else if $r.Gt }}
		if ({{ .Accessor }} <= {{ $r.Gt }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGt, {{ $r.GetGt }});
	{{ else if $r.Gte }}
		if ({{ .Accessor }} < {{ $r.Gte }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetGte }});
	{{ end }}
`
