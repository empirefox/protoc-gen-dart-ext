package pgvt

const constTpl = `{{ $r := .Pgv.Rules }}

{{ if $r.Const }}{{ $constLiteral := lit . $r.GetConst }}
	if ({{ .Accessor }} != {{ $constLiteral }})
		throw ConstError({{ .Err3Args }}, {{ .InfoAccessor }}.l10n.validateEq, {{ $constLiteral }});
{{ end }}
`
