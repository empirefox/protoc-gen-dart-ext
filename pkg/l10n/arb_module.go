package l10n

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"unsafe"

	"github.com/BurntSushi/toml"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	pgs "github.com/lyft/protoc-gen-star"
	"golang.org/x/text/language"
)

type ArbModule struct {
	pgs.ModuleBase
	Dart   *dart.Dart
	Locale language.Tag
	Outdir string
}

func (m *ArbModule) Name() string { return "l10n" }

func (m *ArbModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	sortedPkgs := util.FilesPackages(targets)

	gtt := arb.GttArchive{
		ZipFile:      "$GTT_DIR/archive.zip",
		OriginLocale: m.Locale,
		ArbInfos:     make([]*arb.GttArbInfo, len(sortedPkgs)),
	}

	for i, pkg := range sortedPkgs {
		a, err := PackageToArb(m.Dart, m.Locale, pkg)
		if err != nil {
			m.Failf("pb to arb err: %v", err)
		}

		b, err := json.MarshalIndent(a, "", "\t")
		if err != nil {
			m.Failf("marshal arb err: %v", err)
		}

		filebase := pgs.Name(a.Entity).SnakeCase().String() + ".arb"
		m.OverwriteCustomFile(
			filepath.Join(m.Outdir, filebase),
			*(*string)(unsafe.Pointer(&b)),
			0644,
		)

		gtt.ArbInfos[i] = &arb.GttArbInfo{
			OriginFile:    filepath.Join("$GTT_DIR", filebase),
			InZipFileName: filebase,
			Entity:        a.Entity,
		}
	}

	var buf bytes.Buffer
	buf.Grow(1 << 10)
	err := toml.NewEncoder(&buf).Encode(&gtt)
	if err != nil {
		m.Failf("encode gtt err: %v", err)
	}

	m.OverwriteCustomFile(
		filepath.Join(m.Outdir, "gtt.toml"),
		buf.String(),
		0644,
	)

	return m.Artifacts()
}

// TODO
func PackageToArb(d *dart.Dart, locale language.Tag, pkg pgs.Package) (*arb.Arb, error) {
	return nil, nil
}
