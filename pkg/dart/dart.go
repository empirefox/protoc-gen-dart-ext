package dart

import (
	"fmt"
	"path/filepath"

	pgs "github.com/lyft/protoc-gen-star"
)

type UsedNames []map[Qualifier]bool

func (un UsedNames) Add(name Qualifier) {
	if un[0] == nil {
		un[0] = make(map[Qualifier]bool)
	}
	un[0][name] = true
}
func (un *UsedNames) AddMap(m map[Qualifier]bool) {
	if len(*un) == 0 {
		return
	}
	*un = append(*un, m)
}
func (un *UsedNames) AddSlice(names []Qualifier) {
	m := make(map[Qualifier]bool, len(names))
	for _, n := range names {
		m[n] = true
	}
	un.AddMap(m)
}
func (un UsedNames) Used(name Qualifier) bool {
	for _, m := range un {
		if m[name] {
			return true
		}
	}
	return false
}

type Names interface {
	Name() Qualifier
}

type names struct {
	name      Qualifier
	usedNames UsedNames
}

func (ns *names) Name() Qualifier { return ns.name }

type FieldNames interface {
	Names
	HasMethodName() Qualifier
	ClearMethodName() Qualifier
	EnsureMethodName() Qualifier
}

type fieldNames struct {
	fieldName        Qualifier
	hasMethodName    Qualifier
	clearMethodName  Qualifier
	ensureMethodName Qualifier
}

func (ns *fieldNames) Name() Qualifier             { return ns.fieldName }
func (ns *fieldNames) HasMethodName() Qualifier    { return ns.hasMethodName }
func (ns *fieldNames) ClearMethodName() Qualifier  { return ns.clearMethodName }
func (ns *fieldNames) EnsureMethodName() Qualifier { return ns.ensureMethodName }

type OneOfNames interface {
	Names
	WhichOneofMethodName() Qualifier
	ClearMethodName() Qualifier
	OneofEnumName() Qualifier
}

type oneOfNames struct {
	oneofName            Qualifier
	whichOneofMethodName Qualifier
	clearMethodName      Qualifier
	oneofEnumName        Qualifier
}

func (ns *oneOfNames) Name() Qualifier                 { return ns.oneofName }
func (ns *oneOfNames) WhichOneofMethodName() Qualifier { return ns.whichOneofMethodName }
func (ns *oneOfNames) ClearMethodName() Qualifier      { return ns.clearMethodName }
func (ns *oneOfNames) OneofEnumName() Qualifier        { return ns.oneofEnumName }

type Dart struct {
	pkgs       map[pgs.Package]Qualifier
	files      map[pgs.File]*names
	messages   map[pgs.Message]*names
	fields     map[pgs.Field]*fieldNames
	oneofs     map[pgs.OneOf]*oneOfNames
	enums      map[pgs.Enum]*names
	enumValues map[pgs.EnumValue]*names
}

func NewDart() *Dart {
	return &Dart{
		pkgs:       make(map[pgs.Package]Qualifier),
		files:      make(map[pgs.File]*names),
		messages:   make(map[pgs.Message]*names),
		fields:     make(map[pgs.Field]*fieldNames),
		oneofs:     make(map[pgs.OneOf]*oneOfNames),
		enums:      make(map[pgs.Enum]*names),
		enumValues: make(map[pgs.EnumValue]*names),
	}
}

func (d *Dart) EntityNameOf(p pgs.Package) Qualifier {
	name, ok := d.pkgs[p]
	if !ok {
		name = entityNameOfPkg(p)
		d.pkgs[p] = name
	}
	return name
}

func entityNameOfPkg(pkg pgs.Package) Qualifier {
	pkgName := pkg.ProtoName().String()
	ext := filepath.Ext(pkgName)
	if ext != "" {
		pkgName = ext[1:]
	}
	return Qualifier(pkgName).ToCamel()
}

func (d *Dart) NameOf(e pgs.Entity) Qualifier {
	return d.entityOf(e).Name()
}

func (d *Dart) FieldNames(e pgs.Field) FieldNames {
	return d.entityOfField(e)
}

func (d *Dart) OneOfNames(e pgs.OneOf) OneOfNames {
	return d.entityOfOneof(e)
}

