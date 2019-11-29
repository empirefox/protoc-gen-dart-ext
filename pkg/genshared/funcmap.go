package genshared

import (
	"bytes"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/lynn9388/supsub"
	strcase "github.com/stoewer/go-strcase"
)

var Funcs = template.FuncMap{
	"fileBaseName":    FileBaseName,
	"powerKebab":      strcase.KebabCase,
	"powerCamel":      util.PowerCamel,
	"powerLowerCamel": util.PowerLowerCamel,
	"powerSnake":      strcase.SnakeCase,
	"proto":           proto.Marshal,
	"bytes":           util.BytesLiteral,
	"superscript":     supsub.ToSup,
	"subscript":       supsub.ToSub,
}

type ConstantsTemplater interface {
	ConstTemplateName() string
}

func RenderConstants(tpl *template.Template) func(ConstantsTemplater) (string, error) {
	return func(t ConstantsTemplater) (string, error) {
		var b bytes.Buffer
		var err error

		name := t.ConstTemplateName()
		for _, item := range tpl.Templates() {
			if item.Name() == name {
				err = tpl.ExecuteTemplate(&b, name, t)
				break
			}
		}

		return b.String(), err
	}
}

type Templater interface {
	TemplateName() string
}

func RenderTemplater(tpl *template.Template) func(Templater) (string, error) {
	return func(t Templater) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, t.TemplateName(), t)
		return b.String(), err
	}
}

func VersionEntityFuncs(ve *VersionEntity) template.FuncMap {
	return template.FuncMap{
		"entity":  ve.LowerCamel,
		"Entity":  ve.UpperCamel,
		"entityV": ve.LowerCamelV,
		"EntityV": ve.UpperCamelV,
	}
}

func JoinFuncs(ms ...template.FuncMap) template.FuncMap {
	total := 0
	for _, m := range ms {
		total += len(m)
	}
	result := make(template.FuncMap, total)
	for _, m := range ms {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

func FileBaseName(name string) string {
	return pgs.FilePath(name).BaseName()
}
