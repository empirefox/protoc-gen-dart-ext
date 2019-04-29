package main

import (
	"sort"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const exportsOutTplStr = `// ignore_for_file: lines_longer_than_80_chars
import 'package:pgde/exports/register.dart';

void register() {
  {{- range .ExportsOut }}
  registerPkgBytes('{{ .Name }}', const <int>[{{ .Package | proto | bytes }}]);
  {{- end }}
}
`

var (
	exportsOutTpl = &genshared.Template{
		Template: template.Must(template.
			New("exports_out").
			Funcs(genshared.Funcs).
			Parse(exportsOutTplStr)),
		IgnoreIfData: func(v interface{}) bool {
			return v == nil || len(v.(*Data).ExportsOut) == 0
		},
	}
)

type ExportsPackage struct {
	*exports.Package
	Name string
}

type ExportsOut []*ExportsPackage

func NewSingleEntitiesExportsOut(pkgToNty map[string]*exports.Entity) ExportsOut {
	if len(pkgToNty) == 0 {
		return nil
	}
	data := make(ExportsOut, len(pkgToNty))
	i := 0
	for pkgName, nty := range pkgToNty {
		sort.Sort(nty)
		data[i] = &ExportsPackage{
			Name: pkgName,
			Package: &exports.Package{
				Entities: []*exports.Entity{nty},
			},
		}
		i++
	}
	return data
}

func NewSingleEntityExportsOut(pkgName string, nty *exports.Entity) ExportsOut {
	if nty == nil {
		return nil
	}
	sort.Sort(nty)
	return ExportsOut{
		&ExportsPackage{
			Name: pkgName,
			Package: &exports.Package{
				Entities: []*exports.Entity{nty},
			},
		},
	}
}
