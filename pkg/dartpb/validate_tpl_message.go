package dartpb

import "github.com/empirefox/protoc-gen-dart-ext/pkg/dart"

func (vf *ValidateField) MessageFullValidatorClass() (dart.Qualifier, error) {
	targetMessage := vf.Package.Group.PgsToMsg[vf.Pgs.Type().Embed()]
	targetValidatorClass := targetMessage.Validator().ClassName
	as, err := vf.Validator().ImportValidatorAs(targetMessage.Validator())
	if err != nil {
		return "", nil
	}

	return as.Append(targetValidatorClass), nil
}
