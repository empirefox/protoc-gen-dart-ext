package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/envoyproxy/protoc-gen-validate/templates/shared"
	pgs "github.com/lyft/protoc-gen-star"
)

func (f *ValidateField) MessageFullValidatorClass() (dart.Qualifier, error) {
	target := f.ElementOrEmbed()
	yes, err := shared.Disabled(target)
	if err != nil {
		return "", err
	}

	if yes {
		return "", nil
	}

	name := f.File.Dart.MessageNames(target).ValidatorName()
	return f.Validators().ImportManager.ValidateFileDot(target, name)
}

func (f *ValidateField) IsOfMessageType() bool {
	return f.Pgv.Field.Type().ProtoType() == pgs.MessageT
}
