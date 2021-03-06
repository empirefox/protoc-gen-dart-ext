package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
	pgs "github.com/lyft/protoc-gen-star"
)

func (r *RPCSUtil) SrcId(f pgs.Field) (dart.Qualifier, error) {
	m, err := r.TypeOf(f)
	if err != nil {
		return "", err
	}

	ok, err := r.CheckNodeType(m, form.Node_element)
	if err != nil {
		return "", err
	}

	if ok {
		m, err = r.ParentMessageWithType(m, form.Node_view, form.Node_selectManyView)
		if err != nil {
			return "", err
		}

		return r.ProtoRef(m).Dot("Id"), nil
	}

	ok, err = r.CheckNodeType(m, form.Node_view, form.Node_selectManyView)
	if err != nil {
		return "", err
	}

	if ok && r.IsGroup(m) {
		hasElement, err := r.HasElement(m)
		if err != nil {
			return "", err
		}
		if hasElement {
			return r.ProtoRef(m).Dot("Id"), nil
		}
	}

	pn, err := r.Payload(f)
	if err != nil {
		return "", err
	}
	return pn.Dot("SrcId"), nil
}

func (r *RPCSUtil) SrcResp(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	return r.anyViewDot(d, f, "SrcResp")
}

func (r *RPCSUtil) GetResp(f pgs.Field) (dart.Qualifier, error) {
	return r.payloadDot(f, "GetResp")
}
func (r *RPCSUtil) ListAddOrSaveReq(f pgs.Field) (dart.Qualifier, error) {
	ok, err := r.TypeIsGroupView(f)
	if err != nil {
		return "", err
	}
	if ok {
		pn, err := r.Payload(f)
		if err != nil {
			return "", err
		}
		return pn.Dot("SrcElement"), nil
	}
	return r.ProtoRefFieldType(f)
}
func (r *RPCSUtil) AddResp(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	return r.anyViewDot(d, f, "AddResp")
}
func (r *RPCSUtil) SrcIds(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	return r.payloadDot(f, "SrcIds")
}
func (r *RPCSUtil) DstResp(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	ok, err := r.isBareSelectMany(f)
	if err != nil {
		return "", err
	}

	if ok {
		return r.anyViewDot(d, f, "SrcResp")
	}
	return r.anyViewDot(d, f, "DstResp")
}
func (r *RPCSUtil) SelectResp(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	ok, err := r.isBareSelectMany(f)
	if err != nil {
		return "", err
	}

	if ok {
		return r.Empty()
	}
	return r.anyViewDot(d, f, "SelectResp")
}
func (r *RPCSUtil) SelectManyResp(d *dart.Dart, f pgs.Field) (dart.Qualifier, error) {
	ok, err := r.isBareSelectMany(f)
	if err != nil {
		return "", err
	}

	if ok {
		return r.Empty()
	}
	return r.anyViewDot(d, f, "DstResp")
}

func (r *RPCSUtil) payloadDot(f pgs.Field, name dart.Qualifier) (dart.Qualifier, error) {
	pn, err := r.Payload(f)
	if err != nil {
		return "", err
	}
	return pn.Dot(name), nil
}

func (r *RPCSUtil) anyViewDot(d *dart.Dart, f pgs.Field, name dart.Qualifier) (dart.Qualifier, error) {
	ok, err := r.TypeIsElement(f)
	if err != nil {
		return "", err
	}

	if ok {
		v, err := r.TypeOf(f)
		if err != nil {
			return "", err
		}

		v, err = r.ParentMessageWithType(v, form.Node_selectManyView)
		if err != nil {
			return "", err
		}

		if r.IsGroup(v) {
			return d.NameOf(v).PayloadSuffix().Dot(name), nil
		}

		l, err := r.ParentMessageWithType(v, form.Node_leaf)
		if err != nil {
			return "", err
		}

		return d.NameOf(l).PayloadSuffix().
			DotStringer(v.Name()).
			Dot(name), nil
	}

	return r.anyViewDotFromViewField(f, name)
}

func (r *RPCSUtil) anyViewDotFromViewField(f pgs.Field, name dart.Qualifier) (dart.Qualifier, error) {
	pn, err := r.Payload(f)
	if err != nil {
		return "", err
	}

	typ, err := r.TypeOf(f)
	if err != nil {
		return "", err
	}

	if r.IsGroup(typ) {
		return pn.Dot(name), nil
	}

	return pn.Dot(dart.Qualifier(typ.Name())).Dot(name), nil
}

func (r *RPCSUtil) isBareSelectMany(f pgs.Field) (bool, error) {
	typ, err := r.TypeOf(f)
	if err != nil {
		return false, err
	}

	isElement, err := r.HasElement(typ)
	if err != nil {
		return false, err
	}

	isSelectMany, err := r.CheckNodeType(typ, form.Node_selectManyView)
	if err != nil {
		return false, err
	}

	if isElement && !isSelectMany {
		return false, fmt.Errorf("Element must be defined under selectManyView: %s",
			typ.FullyQualifiedName())
	}

	return isSelectMany && !isElement, nil
}
