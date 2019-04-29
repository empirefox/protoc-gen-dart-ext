package arb

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/emirpasic/gods/maps/linkedhashmap"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/util"
	"github.com/mailru/easyjson"
	"github.com/relvacode/iso8601"
	"golang.org/x/text/language"
)

//easyjson:json
type Arb struct {
	LastModified iso8601.Time `json:"@@last_modified,omitempty"`
	Context      string       `json:"@@context,omitempty"`
	Locale       language.Tag `json:"@@locale,omitempty"`
	Author       string       `json:"@@author,omitempty"`
	Entity       string       `json:"@@x-entity,omitempty"`

	// Imports resoures import from other resources
	// which export unique id of entity.
	Imports []*ArbImport `json:"@@x-imports,omitempty"`

	Nils []string `json:"@@x-nils,omitempty"`

	Resources []*ArbResource `json:"-"`

	resourceMap map[string]*ArbResource   `json:"-"`
	attrs       map[string]*ArbAttributes `json:"-"`
}

func (a *Arb) LocaleLang() (lang language.Base) {
	lang, _ = a.Locale.Base()
	return
}
func (a *Arb) LocaleRegion() (region language.Region) {
	region, _ = a.Locale.Region()
	return
}

func (a *Arb) SetAttrs(en *Arb) { a.attrs = en.Attrs() }

func (a *Arb) Attrs() map[string]*ArbAttributes {
	if a.attrs == nil {
		a.attrs = make(map[string]*ArbAttributes, len(a.Resources))
		for _, r := range a.Resources {
			a.attrs[r.Id] = r.attr
		}
	}
	return a.attrs
}

func (a *Arb) resources() map[string]*ArbResource {
	if a.resourceMap == nil {
		a.resourceMap = make(map[string]*ArbResource, len(a.Resources))
		for _, r := range a.Resources {
			a.resourceMap[r.Id] = r
		}
	}
	return a.resourceMap
}

func (a *Arb) MarshalJSON() ([]byte, error) {
	m := linkedhashmap.New()
	a.addToJsonMap(m)
	for _, r := range a.Resources {
		r.addToJsonMap(m)
	}
	return m.ToJSON()
}
func (a *Arb) UnmarshalJSON(b []byte) error {
	err := easyjson.Unmarshal(b, interface{}(a).(easyjson.Unmarshaler))
	if err != nil {
		return err
	}

	// need this to preserve insert-order
	m := linkedhashmap.New()
	err = m.FromJSON(b)
	if err != nil {
		return err
	}

	var temp map[string]json.RawMessage
	err = json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	if a.attrs == nil {
		a.attrs = make(map[string]*ArbAttributes, len(temp))
	}
	a.Resources = make([]*ArbResource, 0, len(temp))
	for _, ik := range m.Keys() {
		k, ok := ik.(string)
		if !ok {
			return fmt.Errorf("required string key, but got: %#v", k)
		}

		v := temp[k]
		if strings.HasPrefix(k, "@@") {
			continue
		}
		if strings.HasPrefix(k, "@") {
			var attr ArbAttributes
			err = json.Unmarshal(v, &attr)
			if err != nil {
				return err
			}
			a.attrs[k[1:]] = &attr
			continue
		}
		r := ArbResource{Id: k, Arb: a}
		err = json.Unmarshal(v, &r.Value)
		if err != nil {
			return err
		}
		a.Resources = append(a.Resources, &r)
	}
	for _, r := range a.Resources {
		r.attr = a.attrs[r.Id]
	}
	return nil
}

func (a *Arb) addToJsonMap(m *linkedhashmap.Map) {
	if !a.LastModified.IsZero() {
		m.Put("@@last_modified", a.LastModified.UTC().Format(time.RFC3339))
	}
	if a.Context != "" {
		m.Put("@@context", a.Context)
	}
	if !a.Locale.IsRoot() {
		m.Put("@@locale", a.Locale)
	}
	if a.Author != "" {
		m.Put("@@author", a.Author)
	}

	if a.Entity != "" {
		m.Put("@@x-entity", a.Entity)
	}
}

type ArbResource struct {
	Arb   *Arb
	Id    string
	Value string
	attr  *ArbAttributes
}

func NewResource(a *Arb, id, value string, attr *ArbAttributes) *ArbResource {
	return &ArbResource{
		Arb:   a,
		Id:    id,
		Value: value,
		attr:  attr,
	}
}

func (ar *ArbResource) SameWith() (string, error) {
	other := ar.Attr().MaybeSameWith
	if other == "" {
		return "", nil
	}
	r, ok := ar.Arb.resources()[other]
	if !ok {
		return "", fmt.Errorf("other resource not found: x-maybe_same_with=%s", other)
	}
	if ar.Value == "" || r.Value == ar.Value {
		return other, nil
	}
	return "", nil
}

func (ar *ArbResource) Attr() *ArbAttributes {
	attr, ok := ar.Arb.attrs[ar.Id]
	if ok {
		return attr
	}
	return ar.attr
}

func (ar *ArbResource) addToJsonMap(m *linkedhashmap.Map) {
	m.Put(ar.Id, ar.Value)
	if ar.attr != nil {
		m.Put("@"+ar.Id, ar.attr)
	}
}

func ShallowMerge(from []*Arb, to *Arb) {
	total := len(to.Imports)
	for _, a := range from {
		total += len(a.Imports)
	}
	newImp := make([]*ArbImport, 0, total)
	newImp = append(newImp, to.Imports...)

	total = len(to.Nils)
	for _, a := range from {
		total += len(a.Nils)
	}
	newNils := make([]string, 0, total)
	newNils = append(newNils, to.Nils...)

	total = len(to.Resources)
	for _, a := range from {
		total += len(a.Resources)
	}
	newRs := make([]*ArbResource, 0, total)
	newRs = append(newRs, to.Resources...)

	for _, a := range from {
		if a.LastModified.After(to.LastModified.Time) {
			to.LastModified = a.LastModified
		}
		// TODO merge authors
		for _, r1 := range a.Imports {
			if r1.Id == "" {
				continue
			}
			r2 := *r1
			r2.Id = util.JoinEntityId(a.Entity, r2.Id)
			newImp = append(newImp, &r2)
		}
		for _, n := range a.Nils {
			if n == "" {
				continue
			}
			newNils = append(newNils, util.JoinEntityId(a.Entity, n))
		}
		for _, r1 := range a.Resources {
			if r1.Id == "" {
				continue
			}
			r2 := *r1
			if !isEntity(from, r2.Id) && r2.Id != to.Entity {
				// ignore package entity name
				r2.Id = util.JoinEntityId(a.Entity, r2.Id)
			}
			attr := *r1.Attr()
			if attr.MaybeSameWith != "" &&
				!isEntity(from, attr.MaybeSameWith) &&
				attr.MaybeSameWith != to.Entity {
				attr.MaybeSameWith = util.JoinEntityId(a.Entity, attr.MaybeSameWith)
			}
			r2.attr = &attr
			r2.Arb = to
			newRs = append(newRs, &r2)
		}
	}
	to.Imports = newImp
	to.Nils = newNils
	to.Resources = newRs
	to.resourceMap = nil
	to.attrs = nil
}

func isEntity(as []*Arb, id string) bool {
	for _, a := range as {
		if a.Entity == id {
			return true
		}
	}
	return false
}
