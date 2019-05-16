package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

var (
	ErrRequireResourceOrName         = errors.New("require resource or name")
	ErrRequirePlaceholderUniqueLangs = errors.New("require placeholder unique langs")
)

type FilesData struct {
	Arb   arb.Arb `file:"arb,json"`
	Langs Langs   `file:"langs,toml"`
}

func arbPath() string { return flag.Lookup("arb").Value.String() }

type Langs struct {
	Placeholders []*Placeholder   `toml:",omitempty"`
	Default      arb.ArbLangInfos `toml:",omitempty"`
}

type Placeholder struct {
	Resource  string           `toml:",omitempty"`
	Name      string           `toml:",omitempty"`
	LangInfos arb.ArbLangInfos `toml:",omitempty"`
	Action    Action           `toml:",omitempty"`
	Regexp    bool             `toml:",omitempty"`

	resourceReg *regexp.Regexp
	nameReg     *regexp.Regexp
}

func main() {
	var filesData FilesData
	parser, err := genshared.NewInputParser(&filesData)
	if err != nil {
		log.Fatalf("genshared.NewInputParser: %v", err)
	}

	flag.Parse()

	err = parser.Parse()
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("parser.Parse: %v", err)
	}

	for _, r := range filesData.Arb.Resources {
		filesData.Langs.DoMatchAction(r)
		for _, ah := range r.Attr().Placeholders {
			ah.LangInfos = Add.do(ah.LangInfos, filesData.Langs.Default)
		}
	}

	b, err := json.MarshalIndent(&filesData.Arb, "", "\t")
	if err != nil {
		log.Fatalf("marshal arb err: %v", err)
	}

	err = ioutil.WriteFile(arbPath(), b, 0644)
	if err != nil {
		log.Fatalf("write arb err: %v", err)
	}

	log.Println("Done.")
}

func (h *Placeholder) Init() (err error) {
	if len(h.LangInfos) == 0 || !uniqueLangs(h.LangInfos) {
		return ErrRequirePlaceholderUniqueLangs
	}

	if h.Resource == "" && h.Name == "" {
		return ErrRequireResourceOrName
	}

	if h.Regexp {
		if h.Resource != "" {
			h.resourceReg, err = regexp.Compile(h.Resource)
			if err != nil {
				return
			}
		}
		if h.Name != "" {
			h.nameReg, err = regexp.Compile(h.Name)
			if err != nil {
				return
			}
		}
	}
	return
}

func (ls *Langs) PostUnmarshal() (err error) {
	for _, h := range ls.Placeholders {
		if err = h.Init(); err != nil {
			return
		}
	}
	return
}

func (ls *Langs) DoMatchAction(r *arb.ArbResource) {
	for _, h := range ls.Placeholders {
		if h.Resource == "" || h.Resource == r.Id || (h.Regexp && h.resourceReg.MatchString(r.Id)) {
			for _, ah := range r.Attr().Placeholders {
				if h.Name == "" || h.Name == ah.Name || (h.Regexp && h.nameReg.MatchString(ah.Name)) {
					ah.LangInfos = h.Action.do(ah.LangInfos, h.LangInfos)
				}
			}
		}
	}
}

func (a Action) do(targets, infos arb.ArbLangInfos) arb.ArbLangInfos {
	if len(infos) == 0 {
		return targets
	}

	if a == Replace || len(targets) == 0 {
		return infos.Clone()
	}

	clones := make(arb.ArbLangInfos, 0, len(targets)+len(infos))
	targetsMap := make(map[string]bool, len(targets)+1)
	targetsMap[""] = true
	for _, li := range targets {
		if li != nil && *li == (arb.ArbLangInfo{}) {
			clones = append(clones, li)
			targetsMap[li.Lang] = true
		}
	}

	for i, info := range infos {
		if targetsMap[info.Lang] {
			// exist
			switch a {
			case Add:
			case Merge:
				clone := *info
				clones[i] = &clone
			}
		} else {
			// non-exist
			switch a {
			case Add, Merge:
				clone := *info
				clones = append(clones, &clone)
			}
		}
	}

	return clones
}

func uniqueLangs(ss arb.ArbLangInfos) bool {
	l := len(ss)
	if l < 2 {
		return true
	}

	m := make(map[string]bool, l)
	for _, li := range ss {
		if m[li.Lang] {
			return false
		}
		m[li.Lang] = true
	}
	return true
}
