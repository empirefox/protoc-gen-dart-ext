package pglt

import (
	"fmt"
	"text/template"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/messageformat"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dartpb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"golang.org/x/text/language"
)

func Register(tpl *template.Template) {
	funcs := genshared.JoinFuncs(
		sprig.HermeticTxtFuncMap(),
		dart.Funcs,
		genshared.Funcs,
		messageformat.Funcs,
		template.FuncMap{
			"render":     genshared.RenderTemplater(tpl),
			"renderJoin": genshared.RenderJoin(tpl),
			"gttToData":  GttToData,
			"resource":   NewArbResource,
			"l10nClass": func(ntyName string) string {
				return dart.Qualifier(ntyName).ToCamel().String() + "Localizations"
			},
			"matchFn": func(pluralLib *dart.ImportFile, tag language.Tag) dart.Qualifier {
				return pluralLib.AsDot(dart.Qualifier("match" + util.PowerCamel(tag.String())))
			},
			"node": func(nd messageformat.Node) Node { return Node{nd} },
		})

	template.Must(tpl.Funcs(funcs).Parse(fileTplStr))
	template.Must(tpl.New("fileHeader").Parse(fileHeaderTplStr))
	template.Must(tpl.New("fileBody").Parse(fileBodyTplStr))
	template.Must(tpl.New("delegate").Parse(delegateTplStr))
	template.Must(tpl.New("arbBase").Parse(arbBaseTplStr))
	template.Must(tpl.New("msgOrEnum").Parse(msgOrEnumTplStr))
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

type ArbResource struct {
	*arb.ArbResource
	im *dart.ImportManager
}

func NewArbResource(im *dart.ImportManager, r *arb.ArbResource) ArbResource {
	return ArbResource{
		ArbResource: r,
		im:          im,
	}
}

func (ar ArbResource) PluralLib(varname string) (*dart.ImportFile, error) {
	imp, err := ar.Attr().DartParamImport(varname)
	if err != nil {
		return nil, err
	}
	return ar.im.ImportRoot(imp)
}

type DartArb struct {
	*arb.Arb
	*dart.ImportManager
	Delegate []arb.SupportedLocale
	File     *dartpb.File
}

type Entity struct {
	BaseArb DartArb
	Arbs    []DartArb
}

type RenderData struct {
	*dart.ImportManager
	Entities []Entity
}

type GttData struct {
	File          *dartpb.File
	ImportManager *dart.ImportManager
	Gtt           []*arb.GttArbs
}

func GttToData(data GttData) (*RenderData, error) {
	entities := make([]Entity, len(data.Gtt))
	for i, as := range data.Gtt {
		list := make([]DartArb, len(as.List))
		delegate := arb.SupportedLocales(as.List)
		for i, a := range as.List {
			err := a.ParseDartParams(data.ImportManager)
			if err != nil {
				return nil, fmt.Errorf("arb file: %v", err)
			}

			list[i] = DartArb{
				Arb:           a,
				ImportManager: data.ImportManager,
				Delegate:      delegate,
				File:          data.File,
			}
		}
		entities[i] = Entity{
			BaseArb: list[0],
			Arbs:    list,
		}
	}
	return &RenderData{
		ImportManager: data.ImportManager,
		Entities:      entities,
	}, nil
}

type Node struct {
	messageformat.Node
}

func (nd Node) TemplateName() string {
	return nd.Node.Type()
}

func (nd Node) TemplateData() interface{} {
	return nd.Node
}
