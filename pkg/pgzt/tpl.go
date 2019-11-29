package pgzt

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const fileTpl = `{{ renderJoin . "fileHeader" "fileBody+1" }}`

const fileHeaderTpl = genshared.DartHead + `
/// Code generated by protoc-gen-dart-ext. DO NOT EDIT.
/// source: {{ .Pgs.InputPath }}
// ignore_for_file: non_constant_identifier_names,unnecessary_brace_in_string_interps,unused_local_variable

{{ .Zeros.ImportManager.RenderImports }}
`

const fileBodyTpl = `
{{ range .Messages }}
	{{ if not .Zero.Disabled }}
		{{ template "msg" .Zero }}
	{{ end }}
{{ end }}
`

const msgTpl = `
/// Zero for [{{ .FullPbClass }}] protobuf objects.
const {{ .Names.ZeroName }} = const _{{ .DartName }}();
class _{{ .DartName }} implements {{ .PgdeFile.AsDot "ZeroAction" }}<{{ .FullPbClass }}> {
	const _{{ .DartName }}();

	{{ range .NonOneOfFields }}
		{{ renderConstants (.Zero.ForAction "all") }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOfConst" (.Zero.ForAction "all") }}
	{{ end }}

	{{ range .Actions }}
		{{ template "action" ($.ForAction .) }}
	{{ end }}
}
`

const actionTpl = `
@override
void {{ .Action }}({{ .FullPbClass }} proto) {
	{{- range .NonOneOfFields }}
		{{ template "field" (.Zero.ForAction $.Action) }}
	{{ end }}
	{{ range .OneOfs }}
		{{ template "oneOf" (.Zero.ForAction $.Action) }}
	{{- end -}}
}
`

const oneOfConstTpl = `
{{ if .Default.AddBuild }}
	{{ renderConstants (.Default.ForAction .Action) }}
{{ end }}`
const oneOfTpl = `
{{ if .Default.AddBuild }}
	if (proto.{{ .Names.WhichOneofMethodName }}() == {{ .FullPbWhichEnum }}.notSet) {
		{{ template "field" (.Default.ForAction .Action) }}
	}
{{ end }}`

const fieldConstTpl = `{{ if and .AddBuild }}{{ renderConstants . }}{{ end }}`
const fieldTpl = `
{{ if and .AddBuild }}
	{{ render . }}
{{ else }}
	// name={{ .DartName }}, AddBuild={{ .AddBuild }}
{{ end }}`

const mapConstTpl = `
{{ if .SetTo }}
	{{ .TypeForField }} get {{ .Member }} => 
	{{ if .RangeElemIsMsg }}
		_{{ .Member }}.map((k, v) => MapEntry(k, v.clone()));
		static final {{ .TypeForField }} _{{ .Member }} =
	{{ end }}
	<{{ .TypeForRangeKey }}, {{ .TypeForRangeElem }}>{
		{{ range .SetTo }}
			{{ .Literal .Key }}: {{ .Literal .Value }},
		{{ end }}
	};
{{ end }}
{{ if not .Rules.GetMessage.GetSkip }}
	{{ renderConstants (.Elem .DartName "v" (.DartName.Index "k")) }}
{{ end }}`
const mapTpl = `
{{ if or .SetTo (not .Rules.GetMessage.GetSkip) }}
	final {{ .DartName }} = {{ .Getter }};
{{ end }}
{{ if .SetTo }}
	if ({{ .DartName }}.isEmpty) {{ .DartName }}.addAll({{ .Member }});
{{ end }}
{{ if not .Rules.GetMessage.GetSkip }}
	{{ .DartName }}.forEach((k, v) => {
		{{ template "field" (.Elem .DartName "v" (.DartName.Index "k")) }}
	});
{{ end }}`

