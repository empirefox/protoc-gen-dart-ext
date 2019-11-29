package dartpb

import (
	"fmt"
	"time"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

type FileConfig struct {
	Dart      *dart.Dart
	Reflected map[string]*desc.FileDescriptor
	Locale    language.Tag
	Pgs       pgs.File
	ZeroPlan  string
}

type File struct {
	FileConfig

	Names dart.FileNames

	Zeros *Zeros

	Translator *Translator

	Validators *Validators

	Messages []*Message
	Enums    []*Enum
}

func NewFile(c FileConfig) (*File, error) {
	names := c.Dart.FileNames(c.Pgs)

	f := &File{
		FileConfig: c,
		Names:      names,

		Translator: &Translator{
			Dart: c.Dart,
			Arb: &arb.Arb{
				LastModified: iso8601.Time{time.Now()},
				Locale:       c.Locale,
				Entity:       names.EntityName().String(),
			},
		},
	}

	imz, err := dart.NewDefaultImportManager(c.Dart, dart.ZeroFile.RootFilePath(c.Pgs))
	if err != nil {
		return nil, err
	}
	z := Zeros{
		File:             f,
		Plan:             c.ZeroPlan,
		ImportManager:    imz,
		MemberPrefix:     "_$",
		MemberElemSuffix: "_e",
	}
	f.Zeros = &z

	imv, err := dart.NewDefaultImportManager(c.Dart, dart.ValidateFile.RootFilePath(c.Pgs))
	if err != nil {
		return nil, err
	}
	v := Validators{
		File:                  f,
		BuildContextAccessor:  "_ctx",
		InfoAccessor:          "_info",
		L10nAccessor:          "_l10n",
		FieldAccessor:         "_v",
		EnumFieldL10nAccessor: "_enumL10n",
		ImportManager:         imv,
	}
	f.Validators = &v

	for _, pgsNty := range c.Pgs.AllEnums() {
		err := f.addEnum(pgsNty)
		if err != nil {
			return nil, err
		}
	}
	for _, pgsNty := range c.Pgs.AllMessages() {
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

func (f *File) ReflectMessage(pgsNty pgs.Message) (*dynamic.Message, error) {
	if pgsNty == nil {
		return nil, fmt.Errorf("reflect nil pgs message")
	}
	fd, ok := f.Reflected[pgsNty.File().Name().String()]
	if !ok {
		return nil, fmt.Errorf("reflect proto file not found: %s", pgsNty.File().Name())
	}
	md := fd.FindMessage(pgsNty.FullyQualifiedName())
	if md == nil {
		return nil, fmt.Errorf("reflect proto message not found: %s", pgsNty.FullyQualifiedName())
	}
	return dynamic.NewMessage(md), nil
}

func (f *File) SetL10nTo(lang *arb.Arb) {
	for _, nty := range f.Messages {
		nty.L10n.setDartTo(lang)
	}
}
