// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pgde/zero/zero.proto

package zero

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Zero struct {
	// Types that are valid to be assigned to Type:
	//	*Zero_Float
	//	*Zero_Double
	//	*Zero_Int32
	//	*Zero_Int64
	//	*Zero_Uint32
	//	*Zero_Uint64
	//	*Zero_Sint32
	//	*Zero_Sint64
	//	*Zero_Fixed32
	//	*Zero_Fixed64
	//	*Zero_Sfixed32
	//	*Zero_Sfixed64
	//	*Zero_Bool
	//	*Zero_String_
	//	*Zero_Bytes
	//	*Zero_Enum
	//	*Zero_Message
	//	*Zero_Repeated
	//	*Zero_Map
	//	*Zero_Any
	//	*Zero_Duration
	//	*Zero_Time
	//	*Zero_NoChange
	Type                 isZero_Type `protobuf_oneof:"type"`
	OnCreate             *Affect     `protobuf:"bytes,30,opt,name=onCreate,proto3" json:"onCreate,omitempty"`
	OnLoad               *Affect     `protobuf:"bytes,31,opt,name=onLoad,proto3" json:"onLoad,omitempty"`
	OnSave               *Affect     `protobuf:"bytes,32,opt,name=onSave,proto3" json:"onSave,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Zero) Reset()         { *m = Zero{} }
func (m *Zero) String() string { return proto.CompactTextString(m) }
func (*Zero) ProtoMessage()    {}
func (*Zero) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{0}
}

func (m *Zero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Zero.Unmarshal(m, b)
}
func (m *Zero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Zero.Marshal(b, m, deterministic)
}
func (m *Zero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Zero.Merge(m, src)
}
func (m *Zero) XXX_Size() int {
	return xxx_messageInfo_Zero.Size(m)
}
func (m *Zero) XXX_DiscardUnknown() {
	xxx_messageInfo_Zero.DiscardUnknown(m)
}

var xxx_messageInfo_Zero proto.InternalMessageInfo

type isZero_Type interface {
	isZero_Type()
}

type Zero_Float struct {
	Float float32 `protobuf:"fixed32,1,opt,name=float,proto3,oneof"`
}

type Zero_Double struct {
	Double float64 `protobuf:"fixed64,2,opt,name=double,proto3,oneof"`
}

type Zero_Int32 struct {
	Int32 int32 `protobuf:"varint,3,opt,name=int32,proto3,oneof"`
}

type Zero_Int64 struct {
	Int64 int64 `protobuf:"varint,4,opt,name=int64,proto3,oneof"`
}

type Zero_Uint32 struct {
	Uint32 uint32 `protobuf:"varint,5,opt,name=uint32,proto3,oneof"`
}

type Zero_Uint64 struct {
	Uint64 uint64 `protobuf:"varint,6,opt,name=uint64,proto3,oneof"`
}

type Zero_Sint32 struct {
	Sint32 int32 `protobuf:"zigzag32,7,opt,name=sint32,proto3,oneof"`
}

type Zero_Sint64 struct {
	Sint64 int64 `protobuf:"zigzag64,8,opt,name=sint64,proto3,oneof"`
}

type Zero_Fixed32 struct {
	Fixed32 uint32 `protobuf:"fixed32,9,opt,name=fixed32,proto3,oneof"`
}

type Zero_Fixed64 struct {
	Fixed64 uint64 `protobuf:"fixed64,10,opt,name=fixed64,proto3,oneof"`
}

type Zero_Sfixed32 struct {
	Sfixed32 int32 `protobuf:"fixed32,11,opt,name=sfixed32,proto3,oneof"`
}

type Zero_Sfixed64 struct {
	Sfixed64 int64 `protobuf:"fixed64,12,opt,name=sfixed64,proto3,oneof"`
}

type Zero_Bool struct {
	Bool bool `protobuf:"varint,13,opt,name=bool,proto3,oneof"`
}

type Zero_String_ struct {
	String_ string `protobuf:"bytes,14,opt,name=string,proto3,oneof"`
}

type Zero_Bytes struct {
	Bytes []byte `protobuf:"bytes,15,opt,name=bytes,proto3,oneof"`
}

type Zero_Enum struct {
	Enum int32 `protobuf:"varint,16,opt,name=enum,proto3,oneof"`
}

type Zero_Message struct {
	Message *Message `protobuf:"bytes,17,opt,name=message,proto3,oneof"`
}

type Zero_Repeated struct {
	Repeated *Repeated `protobuf:"bytes,18,opt,name=repeated,proto3,oneof"`
}

type Zero_Map struct {
	Map *Map `protobuf:"bytes,19,opt,name=map,proto3,oneof"`
}

type Zero_Any struct {
	Any *any.Any `protobuf:"bytes,20,opt,name=any,proto3,oneof"`
}

type Zero_Duration struct {
	Duration *Duration `protobuf:"bytes,21,opt,name=duration,proto3,oneof"`
}

type Zero_Time struct {
	Time *Time `protobuf:"bytes,22,opt,name=time,proto3,oneof"`
}

type Zero_NoChange struct {
	NoChange bool `protobuf:"varint,23,opt,name=noChange,proto3,oneof"`
}

func (*Zero_Float) isZero_Type() {}

func (*Zero_Double) isZero_Type() {}

func (*Zero_Int32) isZero_Type() {}

func (*Zero_Int64) isZero_Type() {}

func (*Zero_Uint32) isZero_Type() {}

func (*Zero_Uint64) isZero_Type() {}

func (*Zero_Sint32) isZero_Type() {}

func (*Zero_Sint64) isZero_Type() {}

func (*Zero_Fixed32) isZero_Type() {}

func (*Zero_Fixed64) isZero_Type() {}

func (*Zero_Sfixed32) isZero_Type() {}

func (*Zero_Sfixed64) isZero_Type() {}

func (*Zero_Bool) isZero_Type() {}

func (*Zero_String_) isZero_Type() {}

func (*Zero_Bytes) isZero_Type() {}

func (*Zero_Enum) isZero_Type() {}

func (*Zero_Message) isZero_Type() {}

func (*Zero_Repeated) isZero_Type() {}

func (*Zero_Map) isZero_Type() {}

func (*Zero_Any) isZero_Type() {}

func (*Zero_Duration) isZero_Type() {}

func (*Zero_Time) isZero_Type() {}

func (*Zero_NoChange) isZero_Type() {}

func (m *Zero) GetType() isZero_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *Zero) GetFloat() float32 {
	if x, ok := m.GetType().(*Zero_Float); ok {
		return x.Float
	}
	return 0
}

func (m *Zero) GetDouble() float64 {
	if x, ok := m.GetType().(*Zero_Double); ok {
		return x.Double
	}
	return 0
}

func (m *Zero) GetInt32() int32 {
	if x, ok := m.GetType().(*Zero_Int32); ok {
		return x.Int32
	}
	return 0
}

func (m *Zero) GetInt64() int64 {
	if x, ok := m.GetType().(*Zero_Int64); ok {
		return x.Int64
	}
	return 0
}

func (m *Zero) GetUint32() uint32 {
	if x, ok := m.GetType().(*Zero_Uint32); ok {
		return x.Uint32
	}
	return 0
}

func (m *Zero) GetUint64() uint64 {
	if x, ok := m.GetType().(*Zero_Uint64); ok {
		return x.Uint64
	}
	return 0
}

func (m *Zero) GetSint32() int32 {
	if x, ok := m.GetType().(*Zero_Sint32); ok {
		return x.Sint32
	}
	return 0
}

func (m *Zero) GetSint64() int64 {
	if x, ok := m.GetType().(*Zero_Sint64); ok {
		return x.Sint64
	}
	return 0
}

func (m *Zero) GetFixed32() uint32 {
	if x, ok := m.GetType().(*Zero_Fixed32); ok {
		return x.Fixed32
	}
	return 0
}

func (m *Zero) GetFixed64() uint64 {
	if x, ok := m.GetType().(*Zero_Fixed64); ok {
		return x.Fixed64
	}
	return 0
}

func (m *Zero) GetSfixed32() int32 {
	if x, ok := m.GetType().(*Zero_Sfixed32); ok {
		return x.Sfixed32
	}
	return 0
}

func (m *Zero) GetSfixed64() int64 {
	if x, ok := m.GetType().(*Zero_Sfixed64); ok {
		return x.Sfixed64
	}
	return 0
}

func (m *Zero) GetBool() bool {
	if x, ok := m.GetType().(*Zero_Bool); ok {
		return x.Bool
	}
	return false
}

func (m *Zero) GetString_() string {
	if x, ok := m.GetType().(*Zero_String_); ok {
		return x.String_
	}
	return ""
}

func (m *Zero) GetBytes() []byte {
	if x, ok := m.GetType().(*Zero_Bytes); ok {
		return x.Bytes
	}
	return nil
}

func (m *Zero) GetEnum() int32 {
	if x, ok := m.GetType().(*Zero_Enum); ok {
		return x.Enum
	}
	return 0
}

func (m *Zero) GetMessage() *Message {
	if x, ok := m.GetType().(*Zero_Message); ok {
		return x.Message
	}
	return nil
}

func (m *Zero) GetRepeated() *Repeated {
	if x, ok := m.GetType().(*Zero_Repeated); ok {
		return x.Repeated
	}
	return nil
}

func (m *Zero) GetMap() *Map {
	if x, ok := m.GetType().(*Zero_Map); ok {
		return x.Map
	}
	return nil
}

func (m *Zero) GetAny() *any.Any {
	if x, ok := m.GetType().(*Zero_Any); ok {
		return x.Any
	}
	return nil
}

func (m *Zero) GetDuration() *Duration {
	if x, ok := m.GetType().(*Zero_Duration); ok {
		return x.Duration
	}
	return nil
}

func (m *Zero) GetTime() *Time {
	if x, ok := m.GetType().(*Zero_Time); ok {
		return x.Time
	}
	return nil
}

func (m *Zero) GetNoChange() bool {
	if x, ok := m.GetType().(*Zero_NoChange); ok {
		return x.NoChange
	}
	return false
}

func (m *Zero) GetOnCreate() *Affect {
	if m != nil {
		return m.OnCreate
	}
	return nil
}

func (m *Zero) GetOnLoad() *Affect {
	if m != nil {
		return m.OnLoad
	}
	return nil
}

func (m *Zero) GetOnSave() *Affect {
	if m != nil {
		return m.OnSave
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Zero) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Zero_Float)(nil),
		(*Zero_Double)(nil),
		(*Zero_Int32)(nil),
		(*Zero_Int64)(nil),
		(*Zero_Uint32)(nil),
		(*Zero_Uint64)(nil),
		(*Zero_Sint32)(nil),
		(*Zero_Sint64)(nil),
		(*Zero_Fixed32)(nil),
		(*Zero_Fixed64)(nil),
		(*Zero_Sfixed32)(nil),
		(*Zero_Sfixed64)(nil),
		(*Zero_Bool)(nil),
		(*Zero_String_)(nil),
		(*Zero_Bytes)(nil),
		(*Zero_Enum)(nil),
		(*Zero_Message)(nil),
		(*Zero_Repeated)(nil),
		(*Zero_Map)(nil),
		(*Zero_Any)(nil),
		(*Zero_Duration)(nil),
		(*Zero_Time)(nil),
		(*Zero_NoChange)(nil),
	}
}

type Affect struct {
	// Types that are valid to be assigned to On:
	//	*Affect_All
	//	*Affect_None
	//	*Affect_Plan
	On                   isAffect_On `protobuf_oneof:"on"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Affect) Reset()         { *m = Affect{} }
