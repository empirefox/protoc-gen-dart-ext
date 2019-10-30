package dart

import (
	"fmt"
	"sort"
	"strings"
)

type ImportFileClass struct {
	File       *ImportFile
	SimpleName Qualifier
	Instance   Qualifier
}

func (c *ImportFileClass) FullName() Qualifier {
	return c.File.As.Append(c.SimpleName)
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

func (ifile *ImportFile) EnableIgnoreEmptyShow() *ImportFile {
	ifile.IgnoreEmptyShow = true
	return ifile
}

func (ifile *ImportFile) EnableShow() *ImportFile {
	ifile.UseShow = true
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
	// $i
	prefix string
	// add this to prefix to create new ifile
	depPrefix string
	byName    map[string]*ImportFile
	files     []*ImportFile
}

func NewDefaultImportManager() *ImportManager {
	return NewImportManager("$", "_")
}

func NewImportManager(prefix, depPrefix string, predefine ...string) *ImportManager {
	im := &ImportManager{
		prefix:    prefix,
		depPrefix: depPrefix,
		byName:    make(map[string]*ImportFile, 16),
	}
	im.getFile("").As = ""
	for _, f := range predefine {
		im.getFile(f)
	}
	return im
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

func (m *ImportManager) ClassInstance(fileName string, simpleClass Qualifier) Qualifier {
	return m.getFile(fileName).ClassInstance(simpleClass)
}

func (m *ImportManager) getAs(seq int) string {
	return fmt.Sprintf("%s%d", m.prefix, seq)
}

func (m *ImportManager) GetFile(fileName string) *ImportFile {
	return m.getFile(fileName)
}

func (m *ImportManager) getFile(fileName string) *ImportFile {
	ifile, ok := m.byName[fileName]
	if !ok {
		seq := len(m.byName) - 1
		ifile := &ImportFile{
			Manager: m,
			Name:    fileName,
			As:      Qualifier(m.getAs(seq)),
			show:    make(ImportFileShow),
			byName:  make(map[Qualifier]*ImportFileClass),
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
