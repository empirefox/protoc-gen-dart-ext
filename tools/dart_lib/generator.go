package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const goTplStr = genshared.DartHead + `
package dart

{{ range . }}
	{{ template "lib" . }}
{{ end }}`

const libTplStr = `
type {{ .Name }} struct {
	Library
	{{ range .AllExports }}
		_{{ . }}  Qualifier
	{{ end }}
}

func New{{ .Name }}(im *ImportManager) *{{ .Name }} {
	return &{{ .Name }}{
		Library: Library{
			im:         im,
			importPath: "{{ .ImportPath }}",
		},
	}
}

{{ range .AllExports }}
	func (m *{{ $.Name }}) {{ $.MethodName . }}() Qualifier {
		if m._{{ . }} == "" {
			m._{{ . }} = m.File().AsDot("{{ . }}")
		}
		return m._{{ . }}
	}
{{ end }}
`

var (
	goTpl = template.Must(template.New("go").Parse(goTplStr))
)

func init() {
	template.Must(goTpl.New("lib").Parse(libTplStr))
}

func main() {
	rs := genshared.NewGoTplGroupRenderer(goTpl)
	flag.Parse()
	rs.Open()
	defer rs.Close()

	err := rs.Render(data)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type Lib struct {
	Name       Qualifier
	ImportPath string
	Presets    []Qualifier
	Exports    []Qualifier
	usedNames  UsedNames
	allExports []Qualifier
}

func (lib *Lib) AllExports() []Qualifier {
	if lib.allExports == nil {
		lib.allExports = append(lib.Presets, lib.Exports...)
	}
	return lib.allExports
}

func (lib *Lib) MethodName(name Qualifier) Qualifier {
	return lib.methodName(name.ToCamel())
}

func (lib *Lib) methodName(name Qualifier) Qualifier {
	if lib.usedNames.Used(name) {
		return lib.methodName(name + "_")
	}
	lib.usedNames.Add(name)
	return name
}

var builtinNames = map[Qualifier]bool{
	"Library":    true,
	"ImportPath": true,
	"File":       true,
}

var data = []*Lib{
	flutterLib("animation"),
	flutterLib("foundation"),
	flutterLib("material"),
}

func flutterLib(name Qualifier) *Lib {
	return &Lib{
		Name:       name.ToCamel(),
		ImportPath: fmt.Sprintf("package:flutter/%s.dart", name),
		Presets:    nil,
		Exports:    loadExports(os.ExpandEnv(fmt.Sprintf("$MAKEFILE_DIR/pkg/dart/%s.export.json", name))),
		usedNames:  UsedNames{nil, builtinNames},
	}
}

func loadExports(js string) (e []Qualifier) {
	b, err := ioutil.ReadFile(js)
	if err != nil {
		log.Fatalf("ReadFile: %v", err)
	}

	if err = json.Unmarshal(b, &e); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return
}
