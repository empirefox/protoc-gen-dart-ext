package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
)

func (vf *ValidateField) RepeatedElemIsBytes() bool {
	return vf.Pgs.Type().Element().ProtoType() == pgs.BytesT
}

func (vf *ValidateField) RepeatedElemType() (dart.Qualifier, error) {
	return vf.ImportManager().TypeForNonRangeFieldType(vf.Pgs.Type().Element())
}
