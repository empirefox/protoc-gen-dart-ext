package main

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/module"
	pgs "github.com/lyft/protoc-gen-star"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(
		module.New(),
	).Render()
}