func (m *Affect) String() string { return proto.CompactTextString(m) }
func (*Affect) ProtoMessage()    {}
func (*Affect) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{1}
}

func (m *Affect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Affect.Unmarshal(m, b)
}
func (m *Affect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Affect.Marshal(b, m, deterministic)
}
func (m *Affect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Affect.Merge(m, src)
}
func (m *Affect) XXX_Size() int {
	return xxx_messageInfo_Affect.Size(m)
}
func (m *Affect) XXX_DiscardUnknown() {
	xxx_messageInfo_Affect.DiscardUnknown(m)
}

var xxx_messageInfo_Affect proto.InternalMessageInfo

type isAffect_On interface {
	isAffect_On()
}

type Affect_All struct {
	All bool `protobuf:"varint,1,opt,name=all,proto3,oneof"`
}

type Affect_None struct {
	None bool `protobuf:"varint,2,opt,name=none,proto3,oneof"`
}

type Affect_Plan struct {
	Plan string `protobuf:"bytes,3,opt,name=plan,proto3,oneof"`
}

func (*Affect_All) isAffect_On() {}

func (*Affect_None) isAffect_On() {}

func (*Affect_Plan) isAffect_On() {}

func (m *Affect) GetOn() isAffect_On {
	if m != nil {
		return m.On
	}
	return nil
}

