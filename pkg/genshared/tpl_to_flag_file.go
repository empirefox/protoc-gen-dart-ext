package genshared

import (
	"flag"
	"log"
	"os"
	"text/template"

	multierror "github.com/hashicorp/go-multierror"
)

type Renderer interface {
	Open()
	Render(data interface{}) error
	Close() error
}

//go:generate pie MultiRenderer.Append.Extend
type MultiRenderer []Renderer

func (mr MultiRenderer) Open() {
	for _, r := range mr {
		r.Open()
	}
}
func (mr MultiRenderer) Render(data interface{}) (err error) {
	for _, r := range mr {
		if err = r.Render(data); err != nil {
			return
		}
	}
	return
}
func (mr MultiRenderer) Close() error {
	var err, result error
	for _, r := range mr {
		if err = r.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

type Template struct {
	*template.Template
	RootFromData func(v interface{}) interface{}
	IgnoreIfData func(v interface{}) bool
}

type GroupRenderer []*SingleRenderer

func NewTplGroupRenderer(tpls ...*Template) GroupRenderer {
	rs := make(GroupRenderer, len(tpls))
	for i, tpl := range tpls {
		rs[i] = NewTplSingleRenderer(tpl)
	}
	return rs
}

func NewGoTplGroupRenderer(tpls ...*template.Template) GroupRenderer {
	rs := make(GroupRenderer, len(tpls))
	for i, tpl := range tpls {
		rs[i] = NewTplSingleRenderer(&Template{Template: tpl})
	}
	return rs
}

func (rs GroupRenderer) Render(data interface{}) (err error) {
	for _, r := range rs {
		if r.tpl.IgnoreIfData != nil && r.tpl.IgnoreIfData(data) {
			r.ignored = true
			continue
		}
		if r.tpl.RootFromData != nil {
			data = r.tpl.RootFromData(data)
		}
		if err = r.Render(data); err != nil {
			return
		}
	}
	return
}

func (rs GroupRenderer) Open() {
	for _, r := range rs {
		r.Open()
	}
}

func (rs GroupRenderer) Close() error {
	var err, result error
	for _, r := range rs {
		if err = r.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

type SingleRenderer struct {
	f   *os.File
	of  *OutputFile
	tpl *Template

	ignored bool
}

func NewTplSingleRenderer(tpl *Template) *SingleRenderer {
	return NewSingleRenderer(tpl.Name(), tpl)
}

func NewSingleRenderer(flagName string, tpl *Template) *SingleRenderer {
	return &SingleRenderer{
		of:  NewOutputFile(flagName),
		tpl: tpl,
	}
}

func (r *SingleRenderer) Render(data interface{}) error {
	r.Open()
	return r.tpl.Execute(r.f, data)
}

func (r *SingleRenderer) Open() {
	if r.f == nil {
		r.f = r.of.Open()
	}
}

func (r *SingleRenderer) Close() error {
	if r.f == nil {
		return nil
	}
	if r.ignored {
		r.ignored = false
		defer os.Remove(r.of.Path)
	}
	return r.f.Close()
}

type OutputFile struct {
	Flag string
	Path string
}

func NewOutputFile(flagName string) *OutputFile {
	of := OutputFile{Flag: flagName}
	flag.StringVar(&of.Path, flagName, "", "the required "+flagName+" output file")
	return &of
}

func (of *OutputFile) Open() *os.File {
	if of.Path == "" {
		log.Fatalf("%s argument must be set\n", of.Flag)
	}
	f, err := os.OpenFile(of.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("OpenFile: %v\n", err)
	}
	return f
}

type VariantPathFunc func(pathPrefix, variant string) string
type VariantDataFunc func(data interface{}, variant string) interface{}

type VariantRenderer struct {
	flag         string
	tpl          *Template
	variantsFunc func() []string

	dataFunc VariantDataFunc
	variants []string

	pathFunc   VariantPathFunc
	pathPrefix string

	fs     []*os.File
	opened bool

	ignored bool
}

func NewGoTplVariantRenderer(
	tpl *template.Template,
	variantsFunc func() []string,
	dataFunc VariantDataFunc,
	makePath VariantPathFunc) *VariantRenderer {
	return NewTplVariantRenderer(&Template{Template: tpl}, variantsFunc, dataFunc, makePath)
}

func NewTplVariantRenderer(
	tpl *Template,
	variantsFunc func() []string,
	dataFunc VariantDataFunc,
	makePath VariantPathFunc) *VariantRenderer {
	return NewVariantRenderer(tpl.Name(), tpl, variantsFunc, dataFunc, makePath)
}

func NewVariantRenderer(
	flagName string,
	tpl *Template,
	variantsFunc func() []string,
	dataFunc VariantDataFunc,
	makePath VariantPathFunc) *VariantRenderer {
	vr := VariantRenderer{
		flag:         flagName,
		tpl:          tpl,
		variantsFunc: variantsFunc,
		pathFunc:     makePath,
		dataFunc:     dataFunc,
	}
	flag.StringVar(&vr.pathPrefix, flagName, "", "the required "+flagName+" output files pathFunc prefix")
	return &vr
}

type VariantRenderData struct {
	Data    interface{}
	Variant string
}

func (vr *VariantRenderer) Render(data interface{}) (err error) {
	vr.Open()

	if vr.tpl.IgnoreIfData != nil && vr.tpl.IgnoreIfData(data) {
		vr.ignored = true
		return
	}

	if vr.tpl.RootFromData != nil {
		data = vr.tpl.RootFromData(data)
	}

	for i, v := range vr.getVariants() {
		var result interface{}
		if vr.dataFunc != nil {
			result = vr.dataFunc(data, v)
		} else {
			result = VariantRenderData{data, v}
		}
		if err = vr.tpl.Execute(vr.fs[i], result); err != nil {
			return
		}
	}
	return
}

func (vr *VariantRenderer) Open() {
	if vr.opened {
		return
	}
	vr.opened = true

	for i, v := range vr.getVariants() {
		f, err := os.OpenFile(vr.pathFunc(vr.pathPrefix, v), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatalf("OpenFile: %v\n", err)
		}
		vr.fs[i] = f
	}
}

func (vr *VariantRenderer) Close() error {
	if !vr.opened {
		return nil
	}

	if vr.fs == nil {
		return nil
	}

	var err, result error
	for _, f := range vr.fs {
		if err = f.Close(); err != nil {
			result = multierror.Append(result, err)
		}
		if vr.ignored {
			os.Remove(f.Name())
		}
	}
	vr.fs = nil
	return result
}

func (vr *VariantRenderer) getVariants() []string {
	if vr.variants == nil {
		vr.variants = vr.variantsFunc()
		vr.fs = make([]*os.File, len(vr.variants))
	}
	return vr.variants
}
