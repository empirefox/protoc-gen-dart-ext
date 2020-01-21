package dartpb

import (
	"fmt"
	"log"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/l10n"
	pgs "github.com/lyft/protoc-gen-star"
)

type Translator struct {
	Dart *dart.Dart
	Arb  *arb.Arb
}

func (tr *Translator) FileClassDotResourceId(refNty pgs.Entity) dart.Qualifier {
	return tr.FileClassName(refNty).Dot(tr.ResourceId(refNty))
}

func (tr *Translator) FileClassName(refNty pgs.Entity) dart.Qualifier {
	return tr.Dart.FileNames(refNty.File()).L10nName()
}

func (tr *Translator) RefExtension(refNty pgs.Entity, isEnum bool) (ext *l10n.MsgOrEnum, err error) {
	ext = new(l10n.MsgOrEnum)
	if isEnum {
		_, err = refNty.Extension(l10n.E_Enum, ext)
	} else {
		_, err = refNty.Extension(l10n.E_Message, ext)
	}
	return
}

func (tr *Translator) ResourceId(pgsNty pgs.Entity) dart.Qualifier {
	switch nty := pgsNty.(type) {
	case pgs.Message, pgs.Enum:
		return tr.Dart.NameOf(nty)
	case pgs.OneOf:
		return tr.Dart.NameOf(nty.Message()) + tr.Dart.NameOf(nty).ToCamel()
	case pgs.Field:
		return tr.Dart.NameOf(nty.Message()) + tr.Dart.NameOf(nty).ToCamel()
	case pgs.EnumValue:
		return tr.Dart.NameOf(nty.Enum()) + tr.Dart.NameOf(nty).ToCamel()
	default:
		panic(fmt.Errorf("resource id from `%s` type error: %T", tr.Dart.NameOf(nty), pgsNty))
	}
}

type L10nMsgOrEnum struct {
	Extension l10n.MsgOrEnum
	IsError   bool
	IsEnum    bool
	Message   *Message
	Enum      *Enum
}

func (l *L10nMsgOrEnum) Children() interface{} {
	if l.IsEnum {
		return l.Enum.Values
	}
	return l.Message.Fields
}

func (l *L10nMsgOrEnum) Entity() *Entity {
	if l.IsEnum {
		return &l.Enum.Entity
	}
	return &l.Message.Entity
}

func (l *L10nMsgOrEnum) PgsEntity() pgs.Entity {
	if l.IsEnum {
		return l.Enum.Pgs
	}
	return l.Message.Pgs
}

type L10nEnumValue struct {
	Extension l10n.Value
	Entity    *EnumValue
}

type L10nOneOf struct {
	Extension l10n.Oneof
	Entity    *OneOf
}

type L10nField struct {
	Extension l10n.Field
	Entity    *Field

	IsRefEnum  bool
	RefEnum    pgs.Enum
	RefMessage pgs.Message

	// IsRefWKT if set to true, HasRef and IsRefOtherFile both return false
	IsRefWKT bool
}

func (l *L10nField) RefEntity() pgs.Entity {
	if l.IsRefEnum {
		return l.RefEnum
	}
	return l.RefMessage
}

func (l *L10nField) HasRef() bool {
	return l.RefEntity() != nil && !l.IsRefWKT
}

func (l *L10nField) IsRefOtherFile() bool {
	return l.HasRef() &&
		l.RefEnum != nil && l.RefEnum.File() != l.Entity.Pgs.File() ||
		l.RefMessage != nil && l.RefMessage.File() != l.Entity.Pgs.File()
}

func (l *L10nField) IsAssetDisabled(assetName string) bool {
	return l.DefaultAssetValue(assetName) == "-"
}

func dartNameAsValueIfEmpty(s string, dn dart.Qualifier) string {
	if s == "" {
		s = dn.ToDelimited(' ').String()
	}
	return s
}

func getDescOrLabel(desc, label string) string {
	if desc == "$" {
		return label
	}
	return desc
}

func (l *L10nMsgOrEnum) File() *File { return l.Entity().File }
func (l *L10nEnumValue) File() *File { return l.Entity.File }
func (l *L10nOneOf) File() *File     { return l.Entity.File }
func (l *L10nField) File() *File     { return l.Entity.File }

func (l *L10nMsgOrEnum) Translator() *Translator { return l.File().Translator }
func (l *L10nEnumValue) Translator() *Translator { return l.File().Translator }
func (l *L10nOneOf) Translator() *Translator     { return l.File().Translator }
func (l *L10nField) Translator() *Translator     { return l.File().Translator }

