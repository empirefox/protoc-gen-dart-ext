package pgvt

const repeatedConstTpl = `{{- $r := .Pgv.Rules }}
{{- $elemField := .Elem "" "" }}
{{ if or $r.GetUnique (ne $elemField.Pgv.Typ "none") }}
	{{ renderConstants $elemField }}
{{ end }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Pgv.Rules }}
{{ if or $r.GetMinItems $r.GetMaxItems }}
	final _rl = {{ .PgdeFile.AsDot "Lists" }}.len({{ .Accessor }});
{{ end }}
{{ if $r.GetMinItems }}
	{{ if eq $r.GetMinItems $r.GetMaxItems }}
		if (_rl != {{ $r.GetMinItems }})
			throw {{ .PgdeFile.AsDot "ItemsLenConstError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinItems }});
	{{ else if $r.MaxItems }}
		if (_rl < {{ $r.GetMinItems }} || _rl > {{ $r.GetMaxItems }})
			throw {{ .PgdeFile.AsDot "RangeError" }}(ErrorRange.outEE ({{ .Err4Args }}, {{ $r.GetMinItems }}, {{ $r.GetMaxItems }}));
	{{ else }}
		if (_rl < {{ $r.GetMinItems }})
			throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinItems }});
	{{ end }}
{{ else if $r.MaxItems }}
	if (_rl > {{ $r.GetMaxItems }})
		throw {{ .PgdeFile.AsDot "ConstError" }}({{ .Err4Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxItems }});
{{ end }}

{{ if $r.GetUnique }}
	final Set<
		{{- if .RepeatedElemIsBytes -}}
			{{ .PgdeFile.AsDot "Bytes" }}
		{{- else -}}
			{{ .RepeatedElemType }}
		{{- end -}}
	> 
		_unique = {{ .CollectionLib.AsDot "HashSet" }}();
{{ end }}

{{ if or $r.GetUnique (ne (.Elem "" "").Pgv.Typ "none") }}
	for (var _ridx = 0; _ridx < {{ .Accessor }}.length; _ridx++) {
		final _ritem = {{ .Accessor }}[_ridx];
		{{ if $r.GetUnique }}
			if (!_unique.add(
				{{- if .RepeatedElemIsBytes -}}
					{{ .PgdeFile.AsDot "Bytes" }}(_ritem)
				{{- else -}}
					_ritem
				{{- end -}}
			)) throw {{ .PgdeFile.AsDot "BeSomethingError" }}({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateUnique(_ridx+1));
		{{ end }}

		{{ render (.Elem "_ritem" "_ridx") }}
	}
{{ end }}
`