func (m *Affect) GetAll() bool {
	if x, ok := m.GetOn().(*Affect_All); ok {
		return x.All
	}
	return false
}

func (m *Affect) GetNone() bool {
	if x, ok := m.GetOn().(*Affect_None); ok {
		return x.None
	}
	return false
}

func (m *Affect) GetPlan() string {
	if x, ok := m.GetOn().(*Affect_Plan); ok {
		return x.Plan
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Affect) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Affect_All)(nil),
		(*Affect_None)(nil),
		(*Affect_Plan)(nil),
	}
}

type Message struct {
	Json                 []byte   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	Skip                 bool     `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{2}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

func (m *Message) GetSkip() bool {
	if m != nil {
		return m.Skip
	}
	return false
}

type Repeated struct {
	Json                 []byte   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	Items                *Zero    `protobuf:"bytes,2,opt,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repeated) Reset()         { *m = Repeated{} }
func (m *Repeated) String() string { return proto.CompactTextString(m) }
func (*Repeated) ProtoMessage()    {}
func (*Repeated) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{3}
}

func (m *Repeated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repeated.Unmarshal(m, b)
}
func (m *Repeated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repeated.Marshal(b, m, deterministic)
}
func (m *Repeated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repeated.Merge(m, src)
}
func (m *Repeated) XXX_Size() int {
	return xxx_messageInfo_Repeated.Size(m)
}
func (m *Repeated) XXX_DiscardUnknown() {
	xxx_messageInfo_Repeated.DiscardUnknown(m)
}

