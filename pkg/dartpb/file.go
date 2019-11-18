package dartpb

import (
	"time"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

type File struct {
	Dart *dart.Dart
	Pgs  pgs.File

	Names dart.FileNames

	Translator *Translator

	Validators *Validators

	Messages []*Message
	Enums    []*Enum
}

func NewFile(d *dart.Dart, locale language.Tag, pgsFile pgs.File) (*File, error) {
	names := d.FileNames(pgsFile)

	f := &File{
		Dart:  d,
		Pgs:   pgsFile,
		Names: names,

		Translator: &Translator{
			Dart: d,
			Arb: &arb.Arb{
				LastModified: iso8601.Time{time.Now()},
				Locale:       locale,
				Entity:       names.EntityName().String(),
			},
		},
		Validators: &Validators{
			BuildContextAccessor:  "_ctx",
			InfoAccessor:          "_info",
			L10nAccessor:          "_l10n",
			FieldAccessor:         "_v",
			EnumFieldL10nAccessor: "_enumL10n",
			ImportManager: dart.NewDefaultImportManager(d,
				dart.ValidateFile.RootFilePath(pgsFile)),
		},
	}
	f.Validators.File = f

	validateImportPb, err := f.Validators.ImportManager.
		Resolve(dart.PbFile.RootFilePath(pgsFile))
	if err != nil {
		return nil, err
	}

	validateImportL10n, err := f.Validators.ImportManager.
		Resolve(dart.L10nFile.RootFilePath(pgsFile))
	if err != nil {
		return nil, err
	}

	f.Validators.CollectionLib = f.Validators.ImportManager.
		Import("dart:collection")
	f.Validators.ConvertLib = f.Validators.ImportManager.
		Import("dart:convert")
	f.Validators.FoundationFile = f.Validators.ImportManager.
		Import("package:flutter/foundation.dart")
	f.Validators.MaterialFile = f.Validators.ImportManager.
		Import("package:flutter/material.dart")
	f.Validators.FixnumFile = f.Validators.ImportManager.
		Import("package:fixnum/fixnum.dart")
	f.Validators.EmailValidatorFile = f.Validators.ImportManager.
		Import("package:email_validator/email_validator.dart")
	f.Validators.PgdeFile = f.Validators.ImportManager.
		Import("package:pgde/pgde.dart")
	f.Validators.PbFile = f.Validators.ImportManager.Import(validateImportPb)
	f.Validators.L10nFile = f.Validators.ImportManager.Import(validateImportL10n)

	for _, pgsNty := range pgsFile.AllEnums() {
		err := f.addEnum(pgsNty)
		if err != nil {
			return nil, err
		}
	}
	for _, pgsNty := range pgsFile.AllMessages() {
		err := f.addMessage(pgsNty)
		if err != nil {
			return nil, err
		}
	}

	// add resources to arb using ref
	for _, nty := range f.Messages {
		nty.L10n.addAssetsResources()
	}
	for _, nty := range f.Enums {
		nty.L10n.addAssetsResources()
	}

	return f, nil
}

func (f *File) SetL10nTo(lang *arb.Arb) {
	for _, nty := range f.Messages {
		nty.L10n.setDartTo(lang)
	}
}
