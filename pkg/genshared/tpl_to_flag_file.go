package genshared

import (
	"flag"
	"log"
	"os"
	"text/template"

	multierror "github.com/hashicorp/go-multierror"
)

type Template struct {
	*template.Template
	RootFromData func(v interface{}) interface{}
	IgnoreIfData func(v interface{}) bool
}

type Renderers []*Renderer

func NewTplRenderers(tpls ...*Template) Renderers {
	rs := make(Renderers, len(tpls))
	for i, tpl := range tpls {
		rs[i] = NewTplRenderer(tpl)
	}
	return rs
}

func NewGoTplRenderers(tpls ...*template.Template) Renderers {
	rs := make(Renderers, len(tpls))
	for i, tpl := range tpls {
		rs[i] = NewTplRenderer(&Template{Template: tpl})
	}
	return rs
}

func (rs Renderers) Render(data interface{}) (err error) {
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

func (rs Renderers) OpenAll() {
	for _, r := range rs {
		r.File()
	}
}

func (rs Renderers) Close() error {
	var err, result error
	for _, r := range rs {
		if err = r.Close(); err != nil {
			result = multierror.Append(result, err)
		}
	}
	return result
}

type Renderer struct {
	f   *os.File
	of  *OutputFile
	tpl *Template

	ignored bool
}

func NewTplRenderer(tpl *Template) *Renderer {
	return NewRenderer(tpl.Name(), tpl)
}

func NewRenderer(flagName string, tpl *Template) *Renderer {
	return &Renderer{
		of:  NewOutputFile(flagName),
		tpl: tpl,
	}
}

func (r *Renderer) Render(data interface{}) error {
	return r.tpl.Execute(r.File(), data)
}

func (r *Renderer) File() *os.File {
	if r.f == nil {
		r.f = r.of.Open()
	}
	return r.f
}

func (r *Renderer) Close() error {
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
