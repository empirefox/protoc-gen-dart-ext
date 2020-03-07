package main

import (
	"flag"
	"log"
	"strings"
	"text/template"
	"time"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/units/atom"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/relvacode/iso8601"
)

var (
	entity = &genshared.VersionEntity{
		LowerCamelCase: "atom",
		Version:        1,
	}
)

const protoTplStr = genshared.UnitsProtoHead + `
import "pgde/zero/zero.proto";

// Part of:
// https://physics.nist.gov/cuu/Units/units.html
// https://physics.nist.gov/cuu/Units/outside.html
enum {{ Entity }} {
  option (pgde.zero.defaultNotSet) = true;

  noAtom = 0;

{{- range $idx, $e := .Atoms }}
  // {{ $e.PrintSymbol false }}
  {{ field $e }} = {{ add1 $idx }};
{{- end }}
}
`

const dartTplStr = genshared.DartHead + `
import '../plural.dart';
import '../l10n.dart';

abstract class _Valuer {
  String of(PgdeLocalizations l, Form p);
}

class _No{{ Entity }} implements _Valuer {
  const _No{{ Entity }}();
  String of(PgdeLocalizations l, Form p) => '';
}

{{ range .Atoms }}
  class _{{ Field . }} implements _Valuer {
    const _{{ Field . }}();
    String of(PgdeLocalizations l, Form p) => l.{{ entity }}{{ Field . }}(p);
  }
{{ end }}

class {{ Entity }} {
	static const no{{ Entity }} = const {{ Entity }}._('', const _No{{ Entity }}());
	{{ range .Atoms }}
	  {{ $symbol := .PrintSymbol false | dartRawStr }}
	  static const {{ field . }} = const {{ Entity }}._({{ $symbol }}, const _{{ Field . }}());
	{{ end }}

	static const $byNumber = <{{ Entity }}>[
	  no{{ Entity }},
	  {{- range .Atoms }}
		{{ field . }},
	  {{ end }}
	];

  final String symbol;
  final _Valuer _v;
  const {{ Entity }}._(this.symbol, this._v);
  const Atom.symbol(this.symbol) : _v = null;
  String l10n(PgdeLocalizations l10n, Form form) => l10n == null ? symbol : _v?.of(l10n, form) ?? symbol;
}
`

const arbTplStr = `{{ .Arb | toPrettyJson }}`

var (
	toField = func(a *atom.Atom) (s string) {
		if a.Symbol == "''" || a.Symbol == "'" {
			s = a.Name + "Angle"
		} else {
			s = a.Name
		}
		return util.PowerCamel(s)
	}
	toLowerField = func(a *atom.Atom) (s string) {
		if a.Symbol == "''" || a.Symbol == "'" {
			s = a.Name + "Angle"
		} else {
			s = a.Name
		}
		return util.PowerLowerCamel(s)
	}

	sharedFuncs = genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		genshared.VersionEntityFuncs(entity),
		template.FuncMap{
			"field": toLowerField,
			"Field": toField,
		},
	)

	protoTpl = template.Must(template.New("proto").Funcs(sharedFuncs).Parse(protoTplStr))
	dartTpl  = template.Must(template.New("dart").Funcs(sharedFuncs).Parse(dartTplStr))
	arbTpl   = template.Must(template.New("arb").Funcs(sharedFuncs).Parse(arbTplStr))
)

func main() {
	gr := genshared.NewGoTplGroupRenderer(protoTpl, dartTpl, arbTpl)

	flag.Parse()

	gr.Open()
	defer gr.Close()

	err := gr.Render(&Data{Atoms: atom.Atoms})
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type Data struct {
	Atoms []*atom.Atom
}

func (d *Data) Arb() *arb.Arb {
	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Entity:       entity.UpperCamel(),
		Resources:    make([]*arb.ArbResource, 0, len(d.Atoms)),
	}

	for _, ua := range d.Atoms {
		ar := arb.NewResource(a, toLowerField(ua), toICU(ua), &arb.ArbAttributes{
			Type:        "text",
			Description: "SI unit",
			Placeholders: arb.ArbPlaceholders{
				&arb.ArbPlaceholder{
					Name:      "form",
					LangInfos: arb.ArbLangInfos{&arb.DartMatchedPluralLangInfo},
				},
			},
		})
		a.Resources = append(a.Resources, ar)
	}
	return a
}

var forms = []string{"zero", "one", "two", "few", "many", "other"}

func toICU(ua *atom.Atom) string {
	var builder strings.Builder
	builder.WriteString("{form, plural, ")
	for _, c := range forms {
		builder.WriteString(c)
		builder.WriteByte('{')
		if c == "one" {
			builder.WriteString(ua.Name)
		} else {
			builder.WriteString(ua.GetPlural())
		}
		builder.WriteByte('}')
	}
	// other
	builder.WriteString("other{")
	builder.WriteString(ua.GetPlural())
	builder.WriteByte('}')

	builder.WriteByte('}')
	return builder.String()
}
