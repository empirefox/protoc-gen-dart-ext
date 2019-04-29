package arb

import (
	"strings"
	"text/template"
)

var Funcs = template.FuncMap{
	"arbDartParams": dartParams,
}

func dartParams(holders ArbPlaceholders, withType bool) string {
	ps := make([]string, len(holders))
	for i, lp := range holders.LangPrams("dart") {
		if withType {
			ps[i] = lp.Info + " " + lp.Name
		} else {
			ps[i] = lp.Name
		}
	}
	return strings.Join(ps, ",")
}
