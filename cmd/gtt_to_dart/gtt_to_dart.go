package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

var (
	withDelegate   = flag.Bool("with_delegate", true, "generate delegate")
	dartImportFile = flag.String("dart_import_file", "",
		"the dart import full path or file base name to overwrite `dart_out`,\n"+
			" the `package:pkga` can be auto computed")
)

type FilesData struct {
	Gtt     arb.GttArchive   `file:"gtt,toml"`
	Exports *exports.Exports `file:"resolve,proto"`
}

func main() {
	var filesData FilesData
	parser, err := genshared.NewInputParser(&filesData)
	if err != nil {
		log.Fatalf("genshared.NewInputParser: %v", err)
	}

	// flags:
	rs := genshared.NewTplGroupRenderer(
		&genshared.Template{Template: dartOutTpl},
		exportsOutTpl,
		importsOutTpl)

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

	if *dartImportFile == "" {
		*dartImportFile = filepath.Base(flag.Lookup(dart_out).Value.String())
	}

	data := Data{
		BaseArb:    BaseArb{Arb: as[0], Delegate: delegate},
		Arbs:       make([]*dart.Resolved, len(as)),
		Gtt:        &filesData.Gtt,
		ExportsOut: NewSingleEntityExportsOut(*dartImportFile, as[0].ExportProto()),
	}

	resolved, err := dart.ResolveWithExports(filesData.Exports, as[0])
	if err != nil {
		log.Fatalf("resolve arb: %v", err)
	}
	resolved.Imports.AddNoAs("package:pgde/plural.dart")
	data.Imports = resolved.Imports.ToDartSource()
	data.BaseArb.Imports = resolved.Entries.ToDartSource()
	data.ImportsOut = NewSingleEntityImportsOut(resolved.ImportProto())

	for i, a := range as {
		data.Arbs[i] = resolved.WithArb(a)
	}

	err = rs.Render(&data)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	log.Println("Done.")
}

type BaseArb struct {
	*arb.Arb
	Delegate []arb.SupportedLocale
	Imports  string
}

type Data struct {
	BaseArb    BaseArb
	Arbs       []*dart.Resolved
	Gtt        *arb.GttArchive
	Imports    string
	ExportsOut ExportsOut
	ImportsOut ImportsOut
}
