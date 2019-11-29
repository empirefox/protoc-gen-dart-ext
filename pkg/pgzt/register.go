package pgzt

import (
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

func Register(tpl *template.Template) {
	funcs := genshared.JoinFuncs(
		template.FuncMap{
			"renderConstants": genshared.RenderConstants(tpl),
			"render":          genshared.RenderTemplater(tpl),
			"renderJoin":      genshared.RenderJoin(tpl),
		})

	template.Must(tpl.Funcs(funcs).Parse(fileTpl))
	template.Must(tpl.New("fileHeader").Parse(fileHeaderTpl))
	template.Must(tpl.New("fileBody").Parse(fileBodyTpl))

	template.Must(tpl.New("msg").Parse(msgTpl))
	template.Must(tpl.New("action").Parse(actionTpl))
	template.Must(tpl.New("oneOf").Parse(oneOfTpl))
	template.Must(tpl.New("oneOfConst").Parse(oneOfConstTpl))

	template.Must(tpl.New("field").Parse(fieldTpl))
	template.Must(tpl.New("fieldConst").Parse(fieldConstTpl))

	template.Must(tpl.New("repeated").Parse(repeatedTpl))
	template.Must(tpl.New("repeatedConst").Parse(repeatedConstTpl))
	template.Must(tpl.New("map").Parse(mapTpl))
	template.Must(tpl.New("mapConst").Parse(mapConstTpl))

	template.Must(tpl.New("message").Parse(messageTpl))
	template.Must(tpl.New("messageConst").Parse(messageConstTpl))

	template.Must(tpl.New("wkt").Parse(wktTpl))
	template.Must(tpl.New("wktConst").Parse(wktConstTpl))

	template.Must(tpl.New("evalTime").Parse(evalTimeTpl))
	template.Must(tpl.New("evalTimeConst").Parse(evalTimeConstTpl))

	template.Must(tpl.New("int64").Parse(int64Tpl))
	template.Must(tpl.New("int64Const").Parse(int64ConstTpl))

	template.Must(tpl.New("string").Parse(stringTpl))
	template.Must(tpl.New("stringConst").Parse(stringConstTpl))
	template.Must(tpl.New("bytes").Parse(bytesTpl))
	template.Must(tpl.New("bytesConst").Parse(bytesConstTpl))

	template.Must(tpl.New("enum").Parse(enumTpl))
	template.Must(tpl.New("bool").Parse(boolTpl))
	template.Must(tpl.New("num").Parse(numTpl))

	template.Must(tpl.New("none").Parse(noneTpl))
}
