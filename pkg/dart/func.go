package dart

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/arb"

	pgs "github.com/lyft/protoc-gen-star"
)

var Funcs = template.FuncMap{
	"dartPgsPkgName":            PackageName,
	"dartPgsFileNameFromEntity": FileName,
	"dartRawStr":                RawString,
	"dartArbParams":             ArbParams,
}

func (d *Dart) Funcs() template.FuncMap {
	return template.FuncMap{
		"dartNameOf": d.NameOf,
	}
}

func PackageName(e pgs.Package) pgs.Name {
	return pgs.Name(pgs.FilePath(e.ProtoName()).BaseName()).UpperCamelCase()
}

func FileName(e pgs.Entity) pgs.Name {
	return pgs.Name(e.File().InputPath().BaseName()).UpperCamelCase()
}

func RawString(s string) string {
	const squot = '\''
	if !strings.ContainsRune(s, squot) {
		return fmt.Sprintf(`r'%s'`, s)
	}

	var output strings.Builder
	started := false
	for _, r := range s {
		if r == squot {
			if started {
				output.WriteRune(r)
			}
			output.WriteString(` "'" `)
			started = false
		} else if !started {
			output.WriteString(`r'`)
			output.WriteRune(r)
			started = true
		} else {
			output.WriteRune(r)
		}
	}
	if started {
		output.WriteRune(squot)
	}
	return output.String()
}

func ArbParams(holders arb.ArbPlaceholders, withType bool) string {
	ps := make([]string, len(holders))
	for i, lp := range holders.LangParams("dart") {
		if withType {
			ps[i] = lp.Info + " " + lp.Name
		} else {
			ps[i] = lp.Name
		}
	}
	return strings.Join(ps, ",")
}
