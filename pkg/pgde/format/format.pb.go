// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pgde/format/format.proto

package format

import (
	units "github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/units"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

// Support field: `Datetime`, `Duration`, `[s,fixed]int64`(will mark it as
// time) or other `int` types. `int` values must be from utc.
//
// - `*Datetime[24H]` must set on `Datetime` or `int64` type.
// - `ofDay[24H]` must set on `Datetime` or `uint` types.
// - `duration` must set on `Duration` or any `int` types.
type TimeFormat_Builtin int32

const (
	// Formats the date using a medium-width format.
	//  Examples:
	// - US English: Wed, Sep 27
	// - Russian: ср, сент. 27
	TimeFormat_mediumDate TimeFormat_Builtin = 0
	// Formats day of week, month, day of month and year in a long-width format.
	// Examples:
	// - US English: Wednesday, September 27, 2017
	// - Russian: Среда, Сентябрь 27, 2017
	TimeFormat_fullDate TimeFormat_Builtin = 1
	// 3:00pm or 15:00pm
	// default parse the int value as minute
	TimeFormat_ofDay TimeFormat_Builtin = 2
	// 15:00pm
	TimeFormat_ofDay24H          TimeFormat_Builtin = 3
	TimeFormat_mediumDatetime    TimeFormat_Builtin = 4
	TimeFormat_mediumDatetime24H TimeFormat_Builtin = 5
	TimeFormat_fullDatetime      TimeFormat_Builtin = 6
	TimeFormat_fullDatetime24H   TimeFormat_Builtin = 7
	TimeFormat_duration          TimeFormat_Builtin = 8
)

// Enum value maps for TimeFormat_Builtin.
var (
	TimeFormat_Builtin_name = map[int32]string{
		0: "mediumDate",
		1: "fullDate",
		2: "ofDay",
		3: "ofDay24H",
		4: "mediumDatetime",
		5: "mediumDatetime24H",
		6: "fullDatetime",
		7: "fullDatetime24H",
		8: "duration",
	}
	TimeFormat_Builtin_value = map[string]int32{
		"mediumDate":        0,
		"fullDate":          1,
		"ofDay":             2,
		"ofDay24H":          3,
		"mediumDatetime":    4,
		"mediumDatetime24H": 5,
		"fullDatetime":      6,
		"fullDatetime24H":   7,
		"duration":          8,
	}
)

func (x TimeFormat_Builtin) Enum() *TimeFormat_Builtin {
	p := new(TimeFormat_Builtin)
	*p = x
	return p
}

func (x TimeFormat_Builtin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeFormat_Builtin) Descriptor() protoreflect.EnumDescriptor {
	return file_pgde_format_format_proto_enumTypes[0].Descriptor()
}

func (TimeFormat_Builtin) Type() protoreflect.EnumType {
	return &file_pgde_format_format_proto_enumTypes[0]
}

func (x TimeFormat_Builtin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeFormat_Builtin.Descriptor instead.
func (TimeFormat_Builtin) EnumDescriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{1, 0}
}

type TimeFormat_IntUnit int32

const (
	// invalid if field is of an int type
	TimeFormat_und TimeFormat_IntUnit = 0
	// Dart lost the nano part, eg. 1,100 becomes 1,000.
	TimeFormat_nanosecond  TimeFormat_IntUnit = 1
	TimeFormat_microsecond TimeFormat_IntUnit = 2
	TimeFormat_millisecond TimeFormat_IntUnit = 3
	TimeFormat_second      TimeFormat_IntUnit = 4
	TimeFormat_minute      TimeFormat_IntUnit = 5
	TimeFormat_hour        TimeFormat_IntUnit = 6
)

// Enum value maps for TimeFormat_IntUnit.
var (
	TimeFormat_IntUnit_name = map[int32]string{
		0: "und",
		1: "nanosecond",
		2: "microsecond",
		3: "millisecond",
		4: "second",
		5: "minute",
		6: "hour",
	}
	TimeFormat_IntUnit_value = map[string]int32{
		"und":         0,
		"nanosecond":  1,
		"microsecond": 2,
		"millisecond": 3,
		"second":      4,
		"minute":      5,
		"hour":        6,
	}
)

func (x TimeFormat_IntUnit) Enum() *TimeFormat_IntUnit {
	p := new(TimeFormat_IntUnit)
	*p = x
	return p
}

func (x TimeFormat_IntUnit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeFormat_IntUnit) Descriptor() protoreflect.EnumDescriptor {
	return file_pgde_format_format_proto_enumTypes[1].Descriptor()
}

func (TimeFormat_IntUnit) Type() protoreflect.EnumType {
	return &file_pgde_format_format_proto_enumTypes[1]
}

func (x TimeFormat_IntUnit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeFormat_IntUnit.Descriptor instead.
func (TimeFormat_IntUnit) EnumDescriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{1, 1}
}

type NumberFormat_Builtin int32

const (
	NumberFormat_decimal        NumberFormat_Builtin = 0
	NumberFormat_percent        NumberFormat_Builtin = 1
	NumberFormat_scientific     NumberFormat_Builtin = 2
	NumberFormat_currency       NumberFormat_Builtin = 3
	NumberFormat_currencySimple NumberFormat_Builtin = 4
)

// Enum value maps for NumberFormat_Builtin.
var (
	NumberFormat_Builtin_name = map[int32]string{
		0: "decimal",
		1: "percent",
		2: "scientific",
		3: "currency",
		4: "currencySimple",
	}
	NumberFormat_Builtin_value = map[string]int32{
		"decimal":        0,
		"percent":        1,
		"scientific":     2,
		"currency":       3,
		"currencySimple": 4,
	}
)

func (x NumberFormat_Builtin) Enum() *NumberFormat_Builtin {
	p := new(NumberFormat_Builtin)
	*p = x
	return p
}

func (x NumberFormat_Builtin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NumberFormat_Builtin) Descriptor() protoreflect.EnumDescriptor {
	return file_pgde_format_format_proto_enumTypes[2].Descriptor()
}

func (NumberFormat_Builtin) Type() protoreflect.EnumType {
	return &file_pgde_format_format_proto_enumTypes[2]
}

func (x NumberFormat_Builtin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NumberFormat_Builtin.Descriptor instead.
func (NumberFormat_Builtin) EnumDescriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{3, 0}
}

type BytesFormat_Builtin int32

const (
	BytesFormat_ip       BytesFormat_Builtin = 0
	BytesFormat_ipv4     BytesFormat_Builtin = 1
	BytesFormat_ipv6     BytesFormat_Builtin = 2
	BytesFormat_ipPort   BytesFormat_Builtin = 3
	BytesFormat_ipv4Port BytesFormat_Builtin = 4
	BytesFormat_ipv6Port BytesFormat_Builtin = 5
	BytesFormat_cidr     BytesFormat_Builtin = 6
	BytesFormat_cidrv4   BytesFormat_Builtin = 7
	BytesFormat_cidrv6   BytesFormat_Builtin = 8
)

// Enum value maps for BytesFormat_Builtin.
var (
	BytesFormat_Builtin_name = map[int32]string{
		0: "ip",
		1: "ipv4",
		2: "ipv6",
		3: "ipPort",
		4: "ipv4Port",
		5: "ipv6Port",
		6: "cidr",
		7: "cidrv4",
		8: "cidrv6",
	}
	BytesFormat_Builtin_value = map[string]int32{
		"ip":       0,
		"ipv4":     1,
		"ipv6":     2,
		"ipPort":   3,
		"ipv4Port": 4,
		"ipv6Port": 5,
		"cidr":     6,
		"cidrv4":   7,
		"cidrv6":   8,
	}
)

func (x BytesFormat_Builtin) Enum() *BytesFormat_Builtin {
	p := new(BytesFormat_Builtin)
	*p = x
	return p
}

func (x BytesFormat_Builtin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BytesFormat_Builtin) Descriptor() protoreflect.EnumDescriptor {
	return file_pgde_format_format_proto_enumTypes[3].Descriptor()
}

func (BytesFormat_Builtin) Type() protoreflect.EnumType {
	return &file_pgde_format_format_proto_enumTypes[3]
}

func (x BytesFormat_Builtin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BytesFormat_Builtin.Descriptor instead.
func (BytesFormat_Builtin) EnumDescriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{4, 0}
}

type Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*Format_Time
	//	*Format_Currency
	//	*Format_Number
	//	*Format_Bytes
	Type isFormat_Type `protobuf_oneof:"type"`
}

