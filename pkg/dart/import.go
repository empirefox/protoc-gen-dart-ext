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

func (ifile *ImportFile) AsDotString(simpleClass string) Qualifier {
	return ifile.AsDot(Qualifier(simpleClass))
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

type ImportManagerCommonFiles interface {
	ImportManager() *ImportManager

	Animation() *Animation
	Foundation() *Foundation
	Material() *Material

	CollectionLib() *ImportFile
	ConvertLib() *ImportFile
	WrappersFile() *ImportFile
	PgdeFile() *ImportFile
	Int64Type() Qualifier
	AnyType() Qualifier
	DurationType() Qualifier
	TimestampType() Qualifier
}

var _ ImportManagerCommonFiles = new(ImportManager)

type ImportManager struct {
	d *Dart

	animation  *Animation
	foundation *Foundation
	material   *Material

	collectionLib *ImportFile
	convertLib    *ImportFile
	wrappersFile  *ImportFile
	pgdeFile      *ImportFile

	int64Type       Qualifier
	anyType         Qualifier
	durationType    Qualifier
	timestampType   Qualifier
	durationBeTime  Qualifier
	timestampBeTime Qualifier

	// The owner file' path from protoc out path. Cannot be any absolute path.
	RootFilePath string

	// $i
	prefix string

	// add this to prefix to create new ifile
	depPrefix string

	byName map[string]*ImportFile
	files  []*ImportFile
}

func NewDefaultImportManager(d *Dart, rootFilePath string) (*ImportManager, error) {
	return NewImportManager(d, rootFilePath, "$", "_")
}

func NewImportManager(d *Dart,
	rootFilePath, prefix, depPrefix string,
	predefine ...string) (*ImportManager, error) {
	if _, err := filepath.Rel(rootFilePath, "google"); err != nil {
		return nil, err
	}

	im := &ImportManager{
		d:            d,
		RootFilePath: filepath.Clean(rootFilePath),
		prefix:       prefix,
		depPrefix:    depPrefix,
		byName:       make(map[string]*ImportFile, 16),
	}
	im.getFile("").As = ""
	for _, f := range predefine {
		im.getFile(f)
	}
	im.animation = NewAnimation(im)
	im.foundation = NewFoundation(im)
	im.material = NewMaterial(im)
	return im, nil
}

func (im *ImportManager) ImportManager() *ImportManager { return im }

func (im *ImportManager) Animation() *Animation   { return im.animation }
func (im *ImportManager) Foundation() *Foundation { return im.foundation }
func (im *ImportManager) Material() *Material     { return im.material }

func (im *ImportManager) CollectionLib() *ImportFile {
	if im.collectionLib == nil {
		im.collectionLib = im.Import("dart:collection")
	}
	return im.collectionLib
}
func (im *ImportManager) ConvertLib() *ImportFile {
	if im.convertLib == nil {
		im.convertLib = im.Import("dart:convert")
	}
	return im.convertLib
}
func (im *ImportManager) WrappersFile() *ImportFile {
	if im.wrappersFile == nil {
		im.wrappersFile, _ = im.ImportRoot("google/protobuf/wrappers.pb.dart")
	}
	return im.wrappersFile
}
func (im *ImportManager) PgdeFile() *ImportFile {
	if im.pgdeFile == nil {
		im.pgdeFile = im.Import("package:pgde/pgde.dart")
	}
	return im.pgdeFile
}

func (im *ImportManager) Int64Type() Qualifier {
	if im.int64Type == "" {
		im.int64Type = im.Import("package:fixnum/fixnum.dart").AsDot("Int64")
	}
	return im.int64Type
}
func (im *ImportManager) AnyType() Qualifier {
	if im.anyType == "" {
		f, _ := im.ImportRoot("google/protobuf/any.pb.dart")
		im.anyType = f.AsDot("Any")
	}
	return im.anyType
}
func (im *ImportManager) DurationType() Qualifier {
	if im.durationType == "" {
		f, _ := im.ImportRoot("google/protobuf/duration.pb.dart")
		im.durationType = f.AsDot("Duration")
	}
	return im.durationType
}
func (im *ImportManager) TimestampType() Qualifier {
	if im.timestampType == "" {
		f, _ := im.ImportRoot("google/protobuf/timestamp.pb.dart")
		im.timestampType = f.AsDot("Timestamp")
	}
	return im.timestampType
}
func (im *ImportManager) DurationBeTime() Qualifier {
	if im.durationBeTime == "" {
		f, _ := im.ImportRoot("google/protobuf/duration.fmt.dart")
		im.durationBeTime = f.AsDot("beTime")
	}
	return im.durationBeTime
}
func (im *ImportManager) TimestampBeTime() Qualifier {
	if im.timestampBeTime == "" {
		f, _ := im.ImportRoot("google/protobuf/timestamp.fmt.dart")
		im.timestampBeTime = f.AsDot("beTime")
	}
	return im.timestampBeTime
}

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
	return ResolveImport(m.RootFilePath, target)
}

