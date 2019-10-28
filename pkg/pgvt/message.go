package pgvt

const messageTpl = `
	{{ if .Pgv.MessageRules.GetSkip }}
		// skipping validation for {{ .DartName }}
	{{ else }}
		{{ template "required" . }}
		{{ if (isOfMessageType .) }}
			{{ .FullFieldTypeMessage }}.assertProto({{ .Validator.BuildContextAccessor }}, {{ .Validator.InfoAccessor }}.clone({{ .Accessor }}));
		{{ end }}
	{{ end }}
`
