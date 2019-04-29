package l10n

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/imports"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

const ExportImportProtoPrefix = "proto:"

func PackageToArb(d *dart.Dart, locale language.Tag, pkg pgs.Package) (*arb.Arb, error) {
	as := make([]*arb.Arb, 0, 128)
	for _, f := range pkg.Files() {
		for _, nty := range f.AllEnums() {
			a, err := EnumToArb(d, locale, nty)
			if err != nil {
				return nil, err
			}
			as = append(as, a)
		}
		for _, nty := range f.AllMessages() {
			a, err := MessageToArb(d, locale, nty)
			if err != nil {
				return nil, err
			}
			as = append(as, a)
		}
	}

	to := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       locale,
		Entity:       entityNameOfPkg(pkg),
	}
	arb.ShallowMerge(as, to)
	return to, nil
}

func entityNameOfPkg(pkg pgs.Package) string {
	pkgName := pkg.ProtoName().String()
	ext := filepath.Ext(pkgName)
	if ext != "" {
		pkgName = ext[1:]
	}
	return pgs.Name(pkgName).UpperCamelCase().String()
}

func MessageToArb(d *dart.Dart, locale language.Tag, nty pgs.Message) (*arb.Arb, error) {
	var ext Arb
	_, err := nty.Extension(E_MsgArb, &ext)
	if err != nil {
		return nil, err
	}

	if ext.Ignore {
		return nil, nil
	}

	fieldsLen := len(nty.Fields())
	fieldsAssetsLen := fieldsLen * 8
	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       locale,
		Entity:       d.NameOf(nty).String(),
		Imports:      make([]*arb.ArbImport, 0, fieldsLen),
		Nils:         make([]string, 0, fieldsAssetsLen),
		Resources:    make([]*arb.ArbResource, 0, fieldsAssetsLen+len(nty.OneOfs())*4+6),
	}

	addResource(d, a, &ext, nty, true)
	addAssets(a, a.Entity, &ext, false)

	for _, child := range nty.OneOfs() {
		var childExt OneofArb
		_, err = nty.Extension(E_OneofArb, &childExt)
		if err != nil {
			return nil, err
		}

		if childExt.Ignore {
			continue
		}
		addResource(d, a, &childExt, child, false)
		addAssets(a, d.NameOf(child).String(), &childExt, false)
	}

	for _, child := range nty.Fields() {
		var childExt FieldArb
		_, err = nty.Extension(E_FieldArb, &childExt)
		if err != nil {
			return nil, err
		}

		if childExt.Ignore {
			continue
		}

		var refEntity pgs.Entity
		var refDesc *proto.ExtensionDesc
		if child.Type().IsRepeated() || child.Type().IsMap() {
			elem := child.Type().Element()
			if elem.IsEnum() {
				refEntity, refDesc = elem.Enum(), E_EnumArb
			} else if elem.IsEmbed() {
				refEntity, refDesc = elem.Embed(), E_MsgArb
			}
		} else if child.Type().IsEnum() {
			refEntity, refDesc = child.Type().Enum(), E_EnumArb
		} else if child.Type().IsEmbed() {
			refEntity, refDesc = child.Type().Embed(), E_MsgArb
		}

		ntyName := d.NameOf(child).String()
		if refEntity != nil {
			var refExt Arb
			_, err = refEntity.Extension(refDesc, &refExt)
			if err != nil {
				return nil, err
			}

			assets := make(imports.Assets, 0, len(imports.Field_Asset_name))
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetValue, refExt.Value, imports.Field_Value)
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetLabel, refExt.Label, imports.Field_Label)
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetHelper, refExt.Helper, imports.Field_Helper)
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetHint, refExt.Hint, imports.Field_Hint)
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetPrefix, refExt.Prefix, imports.Field_Prefix)
			addAssetToRef(d, a, &assets, child, &childExt, childExt.GetSuffix, refExt.Suffix, imports.Field_Suffix)
			if len(assets) != 0 {
				a.Imports = append(a.Imports, &arb.ArbImport{
					Id:     child.Name().String(),
					Ref:    ExportImportProtoPrefix + refEntity.FullyQualifiedName(),
					Assets: assets.Strings(),
				})
			}
		} else {
			addResource(d, a, &childExt, child, false)
			addAssets(a, ntyName, &childExt, child.Type().ProtoType() == pgs.BoolT)
		}

	}

	return a, nil
}

