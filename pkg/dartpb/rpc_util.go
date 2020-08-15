package dartpb

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/sets/treeset"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
	pgs "github.com/lyft/protoc-gen-star"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

type RPCSUtil struct {
	Pkg     pgs.Package
	Imports *treeset.Set
}

func (r *RPCSUtil) TypeIsElement(f pgs.Field) (ok bool, err error) {
	v, err := r.TypeOf(f)
	if err != nil {
		return false, err
	}
	return r.CheckNodeType(v, form.Node_element)
}

func (r *RPCSUtil) TypeHasElement(f pgs.Field) (ok bool, err error) {
	v, err := r.TypeOf(f)
	if err != nil {
		return false, err
	}
	return r.HasElement(v)
}

func (r *RPCSUtil) HasElement(v pgs.Message) (ok bool, err error) {
	for _, c := range v.Messages() {
		ok, err = r.CheckNodeType(c, form.Node_element)
		if err != nil || ok {
			return
		}
	}
	return
}

func (r *RPCSUtil) TypeIsGroupView(f pgs.Field) (ok bool, err error) {
	v, err := r.TypeOf(f)
	if err != nil {
		return false, err
	}
	return r.IsGroupView(v)
}

func (r *RPCSUtil) IsGroupView(v pgs.Message) (ok bool, err error) {
	ok, err = r.IsView(v)
	if !ok {
		return
	}

	return r.IsGroup(v), nil
}

func (r *RPCSUtil) IsView(v pgs.Message) (ok bool, err error) {
	var ext form.Node
	ok, err = v.Extension(form.E_Node, &ext)
	if !ok || err != nil {
		return
	}

	return ext.Type == form.Node_view || ext.Type == form.Node_selectManyView, nil
}

func (r *RPCSUtil) Payload(f pgs.Field) (dart.Qualifier, error) {
	dt, err := r.TypeOf(f)
	if err != nil {
		return "", err
	}

	var ext form.Node
	ok, err := dt.Extension(form.E_Node, &ext)
	if !ok || err != nil {
		return "", fmt.Errorf("Field type is not a node: %s", f.FullyQualifiedName())
	}

	switch ext.Type {
	case form.Node_leaf:
		return r.ProtoRef(dt).PayloadSuffix(), nil
	case form.Node_view, form.Node_selectManyView:
		if r.IsGroup(dt) {
			return r.ProtoRef(dt).PayloadSuffix(), nil
		}

		p, err := r.ParentMessage(dt)
		if err != nil {
			return "", err
		}

		return r.ProtoRef(p).PayloadSuffix(), nil
	}

	return "", fmt.Errorf("Payload only works for leaf, view: %s",
		f.FullyQualifiedName())
}

func (r *RPCSUtil) ParentMessageWithType(embed pgs.Message, nt ...form.Node_Type) (pgs.Message, error) {
	p, err := r.ParentMessage(embed)
	if err != nil {
		return nil, err
	}

	ok, err := r.CheckNodeType(p, nt...)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Parent must be type of %s: %s", nt,
			embed.FullyQualifiedName())
	}

	return p, nil
}

func (r *RPCSUtil) ParentMessage(embed pgs.Message) (pgs.Message, error) {
	p, ok := embed.Parent().(pgs.Message)
	if !ok {
		return nil, fmt.Errorf("Must be defined under message: %s",
			embed.FullyQualifiedName())
	}
	return p, nil
}

func (r *RPCSUtil) TypeNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	t, err := r.ViewTypeOf(pf)
	if err != nil {
		return "", err
	}

	t, err = r.ParentMessageWithType(t, form.Node_leaf)
	if err != nil {
		return "", err
	}

	return r.ProtoRef(t), nil
}

func (r *RPCSUtil) ViewNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	t, err := r.ViewTypeOf(pf)
	if err != nil {
		return "", err
	}

	return r.ProtoRef(t), nil
}

func (r *RPCSUtil) ViewTypeOf(pf pgs.Field) (pgs.Message, error) {
	embed, err := r.TypeOf(pf)
	if err != nil {
		return nil, err
	}

	ok, err := r.IsSelectMany(pf)
	if err != nil {
		return nil, err
	}
	if ok {
		return r.ParentMessageWithType(embed, form.Node_view, form.Node_selectManyView)
	}

	ok, err = r.CheckNodeType(embed, form.Node_view, form.Node_selectManyView)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Type must be %s: %s or %s",
			form.Node_view, form.Node_selectManyView, pf.FullyQualifiedName())
	}

	return embed, nil
}

