package dartpb

import "github.com/empirefox/protoc-gen-dart-ext/pkg/dart"

func (vf *ValidateField) BoolL10nValue(v bool) dart.Qualifier {
	if v {
		return vf.L10nField() + "BoolTrue"
	}
	return vf.L10nField() + "BoolFalse"
}
