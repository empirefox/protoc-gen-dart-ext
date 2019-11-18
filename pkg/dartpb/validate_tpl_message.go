package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	pgs "github.com/lyft/protoc-gen-star"
)

func (vf *ValidateField) MessageFullValidatorClass() (dart.Qualifier, error) {
	pgsType := vf.Pgs.Type()
	var target pgs.Message
	if pgsType.IsRepeated() || pgsType.IsMap() {
		target = pgsType.Element().Embed()
	} else {
		target = pgsType.Embed()
	}
	targetClass := vf.File.Dart.MessageNames(target).ValidatorName()
	return vf.Validators().ImportManager.ValidateFileDot(target, targetClass)
}
