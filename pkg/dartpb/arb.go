package dartpb

import (
	"log"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/l10n"
)

const (
	MsgAssetAsFieldPrefix       = "mAsset"
	EnumAssetAsFieldPrefix      = "eAsset"
	EnumValueAssetAsFieldPrefix = "vAsset"
	FieldAssetAsFieldPrefix     = "fAsset"
	OneOfAssetAsFieldPrefix     = "oAsset"
)

func (r *ArbRef) IsSamePkg() bool { return IsArbSamePkg(r.From, r.To) }
func IsArbSamePkg(from *FieldArb, to *MsgOrEnumArb) bool {
	return to == nil || from.Package() == to.Package()
}

type MsgOrEnumArb struct {
	Extension l10n.Arb
	Message   *Message
	Enum      *Enum
	IsEnum    bool
	RefBy     []*FieldArb
}

func (a *MsgOrEnumArb) Entity() *Entity {
	if a.IsEnum {
		return &a.Enum.Entity
	}
	return &a.Message.Entity
}

func (a *MsgOrEnumArb) PgsEntity() pgs.Entity {
	if a.IsEnum {
		return a.Enum.Pgs
	}
	return a.Message.Pgs
}

func (a *MsgOrEnumArb) LabelAsDesc() bool { return a.Extension.Desc == "$" }
func (a *MsgOrEnumArb) IsDescNull() bool {
	return a.Extension.Desc == "" || a.LabelAsDesc() && a.Extension.Label == ""
}

func (a *MsgOrEnumArb) Value(lang *arb.Arb) string  { return a.getRawArbString(lang, "") }
func (a *MsgOrEnumArb) Label(lang *arb.Arb) string  { return a.getRawArbString(lang, "Label") }
func (a *MsgOrEnumArb) Helper(lang *arb.Arb) string { return a.getRawArbString(lang, "Helper") }
func (a *MsgOrEnumArb) Hint(lang *arb.Arb) string   { return a.getRawArbString(lang, "Hint") }
func (a *MsgOrEnumArb) Prefix(lang *arb.Arb) string { return a.getRawArbString(lang, "Prefix") }
func (a *MsgOrEnumArb) Suffix(lang *arb.Arb) string { return a.getRawArbString(lang, "Suffix") }

type EnumValueArb struct {
	Extension l10n.ValueArb
	Entity    *EnumValue
}

func (a *EnumValueArb) Value(lang *arb.Arb) string { return a.getRawArbString(lang, "") }

type OneOfArb struct {
	Extension l10n.OneofArb
	Entity    *OneOf
}

func (a *OneOfArb) LabelAsDesc() bool { return a.Extension.Desc == "$" }
func (a *OneOfArb) IsDescNull() bool {
	return a.Extension.Desc == "" || a.LabelAsDesc() && a.Extension.Label == ""
}

func (a *OneOfArb) Value(lang *arb.Arb) string  { return a.getRawArbString(lang, "") }
func (a *OneOfArb) Label(lang *arb.Arb) string  { return a.getRawArbString(lang, "Label") }
func (a *OneOfArb) Prefix(lang *arb.Arb) string { return a.getRawArbString(lang, "Prefix") }
func (a *OneOfArb) Suffix(lang *arb.Arb) string { return a.getRawArbString(lang, "Suffix") }

type FieldArb struct {
	Extension l10n.FieldArb
	Entity    *Field
	Ref       *MsgOrEnumArb
	UseRef    bool
}

func (a *FieldArb) RefSamePkg() bool { return IsArbSamePkg(a, a.Ref) }

func (a *FieldArb) refResourceId() string {
	if a.RefSamePkg() {
		return a.Ref.ResourceId()
	}
	return a.L10n().ImportAsInstanceName(a.Ref.L10n()) + "." + a.Ref.ResourceId()
}

func (a *FieldArb) Value(lang *arb.Arb) string {
	r := lang.ResourceMap()[a.ResourceId()]
	if r.Value == "" {
		if a.Ref != nil {
			a.UseRef = true
			return a.refResourceId()
		}
	}

	return dart.RawString(r.Value)
}