const repeatedConstTpl = `
{{ if .SetTo }}
	{{ .TypeForField }} get {{ .Member }} => 
	{{ if .RangeElemIsMsg }}
		_{{ .Member }}.map((v) => v.clone());
		static final {{ .TypeForField }} _{{ .Member }} =
	{{ end }}
	<{{ .TypeForRangeElem }}>[
		{{ range .SetTo }}
			{{ .Literal . }},
		{{ end }}
	];
{{ end }}
{{ if not .Rules.GetMessage.GetSkip }}
	{{ renderConstants (.Elem .DartName (.DartName.Index "idx") "") }}
{{ end }}`
const repeatedTpl = `
{{ if or .SetTo (not .Rules.GetMessage.GetSkip) }}
	final {{ .DartName }} = {{ .Getter }};
{{ end }}
{{ if .SetTo }}
	if ({{ .DartName }}.isEmpty) {{ .DartName }}.setAll(0, {{ .Member }});
{{ end }}
{{ if not .Rules.GetMessage.GetSkip }}
	for (var idx = 0; idx < {{ .DartName }}.length; idx++) {
		{{ template "field" (.Elem .DartName (.DartName.Index "idx") "") }}
	}
{{ end }}`

const messageConstTpl = `
{{ if .SetTo }}
	void {{ .Member }}({{ .MessageFullPbClass }} p) => p.mergeFromMessage(_{{ .Member }});
	static final {{ .MessageFullPbClass }} _{{ .Member }} = {{ .Literal .SetTo }};
{{ end }}`
const messageTpl = `
{{ if .SetTo }}
	if ({{ .MessageHasMethodOr }} {{ .Getter }} == {{ .MessageFullPbClass }}.getDefault()) {
		{{ .Member }}({{ .Setter }});
	}
{{ end }}
{{ if not .Rules.GetMessage.GetSkip }}
	{{- with .MessageFullZeroAccessor }}
		{{ . }}.{{ $.Action }}(
			{{- if $.Pgs.Type.IsEmbed -}}
				{{- $.Owner }}.{{ $.Names.EnsureMethodName -}}()
			{{- else -}}
				{{ $.Getter -}}
			{{- end -}}
		);
	{{- end }}
{{ end }}`

const wktConstTpl = `{{ .MessageFullPbClass }} get {{ .Member }} => {{ .Literal .SetTo }};`
const wktTpl = `if ({{ .MessageHasMethodOr }} {{ .WktFieldsIsZero }}) {{ .Setter }} = {{ .Member }};`

const evalTimeConstTpl = `
{{ if .SetTo }}
	{{ .MessageFullPbClass }} get {{ .Member }}() {
		final t = DateTime.now().microsecondsSinceEpoch 
			{{- if .SetTo -}}
				{{ printf "%+d" .SetTo.Microseconds }}
			{{- end -}}
		;
		return {{ .TimestampType }}..seconds=t~/Duration.microsecondsPerSecond
			..nanos=t%Duration.microsecondsPerSecond;
	}
{{ end }}`
const evalTimeTpl = wktTpl

const int64ConstTpl = `static const {{ .Member }} = const {{ .Literal .SetTo }};`
const int64Tpl = `if ({{ .Getter }}.isZero) {{ .Setter }} = {{ .Member }};`

const stringConstTpl = `static const {{ .Member }} = {{ .Literal .SetTo }};`
const stringTpl = `if ({{ .Getter }}.isEmpty) {{ .Setter }} = {{ .Member }};`

const bytesConstTpl = `List<int> get {{ .Member }} => {{ .Literal .SetTo }};`
const bytesTpl = stringTpl

const enumTpl = `if ({{ .Getter }} == {{ .EnumFullPbDefault }}) {{ .Setter }} = {{ .Literal .SetTo }};`
const boolTpl = `if (!{{ .Getter }}) {{ .Setter }} = {{ .Literal .SetTo }};`
const numTpl = `if ({{ .Getter }} == 0) {{ .Setter }} = {{ .Literal .SetTo }};`
const noneTpl = `// no zero rules for {{ .Getter }}
`