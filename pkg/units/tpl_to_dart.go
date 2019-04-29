package units

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const dartUnitTplStr = `const Unit(
	{{ if .Dots }}dots: {{ template "cells" .Dots }},{{ end }}
	{{ if .Per }}per: {{ template "cells" .Per }},{{ end }}
)`

const dartCellsTplStr = `[
{{ range . }}
	Cell(
		{{ if .Exponent }}exponent: '{{ .Exponent | toString | superscript }}',{{ end }}
		{{ if .Prefix.Valid }}prefix: {{ template "prefix" .Prefix }},{{ end }}
		{{ if .GetCurrency.Valid }}unit: {{ template "currency" .GetCurrency }},{{ end }}
		{{ if .GetAtom.Valid }}unit: {{ template "atom" .GetAtom }},{{ end }}
		{{ if .GetSymbol }}unit: {{ template "symbol" .GetSymbol }},{{ end }}
	),
{{ end }}
]`

const dartPrefixTplStr = `const CellPrefix.{{ .Type }}(PrefixV1.{{ .Is }})`

const dartCurrencyTplStr = `const CurrencyUnit.{{ .Type }}(CurrencyV1.{{ .Is }})`
const dartAtomTplStr = `const AtomUnit.{{ .Type }}(AtomV1.{{ .Is }})`
const dartSymbolTplStr = `const SymbolUnit({{ . | dartRawStr }})`

var dartOutTpl *template.Template

func init() {
	funcs := genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
	)
	dartOutTpl = template.Must(template.New("dart_out").Funcs(funcs).Parse(dartUnitTplStr))
	template.Must(dartOutTpl.New("cells").Parse(dartCellsTplStr))
	template.Must(dartOutTpl.New("prefix").Parse(dartPrefixTplStr))
	template.Must(dartOutTpl.New("currency").Parse(dartCurrencyTplStr))
	template.Must(dartOutTpl.New("atom").Parse(dartAtomTplStr))
	template.Must(dartOutTpl.New("symbol").Parse(dartSymbolTplStr))
}

func (u *Unit) ToDartSource() (string, error) {
	var buf bytes.Buffer
	err := dartOutTpl.Execute(&buf, u)
	return buf.String(), err
}

func (p *CellPrefix) Valid() bool {
	return p.GetIs() != PrefixV1_noPrefix
}
func (c *CurrencyUnit) Valid() bool {
	return c.GetIs() != CurrencyV1_XXX
}
func (a *AtomUnit) Valid() bool {
	return a.GetIs() != AtomV1_noAtom
}
