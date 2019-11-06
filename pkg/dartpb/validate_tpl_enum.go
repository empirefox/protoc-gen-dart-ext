package dartpb

import (
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

func (vf *ValidateField) EnumType() (dart.Qualifier, error) {
	return vf.Validators().FullPbEnum(vf.Pgs.Type().Enum())
}

func (vf *ValidateField) EnumLiteralValue(v int32) (dart.Qualifier, error) {
	value, err := vf.parseEnumValue(v)
	if err != nil {
		return "", err
	}

	targetFullEnumClass, err := vf.Validators().FullPbEnum(value.Enum())
	if err != nil {
		return "", err
	}

	return targetFullEnumClass.Dot(vf.File.Dart.NameOf(value)), nil
}

func (vf *ValidateField) EnumFullL10nClass() (dart.Qualifier, error) {
	nty := vf.Pgs.Type().Enum()
	if nty.File() == vf.Pgs.File() {
		return "", nil
	}

	return vf.Validators().ImportManager.L10nFileDot(nty, vf.File.Dart.EnumNames(nty).Name())

}

func (vf *ValidateField) EnumL10nValues(s []int32) (dart.Qualifier, error) {
	nty := vf.Pgs.Type().Enum()
	l10nAccessor := vf.L10nAccessor()
	if nty.File() != vf.Pgs.File() {
		l10nAccessor = vf.EnumFieldL10nAccessor()
	}

	values := make([]string, len(s))
	for i, v := range s {
		ev, err := vf.parseEnumValue(v)
		if err != nil {
			return "", err
		}
		values[i] = l10nAccessor.Dot(vf.Translator().ResourceId(ev)).String()
	}
	return dart.Qualifier("<String>[" + strings.Join(values, ", ") + "]"), nil
}
