package dartpb

import (
	_ "unsafe"

	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

//go:linkname rulesContext github.com/envoyproxy/protoc-gen-validate/templates/shared.rulesContext
func rulesContext(pgsNty pgs.Field) (shared.RuleContext, error)