func addAssetToRef(d *dart.Dart, a *arb.Arb, assets *imports.Assets, nty pgs.Field, ext arbExt, get func() string, ref string, asset imports.Field_Asset) {
	ntyName := d.NameOf(nty).String()
	v := get()
	if v == "" {
		if ref == "" {
			if asset == imports.Field_Value {
				// if Value, add with field name
				addResource(d, a, ext, nty, false)
				return
			}
			addToNils(a, ntyName+asset.Suffix())
			return
		}
		if ref == "-" {
			addToNils(a, ntyName+asset.Suffix())
			return
		}
		*assets = append(*assets, asset)
		return
	}

	if v == "-" {
		addToNils(a, ntyName+asset.Suffix())
		return
	}

	if asset == imports.Field_Value {
		addResource(d, a, ext, nty, false)
		return
	}

	addAsset(a, ntyName, get, asset)
}

func addToNils(a *arb.Arb, id string) {
	a.Nils = append(a.Nils, id)
}

func EnumToArb(d *dart.Dart, locale language.Tag, nty pgs.Enum) (*arb.Arb, error) {
	var ext Arb
	_, err := nty.Extension(E_EnumArb, &ext)
	if err != nil {
		return nil, err
	}

	if ext.Ignore {
		return nil, nil
	}

	a := &arb.Arb{
		LastModified: iso8601.Time{time.Now()},
		Locale:       locale,
		Entity:       d.NameOf(nty).String(),
		Resources:    make([]*arb.ArbResource, 0, len(nty.Values())+1),
	}
	addResource(d, a, &ext, nty, true)
	addAssets(a, a.Entity, &ext, false)

	for _, child := range nty.Values() {
		var childExt ValueArb
		_, err = nty.Extension(E_ValueArb, &childExt)
		if err != nil {
			return nil, err
		}

		if childExt.Ignore {
			continue
		}
		addResource(d, a, &childExt, child, false)
	}

	return a, nil
}

var _ arbExt = new(ValueArb)

type arbExt interface {
	GetValue() string
	GetDesc() string
}

var _ assetsArbExt = new(OneofArb)

type assetsArbExt interface {
	arbExt
	GetLabel() string
	GetPrefix() string
	GetSuffix() string
}

var _ entityArbExt = new(Arb)

type entityArbExt interface {
	assetsArbExt
	GetHelper() string
	GetHint() string
}

var _ fieldArbExt = new(FieldArb)

type fieldArbExt interface {
	entityArbExt
	GetBoolTrue() string
	GetBoolFalse() string
}

// addResource only add Value
func addResource(d *dart.Dart, a *arb.Arb, ext arbExt, nty pgs.Entity, export bool) {
	value := ext.GetValue()
	if value == "-" && !export {
		return
	}

	if value == "" {
		value = strings.Join(nty.Name().UpperCamelCase().Split(), " ")
	}

	attr := &arb.ArbAttributes{
		Type:        "text",
		Description: ext.GetDesc(),
	}
	if attr.Description == "$" {
		assetsExt, ok := ext.(assetsArbExt)
		if ok {
			attr.Description = assetsExt.GetLabel()
		} else {
			attr.Description = ""
		}
	}
	if export {
		attr.Export = ExportImportProtoPrefix + nty.FullyQualifiedName()
	}

	a.Resources = append(a.Resources, arb.NewResource(a, d.NameOf(nty).String(), value, attr))
}

// addAssets do not add Value
func addAssets(a *arb.Arb, ntyName string, ext arbExt, isBool bool) {
	assetsExt, ok := ext.(assetsArbExt)
	if !ok {
		return
	}

	addAsset(a, ntyName, assetsExt.GetLabel, imports.Field_Label)
	addAsset(a, ntyName, assetsExt.GetPrefix, imports.Field_Prefix)
	addAsset(a, ntyName, assetsExt.GetSuffix, imports.Field_Suffix)

	entityExt, ok := ext.(entityArbExt)
	if !ok {
		return
	}

	addAsset(a, ntyName, entityExt.GetHelper, imports.Field_Helper)
	addAsset(a, ntyName, entityExt.GetHint, imports.Field_Hint)

	if !isBool {
		return
	}

	fieldExt, ok := ext.(fieldArbExt)
	if !ok {
		return
	}

	addAsset(a, ntyName, fieldExt.GetBoolTrue, imports.Field_BoolTrue)
	addAsset(a, ntyName, fieldExt.GetBoolFalse, imports.Field_BoolFalse)
}

func addAsset(a *arb.Arb, ntyName string, get func() string, asset imports.Field_Asset) {
	if value := get(); value != "" && value != "-" {
		ar := arb.NewResource(a, ntyName+asset.Suffix(), value, &arb.ArbAttributes{
			Type:        "text",
			Description: asset.String() + " of " + ntyName,
		})

		a.Resources = append(a.Resources, ar)
	}
}
