package pgvt

const messageTpl = `
	{{ if .Pgv.MessageRules.GetSkip }}
		// skipping validation for {{ .DartName }}
	{{ else }}
		{{ template "required" . }}
		{{ if .IsOfMessageType }}
			{{ $target := .MessageFullValidatorClass }}
			{{ if $target }}
				{{ $target }}({{ .BuildContextAccessor }}, {{ .InfoAccessor }}.clone({{ .Accessor }})).assertProto();
			{{ end }}
		{{ end }}
	{{ end }}
`
