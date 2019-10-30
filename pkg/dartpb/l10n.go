package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ArbRef struct {
	From *FieldArb
	To   *MsgOrEnumArb
}

type PackageArbRefManager struct {
	Imports []*ArbRef
	Exports []*ArbRef
	Direct  []*ArbRef
}

type L10n struct {
	RootFilePath  string
	ClassName     dart.Qualifier
	Arb           *arb.Arb
	RefManager    *PackageArbRefManager
	ImportManager *dart.ImportManager
}

func (l *L10n) As(i *L10n) (dart.Qualifier, error) {
	// import 'lib.l10n.dart' as `$i0`
	resolved, err := ResolveImport(l.RootFilePath, i.RootFilePath)
	if err != nil {
		return "", err
	}
	return l.ImportManager.GetAs(resolved), nil
}

func (l *L10n) ImportAsInstanceType(i *L10n) (dart.Qualifier, error) {
	// `$i0.LibLocalization`
	as, err := l.As(i)
	if err != nil {
		return "", err
	}
	return as.Append(i.ClassName), nil
}

func (l *L10n) ImportAsInstanceName(i *L10n) (dart.Qualifier, error) {
	// $i0.LibLocalization `_$i0_0`
	resolved, err := ResolveImport(l.RootFilePath, i.RootFilePath)
	if err != nil {
		return "", err
	}
	return l.ImportManager.ClassInstance(resolved, i.ClassName), nil
}
