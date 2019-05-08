package dart

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/imports"
)

//go:generate pie ImportEntries.ToStrings.Unselect
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

func (s *ImportItems) DartType(imp, typ string) string {
	as := s.Add(imp)
	if as == 0 {
		return typ
	}
	return fmt.Sprintf("$i%d.%s", as, typ)
}

func (s *ImportItems) AddNoAs(imp string) {
	if s.SetNoAs(imp) {
		return
	}

	s2 := make(ImportItems, len(*s)+1)
	s2[0] = &ImportItem{Path: imp}
	copy(s2[1:], *s)
	*s = s2
}

func (s *ImportItems) SetNoAs(imp string) bool {
	if imp == "" {
		return false
	}

	for _, item := range *s {
		if item.Path == imp {
			item.As = 0
			return true
		}
	}
	return false
}

func (s *ImportItems) Add(imp string) int {
	if imp == "" {
		return 0
	}

	for _, item := range *s {
		if item.Path == imp {
			return item.As
		}
	}

	as := len(*s) + 1
	*s = append(*s, &ImportItem{Path: imp, As: as})
	return as
}

func (s ImportItems) ToDartSource() string {
	if len(s) == 0 {
		return ""
	}
	var builder strings.Builder
	builder.Grow(1 << 10)
	for _, item := range s {
		builder.WriteString("import '")
		builder.WriteString(item.Path)
		if item.As == 0 {
			builder.WriteString("';\n")
		} else {
			builder.WriteString("' as $i")
			builder.WriteString(strconv.Itoa(item.As))
			builder.WriteString(";\n")
		}
	}
	return builder.String()
}

type ImportItem struct {
	Path string
	As   int
}

type Resource struct {
	*arb.ArbResource
	Resolved *Resolved
}

func (r Resource) TplAsString(varname string) string {
	lp := r.Attr().Placeholders.LangParam("dart", varname)
	if lp.Info == "String" {
		return varname
	}
	return fmt.Sprintf(`'$%s'`, varname)
}

func (r Resource) TplSelectCaseCond(varname, key string) string {
	// varname: form
	// key: zero one two...
	lp := r.Attr().Placeholders.LangParam("dart", varname)
	if lp.Info == "String" {
		return RawString(key)
	}
	if lp.Info == "" {
		return r.Resolved.Imports.DartType(lp.Import, key)
	}
	return r.Resolved.Imports.DartType(lp.Import, lp.Info) + "." + key
}

type Resolved struct {
	*arb.Arb
	Imports ImportItems
	Entries ImportEntries
}

func ResolveWithExports(es *exports.Exports, a *arb.Arb) (*Resolved, error) {
	var resolver *arb.Resolver
	if es != nil {
		resolver = arb.NewResolver(es)
	}
	return Resolve(resolver, a)
}

func Resolve(resolver *arb.Resolver, a *arb.Arb) (*Resolved, error) {
	selfResolver := a.ExportResolver()

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
		var nty *arb.ResolveEntry
		if selfResolver != nil {
			nty, err = selfResolver.Resolve(r.Ref)
		} else {
			err = fmt.Errorf("resolve failed: %s", r.Ref)
		}
		if err != nil {
			self = false
			if resolver != nil {
				nty, err = resolver.Resolve(r.Ref)
			}
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
	paths := ra.Entries.
		Unselect(func(nty *ImportEntry) bool { return nty.self }).
		ToStrings(func(nty *ImportEntry) string { return nty.resolved.Import }).
		Unique().
		Sort()

	ra.Imports = make(ImportItems, len(paths))
	for i, p := range paths {
		ra.Imports[i] = &ImportItem{
			Path: p,
			As:   i + 1,
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

	for _, r := range a.Resources {
		for _, p := range r.Attr().Placeholders.LangParams("dart") {
			ra.Imports.Add(p.Import)
		}
	}

	return &ra, nil
}

func (ra *Resolved) Resources() []Resource {
	rs := make([]Resource, len(ra.Arb.Resources))
	for i, r := range ra.Arb.Resources {
		rs[i] = Resource{ArbResource: r, Resolved: ra}
	}
	return rs
}

func (ra *Resolved) WithArb(a *arb.Arb) *Resolved {
	return &Resolved{
		Arb:     a,
		Imports: ra.Imports,
		Entries: ra.Entries,
	}
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
