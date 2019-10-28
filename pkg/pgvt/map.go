package pgvt

const mapConstTpl = `{{ $r := .Rules }}
{{ if or $r.GetNoSparse (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{ renderConstants (.Key "key" "Key") }}
		{{ renderConstants (.Elem "value" "Value") }}
{{- end -}}
`

const mapTpl = `{{ $r := .Rules }}
{{ if or $r.GetMinPairs $r.GetMaxPairs }}
	final _vl = _v.length;
{{ end }}
{{ if $r.GetMinPairs }}
	{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
		if (_vl != {{ $r.GetMinPairs }})
			throw {{ .PgdeFile.As }}.ItemsLenConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $r.GetMinPairs }});
	{{ else if $r.MaxPairs }}
		if (_vl < {{ $r.GetMinPairs }} || _vl > {{ $r.GetMaxPairs }})
			throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE(info, {{ $.Number }}, {{ $.L10nField }}, {{ $r.GetMinPairs }}, {{ $r.GetMaxPairs }}));
	{{ else }}
		if (_vl < {{ $r.GetMinPairs }})
			throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateGte, {{ $r.GetMinPairs }});
	{{ end }}
{{ else if $r.MaxPairs }}
	if (_vl > {{ $r.GetMaxPairs }})
		throw {{ .PgdeFile.As }}.ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateLte, {{ $r.GetMaxPairs }});
{{ end }}

{{ if or $r.GetNoSparse (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
	_v.forEach((key, val) {
		{{ if $r.GetNoSparse }}
			if (val == null) throw BeSomethingError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateNoSparse(key));
		{{ end }}

		{{ render (.Key "key" "key") }}

		{{ render (.Elem "val" "key") }}
	});
{{ end }}
`
