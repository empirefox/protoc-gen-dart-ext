package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateOneOf struct {
	*OneOf
	Required bool
}

func (nty *ValidateOneOf) Names() dart.OneOfNames {
	return nty.File.Dart.OneOfNames(nty.Pgs)
}

func (nty *ValidateOneOf) FullPbWhichEnum() (dart.Qualifier, error) {
	return nty.Validators().ImportManager.PbFileDot(nty.Pgs, nty.Names().OneofEnumName())
}
