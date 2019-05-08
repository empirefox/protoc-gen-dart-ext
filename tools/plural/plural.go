package main

import (
	"flag"
	"log"
	"text/template"

	"github.com/Masterminds/sprig"

	"github.com/empirefox/makeplural/plural"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

// plural templates

const bodyTplStr = `
  {{- if .NeedFinvtw }}
    {{- if .P.Use }}
    final p = w == 0;
    {{- end }}
  {{- else }}
    {{- if .N.Use }}
    final n = value.abs();
    {{- end }}

    {{- if .I.Use }}
      {{- if .N.Use }}
      final i = n.toInt();
      {{- else }}
      final i = value.abs().toInt();
      {{- end }}
    {{- end }}

    {{- if .P.Use }}
      {{- if .I.Use }}
      final p = i == n;
      {{- else }}
      final p = n.toInt() == n;
      {{- end }}
    {{- end }}
  {{- end }}

  {{- range .Vars }}
    final {{ .Name }} = {{ .Symbol.Name }} % {{ .Mod }};
  {{- end }}

  if (ordinal) {
    {{- if .HasOrdinal }}
      {{ template "cases" .Ordinal }}
    {{- else }}
      return Form.other;
    {{- end }}
  }

  {{ template "cases" .Cardinal }}
`

const casesTplStr = `
{{- range . }}
  if ({{ .Cond }}) return Form.{{ .Form }};
{{- end }}
  return Form.other;
`

const ruleTplStr = `
{{- if .NeedFinvtw }}
  return finvtw(value, (f, i, n, v, t, w) {
    {{ template "body" . }}
  });
{{- else }}
  {{ template "body" . }}
{{- end }}
`

const pluralTplStr = genshared.DartHead + `
library plural;

import './finvtw.dart';

enum Form { other, zero, one, two, few, many }

typedef PluralFunc = Form Function(num value, bool ordinal);

{{- range .Cultures }}
Form match{{ .Langs | join "_" | powerCamel }}(num value, bool ordinal) {
  {{ template "rule" . }}
}
{{- end }}

Form matchOther(num value, bool ordinal) => Form.other;

final Map<String, PluralFunc> rules = {
  {{- range .Cultures }}{{ $func := .Langs | join "_" | powerCamel }}
    {{- range .Langs }}
    "{{ . }}": match{{ $func }},
    {{- end }}
  {{- end }}
  {{- range .Others }}
  "{{ . }}": matchOther,
  {{- end }}
};
`

// plural_test templates

const cardinalTplStr = `
{{- range .Integers }}
  _testNamedKey(fn, {{ . }}, Form.{{ $.Expected }}, 'fn({{ . }}, false)', false);
{{- end }}
{{- range .Decimals }}
  _testNamedKey(fn, {{ . }}, Form.{{ $.Expected }}, 'fn({{ . }}, false)', false);
{{- end }}
`

const ordinalTplStr = `
{{- range .Integers }}
  _testNamedKey(fn, {{ . }}, Form.{{ $.Expected }}, 'fn({{ . }}, true)', true);
{{- end }}
{{- range .Decimals }}
  _testNamedKey(fn, {{ . }}, Form.{{ $.Expected }}, 'fn({{ . }}, true)', true);
{{- end }}
`

const pluralTestTplStr = genshared.DartHead + `
import "package:test/test.dart";
import '../lib/plural/plural.dart';

void _testNamedKey(
    PluralFunc fn, num input, Form expected, String name, bool ordinal) {
  test('plural func return right Form', () {
    expect(fn(input, ordinal), equals(expected),
        reason: '$name, ordinal=$ordinal');
  });
}

void main() {
  {{- range .Cultures }}{{ $tests := .Tests }}
    {{- range .Langs }}
      group('{{ . }}', () {
        final fn = rules['{{ . }}'];
        {{- range $tests.Cardinal }}
          {{ template "cardinal" . }}
        {{- end }}
        {{- range $tests.Ordinal }}
          {{ template "ordinal" . }}
        {{- end }}
      });
    {{- end }}
  {{- end }}
}
`

var pluralTpl *template.Template
var pluralTestTpl *template.Template

func init() {
	pluralTpl = template.Must(template.New("dart").
		Funcs(sprig.HermeticTxtFuncMap()).
		Funcs(genshared.Funcs).
		Parse(pluralTplStr))
	template.Must(pluralTpl.New("rule").Parse(ruleTplStr))
	template.Must(pluralTpl.New("body").Parse(bodyTplStr))
	template.Must(pluralTpl.New("cases").Parse(casesTplStr))

	pluralTestTpl = template.Must(template.New("dart_test").Parse(pluralTestTplStr))
	template.Must(pluralTestTpl.New("cardinal").Parse(cardinalTplStr))
	template.Must(pluralTestTpl.New("ordinal").Parse(ordinalTplStr))
}

func main() {
	rs := genshared.NewGoTplGroupRenderer(pluralTpl, pluralTestTpl)
	flag.Parse()
	rs.Open()
	defer rs.Close()

	err := rs.Render(&plural.Info)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}
