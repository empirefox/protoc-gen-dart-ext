package pgvt

const messageTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}
	{{ template "required" . }}

	{{ if $r.GetSkip }}
		// skipping validation for {{ $f.Name }}
	{{ else }}
		{{ .MessageValidator }}.assertProto({{ .Accessor }});
	{{ end }}
`
