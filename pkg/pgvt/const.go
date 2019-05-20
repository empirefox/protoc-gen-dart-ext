package pgvt

const constTpl = `{{ $r := .Rules }}

{{ if $r.Const }}{{ $constLiteral := lit $r.GetConst }}
	if (_v != {{ $constLiteral }})
		throw ConstError(info, {{ $.Number }}, {{ $.L10nField }}, info.l10n.validateEq, {{ $constLiteral }});
{{ end }}
`
