package main

import (
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/messageformat"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const dartOutTplStr = genshared.DartHead + `
{{ if .BaseArb.Delegate }}
	import 'dart:collection' as {{ .EmailValidatorFile.As }} show HashSet;
	import 'package:flutter/foundation.dart' as {{ .PgdeFile.As }} show SynchronousFuture;
	import 'package:flutter/material.dart' as {{ .L10nFile.As }} show BuildContext, Locale, Localizations, LocalizationsDelegate;
{{ end }}

{{ .BaseArb.RenderImports }}

{{ if .BaseArb.Delegate }}
	{{ template "delegate" .BaseArb }}
{{ end }}

/// last_modified: {{ .BaseArb.LastModified.Time }}
{{ template "arbBase" .BaseArb }}

{{ range .Arbs }}
	{{ template "arb" . }}
{{ end }}
`

const delegateTplStr = `{{ $Entity := powerCamel .Entity }}
class _{{ $Entity }}LocalizationDelegate extends {{ .L10nFile.As }}.LocalizationsDelegate<{{ $Entity }}Localization> {
  const _{{ $Entity }}LocalizationDelegate();

  @override
  bool isSupported({{ .L10nFile.As }}.Locale locale) => kSupportedLanguages.contains(locale.languageCode);

  @override
  Future<{{ $Entity }}Localization> load({{ .L10nFile.As }}.Locale locale) => {{ .PgdeFile.As }}.SynchronousFuture<{{ $Entity }}Localization>(_getTranslation(locale));

  @override
  bool shouldReload(_{{ $Entity }}LocalizationDelegate old) => false;

  static final Set<String> kSupportedLanguages = {{ .EmailValidatorFile.As }}.HashSet<String>.from(const <String>[
	{{ range .Delegate }}{{ $lang := .Lang }}
		{{ if .Fallback }}
			'{{ $lang }}',
		{{ else }}
			{{ range .Regions }}
				'{{ $lang }}-{{ . }}',
			{{ end }}
		{{ end }}
	{{ end }}
  ]);

  static {{ $Entity }}Localization _getTranslation({{ .L10nFile.As }}.Locale locale) {
    switch (locale.languageCode) {
		{{ range .Delegate }}{{ $lang := .Lang }}
		case '{{ $lang }}': {
			{{ if .Regions }}
				switch (locale.countryCode) {
					{{ range .Regions }}
					case '{{ . }}':
						return {{ $Entity }}Localization{{ $lang.String | powerCamel }}{{ . | powerCamel }}();
					{{ end }}
				}
			{{ end }}
			{{ if .Fallback }}
				return {{ $Entity }}Localization{{ $lang.String | powerCamel }}();
			{{ end }}
		}
		{{ end }}
    }
    assert(false, 'getTranslation() called for unsupported locale "$locale"');
    return null;
  }
}
`

const arbBaseTplStr = `{{ $Entity := powerCamel .Entity }}
abstract class {{ $Entity }}Localization {
	{{ if .Delegate }}
		static const delegate = _{{ $Entity }}LocalizationDelegate();
	{{ end }}

	{{ .Const }} {{ $Entity }}Localization();

	{{ range .Resources }}
		{{ template "resourceBase" . }}
	{{ end }}

	{{ range .InstanceClasses }}
		{{ .FullName }} {{ .Instance }};
	{{ end }}

	static {{ $Entity }}Localization of({{ .L10nFile.As }}.BuildContext context) =>
		{{ .L10nFile.As }}.Localizations.of<{{ $Entity }}Localization>(context, {{ $Entity }}Localization)
		  {{ range .InstanceClasses }}
		  	..{{ .Instance }} = {{ .FullName }}.of(context)
		  {{ end }}
		;
}
`

const arbTplStr = `{{ $Entity := powerCamel .Entity }}
class {{ $Entity }}Localization{{ powerCamel .Locale.String }} extends {{ $Entity }}Localization {
	{{ .Const }} {{ $Entity }}Localization{{ powerCamel .Locale.String }}();

	{{ range .Resources }}
		{{ template "resource" . }}
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
	String {{ .Id }}({{ .Attr.DartParams.JoinWithType }});
	{{ if $kSameWith }}
		=> {{ $kSameWith }}({{ .Attr.DartParams.JoinWithoutType }});
	{{ else }}
		{
			{{ template "root" (mfPluralNodeWithData .Arb.Locale .Value .) }}
		}
	{{ end }}
{{ else }}
	String get {{ .Id }} => {{ if $kSameWith }} {{ $kSameWith }} {{ else }} {{ dartRawStr .Value }} {{ end }};
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
const nodeTplStr = `{{ renderNode .Type . }}`

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
			return {{ dartRawStr . }};
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
{{ $kPluralAs := .Data.Attr.DartParamImport .Varname | getAs }}
switch ({{ $kPluralAs }}.match{{ powerCamel .Culture }}(
	{{ .Data.Attr.DartPlaceholderReplace .Varname }}, true)) {
	{{ range .Choices }}
	case {{ $kPluralAs }}.Form.{{ .Key }}:
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
{{ $kPluralAs := .Data.Attr.DartParamImport .Varname | getAs }}
switch ({{ $kPluralAs }}.match{{ powerCamel .Culture }}(
	{{ .Data.Attr.DartPlaceholderReplace .Varname }}
	{{- if .Offset }}{{ .Offset | sub 0 | printf "%+d" }}{{ end }},
	false)) {

	{{- range .Choices }}
	case {{ $kPluralAs }}.Form.{{ .Key }}:
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

const dart_out = "dart_out"

var (
	dartOutTpl = template.New(dart_out)
	im         = dart.NewDefaultImportManager()
)

func init() {
	funcs := genshared.JoinFuncs(messageformat.Funcs,
		dart.Funcs,
		sprig.HermeticTxtFuncMap(),
		genshared.FuncsWithRender("renderNode", dartOutTpl),
		template.FuncMap{
			"getAs": im.GetAs,
		})
	template.Must(dartOutTpl.Funcs(funcs).Parse(dartOutTplStr))
	template.Must(dartOutTpl.New("delegate").Parse(delegateTplStr))
	template.Must(dartOutTpl.New("arbBase").Parse(arbBaseTplStr))
	template.Must(dartOutTpl.New("arb").Parse(arbTplStr))
	template.Must(dartOutTpl.New("resource").Parse(resourceTplStr))
	template.Must(dartOutTpl.New("resourceBase").Parse(resourceBaseTplStr))
	template.Must(dartOutTpl.New("root").Parse(rootTplStr))
	template.Must(dartOutTpl.New("node").Parse(nodeTplStr))
	template.Must(dartOutTpl.New("container").Parse(containerTplStr))
	template.Must(dartOutTpl.New("literal").Parse(literalTplStr))
	template.Must(dartOutTpl.New("var").Parse(varTplStr))
	template.Must(dartOutTpl.New("select").Parse(selectTplStr))
	template.Must(dartOutTpl.New("selectordinal").Parse(selectOrdinalTplStr))
	template.Must(dartOutTpl.New("plural").Parse(pluralTplStr))
	template.Must(dartOutTpl.New("equalsPlural").Parse(equalsPluralTplStr))
	template.Must(dartOutTpl.New("formedPlural").Parse(formedPluralTplStr))
}
