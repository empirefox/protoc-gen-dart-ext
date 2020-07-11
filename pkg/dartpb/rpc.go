package dartpb

import (
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
	Extension   form.Node
	IdField     pgs.Field
	IdMessage   pgs.Message
	ViewElement *RPCMessage
	ParentLeaf  *RPCMessage
	ParentView  *RPCMessage
	// LeafViews requires IdField
	LeafViews []*RPCMessage
	ViewGroup pgs.OneOf
}

func (m *RPCMessage) ViewUnderLeaf() bool { return m.IsView() && m.ParentLeaf != nil }
func (m *RPCMessage) ViewHasGroup() bool  { return m != nil && m.ViewGroup != nil }
func (m *RPCMessage) LeafHasView() bool   { return m != nil && len(m.LeafViews) != 0 }
func (m *RPCMessage) LeafHasViewElement() bool {
	if m == nil {
		return false
	}
	for _, v := range m.LeafViews {
		if v.ViewElement != nil {
			return true
		}
	}
	return false
}

func (m *RPCMessage) RPCS() *RPCS { return m.File.RPCS }

func (m *RPCMessage) IsUndefined() bool { return m == nil || m.Extension.Type == form.Node_undefined }
func (m *RPCMessage) IsEntry() bool     { return m != nil && m.Extension.Type == form.Node_entry }
func (m *RPCMessage) IsLeaf() bool      { return m != nil && m.Extension.Type == form.Node_leaf }
func (m *RPCMessage) IsView() bool      { return m != nil && m.Extension.Type == form.Node_view }
func (m *RPCMessage) IsElement() bool   { return m != nil && m.Extension.Type == form.Node_element }

func (m *RPCMessage) ProtoType(oof pgs.Field) (dart.Qualifier, error) {
	return m.File.RPCS.ProtoType(oof)
}
func (m *RPCMessage) IdTypeOfFieldType(oof pgs.Field) (dart.Qualifier, error) {
	return m.File.RPCS.IdTypeOfFieldType(oof)
}

func (m *RPCMessage) ViewNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	return m.RPCS().ViewNameOfFieldType(pf)
}
func (m *RPCMessage) TypeNameOfFieldType(pf pgs.Field) (dart.Qualifier, error) {
	return m.RPCS().TypeNameOfFieldType(pf)
}

func (m *RPCMessage) ProtoRef() dart.Qualifier { return m.RPCS().ProtoRef(m.Pgs) }

// For leaf, group-view
func (m *RPCMessage) PayloadName() dart.Qualifier { return m.DartName.PayloadSuffix() }

func (m *RPCMessage) BackendErrorName() (dart.Qualifier, error) {
	return m.RPCS().ProtoRefRaw("pgde/error/error.proto",
		"pgde.error", "pgde.error.BackendError")
}

type RPCField struct {
	*Field
	EntryExt form.EntryField
	LeafExt  form.Field
}

func (f *RPCField) RPCS() *RPCS { return f.File.RPCS }

func (f *RPCField) IsUndefined() bool { return f.Message.RPC.IsUndefined() }
func (f *RPCField) IsEntry() bool     { return f.Message.RPC.IsEntry() }
func (f *RPCField) IsLeaf() bool      { return f.Message.RPC.IsLeaf() }
func (f *RPCField) IsView() bool      { return f.Message.RPC.IsView() }
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

func (f *RPCField) Payload() (dart.Qualifier, error) {
	return f.RPCS().Payload(f.Pgs)
}

func (f *RPCField) SrcId() (dart.Qualifier, error)   { return f.RPCS().SrcId(f.Pgs) }
func (f *RPCField) SrcResp() (dart.Qualifier, error) { return f.RPCS().SrcResp(f.File.Dart, f.Pgs) }
func (f *RPCField) GetResp() (dart.Qualifier, error) { return f.RPCS().GetResp(f.Pgs) }
func (f *RPCField) ListAddOrSaveReq() (dart.Qualifier, error) {
	return f.RPCS().ListAddOrSaveReq(f.Pgs)
}
func (f *RPCField) AddResp() (dart.Qualifier, error) { return f.RPCS().AddResp(f.File.Dart, f.Pgs) }
func (f *RPCField) SrcIds() (dart.Qualifier, error)  { return f.RPCS().SrcIds(f.File.Dart, f.Pgs) }
func (f *RPCField) DstResp() (dart.Qualifier, error) { return f.RPCS().DstResp(f.File.Dart, f.Pgs) }
func (f *RPCField) SelectResp() (dart.Qualifier, error) {
	return f.RPCS().SelectResp(f.File.Dart, f.Pgs)
}

func (f *RPCField) BackendErrorName() (dart.Qualifier, error) {
	return f.Message.RPC.BackendErrorName()
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
func (f *RPCField) AddMany() bool {
	return f.IsSelectMany() && f.EntryExt.GetSelectMany().Dst.GetAddMany()
}
func (f *RPCField) Save() bool { return f.IsList() && f.EntryExt.GetList().Edit }
