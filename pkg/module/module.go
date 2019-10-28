package module

import (
	"strings"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/l10n"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
	funk "github.com/thoas/go-funk"

	pgs "github.com/lyft/protoc-gen-star"
)

const (
	dartExtKey = "dart_ext"
	tplPrefix  = "// DO NOT EDIT. This is code generated via protoc-gen-dart-ext.\n"
)

type line struct {
	ext    string
	tplStr string
	tpl    *template.Template
}

type DartExtModule struct {
	*pgs.ModuleBase
	ctx       pgsgo.Context
	dart      *dart.Dart
	fileLines []*line
	pkgLine   *line
}

func New() *DartExtModule {
	return &DartExtModule{
		ModuleBase: &pgs.ModuleBase{},
		dart:       dart.NewDart(),
		fileLines:  make([]*line, 0, len(fileLines)),
	}
}

func (p *DartExtModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())

	params := p.ctx.Params().Str(dartExtKey)
	if l, ok := pkgLines[params]; ok {
		l.tpl = p.mustTpl(params, &l)
		p.pkgLine = &l
		return
	}

	tasks := strings.Split(params, "+")
	if len(tasks) == 0 {
		panic(dartExtKey + " must be set.\n" + supportedComment())
	}

	for _, task := range tasks {
		l, ok := fileLines[task]
		if !ok {
			panic(dartExtKey + " does not support: " + task + "\n" + supportedComment())
		}
		l.tpl = p.mustTpl(task, &l)
		p.fileLines = append(p.fileLines, &l)
	}
}

func (p *DartExtModule) mustTpl(task string, l *line) *template.Template {
	tplStr := tplPrefix + supportedComment() + l.tplStr
	return template.Must(template.New(task).Funcs(p.funcMap()).Parse(tplStr))
}

func (p *DartExtModule) funcMap() template.FuncMap {
	hasOneOf := func(f pgs.File) bool {
		for _, msg := range f.AllMessages() {
			if len(msg.OneOfs()) > 0 {
				return true
			}
		}
		return false
	}
	hasEnum := func(f pgs.File) (bool, error) {
		for _, e := range f.AllEnums() {
			l, err := l10n.FromEnum(e)
			if err != nil {
				return false, err
			}
			if !l.GetIgnore() {
				return true, nil
			}
		}
		return false, nil
	}
	return map[string]interface{}{
		"nameOf":       p.dart.NameOf,
		"pkgDartName":  p.dart.PackageDartName,
		"fileDartName": p.dart.FileDartName,
		"enumL10n":     l10n.FromEnum,
		"hasOneof":     hasOneOf,
		"hasOneofPkg": func(pkg pgs.Package) bool {
			for _, f := range pkg.Files() {
				if hasOneOf(f) {
					return true
				}
			}
			return false
		},
		"hasEnum": hasEnum,
		"hasEnumPkg": func(pkg pgs.Package) (bool, error) {
			for _, f := range pkg.Files() {
				if yes, err := hasEnum(f); yes || err != nil {
					return yes, err
				}
			}
			return false, nil
		},
	}
}

// Name satisfies the generator.Plugin interface.
func (p *DartExtModule) Name() string { return dartExtKey }

func (p *DartExtModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	if p.pkgLine != nil {
		for _, pkg := range targetPackages(targets) {
			p.genPkg(pkg)
		}
	} else {
		for _, t := range targets {
			p.genFile(t)
		}
	}

	return p.Artifacts()
}

func (p *DartExtModule) genFile(f pgs.File) {
	// For dart, use input path as output path.
	for _, l := range p.fileLines {
		name := f.InputPath().SetExt(l.ext)
		p.AddGeneratorTemplateFile(name.String(), l.tpl, f)
	}
}

func (p *DartExtModule) genPkg(pkg pgs.Package) {
	data := PkgData{Package: pkg, Ext: p.pkgLine.ext}
	p.AddGeneratorTemplateFile(data.OutputPath(), p.pkgLine.tpl, &data)
}

var fileLines = map[string]line{
	"defaults": {
		ext:    ".default.dart",
		tplStr: defaultsTplStr,
	},
}

var pkgLines = map[string]line{
	"pkg_l10n_base": {
		ext: ".l10n.base.dart",
		// TODO icons!!!
		tplStr: l10n_base,
	},
}

func supportedComment() string {
	comment := "// Supported `" + dartExtKey + "=` parameters:\n"
	comment += "//   1. file(any compose of): " + namesAdd(fileLines) + "\n"
	comment += "//   2. package(only one of): " + namesSingle(pkgLines) + "\n"
	return comment
}
func namesAdd(lines map[string]line) string {
	return "`" + strings.Join(funk.Keys(lines).([]string), "+") + "`"
}
func namesSingle(lines map[string]line) string {
	return "`" + strings.Join(funk.Keys(lines).([]string), "`, `") + "`"
}

func targetPackages(targets map[string]pgs.File) map[string]pgs.Package {
	pkgs := make(map[string]pgs.Package)
	for _, t := range targets {
		pkgs[t.Package().ProtoName().String()] = t.Package()
	}
	return pkgs
}