func (x *Format) Reset() {
	*x = Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_format_format_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Format) ProtoMessage() {}

func (x *Format) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_format_format_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Format.ProtoReflect.Descriptor instead.
func (*Format) Descriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{0}
}

func (m *Format) GetType() isFormat_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Format) GetTime() *TimeFormat {
	if x, ok := x.GetType().(*Format_Time); ok {
		return x.Time
	}
	return nil
}

func (x *Format) GetCurrency() *CurrencyFormat {
	if x, ok := x.GetType().(*Format_Currency); ok {
		return x.Currency
	}
	return nil
}

func (x *Format) GetNumber() *NumberFormat {
	if x, ok := x.GetType().(*Format_Number); ok {
		return x.Number
	}
	return nil
}

func (x *Format) GetBytes() *BytesFormat {
	if x, ok := x.GetType().(*Format_Bytes); ok {
		return x.Bytes
	}
	return nil
}

type isFormat_Type interface {
	isFormat_Type()
}

type Format_Time struct {
	Time *TimeFormat `protobuf:"bytes,1,opt,name=time,proto3,oneof"`
}

type Format_Currency struct {
	Currency *CurrencyFormat `protobuf:"bytes,2,opt,name=currency,proto3,oneof"`
}

type Format_Number struct {
	Number *NumberFormat `protobuf:"bytes,3,opt,name=number,proto3,oneof"`
}

