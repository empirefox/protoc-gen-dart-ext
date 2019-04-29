package arb

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/language"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

type GttArchive struct {
	// ZipFile zip file path which expanded to env
	ZipFile string

	// ArbInfos each arb file info in zip
	ArbInfos []*GttArbInfo

	// OriginLocale locale to set to Arb if Arb.Locale is empty
	OriginLocale language.Tag `toml:",omitempty"`

	PackageEntity  string `toml:",omitempty"`
	PackageContext string `toml:",omitempty"`
	PackageAuthor  string `toml:",omitempty"`
}

type GttArbInfo struct {
	// OriginFile origin arb file which expanded to env, eg $HOME/app_en.arb
	OriginFile string

	// InZipFileName name under locales in the zip file, eg app.arb
	InZipFileName string

	// Entity overwrite Arb.Entity if set
	Entity string `toml:",omitempty"`
}

func FromArchive(ga *GttArchive) ([]*Arb, error) {
	entitiesLen := len(ga.ArbInfos)
	if entitiesLen == 0 {
		return nil, fmt.Errorf("GttArchive.ArbInfos must be set")
	}

	origins := make([]*Arb, entitiesLen)
	// InZipFileName => Arb from OriginFile
	originMap := make(map[string]*Arb, entitiesLen)
	for i, gai := range ga.ArbInfos {
		b, err := ioutil.ReadFile(os.ExpandEnv(gai.OriginFile))
		if err != nil {
			return nil, err
		}

		a := new(Arb)
		err = json.Unmarshal(b, a)
		if err != nil {
			return nil, err
		}

		if !ga.OriginLocale.IsRoot() {
			if !a.Locale.IsRoot() && a.Locale != ga.OriginLocale {
				return nil, fmt.Errorf("locale not match in %s, expected %s, but got %s",
					gai.OriginFile, ga.OriginLocale, a.Locale)
			}
			a.Locale = ga.OriginLocale
		}

		if a.Locale.IsRoot() {
			return nil, fmt.Errorf("GttArchive.OriginLocale or Arb.Locale must be set: %s",
				gai.OriginFile)
		}

		if gai.Entity != "" {
			a.Entity = gai.Entity
		}

		origins[i] = a
		originMap[gai.InZipFileName] = a
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

	merge := func(from []*Arb, localeTag language.Tag) *Arb {
		var pkgArb *Arb
		if ga.PackageEntity != "" {
			pkgArb = &Arb{
				Entity:  ga.PackageEntity,
				Context: ga.PackageContext,
				Author:  ga.PackageAuthor,
				Locale:  localeTag,
			}
		} else {
			pkgArb = from[0]
			from = from[1:]
		}
		if len(from) != 0 {
			ShallowMerge(from, pkgArb)
		}
		return pkgArb
	}

	locales := make([]*Arb, len(fis)+1)
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
			name := "/archive/" + fi.Name() + "/" + gai.InZipFileName
			rcs, err := gafs.Open(name)
			if err != nil {
				return nil, err
			}
			defer rcs.Close()

			a := new(Arb)
			err = json.NewDecoder(rcs).Decode(a)
			if err != nil {
				return nil, fmt.Errorf("decode %s, err: %v", name, err)
			}

			if originMap[gai.InZipFileName].Entity != "" {
				a.Entity = originMap[gai.InZipFileName].Entity
			}
			as[j] = a
		}

		locales[i+1] = merge(as, localeTag)
	}

	locales[0] = merge(origins, ga.OriginLocale)

	for _, a := range locales {
		a.SetAttrs(locales[0])
	}
	return locales, nil
}
