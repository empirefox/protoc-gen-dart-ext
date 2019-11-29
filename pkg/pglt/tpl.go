package pglt

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const fileTplStr = `{{ renderJoin (gttToData .) "fileHeader" "fileBody+1" }}`

const fileHeaderTplStr = genshared.DartHead + `// ignore_for_file: non_constant_identifier_names
{{ .RenderImports }}`

const fileBodyTplStr = `
{{ range .Entities }}
	{{ template "delegate" .BaseArb }}

	{{ range .Arbs }}
		{{ template "arb" . }}
	{{ end }}

	/// last_modified: {{ .BaseArb.LastModified.Time }}
	{{ template "arbBase" .BaseArb }}
{{ end }}
`

const delegateTplStr = `{{ $L10nClass := l10nClass .Entity }}
class _{{ $L10nClass }}Delegate extends {{ .MaterialLib.AsDot "LocalizationsDelegate" }}<{{ $L10nClass }}> {
  const _{{ $L10nClass }}Delegate();

  @override
  bool isSupported({{ .MaterialLib.AsDot "Locale" }} locale) => kSupportedLanguages.contains(locale.languageCode);

  @override
  Future<{{ $L10nClass }}> load({{ .MaterialLib.AsDot "Locale" }} locale) => {{ .FoundationLib.AsDot "SynchronousFuture" }}<{{ $L10nClass }}>(_getTranslation(locale));

  @override
  bool shouldReload(_{{ $L10nClass }}Delegate old) => false;

  static final Set<String> kSupportedLanguages = {{ .CollectionLib.AsDot "HashSet" }}<String>.from(const <String>[
	{{ range .Delegate }}
		'{{ .Lang }}',
	{{ end }}
  ]);

  static {{ $L10nClass }} _getTranslation({{ .MaterialLib.AsDot "Locale" }} locale) {
    switch (locale.languageCode) {
		{{ range .Delegate }}{{ $lang := .Lang }}
		case '{{ $lang }}': {
			{{ if .Regions }}
				switch (locale.countryCode) {
					{{ range .Regions }}
					case '{{ . }}':
						return {{ $L10nClass }}{{ $lang.String | powerCamel }}{{ . | powerCamel }}();
					{{ end }}
				}
			{{ end }}
			return {{ $L10nClass }}{{ $lang.String | powerCamel }}{{ .Fallback | powerCamel }}();
		}
		{{ end }}
    }
    assert(false, 'getTranslation() called for unsupported locale "$locale"');
    return null;
  }
}
`

const arbBaseTplStr = `{{ $L10nClass := l10nClass .Entity }}
abstract class {{ $L10nClass }} {
	{{ if .Delegate }}
		static const delegate = _{{ $L10nClass }}Delegate();
	{{ end }}

	{{ $L10nClass }}();

	{{ range .Resources }}
		{{ template "resourceBase" . }}
	{{ end }}

	{{ range .InstanceClasses }}
		{{ .FullName }} {{ .Instance }};
	{{ end }}

	static {{ $L10nClass }} of({{ .MaterialLib.AsDot "BuildContext" }} context) =>
		{{ .MaterialLib.AsDot "Localizations" }}.of<{{ $L10nClass }}>(context, {{ $L10nClass }})
		  {{ range .InstanceClasses }}
		  	..{{ .Instance }} = {{ .FullName }}.of(context)
		  {{ end }}
		;
}
`

const arbTplStr = `{{ $L10nClass := l10nClass .Entity }}
class {{ $L10nClass }}{{ powerCamel .Locale.String }} extends {{ $L10nClass }} {
	{{ $L10nClass }}{{ powerCamel .Locale.String }}();

	{{ range .Resources }}
		{{ template "resource" (resource $.ImportManager .) }}
	{{ end }}
}
`

const resourceBaseTplStr = `
{{ if and .Attr .Attr.DartParams.JoinWithType }}
	String {{ .Id }}({{ .Attr.DartParams.JoinWithType }});
{{ else }}
	String get {{ .Id }};
{{ end }}
`

