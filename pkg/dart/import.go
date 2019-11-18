package dart

import (
	"fmt"
	"path/filepath"
	"sort"
	"strings"

	pgs "github.com/lyft/protoc-gen-star"
)

type ImportFileClass struct {
	File       *ImportFile
	SimpleName Qualifier
	Instance   Qualifier
}

func (c *ImportFileClass) FullName() Qualifier {
	return c.File.As.Dot(c.SimpleName)
}

//go:generate pie ImportFileShow.Keys
type ImportFileShow map[string]emptyStruct
type emptyStruct struct{}

type ImportFile struct {
	Manager *ImportManager
	Name    string
	As      Qualifier

	UseShow         bool
	IgnoreEmptyShow bool
	show            ImportFileShow

	byName map[Qualifier]*ImportFileClass
	list   []*ImportFileClass
}

func (ifile *ImportFile) AsDot(simpleClass Qualifier) Qualifier {
	return ifile.AddShow(simpleClass).As.Dot(simpleClass)
}

func (ifile *ImportFile) RenderImport() string {
	if ifile.As == "" {
		return ""
	}

	show := ifile.Show()
	if ifile.IgnoreEmptyShow && show == "" {
		return ""
	}

	var b strings.Builder
	b.WriteString(`import '`)
	b.WriteString(ifile.Name)
	b.WriteByte('\'')
	if ifile.As != "" {
		b.WriteString(` as `)
		b.WriteString(ifile.As.String())
	}
	if ifile.UseShow && show != "" {
		b.WriteString(` show `)
		b.WriteString(show)
	}
	b.WriteByte(';')

	return b.String()
}

func (ifile *ImportFile) Show() string {
	ss := ifile.show.Keys()
	sort.Strings(ss)
	return strings.Join(ss, ", ")
}

func (ifile *ImportFile) DisableIgnoreEmptyShow() *ImportFile {
	ifile.IgnoreEmptyShow = false
	return ifile
}

func (ifile *ImportFile) DisableShow() *ImportFile {
	ifile.UseShow = false
	return ifile
}

func (ifile *ImportFile) AddShow(simpleClass Qualifier) *ImportFile {
	ifile.show[simpleClass.String()] = emptyStruct{}
	return ifile
}

func (ifile *ImportFile) Classes() []*ImportFileClass { return ifile.list }

func (ifile *ImportFile) ClassInstance(simpleClass Qualifier) Qualifier {
	return ifile.getClass(simpleClass).Instance
}

func (ifile *ImportFile) getClass(simpleClass Qualifier) *ImportFileClass {
	c, ok := ifile.byName[simpleClass]
	if !ok {
		name := fmt.Sprintf("%s%s_%d", ifile.Manager.depPrefix, ifile.As, len(ifile.byName))
		c = &ImportFileClass{
			File:       ifile,
			SimpleName: simpleClass,
			Instance:   Qualifier(name),
		}
		ifile.byName[simpleClass] = c
		ifile.list = append(ifile.list, c)
		ifile.AddShow(simpleClass)
	}
	return c
}

type ImportManager struct {
	d *Dart

	rootFilePath string

	// $i
	prefix string

	// add this to prefix to create new ifile
	depPrefix string

	byName map[string]*ImportFile
	files  []*ImportFile
}

func NewDefaultImportManager(d *Dart, rootFilePath string) *ImportManager {
	return NewImportManager(d, rootFilePath, "$", "_")
}

func NewImportManager(d *Dart,
	rootFilePath, prefix, depPrefix string,
	predefine ...string) *ImportManager {
	im := &ImportManager{
		d:            d,
		rootFilePath: rootFilePath,
		prefix:       prefix,
		depPrefix:    depPrefix,
		byName:       make(map[string]*ImportFile, 16),
	}
	im.getFile("").As = ""
	for _, f := range predefine {
		im.getFile(f)
	}
	return im
}

func (m *ImportManager) RootFilePath() string { return m.rootFilePath }

func ResolveImport(source, target string) (string, error) {
	if source == target {
		return "", nil
	}
	if strings.HasPrefix(target, "dart:") || strings.HasPrefix(target, "package:") {
		return target, nil
	}
	return filepath.Rel(filepath.Dir(source), target)
}

