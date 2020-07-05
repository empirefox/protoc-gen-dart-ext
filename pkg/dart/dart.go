package dart

import (
	"fmt"
	"strings"

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

type ScopedNames interface {
	Names
	getUsedNames() UsedNames
}

type FileNames interface {
	ScopedNames
	EntityName() Qualifier
	L10nName() Qualifier
}

type fileNames struct {
	fileName   Qualifier
	entityName Qualifier
	l10nName   Qualifier
	usedNames  UsedNames
}

func (ns *fileNames) Name() Qualifier         { return ns.fileName }
func (ns *fileNames) EntityName() Qualifier   { return ns.entityName }
func (ns *fileNames) L10nName() Qualifier     { return ns.l10nName }
func (ns *fileNames) getUsedNames() UsedNames { return ns.usedNames }

type MessageNames interface {
	ScopedNames
	ZeroName() Qualifier
	ValidatorName() Qualifier
	FormInputsName() Qualifier
}

type messageNames struct {
	zeroName       Qualifier
	messageName    Qualifier
	validatorName  Qualifier
	formInputsName Qualifier
	usedNames      UsedNames
}

func (ns *messageNames) Name() Qualifier           { return ns.messageName }
func (ns *messageNames) ZeroName() Qualifier       { return ns.zeroName }
func (ns *messageNames) ValidatorName() Qualifier  { return ns.validatorName }
func (ns *messageNames) FormInputsName() Qualifier { return ns.formInputsName }
func (ns *messageNames) getUsedNames() UsedNames   { return ns.usedNames }

type FieldNames interface {
	Names
	ServiceName() Qualifier
	PayloadName() Qualifier
	HasMethodName() Qualifier
	ClearMethodName() Qualifier
	EnsureMethodName() Qualifier
}

type fieldNames struct {
	messageName      Qualifier
	fieldName        Qualifier
	hasMethodName    Qualifier
	clearMethodName  Qualifier
	ensureMethodName Qualifier
}

func (ns *fieldNames) ServiceName() Qualifier {
	return ns.messageName + ns.fieldName.ToCamel() + "Service"
}
func (ns *fieldNames) PayloadName() Qualifier {
	return ns.messageName + ns.fieldName.ToCamel() + "Payload"
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

type EnumNames interface {
	ScopedNames
}

type enumNames struct {
	enumName  Qualifier
	usedNames UsedNames
}

func (ns *enumNames) Name() Qualifier         { return ns.enumName }
func (ns *enumNames) getUsedNames() UsedNames { return ns.usedNames }

type ValueNames interface {
	Names
}

type valueNames struct {
	valueName Qualifier
}

func (ns *valueNames) Name() Qualifier { return ns.valueName }

type Dart struct {
	files      map[pgs.File]*fileNames
	messages   map[pgs.Message]*messageNames
	fields     map[pgs.Field]*fieldNames
	oneofs     map[pgs.OneOf]*oneOfNames
	enums      map[pgs.Enum]*enumNames
	enumValues map[pgs.EnumValue]*valueNames
}

func NewDart() *Dart {
	return &Dart{
		files:      make(map[pgs.File]*fileNames),
		messages:   make(map[pgs.Message]*messageNames),
		fields:     make(map[pgs.Field]*fieldNames),
		oneofs:     make(map[pgs.OneOf]*oneOfNames),
		enums:      make(map[pgs.Enum]*enumNames),
		enumValues: make(map[pgs.EnumValue]*valueNames),
	}
}

// NameOf returns dart name of entity. Return computed class name of File.
func (d *Dart) NameOf(e pgs.Entity) Qualifier           { return d.namesOf(e).Name() }
func (d *Dart) NamesOf(e pgs.Entity) Names              { return d.namesOf(e) }
func (d *Dart) FileNames(e pgs.File) FileNames          { return d.namesOfFile(e) }
func (d *Dart) MessageNames(e pgs.Message) MessageNames { return d.namesOfMessage(e) }
func (d *Dart) FieldNames(e pgs.Field) FieldNames       { return d.namesOfField(e) }
func (d *Dart) OneOfNames(e pgs.OneOf) OneOfNames       { return d.namesOfOneof(e) }
func (d *Dart) EnumNames(e pgs.Enum) EnumNames          { return d.namesOfEnum(e) }
func (d *Dart) ValueNames(e pgs.EnumValue) ValueNames   { return d.namesOfEnumValue(e) }

func (d *Dart) namesOf(e pgs.Entity) Names {
	switch i := e.(type) {
	case pgs.File:
		return d.namesOfFile(i)
	case pgs.Message:
		return d.namesOfMessage(i)
	case pgs.Field:
		return d.namesOfField(i)
	case pgs.OneOf:
		return d.namesOfOneof(i)
	case pgs.Enum:
		return d.namesOfEnum(i)
	case pgs.EnumValue:
		return d.namesOfEnumValue(i)
	}
	return nil
}

func (d *Dart) namesOfFile(i pgs.File) *fileNames {
	nty, ok := d.files[i]
	if !ok {
		fileName := i.InputPath().BaseName()
		typeName := Qualifier(strings.ReplaceAll(fileName, ".", "_")).ToCamel()
		nty = &fileNames{
			entityName: typeName,
			l10nName:   typeName + "Localizations",
			usedNames:  usedTopLevelNames(),
		}
		d.files[i] = nty
	}
	return nty
}

func (d *Dart) namesOfMessage(i pgs.Message) *messageNames {
	nty, ok := d.messages[i]
	if !ok {
		parent := d.namesOf(i.Parent()).(ScopedNames)
		name := messageOrEnumClassName(Qualifier(i.Name()), parent.getUsedNames(), parent.Name())
		parent.getUsedNames().Add(name)

		var reserved map[Qualifier]bool // compute?
		existingNames := reservedMemberNames()
		existingNames.AddMap(reserved)

		nty = &messageNames{
			messageName:    name, // check dart_options.dart_name?
			zeroName:       "zero" + name,
			validatorName:  name + "Validator",
			formInputsName: name + "FormInputs",
			usedNames:      existingNames,
		}
		d.messages[i] = nty
	}
	return nty
}

func (d *Dart) namesOfOneof(i pgs.OneOf) *oneOfNames {
	nty, ok := d.oneofs[i]
	if !ok {
		parent := d.namesOfMessage(i.Message())
		oneofNameVariants := func(name Qualifier) []Qualifier {
			return []Qualifier{
				_defaultWhichMethodName(name),
				_defaultClearMethodName(name),
			}
		}
		oneofName := disambiguateName(Qualifier(i.Name()).ToCamel(),
			parent.usedNames, new(defaultSuffixes), oneofNameVariants)
		parent.usedNames.Add(oneofName)

		f := d.namesOfFile(i.File())
		oneofEnumName := oneofEnumClassName(Qualifier(i.Name()),
			f.usedNames,
			parent.Name())
		f.usedNames.Add(oneofEnumName)

		enumMapName := disambiguateName(
			Qualifier(fmt.Sprintf("_$%sByTag", oneofEnumName)),
			parent.usedNames, new(defaultSuffixes),
			nil)
		parent.usedNames.Add(enumMapName)

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

func (d *Dart) namesOfField(i pgs.Field) *fieldNames {
	nty, ok := d.fields[i]
	if !ok {
		parent := d.namesOfMessage(i.Message())
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
		name := disambiguateName(_fieldMethodSuffix(i),
			parent.usedNames, suffix, generateNameVariants)
		parent.usedNames.Add(name)

		var ensureMethodName Qualifier
		if i.Type().IsEmbed() {
			ensureMethodName = _defaultEnsureMethodName(name)
		}

		nty = &fieldNames{
			messageName:      d.NameOf(i.Message()),
			fieldName:        _defaultFieldName(name),
			hasMethodName:    _defaultHasMethodName(name),
			clearMethodName:  _defaultClearMethodName(name),
			ensureMethodName: ensureMethodName,
		}
		d.fields[i] = nty
	}
	return nty
}

func (d *Dart) namesOfEnum(i pgs.Enum) *enumNames {
	nty, ok := d.enums[i]
	if !ok {
		parent := d.namesOf(i.Parent()).(ScopedNames)
		name := messageOrEnumClassName(Qualifier(i.Name()), parent.getUsedNames(), parent.Name())
		parent.getUsedNames().Add(name)

		nty = &enumNames{
			enumName:  name,
			usedNames: reservedEnumNames(),
		}
		d.enums[i] = nty
	}
	return nty
}

func (d *Dart) namesOfEnumValue(i pgs.EnumValue) *valueNames {
	nty, ok := d.enumValues[i]
	if !ok {
		parent := d.namesOfEnum(i.Enum())
		name := disambiguateName(avoidInitialUnderscore(Qualifier(i.Name())),
			parent.usedNames, new(enumSuffixes), nil)
		parent.usedNames.Add(name)

		nty = &valueNames{
			valueName: name,
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

func _defaultFieldName(fieldMethodSuffix Qualifier) Qualifier {
	return Qualifier(strings.ToLower(string(fieldMethodSuffix[:1]))) + fieldMethodSuffix[1:]
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

func _fieldMethodSuffix(field pgs.Field) Qualifier {
	name := Qualifier(field.Name()).ToCamel()
	return name
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
