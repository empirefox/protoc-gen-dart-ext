package util

import (
	"sort"

	pgs "github.com/lyft/protoc-gen-star"
)

func FilesPackages(files map[string]pgs.File) []pgs.Package {
	m := make(map[pgs.Package]struct{}, len(files))
	for _, t := range files {
		m[t.Package()] = struct{}{}
	}
	pkgs := make([]pgs.Package, 0, len(m))
	for pkg := range m {
		pkgs = append(pkgs, pkg)
	}
	sort.Slice(pkgs, func(i, j int) bool { return pkgs[i].ProtoName() < pkgs[j].ProtoName() })
	return pkgs
}
