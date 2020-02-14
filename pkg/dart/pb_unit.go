package dart

import (
	"fmt"
	"strings"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/units"
	"github.com/lynn9388/supsub"
)

func (im *ImportManager) RenderUnit(u *units.Unit) Qualifier {
	if u == nil || len(u.Dots) == 0 && len(u.Per) == 0 {
		return "null"
	}
	return "const " + im.PgdeFile().AsDot("Unit").InitWith(
		im.renderUnitShow(u.Show),
		im.renderUnitCells(u.Dots),
		im.renderUnitCells(u.Per),
	)
}

func (im *ImportManager) renderUnitShow(s *units.Show) Qualifier {
	if s == nil {
		s = new(units.Show)
	}
	if !s.Off && s.Prefix == 0 && s.Atom == 0 {
		return im.PgdeFile().AsDot("Unit").Dot("showSymbol")
	}
	showType := im.PgdeFile().AsDot("Show")
	return "const " + showType.InitWithString(
		fmt.Sprintf("%v", s.Off),
		string(showType.DotStringer(s.Prefix))+"Prefix",
		string(showType.DotStringer(s.Atom)),
	)
}

func (im *ImportManager) renderUnitCells(cells []*units.Cell) Qualifier {
	if len(cells) == 0 {
		return "null"
	}

	s := make([]string, len(cells))
	for i, c := range cells {
		s[i] = string(im.renderUnitCell(c))
	}
	return Qualifier("const [" + strings.Join(s, ", ") + "]")
}

func (im *ImportManager) renderUnitCell(c *units.Cell) Qualifier {
	if c == nil {
		c = new(units.Cell)
	}
	args := make([]string, 0, 3)
	args = append(args, string(im.PgdeFile().
		AsDot("Prefix").
		DotStringer(c.Prefix)))

	atomType := im.PgdeFile().AsDot("Atom")
	switch t := c.Type.(type) {
	case *units.Cell_Atom:
		args = append(args, string(atomType.DotStringer(t.Atom)))
	case *units.Cell_Symbol:
		args = append(args, "const "+string(atomType.Dot("symbol").
			InitWithString(RawString(t.Symbol))))
	case nil:
		args = append(args, string(atomType.DotStringer(units.Atom_noAtom)))
	}

	if c.Exponent != 0 || c.Exponent != 1 {
		args = append(args, RawString(supsub.ToSup(fmt.Sprintf("%d", c.Exponent))))
	}
	return "const " + im.PgdeFile().AsDot("Cell").InitWithString(args...)
}
