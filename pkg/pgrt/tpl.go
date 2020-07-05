package pgrt

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const fileTpl = `{{ renderJoin . "fileHeader" "fileBody+1" }}`

const fileHeaderTpl = genshared.DartHead + `
syntax = "{{ .Pgs.Descriptor.Syntax }}";

package {{ .Pgs.Descriptor.GetPackage }};
{{ with .Pgs.Descriptor.Options.GetGoPackage }}option go_package = "{{ . }}";{{ end }}

{{- range .RPCS.Imports.Values }}
	import "{{ . }}";
{{- end }}`

const fileBodyTpl = `
{{- range .Messages }}{{ with .RPC }}
	{{- if .IsEntry }}
		{{ template "entry" . }}
	{{- else if .IsLeaf }}
		{{ template "leaf" . }}
	{{- end }}
{{- end }}{{ end }}`

const leafTpl = `
{{- range .Fields }}{{ with .RPC }}
	{{- if .IsSelectOne }}
		{{ template "selectOne" . }}
	{{- else if .IsSelectMany }}
		{{ template "selectMany" . }}
	{{- end }}
{{- end }}{{ end }}`

const entryTpl = `
{{- range .Fields }}{{ with .RPC }}
	{{- if .IsForm }}
		{{ template "form" . }}
	{{- else if .IsList }}
		{{ template "list" . }}
	{{- else if .IsSelectOne }}
		{{ template "selectOne" . }}
	{{- else if .IsSelectMany }}
		{{ template "selectMany" . }}
	{{- end }}
{{- end }}{{ end }}`

const formTpl = `
service {{ .ServiceName }} {
	rpc Get({{ .Empty }}) returns ({{ .PayloadName }}.GetResp);
}

message {{ .PayloadName }} {
	message GetResp {
		{{ .TypeName }} data = 1;
		{{ .BackendErrorName }} error = 2;
	}
}`

const listTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	rpc Get({{ .PayloadName }}.SrcId) returns ({{ .PayloadName }}.GetResp);
	{{ if .Add -}}
		rpc Add({{ .RefTypeName }}) returns ({{ .PayloadName }}.AddResp);
	{{- end }}
	{{ if .Save -}}
		rpc Save({{ .RefTypeName }}) returns ({{ .BackendErrorName }});
	{{- end }}
	{{- template "srcRPCRemove" . }}
}

message {{ .PayloadName }} {
	{{- template "srcResp" . }}
	{{- if .IsGroup }}
		message Element {
			oneof is {
				{{ range .Group.Fields -}}
					{{ $.TypeNameOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
				{{- end }}
			}
		}
	{{- end }}
	message GetResp {
		{{ if .IsGroup }}Element{{ else }}{{ .NonGroupTypeName }}{{ end }} data = 1;
		{{ .BackendErrorName }} error = 2;
	}
	{{- if .Add }}
		message AddResp {
			{{ .ViewName }} data = 1;
			{{ .BackendErrorName }} error = 2;
		}
	{{- end }}
}`

const srcRPCLineTpl = `rpc Src({{ .Empty }}) returns ({{ .PayloadName }}.SrcResp);`

const srcRPCRemoveTpl = `
{{ if .Remove -}}
	rpc Remove({{ .PayloadName }}.SrcId) returns ({{ .BackendErrorName }});
{{- end }}
{{ if .RemoveAll -}}
	rpc RemoveAll({{ .Empty }}) returns ({{ .BackendErrorName }});
{{- end }}`

const srcRespTpl = `
message SrcId {
	{{- if .IsGroup }}
		oneof is {
			{{ range .Group.Fields -}}
				{{ $.IdTypeOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
			{{- end }}
		}
	{{- else }}
		{{ .IdTypeOfFieldType .Pgs }} id = 1;
	{{- end }}
}
{{ if .IsGroupByType -}}
	message Group {
		{{ range .Group.Fields -}}
			repeated {{ $.ViewNameOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
		{{- end }}
	}
{{- end }}
message SrcResp {
	{{ if .IsGroupByType }}Group{{ else }}repeated {{ .ViewName }}{{ end }} data = 1;
	{{ .BackendErrorName }} error = 2;
}`

const selectOneTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	rpc Select({{ .PayloadName }}.SrcId) returns ({{ .BackendErrorName }});
	{{ template "srcRPCRemove" . }}
}

message {{ .PayloadName }} {
	{{- template "srcResp" . }}
}`

// type is Element
const selectManyTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	{{ if .Add -}}
		rpc Select({{ .PayloadName }}.SrcId) returns ({{ .PayloadName }}.SelectResp);
	{{- end }}

	rpc Dst({{ .Empty }}) returns ({{ .PayloadName }}.DstResp);
	{{ if .Remove -}}
		rpc Remove({{ .PayloadName }}.DstId) returns ({{ .BackendErrorName }});
	{{- end }}
	{{ if .RemoveAll -}}
		rpc RemoveAll({{ .Empty }}) returns ({{ .BackendErrorName }});
	{{- end }}
}

message {{ .PayloadName }} {
	{{- template "srcResp" . }}
	message SelectResp {
		{{ .ElementName }} data = 1;
		{{ .BackendErrorName }} error = 2;
	}
	message DstResp {
		repeated {{ .ElementName }} data = 1;
		{{ .BackendErrorName }} error = 2;
	}
	{{ if .Remove -}}
		message DstId {
			{{ .ElementIdFieldType }} id = 1;
		}
	{{- end }}
}`
