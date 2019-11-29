package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateOneOf struct {
	*OneOf
	dart.ImportManagerCommonFiles
	Required bool
}

func (nty *ValidateOneOf) InfoAccessor() dart.Qualifier { return nty.Validators().InfoAccessor }
func (nty *ValidateOneOf) L10nAccessor() dart.Qualifier { return nty.Validators().L10nAccessor }

func (im *ValidateOneOf) FullPbWhichEnum() (dart.Qualifier, error) {
	return im.ImportManager().FullPbWhichEnum(im.Pgs)
}