func (m *ImportManager) RootPathFileDot(target string, simpleName Qualifier) (Qualifier, error) {
	f, err := m.ImportRoot(target)
	if err != nil {
		return "", err
	}
	return f.AsDot(simpleName), nil
}

func (im *ImportManager) ImportPbFile(nty pgs.Entity) (f *ImportFile) {
	f, _ = im.ImportRoot(PbFile.RootFilePath(nty.File()))
	return
}
func (im *ImportManager) ImportZeroFile(nty pgs.Entity) (f *ImportFile) {
	f, _ = im.ImportRoot(ZeroFile.RootFilePath(nty.File()))
	return
}
func (im *ImportManager) ImportL10nFile(nty pgs.Entity) (f *ImportFile) {
	f, _ = im.ImportRoot(L10nFile.RootFilePath(nty.File()))
	return
}
func (im *ImportManager) ImportValidateFile(nty pgs.Entity) (f *ImportFile) {
	f, _ = im.ImportRoot(ValidateFile.RootFilePath(nty.File()))
	return
}
func (im *ImportManager) ImportFormInputsFile(nty pgs.Entity) (f *ImportFile) {
	f, _ = im.ImportRoot(FormInputsFile.RootFilePath(nty.File()))
	return
}

type FileType string

var (
	PbFile         FileType = ".pb.dart"
	ZeroFile       FileType = ".zero.dart"
	L10nFile       FileType = ".l10n.dart"
	ValidateFile   FileType = ".validate.dart"
	FormInputsFile FileType = ".inputs.dart"
)

func (ft FileType) RootFilePath(nty pgs.Entity) string {
	return nty.File().InputPath().SetExt(string(ft)).String()
}

type DartFileDotFunc func(nty pgs.Entity, simpleName Qualifier) (Qualifier, error)

func (m *ImportManager) PbFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(PbFile.RootFilePath(nty), simpleName)
}
func (m *ImportManager) ZeroFileDot(nty pgs.Entity, simpleName Qualifier) (Qualifier, error) {
	return m.RootPathFileDot(ZeroFile.RootFilePath(nty), simpleName)
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

func (m *ImportManager) TypeForFieldType(t pgs.FieldType) (Qualifier, error) {
	if t.IsMap() {
		key, err := m.TypeForNonRangeFieldType(t.Key())
		if err != nil {
			return "", err
		}
		value, err := m.TypeForNonRangeFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return Qualifier(fmt.Sprintf("Map<%s, %s>", key, value)), nil
	}

	if t.IsRepeated() {
		value, err := m.TypeForNonRangeFieldType(t.Element())
		if err != nil {
			return "", err
		}
		return Qualifier(fmt.Sprintf("List<%s>", value)), nil
	}

	return m.TypeForNonRangeFieldType(t)
}

type EmbedOrEnumOrCoreFieldType interface {
	IsEmbed() bool
	Embed() pgs.Message
	IsEnum() bool
	Enum() pgs.Enum
	ProtoType() pgs.ProtoType
}

func (m *ImportManager) TypeForNonRangeFieldType(ft EmbedOrEnumOrCoreFieldType) (Qualifier, error) {
	if ft.IsEmbed() {
		return m.FullPbClassOrEnum(ft.Embed())
	}
	if ft.IsEnum() {
		return m.FullPbClassOrEnum(ft.Enum())
	}
	// make sure it will never return dynamic
	return m.TypeForCoreFieldType(ft.ProtoType()), nil
}

func (m *ImportManager) TypeForCoreFieldType(t pgs.ProtoType) Qualifier {
	switch t {
	case pgs.Int32T, pgs.UInt32T, pgs.SInt32, pgs.Fixed32T, pgs.SFixed32:
		return "int"
	case pgs.Int64T, pgs.UInt64T, pgs.SInt64, pgs.Fixed64T, pgs.SFixed64:
		return m.Int64Type()
	case pgs.DoubleT, pgs.FloatT:
		return "double"
	case pgs.BoolT:
		return "bool"
	case pgs.StringT:
		return "String"
	case pgs.BytesT:
		return "List<int>"
	}
	return "dynamic"
}

func (m *ImportManager) FullPbClassOrEnum(nty pgs.Entity) (Qualifier, error) {
	return m.PbFileDot(nty, m.d.NameOf(nty))
}
func (m *ImportManager) FullPbWhichEnum(nty pgs.OneOf) (Qualifier, error) {
	return m.PbFileDot(nty, m.d.OneOfNames(nty).OneofEnumName())
}
