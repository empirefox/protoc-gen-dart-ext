package pgvt

const mapConstTpl = `{{ $r := .Pgv.Rules }}
{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
		{{ renderConstants (.Key "_mk" "_mk") }}
		{{ renderConstants (.Elem "_mv" "_mk") }}
{{- end -}}
`

const mapTpl = `{{ $r := .Pgv.Rules }}
{{ if or $r.GetMinPairs $r.GetMaxPairs }}
	final _ml = {{ .Accessor }}.length;
{{ end }}
{{ if $r.GetMinPairs }}
	{{ if eq $r.GetMinPairs $r.GetMaxPairs }}
		if (_ml != {{ $r.GetMinPairs }})
			throw {{ .PgdeFile.As }}.ItemsLenConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $r.GetMinPairs }});
	{{ else if $r.MaxPairs }}
		if (_ml < {{ $r.GetMinPairs }} || _ml > {{ $r.GetMaxPairs }})
			throw {{ .PgdeFile.As }}.RangeError(ErrorRange.outEE ({{ .Err3Args }}, {{ $r.GetMinPairs }}, {{ $r.GetMaxPairs }}));
	{{ else }}
		if (_ml < {{ $r.GetMinPairs }})
			throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateGte, {{ $r.GetMinPairs }});
	{{ end }}
{{ else if $r.MaxPairs }}
	if (_ml > {{ $r.GetMaxPairs }})
		throw {{ .PgdeFile.As }}.ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateLte, {{ $r.GetMaxPairs }});
{{ end }}

{{ if or (ne (.Elem "" "").Typ "none") (ne (.Key "" "").Typ "none") }}
	{{ .Accessor }}.forEach((_mk, _mv) {
		{{ if $r.GetNoSparse }}
			// TODO(alway no sparse here)
			// if (val == null) throw BeSomethingError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateNoSparse(key));
		{{ end }}

		{{ render (.Key "_mk" "_mk") }}

		{{ render (.Elem "_mv" "_mk") }}
	});
{{ end }}
`