func (r *RPCSUtil) hasView(m pgs.Message) (bool, error) {
	for _, c := range m.Messages() {
		var ext form.Node
		_, err := c.Extension(form.E_Node, &ext)
		if err != nil {
			return false, err
		}

		if ext.Type == form.Node_view || ext.Type == form.Node_selectManyView {
			return true, nil
		}
	}
	return false, nil
}

const ElementName = "Element"

func (r *RPCSUtil) findElement(m pgs.Message) pgs.Message {
	return r.findEmbed(m, ElementName)
}
func (r *RPCSUtil) findEmbedId(m pgs.Message) pgs.Message {
	return r.findEmbed(m, "Id")
}

func (r *RPCSUtil) findEmbed(m pgs.Message, name dart.Qualifier) pgs.Message {
	for _, mm := range m.Messages() {
		if dart.Qualifier(mm.Name()) == name {
			return mm
		}
	}
	return nil
}

func (r *RPCSUtil) findOneof(m pgs.Message, ooName pgs.Name) (pgs.OneOf, error) {
	for _, oo := range m.OneOfs() {
		if oo.Name() == ooName {
			return oo, nil
		}
	}
	return nil, fmt.Errorf("oneof [%s] not found: %s", ooName, m.FullyQualifiedName())
}

func (r *RPCSUtil) IsGroup(m pgs.Message) bool {
	_, err := r.findOneof(m, "group")
	return err == nil
}

func (r *RPCSUtil) IsSelectMany(f pgs.Field) (ok bool, err error) {
	var nd form.Node
	ok, err = f.Message().Extension(form.E_Node, &nd)
	if !ok || err != nil {
		return
	}

	if nd.Type == form.Node_entry {
		var ext form.EntryField
		ok, err = f.Extension(form.E_EntryField, &ext)
		if !ok || err != nil {
			return
		}
		return ext.GetSelectMany() != nil, nil
	}

	if nd.Type == form.Node_leaf {
		var ext form.Field
		ok, err = f.Extension(form.E_Field, &ext)
		if !ok || err != nil {
			return
		}
		return ext.GetSelectMany() != nil, nil
	}

	return false, nil
}

func (r *RPCSUtil) IsId(f pgs.Field) (ok bool, err error) {
	_, err = f.Extension(form.E_Id, &ok)
	return
}

func (r *RPCSUtil) SelectManyElement(m pgs.Message) (pgs.Message, error) {
	for _, n := range m.Messages() {
		if n.Name() == "Element" {
			return n, nil
		}
	}
	return nil, fmt.Errorf("Element not found: %s", m.FullyQualifiedName())
}

func (r *RPCSUtil) ElementTypeOf(pf pgs.Field) (pgs.Message, error) {
	pgsTyp := pf.Type()
	if !pgsTyp.IsRepeated() {
		return nil, fmt.Errorf("Get Element type from non-repeated field: %s",
			pf.FullyQualifiedName())
	}

	elemTyp := pgsTyp.Element()
	if !elemTyp.IsEmbed() {
		return nil, fmt.Errorf("Get ViewType from non-message field: %s",
			pf.FullyQualifiedName())
	}

	return elemTyp.Embed(), nil
}

func (r *RPCSUtil) TypeOf(pf pgs.Field) (m pgs.Message, err error) {
	var isEmbed bool
	pgsTyp := pf.Type()
	if pgsTyp.IsRepeated() {
		elemTyp := pgsTyp.Element()
		isEmbed = elemTyp.IsEmbed()
		if isEmbed {
			m = elemTyp.Embed()
		}
	} else {
		isEmbed = pgsTyp.IsEmbed()
		if isEmbed {
			m = pgsTyp.Embed()
		}
	}

	if !isEmbed {
		err = fmt.Errorf("Get ViewType from non-message field: %s",
			pf.FullyQualifiedName())
	}
	return
}

func (r *RPCSUtil) IdTypeOfFieldType(f pgs.Field) (dart.Qualifier, error) {
	vt, err := r.TypeOf(f)
	if err != nil {
		return "", nil
	}
	return r.IdTypeOfViewOrType(vt)
}

func (r *RPCSUtil) IdTypeOfViewOrType(m pgs.Message) (dart.Qualifier, error) {
	pf, err := r.FindIdFieldFromView(m)
	if err != nil {
		return "", err
	}

	if pf == nil {
		return "", fmt.Errorf("No id found for: %s", m.FullyQualifiedName())
	}

	return r.ProtoType(pf)
}

