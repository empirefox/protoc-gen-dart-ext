package pgvt

const numConstTpl = `
{{ template "inConst" . }}
`

const numTpl = `
{{ template "const" . }}
{{ template "ltgt" . }}
{{ template "in" . }}
`
