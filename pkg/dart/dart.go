package dart

import (
	"fmt"

	pgs "github.com/lyft/protoc-gen-star"
)

type UsedNames []map[pgs.Name]bool

func (un UsedNames) Add(name pgs.Name) {
	if un[0] == nil {
		un[0] = make(map[pgs.Name]bool)
	}
	un[0][name] = true
}
func (un *UsedNames) AddMap(m map[pgs.Name]bool) {
	if len(*un) == 0 {
		return
	}
	*un = append(*un, m)
}
func (un *UsedNames) AddSlice(names []pgs.Name) {
	m := make(map[pgs.Name]bool, len(names))
	for _, n := range names {
		m[n] = true
	}
	un.AddMap(m)
}
func (un UsedNames) Used(name pgs.Name) bool {
	for _, m := range un {
		if m[name] {
			return true
		}
	}
	return false
}

type Entity struct {
	Name      pgs.Name
	usedNames UsedNames
}

type Dart struct {
	files      map[pgs.File]*Entity
	messages   map[pgs.Message]*Entity
	fields     map[pgs.Field]*Entity
	oneofs     map[pgs.OneOf]*Entity
	enums      map[pgs.Enum]*Entity
	enumValues map[pgs.EnumValue]*Entity
}

func NewDart() *Dart {
	return &Dart{
		files:      make(map[pgs.File]*Entity),
		messages:   make(map[pgs.Message]*Entity),
		fields:     make(map[pgs.Field]*Entity),
		oneofs:     make(map[pgs.OneOf]*Entity),
		enums:      make(map[pgs.Enum]*Entity),
		enumValues: make(map[pgs.EnumValue]*Entity),
	}
}

func (d *Dart) NameOf(e pgs.Entity) pgs.Name {
	return d.entityOf(e).Name
}

func (d *Dart) entityOf(e pgs.Entity) *Entity {
	switch i := e.(type) {
	case pgs.File:
		return d.entityOfFile(i)
	case pgs.Message:
		return d.entityOfMessage(i)
	case pgs.Field:
		return d.entityOfField(i)
	case pgs.OneOf:
		return d.entityOfOneof(i)
	case pgs.Enum:
		return d.entityOfEnum(i)
	case pgs.EnumValue:
		return d.entityOfEnumValue(i)
	}
	return nil
}

func (d *Dart) entityOfFile(i pgs.File) *Entity {
	nty, ok := d.files[i]
	if !ok {
		nty = &Entity{
			Name:      "",
			usedNames: usedTopLevelNames(),
		}
		d.files[i] = nty
	}
	return nty
}

func (d *Dart) entityOfMessage(i pgs.Message) *Entity {
	nty, ok := d.messages[i]
	if !ok {
		parentEntry := d.entityOf(i.Parent())
		name := messageOrEnumClassName(i.Name(), parentEntry.usedNames, parentEntry.Name)
		parentEntry.usedNames.Add(name)

		var reserved map[pgs.Name]bool // compute?
		existingNames := reservedMemberNames()
		existingNames.AddMap(reserved)

		nty = &Entity{
			Name:      name, // check dart_options.dart_name?
			usedNames: existingNames,
		}
		d.messages[i] = nty
	}
	return nty
}

func (d *Dart) entityOfOneof(i pgs.OneOf) *Entity {
	nty, ok := d.oneofs[i]
	if !ok {
		parentEntry := d.entityOf(i.Message())
		oneofNameVariants := func(name pgs.Name) []pgs.Name {
			return []pgs.Name{
				_defaultWhichMethodName(name),
				_defaultClearMethodName(name),
			}
		}
		name := disambiguateName(i.Name().UpperCamelCase(),
			parentEntry.usedNames, new(defaultSuffixes), oneofNameVariants)
		parentEntry.usedNames.Add(name)

		nty = &Entity{
			Name:      name.UpperCamelCase(),
			usedNames: nil,
		}
		d.oneofs[i] = nty
	}
	return nty
}

func (d *Dart) entityOfField(i pgs.Field) *Entity {
	nty, ok := d.fields[i]
	if !ok {
		parentEntry := d.entityOf(i.Message())
		suffix := newMemberNamesSuffix(i.Descriptor().GetNumber())
		var generateNameVariants generateVariantsFunc
		if !i.Required() {
			generateNameVariants = func(name pgs.Name) []pgs.Name {
				return []pgs.Name{
					name.LowerCamelCase(),
					_defaultHasMethodName(name),
					_defaultClearMethodName(name),
				}
			}
		}
		name := disambiguateName(i.Name(), parentEntry.usedNames, suffix, generateNameVariants)
		parentEntry.usedNames.Add(name)

		nty = &Entity{
			Name:      name.LowerCamelCase(),
			usedNames: nil,
		}
		d.fields[i] = nty
	}
	return nty
}