type Format_Bytes struct {
	Bytes *BytesFormat `protobuf:"bytes,4,opt,name=bytes,proto3,oneof"`
}

func (*Format_Time) isFormat_Type() {}

func (*Format_Currency) isFormat_Type() {}

func (*Format_Number) isFormat_Type() {}

func (*Format_Bytes) isFormat_Type() {}

type TimeFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type TimeFormat_Builtin `protobuf:"varint,1,opt,name=type,proto3,enum=pgde.format.TimeFormat_Builtin" json:"type,omitempty"`
	// only valid on int types, and must set if type is int
	IntUnit TimeFormat_IntUnit `protobuf:"varint,2,opt,name=intUnit,proto3,enum=pgde.format.TimeFormat_IntUnit" json:"intUnit,omitempty"`
}

func (x *TimeFormat) Reset() {
	*x = TimeFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_format_format_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimeFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeFormat) ProtoMessage() {}

func (x *TimeFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_format_format_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeFormat.ProtoReflect.Descriptor instead.
func (*TimeFormat) Descriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{1}
}

func (x *TimeFormat) GetType() TimeFormat_Builtin {
	if x != nil {
		return x.Type
	}
	return TimeFormat_mediumDate
}

func (x *TimeFormat) GetIntUnit() TimeFormat_IntUnit {
	if x != nil {
		return x.IntUnit
	}
	return TimeFormat_und
}

type CurrencyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*CurrencyFormat_Code
	//	*CurrencyFormat_Symbol
	//	*CurrencyFormat_Name
	Type isCurrencyFormat_Type `protobuf_oneof:"type"`
	// only valid for int currency
	// move decimal point left when showing
	// move decimal point right back when saving
	// need promise the visible fractional digits when showing
	// showing: 1 cent => 0.01 dolar
	// saving: 0.01 dolar => 1 cent
	CentMode bool `protobuf:"varint,4,opt,name=centMode,proto3" json:"centMode,omitempty"`
}

func (x *CurrencyFormat) Reset() {
	*x = CurrencyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_format_format_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyFormat) ProtoMessage() {}

func (x *CurrencyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_format_format_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyFormat.ProtoReflect.Descriptor instead.
func (*CurrencyFormat) Descriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{2}
}

func (m *CurrencyFormat) GetType() isCurrencyFormat_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *CurrencyFormat) GetCode() Currency {
	if x, ok := x.GetType().(*CurrencyFormat_Code); ok {
		return x.Code
	}
	return Currency_XXX
}

