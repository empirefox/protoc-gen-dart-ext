package pgvt

const repeatedConstTpl = `{{ $r := .Rules }}
{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
	{{ renderConstants (.Elem "" "") }}
{{ end }}`

const repeatedTpl = `{{ $f := .Field }}{{ $r := .Rules }}
{{ if or $r.GetMinItems $r.GetMaxItems }}
	final _vl = {{ .PgdeFile.As }}.Lists.len(_v);
{{ end }}
{{ if $r.GetMinItems }}
	{{ if eq $r.GetMinItems $r.GetMaxItems }}
		if (_vl != {{ $r.GetMinItems }})
			throw {{ .PgdeFile.As }}.ItemsLenConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetMinItems }});
	{{ else if $r.MaxItems }}
		if (_vl < {{ $r.GetMinItems }} || _vl > {{ $r.GetMaxItems }})
			throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetMinItems }}, {{ $r.GetMaxItems }}));
	{{ else }}
		if (_vl < {{ $r.GetMinItems }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetMinItems }});
	{{ end }}
{{ else if $r.MaxItems }}
	if (_vl > {{ $r.GetMaxItems }})
		throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetMaxItems }});
{{ end }}

{{ if $r.GetUnique }}
	final Set<
		{{- if isBytes $f.Type.Element -}}
			{{ .PgdeFile.As }}.Bytes
		{{- else -}}
			{{ (typ $f).Element }}
		{{- end -}}
	> 
		_unique = HashSet();
{{ end }}

{{ if or $r.GetUnique (ne (.Elem "" "").Typ "none") }}
	for var idx = 0; idx < _v.length; idx++ {
		final item = _v[idx];
		{{ if $r.GetUnique }}
			if (!_unique.add(
				{{- if isBytes $f.Type.Element -}}
					{{ .PgdeFile.As }}.Bytes(item)
				{{- else -}}
					item
				{{- end -}}
			)) throw {{ .PgdeFile.As }}.BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateUnique(idx+1));
		{{ end }}

		{{ render (.Elem "item" "idx") }}
	}
{{ end }}
`
