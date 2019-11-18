package arb

import (
	"sort"

	"golang.org/x/text/language"
)

type SupportedLocale struct {
	Lang language.Base

	// Regions contain regions or scripts
	Regions  []string
	Fallback string
}

func SupportedLocales(as []*Arb) []SupportedLocale {
	locales := make(localeMap, len(as))
	for _, a := range as {
		locales.add(a.Locale)
	}
	return locales.list()
}

var noRegion language.Region

type localeMap map[language.Base]baseRegions

func (m localeMap) add(tag language.Tag) {
	if tag.IsRoot() {
		return
	}

	lang, script, region := tag.Raw()
	_, ok := m[lang]
	if !ok {
		m[lang] = make(baseRegions, 0, 24)
	}

	baseTag, _ := language.Compose(lang)
	if region == noRegion {
		m[lang] = append(m[lang], &baseRegion{
			region: script.String(),
			c:      language.Comprehends(tag, baseTag),
		})
	} else {
		m[lang] = append(m[lang], &baseRegion{
			region: region.String(),
			c:      language.Comprehends(tag, baseTag),
		})
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

type baseRegion struct {
	region string
	c      language.Confidence
}

type baseRegions []*baseRegion

func (m baseRegions) list() (regions []string, fallback string) {
	var best *baseRegion

	regions = make([]string, 0, len(m))
	for _, br := range m {
		if best == nil {
			best = br
			continue
		}

		if br.c > best.c {
			best, br = br, best
		}
		regions = append(regions, br.region)
	}
	sort.Strings(regions)

	fallback = best.region
	if fallback == "Zzzz" {
		fallback = ""
	}
	return
}
