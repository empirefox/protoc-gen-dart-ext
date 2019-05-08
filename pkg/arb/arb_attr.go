package arb

import (
	"encoding/json"
	"fmt"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

//easyjson:json
type ArbAttributes struct {
	Type         string          `json:"type,omitempty"`
	Context      string          `json:"context,omitempty"`
	Description  string          `json:"description,omitempty"`
	Placeholders ArbPlaceholders `json:"placeholders,omitempty"`

	// MaybeSameWith gives the other resource id. If the id is exist,
	// return the id when the self value is not set or equals the id's.
	MaybeSameWith string `json:"x-maybe_same_with,omitempty"`

	Export string `json:"x-export,omitempty"`
}

type ArbPlaceholders []*ArbPlaceholder

func (s ArbPlaceholders) LangParam(lang, param string) *LangParam {
	for _, holder := range s {
		if holder.LangInfos != nil && holder.Name == param {
			if li := holder.LangInfos.GetLang(lang); li != nil {
				return &LangParam{Name: param, Info: li.Info, Import: li.Import}
			}
			return nil
		}
	}
	return nil
}

func (s ArbPlaceholders) LangParams(lang string) []*LangParam {
	ps := make([]*LangParam, len(s))
	for i, holder := range s {
		if li := holder.LangInfos.GetLang(lang); li != nil {
			ps[i] = &LangParam{Name: holder.Name, Info: li.Info, Import: li.Import}
		} else {
			ps[i] = &LangParam{Name: holder.Name}
		}
	}
	return ps
}

func (s ArbPlaceholders) MarshalJSON() ([]byte, error) {
	m := linkedhashmap.New()
	for _, holder := range s {
		m.Put(holder.Name, holder)
	}
	return m.ToJSON()
}
func (s *ArbPlaceholders) UnmarshalJSON(data []byte) error {
	// need this to preserve insert-order
	m := linkedhashmap.New()
	err := m.FromJSON(data)
	if err != nil {
		return err
	}

	var temp map[string]*ArbPlaceholder
	err = json.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	*s = make(ArbPlaceholders, len(temp))
	for i, k := range m.Keys() {
		key, ok := k.(string)
		if !ok {
			return fmt.Errorf("required string key, but got: %#v", k)
		}
		holder := temp[key]
		holder.Name = key
		(*s)[i] = holder
	}
	return nil
}

//easyjson:json
type ArbPlaceholder struct {
	Name        string       `json:"-"`
	LangInfos   ArbLangInfos `json:"x-lang-info,omitempty"`
	Description string       `json:"description,omitempty"`
	Example     interface{}  `json:"example,omitempty"`
}

//easyjson:json
type ArbLangInfos []*ArbLangInfo

func (s ArbLangInfos) GetLang(lang string) *ArbLangInfo {
	for _, li := range s {
		if li.Lang == lang {
			return li
		}
	}
	return nil
}

//easyjson:json
type ArbLangInfo struct {
	Lang   string
	Info   string
	Import string
}

type LangParam struct {
	Name   string
	Info   string
	Import string
}
