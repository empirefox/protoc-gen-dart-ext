package util

import (
	"sort"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc"
	pgs "github.com/lyft/protoc-gen-star"
)

func ProtoFilesReflect(pkgs map[string]pgs.Package) (map[string]*desc.FileDescriptor, error) {
	total := 0
	for _, pkg := range pkgs {
		total += len(pkg.Files())
	}

	files := make([]*descriptor.FileDescriptorProto, total)
	i := 0
	for _, pkg := range pkgs {
		for _, f := range pkg.Files() {
			files[i] = f.Descriptor()
			i++
		}
	}

	sort.Slice(files, func(i, j int) bool {
		name := files[j].GetName()
		for _, dep := range files[i].Dependency {
			if dep == name {
				return false
			}
		}
		return true
	})

	return desc.CreateFileDescriptors(files)
}
