package dart

import (
	"fmt"
	"sort"
	"strings"
)

type ImportFileClass struct {
	File       *ImportFile
	SimpleName string
	Instance   string
}

func (c *ImportFileClass) FullName() string {
	return c.File.As + "." + c.SimpleName
}

//go:generate pie ImportFileShow.Keys
type ImportFileShow map[string]emptyStruct
type emptyStruct struct{}

type ImportFile struct {
	Manager  *ImportManager
	Name     string
	As string

	IsShow bool
	show   ImportFileShow

	byName map[string]*ImportFileClass
	list   []*ImportFileClass
}

func (ifile *ImportFile) Show() string {
	ss := ifile.show.Keys()
	sort.Strings(ss)
	return strings.Join(ss, ", ")
}

func (ifile *ImportFile) EnableShow() *ImportFile {
	ifile.IsShow = true
	return ifile
}

func (ifile *ImportFile) AddShow(simpleClass string) *ImportFile {
	ifile.show[simpleClass] = emptyStruct{}
	return ifile
}

func (ifile *ImportFile) Classes() []*ImportFileClass { return ifile.list }

func (ifile *ImportFile) ClassInstance(simpleClass string) string {
	return ifile.getClass(simpleClass).Instance
}

func (ifile *ImportFile) getClass(simpleClass string) *ImportFileClass {
	c, ok := ifile.byName[simpleClass]
	if !ok {
		name := fmt.Sprintf("%s%s_%d", ifile.Manager.depPrefix, ifile.As, len(ifile.byName))
		c = &ImportFileClass{
			File:       ifile,
			SimpleName: simpleClass,
			Instance:   name,
		}
		ifile.byName[simpleClass] = c
		ifile.list = append(ifile.list, c)
		ifile.show[simpleClass] = emptyStruct{}
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
	for _, f := range predefine {
		im.GetAs(f)
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

func (m *ImportManager) GetAs(fileName string) string {
	if fileName == "" {
		return ""
	}
	return m.getFile(fileName).As
}

func (m *ImportManager) ClassInstance(fileName, simpleClass string) string {
	if fileName == "" {
		return ""
	}
	return m.getFile(fileName).ClassInstance(simpleClass)
}

func (m *ImportManager) getAs(seq int) string {
	return fmt.Sprintf("%s%d", m.prefix, seq)
}

func (m *ImportManager) GetFile(fileName string) *ImportFile {
	if fileName == "" {
		return nil
	}
	return m.getFile(fileName)
}

func (m *ImportManager) getFile(fileName string) *ImportFile {
	ifile, ok := m.byName[fileName]
	if !ok {
		seq := len(m.byName)
		ifile := &ImportFile{
			Manager:  m,
			Name:     fileName,
			As: m.getAs(seq),
			show:     make(ImportFileShow),
			byName:   make(map[string]*ImportFileClass),
		}
		m.byName[fileName] = ifile
		m.files = append(m.files, ifile)
	}
	return ifile
}
