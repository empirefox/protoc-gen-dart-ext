package main

import (
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type StaticModule struct {
	*pgs.ModuleBase
	ctx       pgsgo.Context
	files     map[string][]byte
	generates map[string][]byte
}

func Static(files map[string][]byte) *StaticModule {
	return &StaticModule{
		ModuleBase: &pgs.ModuleBase{},
		files:      files,
		generates:  make(map[string][]byte, len(files)),
	}
}

func (p *StaticModule) InitContext(c pgs.BuildContext) {
	p.ModuleBase.InitContext(c)
	p.ctx = pgsgo.InitContext(c.Parameters())
	for name := range p.Parameters() {
		content, ok := p.files[name]
		if !ok {
			p.Failf("file name not supported: %s", name)
		}
		p.generates[name] = content
	}
}

func (p *StaticModule) Name() string { return "static" }

func (p *StaticModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {
	for name, content := range p.generates {
		p.AddGeneratorFile(name, string(content))
	}
	return p.Artifacts()
}
