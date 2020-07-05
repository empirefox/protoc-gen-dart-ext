package dartpb

import (
	"fmt"

	"github.com/empirefox/protoc-gen-dart-ext/pkg/dart"
	"github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/form"
	pgs "github.com/lyft/protoc-gen-star"
)

type RPCS struct {
	RPCSUtil
	File         *File
	RootFilePath string
}

type RPCMessage struct {
	*Message
	Extension form.Node
}

func (m *RPCMessage) RPCS() *RPCS { return m.File.RPCS }

func (m *RPCMessage) IsUndefined() bool { return m.Extension.Type == form.Node_undefined }
func (m *RPCMessage) IsEntry() bool     { return m.Extension.Type == form.Node_entry }
func (m *RPCMessage) IsLeaf() bool      { return m.Extension.Type == form.Node_leaf }
func (m *RPCMessage) IsView() bool      { return m.Extension.Type == form.Node_view }
func (m *RPCMessage) IsTypeName() bool  { return m.Extension.Type == form.Node_typeName }
func (m *RPCMessage) IsElement() bool   { return m.Extension.Type == form.Node_element }

type RPCField struct {
	*Field
	EntryExt form.EntryField
	LeafExt  form.Field
}

func (f *RPCField) IsUndefined() bool { return f.Message.RPC.IsUndefined() }
func (f *RPCField) IsEntry() bool     { return f.Message.RPC.IsEntry() }
func (f *RPCField) IsLeaf() bool      { return f.Message.RPC.IsLeaf() }
func (f *RPCField) IsView() bool      { return f.Message.RPC.IsView() }
func (f *RPCField) IsTypeName() bool  { return f.Message.RPC.IsTypeName() }
func (f *RPCField) IsElement() bool   { return f.Message.RPC.IsElement() }

func (f *RPCField) IsForm() bool {
	return f.EntryExt.GetForm() != nil || f.EntryExt.Is == nil
}
func (f *RPCField) IsList() bool { return f.EntryExt.GetList() != nil }
func (f *RPCField) IsSelectOne() bool {
	if f.IsLeaf() {
		return f.LeafExt.GetSelectOne() != nil
	}
	return f.EntryExt.GetSelectOne() != nil
}
func (f *RPCField) IsSelectMany() bool {
	if f.IsLeaf() {
		return f.LeafExt.GetSelectMany() != nil
	}
	return f.EntryExt.GetSelectMany() != nil
}

func (f *RPCField) ServiceName() dart.Qualifier { return f.Names.ServiceName() }
func (f *RPCField) PayloadName() dart.Qualifier { return f.Names.PayloadName() }

func (f *RPCField) Group() (pgs.OneOf, error) {
	v, err := f.ViewType()
	if err != nil {
		return nil, err
	}
	return f.File.RPCS.findOneof(v, "group")
}

func (f *RPCField) IsGroup() (bool, error) {
	vm, err := f.ViewType()
	if err != nil {
		return false, err
	}
	return f.File.RPCS.IsGroup(vm), nil
}

func (f *RPCField) IsGroupByType() (bool, error) {
	isGroup, err := f.IsGroup()
	if err != nil {
		return false, err
	}

	if !isGroup {
		return false, nil
	}

	if f.IsLeaf() {
		if f.IsSelectOne() {
			return f.LeafExt.GetSelectOne().GroupByType, nil
		}
		if f.IsSelectMany() {
			return f.LeafExt.GetSelectMany().Src.GetGroupByType(), nil
		}
		return false, nil
	}

	if f.IsList() {
		return f.EntryExt.GetList().GroupByType, nil
	}
	if f.IsSelectOne() {
		return f.EntryExt.GetSelectOne().GroupByType, nil
	}
	if f.IsSelectMany() {
		return f.EntryExt.GetSelectMany().Src.GetGroupByType(), nil
	}
	return false, nil
}

func (f *RPCField) RefTypeName() (dart.Qualifier, error) {
	isGroup, err := f.IsGroup()
	if err != nil {
		return "", err
	}

	if isGroup {
		// the generated Element
		return f.PayloadName().Dot("Element"), nil
	}
	return f.NonGroupTypeName()
}

func (f *RPCField) NonGroupTypeName() (dart.Qualifier, error) {
	m, err := f.NonGroupType()
	if err != nil {
		return "", err
	}
	return f.File.RPCS.ProtoRef(m), nil
}

func (f *RPCField) NonGroupType() (m pgs.Message, err error) {
	vm, err := f.ViewType()
	if err != nil {
		return nil, err
	}

	m, ok := vm.Parent().(pgs.Message)
	if !ok {
		return nil, fmt.Errorf("view must be defined within leaf")
	}
	return m, nil
}

func (f *RPCField) ElementType() (pgs.Message, error) {
	return f.File.RPCS.ElementTypeOf(f.Pgs)
}

func (f *RPCField) ElementIdFieldType() (dart.Qualifier, error) {
	id, err := f.ElementIdField()
	if err != nil {
		return "", err
	}
	return f.File.RPCS.IdType(id)
}

