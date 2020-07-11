package pgrt

import (
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

func Register(tpl *template.Template) {
	funcs := template.FuncMap{
		"renderJoin": genshared.RenderJoin(tpl),
	}

	template.Must(tpl.Funcs(funcs).Parse(fileTpl))
	template.Must(tpl.New("fileHeader").Parse(fileHeaderTpl))
	template.Must(tpl.New("fileBody").Parse(fileBodyTpl))
	template.Must(tpl.New("groupView").Parse(groupViewTpl))
	template.Must(tpl.New("leaf").Parse(leafTpl))
	template.Must(tpl.New("entry").Parse(entryTpl))
	template.Must(tpl.New("form").Parse(formTpl))
	template.Must(tpl.New("list").Parse(listTpl))
	template.Must(tpl.New("srcRPCLine").Parse(srcRPCLineTpl))
	template.Must(tpl.New("selectOne").Parse(selectOneTpl))
	template.Must(tpl.New("selectMany").Parse(selectManyTpl))
}
