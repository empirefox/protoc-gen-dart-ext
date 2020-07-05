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

func (r *RPCSUtil) ParentMessage(embed pgs.Message) (pgs.Message, error) {
	p, ok := embed.Parent().(pgs.Message)
	if !ok {
		return nil, fmt.Errorf("Must be defined under message: %s",
			embed.FullyQualifiedName())
	}
	return p, nil
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

	return r.IdType(pf)
}

func (r *RPCSUtil) IdType(pf pgs.Field) (dart.Qualifier, error) {
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
