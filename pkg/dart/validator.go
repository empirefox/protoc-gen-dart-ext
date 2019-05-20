package dart

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/exports"
)

type ValidatorResolveEntry struct {
	// Ref the remote FullyQualifiedName
	Ref string

	// Import the remote path
	Import string

	Entity string
}

type ValidatorResolver struct {
	entries map[string]*ValidatorResolveEntry
}

func NewValidatorResolver(e *exports.Exports) *ValidatorResolver {
	entries := make(map[string]*ValidatorResolveEntry, 256)
	for _, v := range e.Validators {
		for _, full := range v.Entities {
			nty := &ValidatorResolveEntry{
				Ref:    full,
				Import: v.Path,
				Entity: entityFromFull(full),
			}
			entries[full] = nty
		}
	}
	return &ValidatorResolver{
		entries: entries,
	}
}

func (r *ValidatorResolver) Resolve(ref string) (*ValidatorResolveEntry, error) {
	nty, ok := r.entries[ref]
	if !ok {
		return nil, fmt.Errorf("Ref not found for validator: %s", ref)
	}
	return nty, nil
}

func entityFromFull(path string) string {
	path = filepath.Base(path)
	for i := len(path) - 1; i >= 0 && !os.IsPathSeparator(path[i]); i-- {
		if path[i] == '.' {
			return path[i+1:]
		}
	}
	return path
}
