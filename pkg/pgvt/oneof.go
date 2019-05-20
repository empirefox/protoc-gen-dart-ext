package pgvt

const oneOfTpl = `
{{- $OneofType := dartNameOf . | printf "$0.%s_%s" -}}
{{- $oneofName := .Name.UpperCamelCase -}}

switch (info.proto.which{{ $oneofName }}()) {
	{{ range .Fields -}}
	case {{ $OneofType }}_{{ $oneofName }}.{{ .Name }}:
		{{ $ctx := . | context | dartContext }}
		assertField_{{ .Name }}({{ $ctx.Accessor }});
		break;
	{{ end -}}
	default:
		if (required)
			throw $pgde.OneofRequiredError(info, {{ .L10nOneofName }});
}
`

type OneofContext struct {
}

func (ctx *OneofContext) L10nOneofName() string { return "_l10n." }
