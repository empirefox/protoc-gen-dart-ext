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
	{{- else if .ViewHasGroup }}
		{{ template "groupView" . }}
	{{- end }}
{{- end }}{{ end }}`

const groupViewTpl = `
message {{ .PayloadName }} {
	{{- if not .IdMessage -}}
		message SrcId {
			{{ range .ViewGroup.Fields -}}
				{{ $.IdTypeOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
			{{- end }}
		}
	{{- end }}
	message SrcResp {
		repeated {{ .ProtoRef }} data = 1;
		{{ .OperateError }} error = 2;
	}
	message Group {
		{{ range .ViewGroup.Fields -}}
			repeated {{ $.ViewNameOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
		{{- end }}
	}
	message GroupResp {
		Group data = 1;
		{{ .OperateError }} error = 2;
	}
	message SrcElement {
		oneof is {
			{{ range .ViewGroup.Fields -}}
				{{ $.TypeNameOfFieldType . }} {{ .Name }} = {{ .Descriptor.GetNumber }};
			{{- end }}
		}
	}
	message GetResp {
		SrcElement data = 1;
		{{ .OperateError }} error = 2;
	}
	message AddResp {
		{{ .ProtoRef }} data = 1;
		{{ .OperateError }} error = 2;
	}
	{{ if .ViewElement -}}
		message SrcIds {
			repeated {{ .ProtoRef }}.Id ids = 1;
		}
		message DstResp {
			repeated {{ .ProtoRef }}.Element data = 1;
			{{ .OperateError }} error = 2;
		}
		message SelectResp {
			{{ .ProtoRef }}.Element data = 1;
			{{ .SubmitError }} error = 2;
		}
	{{- end -}}
}`

const leafTpl = `
message {{ .PayloadName }} {
	{{- if .IdField -}}
		message SrcId {
			{{ .ProtoType .IdField }} id = {{ .IdField.Descriptor.GetNumber }};
		}
		{{- if .LeafHasSelectManyView -}}
			message SrcIds {
				repeated {{ .ProtoType .IdField }} ids = {{ .IdField.Descriptor.GetNumber }};
			}
		{{- end }}
	{{- end }}
	message GetResp {
		{{ .ProtoRef }} data = 1;
		{{ .OperateError }} error = 2;
	}
	{{ range .LeafViews -}}
		message {{ .Pgs.Name }} {
			message SrcResp {
				repeated {{ .ProtoRef }} data = 1;
				{{ .OperateError }} error = 2;
			}
			message AddResp {
				{{ .ProtoRef }} data = 1;
				{{ .SubmitError }} error = 2;
			}
			{{ if and .IsSelectManyView .ViewElement -}}
				message DstResp {
					repeated {{ .ProtoRef }}.Element data = 1;
					{{ .OperateError }} error = 2;
				}
				message SelectResp {
					{{ .ProtoRef }}.Element data = 1;
					{{ .SubmitError }} error = 2;
				}
			{{- end -}}
		}
	{{- end -}}
}

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
	rpc Get({{ .Empty }}) returns ({{ .GetResp }});
}`

const listTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	rpc Get({{ .SrcId }}) returns ({{ .GetResp }});
	{{ if .Add -}}
		rpc Add({{ .ListAddOrSaveReq }}) returns ({{ .AddResp }});
	{{- end }}
	{{ if .Save -}}
		rpc Save({{ .ListAddOrSaveReq }}) returns ({{ .SubmitError }});
	{{- end }}
	{{ if .Remove -}}
		rpc Remove({{ .SrcId }}) returns ({{ .OperateError }});
	{{- end }}
	{{ if .RemoveAll -}}
		rpc RemoveAll({{ .Empty }}) returns ({{ .OperateError }});
	{{- end -}}
}`

const srcRPCLineTpl = `rpc Src({{ .SrcReq }}) returns ({{ .SrcResp }});`

const selectOneTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	rpc Select({{ .SrcId }}) returns ({{ .SubmitError }});
	{{ if .Remove -}}
		rpc Remove({{ .Empty }}) returns ({{ .OperateError }});
	{{- end -}}
}`

// type is Element
const selectManyTpl = `
service {{ .ServiceName }} {
	{{ template "srcRPCLine" . }}
	{{ if .Add -}}
		rpc Select({{ .SrcId }}) returns ({{ .SelectResp }});
	{{- end }}
	{{ if .AddMany -}}
		rpc SelectMany({{ .SrcIds }}) returns ({{ .SelectManyResp }});
	{{- end }}

	rpc Dst({{ .DstReq }}) returns ({{ .DstResp }});
	{{ if .Remove -}}
		rpc Remove({{ .SrcId }}) returns ({{ .OperateError }});
	{{- end }}
	{{ if .RemoveAll -}}
		rpc RemoveAll({{ .Empty }}) returns ({{ .OperateError }});
	{{- end -}}
}`
