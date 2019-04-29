package genshared

import (
	"bytes"
	"fmt"
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
	"bytes":           ToBytes,
	"superscript":     supsub.ToSup,
	"subscript":       supsub.ToSub,
}

func FuncsWithRender(pipe string, tpl *template.Template) template.FuncMap {
	return JoinFuncs(RenderFuncs(pipe, tpl), Funcs)
}

func RenderFuncs(pipe string, tpl *template.Template) template.FuncMap {
	return template.FuncMap{pipe: func(tplName string, data interface{}) (string, error) {
		var b bytes.Buffer
		err := tpl.ExecuteTemplate(&b, tplName, data)
		return b.String(), err
	}}
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

func ToBytes(b []byte) string {
	s := fmt.Sprintf("%#v", b)[7:]
	return s[:len(s)-1]
}
