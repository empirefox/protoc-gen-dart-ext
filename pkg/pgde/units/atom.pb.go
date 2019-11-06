// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pgde/units/atom.proto

package units

import (
	fmt "fmt"
	_ "github.com/empirefox/protoc-gen-dart-ext/pkg/pgde/zero"
	proto "github.com/golang/protobuf/proto"
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

var Atom_name = map[int32]string{
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

var Atom_value = map[string]int32{
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

func (x Atom) String() string {
	return proto.EnumName(Atom_name, int32(x))
}

func (Atom) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d52e034a22a73096, []int{0}
}

func init() {
	proto.RegisterEnum("pgde.units.Atom", Atom_name, Atom_value)
}

func init() { proto.RegisterFile("pgde/units/atom.proto", fileDescriptor_d52e034a22a73096) }

var fileDescriptor_d52e034a22a73096 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x53, 0xcb, 0x56, 0xd4, 0x40,
	0x10, 0x15, 0x61, 0x78, 0x34, 0x20, 0x97, 0x08, 0x3e, 0xf0, 0xfd, 0x56, 0x74, 0x26, 0x2a, 0x2a,
	0xe2, 0x63, 0x81, 0xae, 0xdd, 0xb9, 0xd1, 0x5d, 0x27, 0x29, 0x92, 0x76, 0x3a, 0xdd, 0xb1, 0x53,
	0x0d, 0x0c, 0x5f, 0xe7, 0xd2, 0x0f, 0x71, 0xe9, 0x47, 0x78, 0xaa, 0x67, 0xce, 0x71, 0x93, 0x73,
	0x53, 0xe7, 0xd6, 0xbd, 0xb7, 0xea, 0x54, 0xab, 0xed, 0xae, 0xae, 0x28, 0x8f, 0xce, 0x70, 0x9f,
	0x6b, 0xf6, 0xed, 0xa8, 0x0b, 0x9e, 0x7d, 0xa6, 0xa4, 0x3c, 0x4a, 0xe5, 0x9d, 0xad, 0x44, 0x39,
	0xa3, 0xe0, 0xd3, 0x67, 0xca, 0xd8, 0xfd, 0x3b, 0x50, 0x0b, 0x87, 0xec, 0xdb, 0x4c, 0xa9, 0x45,
	0xe7, 0x05, 0xe1, 0x5c, 0xb6, 0xa2, 0x06, 0x2d, 0x31, 0x05, 0xcc, 0x65, 0xcb, 0x6a, 0xe1, 0xc8,
	0x7b, 0xc6, 0x79, 0x41, 0xc6, 0x95, 0x0d, 0xe6, 0x05, 0x4d, 0x74, 0xa8, 0xb0, 0x20, 0xa8, 0x35,
	0x96, 0x30, 0xc8, 0xa0, 0xd6, 0x9c, 0x8e, 0x6c, 0x4a, 0x6d, 0xbf, 0x48, 0x65, 0x31, 0x5b, 0x57,
	0x2b, 0xd6, 0xd4, 0x0d, 0x7f, 0x23, 0x1d, 0xb0, 0x94, 0xad, 0xaa, 0xa5, 0x86, 0x4a, 0xd6, 0x81,
	0xb0, 0x9c, 0x2d, 0xa9, 0x79, 0x01, 0x2b, 0xe2, 0x64, 0x8d, 0x38, 0x29, 0x09, 0x50, 0x6b, 0x6b,
	0xbd, 0xc3, 0xaa, 0xe0, 0x42, 0x87, 0x40, 0x16, 0x6b, 0xe2, 0x51, 0x07, 0xdd, 0x62, 0x5d, 0xba,
	0xd8, 0x3b, 0x5c, 0x90, 0xae, 0xce, 0x47, 0x57, 0x61, 0x43, 0xa0, 0x8f, 0xae, 0x24, 0x40, 0x9a,
	0x7a, 0x2a, 0xbd, 0xab, 0xb0, 0x29, 0xb8, 0x35, 0x2e, 0x32, 0x21, 0x13, 0x81, 0xc6, 0xc7, 0x80,
	0x8b, 0x22, 0x50, 0xe9, 0x09, 0xb6, 0xa4, 0x74, 0x42, 0x34, 0xc6, 0x76, 0x1a, 0xd5, 0x3b, 0x6e,
	0x70, 0x29, 0x8d, 0x25, 0x59, 0x2f, 0x4b, 0xd6, 0x92, 0x1c, 0xc7, 0x30, 0xc1, 0x95, 0x6c, 0x43,
	0xad, 0x4e, 0x65, 0x0f, 0x5d, 0x6d, 0x09, 0x57, 0xa5, 0x30, 0xd5, 0x9e, 0x16, 0x76, 0xc4, 0xac,
	0xa2, 0x3a, 0x10, 0xe1, 0x9a, 0x60, 0xdd, 0x76, 0x14, 0x08, 0xd7, 0x65, 0x27, 0x64, 0xa9, 0xe4,
	0xe0, 0xdd, 0xb1, 0xb7, 0x8c, 0x1b, 0x12, 0xa0, 0x20, 0x8b, 0x9b, 0x42, 0x1b, 0x93, 0x3d, 0x36,
	0x0e, 0xb7, 0xd2, 0x12, 0xbd, 0x25, 0xdc, 0x4e, 0xbe, 0xda, 0x55, 0x64, 0x35, 0xee, 0xc8, 0x4f,
	0x47, 0x41, 0x72, 0xe0, 0xae, 0x78, 0x76, 0x14, 0xbe, 0x36, 0x3e, 0xf6, 0xda, 0x55, 0xb8, 0x97,
	0x94, 0x0c, 0xe3, 0xbe, 0x74, 0x17, 0x13, 0x26, 0x3c, 0x90, 0x85, 0x97, 0x8d, 0x0e, 0xba, 0x94,
	0x7d, 0x3e, 0x4c, 0x33, 0xfa, 0x50, 0xe1, 0x91, 0x98, 0x05, 0x5d, 0x19, 0xed, 0xf0, 0x58, 0x48,
	0x3d, 0xd3, 0xec, 0xf7, 0x89, 0x8c, 0xdf, 0x50, 0xe0, 0x33, 0xec, 0xa6, 0x03, 0xa0, 0x13, 0x59,
	0xf0, 0x53, 0xc1, 0x9d, 0xee, 0x4b, 0x6d, 0xf1, 0x4c, 0x28, 0x3f, 0x7c, 0xb4, 0x84, 0x61, 0x92,
	0xd4, 0xcc, 0x18, 0xa5, 0xa4, 0x3e, 0x5a, 0xdf, 0x16, 0xc8, 0xa5, 0x9c, 0xe6, 0x7b, 0x2e, 0xdc,
	0x23, 0x1d, 0x74, 0x85, 0x17, 0x12, 0xd0, 0x37, 0x2d, 0x5e, 0x0a, 0xb5, 0x37, 0xd4, 0x92, 0xeb,
	0xb1, 0x27, 0x84, 0x13, 0x2a, 0x28, 0xe0, 0x95, 0x40, 0xa6, 0xde, 0x6a, 0xbc, 0x9e, 0xa6, 0x70,
	0x61, 0x82, 0x37, 0xd9, 0xa6, 0x5a, 0x9f, 0xee, 0xf2, 0x33, 0xd9, 0xde, 0xc4, 0x1e, 0xfb, 0xe9,
	0x46, 0x62, 0x4b, 0x0e, 0x6f, 0x45, 0xd4, 0xc6, 0x53, 0x1c, 0xc8, 0x18, 0x05, 0x95, 0x3f, 0x23,
	0xc9, 0x8d, 0xbc, 0x9b, 0xdd, 0xc8, 0x04, 0xef, 0x67, 0x6e, 0xc7, 0x14, 0x18, 0x1f, 0xa4, 0x73,
	0xac, 0x59, 0x5b, 0x7c, 0xdc, 0x19, 0xfc, 0xfe, 0xf3, 0x6b, 0x7e, 0xee, 0xd3, 0xc1, 0xf7, 0xfd,
	0xda, 0x70, 0x13, 0x8b, 0x51, 0xe9, 0xdb, 0x9c, 0xda, 0xce, 0x04, 0x3a, 0xf2, 0xa7, 0x79, 0x7a,
	0x0c, 0xe5, 0xb0, 0x26, 0x37, 0xac, 0x74, 0xe0, 0x21, 0x9d, 0x72, 0xde, 0x8d, 0xeb, 0xfc, 0xff,
	0xb3, 0x2a, 0x16, 0x13, 0x67, 0xef, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x25, 0xdd, 0xb5, 0x0b,
	0x6b, 0x03, 0x00, 0x00,
}