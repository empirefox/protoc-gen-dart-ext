package arb

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
)

func (a *Arb) ParseDartParams(mgr *dart.ImportManager) error {
	for _, r := range a.Resources {
		attr := r.Attr()
		ps, err := NewDartParams(mgr, attr.Placeholders)
		if err != nil {
			return err
		}
		attr.DartParams = ps
	}
	return nil
}

func (attr *ArbAttributes) getDartParamByName(varname string) (*DartParam, error) {
	dp := attr.DartParams.GetByName(varname)
	if dp == nil {
		return nil, fmt.Errorf("varname=%s not found in placeholders", varname)
	}
	return dp, nil
}

func (attr *ArbAttributes) DartParamImport(varname string) (string, error) {
	dp, err := attr.getDartParamByName(varname)
	if err != nil {
		return "", err
	}
	return dp.Raw.Import, nil
}

func (attr *ArbAttributes) DartPlaceholderReplace(varname string) (string, error) {
	dp, err := attr.getDartParamByName(varname)
	if err != nil {
		return "", err
	}
	return dp.Replace, nil
}

func (attr *ArbAttributes) DartSelectCaseCond(varname, key string) (string, error) {
	dp, err := attr.getDartParamByName(varname)
	if err != nil {
		return "", err
	}

	switch dp.Type {
	case "String":
		return dart.RawString(key), nil
	case "num", "double", "int":
		return key, nil
	case "":
		return "", fmt.Errorf("select=%s, case=%s, type is required")
	}

	return dp.Type + "." + key, nil
}

type DartParam struct {
	Raw     *LangParam
	IsEmpty bool
	Type    string
	Replace string
}

type DartParams struct {
	All             []*DartParam
	ByName          map[string]*DartParam
	JoinWithType    string
	JoinWithoutType string
}

func NewDartParams(mgr *dart.ImportManager, holders ArbPlaceholders) (*DartParams, error) {
	all := make([]*DartParam, len(holders))
	for i, lp := range holders.LangParams("dart") {
		if lp.Type == "" && lp.Replace == "" && lp.Import == "" {
			all[i] = &DartParam{
				Raw:     lp,
				IsEmpty: true,
				Replace: lp.Name,
			}
			continue
		}

		// Type
		if strings.Contains(lp.Type, "&") {
			if !checkValidTypeReg.MatchString(lp.Type) {
				return nil, fmt.Errorf("Invalid Type: %#v", *lp)
			}
		}
		typ, err := replaceAs(mgr, lp.Type, lp)
		if err != nil {
			return nil, err
		}

		// Replace
		replace := lp.Replace
		if strings.Contains(replace, "&&&&") {
			return nil, fmt.Errorf("Invalid Replace: %#v", *lp)
		}

		// replace &&&
		if strings.Contains(replace, "&&&.") {
			if lp.Import == "" {
				return nil, fmt.Errorf("Import is required: %#v", *lp)
			}
			classes := searchClassNames(replace)
			for _, class := range classes {
				instance, err := mgr.ResolveClassInstance(lp.Import, dart.Qualifier(class))
				if err != nil {
					return nil, err
				}
				replace = strings.ReplaceAll(replace, "&&&."+class, instance.String())
			}
		}
		if strings.Contains(replace, "&&&") {
			return nil, fmt.Errorf("Invalid &&&: %#v", *lp)
		}

		// replace &&
		replace, err = replaceAs(mgr, replace, lp)
		if err != nil {
			return nil, err
		}

		// replace &
		if strings.Contains(replace, "&") {
			replace = strings.ReplaceAll(replace, "&", lp.Name)
		}

		all[i] = &DartParam{
			Raw:     lp,
			Type:    typ,
			Replace: replace,
		}
	}

	byName := make(map[string]*DartParam, len(all))
	for _, p := range all {
		byName[p.Raw.Name] = p
	}

	params := DartParams{
		All:    all,
		ByName: byName,
	}
	params.JoinWithType, params.JoinWithoutType = joinFuncParams(all)
	return &params, nil
}

func (ps *DartParams) GetByName(name string) *DartParam { return ps.ByName[name] }

var (
	searchClassNamesReg  = regexp.MustCompile(`&&&\.([_$0-9A-Za-z]+)\.[_$0-9A-Za-z]+.`)
	searchDirectNamesReg = regexp.MustCompile(`&&\.([&_$0-9A-Za-z]+)`)
	checkValidTypeReg    = regexp.MustCompile(`^&&(\.[_$0-9A-Za-z^&]+)?$`)
)

func searchClassNames(s string) []string {
	match := searchClassNamesReg.FindAllStringSubmatch(s, -1)
	result := make([]string, len(match))
	for i, found := range match {
		result[i] = found[1]
	}
	return result
}

func searchDirectNames(s string) []string {
	match := searchDirectNamesReg.FindAllStringSubmatch(s, -1)
	result := make([]string, len(match))
	for i, found := range match {
		result[i] = found[1]
	}
	return result
}

func replaceAs(m *dart.ImportManager, s string, lp *LangParam) (string, error) {
	if strings.Contains(s, "&&.") {
		if lp.Import == "" {
			return "", fmt.Errorf("Import is required: %#v", *lp)
		}

		f, err := m.ImportRoot(lp.Import)
		if err != nil {
			return "", err
		}

		names := searchDirectNames(s)
		for _, name := range names {
			if name == "&" {
				name = lp.Name
			} else if strings.Contains(name, "&") {
				return "", fmt.Errorf("Import only allow &&.&(.name)?: %#v", *lp)
			}
			f.AddShow(dart.Qualifier(name))
		}
		s = strings.ReplaceAll(s, "&&", f.As.String())
	}

	if strings.Contains(s, "&&") {
		return "", fmt.Errorf("Invalid &&: %#v", *lp)
	}
	return s, nil
}

func joinFuncParams(ps []*DartParam) (with string, without string) {
	psWith := make([]string, 0, len(ps))
	psWithout := make([]string, 0, len(ps))
	for _, p := range ps {
		if p.Type != "" || p.IsEmpty {
			psWith = append(psWith, p.Type+" "+p.Raw.Name)
			psWithout = append(psWithout, p.Raw.Name)
		}
	}
	return strings.Join(psWith, ","), strings.Join(psWithout, ",")
}
