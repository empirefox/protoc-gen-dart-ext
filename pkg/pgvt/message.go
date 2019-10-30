package pgvt

const messageTpl = `
	{{ if .Pgv.MessageRules.GetSkip }}
		// skipping validation for {{ .DartName }}
	{{ else }}
		{{ template "required" . }}
		{{ if (isOfMessageType .) }}
			{{ .MessageFullValidatorClass }}({{ .BuildContextAccessor }}, {{ .InfoAccessor }}.clone({{ .Accessor }})).assertProto();
		{{ end }}
	{{ end }}
`
