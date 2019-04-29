package module

import (
	"path/filepath"

	pgs "github.com/lyft/protoc-gen-star"
)

type PkgData struct {
	pgs.Package
	Ext string
}

func (d *PkgData) PackageDir() string {
	return d.Files()[0].InputPath().Dir().String()
}

func (d *PkgData) OutputName() string {
	return pgs.FilePath(d.ProtoName()).SetExt(d.Ext).String()
}

func (d *PkgData) OutputPath() string {
	return filepath.Join(d.PackageDir(), d.OutputName())
}

func (d *PkgData) AllEnums() (a []pgs.Enum) {
	for _, f := range d.Files() {
		a = append(a, f.AllEnums()...)
	}
	return
}

func (d *PkgData) AllMessages() (a []pgs.Message) {
	for _, f := range d.Files() {
		a = append(a, f.AllMessages()...)
	}
	return
}
