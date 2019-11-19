package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"

	"github.com/Masterminds/sprig"
	"github.com/empirefox/makeplural/plural"
	"github.com/empirefox/messageformat"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"golang.org/x/text/language"
)

const (
	outputTplStr = `{{ . | toPrettyJson }}`
)

var (
	langs = flag.String("lang", "", "language subset")
	force = flag.Bool("force", false, "force copy even when no plural found")

	outputTpl = template.Must(template.New("output").
			Funcs(sprig.HermeticTxtFuncMap()).
			Parse(outputTplStr))

	filesData FilesData
)

type FilesData struct {
	Arb arb.Arb `file:"input,json"`
}

func main() {
	parser, err := genshared.NewInputParser(&filesData)
	if err != nil {
		log.Fatalf("genshared.NewInputParser: %v", err)
	}

	vr := genshared.NewGoTplVariantRenderer(outputTpl, getLangs, makeArbData, makeArbPath)

	flag.Parse()

	err = parser.Parse()
	if err != nil {
		flag.PrintDefaults()
		fmt.Println("\nThe output path format: %N=input_base_name %E=ext, %D=dir, %P=input_path_no_ext, %V=variant")
		log.Fatalf("parser.Parse: %v", err)
	}

	icuparser, err := messageformat.New()
	if err != nil {
		log.Fatalf("create ICU parser: %v", err)
	}
	icuparser.SkipForPlural(true)

	hasSelect := false
	for _, r := range filesData.Arb.Resources {
		mf, err := icuparser.Parse(r.Value, nil)
		if err != nil {
			log.Fatalf("ICU parse: %s, err: %v", r.Value, err)
		}
		if mf.HasSelect() {
			hasSelect = true
			break
		}
	}

	if !hasSelect {
		if *force {
			outputPath := flag.Lookup("output").Value.String()
			varaints := getLangs()
			dsts := make([]string, len(varaints))
			for i, lang := range varaints {
				dsts[i] = os.ExpandEnv(makeArbPath(outputPath, lang))
			}
			err = util.CopyFile(flag.Lookup("input").Value.String(), dsts...)
			if err != nil {
				log.Fatalf("copy file: %v", err)
			}
			return
		}

		log.Println("No plural found, so skip generate variants!")
		return
	}

	vr.Open()
	defer vr.Close()

	err = vr.Render(&filesData.Arb)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

func makeArbData(data interface{}, variant string) interface{} {
	tag, err := language.Parse(variant)
	if err != nil {
		log.Fatalf("parse language tag for %s, err: %v", variant, err)
	}

	_, lang, ok := plural.Info.Find(tag)
	if !ok {
		log.Fatalf("plural culture not found: %s", variant)
	}

	p, err := messageformat.NewWithCulture(lang)
	if err != nil {
		log.Fatalf("create icu parser(%s), err: %v", variant, err)
	}

	a := data.(*arb.Arb).Clone()
	a.Locale = lang
	for _, r := range a.Resources {
		mf, err := p.Parse(r.Value, nil)
		if err != nil {
			log.Fatalf("parse icu for `%s`, err: %v", r.Value, err)
		}

		newValue, err := mf.ToString()
		if err != nil {
			log.Fatalf("print parsed icu `%s`, err: %v", r.Value, err)
		}

		r.Value = newValue
	}
	return a
}

var outputPathReg = regexp.MustCompile(`%[NEDPV]`)

func makeArbPath(outputPath, variant string) string {
	inputPath := flag.Lookup("input").Value.String()
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	mapping := map[string]string{
		"%N": strings.TrimSuffix(base, ext),
		"%E": ext,
		"%D": filepath.Dir(inputPath),
		"%P": strings.TrimSuffix(inputPath, ext),
		"%V": variant,
	}
	out := outputPathReg.ReplaceAllStringFunc(outputPath,
		func(in string) string { return mapping[in] })
	return out
}

func getLangs() []string {
	if *langs == "" {
		plural.Info.Langs()
	}

	raw := strings.Split(*langs, ",")
	sort.Strings(raw)
	parseFailed, findFailed, ok := plural.Info.Validate(raw)
	if !ok {
		log.Fatalf("language parse failed: %v\n\tfind failed: %v",
			parseFailed, findFailed)
	}
	return raw
}