func (r *RPCSUtil) ProtoType(pf pgs.Field) (dart.Qualifier, error) {
	n := pf.Type().ProtoType().String()
	if strings.HasPrefix(n, "TYPE_") {
		n = strings.ToLower(n[5:])
	}
	return dart.Qualifier(n), nil
}

func (r *RPCSUtil) FindIdFieldFromView(v pgs.Message) (pgs.Field, error) {
	vid, err := r.IdField(v)
	if err != nil {
		return nil, err
	}

	if vid != nil {
		return vid, nil
	}

	m, ok := v.Parent().(pgs.Message)
	if !ok {
		return nil, nil
	}

	return r.IdField(m)
}

func (r *RPCSUtil) IdField(m pgs.Message) (pgs.Field, error) {
	for _, f := range m.NonOneOfFields() {
		ok, err := r.IsId(f)
		if err != nil {
			return nil, err
		}
		if ok {
			return f, nil
		}
	}
	return nil, nil
}

func (r *RPCSUtil) Empty() (dart.Qualifier, error) {
	return r.ProtoRefWKT(pgs.EmptyWKT)
}
func (r *RPCSUtil) UInt32Value() (dart.Qualifier, error) {
	return r.ProtoRefWKT(pgs.UInt32ValueWKT)
}

func (r *RPCSUtil) IdOrEmpty(m pgs.Message) (dart.Qualifier, error) {
	id, err := r.Id(m)
	if err != nil {
		return "", err
	}

	if id != "" {
		return id, nil
	}
	return r.ProtoRefWKT(pgs.EmptyWKT)
}

func (r *RPCSUtil) Id(m pgs.Message) (dart.Qualifier, error) {
	id, err := r.IdField(m)
	if err != nil {
		return "", err
	}
	if id != nil {
		it := id.Type()
		pt := it.ProtoType()
		if pt.IsNumeric() || pt == pgs.StringT {
			return dart.Qualifier(pt.String()), nil
		}
		if it.IsEmbed() {
			return r.ProtoRef(it.Embed()), nil
		}
		return "", fmt.Errorf("Unsupported id type: %s", id.FullyQualifiedName())
	}
	return "", nil
}

func (r *RPCSUtil) AddImportStr(p string) {
	if !r.Imports.Contains(p) {
		r.Imports.Add(p)
	}
}

func (r *RPCSUtil) AddImport(p dart.Qualifier) { r.AddImportStr(string(p)) }

func (r *RPCSUtil) ProtoRefWKT(t pgs.WellKnownType) (dart.Qualifier, error) {
	return r.ProtoRefFQN(dart.Qualifier(pgs.WellKnownTypePackage + "." + t.Name()))
}

func (r *RPCSUtil) ProtoRefRaw(protofile, pkg, fqn dart.Qualifier) (dart.Qualifier, error) {
	if pgs.Name(pkg) == r.Pkg.ProtoName() {
		fqn = fqn[len(pkg)+1:]
		if fqn[0] == '.' {
			fqn = fqn[1:]
		}
	}

	r.AddImport(protofile)
	return fqn, nil

}

func (r *RPCSUtil) ProtoRefFQN(fqn dart.Qualifier) (dart.Qualifier, error) {
	fn := protoreflect.FullName(fqn)
	mt, err := protoregistry.GlobalTypes.FindMessageByName(fn)
	if err != nil {
		return "", err
	}

	pf := mt.Descriptor().ParentFile()
	if pgs.Name(pf.Package()) == r.Pkg.ProtoName() {
		fqn = fqn[len(pf.Package())+1:]
		if fqn[0] == '.' {
			fqn = fqn[1:]
		}
	}

	r.AddImportStr(pf.Path())
	return fqn, nil
}

func (r *RPCSUtil) ProtoRefFieldType(f pgs.Field) (dart.Qualifier, error) {
	m, err := r.TypeOf(f)
	if err != nil {
		return "", err
	}
	return r.ProtoRef(m), nil
}

func (r *RPCSUtil) ProtoRef(m pgs.Message) dart.Qualifier {
	ref := dart.Qualifier(m.FullyQualifiedName())
	if m.Package() == r.Pkg {
		ref = ref[len(m.Package().ProtoName())+1:]
		if ref[0] == '.' {
			ref = ref[1:]
		}
	}

	r.AddImport(dart.Qualifier(m.File().Name()))
	return ref
}
