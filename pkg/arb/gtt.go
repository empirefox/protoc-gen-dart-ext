package arb

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"golang.org/x/text/language"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

type GttArbs struct {
	Info       GttArbInfo
	List       []*Arb
	ByLocale   map[language.Tag]*Arb
	Overwrites map[*Arb]*GttArbOverwrite
}

type GttArbOverwrite struct {
	Path string
	From string
}

type GttArchive struct {
	// ZipFile zip file path which expanded to env
	ZipFile string

	// ArbInfos each arb file info in zip. If using `protoc --dart-ext_out`
	// the first one is considered as the right arb for file.
	ArbInfos []*GttArbInfo

	MergeTo GttMergeTo `toml:",omitempty"`
}

type GttMergeTo struct {
	GttArbInfo
	Context string `toml:",omitempty"`
	Author  string `toml:",omitempty"`
}

type GttArbInfo struct {
	// Name name under locales in the zip file, eg app.arb
	Name string

	// Entity generated from Name if empty
	Entity string `toml:",omitempty"`
}

func (ai *GttArbInfo) GetEntity() string {
	if ai.Entity != "" {
		return ai.Entity
	}
	return util.PowerCamel(strings.TrimSuffix(ai.Name, filepath.Ext(ai.Name)))
}

func FromArchive(ga *GttArchive) ([]*GttArbs, error) {
	entitiesLen := len(ga.ArbInfos)
	if entitiesLen == 0 {
		return nil, fmt.Errorf("GttArchive.ArbInfos must be set")
	}

	r, err := zip.OpenReader(os.ExpandEnv(ga.ZipFile))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	gafs := zipfs.New(r, ga.ZipFile)
	fis, err := gafs.ReadDir("/archive")
	if err != nil {
		return nil, fmt.Errorf("read %s:/archive err: %v", ga.ZipFile, err)
	}

	if ga.MergeTo.Name != "" {
		ows := make(map[*Arb]*GttArbOverwrite, len(fis))
		locales := make([]*Arb, len(fis))
		for i, fi := range fis {
			if !fi.IsDir() {
				return nil, fmt.Errorf("%s:/archive/%s must be a dir", ga.ZipFile, fi.Name())
			}

			localeTag, err := language.Parse(fi.Name())
			if err != nil {
				return nil, err
			}

			as := make([]*Arb, len(ga.ArbInfos))
			for j, gai := range ga.ArbInfos {
				a, ow, err := decodeGttArb(gafs, gai, fi)
				if err != nil {
					return nil, err
				}

				if ow != nil {
					ows[a] = ow
				}
				as[j] = a
			}

			pkgArb := &Arb{
				Entity:  ga.MergeTo.GetEntity(),
				Context: ga.MergeTo.Context,
				Author:  ga.MergeTo.Author,
				Locale:  localeTag,
			}
			if len(as) != 0 {
				ShallowMerge(as, pkgArb)
			}

			locales[i] = pkgArb
		}

		return []*GttArbs{&GttArbs{
			Info:       ga.MergeTo.GttArbInfo,
			List:       locales,
			ByLocale:   byLocale(locales),
			Overwrites: ows,
		}}, nil
	}

	out := make([]*GttArbs, len(ga.ArbInfos))
	for i, gai := range ga.ArbInfos {
		ows := make(map[*Arb]*GttArbOverwrite, len(fis))
		as := make([]*Arb, len(fis))
		for j, fi := range fis {
			localeTag, err := language.Parse(fi.Name())
			if err != nil {
				return nil, err
			}

			a, ow, err := decodeGttArb(gafs, gai, fi)
			if err != nil {
				return nil, err
			}

			if ow != nil {
				ows[a] = ow
			}

			a.Locale = localeTag
			as[j] = a
		}

		out[i] = &GttArbs{
			Info:       *gai,
			List:       as,
			ByLocale:   byLocale(as),
			Overwrites: ows,
		}
	}
	return out, nil
}

func byLocale(list []*Arb) (m map[language.Tag]*Arb) {
	m = make(map[language.Tag]*Arb, len(list))
	for _, a := range list {
		m[a.Locale] = a
	}
	return
}

func decodeGttArb(gafs vfs.FileSystem, gai *GttArbInfo, fi os.FileInfo) (*Arb, *GttArbOverwrite, error) {
	name := "/archive/" + fi.Name() + "/" + gai.Name
	rcs, err := gafs.Open(name)
	if err != nil {
		return nil, nil, err
	}
	defer rcs.Close()

	a := new(Arb)
	err = json.NewDecoder(rcs).Decode(a)
	if err != nil {
		return nil, nil, fmt.Errorf("decode %s, err: %v", name, err)
	}

	entity := gai.GetEntity()
	if entity == a.Entity {
		return a, nil, nil
	}

	ow := &GttArbOverwrite{Path: name, From: a.Entity}
	a.Entity = entity
	return a, ow, nil
}

type GttArchiveEntity struct {
	GttArbInfo

	// ZipFile zip file path which expanded to env
	ZipFile string
}

func FromArchiveEntity(gae *GttArchiveEntity) ([]*GttArbs, error) {
	return FromArchive(&GttArchive{
		ZipFile:  gae.ZipFile,
		ArbInfos: []*GttArbInfo{&gae.GttArbInfo},
	})
}