func (l *L10nMsgOrEnum) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "":
		v = l.Extension.Value
	case "Label":
		v = l.Extension.Label
	case "Helper":
		v = l.Extension.Helper
	case "Hint":
		v = l.Extension.Hint
	case "Prefix":
		v = l.Extension.Prefix
	case "Suffix":
		v = l.Extension.Suffix
	default:
		log.Fatalf("Message or Enum[%s] asset name typo: %s",
			l.Entity().DartName,
			assetName)
	}
	return v
}

func (l *L10nOneOf) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "Label":
		v = l.Extension.Label
	case "Prefix":
		v = l.Extension.Prefix
	case "Suffix":
		v = l.Extension.Suffix
	default:
		log.Fatalf("oneof[%s.%s] asset name typo: %s",
			l.Entity.Message.DartName,
			l.Entity.DartName,
			assetName)
	}
	return v
}

func (l *L10nField) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "":
		v = l.Extension.Value
	case "Label":
		v = l.Extension.Label
	case "Helper":
		v = l.Extension.Helper
	case "Hint":
		v = l.Extension.Hint
	case "Prefix":
		v = l.Extension.Prefix
	case "Suffix":
		v = l.Extension.Suffix
	case "BoolTrue":
		v = l.Extension.BoolTrue
	case "BoolFalse":
		v = l.Extension.BoolFalse
	default:
		log.Fatalf("field[%s.%s] asset name typo: %s",
			l.Entity.Message.DartName,
			l.Entity.DartName,
			assetName)
	}
	return v
}

func (l *L10nMsgOrEnum) ResourceId() dart.Qualifier {
	return l.Translator().ResourceId(l.PgsEntity())
}
func (l *L10nEnumValue) ResourceId() dart.Qualifier {
	return l.Translator().ResourceId(l.Entity.Pgs)
}
func (l *L10nOneOf) ResourceId() dart.Qualifier {
	return l.Translator().ResourceId(l.Entity.Pgs)
}
func (l *L10nField) ResourceId() dart.Qualifier {
	return l.Translator().ResourceId(l.Entity.Pgs)
}