var xxx_messageInfo_Repeated proto.InternalMessageInfo

func (m *Repeated) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

func (m *Repeated) GetItems() *Zero {
	if m != nil {
		return m.Items
	}
	return nil
}

type Map struct {
	Json                 []byte   `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
	Values               *Zero    `protobuf:"bytes,2,opt,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}
func (*Map) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{4}
}

func (m *Map) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Map.Unmarshal(m, b)
}
func (m *Map) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Map.Marshal(b, m, deterministic)
}
func (m *Map) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Map.Merge(m, src)
}
func (m *Map) XXX_Size() int {
	return xxx_messageInfo_Map.Size(m)
}
func (m *Map) XXX_DiscardUnknown() {
	xxx_messageInfo_Map.DiscardUnknown(m)
}

var xxx_messageInfo_Map proto.InternalMessageInfo

func (m *Map) GetJson() []byte {
	if m != nil {
		return m.Json
	}
	return nil
}

func (m *Map) GetValues() *Zero {
	if m != nil {
		return m.Values
	}
	return nil
}

// can be set to any int field, and can set format to time, default is second
type Duration struct {
	// Types that are valid to be assigned to Is:
	//	*Duration_Wkt
	//	*Duration_Parse
	Is                   isDuration_Is `protobuf_oneof:"is"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Duration) Reset()         { *m = Duration{} }
func (m *Duration) String() string { return proto.CompactTextString(m) }
func (*Duration) ProtoMessage()    {}
func (*Duration) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{5}
}

func (m *Duration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Duration.Unmarshal(m, b)
}
func (m *Duration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Duration.Marshal(b, m, deterministic)
}
func (m *Duration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Duration.Merge(m, src)
}
func (m *Duration) XXX_Size() int {
	return xxx_messageInfo_Duration.Size(m)
}
func (m *Duration) XXX_DiscardUnknown() {
	xxx_messageInfo_Duration.DiscardUnknown(m)
}

var xxx_messageInfo_Duration proto.InternalMessageInfo

type isDuration_Is interface {
	isDuration_Is()
}

type Duration_Wkt struct {
	Wkt *duration.Duration `protobuf:"bytes,1,opt,name=wkt,proto3,oneof"`
}

type Duration_Parse struct {
	Parse string `protobuf:"bytes,2,opt,name=parse,proto3,oneof"`
}

func (*Duration_Wkt) isDuration_Is() {}

func (*Duration_Parse) isDuration_Is() {}

func (m *Duration) GetIs() isDuration_Is {
	if m != nil {
		return m.Is
	}
	return nil
}

func (m *Duration) GetWkt() *duration.Duration {
	if x, ok := m.GetIs().(*Duration_Wkt); ok {
		return x.Wkt
	}
	return nil
}

func (m *Duration) GetParse() string {
	if x, ok := m.GetIs().(*Duration_Parse); ok {
		return x.Parse
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Duration) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Duration_Wkt)(nil),
		(*Duration_Parse)(nil),
	}
}

// can be set to any int field, and can set format to time, default is second
type Time struct {
	// Types that are valid to be assigned to Is:
	//	*Time_Wkt
	//	*Time_Parse
	//	*Time_Now
	//	*Time_NowAdd
	//	*Time_NowAddParse
	Is                   isTime_Is `protobuf_oneof:"is"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Time) Reset()         { *m = Time{} }
