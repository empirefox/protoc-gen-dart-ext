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
		{{ if .Prefix.Valid }}prefix: PrefixV1.{{ .Prefix.String }},{{ end }}
		{{ if .GetAtom.Valid }}atom: AtomV1.{{ .GetAtom.String }},{{ end }}
		{{ if .GetSymbol }}atom: const AtomV1.symbol({{ .GetSymbol | dartRawStr }}){{ end }}
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

func (p PrefixV1) Valid() bool { return p != PrefixV1_noPrefix }
func (a AtomV1) Valid() bool   { return a != AtomV1_noAtom }
