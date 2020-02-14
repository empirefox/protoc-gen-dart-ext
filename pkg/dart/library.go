package dart

type Library struct {
	im         *ImportManager
	file       *ImportFile
	importPath string
}

func (lib *Library) ImportPath() string { return lib.importPath }

func (lib *Library) File() *ImportFile {
	if lib.file == nil {
		lib.file = lib.im.Import(lib.importPath)
	}
	return lib.file
}
