// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/units/units.proto

package units // import "github.com/empirefox/protoc-gen-dart-ext/pkg/units"

/*
https://physics.nist.gov/cuu/Units/units.html
*/

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PrefixType int32

const (
	PrefixType_symbol PrefixType = 0
	PrefixType_title  PrefixType = 1
)

var PrefixType_name = map[int32]string{
	0: "symbol",
	1: "title",
}
var PrefixType_value = map[string]int32{
	"symbol": 0,
	"title":  1,
}

func (x PrefixType) String() string {
	return proto.EnumName(PrefixType_name, int32(x))
}
func (PrefixType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_units_da54191d495fc701, []int{0}
}

type Show_AtomType int32

const (
	Show_symbol Show_AtomType = 0
	// plural Form.one
	Show_one Show_AtomType = 1
	// plural Form.other
	Show_other Show_AtomType = 2
	// parse Form from plural tool
	Show_parse Show_AtomType = 3
)

var Show_AtomType_name = map[int32]string{
	0: "symbol",
	1: "one",
	2: "other",
	3: "parse",
}
var Show_AtomType_value = map[string]int32{
	"symbol": 0,
	"one":    1,
	"other":  2,
	"parse":  3,
}

func (x Show_AtomType) String() string {
	return proto.EnumName(Show_AtomType_name, int32(x))
}
func (Show_AtomType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_units_da54191d495fc701, []int{1, 0}
}

type Unit struct {
	Show                 *Show    `protobuf:"bytes,1,opt,name=show,proto3" json:"show,omitempty"`
	Ordinal              bool     `protobuf:"varint,2,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Dots                 []*Cell  `protobuf:"bytes,3,rep,name=dots,proto3" json:"dots,omitempty"`
	Per                  []*Cell  `protobuf:"bytes,4,rep,name=per,proto3" json:"per,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Unit) Reset()         { *m = Unit{} }
func (m *Unit) String() string { return proto.CompactTextString(m) }
func (*Unit) ProtoMessage()    {}
func (*Unit) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_da54191d495fc701, []int{0}
}
func (m *Unit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Unit.Unmarshal(m, b)
}
func (m *Unit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Unit.Marshal(b, m, deterministic)
}
func (dst *Unit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Unit.Merge(dst, src)
}
func (m *Unit) XXX_Size() int {
	return xxx_messageInfo_Unit.Size(m)
}
func (m *Unit) XXX_DiscardUnknown() {
	xxx_messageInfo_Unit.DiscardUnknown(m)
}

var xxx_messageInfo_Unit proto.InternalMessageInfo

func (m *Unit) GetShow() *Show {
	if m != nil {
		return m.Show
	}
	return nil
}

func (m *Unit) GetOrdinal() bool {
	if m != nil {
		return m.Ordinal
	}
	return false
}

func (m *Unit) GetDots() []*Cell {
	if m != nil {
		return m.Dots
	}
	return nil
}

func (m *Unit) GetPer() []*Cell {
	if m != nil {
		return m.Per
	}
	return nil
}

type Show struct {
	Off                  bool          `protobuf:"varint,1,opt,name=off,proto3" json:"off,omitempty"`
	Prefix               PrefixType    `protobuf:"varint,2,opt,name=prefix,proto3,enum=units.PrefixType" json:"prefix,omitempty"`
	Atom                 Show_AtomType `protobuf:"varint,3,opt,name=atom,proto3,enum=units.Show_AtomType" json:"atom,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Show) Reset()         { *m = Show{} }
func (m *Show) String() string { return proto.CompactTextString(m) }
func (*Show) ProtoMessage()    {}
func (*Show) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_da54191d495fc701, []int{1}
}
func (m *Show) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Show.Unmarshal(m, b)
}
func (m *Show) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Show.Marshal(b, m, deterministic)
}
func (dst *Show) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Show.Merge(dst, src)
}
func (m *Show) XXX_Size() int {
	return xxx_messageInfo_Show.Size(m)
}
func (m *Show) XXX_DiscardUnknown() {
	xxx_messageInfo_Show.DiscardUnknown(m)
}

var xxx_messageInfo_Show proto.InternalMessageInfo

func (m *Show) GetOff() bool {
	if m != nil {
		return m.Off
	}
	return false
}

func (m *Show) GetPrefix() PrefixType {
	if m != nil {
		return m.Prefix
	}
	return PrefixType_symbol
}

func (m *Show) GetAtom() Show_AtomType {
	if m != nil {
		return m.Atom
	}
	return Show_symbol
}

type Cell struct {
	Exponent int32    `protobuf:"varint,1,opt,name=exponent,proto3" json:"exponent,omitempty"`
	Prefix   PrefixV1 `protobuf:"varint,2,opt,name=prefix,proto3,enum=units.PrefixV1" json:"prefix,omitempty"`
	// Types that are valid to be assigned to Unit:
	//	*Cell_Atom
	//	*Cell_Symbol
	Unit                 isCell_Unit `protobuf_oneof:"unit"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Cell) Reset()         { *m = Cell{} }
