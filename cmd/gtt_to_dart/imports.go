package main

import (
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/imports"
)

const importsOutTplStr = `{{ .ImportsOut | proto | printf "%s" }}`

var (
	importsOutTpl = &genshared.Template{
		Template: template.Must(template.
			New("imports_package_out").
			Funcs(genshared.Funcs).
			Parse(importsOutTplStr)),
		IgnoreIfData: func(v interface{}) bool {
			return v == nil || v.(*Data).ImportsOut.Package == nil
		},
	}
)

type ImportsOut struct {
	*imports.Package
}

func NewSingleEntityImportsOut(nty *imports.Entity) ImportsOut {
	if nty == nil {
		return ImportsOut{}
	}
	return ImportsOut{
		&imports.Package{
			Entities: []*imports.Entity{nty},
		},
	}
}
