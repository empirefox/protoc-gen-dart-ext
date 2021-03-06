// DO NOT EDIT. Generated by protoc-gen-dart-ext/tools.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: pgde/units/atom.proto

package units

import (
	_ "github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/zero"
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

// Part of:
// https://physics.nist.gov/cuu/Units/units.html
// https://physics.nist.gov/cuu/Units/outside.html
type Atom int32

const (
	Atom_noAtom Atom = 0
	// m
	Atom_meter Atom = 1
	// ft
	Atom_foot Atom = 2
	// in
	Atom_inch Atom = 3
	// yd
	Atom_yard Atom = 4
	// mi
	Atom_mile Atom = 5
	// NM
	Atom_nauticalMile Atom = 6
	// LY
	Atom_lightYear Atom = 7
	// ha
	Atom_hectare Atom = 8
	// a
	Atom_are Atom = 9
	// L
	Atom_liter Atom = 10
	// gal
	Atom_gallon Atom = 11
	// bbl
	Atom_barrel Atom = 12
	// g
	Atom_gram Atom = 13
	// t
	Atom_ton Atom = 14
	// lbs
	Atom_pound Atom = 15
	// oz
	Atom_ounce Atom = 16
	// s
	Atom_second Atom = 17
	// min
	Atom_minute Atom = 18
	// h
	Atom_hour Atom = 19
	// d
	Atom_day Atom = 20
	// week
	Atom_week Atom = 21
	// month
	Atom_month Atom = 22
	// yr
	Atom_year Atom = 23
	// century
	Atom_century Atom = 24
	// ''
	Atom_secondAngle Atom = 25
	// '
	Atom_minuteAngle Atom = 26
	// °
	Atom_degree Atom = 27
	// A
	Atom_ampere Atom = 28
	// eV
	Atom_electronvolt Atom = 29
	// B
	Atom_bel Atom = 30
	// K
	Atom_kelvin Atom = 31
	// mol
	Atom_mole Atom = 32
	// cd
	Atom_candela Atom = 33
	// %
	Atom_percent Atom = 34
	// ‰
	Atom_perThousand Atom = 35
	// b
	Atom_bit Atom = 36
	// B
	Atom_byte Atom = 37
	// character
	Atom_character Atom = 38
	// word
	Atom_word Atom = 39
	// rad
	Atom_radian Atom = 40
	// sr
	Atom_steradian Atom = 41
	// Hz
	Atom_hertz Atom = 42
	// N
	Atom_newton Atom = 43
	// Pa
	Atom_pascal Atom = 44
	// J
	Atom_joule Atom = 45
	// W
	Atom_watt Atom = 46
	// C
	Atom_coulomb Atom = 47
	// V
	Atom_volt Atom = 48
	// F
	Atom_farad Atom = 49
	// Ω
	Atom_ohm Atom = 50
	// S
	Atom_siemens Atom = 51
	// Wb
	Atom_weber Atom = 52
	// T
	Atom_tesla Atom = 53
	// H
	Atom_henry Atom = 54
	// °C
	Atom_degreeCelsius Atom = 55
	// lm
	Atom_lumen Atom = 56
	// lx
	Atom_lux Atom = 57
	// Bq
	Atom_becquerel Atom = 58
	// Gy
	Atom_gray Atom = 59
	// Sv
	Atom_sievert Atom = 60
	// kat
	Atom_katal Atom = 61
)

// Enum value maps for Atom.
var (
	Atom_name = map[int32]string{
		0:  "noAtom",
		1:  "meter",
		2:  "foot",
		3:  "inch",
		4:  "yard",
		5:  "mile",
		6:  "nauticalMile",
		7:  "lightYear",
		8:  "hectare",
		9:  "are",
		10: "liter",
		11: "gallon",
		12: "barrel",
		13: "gram",
		14: "ton",
		15: "pound",
		16: "ounce",
		17: "second",
		18: "minute",
		19: "hour",
		20: "day",
		21: "week",
		22: "month",
		23: "year",
		24: "century",
		25: "secondAngle",
		26: "minuteAngle",
		27: "degree",
		28: "ampere",
		29: "electronvolt",
		30: "bel",
		31: "kelvin",
		32: "mole",
		33: "candela",
		34: "percent",
		35: "perThousand",
		36: "bit",
		37: "byte",
		38: "character",
		39: "word",
		40: "radian",
		41: "steradian",
		42: "hertz",
		43: "newton",
		44: "pascal",
		45: "joule",
		46: "watt",
		47: "coulomb",
		48: "volt",
		49: "farad",
		50: "ohm",
		51: "siemens",
		52: "weber",
		53: "tesla",
		54: "henry",
		55: "degreeCelsius",
		56: "lumen",
		57: "lux",
		58: "becquerel",
		59: "gray",
		60: "sievert",
		61: "katal",
	}
	Atom_value = map[string]int32{
		"noAtom":        0,
		"meter":         1,
		"foot":          2,
		"inch":          3,
		"yard":          4,
		"mile":          5,
		"nauticalMile":  6,
		"lightYear":     7,
		"hectare":       8,
		"are":           9,
		"liter":         10,
		"gallon":        11,
		"barrel":        12,
		"gram":          13,
		"ton":           14,
		"pound":         15,
		"ounce":         16,
		"second":        17,
		"minute":        18,
		"hour":          19,
		"day":           20,
		"week":          21,
		"month":         22,
		"year":          23,
		"century":       24,
		"secondAngle":   25,
		"minuteAngle":   26,
		"degree":        27,
		"ampere":        28,
		"electronvolt":  29,
		"bel":           30,
		"kelvin":        31,
		"mole":          32,
		"candela":       33,
		"percent":       34,
		"perThousand":   35,
		"bit":           36,
		"byte":          37,
		"character":     38,
		"word":          39,
		"radian":        40,
		"steradian":     41,
		"hertz":         42,
		"newton":        43,
		"pascal":        44,
		"joule":         45,
		"watt":          46,
		"coulomb":       47,
		"volt":          48,
		"farad":         49,
		"ohm":           50,
		"siemens":       51,
		"weber":         52,
		"tesla":         53,
		"henry":         54,
		"degreeCelsius": 55,
		"lumen":         56,
		"lux":           57,
		"becquerel":     58,
		"gray":          59,
		"sievert":       60,
		"katal":         61,
	}
)

func (x Atom) Enum() *Atom {
	p := new(Atom)
	*p = x
	return p
}

func (x Atom) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Atom) Descriptor() protoreflect.EnumDescriptor {
	return file_pgde_units_atom_proto_enumTypes[0].Descriptor()
}

