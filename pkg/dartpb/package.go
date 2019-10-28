package dartpb

import (
	"time"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

type PackageGroup struct {
	*dart.Dart
	List       []*Package
	PgsToPkg   map[pgs.Package]*Package
	PgsToMsg   map[pgs.Message]*Message
	PgsToOneOf map[pgs.OneOf]*OneOf
	PgsToField map[pgs.Field]*Field
	PgsToEnum  map[pgs.Enum]*Enum
	PgsToValue map[pgs.EnumValue]*EnumValue

	// SplitHeader split the following header
	SplitHeader string
}

func NewPackageGroup(d *dart.Dart, locale language.Tag, pkgs []pgs.Package, splitHeader string) (g *PackageGroup, err error) {
	g = &PackageGroup{
		Dart:       d,
		List:       make([]*Package, len(pkgs)),
		PgsToPkg:   make(map[pgs.Package]*Package, 32),
		PgsToMsg:   make(map[pgs.Message]*Message, 128),
		PgsToOneOf: make(map[pgs.OneOf]*OneOf, 128),
		PgsToField: make(map[pgs.Field]*Field, 1024),
		PgsToEnum:  make(map[pgs.Enum]*Enum, 128),
		PgsToValue: make(map[pgs.EnumValue]*EnumValue, 1024),

		SplitHeader: splitHeader,
	}

	// build tree
	for i, pkg := range pkgs {
		g.List[i], err = newPackage(g, locale, pkg)
		if err != nil {
			return
		}
	}

	// set ref to tree
	for _, p := range g.List {
		for _, msg := range p.Messages {
			for _, f := range msg.Fields {
				if f.Pgs.Type().IsRepeated() || f.Pgs.Type().IsMap() {
					elem := f.Pgs.Type().Element()
					if elem.IsEnum() {
						f.Arb.setRef(g.PgsToEnum[elem.Enum()].Arb)
					} else if elem.IsEmbed() {
						f.Arb.setRef(g.PgsToMsg[elem.Embed()].Arb)
					}
				} else if f.Pgs.Type().IsEnum() {
					f.Arb.setRef(g.PgsToEnum[f.Pgs.Type().Enum()].Arb)
				} else if f.Pgs.Type().IsEmbed() {
					f.Arb.setRef(g.PgsToMsg[f.Pgs.Type().Embed()].Arb)
				}
			}
		}
	}

	// add resources to arb using ref
	for _, p := range g.List {
		for _, nty := range p.Messages {
			nty.Arb.addAssetsResources()
		}
		for _, nty := range p.Enums {
			nty.Arb.addAssetsResources()
		}
	}

	// add fields to NonOneOfFields or OneOf.Fields
	for _, p := range g.List {
		for _, msg := range p.Messages {
			for _, nty := range msg.Fields {
				if nty.Pgs.InOneOf() {
					oneof := g.PgsToOneOf[nty.Pgs.OneOf()]
					oneof.Fields = append(oneof.Fields, nty)
				} else {
					msg.NonOneOfFields = append(msg.NonOneOfFields, nty)
				}
			}
		}
	}

	return
}

type Package struct {
	Group    *PackageGroup
	DartName pgs.Name
	Pgs      pgs.Package

	L10n *L10n

	Validator *Validator

	Messages []*Message
	Enums    []*Enum
}

func newPackage(g *PackageGroup, locale language.Tag, pgsPkg pgs.Package) (*Package, error) {
	pkgNtyDartName := g.EntityNameOf(pgsPkg)
	p := &Package{
		Group:    g,
		DartName: pkgNtyDartName,
		L10n: &L10n{
			RootFilePath: "pgde/" + pgsPkg.ProtoName().String() + ".l10n.dart",
			ClassName:    pkgNtyDartName.UpperCamelCase().String() + "Localization",
			Arb: &arb.Arb{
				LastModified: iso8601.Time{time.Now()},
				Locale:       locale,
				Entity:       pkgNtyDartName.String(),
			},
			RefManager:    new(PackageArbRefManager),
			ImportManager: dart.NewDefaultImportManager(),
		},
		Validator: &Validator{
			RootFilePath:         "pgde/" + pgsPkg.ProtoName().String() + ".validate.dart",
			ClassName:            pkgNtyDartName.UpperCamelCase().String() + "Validator",
			BuildContextAccessor: "_context",
			InfoAccessor:         "_info",
			L10nAccessor:         "_l10n",
			FieldAccessor:        "_v",
			ImportManager:        dart.NewDefaultImportManager(),
		},
		Pgs: pgsPkg,
	}
	p.Validator.Package = p

	p.Validator.MaterialFile = p.Validator.ImportManager.
		GetFile("package:flutter/material.dart").
		AddShow("BuildContext").
		AddShow("Localizations").
		EnableShow()
	p.Validator.FixnumFile = p.Validator.ImportManager.
		GetFile("package:fixnum/fixnum.dart").
		AddShow("Int64").
		EnableShow()
	p.Validator.EmailValidatorFile = p.Validator.ImportManager.GetFile("package:email_validator/email_validator.dart")
	p.Validator.PgdeFile = p.Validator.ImportManager.GetFile("package:pgde/pgde.dart")
	p.Validator.L10nFile = p.Validator.ImportManager.GetFile(ResolveImport(p.Validator.RootFilePath, p.L10n.RootFilePath))

	p.Validator.FullL10nClass = p.Validator.L10nFile.As + "." + p.L10n.ClassName

	g.PgsToPkg[pgsPkg] = p

	for _, f := range pgsPkg.Files() {
		for _, pgsNty := range f.AllEnums() {
			err := p.addEnum(g, pgsNty)
			if err != nil {
				return nil, err
			}
		}
		for _, pgsNty := range f.AllMessages() {
			err := p.addMessage(g, pgsNty)
			if err != nil {
				return nil, err
			}
		}
	}

	return p, nil
}

func (p *Package) SetL10nDartTo(lang *arb.Arb) {
	for _, nty := range p.Messages {
		nty.Arb.setDartTo(lang)
	}
}
