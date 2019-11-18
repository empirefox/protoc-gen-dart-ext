package main

import (
	"flag"
	"log"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pglt"
)

type FilesData struct {
	Gtt arb.GttArchive `file:"gtt,toml"`
}

const dart_out = "dart_out"

var (
	dartOutTpl = template.New(dart_out)
)

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

	gtt, err := arb.FromArchive(&filesData.Gtt)
	if err != nil {
		log.Fatalf("arb file: %v", err)
	}

	for _, as := range gtt {
		for _, a := range as.List {
			ow, ok := as.Overwrites[a]
			if ok {
				log.Printf("gtt arb overwrite: [%s] => [%s], in zip: %s\n",
					ow.From, as.Info.Entity, ow.Path)
			}
		}
	}

	d := dart.NewDart()
	im := dart.NewImportManager(d, flag.Lookup(dart_out).Value.String(), "$", "$")

	pglt.Register(dartOutTpl)

	err = rs.Render(pglt.GttData{
		ImportManager: im,
		Gtt:           gtt,
	})
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}
