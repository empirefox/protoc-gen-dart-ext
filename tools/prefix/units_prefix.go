package main

import (
	"flag"
	"log"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/units/prefix"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

var (
	entity = &genshared.VersionEntity{
		LowerCamelCase: "prefix",
		Version:        1,
	}
)

const protoTplStr = genshared.UnitsProtoHead + `
import "pgde/zero/zero.proto";

// https://physics.nist.gov/cuu/Units/prefixes.html
// https://physics.nist.gov/cuu/Units/binary.html
enum {{ Entity }} {
  option (pgde.zero.defaultNotSet) = true;

  no{{ Entity }} = 0;

  // SI
{{- range $idx, $prefix := .SIPrefixes }}
  // {{ $prefix.Symbol }}, {{ $prefix.Base }}^{{ $prefix.Exponent }}
  {{ $prefix.Name }} = {{ add1 $idx }};
{{- end }}

  // binary
{{- range $idx, $prefix := .BinaryPrefixes }}
  // {{ $prefix.Symbol }}, {{ $prefix.Base }}^{{ $prefix.Exponent }}
  {{ $prefix.Name }} = {{ add $idx (len $.SIPrefixes) 1 }};
{{- end }}
}
`

const dartTplStr = genshared.DartHead + `
import '../l10n.dart';

abstract class _Valuer {
  String of(PgdeLocalizations l);
}

class _No{{ Entity }} implements _Valuer {
  const _No{{ Entity }}();
  String of(PgdeLocalizations l) => '';
}

{{- range .AllPrefix }}
  class _{{ .Name | powerCamel }} implements _Valuer {
    const _{{ .Name | powerCamel }}();
    String of(PgdeLocalizations l) => l.{{ entity }}{{ .Name | powerCamel }};
  }
{{- end }}

class {{ Entity }} {
  static const no{{ Entity }} = const {{ Entity }}._('', 1, 1, const _No{{ Entity }}());
  {{- range .AllPrefix }}
	static const {{ .Name }} = const {{ Entity }}._({{ dartRawStr .Symbol }}, {{ .Base }}, {{ .Exponent }}, const _{{ .Name | powerCamel }}());
  {{- end }}

  static const $byNumber = <{{ Entity }}>[
	no{{ Entity }},
	{{- range .AllPrefix }}
		{{ .Name }},
	{{- end }}
  ];

  final String symbol;
  final int base;
  final int exponent;
  final _Valuer _v;
  const {{ Entity }}._(this.symbol, this.base, this.exponent, this._v);
  String l10n(PgdeLocalizations l10n) => _v.of(l10n) ?? symbol;
}
`

const arbTplStr = `{{ toPrettyJson .Arb }}`

var (
	sharedFuncs = genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		genshared.VersionEntityFuncs(entity),
	)

	protoTpl = template.Must(template.New("proto").Funcs(sharedFuncs).Parse(protoTplStr))
	dartTpl  = template.Must(template.New("dart").Funcs(sharedFuncs).Parse(dartTplStr))
	arbTpl   = template.Must(template.New("arb").Funcs(sharedFuncs).Parse(arbTplStr))
)

func main() {
	rs := genshared.NewGoTplGroupRenderer(protoTpl, dartTpl, arbTpl)
	flag.Parse()
	rs.Open()
	defer rs.Close()

	err := rs.Render(&Data{
		SIPrefixes:     prefix.SIPrefixes,
		BinaryPrefixes: prefix.BinaryPrefixes,
		AllPrefix:      prefix.AllPrefix(),
	})
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type Data struct {
	SIPrefixes     []*prefix.Prefix
	BinaryPrefixes []*prefix.Prefix
	AllPrefix      []*prefix.Prefix
}

func (d Data) Arb() *arb.Arb {
	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       language.English,
		Entity:       entity.UpperCamel(),
		Resources:    make([]*arb.ArbResource, len(d.AllPrefix)),
	}

	for i, ua := range d.AllPrefix {
		a.Resources[i] = arb.NewResource(a, ua.Name, ua.Name, &arb.ArbAttributes{
			Type:        "text",
			Description: "SI units prefix",
		})
	}
	return a
}
