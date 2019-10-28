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
	ClassName     string
	Arb           *arb.Arb
	RefManager    *PackageArbRefManager
	ImportManager *dart.ImportManager
}

func (l *L10n) As(i *L10n) string {
	// import 'lib.l10n.dart' as `$i0`
	return l.ImportManager.GetAs(i.RootFilePath)
}

func (l *L10n) ImportAsInstanceType(i *L10n) string {
	// `$i0.LibLocalization`
	return l.As(i) + "." + i.ClassName
}

func (l *L10n) ImportAsInstanceName(i *L10n) string {
	// $i0.LibLocalization `_$i0_0`
	return l.ImportManager.ClassInstance(i.RootFilePath, i.ClassName)
}