func (m *Cell) String() string { return proto.CompactTextString(m) }
func (*Cell) ProtoMessage()    {}
func (*Cell) Descriptor() ([]byte, []int) {
	return fileDescriptor_units_da54191d495fc701, []int{2}
}
func (m *Cell) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cell.Unmarshal(m, b)
}
func (m *Cell) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cell.Marshal(b, m, deterministic)
}
func (dst *Cell) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cell.Merge(dst, src)
}
func (m *Cell) XXX_Size() int {
	return xxx_messageInfo_Cell.Size(m)
}
func (m *Cell) XXX_DiscardUnknown() {
	xxx_messageInfo_Cell.DiscardUnknown(m)
}

var xxx_messageInfo_Cell proto.InternalMessageInfo

type isCell_Unit interface {
	isCell_Unit()
}

type Cell_Atom struct {
	Atom AtomV1 `protobuf:"varint,3,opt,name=atom,proto3,enum=units.AtomV1,oneof"`
}
type Cell_Symbol struct {
	Symbol string `protobuf:"bytes,4,opt,name=symbol,proto3,oneof"`
}

func (*Cell_Atom) isCell_Unit()   {}
func (*Cell_Symbol) isCell_Unit() {}

func (m *Cell) GetUnit() isCell_Unit {
	if m != nil {
		return m.Unit
	}
	return nil
}

func (m *Cell) GetExponent() int32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *Cell) GetPrefix() PrefixV1 {
	if m != nil {
		return m.Prefix
	}
	return PrefixV1_noPrefix
}

func (m *Cell) GetAtom() AtomV1 {
	if x, ok := m.GetUnit().(*Cell_Atom); ok {
		return x.Atom
	}
	return AtomV1_noAtom
}

func (m *Cell) GetSymbol() string {
	if x, ok := m.GetUnit().(*Cell_Symbol); ok {
		return x.Symbol
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Cell) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Cell_OneofMarshaler, _Cell_OneofUnmarshaler, _Cell_OneofSizer, []interface{}{
		(*Cell_Atom)(nil),
		(*Cell_Symbol)(nil),
	}
}

func _Cell_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Cell)
	// unit
	switch x := m.Unit.(type) {
	case *Cell_Atom:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Atom))
	case *Cell_Symbol:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Symbol)
	case nil:
	default:
		return fmt.Errorf("Cell.Unit has unexpected type %T", x)
	}
	return nil
}

func _Cell_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Cell)
	switch tag {
	case 3: // unit.atom
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Unit = &Cell_Atom{AtomV1(x)}
		return true, err
	case 4: // unit.symbol
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Unit = &Cell_Symbol{x}
		return true, err
	default:
		return false, nil
	}
}

func _Cell_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Cell)
	// unit
	switch x := m.Unit.(type) {
	case *Cell_Atom:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Atom))
	case *Cell_Symbol:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Symbol)))
		n += len(x.Symbol)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

var E_Unit = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Unit)(nil),
	Field:         919112,
	Name:          "units.unit",
	Tag:           "bytes,919112,opt,name=unit",
	Filename:      "protos/units/units.proto",
}

func init() {
	proto.RegisterType((*Unit)(nil), "units.Unit")
	proto.RegisterType((*Show)(nil), "units.Show")
	proto.RegisterType((*Cell)(nil), "units.Cell")
	proto.RegisterEnum("units.PrefixType", PrefixType_name, PrefixType_value)
	proto.RegisterEnum("units.Show_AtomType", Show_AtomType_name, Show_AtomType_value)
	proto.RegisterExtension(E_Unit)
}

func init() { proto.RegisterFile("protos/units/units.proto", fileDescriptor_units_da54191d495fc701) }

