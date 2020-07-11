package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
	pgs "github.com/lyft/protoc-gen-star"
)

func (r *RPCSUtil) checkHasGroupElement(m pgs.Message, name, ooName pgs.Name) error {
	for _, mm := range m.Messages() {
		if mm.Name() == name {
			return r.checkHasGroupElementType(mm, ooName)
		}
	}
	return fmt.Errorf("%s not found: %s", name, m.FullyQualifiedName())
}

func (r *RPCSUtil) checkHasGroupElementType(m pgs.Message, ooName pgs.Name) error {
	ft, err := r.findFieldTypeFrom(m, "type")
	if err != nil {
		return err
	}

	for _, e := range m.Enums() {
		if e.Name() == "Type" {
			if !ft.IsEnum() || ft.Enum() != e {
				return fmt.Errorf("field [type] must be type of enum Type: %s",
					m.FullyQualifiedName())
			}
			return r.checkElementType(m, e, ooName)
		}
	}
	return fmt.Errorf("Type not found: %s", m.FullyQualifiedName())
}

func (r *RPCSUtil) findFieldTypeFrom(m pgs.Message, n pgs.Name) (pgs.FieldType, error) {
	for _, f := range m.NonOneOfFields() {
		if f.Name() == n {
			return f.Type(), nil
		}
	}
	return nil, fmt.Errorf("field %s not found: %s", n, m.FullyQualifiedName())
}

func (r *RPCSUtil) checkElementType(m pgs.Message, e pgs.Enum, ooName pgs.Name) error {
	// from source
	vs := e.Values()
	if vs[0].Name() != "invalid" {
		return fmt.Errorf("zero value be invalid from Element: %s", e.FullyQualifiedName())
	}

	oo, err := r.findOneof(m, ooName)
	if err != nil {
		return err
	}

	fs := oo.Fields()
	if len(vs) != len(fs)+1 {
		return fmt.Errorf("Values must be sync with %s: %s", ooName, e.FullyQualifiedName())
	}

	for i, v := range vs[1:] {
		if v.Name() != fs[i].Name() {
			return fmt.Errorf("Values must be sync with %s: %s", ooName, e.FullyQualifiedName())
		}
	}
	return nil
}

func (r *RPCSUtil) checkHasOOAllFieldsType(m pgs.Message, ooName pgs.Name, rts ...form.Node_Type) error {
	oo, err := r.findOneof(m, ooName)
	if err != nil {
		return err
	}

	ok, err := r.checkOOAllFieldsType(oo, rts...)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("all %s fields must be types of %v: %s",
			ooName, rts, oo.FullyQualifiedName())
	}
	return nil
}

func (r *RPCSUtil) checkOOAllFieldsType(oo pgs.OneOf, rts ...form.Node_Type) (bool, error) {
	fs := oo.Fields()
	if len(fs) == 0 {
		return false, nil
	}
	for _, f := range fs {
		ok, err := r.checkFieldType(f, rts...)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}

func (r *RPCSUtil) checkFieldType(f pgs.Field, rts ...form.Node_Type) (bool, error) {
	typ := f.Type()
	if !typ.IsEmbed() {
		return false, fmt.Errorf("field type must be a Message: %s", f.FullyQualifiedName())
	}
	return r.CheckNodeType(typ.Embed(), rts...)
}

func (r *RPCSUtil) NodeType(p pgs.Message) (form.Node_Type, error) {
	var node form.Node
	_, err := p.Extension(form.E_Node, &node)
	if err != nil {
		return form.Node_undefined, err
	}
	return node.Type, nil
}

func (r *RPCSUtil) CheckNodeType(p pgs.Message, rts ...form.Node_Type) (bool, error) {
	nt, err := r.NodeType(p)
	if err != nil {
		return false, err
	}

	for _, rt := range rts {
		if nt == rt {
			return true, nil
		}
	}
	return false, nil
}