func (m *ImportManager) Resolve(target string) (string, error) {
	return ResolveImport(m.rootFilePath, target)
}

func (m *ImportManager) RootPathFileDot(target string, simpleName Qualifier) (Qualifier, error) {
	f, err := m.RootPathFile(target)
	if err != nil {
		return "", err
	}
	return f.AsDot(simpleName), nil
}

func (m *ImportManager) RootPathFile(target string) (*ImportFile, error) {
	resolved, err := m.Resolve(target)
	if err != nil {
		return nil, err
	}
	return m.Import(resolved), nil
}

type FileType string

var (
	PbFile         FileType = ".pb.dart"
	L10nFile       FileType = ".l10n.dart"
	ValidateFile   FileType = ".validate.dart"
	FormInputsFile FileType = ".inputs.dart"
)

func (ft FileType) RootFilePath(nty pgs.Entity) string {
	return nty.File().InputPath().SetExt(string(ft)).String()
}

func (m *ImportManager) PbFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(PbFile.RootFilePath(nty), simpleName)
}
func (m *ImportManager) L10nFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(L10nFile.RootFilePath(nty), simpleName)
}
func (m *ImportManager) ValidateFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(ValidateFile.RootFilePath(nty), simpleName)
}
func (m *ImportManager) FormInputsFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(FormInputsFile.RootFilePath(nty), simpleName)
}

func (m *ImportManager) L10nAsInstance(i pgs.Entity) (Qualifier, error) {
	// $i0.LibLocalizations `_$i0_0`
	resolved, err := m.Resolve(L10nFile.RootFilePath(i))
	if err != nil {
		return "", err
	}
	return m.ClassInstance(resolved, m.d.FileNames(i.File()).L10nName()), nil
}

func (m *ImportManager) L10nAsInstanceType(i pgs.Entity) (Qualifier, error) {
	// `$i0.LibLocalizations`
	return m.L10nFileDot(i, m.d.FileNames(i.File()).L10nName())
}

func (m *ImportManager) InstanceClasses() []*ImportFileClass {
	all := make([]*ImportFileClass, 0, 64)
	for _, ifile := range m.files {
		all = append(all, ifile.list...)
	}
	return all
}

func (m *ImportManager) Files() []*ImportFile { return m.files }

func (m *ImportManager) GetAs(fileName string) Qualifier {
	return m.getFile(fileName).As
}

func (m *ImportManager) ResolveClassInstance(rootPath string, simpleClass Qualifier) (Qualifier, error) {
	resolved, err := m.Resolve(rootPath)
	if err != nil {
		return "", err
	}
	return m.ClassInstance(resolved, simpleClass), nil
}

func (m *ImportManager) ClassInstance(fileName string, simpleClass Qualifier) Qualifier {
	return m.getFile(fileName).ClassInstance(simpleClass)
}

func (m *ImportManager) getAs(seq int) string {
	return fmt.Sprintf("%s%d", m.prefix, seq)
}

func (m *ImportManager) ImportRoot(rootPath string) (*ImportFile, error) {
	resolved, err := m.Resolve(rootPath)
	if err != nil {
		return nil, err
	}
	return m.Import(resolved), nil
}

func (m *ImportManager) Import(fileName string) *ImportFile {
	return m.getFile(fileName)
}

func (m *ImportManager) getFile(fileName string) *ImportFile {
	ifile, ok := m.byName[fileName]
	if !ok {
		seq := len(m.byName) - 1
		ifile = &ImportFile{
			Manager:         m,
			Name:            fileName,
			As:              Qualifier(m.getAs(seq)),
			UseShow:         true,
			IgnoreEmptyShow: true,
			show:            make(ImportFileShow),
			byName:          make(map[Qualifier]*ImportFileClass),
		}
		m.byName[fileName] = ifile
		m.files = append(m.files, ifile)
	}
	return ifile
}

func (m *ImportManager) RenderImports() string {
	ss := make([]string, 0, len(m.files))
	for _, f := range m.files {
		s := f.RenderImport()
		if s != "" {
			ss = append(ss, s)
		}
	}
	return strings.Join(ss, "\n")
}
