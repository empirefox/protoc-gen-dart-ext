package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateMessage struct {
	*Message
	dart.ImportManagerCommonFiles
	Disabled bool
}

func (im *ValidateMessage) L10nType() dart.Qualifier { return im.Validators().L10nType() }

func (nty *ValidateMessage) FullPbClass() (dart.Qualifier, error) {
	return nty.Validators().FullPbClass(nty.Pgs)
}

func (nty *ValidateMessage) FullL10nClass() (dart.Qualifier, error) {
	return nty.ImportManager().L10nAsInstanceType(nty.Pgs)
}