func (m *Time) String() string { return proto.CompactTextString(m) }
func (*Time) ProtoMessage()    {}
func (*Time) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d50a5336198306e, []int{6}
}

func (m *Time) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Time.Unmarshal(m, b)
}
func (m *Time) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Time.Marshal(b, m, deterministic)
}
func (m *Time) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Time.Merge(m, src)
}
func (m *Time) XXX_Size() int {
	return xxx_messageInfo_Time.Size(m)
}
func (m *Time) XXX_DiscardUnknown() {
	xxx_messageInfo_Time.DiscardUnknown(m)
}

var xxx_messageInfo_Time proto.InternalMessageInfo

type isTime_Is interface {
	isTime_Is()
}

type Time_Wkt struct {
	Wkt *timestamp.Timestamp `protobuf:"bytes,1,opt,name=wkt,proto3,oneof"`
}

type Time_Parse struct {
	Parse string `protobuf:"bytes,2,opt,name=parse,proto3,oneof"`
}

type Time_Now struct {
	Now bool `protobuf:"varint,3,opt,name=now,proto3,oneof"`
}

type Time_NowAdd struct {
	NowAdd *duration.Duration `protobuf:"bytes,4,opt,name=nowAdd,proto3,oneof"`
}

type Time_NowAddParse struct {
	NowAddParse string `protobuf:"bytes,5,opt,name=nowAddParse,proto3,oneof"`
}

func (*Time_Wkt) isTime_Is() {}

func (*Time_Parse) isTime_Is() {}

func (*Time_Now) isTime_Is() {}

func (*Time_NowAdd) isTime_Is() {}

func (*Time_NowAddParse) isTime_Is() {}

func (m *Time) GetIs() isTime_Is {
	if m != nil {
		return m.Is
	}
	return nil
}

func (m *Time) GetWkt() *timestamp.Timestamp {
	if x, ok := m.GetIs().(*Time_Wkt); ok {
		return x.Wkt
	}
	return nil
}

func (m *Time) GetParse() string {
	if x, ok := m.GetIs().(*Time_Parse); ok {
		return x.Parse
	}
	return ""
}

func (m *Time) GetNow() bool {
	if x, ok := m.GetIs().(*Time_Now); ok {
		return x.Now
	}
	return false
}

func (m *Time) GetNowAdd() *duration.Duration {
	if x, ok := m.GetIs().(*Time_NowAdd); ok {
		return x.NowAdd
	}
	return nil
}

func (m *Time) GetNowAddParse() string {
	if x, ok := m.GetIs().(*Time_NowAddParse); ok {
		return x.NowAddParse
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Time) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Time_Wkt)(nil),
		(*Time_Parse)(nil),
		(*Time_Now)(nil),
		(*Time_NowAdd)(nil),
		(*Time_NowAddParse)(nil),
	}
}

var E_Disabled = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         919113,
	Name:          "pgde.zero.disabled",
	Tag:           "varint,919113,opt,name=disabled",
	Filename:      "pgde/zero/zero.proto",
}

var E_Default = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.OneofOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         919113,
	Name:          "pgde.zero.default",
	Tag:           "varint,919113,opt,name=default",
	Filename:      "pgde/zero/zero.proto",
}

var E_DefaultNotSet = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         919113,
	Name:          "pgde.zero.defaultNotSet",
	Tag:           "varint,919113,opt,name=defaultNotSet",
	Filename:      "pgde/zero/zero.proto",
}

var E_To = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Zero)(nil),
	Field:         919113,
	Name:          "pgde.zero.to",
	Tag:           "bytes,919113,opt,name=to",
	Filename:      "pgde/zero/zero.proto",
}