func (d *Dart) entityOfEnum(i pgs.Enum) *Entity {
	nty, ok := d.enums[i]
	if !ok {
		parentEntry := d.entityOf(i.Parent())
		name := messageOrEnumClassName(i.Name(), parentEntry.usedNames, parentEntry.Name)
		parentEntry.usedNames.Add(name)

		nty = &Entity{
			Name:      name,
			usedNames: reservedEnumNames(),
		}
		d.enums[i] = nty
	}
	return nty
}

func (d *Dart) entityOfEnumValue(i pgs.EnumValue) *Entity {
	nty, ok := d.enumValues[i]
	if !ok {
		parentEntry := d.entityOf(i.Enum())
		name := disambiguateName(avoidInitialUnderscore(i.Name()),
			parentEntry.usedNames, new(enumSuffixes), nil)
		parentEntry.usedNames.Add(name)

		nty = &Entity{
			Name:      name,
			usedNames: nil,
		}
		d.enumValues[i] = nty
	}
	return nty
}

type suffixGetter interface {
	hasNext() bool
	next() pgs.Name
}

type defaultSuffixes struct {
	idx int
}

func (s *defaultSuffixes) hasNext() bool { return true }
func (s *defaultSuffixes) next() pgs.Name {
	defer func() { s.idx++ }()
	if s.idx == 0 {
		return "_"
	}
	return pgs.Name(fmt.Sprintf("_%d", s.idx))
}

type _memberNamesSuffix struct {
	last   pgs.Name
	suffix pgs.Name
}

func newMemberNamesSuffix(number int32) *_memberNamesSuffix {
	return &_memberNamesSuffix{
		suffix: pgs.Name(fmt.Sprintf("_%d", number)),
	}
}
func (s *_memberNamesSuffix) hasNext() bool { return true }
func (s *_memberNamesSuffix) next() pgs.Name {
	s.suffix += s.last
	s.last = s.suffix
	return s.suffix
}

type enumSuffixes struct {
	currentSuffix pgs.Name
}

func (s *enumSuffixes) hasNext() bool { return true }
func (s *enumSuffixes) next() pgs.Name {
	s.currentSuffix += "_"
	return s.currentSuffix
}

type generateVariantsFunc func(name pgs.Name) []pgs.Name

func disambiguateName(name pgs.Name, usedNames UsedNames, getter suffixGetter, generateVariants generateVariantsFunc) pgs.Name {
	if generateVariants == nil {
		generateVariants = func(name pgs.Name) []pgs.Name { return []pgs.Name{name} }
	}

	allVariantsAvailable := func(variants []pgs.Name) bool {
		for _, v := range variants {
			if usedNames.Used(v) {
				return false
			}
		}
		return true
	}

	var usedSuffix pgs.Name
	candidateVariants := generateVariants(name)

	if !allVariantsAvailable(candidateVariants) {
		for getter.hasNext() {
			suffix := getter.next()
			candidateVariants = generateVariants(name + suffix)
			if allVariantsAvailable(candidateVariants) {
				usedSuffix = suffix
				break
			}
		}
	}

	usedNames.AddSlice(candidateVariants)
	return name + usedSuffix
}

func _defaultHasMethodName(fieldMethodSuffix pgs.Name) pgs.Name {
	return "has" + fieldMethodSuffix
}
func _defaultClearMethodName(fieldMethodSuffix pgs.Name) pgs.Name {
	return "clear" + fieldMethodSuffix
}
func _defaultWhichMethodName(oneofMethodSuffix pgs.Name) pgs.Name {
	return "which" + oneofMethodSuffix
}

func oneofEnumClassName(descriptorName pgs.Name, usedNames UsedNames, parent pgs.Name) pgs.Name {
	descriptorName = descriptorName.UpperCamelCase()
	descriptorName = pgs.Name(fmt.Sprintf("%s_%s", parent, descriptorName))
	avoidName := avoidInitialUnderscore(descriptorName)
	return disambiguateName(avoidName, usedNames, new(defaultSuffixes), nil)
}

func messageOrEnumClassName(descriptorName pgs.Name, usedNames UsedNames, parent pgs.Name) pgs.Name {
	if parent != "" {
		descriptorName = pgs.Name(fmt.Sprintf("%s_%s", parent, descriptorName))
	}
	avoidName := avoidInitialUnderscore(descriptorName)
	return disambiguateName(avoidName, usedNames, new(defaultSuffixes), nil)
}

