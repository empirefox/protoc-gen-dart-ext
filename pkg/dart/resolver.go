package dart

import (
	"fmt"
	"sort"
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/imports"
)

type ImportEntries []*ImportEntry

func (s ImportEntries) ToDartSource() string {
	if len(s) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.Grow(1 << 10)
	for _, nty := range s {
		builder.WriteString("String ")
		builder.WriteString(nty.Left())
		builder.WriteString(" => ")
		builder.WriteString(nty.Right())
		builder.WriteString(";\n")
	}
	return builder.String()
}

type ImportEntry struct {
	reserve  *arb.ArbImport
	resolved *arb.ResolveEntry
	asset    imports.Field_Asset
	self     bool
	imp      *ImportItem
}

func (nty *ImportEntry) Left() string {
	if nty.self {
		return "get " + nty.reserve.Id
	}
	return fmt.Sprintf("%s(%s.%sLocalization l10n)",
		nty.reserve.Id, nty.imp.As, nty.resolved.Entity)
}

func (nty *ImportEntry) Right() string {
	if nty.self {
		return nty.resolved.Field
	}
	return "l10n." + nty.resolved.Field
}

type ImportItems []*ImportItem

func (s ImportItems) ToDartSource() string {
	if len(s) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.Grow(1 << 10)
	for _, item := range s {
		builder.WriteString("import '")
		builder.WriteString(item.Path)
		builder.WriteString("' as ")
		builder.WriteString(item.As)
		builder.WriteString(";\n")
	}
	return builder.String()
}

type ImportItem struct {
	Path string
	As   string
}

type Resolved struct {
	*arb.Arb
	Imports ImportItems
	Entries ImportEntries
}

func Resolve(resolver *arb.Resolver, a *arb.Arb) (*Resolved, error) {
	selfResolver := arb.NewResolver(&exports.Exports{
		Packages: map[string]*exports.Package{
			"": &exports.Package{
				Entities: []*exports.Entity{a.ExportProto()},
			},
		},
	})

	ra := Resolved{
		Arb:     a,
		Entries: make(ImportEntries, 0, len(a.Imports)*8),
	}

	for _, r := range a.Imports {
		assets, err := imports.NewAssets(r.Assets)
		if err != nil {
			return nil, err
		}

		self := true
		nty, err := selfResolver.Resolve(r.Ref)
		if err != nil {
			self = false
			nty, err = resolver.Resolve(r.Ref)
		}
		if err != nil {
			return nil, err
		}

		for _, asset := range assets {
			ra.Entries = append(ra.Entries, &ImportEntry{
				reserve:  r,
				resolved: nty,
				asset:    asset,
				self:     self,
			})
		}
	}

	// get sorted import path list
	pathMap := make(map[string]struct{}, len(ra.Entries))
	for _, nty := range ra.Entries {
		if !nty.self {
			pathMap[nty.resolved.Import] = struct{}{}
		}
	}
	paths := make([]string, len(pathMap))
	i := 0
	for p := range pathMap {
		paths[i] = p
		i++
	}
	sort.Strings(paths)

	ra.Imports = make(ImportItems, len(paths))
	for i, p := range paths {
		ra.Imports[i] = &ImportItem{
			Path: p,
			As:   fmt.Sprintf("$i%d", i),
		}
	}

	// equip ImportItem to Entries
	importItemByPath := make(map[string]*ImportItem, len(paths))
	for _, item := range ra.Imports {
		importItemByPath[item.Path] = item
	}

	for _, nty := range ra.Entries {
		if !nty.self {
			nty.imp = importItemByPath[nty.resolved.Import]
		}
	}

	return &ra, nil
}

func (ra *Resolved) ImportProto() *imports.Entity {
	if ra == nil {
		return nil
	}
	fields := make([]*imports.Field, 0, len(ra.Entries))
	for _, nty := range ra.Entries {
		fields = append(fields, &imports.Field{
			Name:   nty.reserve.Id,
			Entity: nty.resolved.Entity,
			Import: nty.imp.Path,
		})
	}
	return &imports.Entity{
		Name:   ra.Entity,
		Fields: fields,
	}
}