const resourceTplStr = `{{ $kSameWith := .SameWith }}
{{ if and .Attr .Attr.DartParams.JoinWithType }}
	String {{ .Id }}({{ .Attr.DartParams.JoinWithType }})
	{{ if $kSameWith }}
		=> {{ $kSameWith }}({{ .Attr.DartParams.JoinWithoutType }});
	{{ else }}
		{
			{{ template "root" (mfPluralNodeWithData .Arb.Locale .Value .) }}
		}
	{{ end }}
{{ else }}
	String get {{ .Id }} => {{ if $kSameWith }} {{ $kSameWith }} {{ else }} {{ dartRawStrOrNull .Value }} {{ end }};
{{ end }}
`

const rootTplStr = `
{{ if .Complexity.IsComplex }}
	final output = StringBuffer();
	{{ template "container" . }}
	return output.toString();
{{ else }}
	{{ template "container" . }}
{{ end }}
`

// Type: literal, var, select, selectordinal, plural
const nodeTplStr = `{{ render (node .) }}`

const containerTplStr = `
{{- range .Children }}
	{{ template "node" . -}}
{{- end -}}
`

const literalTplStr = `
{{ if .Complexity.IsComplex }}
	{{ range .Content }}
		{{ if . -}}
			output.write({{ dartRawStr . }});
		{{ else if $.Varname -}}
			output.write({{ $.Data.Attr.DartPlaceholderReplace $.Varname }});
		{{ else -}}
			output.writeCharCode(35);
		{{ end -}}
	{{ end -}}
{{ else }}
	{{ range .Content }}
		{{ if . -}}
			return {{ dartRawStrOrNull . }};
		{{ else if $.Varname -}}
			return {{ $.Data.Attr.DartPlaceholderReplace $.Varname }};
		{{ else -}}
			return '#';
		{{ end -}}
	{{ end }}
{{ end }}
`
const varTplStr = `output.write({{ $.Data.Attr.DartPlaceholderReplace $.Varname }});`

const selectTplStr = `
{{ if .Choices }}
switch ({{ .Data.Attr.DartPlaceholderReplace .Varname }}) {
	{{ range .Choices }}
	case {{ $.Data.Attr.DartSelectCaseCond $.Varname .Key }}:
		{{ template "node" .Node -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
	{{ end }}
	default:
		{{ template "node" .Other -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
}
{{ else }}
	{{ template "node" .Other -}}
{{ end -}}
`

const selectOrdinalTplStr = `
{{ if .Choices }}
{{ $kPluralLib := .Data.PluralLib .Varname }}
switch (
	{{- if (.Data.Attr.IsMatched .Varname) -}}
		{{ .Data.Attr.DartPlaceholderReplace .Varname }}
	{{- else -}}
		{{- matchFn $kPluralLib .Culture }}(
			{{ .Data.Attr.DartPlaceholderReplace .Varname }},
			true)
	{{- end -}}
	) {
	{{ range .Choices }}
	case {{ $kPluralLib.AsDot "Form" }}.{{ .Key }}:
		{{ template "node" .Node -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
	{{ end }}
	default:
		{{ template "node" .Other -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
}
{{ else }}
	{{ template "node" .Other -}}
{{ end -}}
`

const pluralTplStr = `
{{ if .Equals }}
	{{ template "equalsPlural" . -}}
{{ else }}
	{{ template "formedPlural" . -}}
{{ end }}
`

const equalsPluralTplStr = `
switch ({{ .Data.Attr.DartPlaceholderReplace .Varname }}) {
	{{ range .Equals }}
	case {{ .Key }}:
		{{ template "node" .Node -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
	{{ end }}
	default:
		{{ template "formedPlural" . -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
}
`

const formedPluralTplStr = `
{{ if .Choices }}
{{ $kPluralLib := .Data.PluralLib .Varname }}
switch (
	{{- if (.Data.Attr.IsMatched .Varname) -}}
		{{ .Data.Attr.DartPlaceholderReplace .Varname }}
	{{- else -}}
		{{- matchFn $kPluralLib .Culture }}(
			{{ .Data.Attr.DartPlaceholderReplace .Varname }}
			{{- if .Offset }}{{ .Offset | sub 0 | printf "%+d" }}{{ end }},
			false)
	{{- end -}}
	) {

	{{- range .Choices }}
	case {{ $kPluralLib.AsDot "Form" }}.{{ .Key }}:
		{{ template "node" .Node -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
	{{- end }}
	default:
		{{ template "node" .Other -}}
		{{ if $.Complexity.IsComplex }}break;{{ end }}
}
{{ else }}
	{{ template "node" .Other -}}
{{ end }}
`