func (l *L10nMsgOrEnum) addValueResource() {
	ta := l.Translator().Arb
	value := dartNameAsValueIfEmpty(l.Extension.Value, l.Entity().DartName)
	desc := l.Desc()
	if l.IsEnum {
		desc = "[enum] " + desc
	}
	attr := &arb.ArbAttributes{Type: "text", Description: desc}
	r := arb.NewResource(ta, l.ResourceId().String(), value, attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nEnumValue) addValueResource() {
	ta := l.Translator().Arb
	value := dartNameAsValueIfEmpty(l.Extension.Value, l.Entity.DartName)
	attr := &arb.ArbAttributes{Type: "text", Description: "[enum] " + l.Extension.Desc}
	r := arb.NewResource(ta, l.ResourceId().String(), value, attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nOneOf) addValueResource() {
	ta := l.Translator().Arb
	value := dartNameAsValueIfEmpty(l.Extension.Value, l.Entity.DartName)
	attr := &arb.ArbAttributes{Type: "text", Description: "[oneof] " + l.Desc()}
	r := arb.NewResource(ta, l.ResourceId().String(), value, attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nField) addValueResource() error {
	ta := l.Translator().Arb
	value := l.Extension.Value
	if value == "$" {
		value = dartNameAsValueIfEmpty("", l.Entity.DartName)
	}
	desc, err := l.Desc()
	if err != nil {
		return err
	}
	attr := &arb.ArbAttributes{Type: "text", Description: desc}
	r := arb.NewResource(ta, l.ResourceId().String(), value, attr)
	ta.Resources = append(ta.Resources, r)
	return nil
}

func (l *L10nMsgOrEnum) addAssetResource(assetName string) {
	ta := l.Translator().Arb
	rid := l.ResourceId().String()

	desc := assetName + " of " + rid
	if l.IsEnum {
		desc = "[enum] " + desc
	}
	attr := &arb.ArbAttributes{Type: "text", Description: desc}

	r := arb.NewResource(ta, rid+assetName, l.DefaultAssetValue(assetName), attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nOneOf) addAssetResource(assetName string) {
	ta := l.Translator().Arb
	rid := l.ResourceId().String()

	attr := &arb.ArbAttributes{
		Type:        "text",
		Description: "[oneof] " + assetName + " of " + rid,
	}

	r := arb.NewResource(ta, rid+assetName, l.DefaultAssetValue(assetName), attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nField) addAssetResource(assetName string) {
	assetValue := l.DefaultAssetValue(assetName)
	if assetValue == "-" {
		return
	}

	ta := l.Translator().Arb
	rid := l.ResourceId().String()
	desc := assetName + " of " + rid
	if l.HasRef() {
		refNty := l.RefEntity()
		desc += ". Set empty to redirect to " + assetName + " of " +
			l.Translator().FileClassDotResourceId(refNty).String()
	}

	attr := &arb.ArbAttributes{
		Type:        "text",
		Description: desc,
	}

	r := arb.NewResource(ta, rid+assetName, assetValue, attr)
	ta.Resources = append(ta.Resources, r)
}

func (l *L10nMsgOrEnum) addAssetsResources() (err error) {
	if l.Ignore() {
		if l.IsError {
			l.addValueResource()
		}
		return nil
	}

	l.addValueResource()
	l.addAssetResource("Label")
	l.addAssetResource("Helper")
	l.addAssetResource("Hint")
	l.addAssetResource("Prefix")
	l.addAssetResource("Suffix")

	if l.IsEnum {
		for _, c := range l.Enum.Values {
			if !c.L10n.Ignore() {
				c.L10n.addValueResource()
			}
		}
	} else {
		for _, c := range l.Message.Fields {
			if !c.L10n.Ignore() {
				err = c.L10n.addAssetsResources()
				if err != nil {
					return
				}
			}
		}
		for _, c := range l.Message.OneOfs {
			if !c.L10n.Ignore() {
				c.L10n.addAssetsResources()
			}
		}
	}
	return
}

func (l *L10nOneOf) addAssetsResources() {
	l.addValueResource()
	l.addAssetResource("Label")
	l.addAssetResource("Prefix")
	l.addAssetResource("Suffix")
}

func (l *L10nField) addAssetsResources() (err error) {
	if err = l.addValueResource(); err != nil {
		return
	}
	l.addAssetResource("Label")
	l.addAssetResource("Helper")
	l.addAssetResource("Hint")
	l.addAssetResource("Prefix")
	l.addAssetResource("Suffix")
	l.addAssetResource("BoolTrue")
	l.addAssetResource("BoolFalse")
	return
}

// Desc only used to generate arb description.
func (l *L10nMsgOrEnum) Desc() string {
	return getDescOrLabel(l.Extension.Desc, l.Extension.Label)
}

func (l *L10nOneOf) Desc() string {
	return getDescOrLabel(l.Extension.Desc, l.Extension.Label)
}

// Desc only used to generate arb description.
func (l *L10nField) Desc() (string, error) {
	s := l.Extension.Desc
	if s == "$" {
		s = l.Extension.Label
	}
	if s == "" {
		if l.HasRef() {
			refExt, err := l.Translator().RefExtension(l.RefEntity(), l.IsRefEnum)
			if err != nil {
				return "", err
			}
			return getDescOrLabel(refExt.Desc, refExt.Label), nil
		}
	}
	return s, nil
}

func (l *L10nMsgOrEnum) setDartTo(lang *arb.Arb) {
	if l.Ignore() || l.IsEnum {
		// enum will not redirect
		return
	}

	for _, f := range l.Message.Fields {
		if !f.L10n.Ignore() {
			f.L10n.setDartTo(lang)
		}
	}
}

var fieldAssetNamesToSetDart = []string{"", "Label", "Helper", "Hint", "Prefix", "Suffix"}

func (l *L10nField) setDartTo(lang *arb.Arb) {
	for _, assetName := range fieldAssetNamesToSetDart {
		if l.IsAssetDisabled(assetName) {
			continue
		}

		r := lang.ResourceMap()[l.ResourceId().String()+assetName]
		if r.Value != "" {
			continue
		}

		r.Value = "{redirect}"
		dartInfo := arb.ArbLangInfo{Lang: "dart"}

		if !l.HasRef() {
			dartInfo.Replace = "null"
		} else if l.IsRefOtherFile() {
			refNty := l.RefEntity()
			dartInfo.Replace = dart.Qualifier("&&&").
				Dot(l.Translator().FileClassDotResourceId(refNty)).
				DotString(assetName).String()
			dartInfo.Import = dart.L10nFile.RootFilePath(refNty)
		} else {
			dartInfo.Replace = l.Translator().ResourceId(l.RefEntity()).
				DotString(assetName).String()
		}

		r.Attr().Placeholders = arb.ArbPlaceholders{
			&arb.ArbPlaceholder{
				Name:      "redirect",
				LangInfos: arb.ArbLangInfos{&dartInfo},
			},
		}
	}
}
