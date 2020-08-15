// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pgde/error/error.proto

package error

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type OperateError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *OperateError) Reset() {
	*x = OperateError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_error_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OperateError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OperateError) ProtoMessage() {}

func (x *OperateError) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_error_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OperateError.ProtoReflect.Descriptor instead.
func (*OperateError) Descriptor() ([]byte, []int) {
	return file_pgde_error_error_proto_rawDescGZIP(), []int{0}
}

func (x *OperateError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *OperateError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SubmitError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    []int32 `protobuf:"varint,1,rep,packed,name=path,proto3" json:"path,omitempty"`
	Code    int32   `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string  `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SubmitError) Reset() {
	*x = SubmitError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_error_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitError) ProtoMessage() {}

func (x *SubmitError) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_error_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitError.ProtoReflect.Descriptor instead.
func (*SubmitError) Descriptor() ([]byte, []int) {
	return file_pgde_error_error_proto_rawDescGZIP(), []int{1}
}

func (x *SubmitError) GetPath() []int32 {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *SubmitError) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *SubmitError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pgde_error_error_proto protoreflect.FileDescriptor

var file_pgde_error_error_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3c, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x4f, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x6d, 0x70, 0x69, 0x72, 0x65, 0x66, 0x6f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x61, 0x72, 0x74, 0x2d, 0x65, 0x78, 0x74, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x3b, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pgde_error_error_proto_rawDescOnce sync.Once
	file_pgde_error_error_proto_rawDescData = file_pgde_error_error_proto_rawDesc
)

func file_pgde_error_error_proto_rawDescGZIP() []byte {
	file_pgde_error_error_proto_rawDescOnce.Do(func() {
		file_pgde_error_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_pgde_error_error_proto_rawDescData)
	})
	return file_pgde_error_error_proto_rawDescData
}

var file_pgde_error_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pgde_error_error_proto_goTypes = []interface{}{
	(*OperateError)(nil), // 0: pgde.error.OperateError
	(*SubmitError)(nil),  // 1: pgde.error.SubmitError
}
var file_pgde_error_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pgde_error_error_proto_init() }
func file_pgde_error_error_proto_init() {
	if File_pgde_error_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pgde_error_error_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OperateError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pgde_error_error_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pgde_error_error_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pgde_error_error_proto_goTypes,
		DependencyIndexes: file_pgde_error_error_proto_depIdxs,
		MessageInfos:      file_pgde_error_error_proto_msgTypes,
	}.Build()
	File_pgde_error_error_proto = out.File
	file_pgde_error_error_proto_rawDesc = nil
	file_pgde_error_error_proto_goTypes = nil
	file_pgde_error_error_proto_depIdxs = nil
}
