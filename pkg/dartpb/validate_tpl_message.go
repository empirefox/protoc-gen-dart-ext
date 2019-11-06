package dartpb

import "github.com/empirefox/protoc-gen-dart-ext/pkg/dart"

func (vf *ValidateField) MessageFullValidatorClass() (dart.Qualifier, error) {
	target := vf.Pgs.Type().Embed()
	targetClass := vf.File.Dart.MessageNames(target).ValidatorName()
	return vf.Validators().ImportManager.ValidateFileDot(target, targetClass)
}