func (d *Dart) entityOf(e pgs.Entity) Names {
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

func (d *Dart) entityOfFile(i pgs.File) *names {
	nty, ok := d.files[i]
	if !ok {
		nty = &names{
			name:      "",
			usedNames: usedTopLevelNames(),
		}
		d.files[i] = nty
	}
	return nty
}

func (d *Dart) entityOfMessage(i pgs.Message) *names {
	nty, ok := d.messages[i]
	if !ok {
		parentEntry := d.entityOf(i.Parent()).(*names)
		name := messageOrEnumClassName(Qualifier(i.Name()), parentEntry.usedNames, parentEntry.name)
		parentEntry.usedNames.Add(name)

		var reserved map[Qualifier]bool // compute?
		existingNames := reservedMemberNames()
		existingNames.AddMap(reserved)

		nty = &names{
			name:      name, // check dart_options.dart_name?
			usedNames: existingNames,
		}
		d.messages[i] = nty
	}
	return nty
}

func (d *Dart) entityOfOneof(i pgs.OneOf) *oneOfNames {
	nty, ok := d.oneofs[i]
	if !ok {
		parentEntry := d.entityOfMessage(i.Message())
		oneofNameVariants := func(name Qualifier) []Qualifier {
			return []Qualifier{
				_defaultWhichMethodName(name),
				_defaultClearMethodName(name),
			}
		}
		oneofName := disambiguateName(
			Qualifier(i.Name()).ToCamel(),
			parentEntry.usedNames, new(defaultSuffixes), oneofNameVariants)
		parentEntry.usedNames.Add(oneofName)

		f := d.entityOfFile(i.File())
		oneofEnumName := oneofEnumClassName(Qualifier(i.Name()),
			f.usedNames,
			parentEntry.Name())
		f.usedNames.Add(oneofEnumName)

		enumMapName := disambiguateName(
			Qualifier(fmt.Sprintf("_$%sByTag", oneofEnumName)),
			parentEntry.usedNames, new(defaultSuffixes),
			nil)
		parentEntry.usedNames.Add(enumMapName)

		nty = &oneOfNames{
			oneofName:            oneofName,
			oneofEnumName:        oneofEnumName,
			clearMethodName:      _defaultClearMethodName(oneofName),
			whichOneofMethodName: _defaultWhichMethodName(oneofName),
		}
		d.oneofs[i] = nty
	}
	return nty
}

func (d *Dart) entityOfField(i pgs.Field) *fieldNames {
	nty, ok := d.fields[i]
	if !ok {
		parentEntry := d.entityOfMessage(i.Message())
		suffix := newMemberNamesSuffix(i.Descriptor().GetNumber())
		var generateNameVariants generateVariantsFunc
		if !i.Required() {
			generateNameVariants = func(name Qualifier) []Qualifier {
				return []Qualifier{
					name.ToLowerCamel(),
					_defaultHasMethodName(name),
					_defaultClearMethodName(name),
				}
			}
		}
		name := disambiguateName(Qualifier(i.Name()), parentEntry.usedNames, suffix, generateNameVariants)
		parentEntry.usedNames.Add(name)

		var ensureMethodName Qualifier
		if i.Type().IsEmbed() {
			ensureMethodName = _defaultEnsureMethodName(name)
		}

		nty = &fieldNames{
			fieldName:        name.ToLowerCamel(),
			hasMethodName:    _defaultHasMethodName(name),
			clearMethodName:  _defaultClearMethodName(name),
			ensureMethodName: ensureMethodName,
		}
		d.fields[i] = nty
	}
	return nty
}

func (d *Dart) entityOfEnum(i pgs.Enum) *names {
	nty, ok := d.enums[i]
	if !ok {
		parentEntry := d.entityOf(i.Parent()).(*names)
		name := messageOrEnumClassName(Qualifier(i.Name()), parentEntry.usedNames, parentEntry.name)
		parentEntry.usedNames.Add(name)

		nty = &names{
			name:      name,
			usedNames: reservedEnumNames(),
		}
		d.enums[i] = nty
	}
	return nty
}

func (d *Dart) entityOfEnumValue(i pgs.EnumValue) *names {
	nty, ok := d.enumValues[i]
	if !ok {
		parentEntry := d.entityOfEnum(i.Enum())
		name := disambiguateName(avoidInitialUnderscore(Qualifier(i.Name())),
			parentEntry.usedNames, new(enumSuffixes), nil)
		parentEntry.usedNames.Add(name)

		nty = &names{
			name:      name,
			usedNames: nil,
		}
		d.enumValues[i] = nty
	}
	return nty
}

type suffixGetter interface {
	hasNext() bool
	next() Qualifier
}

type defaultSuffixes struct {
	idx int
}

func (s *defaultSuffixes) hasNext() bool { return true }
func (s *defaultSuffixes) next() Qualifier {
	defer func() { s.idx++ }()
	if s.idx == 0 {
		return "_"
	}
	return Qualifier(fmt.Sprintf("_%d", s.idx))
}

type _memberNamesSuffix struct {
	last   Qualifier
	suffix Qualifier
}

func newMemberNamesSuffix(number int32) *_memberNamesSuffix {
	return &_memberNamesSuffix{
		suffix: Qualifier(fmt.Sprintf("_%d", number)),
	}
}
func (s *_memberNamesSuffix) hasNext() bool { return true }
func (s *_memberNamesSuffix) next() Qualifier {
	s.suffix += s.last
	s.last = s.suffix
	return s.suffix
}

type enumSuffixes struct {
	currentSuffix Qualifier
}

func (s *enumSuffixes) hasNext() bool { return true }
func (s *enumSuffixes) next() Qualifier {
	s.currentSuffix += "_"
	return s.currentSuffix
}

type generateVariantsFunc func(name Qualifier) []Qualifier

func disambiguateName(name Qualifier, usedNames UsedNames, getter suffixGetter, generateVariants generateVariantsFunc) Qualifier {
	if generateVariants == nil {
		generateVariants = func(name Qualifier) []Qualifier { return []Qualifier{name} }
	}

	allVariantsAvailable := func(variants []Qualifier) bool {
		for _, v := range variants {
			if usedNames.Used(v) {
				return false
			}
		}
		return true
	}

	var usedSuffix Qualifier
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

func _defaultHasMethodName(fieldMethodSuffix Qualifier) Qualifier {
	return "has" + fieldMethodSuffix
}
func _defaultClearMethodName(fieldMethodSuffix Qualifier) Qualifier {
	return "clear" + fieldMethodSuffix
}
func _defaultWhichMethodName(oneofMethodSuffix Qualifier) Qualifier {
	return "which" + oneofMethodSuffix
}
func _defaultEnsureMethodName(fieldMethodSuffix Qualifier) Qualifier {
	return "ensure" + fieldMethodSuffix
}

func oneofEnumClassName(descriptorName Qualifier, usedNames UsedNames, parent Qualifier) Qualifier {
	descriptorName = Qualifier(fmt.Sprintf("%s_%s", parent, descriptorName.ToCamel()))
	avoidName := avoidInitialUnderscore(descriptorName)
	return disambiguateName(avoidName, usedNames, new(defaultSuffixes), nil)
}

func messageOrEnumClassName(descriptorName Qualifier, usedNames UsedNames, parent Qualifier) Qualifier {
	if parent != "" {
		descriptorName = Qualifier(fmt.Sprintf("%s_%s", parent, descriptorName))
	}
	avoidName := avoidInitialUnderscore(descriptorName)
	return disambiguateName(avoidName, usedNames, new(defaultSuffixes), nil)
}

func avoidInitialUnderscore(input Qualifier) Qualifier {
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

var toplevelReservedCapitalizedNames = map[Qualifier]bool{
	"List":     true,
	"Function": true,
	"Map":      true,
}

// List of Dart language reserved words in names which cannot be used in a
// subclass of GeneratedMessage.
var _dartReservedWords = map[Qualifier]bool{
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
var _generatedMessageNames = map[Qualifier]bool{
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
var _protobufEnumNames = map[Qualifier]bool{
	"List":    true,
	"valueOf": true,
	"values":  true,
}

var GeneratedMessage_reservedNames = map[Qualifier]bool{
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

var ProtobufEnum_reservedNames = map[Qualifier]bool{
	"==":           true,
	"Object":       true,
	"ProtobufEnum": true,
	"hashCode":     true,
	"initByValue":  true,
	"noSuchMethod": true,
	"runtimeType":  true,
	"toString":     true,
}
