package genshared

import (
	"bytes"
	"strconv"
	"strings"
	"text/template"
)

//go:generate pie renderJoinGroup.SortStableUsing
type renderJoinGroup []*renderJoinElement

type renderJoinElement struct {
	buf   bytes.Buffer
	name  string
	prior int
}

func renderJoinElementGreater(a, b *renderJoinElement) bool {
	return a.prior > b.prior
}

func newRenderJoinElement(s string) (e *renderJoinElement, err error) {
	e = &renderJoinElement{name: s}
	idx := strings.LastIndexAny(s, "+-")
	if idx != -1 {
		e.prior, err = strconv.Atoi(s[idx:])
		if err != nil {
			return
		}
		e.name = s[:idx]
	}
	return
}

func RenderJoin(tpl *template.Template) func(interface{}, ...string) (string, error) {
	return func(data interface{}, tplNames ...string) (r string, err error) {
		group := make(renderJoinGroup, len(tplNames))
		for i, name := range tplNames {
			group[i], err = newRenderJoinElement(name)
			if err != nil {
				return
			}
		}

		for _, e := range group.SortStableUsing(renderJoinElementGreater) {
			err = tpl.ExecuteTemplate(&e.buf, e.name, data)
			if err != nil {
				return
			}
		}

		s := make([][]byte, len(tplNames))
		for i, e := range group {
			s[i] = e.buf.Bytes()
		}
		return string(bytes.Join(s, nil)), nil
	}
}