func avoidInitialUnderscore(input pgs.Name) pgs.Name {
	for input[0] == '_' {
		input = input[1:] + "_"
	}
	return input
}

func usedTopLevelNames() UsedNames {
	return UsedNames{nil, toplevelReservedCapitalizedNames}
}

func reservedEnumNames() UsedNames {
	return UsedNames{nil, ProtobufEnum_reservedNames, _protobufEnumNames}
}

func reservedMemberNames() UsedNames {
	return UsedNames{nil,
		_dartReservedWords,
		GeneratedMessage_reservedNames,
		_generatedMessageNames}
}

var toplevelReservedCapitalizedNames = map[pgs.Name]bool{
	"List":     true,
	"Function": true,
	"Map":      true,
}

// List of Dart language reserved words in names which cannot be used in a
// subclass of GeneratedMessage.
var _dartReservedWords = map[pgs.Name]bool{
	"assert":   true,
	"bool":     true,
	"break":    true,
	"case":     true,
	"catch":    true,
	"class":    true,
	"const":    true,
	"continue": true,
	"default":  true,
	"do":       true,
	"double":   true,
	"else":     true,
	"enum":     true,
	"extends":  true,
	"false":    true,
	"final":    true,
	"finally":  true,
	"for":      true,
	"if":       true,
	"in":       true,
	"int":      true,
	"is":       true,
	"new":      true,
	"null":     true,
	"rethrow":  true,
	"return":   true,
	"super":    true,
	"switch":   true,
	"this":     true,
	"throw":    true,
	"true":     true,
	"try":      true,
	"var":      true,
	"void":     true,
	"while":    true,
	"with":     true,
}

// List of names used in the generated message classes.
//
// This is in addition to GeneratedMessage_reservedNames, which are names from
// the base GeneratedMessage class determined by reflection.
var _generatedMessageNames = map[pgs.Name]bool{
	"create":         true,
	"createRepeated": true,
	"getDefault":     true,
	"List":           true,
	"notSet":         true,
}

// List of names used in the generated enum classes.
//
// This is in addition to ProtobufEnum_reservedNames, which are names from the
// base ProtobufEnum class determined by reflection.
var _protobufEnumNames = map[pgs.Name]bool{
	"List":    true,
	"valueOf": true,
	"values":  true,
}

var GeneratedMessage_reservedNames = map[pgs.Name]bool{
	"hashCode":                   true,
	"noSuchMethod":               true,
	"copyWith":                   true,
	"createEmptyInstance":        true,
	"runtimeType":                true,
	"toString":                   true,
	"freeze":                     true,
	"fromBuffer":                 true,
	"fromJson":                   true,
	"hasRequiredFields":          true,
	"isInitialized":              true,
	"clear":                      true,
	"getTagNumber":               true,
	"check":                      true,
	"writeToBuffer":              true,
	"writeToCodedBufferWriter":   true,
	"mergeFromCodedBufferReader": true,
	"mergeFromBuffer":            true,
	"writeToJson":                true,
	"mergeFromJson":              true,
	"writeToJsonMap":             true,
	"mergeFromJsonMap":           true,
	"addExtension":               true,
	"getExtension":               true,
	"setExtension":               true,
	"hasExtension":               true,
	"clearExtension":             true,
	"getField":                   true,
	"getFieldOrNull":             true,
	"getDefaultForField":         true,
	"setField":                   true,
	"hasField":                   true,
	"clearField":                 true,
	"extensionsAreInitialized":   true,
	"mergeFromMessage":           true,
	"mergeUnknownFields":         true,
	"==":                         true,
	"info_":                      true,
	"GeneratedMessage":           true,
	"Object":                     true,
	"eventPlugin":                true,
	"createMapField":             true,
	"createRepeatedField":        true,
	"unknownFields":              true,
	"clone":                      true,
	"$_get":                      true,
	"$_getI64":                   true,
	"$_getList":                  true,
	"$_getMap":                   true,
	"$_getN":                     true,
	"$_getS":                     true,
	"$_has":                      true,
	"$_setBool":                  true,
	"$_setBytes":                 true,
	"$_setString":                true,
	"$_setFloat":                 true,
	"$_setDouble":                true,
	"$_setSignedInt32":           true,
	"$_setUnsignedInt32":         true,
	"$_setInt64":                 true,
	"$_whichOneof":               true,
	"toBuilder":                  true,
	"toDebugString":              true,
}

var ProtobufEnum_reservedNames = map[pgs.Name]bool{
	"==":           true,
	"Object":       true,
	"ProtobufEnum": true,
	"hashCode":     true,
	"initByValue":  true,
	"noSuchMethod": true,
	"runtimeType":  true,
	"toString":     true,
}
