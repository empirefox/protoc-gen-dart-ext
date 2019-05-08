package arb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
)

//easyjson:json
type ArbImport struct {
	// Id the resource id of arb who need to be resolved
	Id string `json:"id,omitempty"`

	// Ref ref to attr.Export
	Ref string `json:"import,omitempty"`

	// Assets to import
	Assets []string `json:"assets,omitempty"`
}

type ResolveEntry struct {
	Ref string

	// Import the remote path
	Import string

	// Entity the remote entity
	Entity string

	// Field the remote field
	Field string
}

type Resolver struct {
	entries map[string]*ResolveEntry
}

func NewResolver(e *exports.Exports) *Resolver {
	entries := make(map[string]*ResolveEntry, 256)
	for impPath, pkg := range e.Packages {
		for _, entity := range pkg.Entities {
			for _, field := range entity.Fields {
				nty := &ResolveEntry{
					Ref:    field.Ref,
					Import: impPath,
					Entity: entity.Name,
					Field:  field.Name,
				}
				entries[field.Ref] = nty
			}
		}
	}
	return &Resolver{
		entries: entries,
	}
}

func (r *Resolver) Resolve(ref string) (*ResolveEntry, error) {
	nty, ok := r.entries[ref]
	if !ok {
		return nil, fmt.Errorf("Ref not found: %s", ref)
	}
	return nty, nil
}

func (a *Arb) ExportResolver() *Resolver {
	selfExports := a.ExportProto()
	if selfExports == nil {
		return nil
	}
	return NewResolver(&exports.Exports{
		Packages: map[string]*exports.Package{
			"": &exports.Package{
				Entities: []*exports.Entity{selfExports},
			},
		},
	})
}

func (a *Arb) ExportProto() *exports.Entity {
	fields := make([]*exports.Field, 0, len(a.Resources))
	for _, r := range a.Resources {
		export := r.Attr().Export
		if export != "" {
			fields = append(fields, &exports.Field{
				Ref:  export,
				Name: r.Id,
			})
		}
	}
	if len(fields) == 0 {
		return nil
	}
	return &exports.Entity{Name: a.Entity, Fields: fields}
}