var fileDescriptor_units_da54191d495fc701 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x92, 0xcf, 0xae, 0x93, 0x40,
	0x14, 0xc6, 0x2f, 0x17, 0xca, 0xe5, 0x9e, 0x1b, 0x15, 0x27, 0x26, 0x62, 0x93, 0xc6, 0x86, 0x2e,
	0xac, 0x26, 0x85, 0xb4, 0x6a, 0x62, 0x5c, 0x69, 0x4d, 0x8c, 0x3b, 0x0d, 0x6a, 0x17, 0xee, 0xa0,
	0x0c, 0x74, 0x22, 0x30, 0x93, 0x61, 0x1a, 0xe9, 0xd2, 0xbd, 0x2b, 0x9f, 0xc1, 0x87, 0xf1, 0xb1,
	0x9c, 0x3f, 0xb4, 0x16, 0xdd, 0x34, 0x33, 0xdf, 0xf7, 0x9d, 0x9e, 0xdf, 0x19, 0x0e, 0x04, 0x8c,
	0x53, 0x41, 0xdb, 0x78, 0xdf, 0x10, 0xd1, 0xff, 0x46, 0x5a, 0x42, 0x23, 0x7d, 0x19, 0x4f, 0x4b,
	0x4a, 0xcb, 0x0a, 0xc7, 0x5a, 0xcc, 0xf6, 0x45, 0x9c, 0xe3, 0x76, 0xcb, 0x09, 0x13, 0x94, 0x9b,
	0xe0, 0xf8, 0xc1, 0xe0, 0x2f, 0x18, 0xc7, 0x05, 0xe9, 0x7a, 0xeb, 0xfe, 0xc0, 0x4a, 0x05, 0xad,
	0x8d, 0x11, 0x7e, 0xb7, 0xc0, 0xf9, 0x2c, 0x45, 0xf4, 0x10, 0x9c, 0x76, 0x47, 0xbf, 0x05, 0xd6,
	0xd4, 0x9a, 0xdf, 0xac, 0x6e, 0x22, 0x43, 0xf0, 0x51, 0x4a, 0x89, 0x36, 0x50, 0x00, 0x57, 0x94,
	0xe7, 0xa4, 0x49, 0xab, 0xe0, 0x52, 0x66, 0xbc, 0xe4, 0x78, 0x55, 0xa5, 0x39, 0x15, 0x6d, 0x60,
	0x4f, 0xed, 0xb3, 0xd2, 0x37, 0xb8, 0xaa, 0x12, 0x6d, 0xa0, 0x09, 0xd8, 0x0c, 0xf3, 0xc0, 0xf9,
	0xdf, 0x57, 0x7a, 0xf8, 0x4b, 0x32, 0xa8, 0x46, 0xc8, 0x07, 0x9b, 0x16, 0x85, 0x46, 0xf0, 0x12,
	0x75, 0x44, 0x8f, 0xc1, 0x35, 0x73, 0xe8, 0x9e, 0xb7, 0x57, 0x77, 0xfb, 0xe2, 0x0f, 0x5a, 0xfc,
	0x74, 0x60, 0x38, 0xe9, 0x03, 0x68, 0x0e, 0x8e, 0x9a, 0x4b, 0x52, 0xa8, 0xe0, 0xbd, 0xb3, 0x01,
	0xa2, 0xd7, 0x52, 0xd7, 0x59, 0x9d, 0x08, 0x9f, 0x83, 0x77, 0x54, 0x10, 0x80, 0xdb, 0x1e, 0xea,
	0x8c, 0x56, 0xfe, 0x05, 0xba, 0x92, 0xed, 0x1b, 0xec, 0x5b, 0xe8, 0x1a, 0x46, 0x54, 0xec, 0x30,
	0xf7, 0x2f, 0xd5, 0x91, 0xa5, 0xbc, 0xc5, 0xbe, 0x1d, 0xfe, 0x94, 0x98, 0x0a, 0x1a, 0x8d, 0xc1,
	0xc3, 0x1d, 0x93, 0xd1, 0x46, 0x68, 0xd6, 0x51, 0x72, 0xba, 0xa3, 0x47, 0xff, 0x00, 0xdf, 0x19,
	0x00, 0x6f, 0x96, 0x27, 0xdc, 0xd9, 0x00, 0xf7, 0x56, 0x1f, 0x53, 0x5c, 0x9b, 0xe5, 0xbb, 0x0b,
	0x43, 0x2a, 0xdf, 0xbc, 0xa7, 0x93, 0x6f, 0x67, 0xcd, 0xaf, 0xa5, 0xde, 0xdf, 0xd7, 0x2e, 0x38,
	0xaa, 0xe2, 0xc9, 0x0c, 0xe0, 0xef, 0x5b, 0x0c, 0xa6, 0x91, 0xe4, 0x82, 0x88, 0x4a, 0xce, 0xf3,
	0xf2, 0x95, 0x09, 0xa3, 0x49, 0x64, 0x76, 0x28, 0x3a, 0xee, 0x50, 0xf4, 0x96, 0xe0, 0x2a, 0x7f,
	0xcf, 0x04, 0xa1, 0x4d, 0x1b, 0xfc, 0xfe, 0xf1, 0x62, 0xf0, 0xf1, 0xd5, 0x5e, 0x24, 0xba, 0x72,
	0xfd, 0xec, 0xcb, 0xaa, 0x24, 0x62, 0xb7, 0xcf, 0xa2, 0x2d, 0xad, 0x63, 0x5c, 0x33, 0x22, 0x9b,
	0xd2, 0xce, 0x2c, 0xe3, 0x76, 0x51, 0xe2, 0x66, 0x91, 0xa7, 0x5c, 0x2c, 0x70, 0x27, 0x62, 0xf6,
	0xb5, 0x34, 0x7b, 0x96, 0xb9, 0xda, 0x7e, 0xfa, 0x27, 0x00, 0x00, 0xff, 0xff, 0xa9, 0xdc, 0x62,
	0x65, 0xdc, 0x02, 0x00, 0x00,
}
