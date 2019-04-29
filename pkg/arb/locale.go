package arb

import (
	"sort"

	"golang.org/x/text/language"
)

type SupportedLocale struct {
	Lang language.Base

	// Regions contain regions or scripts
	Regions  []string
	Fallback bool
}

func SupportedLocales(as []*Arb) []SupportedLocale {
	locales := make(localeMap, len(as))
	for _, a := range as {
		locales.add(a.Locale)
	}
	return locales.list()
}

var noRegion language.Region

type localeMap map[language.Base]regionMap

func (m localeMap) add(tag language.Tag) {
	if tag.IsRoot() {
		return
	}
	lang, script, region := tag.Raw()
	rm, ok := m[lang]
	if !ok {
		rm = make(regionMap, 24)
		m[lang] = rm
	}
	if region == noRegion {
		rm[script.String()] = struct{}{}
	} else {
		rm[region.String()] = struct{}{}
	}
}

func (m localeMap) list() []SupportedLocale {
	locales := make([]SupportedLocale, 0, len(m))
	for lang, regions := range m {
		list, fallback := regions.list()
		locales = append(locales, SupportedLocale{
			Lang:     lang,
			Regions:  list,
			Fallback: fallback,
		})
	}
	sort.Slice(locales, func(i, j int) bool {
		return locales[i].Lang.String() < locales[j].Lang.String()
	})
	return locales
}

type regionMap map[string]struct{}

func (m regionMap) list() (regions []string, fallback bool) {
	regions = make([]string, 0, len(m))
	for region := range m {
		if region == "Zzzz" {
			fallback = true
		} else {
			regions = append(regions, region)
		}
	}
	sort.Strings(regions)
	return
}