func init() {
	proto.RegisterType((*Zero)(nil), "pgde.zero.Zero")
	proto.RegisterType((*Affect)(nil), "pgde.zero.Affect")
	proto.RegisterType((*Message)(nil), "pgde.zero.Message")
	proto.RegisterType((*Repeated)(nil), "pgde.zero.Repeated")
	proto.RegisterType((*Map)(nil), "pgde.zero.Map")
	proto.RegisterType((*Duration)(nil), "pgde.zero.Duration")
	proto.RegisterType((*Time)(nil), "pgde.zero.Time")
	proto.RegisterExtension(E_Disabled)
	proto.RegisterExtension(E_Default)
	proto.RegisterExtension(E_DefaultNotSet)
	proto.RegisterExtension(E_To)
}

func init() { proto.RegisterFile("pgde/zero/zero.proto", fileDescriptor_1d50a5336198306e) }

var fileDescriptor_1d50a5336198306e = []byte{
	// 882 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0x36, 0xfd, 0x23, 0x2b, 0x4c, 0xd2, 0x34, 0x6c, 0xd6, 0xb1, 0x41, 0xd7, 0x10, 0x02, 0x8a,
	0x69, 0x17, 0x96, 0xd1, 0xc4, 0x10, 0xba, 0x00, 0x03, 0x96, 0xb4, 0x1d, 0x7c, 0xb1, 0xb6, 0x83,
	0x5a, 0x60, 0x40, 0xef, 0xe8, 0x88, 0x56, 0xb5, 0x48, 0xa4, 0x20, 0x51, 0x4d, 0xbc, 0x67, 0xd8,
	0x4b, 0xed, 0x72, 0x6f, 0xb0, 0xc7, 0x19, 0x0e, 0x29, 0xa9, 0x4e, 0x9c, 0xae, 0x37, 0x86, 0xce,
	0xf9, 0x7e, 0x48, 0x9e, 0x73, 0x48, 0xe3, 0x83, 0x22, 0x89, 0xc5, 0xf4, 0x4f, 0x51, 0x2a, 0xf3,
	0x13, 0x14, 0xa5, 0xd2, 0x8a, 0x6c, 0x41, 0x36, 0x80, 0xc4, 0x21, 0x4b, 0x94, 0x4a, 0x32, 0x31,
	0x35, 0xc0, 0xa2, 0x5e, 0x4e, 0x63, 0x51, 0x5d, 0x94, 0x69, 0xa1, 0x55, 0x69, 0xc9, 0x87, 0x8f,
	0x6e, 0x33, 0xb8, 0x5c, 0x35, 0xd0, 0x93, 0x0d, 0x71, 0x5d, 0x72, 0x9d, 0x2a, 0xd9, 0xe0, 0x47,
	0xb7, 0x71, 0x9d, 0xe6, 0xa2, 0xd2, 0x3c, 0x2f, 0x2c, 0xc1, 0xfb, 0xd7, 0xc1, 0xc3, 0x0f, 0xa2,
	0x54, 0xe4, 0x21, 0x1e, 0x2d, 0x33, 0xc5, 0x35, 0x45, 0x0c, 0xf9, 0xfd, 0x79, 0x2f, 0xb2, 0x21,
	0xa1, 0xd8, 0x89, 0x55, 0xbd, 0xc8, 0x04, 0xed, 0x33, 0xe4, 0xa3, 0x79, 0x2f, 0x6a, 0x62, 0x50,
	0xa4, 0x52, 0x9f, 0x1c, 0xd3, 0x01, 0x43, 0xfe, 0x08, 0x14, 0x26, 0x6c, 0xf2, 0xe1, 0x8c, 0x0e,
	0x19, 0xf2, 0x07, 0x4d, 0x3e, 0x9c, 0x81, 0x53, 0x6d, 0x05, 0x23, 0x86, 0xfc, 0x5d, 0x70, 0xb2,
	0x71, 0x8b, 0x84, 0x33, 0xea, 0x30, 0xe4, 0x0f, 0x5b, 0xc4, 0x6a, 0x2a, 0xab, 0x19, 0x33, 0xe4,
	0xef, 0x03, 0x52, 0x75, 0x9a, 0xca, 0x6a, 0x5c, 0x86, 0x7c, 0xd2, 0x22, 0xe1, 0x8c, 0x1c, 0xe2,
	0xf1, 0x32, 0xbd, 0x16, 0xf1, 0xc9, 0x31, 0xdd, 0x62, 0xc8, 0x1f, 0xcf, 0x7b, 0x51, 0x9b, 0xe8,
	0xb0, 0x70, 0x46, 0x31, 0x43, 0xbe, 0xd3, 0x61, 0xe1, 0x8c, 0x3c, 0xc6, 0x6e, 0xd5, 0x0a, 0xb7,
	0x19, 0xf2, 0xf7, 0xe6, 0xbd, 0xa8, 0xcb, 0x7c, 0x46, 0xc3, 0x19, 0xdd, 0x61, 0xc8, 0xbf, 0xff,
	0x19, 0x0d, 0x67, 0xe4, 0x00, 0x0f, 0x17, 0x4a, 0x65, 0x74, 0x97, 0x21, 0xdf, 0x9d, 0xf7, 0x22,
	0x13, 0x99, 0x3d, 0xea, 0x32, 0x95, 0x09, 0xbd, 0xc7, 0x90, 0xbf, 0x65, 0xf6, 0x68, 0x62, 0xa8,
	0xd1, 0x62, 0xa5, 0x45, 0x45, 0xf7, 0x18, 0xf2, 0x77, 0xa0, 0x46, 0x26, 0x04, 0x1f, 0x21, 0xeb,
	0x9c, 0xde, 0x6f, 0x4a, 0x6a, 0x22, 0x12, 0xe0, 0x71, 0x2e, 0xaa, 0x8a, 0x27, 0x82, 0xee, 0x33,
	0xe4, 0x6f, 0x1f, 0x93, 0xa0, 0x9b, 0x9f, 0xe0, 0xb5, 0x45, 0xe0, 0x24, 0x0d, 0x89, 0x3c, 0xc3,
	0x6e, 0x29, 0x0a, 0xc1, 0xb5, 0x88, 0x29, 0x31, 0x82, 0x07, 0x6b, 0x82, 0xa8, 0x81, 0xe0, 0x00,
	0x2d, 0x8d, 0x78, 0x78, 0x90, 0xf3, 0x82, 0x3e, 0x30, 0xec, 0x7b, 0xeb, 0xf6, 0xbc, 0x98, 0xf7,
	0x22, 0x00, 0x89, 0x8f, 0x07, 0x5c, 0xae, 0xe8, 0x81, 0xe1, 0x1c, 0x04, 0x76, 0xb4, 0x82, 0x76,
	0xb4, 0x82, 0x33, 0xb9, 0x02, 0x26, 0x97, 0x2b, 0xd8, 0x40, 0x3b, 0x88, 0xf4, 0x9b, 0x8d, 0x0d,
	0xbc, 0x6c, 0x20, 0xd8, 0x40, 0x4b, 0x23, 0x4f, 0xf1, 0x10, 0x66, 0x93, 0x3e, 0x34, 0xf4, 0xbd,
	0x35, 0xfa, 0xfb, 0x34, 0x87, 0xd3, 0x19, 0x18, 0xda, 0x20, 0xd5, 0x8b, 0x8f, 0x5c, 0x26, 0x82,
	0x7e, 0xdb, 0x14, 0xbb, 0xcb, 0x90, 0x09, 0x76, 0x95, 0x7c, 0x51, 0xc2, 0x91, 0xe8, 0x13, 0x63,
	0xb4, 0xbf, 0x66, 0x74, 0xb6, 0x5c, 0x8a, 0x0b, 0x1d, 0x75, 0x14, 0xf2, 0x03, 0x76, 0x94, 0xfc,
	0x55, 0xf1, 0x98, 0x1e, 0x7d, 0x89, 0xdc, 0x10, 0x2c, 0xf5, 0x1d, 0xff, 0x24, 0x28, 0xfb, 0x1f,
	0x2a, 0x10, 0xce, 0x1d, 0x3c, 0xd4, 0xab, 0x42, 0x78, 0x6f, 0xb0, 0x63, 0x11, 0x42, 0xf0, 0x80,
	0x67, 0x99, 0xb9, 0x59, 0xae, 0x29, 0x51, 0x96, 0x41, 0xa7, 0xa5, 0x92, 0xf6, 0x56, 0x99, 0x89,
	0x81, 0x08, 0xb2, 0x45, 0xc6, 0xa5, 0xb9, 0x52, 0x30, 0x2f, 0x26, 0x3a, 0x1f, 0xe2, 0xbe, 0x92,
	0xde, 0x33, 0x3c, 0x6e, 0x7a, 0x4d, 0x08, 0x1e, 0xfe, 0x51, 0x29, 0x69, 0x1c, 0x77, 0x22, 0xf3,
	0x0d, 0xb9, 0xea, 0x32, 0x2d, 0xac, 0x61, 0x64, 0xbe, 0xbd, 0x57, 0xd8, 0x6d, 0xbb, 0x7d, 0xa7,
	0xe6, 0x29, 0x1e, 0xa5, 0x5a, 0xe4, 0x95, 0x11, 0xdd, 0xac, 0x3a, 0x3c, 0x0a, 0x91, 0x45, 0xbd,
	0x73, 0x3c, 0x78, 0xcd, 0x8b, 0x3b, 0x1d, 0xbe, 0xc7, 0xce, 0x27, 0x9e, 0xd5, 0xe2, 0x8b, 0x16,
	0x0d, 0xec, 0xfd, 0x8e, 0xdd, 0xb6, 0xef, 0x64, 0x82, 0x07, 0x57, 0x97, 0xf6, 0xa5, 0xd9, 0x3e,
	0x7e, 0xb4, 0x31, 0x48, 0x6b, 0xf3, 0x01, 0x3c, 0xb8, 0x2c, 0x05, 0x2f, 0x2b, 0x5b, 0x2b, 0xa8,
	0x8a, 0x0d, 0xa1, 0x2c, 0x69, 0xe5, 0xfd, 0x8d, 0xf0, 0x10, 0x46, 0x84, 0x04, 0xeb, 0xae, 0x87,
	0x1b, 0xae, 0xef, 0xdb, 0x97, 0xef, 0x2b, 0xb6, 0xd0, 0x2d, 0xa9, 0xae, 0x4c, 0x0b, 0x4c, 0xb7,
	0xa4, 0xba, 0x22, 0x27, 0xd8, 0x91, 0xea, 0xea, 0x2c, 0x8e, 0xcd, 0xa3, 0xf6, 0x95, 0x4d, 0x37,
	0x54, 0xe2, 0xe1, 0x6d, 0xfb, 0xf5, 0x9b, 0x59, 0x66, 0xd4, 0x2c, 0xb3, 0x9e, 0xb4, 0x67, 0x38,
	0xfd, 0x09, 0xbb, 0x71, 0x5a, 0xf1, 0x45, 0x26, 0x62, 0x72, 0xb4, 0x61, 0xdd, 0x74, 0xfd, 0x6d,
	0x01, 0xfe, 0x15, 0xfd, 0xe7, 0xaf, 0xe7, 0xa6, 0xc5, 0x9d, 0xe4, 0xf4, 0x47, 0x3c, 0x8e, 0xc5,
	0x92, 0xd7, 0x99, 0x26, 0xdf, 0x6d, 0xa8, 0xdf, 0x4a, 0xa1, 0x96, 0x37, 0xb5, 0xa3, 0xa8, 0xe5,
	0x9f, 0xbe, 0xc4, 0xbb, 0xcd, 0xe7, 0x1b, 0xa5, 0xdf, 0x09, 0x4d, 0x1e, 0x6f, 0x18, 0xbc, 0x92,
	0x75, 0x7e, 0x7b, 0xed, 0x9b, 0xa2, 0xd3, 0x9f, 0x71, 0x5f, 0xab, 0x3b, 0xd6, 0xfe, 0x25, 0x15,
	0x59, 0x7c, 0x53, 0x7b, 0xc7, 0x88, 0xf4, 0xb5, 0x3a, 0x7f, 0xfe, 0x21, 0x4c, 0x52, 0xfd, 0xb1,
	0x5e, 0x04, 0x17, 0x2a, 0x9f, 0x8a, 0xbc, 0x48, 0x4b, 0xb1, 0x54, 0xd7, 0xf6, 0x8f, 0xeb, 0x62,
	0x92, 0x08, 0x39, 0x89, 0x79, 0xa9, 0x27, 0xe2, 0x5a, 0x4f, 0x8b, 0xcb, 0x64, 0xda, 0xfd, 0xab,
	0x2e, 0x1c, 0x43, 0x39, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x43, 0x7c, 0x9a, 0x45, 0x69, 0x07,
	0x00, 0x00,
}
