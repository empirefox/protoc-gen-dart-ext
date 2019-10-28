package arb

import (
	"testing"
)

func TestLangParam(t *testing.T) {
	in := ArbPlaceholders{
		&ArbPlaceholder{
			Name: "form",
			LangInfos: ArbLangInfos{
				&ArbLangInfo{
					Lang:   "dart",
					Type:   "Form",
					Import: "package:pgde/plural.dart",
				},
			},
		},
	}

	lp := in.LangParam("dart", "form")
	if lp == nil {
		t.Fatalf("form field should be found")
	}

	if lp.Type != "Form" {
		t.Fatalf("form field info should be Form, but got: %s", lp.Type)
	}
}
