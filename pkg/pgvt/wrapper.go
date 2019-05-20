package pgvt

const wrapperConstTpl = `{{ $r := .Rules }}			
{{- renderConstants (unwrap .) }}
`

const wrapperTpl = `{{ $r := .Rules }}
{{ .IfHasBegin }}
	{{ render (unwrap . "wrapper") }}
{{ .IfHasEnd }}
`
