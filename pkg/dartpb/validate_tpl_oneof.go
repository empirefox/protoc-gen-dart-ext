package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateOneOf struct {
	*OneOf
	Required bool
}

func (nty *ValidateOneOf) InfoAccessor() dart.Qualifier { return nty.Validators().InfoAccessor }
func (nty *ValidateOneOf) L10nAccessor() dart.Qualifier { return nty.Validators().L10nAccessor }
func (nty *ValidateOneOf) PgdeFile() *dart.ImportFile   { return nty.Validators().PgdeFile }

func (nty *ValidateOneOf) Names() dart.OneOfNames {
	return nty.File.Dart.OneOfNames(nty.Pgs)
}

func (nty *ValidateOneOf) FullPbWhichEnum() (dart.Qualifier, error) {
	return nty.Validators().ImportManager.PbFileDot(nty.Pgs, nty.Names().OneofEnumName())
}
