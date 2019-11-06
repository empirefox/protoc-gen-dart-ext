package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateMessage struct {
	*Message
	Disabled bool
}

func (nty *ValidateMessage) FullL10nClass() (dart.Qualifier, error) {
	return nty.Validators().ImportManager.L10nAsInstanceType(nty.Pgs)
}