func (a *FieldArb) Label(lang *arb.Arb) string     { return a.getRawArbString(lang, "Label") }
func (a *FieldArb) Helper(lang *arb.Arb) string    { return a.getRawArbString(lang, "Helper") }
func (a *FieldArb) Hint(lang *arb.Arb) string      { return a.getRawArbString(lang, "Hint") }
func (a *FieldArb) Prefix(lang *arb.Arb) string    { return a.getRawArbString(lang, "Prefix") }
func (a *FieldArb) Suffix(lang *arb.Arb) string    { return a.getRawArbString(lang, "Suffix") }
func (a *FieldArb) BoolTrue(lang *arb.Arb) string  { return a.getRawArbString(lang, "BoolTrue") }
func (a *FieldArb) BoolFalse(lang *arb.Arb) string { return a.getRawArbString(lang, "BoolFalse") }

// SetRef called after all packages parsed
func (from *FieldArb) setRef(to *MsgOrEnumArb) {
	ref := &ArbRef{
		From: from,
		To:   to,
	}
	if ref.IsSamePkg() {
		from.L10n().RefManager.Direct = append(from.L10n().RefManager.Direct, ref)
	} else {
		from.L10n().RefManager.Imports = append(from.L10n().RefManager.Imports, ref)
		to.L10n().RefManager.Exports = append(to.L10n().RefManager.Exports, ref)
	}

	from.Ref = to
	to.RefBy = append(to.RefBy, from)
}

func (a *FieldArb) IsAssetDisabled(assetName string) bool {
	return a.DefaultAssetValue(assetName) == "-"
}

func dartNameAsValueIfEmpty(s string, dn pgs.Name) string {
	if s == "" {
		s = strings.Join(dn.UpperCamelCase().Split(), " ")
	}
	return s
}

func getDescOrLabel(desc, label string) string {
	if desc == "$" {
		return label
	}
	return desc
}

func getAssetOrRef(s, ref string) string {
	switch s {
	case "-":
		return ""
	case "":
		return ref
	default:
		return s
	}
}

func (a *MsgOrEnumArb) Package() *Package { return a.Package() }
func (a *EnumValueArb) Package() *Package { return a.Entity.Package }
func (a *OneOfArb) Package() *Package     { return a.Entity.Package }
func (a *FieldArb) Package() *Package     { return a.Entity.Package }

func (a *MsgOrEnumArb) L10n() *L10n { return a.Package().L10n }
func (a *EnumValueArb) L10n() *L10n { return a.Package().L10n }
func (a *OneOfArb) L10n() *L10n     { return a.Package().L10n }
func (a *FieldArb) L10n() *L10n     { return a.Package().L10n }

func (a *MsgOrEnumArb) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "":
		v = a.Extension.Value
	case "Label":
		v = a.Extension.Label
	case "Helper":
		v = a.Extension.Helper
	case "Hint":
		v = a.Extension.Hint
	case "Prefix":
		v = a.Extension.Prefix
	case "Suffix":
		v = a.Extension.Suffix
	default:
		log.Fatalf("Message or Enum[%s] asset name typo: %s",
			a.Entity().DartName,
			assetName)
	}
	return v
}

func (a *OneOfArb) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "Label":
		v = a.Extension.Label
	case "Prefix":
		v = a.Extension.Prefix
	case "Suffix":
		v = a.Extension.Suffix
	default:
		log.Fatalf("oneof[%s.%s] asset name typo: %s",
			a.Entity.Message.DartName,
			a.Entity.DartName,
			assetName)
	}
	return v
}

func (a *FieldArb) DefaultAssetValue(assetName string) string {
	var v string
	switch assetName {
	case "Label":
		v = a.Extension.Label
	case "Helper":
		v = a.Extension.Helper
	case "Hint":
		v = a.Extension.Hint
	case "Prefix":
		v = a.Extension.Prefix
	case "Suffix":
		v = a.Extension.Suffix
	case "BoolTrue":
		v = a.Extension.BoolTrue
	case "BoolFalse":
		v = a.Extension.BoolFalse
	default:
		log.Fatalf("field[%s.%s] asset name typo: %s",
			a.Entity.Message.DartName,
			a.Entity.DartName,
			assetName)
	}
	return v
}

func (a *MsgOrEnumArb) ResourceId() string {
	prefix := MsgAssetAsFieldPrefix
	if a.IsEnum {
		prefix = EnumAssetAsFieldPrefix
	}
	return prefix + a.Entity().DartName.String()
}

