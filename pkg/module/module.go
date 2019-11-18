package module

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/BurntSushi/toml"
	"github.com/Masterminds/sprig"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dartpb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pglt"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgvt"
	pgs "github.com/lyft/protoc-gen-star"
	"golang.org/x/text/language"
)

var (
	tomlHeader = []byte(`# Path placeholder usage:
#  %N: base name
#  %O: output of protoc
#  %P: output relative path
#  %E: extension of input file

`)
)

type DartExtModule struct {
	*pgs.ModuleBase
	dart   *dart.Dart
	params pgs.Parameters

	arbTpl      *template.Template
	l10nTpl     *template.Template
	validateTpl *template.Template
}

func New() *DartExtModule {
	return &DartExtModule{
		ModuleBase: &pgs.ModuleBase{},
		dart:       dart.NewDart(),
		arbTpl: template.Must(template.New("arb").
			Funcs(sprig.HermeticTxtFuncMap()).
			Parse(`{{ . | toPrettyJson }}`)),
		l10nTpl:     template.New("l10n"),
		validateTpl: template.New("validate"),
	}
}

func (m *DartExtModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.params = c.Parameters()
	pglt.Register(m.l10nTpl)
	pgvt.Register(m.validateTpl, m.params)
}

// Name satisfies the generator.Plugin interface.
func (m *DartExtModule) Name() string { return "dart-ext" }

var pathReg = regexp.MustCompile(`%[NOPE]`)

func expandPath(input string, mapping map[string]string) string {
	return pathReg.ReplaceAllStringFunc(input, func(in string) string { return mapping[in] })
}

func (m *DartExtModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for _, t := range targets {
		f, err := dartpb.NewFile(m.dart, language.Und, t)
		if err != nil {
			m.Failf("dartpb.NewFile err: ", err)
		}

		inputPath := t.InputPath().String()
		base := filepath.Base(inputPath)
		ext := filepath.Ext(base)
		mapping := map[string]string{
			"%N": strings.TrimSuffix(base, ext),
			"%O": m.BuildContext.OutputPath(),
			"%P": strings.TrimSuffix(inputPath, ext),
			"%E": ext,
		}

		if m.HasParam("arb") {
			gae := arb.GttArchiveEntity{
				ZipFile: "$GTT_DIR/archive.zip",
				GttArbInfo: arb.GttArbInfo{
					Name:   t.InputPath().BaseName() + ".arb",
					Entity: f.Translator.Arb.Entity,
				},
			}
			var buf bytes.Buffer
			buf.Write(tomlHeader)
			err = toml.NewEncoder(&buf).Encode(&gae)
			if err != nil {
				m.Failf("encode GttArchiveEntity err: %v", err)
			}

			name := t.InputPath().SetExt(".gtt.toml").String()
			m.AddGeneratorFile(name, buf.String())

			name = t.InputPath().SetExt(".arb").String()
			m.AddGeneratorTemplateFile(name, m.arbTpl, f.Translator.Arb)
		}

		if m.HasParam("l10n") {
			gttTomlPath := m.params.Str("l10n")
			if gttTomlPath != "" {
				gttTomlPath = expandPath(gttTomlPath, mapping)
				gttTomlPath = os.ExpandEnv(gttTomlPath)
			} else {
				gttTomlPath = t.InputPath().SetExt("gtt.toml").String()
				gttTomlPath = filepath.Join(m.BuildContext.OutputPath(), gttTomlPath)
			}

			gabuf, err := ioutil.ReadFile(gttTomlPath)
			if err != nil {
				m.Failf("read file err: ", err)
			}

			var gae arb.GttArchiveEntity
			err = json.Unmarshal(gabuf, &gae)
			if err != nil {
				m.Failf("unmarshal GttArchiveEntity err: ", err)
			}

			gae.ZipFile = expandPath(gae.ZipFile, mapping)
			gtt, err := arb.FromArchiveEntity(&gae)
			if err != nil {
				m.Failf("FromArchiveEntity err: ", err)
			}

			// FromArchiveEntity only has one entity
			as := gtt[0]
			for _, a := range as.List {
				f.SetL10nTo(a)
				ow, ok := as.Overwrites[a]
				if ok {
					m.Logf("gtt arb overwrite: [%s] => [%s], in zip: %s\n",
						ow.From, as.Info.Entity, ow.Path)
				}
			}

			name := t.InputPath().SetExt(".l10n.dart").String()
			im := dart.NewImportManager(m.dart, name, "$", "$")
			m.AddGeneratorTemplateFile(name, m.l10nTpl, pglt.GttData{
				ImportManager: im,
				Gtt:           gtt,
			})
		}

		if m.HasParam("validate") {
			name := f.Validators.ImportManager.RootFilePath()
			m.AddGeneratorTemplateFile(name, m.validateTpl, f)
		}
	}

	return m.Artifacts()
}

func (m *DartExtModule) HasParam(p string) (ok bool) {
	_, ok = m.params[p]
	return
}
