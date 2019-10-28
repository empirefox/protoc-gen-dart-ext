package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

var (
	withDelegate = flag.Bool("with_delegate", true, "generate delegate")
)

type FilesData struct {
	Gtt arb.GttArchive `file:"gtt,toml"`
}

func main() {
	var filesData FilesData
	parser, err := genshared.NewInputParser(&filesData)
	if err != nil {
		log.Fatalf("genshared.NewInputParser: %v", err)
	}

	// flags:
	rs := genshared.NewTplGroupRenderer(&genshared.Template{Template: dartOutTpl})

	flag.Parse()

	err = parser.Parse()
	if err != nil {
		flag.PrintDefaults()
		log.Fatalf("parser.Parse: %v", err)
	}

	rs.Open()
	defer rs.Close()

	as, err := arb.FromArchive(&filesData.Gtt)
	if err != nil {
		log.Fatalf("arb file: %v", err)
	}

	var delegate []arb.SupportedLocale
	if *withDelegate {
		delegate = arb.SupportedLocales(as)
	}

	arbs := make([]*DartArb, len(as))
	for i, a := range as {
		err = a.ParseDartParams(im)
		if err != nil {
			log.Fatalf("arb file: %v", err)
		}
		arbs[i] = &DartArb{ImportManager: im, Arb: a, Delegate: delegate}
	}

	data := Data{
		BaseArb: arbs[0],
		Arbs:    arbs,
		Gtt:     &filesData.Gtt,
	}

	err = rs.Render(&data)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type DartArb struct {
	*dart.ImportManager
	*arb.Arb
	Delegate []arb.SupportedLocale
}

func (a *DartArb) Const() string {
	if len(a.InstanceClasses()) == 0 {
		return "const"
	}
	return ""
}

type Data struct {
	BaseArb *DartArb
	Arbs    []*DartArb
	Gtt     *arb.GttArchive
}
