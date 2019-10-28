package pgvt

const packageTpl = `
{{ range .Messages }}
	{{ template "msg" .Validate }}
{{ end }}

{{ .Package.Group.SplitHeader }}

// Code generated by protoc-gen-dart-ext. DO NOT EDIT.
// package: {{ .Pgs.ProtoName }}
{{ range .Pgs.Files }}
	// source: {{ .InputPath }}
{{ end }}

{{- range .Validator.ImportManager.Files }}
	import '{{ .Name }}'
		{{- if .As }} as {{ .As }}{{ end -}}
		{{- if .IsShow }} show {{ .Show }}{{ end -}}
	;
{{- end }}
`
