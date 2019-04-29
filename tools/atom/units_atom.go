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
	"github.com/empirefox/protoc-gen-dart-ext/pkg/units/atom"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

var (
	entity = &genshared.VersionEntity{
		LowerCamelCase: "atom",
		Version:        1,
	}
)

const protoTplStr = genshared.UnitsProtoHead + `
import "protos/l10n/l10n.proto";

// Part of:
// https://physics.nist.gov/cuu/Units/units.html
// https://physics.nist.gov/cuu/Units/outside.html
enum {{ EntityV }} {
  option (l10n.enumArb).ignore = true;
  noAtom = 0;

{{- range $idx, $e := .Atoms }}
  // {{ $e.PrintSymbol false }}
  {{ field $e }} = {{ add1 $idx }};
{{- end }}
}
`

const dartTplStr = genshared.DartHead + `
import './units.l10n.dart';

abstract class _Valuer {
	String of(UnitsLocalization l, bool p);
}

class _No{{ Entity }} implements _Valuer {
  const _No{{ Entity }}();
  String of(UnitsLocalization l, bool p) => '';
}

{{ range .Atoms }}
  class _{{ Field . }} implements _Valuer {
    const _{{ Field . }}();
    String of(UnitsLocalization l, bool p) =>
      {{ if .Plural }}
        p ? l?.{{ entity }}{{ FieldPlural . }} : l?.{{ entity }}{{ Field . }}
      {{ else }}
        l?.{{ entity }}{{ Field . }}
      {{ end }}
    ;
  }
{{ end }}

class {{ EntityV }} {
	static const no{{ Entity }} = const {{ EntityV }}._('', const _No{{ Entity }}());
	{{ range .Atoms }}
		{{ $symbol := .PrintSymbol false | dartRawStr }}
		static const {{ field . }} = const {{ EntityV }}._({{ $symbol }}, const _{{ Field . }}());
	{{ end }}

  final String symbol;
  final _Valuer _v;
  const {{ EntityV }}._(this.symbol, this._v);
  String l10n(UnitsLocalization l10n, bool plural) => _v.of(l10n, plural) ?? symbol;
}
`

const arbTplStr = `{{ toPrettyJson .Arb }}`

var (
	toField = func(a *atom.Atom) (s string) {
		if a.Symbol == "''" || a.Symbol == "'" {
			s = a.Name + "Angle"
		} else {
			s = a.Name
		}
		return util.PowerCamel(s)
	}
	toFieldPlural = func(a *atom.Atom) (s string) {
		if a.Symbol == "''" || a.Symbol == "'" {
			s = a.Plural + "Angle"
		} else {
			s = a.Plural
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
	toLowerFieldPlural = func(a *atom.Atom) (s string) {
		if a.Symbol == "''" || a.Symbol == "'" {
			s = a.Plural + "Angle"
		} else {
			s = a.Plural
		}
		return util.PowerLowerCamel(s)
	}

	sharedFuncs = genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		genshared.VersionEntityFuncs(entity),
		template.FuncMap{
			"field":       toLowerField,
			"fieldPlural": toLowerFieldPlural,
			"Field":       toField,
			"FieldPlural": toFieldPlural,
		},
	)

	protoTpl = template.Must(template.New("proto").Funcs(sharedFuncs).Parse(protoTplStr))
	dartTpl  = template.Must(template.New("dart").Funcs(sharedFuncs).Parse(dartTplStr))
	arbTpl   = template.Must(template.New("arb").Funcs(sharedFuncs).Parse(arbTplStr))
)

func main() {
	rs := genshared.NewGoTplRenderers(protoTpl, dartTpl, arbTpl)
	flag.Parse()
	rs.OpenAll()
	defer rs.Close()

	err := rs.Render(&Data{Atoms: atom.Atoms})
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type Data struct {
	Atoms []*atom.Atom
}

func (d Data) Arb() *arb.Arb {
	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       language.English,
		Entity:       entity.UpperCamelV(),
		Resources:    make([]*arb.ArbResource, 0, len(d.Atoms)*2),
	}

	for _, ua := range d.Atoms {
		ar := arb.NewResource(a, toLowerField(ua), ua.Name, &arb.ArbAttributes{
			Type:        "text",
			Description: "SI units (singular)",
		})
		a.Resources = append(a.Resources, ar)
		if ua.Plural != "" {
			ar = arb.NewResource(a, toLowerFieldPlural(ua), ua.Plural, &arb.ArbAttributes{
				Type:          "text",
				Description:   "SI units (plural), empty means no plural form",
				MaybeSameWith: toField(ua),
			})
			a.Resources = append(a.Resources, ar)
		}
	}
	return a
}
