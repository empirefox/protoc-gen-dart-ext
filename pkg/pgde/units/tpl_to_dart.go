package units

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const dartUnitTplStr = `const Unit(
	{{ if .Show }}show: {{ template "unit_show" .Show }},{{ end }}
	{{ if .Dots }}dots: {{ template "unit_cells" .Dots }},{{ end }}
	{{ if .Per }}per: {{ template "unit_cells" .Per }},{{ end }}
)`

const dartShowTplStr = `const Show(
	{{ if .Off }}off: true,{{ end }}
	{{ if .Prefix }}prefix: Show.{{ .Prefix.String }}Prefix,{{ end }}
	{{ if .Atom }}atom: Show.{{ .Atom.String }},{{ end }}
)`

const dartCellsTplStr = `[
{{ range . }}
	Cell(
		{{ if .Exponent }}exponent: '{{ .Exponent | toString | superscript }}',{{ end }}
		{{ if .Prefix.Valid }}prefix: Prefix.{{ .Prefix.String }},{{ end }}
		{{ if .GetAtom.Valid }}atom: Atom.{{ .GetAtom.String }},{{ end }}
		{{ if .GetSymbol }}atom: const Atom.symbol({{ .GetSymbol | dartRawStr }}){{ end }}
	),
{{ end }}
]`

var dartOutTpl *template.Template

func init() {
	funcs := genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
	)
	dartOutTpl = template.Must(template.New("unit").Funcs(funcs).Parse(dartUnitTplStr))
	template.Must(dartOutTpl.New("unit_show").Parse(dartShowTplStr))
	template.Must(dartOutTpl.New("unit_cells").Parse(dartCellsTplStr))
}

func (u *Unit) ToDartSource() (string, error) {
	var buf bytes.Buffer
	err := dartOutTpl.Execute(&buf, u)
	return buf.String(), err
}

func (p Prefix) Valid() bool { return p != Prefix_noPrefix }
func (a Atom) Valid() bool   { return a != Atom_noAtom }