func (a *EnumValueArb) ResourceId() string {
	return EnumValueAssetAsFieldPrefix +
		a.Entity.Enum.DartName.String() +
		a.Entity.DartName.UpperCamelCase().String()
}

func (a *OneOfArb) ResourceId() string {
	return OneOfAssetAsFieldPrefix +
		a.Entity.Message.DartName.String() +
		a.Entity.DartName.UpperCamelCase().String()
}

func (a *FieldArb) ResourceId() string {
	return FieldAssetAsFieldPrefix +
		a.Entity.Message.DartName.String() +
		a.Entity.DartName.UpperCamelCase().String()
}

func (a *MsgOrEnumArb) addValueResource() {
	pa := a.L10n().Arb
	value := dartNameAsValueIfEmpty(a.Extension.Value, a.Entity().DartName)
	desc := a.Desc()
	if a.IsEnum {
		desc = "[enum] " + desc
	}
	attr := &arb.ArbAttributes{Type: "text", Description: desc}
	r := arb.NewResource(pa, a.ResourceId(), value, attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *EnumValueArb) addValueResource() {
	pa := a.L10n().Arb
	value := dartNameAsValueIfEmpty(a.Extension.Value, a.Entity.DartName)
	attr := &arb.ArbAttributes{Type: "text", Description: "[enum] " + a.Extension.Desc}
	r := arb.NewResource(pa, a.ResourceId(), value, attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *OneOfArb) addValueResource() {
	pa := a.L10n().Arb
	value := dartNameAsValueIfEmpty(a.Extension.Value, a.Entity.DartName)
	attr := &arb.ArbAttributes{Type: "text", Description: "[oneof] " + a.Desc()}
	r := arb.NewResource(pa, a.ResourceId(), value, attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *FieldArb) addValueResource() {
	pa := a.L10n().Arb
	value := a.Extension.Value
	if value == "$" {
		value = dartNameAsValueIfEmpty("", a.Entity.DartName)
	} else if value == "" {
		if a.Ref != nil {
			value = dartNameAsValueIfEmpty(a.Ref.Extension.Value, a.Ref.Entity().DartName)
		} else {
			value = dartNameAsValueIfEmpty("", a.Entity.DartName)
		}
	}
	attr := &arb.ArbAttributes{Type: "text", Description: a.Desc()}
	r := arb.NewResource(pa, a.ResourceId(), value, attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *MsgOrEnumArb) addAssetResource(assetName string) {
	pa := a.L10n().Arb
	rid := a.ResourceId()

	desc := assetName + " of " + rid
	if a.IsEnum {
		desc = "[enum] " + desc
	}
	attr := &arb.ArbAttributes{Type: "text", Description: desc}

	r := arb.NewResource(pa, rid+assetName, a.DefaultAssetValue(assetName), attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *OneOfArb) addAssetResource(assetName string) {
	pa := a.L10n().Arb
	rid := a.ResourceId()

	attr := &arb.ArbAttributes{
		Type:        "text",
		Description: "[oneof] " + assetName + " of " + rid,
	}

	r := arb.NewResource(pa, rid+assetName, a.DefaultAssetValue(assetName), attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *FieldArb) addAssetResource(assetName string) {
	assetValue := a.DefaultAssetValue(assetName)
	if assetValue == "-" {
		return
	}

	pa := a.L10n().Arb
	rid := a.ResourceId()
	desc := assetName + " of " + rid
	if a.Ref != nil {
		desc += ". Set empty to redirect to " + assetName + " of " +
			a.Ref.L10n().ClassName + "." + a.Ref.ResourceId()
	}

	attr := &arb.ArbAttributes{
		Type:        "text",
		Description: desc,
	}

	r := arb.NewResource(pa, rid+assetName, assetValue, attr)
	pa.Resources = append(pa.Resources, r)
}

func (a *MsgOrEnumArb) addAssetsResources() {
	a.addValueResource()
	a.addAssetResource("Label")
	a.addAssetResource("Helper")
	a.addAssetResource("Hint")
	a.addAssetResource("Prefix")
	a.addAssetResource("Suffix")

	if a.IsEnum {
		for _, c := range a.Enum.Values {
			c.Arb.addValueResource()
		}
	} else {
		for _, c := range a.Message.Fields {
			c.Arb.addAssetsResources()
		}
		for _, c := range a.Message.OneOfs {
			c.Arb.addAssetsResources()
		}
	}
}

func (a *OneOfArb) addAssetsResources() {
	a.addValueResource()
	a.addAssetResource("Label")
	a.addAssetResource("Prefix")
	a.addAssetResource("Suffix")
}

func (a *FieldArb) addAssetsResources() {
	a.addValueResource()
	a.addAssetResource("Label")
	a.addAssetResource("Helper")
	a.addAssetResource("Hint")
	a.addAssetResource("Prefix")
	a.addAssetResource("Suffix")
	a.addAssetResource("BoolTrue")
	a.addAssetResource("BoolFalse")
}

// Desc only used to generate arb description.
func (a *MsgOrEnumArb) Desc() string {
	s := getDescOrLabel(a.Extension.Desc, a.Extension.Label)
	s = dartNameAsValueIfEmpty(s, a.PgsEntity().Name())
	return s
}

func (a *OneOfArb) Desc() string {
	s := getDescOrLabel(a.Extension.Desc, a.Extension.Label)
	s = dartNameAsValueIfEmpty(s, a.Entity.Pgs.Name())
	return s
}

// Desc only used to generate arb description.
func (a *FieldArb) Desc() string {
	s := a.Extension.Desc
	switch s {
	case "":
		s = a.Ref.Desc()
	case "$":
		s = getAssetOrRef(a.Extension.Label, a.Ref.Extension.Label)
	default:
	}

	s = dartNameAsValueIfEmpty(s, a.Entity.Pgs.Name())
	return s
}

func (a *MsgOrEnumArb) getRawArbString(lang *arb.Arb, assetName string) string {
	return getRawArbString(lang, a.ResourceId(), assetName)
}

func (a *EnumValueArb) getRawArbString(lang *arb.Arb, assetName string) string {
	return getRawArbString(lang, a.ResourceId(), assetName)
}

func (a *OneOfArb) getRawArbString(lang *arb.Arb, assetName string) string {
	return getRawArbString(lang, a.ResourceId(), assetName)
}

func getRawArbString(lang *arb.Arb, resourceId, assetName string) string {
	r, ok := lang.ResourceMap()[resourceId+assetName]
	if !ok || r.Value == "" {
		return "null"
	}
	return r.Value
}

func (a *FieldArb) getRawArbString(lang *arb.Arb, assetName string) string {
	if a.IsAssetDisabled(assetName) {
		return "null"
	}

	r, ok := lang.ResourceMap()[a.ResourceId()+assetName]
	if !ok || r.Value == "" {
		if a.Ref != nil {
			a.UseRef = true
			return a.refResourceId() + assetName
		}
		return "null"
	}

	return dart.RawString(r.Value)
}

func (a *MsgOrEnumArb) setDartTo(lang *arb.Arb) {
	if a.IsEnum || a.Message.Arb.Extension.GetIgnore() {
		return
	}

	for _, f := range a.Message.Fields {
		f.Arb.setDartTo(lang)
	}
}

var fieldAssetNamesToSetDart = []string{"", "Label", "Helper", "Hint", "Prefix", "Suffix"}

func (a *FieldArb) setDartTo(lang *arb.Arb) {
	if a.Extension.GetIgnore() {
		return
	}

	for _, assetName := range fieldAssetNamesToSetDart {
		if a.IsAssetDisabled(assetName) {
			continue
		}

		r := lang.ResourceMap()[a.ResourceId()+assetName]
		if r.Value != "" {
			continue
		}

		r.Value = "{redirect}"
		dartInfo := arb.ArbLangInfo{Lang: "dart"}

		if a.Ref == nil {
			dartInfo.Replace = "null"
		} else if a.RefSamePkg() {
			dartInfo.Replace = a.Ref.ResourceId() + assetName
		} else {
			a.UseRef = true
			dartInfo.Replace = "&&&." + a.Ref.L10n().ClassName + "." +
				a.Ref.ResourceId() + assetName
			dartInfo.Import = a.Ref.L10n().FileName
		}

		r.Attr().Placeholders = arb.ArbPlaceholders{
			&arb.ArbPlaceholder{
				Name:      "redirect",
				LangInfos: arb.ArbLangInfos{&dartInfo},
			},
		}
	}
}