func (f *RPCField) ElementIdField() (pgs.Field, error) {
	typ, err := f.ElementType()
	if err != nil {
		return nil, err
	}

	id, err := f.File.RPCS.IdField(typ)
	if err != nil {
		return nil, err
	}

	if id == nil {
		return nil, fmt.Errorf("id not found: %s", typ.FullyQualifiedName())
	}
	return id, nil
}

func (f *RPCField) ElementName() (dart.Qualifier, error) {
	m, err := f.ElementType()
	if err != nil {
		return "", err
	}
	return f.File.RPCS.ProtoRef(m), nil
}

func (f *RPCField) ViewName() (dart.Qualifier, error) {
	m, err := f.ViewType()
	if err != nil {
		return "", err
	}
	return f.File.RPCS.ProtoRef(m), nil
}

func (f *RPCField) TypeName() (dart.Qualifier, error) {
	if f.IsForm() {
		return f.File.RPCS.ProtoRef(f.Pgs.Type().Embed()), nil
	}
	return f.TypeNameOfFieldType(f.Pgs)
}

func (f *RPCField) TypeNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	v, err := f.ViewTypeOf(pf)
	if err != nil {
		return "", nil
	}

	m, ok := v.Parent().(pgs.Message)
	if !ok {
		return "", fmt.Errorf("Leaf not found for view: %s", v.FullyQualifiedName())
	}
	return f.File.RPCS.ProtoRef(m), nil
}

func (f *RPCField) ViewNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	m, err := f.ViewTypeOf(pf)
	if err != nil {
		return "", nil
	}
	return f.File.RPCS.ProtoRef(m), nil
}

func (f *RPCField) ViewType() (pgs.Message, error) { return f.ViewTypeOf(f.Pgs) }

func (f *RPCField) ViewTypeOf(pf pgs.Field) (pgs.Message, error) {
	embed, err := f.File.RPCS.TypeOf(pf)
	if err != nil {
		return nil, err
	}

	ok, err := f.File.RPCS.IsSelectMany(pf)
	if err != nil {
		return nil, err
	}
	if ok {
		return f.ParentMessage(embed, form.Node_view)
	}

	ok, err = f.File.RPCS.CheckNodeType(embed, form.Node_view)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Type must be %s: %s", form.Node_view,
			pf.FullyQualifiedName())
	}

	return embed, nil
}

func (f *RPCField) ParentMessage(embed pgs.Message, nt form.Node_Type) (pgs.Message, error) {
	p, err := f.Message.RPC.RPCS().ParentMessage(embed)
	if err != nil {
		return nil, err
	}

	ok, err := f.File.RPCS.CheckNodeType(p, nt)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("Parent must be type of %s: %s", nt,
			embed.FullyQualifiedName())
	}

	return p, nil
}

func (f *RPCField) BackendErrorName() (dart.Qualifier, error) {
	return f.Message.RPC.RPCS().ProtoRefRaw("pgde/error/error.proto",
		"pgde.error", "pgde.error.BackendError")
}

func (f *RPCField) IdTypeOfFieldType(oof pgs.Field) (dart.Qualifier, error) {
	return f.File.RPCS.IdTypeOfFieldType(oof)
}

func (f *RPCField) Empty() (dart.Qualifier, error) { return f.Message.File.RPCS.Empty() }
func (f *RPCField) IdOrEmpty() (dart.Qualifier, error) {
	return f.Message.File.RPCS.IdOrEmpty(f.Message.Pgs)
}

func (f *RPCField) Remove() bool {
	if f.IsLeaf() {
		if ss := f.LeafExt.GetSelectOne(); ss != nil {
			return ss.Remove
		}
		if ss := f.LeafExt.GetSelectMany(); ss != nil {
			return ss.Dst.GetRemove()
		}
		return false
	}

	if f.IsList() {
		return f.EntryExt.GetList().Remove
	}
	if f.IsSelectOne() {
		return f.EntryExt.GetSelectOne().Remove
	}
	if f.IsSelectMany() {
		return f.EntryExt.GetSelectMany().Dst.GetRemove()
	}
	return f.EntryExt.GetForm().Remove
}
func (f *RPCField) RemoveAll() bool {
	if f.IsLeaf() {
		if ss := f.LeafExt.GetSelectMany(); ss != nil {
			return ss.Dst.GetRemoveAll()
		}
		return false
	}

	if f.IsList() {
		return f.EntryExt.GetList().RemoveAll
	}
	if f.IsSelectMany() {
		return f.EntryExt.GetSelectMany().Dst.GetRemoveAll()
	}
	return false
}

func (f *RPCField) Add() bool {
	return f.IsList() && f.EntryExt.GetList().Add ||
		f.IsSelectMany() && f.EntryExt.GetSelectMany().Dst.GetAdd()
}
func (f *RPCField) Save() bool { return f.IsList() && f.EntryExt.GetList().Edit }
