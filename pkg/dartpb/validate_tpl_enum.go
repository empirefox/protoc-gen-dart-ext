package dartpb

import (
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

func (vf *ValidateField) EnumType() (dart.Qualifier, error) {
	return vf.Validator().FullPgsPbEnum(vf.Pgs.Type().Enum())
}

func (vf *ValidateField) EnumLiteralValue(v int32) (dart.Qualifier, error) {
	value, err := vf.parseEnumValue(v)
	if err != nil {
		return "", err
	}

	targetFullEnumClass, err := vf.Validator().FullPbEnum(value.Enum)
	if err != nil {
		return "", err
	}

	return targetFullEnumClass.Append(value.DartName), nil
}

func (vf *ValidateField) EnumFullL10nClass() (dart.Qualifier, error) {
	nty := vf.Package.Group.PgsToEnum[vf.Pgs.Type().Enum()]
	if nty.Package == vf.Package {
		return "", nil
	}

	resolvedL10nImport, err := ResolveImport(vf.Validator().RootFilePath, nty.L10n().RootFilePath)
	if err != nil {
		return "", err
	}
	return vf.Validator().ImportManager.GetAs(resolvedL10nImport).Append(nty.L10n().ClassName), nil
}

func (vf *ValidateField) EnumL10nValues(s []int32) (dart.Qualifier, error) {
	nty := vf.Package.Group.PgsToEnum[vf.Pgs.Type().Enum()]
	l10nAccessor := vf.L10nAccessor()
	if nty.Package != vf.Package {
		l10nAccessor = vf.EnumFieldL10nAccessor()
	}

	values := make([]string, len(s))
	for i, v := range s {
		ev, err := vf.parseEnumValue(v)
		if err != nil {
			return "", err
		}
		values[i] = l10nAccessor.Append(ev.Arb.ResourceId()).String()
	}
	return dart.Qualifier("[" + strings.Join(values, ", ") + "]"), nil
}
