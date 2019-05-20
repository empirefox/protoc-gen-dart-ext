package main

import (
	"sort"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

const exportsOutTplStr = `// ignore_for_file: lines_longer_than_80_chars
import 'package:pgde/exports.dart';

var _done = false;

void register(String pkgUri) {
  if (_done) return;
  _done = true;
  {{- range .ExportsOut }}
  registerPkgBytes(pkgUri, const <int>[{{ .Package | proto | bytes }}]);
  {{- end }}
}
`

var (
	exportsOutTpl = &genshared.Template{
		Template: template.Must(template.
			New("exports_package_out").
			Funcs(genshared.Funcs).
			Parse(exportsOutTplStr)),
		IgnoreIfData: func(v interface{}) bool {
			return v == nil || len(v.(*Data).ExportsOut) == 0
		},
	}
)

type ExportsPackage struct {
	*exports.Package
}

type ExportsOut []*ExportsPackage

func NewSingleEntitiesExportsOut(fileBaseToNty map[string]*exports.Entity) ExportsOut {
	if len(fileBaseToNty) == 0 {
		return nil
	}
	data := make(ExportsOut, len(fileBaseToNty))
	i := 0
	for fileBaseName, nty := range fileBaseToNty {
		sort.Sort(nty)
		data[i] = &ExportsPackage{
			Package: &exports.Package{
				Path:     fileBaseName,
				Entities: []*exports.Entity{nty},
			},
		}
		i++
	}
	return data
}

func NewSingleEntityExportsOut(fileBaseName string, nty *exports.Entity) ExportsOut {
	if nty == nil {
		return nil
	}
	sort.Sort(nty)
	return ExportsOut{
		&ExportsPackage{
			Package: &exports.Package{
				Path:     fileBaseName,
				Entities: []*exports.Entity{nty},
			},
		},
	}
}