func (Atom) Type() protoreflect.EnumType {
	return &file_pgde_units_atom_proto_enumTypes[0]
}

func (x Atom) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Atom.Descriptor instead.
func (Atom) EnumDescriptor() ([]byte, []int) {
	return file_pgde_units_atom_proto_rawDescGZIP(), []int{0}
}

var File_pgde_units_atom_proto protoreflect.FileDescriptor

var file_pgde_units_atom_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x2f, 0x61, 0x74, 0x6f,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x67, 0x64, 0x65, 0x2e, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x1a, 0x14, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x7a, 0x65, 0x72, 0x6f, 0x2f, 0x7a,
	0x65, 0x72, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xec, 0x05, 0x0a, 0x04, 0x41, 0x74,
	0x6f, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x6f, 0x41, 0x74, 0x6f, 0x6d, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x66, 0x6f, 0x6f,
	0x74, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x69, 0x6e, 0x63, 0x68, 0x10, 0x03, 0x12, 0x08, 0x0a,
	0x04, 0x79, 0x61, 0x72, 0x64, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x6d, 0x69, 0x6c, 0x65, 0x10,
	0x05, 0x12, 0x10, 0x0a, 0x0c, 0x6e, 0x61, 0x75, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x4d, 0x69, 0x6c,
	0x65, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x59, 0x65, 0x61, 0x72,
	0x10, 0x07, 0x12, 0x0b, 0x0a, 0x07, 0x68, 0x65, 0x63, 0x74, 0x61, 0x72, 0x65, 0x10, 0x08, 0x12,
	0x07, 0x0a, 0x03, 0x61, 0x72, 0x65, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x6c, 0x69, 0x74, 0x65,
	0x72, 0x10, 0x0a, 0x12, 0x0a, 0x0a, 0x06, 0x67, 0x61, 0x6c, 0x6c, 0x6f, 0x6e, 0x10, 0x0b, 0x12,
	0x0a, 0x0a, 0x06, 0x62, 0x61, 0x72, 0x72, 0x65, 0x6c, 0x10, 0x0c, 0x12, 0x08, 0x0a, 0x04, 0x67,
	0x72, 0x61, 0x6d, 0x10, 0x0d, 0x12, 0x07, 0x0a, 0x03, 0x74, 0x6f, 0x6e, 0x10, 0x0e, 0x12, 0x09,
	0x0a, 0x05, 0x70, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x0f, 0x12, 0x09, 0x0a, 0x05, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x10, 0x10, 0x12, 0x0a, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x10, 0x11,
	0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x10, 0x12, 0x12, 0x08, 0x0a, 0x04,
	0x68, 0x6f, 0x75, 0x72, 0x10, 0x13, 0x12, 0x07, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x10, 0x14, 0x12,
	0x08, 0x0a, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x10, 0x15, 0x12, 0x09, 0x0a, 0x05, 0x6d, 0x6f, 0x6e,
	0x74, 0x68, 0x10, 0x16, 0x12, 0x08, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x10, 0x17, 0x12, 0x0b,
	0x0a, 0x07, 0x63, 0x65, 0x6e, 0x74, 0x75, 0x72, 0x79, 0x10, 0x18, 0x12, 0x0f, 0x0a, 0x0b, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x19, 0x12, 0x0f, 0x0a, 0x0b,
	0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x10, 0x1a, 0x12, 0x0a, 0x0a,
	0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x10, 0x1b, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x6d, 0x70,
	0x65, 0x72, 0x65, 0x10, 0x1c, 0x12, 0x10, 0x0a, 0x0c, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x72, 0x6f,
	0x6e, 0x76, 0x6f, 0x6c, 0x74, 0x10, 0x1d, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x65, 0x6c, 0x10, 0x1e,
	0x12, 0x0a, 0x0a, 0x06, 0x6b, 0x65, 0x6c, 0x76, 0x69, 0x6e, 0x10, 0x1f, 0x12, 0x08, 0x0a, 0x04,
	0x6d, 0x6f, 0x6c, 0x65, 0x10, 0x20, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x65, 0x6c,
	0x61, 0x10, 0x21, 0x12, 0x0b, 0x0a, 0x07, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x10, 0x22,
	0x12, 0x0f, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x54, 0x68, 0x6f, 0x75, 0x73, 0x61, 0x6e, 0x64, 0x10,
	0x23, 0x12, 0x07, 0x0a, 0x03, 0x62, 0x69, 0x74, 0x10, 0x24, 0x12, 0x08, 0x0a, 0x04, 0x62, 0x79,
	0x74, 0x65, 0x10, 0x25, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x72, 0x61, 0x63, 0x74, 0x65,
	0x72, 0x10, 0x26, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x6f, 0x72, 0x64, 0x10, 0x27, 0x12, 0x0a, 0x0a,
	0x06, 0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x10, 0x28, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x74, 0x65,
	0x72, 0x61, 0x64, 0x69, 0x61, 0x6e, 0x10, 0x29, 0x12, 0x09, 0x0a, 0x05, 0x68, 0x65, 0x72, 0x74,
	0x7a, 0x10, 0x2a, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x65, 0x77, 0x74, 0x6f, 0x6e, 0x10, 0x2b, 0x12,
	0x0a, 0x0a, 0x06, 0x70, 0x61, 0x73, 0x63, 0x61, 0x6c, 0x10, 0x2c, 0x12, 0x09, 0x0a, 0x05, 0x6a,
	0x6f, 0x75, 0x6c, 0x65, 0x10, 0x2d, 0x12, 0x08, 0x0a, 0x04, 0x77, 0x61, 0x74, 0x74, 0x10, 0x2e,
	0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6c, 0x6f, 0x6d, 0x62, 0x10, 0x2f, 0x12, 0x08, 0x0a,
	0x04, 0x76, 0x6f, 0x6c, 0x74, 0x10, 0x30, 0x12, 0x09, 0x0a, 0x05, 0x66, 0x61, 0x72, 0x61, 0x64,
	0x10, 0x31, 0x12, 0x07, 0x0a, 0x03, 0x6f, 0x68, 0x6d, 0x10, 0x32, 0x12, 0x0b, 0x0a, 0x07, 0x73,
	0x69, 0x65, 0x6d, 0x65, 0x6e, 0x73, 0x10, 0x33, 0x12, 0x09, 0x0a, 0x05, 0x77, 0x65, 0x62, 0x65,
	0x72, 0x10, 0x34, 0x12, 0x09, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x6c, 0x61, 0x10, 0x35, 0x12, 0x09,
	0x0a, 0x05, 0x68, 0x65, 0x6e, 0x72, 0x79, 0x10, 0x36, 0x12, 0x11, 0x0a, 0x0d, 0x64, 0x65, 0x67,
	0x72, 0x65, 0x65, 0x43, 0x65, 0x6c, 0x73, 0x69, 0x75, 0x73, 0x10, 0x37, 0x12, 0x09, 0x0a, 0x05,
	0x6c, 0x75, 0x6d, 0x65, 0x6e, 0x10, 0x38, 0x12, 0x07, 0x0a, 0x03, 0x6c, 0x75, 0x78, 0x10, 0x39,
	0x12, 0x0d, 0x0a, 0x09, 0x62, 0x65, 0x63, 0x71, 0x75, 0x65, 0x72, 0x65, 0x6c, 0x10, 0x3a, 0x12,
	0x08, 0x0a, 0x04, 0x67, 0x72, 0x61, 0x79, 0x10, 0x3b, 0x12, 0x0b, 0x0a, 0x07, 0x73, 0x69, 0x65,
	0x76, 0x65, 0x72, 0x74, 0x10, 0x3c, 0x12, 0x09, 0x0a, 0x05, 0x6b, 0x61, 0x74, 0x61, 0x6c, 0x10,
	0x3d, 0x1a, 0x05, 0xc8, 0xe4, 0xc0, 0x03, 0x01, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6d, 0x70, 0x69, 0x72, 0x65, 0x66, 0x6f, 0x78,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x64, 0x61, 0x72, 0x74,
	0x2d, 0x65, 0x78, 0x74, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x67, 0x64, 0x65, 0x2f, 0x75, 0x6e,
	0x69, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pgde_units_atom_proto_rawDescOnce sync.Once
	file_pgde_units_atom_proto_rawDescData = file_pgde_units_atom_proto_rawDesc
)

func file_pgde_units_atom_proto_rawDescGZIP() []byte {
	file_pgde_units_atom_proto_rawDescOnce.Do(func() {
		file_pgde_units_atom_proto_rawDescData = protoimpl.X.CompressGZIP(file_pgde_units_atom_proto_rawDescData)
	})
	return file_pgde_units_atom_proto_rawDescData
}

var file_pgde_units_atom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pgde_units_atom_proto_goTypes = []interface{}{
	(Atom)(0), // 0: pgde.units.Atom
}
var file_pgde_units_atom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pgde_units_atom_proto_init() }
func file_pgde_units_atom_proto_init() {
	if File_pgde_units_atom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pgde_units_atom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pgde_units_atom_proto_goTypes,
		DependencyIndexes: file_pgde_units_atom_proto_depIdxs,
		EnumInfos:         file_pgde_units_atom_proto_enumTypes,
	}.Build()
	File_pgde_units_atom_proto = out.File
	file_pgde_units_atom_proto_rawDesc = nil
	file_pgde_units_atom_proto_goTypes = nil
	file_pgde_units_atom_proto_depIdxs = nil
}
