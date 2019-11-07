package pglt

import (
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/messageformat"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	pgs "github.com/lyft/protoc-gen-star"
)

func Register(im *dart.ImportManager, tpl *template.Template, params pgs.Parameters) {
	collectionLib := im.Import("dart:collection")
	foundationLib := im.Import("package:flutter/foundation.dart")
	materialLib := im.Import("package:flutter/material.dart")

	funcs := genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		messageformat.Funcs,
		template.FuncMap{
			"renderNode":    genshared.Render(tpl),
			"renderJoin":    genshared.RenderJoin(tpl),
			"arbsToData":    ArbsToData(im),
			"import":        im.ImportRoot,
			"collectionLib": func() *dart.ImportFile { return collectionLib },
			"foundationLib": func() *dart.ImportFile { return foundationLib },
			"materialLib":   func() *dart.ImportFile { return materialLib },
			"const": func() string {
				if len(im.InstanceClasses()) == 0 {
					return "const"
				}
				return ""
			},
			"l10nClass": func(ntyName string) string {
				return dart.Qualifier(ntyName).ToCamel().String() + "Localizations"
			},
		})

	template.Must(tpl.Funcs(funcs).Parse(fileTplStr))
	template.Must(tpl.New("fileHeader").Parse(fileHeaderTplStr))
	template.Must(tpl.New("fileBody").Parse(fileBodyTplStr))
	template.Must(tpl.New("delegate").Parse(delegateTplStr))
	template.Must(tpl.New("arbBase").Parse(arbBaseTplStr))
	template.Must(tpl.New("arb").Parse(arbTplStr))
	template.Must(tpl.New("resource").Parse(resourceTplStr))
	template.Must(tpl.New("resourceBase").Parse(resourceBaseTplStr))
	template.Must(tpl.New("root").Parse(rootTplStr))
	template.Must(tpl.New("node").Parse(nodeTplStr))
	template.Must(tpl.New("container").Parse(containerTplStr))
	template.Must(tpl.New("literal").Parse(literalTplStr))
	template.Must(tpl.New("var").Parse(varTplStr))
	template.Must(tpl.New("select").Parse(selectTplStr))
	template.Must(tpl.New("selectordinal").Parse(selectOrdinalTplStr))
	template.Must(tpl.New("plural").Parse(pluralTplStr))
	template.Must(tpl.New("equalsPlural").Parse(equalsPluralTplStr))
	template.Must(tpl.New("formedPlural").Parse(formedPluralTplStr))
}

type DartArb struct {
	*arb.Arb
	*dart.ImportManager
	Delegate []arb.SupportedLocale
}

type Data struct {
	BaseArb DartArb
	Arbs    []*arb.Arb
}

func ArbsToData(im *dart.ImportManager) func([]*arb.Arb) (*Data, error) {
	return func(as []*arb.Arb) (*Data, error) {
		for _, a := range as {
			err := a.ParseDartParams(im)
			if err != nil {
				return nil, fmt.Errorf("arb file: %v", err)
			}
		}

		return &Data{
			BaseArb: DartArb{
				Arb:           as[0],
				ImportManager: im,
				Delegate:      arb.SupportedLocales(as),
			},
			Arbs: as,
		}, nil
	}
}