func (x *CurrencyFormat) GetSymbol() Currency {
	if x, ok := x.GetType().(*CurrencyFormat_Symbol); ok {
		return x.Symbol
	}
	return Currency_XXX
}

func (x *CurrencyFormat) GetName() Currency {
	if x, ok := x.GetType().(*CurrencyFormat_Name); ok {
		return x.Name
	}
	return Currency_XXX
}

func (x *CurrencyFormat) GetCentMode() bool {
	if x != nil {
		return x.CentMode
	}
	return false
}

type isCurrencyFormat_Type interface {
	isCurrencyFormat_Type()
}

type CurrencyFormat_Code struct {
	Code Currency `protobuf:"varint,1,opt,name=code,proto3,enum=pgde.format.Currency,oneof"`
}

type CurrencyFormat_Symbol struct {
	Symbol Currency `protobuf:"varint,2,opt,name=symbol,proto3,enum=pgde.format.Currency,oneof"`
}

type CurrencyFormat_Name struct {
	Name Currency `protobuf:"varint,3,opt,name=name,proto3,enum=pgde.format.Currency,oneof"`
}

func (*CurrencyFormat_Code) isCurrencyFormat_Type() {}

func (*CurrencyFormat_Symbol) isCurrencyFormat_Type() {}

func (*CurrencyFormat_Name) isCurrencyFormat_Type() {}

type NumberFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    NumberFormat_Builtin `protobuf:"varint,1,opt,name=type,proto3,enum=pgde.format.NumberFormat_Builtin" json:"type,omitempty"`
	Locale  string               `protobuf:"bytes,2,opt,name=locale,proto3" json:"locale,omitempty"`
	Ordinal bool                 `protobuf:"varint,3,opt,name=ordinal,proto3" json:"ordinal,omitempty"`
	Unit    *units.Unit          `protobuf:"bytes,4,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (x *NumberFormat) Reset() {
	*x = NumberFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_format_format_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NumberFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NumberFormat) ProtoMessage() {}

func (x *NumberFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_format_format_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NumberFormat.ProtoReflect.Descriptor instead.
func (*NumberFormat) Descriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{3}
}

func (x *NumberFormat) GetType() NumberFormat_Builtin {
	if x != nil {
		return x.Type
	}
	return NumberFormat_decimal
}

func (x *NumberFormat) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *NumberFormat) GetOrdinal() bool {
	if x != nil {
		return x.Ordinal
	}
	return false
}

func (x *NumberFormat) GetUnit() *units.Unit {
	if x != nil {
		return x.Unit
	}
	return nil
}

type BytesFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type BytesFormat_Builtin `protobuf:"varint,1,opt,name=type,proto3,enum=pgde.format.BytesFormat_Builtin" json:"type,omitempty"`
}

func (x *BytesFormat) Reset() {
	*x = BytesFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pgde_format_format_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BytesFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BytesFormat) ProtoMessage() {}

func (x *BytesFormat) ProtoReflect() protoreflect.Message {
	mi := &file_pgde_format_format_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BytesFormat.ProtoReflect.Descriptor instead.
func (*BytesFormat) Descriptor() ([]byte, []int) {
	return file_pgde_format_format_proto_rawDescGZIP(), []int{4}
}

func (x *BytesFormat) GetType() BytesFormat_Builtin {
	if x != nil {
		return x.Type
	}
	return BytesFormat_ip
}

var file_pgde_format_format_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*Format)(nil),
		Field:         919111,
		Name:          "pgde.format.to",
		Tag:           "bytes,919111,opt,name=to",
		Filename:      "pgde/format/format.proto",
	},
}

// Extension fields to descriptor.FieldOptions.
var (
	// optional pgde.format.Format to = 919111;
	E_To = &file_pgde_format_format_proto_extTypes[0]
)

var File_pgde_format_format_proto protoreflect.FileDescriptor

var file_pgde_format_format_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x70, 0x67, 0x64, 0x65,
	0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x67, 0x64, 0x65, 0x2f,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x75, 0x6e, 0x69, 0x74,
	0x73, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x01,
	0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x48,
	0x00, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x67, 0x64, 0x65,
	0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x48, 0x00, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x33, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x48, 0x00, 0x52,
	0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x48, 0x00, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x87, 0x03, 0x0a, 0x0a, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f,
	0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x69, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e,
	0x49, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x07, 0x69, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74,
	0x22, 0xa0, 0x01, 0x0a, 0x07, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x0a,
	0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x66, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x66,
	0x44, 0x61, 0x79, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x6f, 0x66, 0x44, 0x61, 0x79, 0x32, 0x34,
	0x48, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x44, 0x61, 0x74,
	0x65, 0x74, 0x69, 0x6d, 0x65, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x6d, 0x65, 0x64, 0x69, 0x75,
	0x6d, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x32, 0x34, 0x48, 0x10, 0x05, 0x12, 0x10,
	0x0a, 0x0c, 0x66, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x10, 0x06,
	0x12, 0x13, 0x0a, 0x0f, 0x66, 0x75, 0x6c, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65,
	0x32, 0x34, 0x48, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x10, 0x08, 0x22, 0x66, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x07,
	0x0a, 0x03, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x6e, 0x61, 0x6e, 0x6f, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x6d, 0x69, 0x6c, 0x6c,
	0x69, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x10,
	0x05, 0x12, 0x08, 0x0a, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x10, 0x06, 0x22, 0xbf, 0x01, 0x0a, 0x0e,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2b,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70,
	0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x48, 0x00, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x67,
	0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x48, 0x00, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2b, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x70, 0x67, 0x64,
	0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x48, 0x00, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x65, 0x6e,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xf4, 0x01,
	0x0a, 0x0c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x35,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x70,
	0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x75, 0x6e, 0x69,
	0x74, 0x73, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x22, 0x55, 0x0a,
	0x07, 0x42, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x64, 0x65, 0x63, 0x69,
	0x6d, 0x61, 0x6c, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74,
	0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x73, 0x63, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x10, 0x03,
	0x12, 0x12, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x10, 0x04, 0x22, 0xb4, 0x01, 0x0a, 0x0b, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x20, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x42, 0x75, 0x69,
	0x6c, 0x74, 0x69, 0x6e, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x6f, 0x0a, 0x07, 0x42, 0x75,
	0x69, 0x6c, 0x74, 0x69, 0x6e, 0x12, 0x06, 0x0a, 0x02, 0x69, 0x70, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x10,
	0x02, 0x12, 0x0a, 0x0a, 0x06, 0x69, 0x70, 0x50, 0x6f, 0x72, 0x74, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x69, 0x70, 0x76, 0x34, 0x50, 0x6f, 0x72, 0x74, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x69,
	0x70, 0x76, 0x36, 0x50, 0x6f, 0x72, 0x74, 0x10, 0x05, 0x12, 0x08, 0x0a, 0x04, 0x63, 0x69, 0x64,
	0x72, 0x10, 0x06, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x69, 0x64, 0x72, 0x76, 0x34, 0x10, 0x07, 0x12,
	0x0a, 0x0a, 0x06, 0x63, 0x69, 0x64, 0x72, 0x76, 0x36, 0x10, 0x08, 0x3a, 0x44, 0x0a, 0x02, 0x74,
	0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xc7, 0x8c, 0x38, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x67, 0x64, 0x65, 0x2e,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x02, 0x74,
	0x6f, 0x42, 0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x6d, 0x70, 0x69, 0x72, 0x65, 0x66, 0x6f, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x61, 0x72, 0x74, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x3b, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pgde_format_format_proto_rawDescOnce sync.Once
	file_pgde_format_format_proto_rawDescData = file_pgde_format_format_proto_rawDesc
)

func file_pgde_format_format_proto_rawDescGZIP() []byte {
	file_pgde_format_format_proto_rawDescOnce.Do(func() {
		file_pgde_format_format_proto_rawDescData = protoimpl.X.CompressGZIP(file_pgde_format_format_proto_rawDescData)
	})
	return file_pgde_format_format_proto_rawDescData
}

var file_pgde_format_format_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_pgde_format_format_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pgde_format_format_proto_goTypes = []interface{}{
	(TimeFormat_Builtin)(0),         // 0: pgde.format.TimeFormat.Builtin
	(TimeFormat_IntUnit)(0),         // 1: pgde.format.TimeFormat.IntUnit
	(NumberFormat_Builtin)(0),       // 2: pgde.format.NumberFormat.Builtin
	(BytesFormat_Builtin)(0),        // 3: pgde.format.BytesFormat.Builtin
	(*Format)(nil),                  // 4: pgde.format.Format
	(*TimeFormat)(nil),              // 5: pgde.format.TimeFormat
	(*CurrencyFormat)(nil),          // 6: pgde.format.CurrencyFormat
	(*NumberFormat)(nil),            // 7: pgde.format.NumberFormat
	(*BytesFormat)(nil),             // 8: pgde.format.BytesFormat
	(Currency)(0),                   // 9: pgde.format.Currency
	(*units.Unit)(nil),              // 10: pgde.units.Unit
	(*descriptor.FieldOptions)(nil), // 11: google.protobuf.FieldOptions
}
var file_pgde_format_format_proto_depIdxs = []int32{
	5,  // 0: pgde.format.Format.time:type_name -> pgde.format.TimeFormat
	6,  // 1: pgde.format.Format.currency:type_name -> pgde.format.CurrencyFormat
	7,  // 2: pgde.format.Format.number:type_name -> pgde.format.NumberFormat
	8,  // 3: pgde.format.Format.bytes:type_name -> pgde.format.BytesFormat
	0,  // 4: pgde.format.TimeFormat.type:type_name -> pgde.format.TimeFormat.Builtin
	1,  // 5: pgde.format.TimeFormat.intUnit:type_name -> pgde.format.TimeFormat.IntUnit
	9,  // 6: pgde.format.CurrencyFormat.code:type_name -> pgde.format.Currency
	9,  // 7: pgde.format.CurrencyFormat.symbol:type_name -> pgde.format.Currency
	9,  // 8: pgde.format.CurrencyFormat.name:type_name -> pgde.format.Currency
	2,  // 9: pgde.format.NumberFormat.type:type_name -> pgde.format.NumberFormat.Builtin
	10, // 10: pgde.format.NumberFormat.unit:type_name -> pgde.units.Unit
	3,  // 11: pgde.format.BytesFormat.type:type_name -> pgde.format.BytesFormat.Builtin
	11, // 12: pgde.format.to:extendee -> google.protobuf.FieldOptions
	4,  // 13: pgde.format.to:type_name -> pgde.format.Format
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	13, // [13:14] is the sub-list for extension type_name
	12, // [12:13] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_pgde_format_format_proto_init() }
func file_pgde_format_format_proto_init() {
	if File_pgde_format_format_proto != nil {
		return
	}
	file_pgde_format_currency_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pgde_format_format_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Format); i {
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
		file_pgde_format_format_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimeFormat); i {
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
		file_pgde_format_format_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyFormat); i {
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
		file_pgde_format_format_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NumberFormat); i {
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
		file_pgde_format_format_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BytesFormat); i {
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
	file_pgde_format_format_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Format_Time)(nil),
		(*Format_Currency)(nil),
		(*Format_Number)(nil),
		(*Format_Bytes)(nil),
	}
	file_pgde_format_format_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CurrencyFormat_Code)(nil),
		(*CurrencyFormat_Symbol)(nil),
		(*CurrencyFormat_Name)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pgde_format_format_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pgde_format_format_proto_goTypes,
		DependencyIndexes: file_pgde_format_format_proto_depIdxs,
		EnumInfos:         file_pgde_format_format_proto_enumTypes,
		MessageInfos:      file_pgde_format_format_proto_msgTypes,
		ExtensionInfos:    file_pgde_format_format_proto_extTypes,
	}.Build()
	File_pgde_format_format_proto = out.File
	file_pgde_format_format_proto_rawDesc = nil
	file_pgde_format_format_proto_goTypes = nil
	file_pgde_format_format_proto_depIdxs = nil
}
