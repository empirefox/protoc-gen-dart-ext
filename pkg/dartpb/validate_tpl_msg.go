package dartpb

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

type ValidateMessage struct {
	*Message
	Disabled bool
}

func (nty *ValidateMessage) ConvertLib() *dart.ImportFile { return nty.Validators().ConvertLib }
func (nty *ValidateMessage) FoundationFile() *dart.ImportFile { return nty.Validators().FoundationFile }
func (nty *ValidateMessage) MaterialFile() *dart.ImportFile   { return nty.Validators().MaterialFile }
func (nty *ValidateMessage) FixnumFile() *dart.ImportFile     { return nty.Validators().FixnumFile }
func (nty *ValidateMessage) EmailValidatorFile() *dart.ImportFile {
	return nty.Validators().EmailValidatorFile
}
func (nty *ValidateMessage) PgdeFile() *dart.ImportFile { return nty.Validators().PgdeFile }
func (nty *ValidateMessage) PbFile() *dart.ImportFile   { return nty.Validators().PbFile }
func (nty *ValidateMessage) L10nFile() *dart.ImportFile { return nty.Validators().L10nFile }

func (nty *ValidateMessage) FullPbClass() (dart.Qualifier, error) {
	return nty.Validators().FullPbClass(nty.Pgs)
}

func (nty *ValidateMessage) FullL10nClass() (dart.Qualifier, error) {
	return nty.Validators().ImportManager.L10nAsInstanceType(nty.Pgs)
}
