package exports

import (
	"sort"
)

var _ sort.Interface = new(Entity)

func (s *Entity) Len() int {
	return len(s.Fields)
}
func (s *Entity) Less(i, j int) bool {
	return s.Fields[i].Name < s.Fields[j].Name
}
func (s *Entity) Swap(i, j int) {
	s.Fields[i], s.Fields[j] = s.Fields[j], s.Fields[i]
}

var _ sort.Interface = new(Package)

func (s *Package) Len() int {
	return len(s.Entities)
}
func (s *Package) Less(i, j int) bool {
	return s.Entities[i].Name < s.Entities[j].Name
}
func (s *Package) Swap(i, j int) {
	s.Entities[i], s.Entities[j] = s.Entities[j], s.Entities[i]
}
