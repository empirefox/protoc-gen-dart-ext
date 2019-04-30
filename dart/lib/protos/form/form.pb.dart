///
//  Generated code. Do not modify.
//  source: protos/form/form.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

import 'color.pb.dart' as $0;

import 'icon.pbenum.dart' as $1;
import 'form.pbenum.dart';

export 'form.pbenum.dart';

class InputOption extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('InputOption', package: const $pb.PackageName('form'))
    ..e<$1.MaterialIcon>(1, 'icon', $pb.PbFieldType.OE, $1.MaterialIcon.noMaterialIcon, $1.MaterialIcon.valueOf, $1.MaterialIcon.values)
    ..hasRequiredFields = false
  ;

  InputOption() : super();
  InputOption.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  InputOption.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  InputOption clone() => InputOption()..mergeFromMessage(this);
  InputOption copyWith(void Function(InputOption) updates) => super.copyWith((message) => updates(message as InputOption));
  $pb.BuilderInfo get info_ => _i;
  static InputOption create() => InputOption();
  InputOption createEmptyInstance() => create();
  static $pb.PbList<InputOption> createRepeated() => $pb.PbList<InputOption>();
  static InputOption getDefault() => _defaultInstance ??= create()..freeze();
  static InputOption _defaultInstance;

  $1.MaterialIcon get icon => $_getN(0);
  set icon($1.MaterialIcon v) { setField(1, v); }
  $core.bool hasIcon() => $_has(0);
  void clearIcon() => clearField(1);
}

enum Field_Type {
  text, 
  typeAhead, 
  signature, 
  number, 
  stepper, 
  slider, 
  rate, 
  radio, 
  boolSwitch, 
  list, 
  datetime, 
  notSet
}

class Field extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Field_Type> _Field_TypeByTag = {
    11 : Field_Type.text,
    12 : Field_Type.typeAhead,
    13 : Field_Type.signature,
    14 : Field_Type.number,
    15 : Field_Type.stepper,
    16 : Field_Type.slider,
    17 : Field_Type.rate,
    18 : Field_Type.radio,
    19 : Field_Type.boolSwitch,
    20 : Field_Type.list,
    21 : Field_Type.datetime,
    0 : Field_Type.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Field', package: const $pb.PackageName('form'))
    ..a<InputDecoration>(1, 'decoration', $pb.PbFieldType.OM, InputDecoration.getDefault, InputDecoration.create)
    ..aOB(2, 'hide')
    ..aOB(3, 'enabled')
    ..aOB(4, 'ensureAgain')
    ..aOB(5, 'emptyAsNoChange')
    ..a<TextInput>(11, 'text', $pb.PbFieldType.OM, TextInput.getDefault, TextInput.create)
    ..a<TypeAheadInput>(12, 'typeAhead', $pb.PbFieldType.OM, TypeAheadInput.getDefault, TypeAheadInput.create)
    ..a<SignatureInput>(13, 'signature', $pb.PbFieldType.OM, SignatureInput.getDefault, SignatureInput.create)
    ..a<NumberInput>(14, 'number', $pb.PbFieldType.OM, NumberInput.getDefault, NumberInput.create)
    ..a<StepperInput>(15, 'stepper', $pb.PbFieldType.OM, StepperInput.getDefault, StepperInput.create)
    ..a<SliderInput>(16, 'slider', $pb.PbFieldType.OM, SliderInput.getDefault, SliderInput.create)
    ..a<RateInput>(17, 'rate', $pb.PbFieldType.OM, RateInput.getDefault, RateInput.create)
    ..a<RadioInput>(18, 'radio', $pb.PbFieldType.OM, RadioInput.getDefault, RadioInput.create)
    ..a<SwitchInput>(19, 'boolSwitch', $pb.PbFieldType.OM, SwitchInput.getDefault, SwitchInput.create)
    ..a<ListInput>(20, 'list', $pb.PbFieldType.OM, ListInput.getDefault, ListInput.create)
    ..a<DatetimePickerInput>(21, 'datetime', $pb.PbFieldType.OM, DatetimePickerInput.getDefault, DatetimePickerInput.create)
    ..oo(0, [11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21])
    ..hasRequiredFields = false
  ;

  Field() : super();
  Field.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Field.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Field clone() => Field()..mergeFromMessage(this);
  Field copyWith(void Function(Field) updates) => super.copyWith((message) => updates(message as Field));
  $pb.BuilderInfo get info_ => _i;
  static Field create() => Field();
  Field createEmptyInstance() => create();
  static $pb.PbList<Field> createRepeated() => $pb.PbList<Field>();
  static Field getDefault() => _defaultInstance ??= create()..freeze();
  static Field _defaultInstance;

  Field_Type whichType() => _Field_TypeByTag[$_whichOneof(0)];
  void clearType() => clearField($_whichOneof(0));

  InputDecoration get decoration => $_getN(0);
  set decoration(InputDecoration v) { setField(1, v); }
  $core.bool hasDecoration() => $_has(0);
  void clearDecoration() => clearField(1);

  $core.bool get hide => $_get(1, false);
  set hide($core.bool v) { $_setBool(1, v); }
  $core.bool hasHide() => $_has(1);
  void clearHide() => clearField(2);

  $core.bool get enabled => $_get(2, false);
  set enabled($core.bool v) { $_setBool(2, v); }
  $core.bool hasEnabled() => $_has(2);
  void clearEnabled() => clearField(3);

  $core.bool get ensureAgain => $_get(3, false);
  set ensureAgain($core.bool v) { $_setBool(3, v); }
  $core.bool hasEnsureAgain() => $_has(3);
  void clearEnsureAgain() => clearField(4);

  $core.bool get emptyAsNoChange => $_get(4, false);
  set emptyAsNoChange($core.bool v) { $_setBool(4, v); }
  $core.bool hasEmptyAsNoChange() => $_has(4);
  void clearEmptyAsNoChange() => clearField(5);

  TextInput get text => $_getN(5);
  set text(TextInput v) { setField(11, v); }
  $core.bool hasText() => $_has(5);
  void clearText() => clearField(11);

  TypeAheadInput get typeAhead => $_getN(6);
  set typeAhead(TypeAheadInput v) { setField(12, v); }
  $core.bool hasTypeAhead() => $_has(6);
  void clearTypeAhead() => clearField(12);

  SignatureInput get signature => $_getN(7);
  set signature(SignatureInput v) { setField(13, v); }
  $core.bool hasSignature() => $_has(7);
  void clearSignature() => clearField(13);

  NumberInput get number => $_getN(8);
  set number(NumberInput v) { setField(14, v); }
  $core.bool hasNumber() => $_has(8);
  void clearNumber() => clearField(14);

  StepperInput get stepper => $_getN(9);
  set stepper(StepperInput v) { setField(15, v); }
  $core.bool hasStepper() => $_has(9);
  void clearStepper() => clearField(15);

  SliderInput get slider => $_getN(10);
  set slider(SliderInput v) { setField(16, v); }
  $core.bool hasSlider() => $_has(10);
  void clearSlider() => clearField(16);

  RateInput get rate => $_getN(11);
  set rate(RateInput v) { setField(17, v); }
  $core.bool hasRate() => $_has(11);
  void clearRate() => clearField(17);

  RadioInput get radio => $_getN(12);
  set radio(RadioInput v) { setField(18, v); }
  $core.bool hasRadio() => $_has(12);
  void clearRadio() => clearField(18);

  SwitchInput get boolSwitch => $_getN(13);
  set boolSwitch(SwitchInput v) { setField(19, v); }
  $core.bool hasBoolSwitch() => $_has(13);
  void clearBoolSwitch() => clearField(19);

  ListInput get list => $_getN(14);
  set list(ListInput v) { setField(20, v); }
  $core.bool hasList() => $_has(14);
  void clearList() => clearField(20);

  DatetimePickerInput get datetime => $_getN(15);
  set datetime(DatetimePickerInput v) { setField(21, v); }
  $core.bool hasDatetime() => $_has(15);
  void clearDatetime() => clearField(21);
}

class TextInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TextInput', package: const $pb.PackageName('form'))
    ..e<TextInput_Type>(1, 'keyboardType', $pb.PbFieldType.OE, TextInput_Type.autoByType, TextInput_Type.valueOf, TextInput_Type.values)
    ..aOB(2, 'obscureText')
    ..aOB(3, 'autocorrect')
    ..a<$core.int>(4, 'maxLines', $pb.PbFieldType.O3)
    ..aOB(5, 'maxLengthEnforced')
    ..hasRequiredFields = false
  ;

  TextInput() : super();
  TextInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  TextInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  TextInput clone() => TextInput()..mergeFromMessage(this);
  TextInput copyWith(void Function(TextInput) updates) => super.copyWith((message) => updates(message as TextInput));
  $pb.BuilderInfo get info_ => _i;
  static TextInput create() => TextInput();
  TextInput createEmptyInstance() => create();
  static $pb.PbList<TextInput> createRepeated() => $pb.PbList<TextInput>();
  static TextInput getDefault() => _defaultInstance ??= create()..freeze();
  static TextInput _defaultInstance;

  TextInput_Type get keyboardType => $_getN(0);
  set keyboardType(TextInput_Type v) { setField(1, v); }
  $core.bool hasKeyboardType() => $_has(0);
  void clearKeyboardType() => clearField(1);

  $core.bool get obscureText => $_get(1, false);
  set obscureText($core.bool v) { $_setBool(1, v); }
  $core.bool hasObscureText() => $_has(1);
  void clearObscureText() => clearField(2);

  $core.bool get autocorrect => $_get(2, false);
  set autocorrect($core.bool v) { $_setBool(2, v); }
  $core.bool hasAutocorrect() => $_has(2);
  void clearAutocorrect() => clearField(3);

  $core.int get maxLines => $_get(3, 0);
  set maxLines($core.int v) { $_setSignedInt32(3, v); }
  $core.bool hasMaxLines() => $_has(3);
  void clearMaxLines() => clearField(4);

  $core.bool get maxLengthEnforced => $_get(4, false);
  set maxLengthEnforced($core.bool v) { $_setBool(4, v); }
  $core.bool hasMaxLengthEnforced() => $_has(4);
  void clearMaxLengthEnforced() => clearField(5);
}

enum TypeAheadInput_Builder {
  codeAccessor, 
  fromEnum, 
  notSet
}

class TypeAheadInput extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, TypeAheadInput_Builder> _TypeAheadInput_BuilderByTag = {
    1 : TypeAheadInput_Builder.codeAccessor,
    2 : TypeAheadInput_Builder.fromEnum,
    0 : TypeAheadInput_Builder.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('TypeAheadInput', package: const $pb.PackageName('form'))
    ..aOS(1, 'codeAccessor')
    ..aOB(2, 'fromEnum')
    ..oo(0, [1, 2])
    ..hasRequiredFields = false
  ;

  TypeAheadInput() : super();
  TypeAheadInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  TypeAheadInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  TypeAheadInput clone() => TypeAheadInput()..mergeFromMessage(this);
  TypeAheadInput copyWith(void Function(TypeAheadInput) updates) => super.copyWith((message) => updates(message as TypeAheadInput));
  $pb.BuilderInfo get info_ => _i;
  static TypeAheadInput create() => TypeAheadInput();
  TypeAheadInput createEmptyInstance() => create();
  static $pb.PbList<TypeAheadInput> createRepeated() => $pb.PbList<TypeAheadInput>();
  static TypeAheadInput getDefault() => _defaultInstance ??= create()..freeze();
  static TypeAheadInput _defaultInstance;

  TypeAheadInput_Builder whichBuilder() => _TypeAheadInput_BuilderByTag[$_whichOneof(0)];
  void clearBuilder() => clearField($_whichOneof(0));

  $core.String get codeAccessor => $_getS(0, '');
  set codeAccessor($core.String v) { $_setString(0, v); }
  $core.bool hasCodeAccessor() => $_has(0);
  void clearCodeAccessor() => clearField(1);

  $core.bool get fromEnum => $_get(1, false);
  set fromEnum($core.bool v) { $_setBool(1, v); }
  $core.bool hasFromEnum() => $_has(1);
  void clearFromEnum() => clearField(2);
}

class SignatureInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SignatureInput', package: const $pb.PackageName('form'))
    ..a<$core.double>(1, 'width', $pb.PbFieldType.OD)
    ..a<$core.double>(2, 'height', $pb.PbFieldType.OD)
    ..a<$0.MaterialColor>(3, 'backgroundColor', $pb.PbFieldType.OM, $0.MaterialColor.getDefault, $0.MaterialColor.create)
    ..a<$0.MaterialColor>(4, 'penColor', $pb.PbFieldType.OM, $0.MaterialColor.getDefault, $0.MaterialColor.create)
    ..a<$core.double>(5, 'penStrokeWidth', $pb.PbFieldType.OD)
    ..hasRequiredFields = false
  ;

  SignatureInput() : super();
  SignatureInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  SignatureInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  SignatureInput clone() => SignatureInput()..mergeFromMessage(this);
  SignatureInput copyWith(void Function(SignatureInput) updates) => super.copyWith((message) => updates(message as SignatureInput));
  $pb.BuilderInfo get info_ => _i;
  static SignatureInput create() => SignatureInput();
  SignatureInput createEmptyInstance() => create();
  static $pb.PbList<SignatureInput> createRepeated() => $pb.PbList<SignatureInput>();
  static SignatureInput getDefault() => _defaultInstance ??= create()..freeze();
  static SignatureInput _defaultInstance;

  $core.double get width => $_getN(0);
  set width($core.double v) { $_setDouble(0, v); }
  $core.bool hasWidth() => $_has(0);
  void clearWidth() => clearField(1);

  $core.double get height => $_getN(1);
  set height($core.double v) { $_setDouble(1, v); }
  $core.bool hasHeight() => $_has(1);
  void clearHeight() => clearField(2);

  $0.MaterialColor get backgroundColor => $_getN(2);
  set backgroundColor($0.MaterialColor v) { setField(3, v); }
  $core.bool hasBackgroundColor() => $_has(2);
  void clearBackgroundColor() => clearField(3);

  $0.MaterialColor get penColor => $_getN(3);
  set penColor($0.MaterialColor v) { setField(4, v); }
  $core.bool hasPenColor() => $_has(3);
  void clearPenColor() => clearField(4);

  $core.double get penStrokeWidth => $_getN(4);
  set penStrokeWidth($core.double v) { $_setDouble(4, v); }
  $core.bool hasPenStrokeWidth() => $_has(4);
  void clearPenStrokeWidth() => clearField(5);
}

class NumberInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('NumberInput', package: const $pb.PackageName('form'))
    ..aOB(1, 'fullUnit')
    ..e<UnitPlace>(2, 'unitPlace', $pb.PbFieldType.OE, UnitPlace.noUnitPlace, UnitPlace.valueOf, UnitPlace.values)
    ..hasRequiredFields = false
  ;

  NumberInput() : super();
  NumberInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  NumberInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  NumberInput clone() => NumberInput()..mergeFromMessage(this);
  NumberInput copyWith(void Function(NumberInput) updates) => super.copyWith((message) => updates(message as NumberInput));
  $pb.BuilderInfo get info_ => _i;
  static NumberInput create() => NumberInput();
  NumberInput createEmptyInstance() => create();
  static $pb.PbList<NumberInput> createRepeated() => $pb.PbList<NumberInput>();
  static NumberInput getDefault() => _defaultInstance ??= create()..freeze();
  static NumberInput _defaultInstance;

  $core.bool get fullUnit => $_get(0, false);
  set fullUnit($core.bool v) { $_setBool(0, v); }
  $core.bool hasFullUnit() => $_has(0);
  void clearFullUnit() => clearField(1);

  UnitPlace get unitPlace => $_getN(1);
  set unitPlace(UnitPlace v) { setField(2, v); }
  $core.bool hasUnitPlace() => $_has(1);
  void clearUnitPlace() => clearField(2);
}

class StepperInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('StepperInput', package: const $pb.PackageName('form'))
    ..aOB(1, 'fullUnit')
    ..e<UnitPlace>(2, 'unitPlace', $pb.PbFieldType.OE, UnitPlace.noUnitPlace, UnitPlace.valueOf, UnitPlace.values)
    ..a<$core.int>(3, 'step', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  StepperInput() : super();
  StepperInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  StepperInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  StepperInput clone() => StepperInput()..mergeFromMessage(this);
  StepperInput copyWith(void Function(StepperInput) updates) => super.copyWith((message) => updates(message as StepperInput));
  $pb.BuilderInfo get info_ => _i;
  static StepperInput create() => StepperInput();
  StepperInput createEmptyInstance() => create();
  static $pb.PbList<StepperInput> createRepeated() => $pb.PbList<StepperInput>();
  static StepperInput getDefault() => _defaultInstance ??= create()..freeze();
  static StepperInput _defaultInstance;

  $core.bool get fullUnit => $_get(0, false);
  set fullUnit($core.bool v) { $_setBool(0, v); }
  $core.bool hasFullUnit() => $_has(0);
  void clearFullUnit() => clearField(1);

  UnitPlace get unitPlace => $_getN(1);
  set unitPlace(UnitPlace v) { setField(2, v); }
  $core.bool hasUnitPlace() => $_has(1);
  void clearUnitPlace() => clearField(2);

  $core.int get step => $_get(2, 0);
  set step($core.int v) { $_setUnsignedInt32(2, v); }
  $core.bool hasStep() => $_has(2);
  void clearStep() => clearField(3);
}

class SliderInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SliderInput', package: const $pb.PackageName('form'))
    ..aOB(1, 'fullUnit')
    ..e<UnitPlace>(2, 'unitPlace', $pb.PbFieldType.OE, UnitPlace.noUnitPlace, UnitPlace.valueOf, UnitPlace.values)
    ..a<$core.int>(3, 'divisions', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  SliderInput() : super();
  SliderInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  SliderInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  SliderInput clone() => SliderInput()..mergeFromMessage(this);
  SliderInput copyWith(void Function(SliderInput) updates) => super.copyWith((message) => updates(message as SliderInput));
  $pb.BuilderInfo get info_ => _i;
  static SliderInput create() => SliderInput();
  SliderInput createEmptyInstance() => create();
  static $pb.PbList<SliderInput> createRepeated() => $pb.PbList<SliderInput>();
  static SliderInput getDefault() => _defaultInstance ??= create()..freeze();
  static SliderInput _defaultInstance;

  $core.bool get fullUnit => $_get(0, false);
  set fullUnit($core.bool v) { $_setBool(0, v); }
  $core.bool hasFullUnit() => $_has(0);
  void clearFullUnit() => clearField(1);

  UnitPlace get unitPlace => $_getN(1);
  set unitPlace(UnitPlace v) { setField(2, v); }
  $core.bool hasUnitPlace() => $_has(1);
  void clearUnitPlace() => clearField(2);

  $core.int get divisions => $_get(2, 0);
  set divisions($core.int v) { $_setUnsignedInt32(2, v); }
  $core.bool hasDivisions() => $_has(2);
  void clearDivisions() => clearField(3);
}

class RateInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RateInput', package: const $pb.PackageName('form'))
    ..e<$1.MaterialIcon>(1, 'icon', $pb.PbFieldType.OE, $1.MaterialIcon.noMaterialIcon, $1.MaterialIcon.valueOf, $1.MaterialIcon.values)
    ..hasRequiredFields = false
  ;

  RateInput() : super();
  RateInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  RateInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  RateInput clone() => RateInput()..mergeFromMessage(this);
  RateInput copyWith(void Function(RateInput) updates) => super.copyWith((message) => updates(message as RateInput));
  $pb.BuilderInfo get info_ => _i;
  static RateInput create() => RateInput();
  RateInput createEmptyInstance() => create();
  static $pb.PbList<RateInput> createRepeated() => $pb.PbList<RateInput>();
  static RateInput getDefault() => _defaultInstance ??= create()..freeze();
  static RateInput _defaultInstance;

  $1.MaterialIcon get icon => $_getN(0);
  set icon($1.MaterialIcon v) { setField(1, v); }
  $core.bool hasIcon() => $_has(0);
  void clearIcon() => clearField(1);
}

enum RadioInput_Builder {
  codeAccessor, 
  fromEnum, 
  notSet
}

class RadioInput extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RadioInput_Builder> _RadioInput_BuilderByTag = {
    2 : RadioInput_Builder.codeAccessor,
    3 : RadioInput_Builder.fromEnum,
    0 : RadioInput_Builder.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RadioInput', package: const $pb.PackageName('form'))
    ..e<RadioInput_Type>(1, 'type', $pb.PbFieldType.OE, RadioInput_Type.radio, RadioInput_Type.valueOf, RadioInput_Type.values)
    ..aOS(2, 'codeAccessor')
    ..aOB(3, 'fromEnum')
    ..oo(0, [2, 3])
    ..hasRequiredFields = false
  ;

  RadioInput() : super();
  RadioInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  RadioInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  RadioInput clone() => RadioInput()..mergeFromMessage(this);
  RadioInput copyWith(void Function(RadioInput) updates) => super.copyWith((message) => updates(message as RadioInput));
  $pb.BuilderInfo get info_ => _i;
  static RadioInput create() => RadioInput();
  RadioInput createEmptyInstance() => create();
  static $pb.PbList<RadioInput> createRepeated() => $pb.PbList<RadioInput>();
  static RadioInput getDefault() => _defaultInstance ??= create()..freeze();
  static RadioInput _defaultInstance;

  RadioInput_Builder whichBuilder() => _RadioInput_BuilderByTag[$_whichOneof(0)];
  void clearBuilder() => clearField($_whichOneof(0));

  RadioInput_Type get type => $_getN(0);
  set type(RadioInput_Type v) { setField(1, v); }
  $core.bool hasType() => $_has(0);
  void clearType() => clearField(1);

  $core.String get codeAccessor => $_getS(1, '');
  set codeAccessor($core.String v) { $_setString(1, v); }
  $core.bool hasCodeAccessor() => $_has(1);
  void clearCodeAccessor() => clearField(2);

  $core.bool get fromEnum => $_get(2, false);
  set fromEnum($core.bool v) { $_setBool(2, v); }
  $core.bool hasFromEnum() => $_has(2);
  void clearFromEnum() => clearField(3);
}

class SwitchInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SwitchInput', package: const $pb.PackageName('form'))
    ..aOB(1, 'checkbox')
    ..hasRequiredFields = false
  ;

  SwitchInput() : super();
  SwitchInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  SwitchInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  SwitchInput clone() => SwitchInput()..mergeFromMessage(this);
  SwitchInput copyWith(void Function(SwitchInput) updates) => super.copyWith((message) => updates(message as SwitchInput));
  $pb.BuilderInfo get info_ => _i;
  static SwitchInput create() => SwitchInput();
  SwitchInput createEmptyInstance() => create();
  static $pb.PbList<SwitchInput> createRepeated() => $pb.PbList<SwitchInput>();
  static SwitchInput getDefault() => _defaultInstance ??= create()..freeze();
  static SwitchInput _defaultInstance;

  $core.bool get checkbox => $_get(0, false);
  set checkbox($core.bool v) { $_setBool(0, v); }
  $core.bool hasCheckbox() => $_has(0);
  void clearCheckbox() => clearField(1);
}

enum ListInput_Builder {
  codeAccessor, 
  fromEnum, 
  notSet
}

class ListInput extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ListInput_Builder> _ListInput_BuilderByTag = {
    2 : ListInput_Builder.codeAccessor,
    3 : ListInput_Builder.fromEnum,
    0 : ListInput_Builder.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ListInput', package: const $pb.PackageName('form'))
    ..e<ListInput_Type>(1, 'type', $pb.PbFieldType.OE, ListInput_Type.checkbox, ListInput_Type.valueOf, ListInput_Type.values)
    ..aOS(2, 'codeAccessor')
    ..aOB(3, 'fromEnum')
    ..oo(0, [2, 3])
    ..hasRequiredFields = false
  ;

  ListInput() : super();
  ListInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  ListInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  ListInput clone() => ListInput()..mergeFromMessage(this);
  ListInput copyWith(void Function(ListInput) updates) => super.copyWith((message) => updates(message as ListInput));
  $pb.BuilderInfo get info_ => _i;
  static ListInput create() => ListInput();
  ListInput createEmptyInstance() => create();
  static $pb.PbList<ListInput> createRepeated() => $pb.PbList<ListInput>();
  static ListInput getDefault() => _defaultInstance ??= create()..freeze();
  static ListInput _defaultInstance;

  ListInput_Builder whichBuilder() => _ListInput_BuilderByTag[$_whichOneof(0)];
  void clearBuilder() => clearField($_whichOneof(0));

  ListInput_Type get type => $_getN(0);
  set type(ListInput_Type v) { setField(1, v); }
  $core.bool hasType() => $_has(0);
  void clearType() => clearField(1);

  $core.String get codeAccessor => $_getS(1, '');
  set codeAccessor($core.String v) { $_setString(1, v); }
  $core.bool hasCodeAccessor() => $_has(1);
  void clearCodeAccessor() => clearField(2);

  $core.bool get fromEnum => $_get(2, false);
  set fromEnum($core.bool v) { $_setBool(2, v); }
  $core.bool hasFromEnum() => $_has(2);
  void clearFromEnum() => clearField(3);
}

class DatetimePickerInput extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('DatetimePickerInput', package: const $pb.PackageName('form'))
    ..e<DatetimePickerInput_Type>(1, 'type', $pb.PbFieldType.OE, DatetimePickerInput_Type.date, DatetimePickerInput_Type.valueOf, DatetimePickerInput_Type.values)
    ..aOS(2, 'format')
    ..hasRequiredFields = false
  ;

  DatetimePickerInput() : super();
  DatetimePickerInput.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  DatetimePickerInput.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  DatetimePickerInput clone() => DatetimePickerInput()..mergeFromMessage(this);
  DatetimePickerInput copyWith(void Function(DatetimePickerInput) updates) => super.copyWith((message) => updates(message as DatetimePickerInput));
  $pb.BuilderInfo get info_ => _i;
  static DatetimePickerInput create() => DatetimePickerInput();
  DatetimePickerInput createEmptyInstance() => create();
  static $pb.PbList<DatetimePickerInput> createRepeated() => $pb.PbList<DatetimePickerInput>();
  static DatetimePickerInput getDefault() => _defaultInstance ??= create()..freeze();
  static DatetimePickerInput _defaultInstance;

  DatetimePickerInput_Type get type => $_getN(0);
  set type(DatetimePickerInput_Type v) { setField(1, v); }
  $core.bool hasType() => $_has(0);
  void clearType() => clearField(1);

  $core.String get format => $_getS(1, '');
  set format($core.String v) { $_setString(1, v); }
  $core.bool hasFormat() => $_has(1);
  void clearFormat() => clearField(2);
}

class InputDecoration extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('InputDecoration', package: const $pb.PackageName('form'))
    ..e<$1.MaterialIcon>(1, 'icon', $pb.PbFieldType.OE, $1.MaterialIcon.noMaterialIcon, $1.MaterialIcon.valueOf, $1.MaterialIcon.values)
    ..e<$1.MaterialIcon>(2, 'prefixIcon', $pb.PbFieldType.OE, $1.MaterialIcon.noMaterialIcon, $1.MaterialIcon.valueOf, $1.MaterialIcon.values)
    ..e<$1.MaterialIcon>(3, 'suffixIcon', $pb.PbFieldType.OE, $1.MaterialIcon.noMaterialIcon, $1.MaterialIcon.valueOf, $1.MaterialIcon.values)
    ..aOB(4, 'hasCounter')
    ..hasRequiredFields = false
  ;

  InputDecoration() : super();
  InputDecoration.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  InputDecoration.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  InputDecoration clone() => InputDecoration()..mergeFromMessage(this);
  InputDecoration copyWith(void Function(InputDecoration) updates) => super.copyWith((message) => updates(message as InputDecoration));
  $pb.BuilderInfo get info_ => _i;
  static InputDecoration create() => InputDecoration();
  InputDecoration createEmptyInstance() => create();
  static $pb.PbList<InputDecoration> createRepeated() => $pb.PbList<InputDecoration>();
  static InputDecoration getDefault() => _defaultInstance ??= create()..freeze();
  static InputDecoration _defaultInstance;

  $1.MaterialIcon get icon => $_getN(0);
  set icon($1.MaterialIcon v) { setField(1, v); }
  $core.bool hasIcon() => $_has(0);
  void clearIcon() => clearField(1);

  $1.MaterialIcon get prefixIcon => $_getN(1);
  set prefixIcon($1.MaterialIcon v) { setField(2, v); }
  $core.bool hasPrefixIcon() => $_has(1);
  void clearPrefixIcon() => clearField(2);

  $1.MaterialIcon get suffixIcon => $_getN(2);
  set suffixIcon($1.MaterialIcon v) { setField(3, v); }
  $core.bool hasSuffixIcon() => $_has(2);
  void clearSuffixIcon() => clearField(3);

  $core.bool get hasCounter => $_get(3, false);
  set hasCounter($core.bool v) { $_setBool(3, v); }
  $core.bool hasHasCounter() => $_has(3);
  void clearHasCounter() => clearField(4);
}

enum Date_Utc {
  unix, 
  rfc3339, 
  notSet
}

class Date extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, Date_Utc> _Date_UtcByTag = {
    1 : Date_Utc.unix,
    2 : Date_Utc.rfc3339,
    0 : Date_Utc.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Date', package: const $pb.PackageName('form'))
    ..aInt64(1, 'unix')
    ..aOS(2, 'rfc3339')
    ..oo(0, [1, 2])
    ..hasRequiredFields = false
  ;

  Date() : super();
  Date.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Date.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Date clone() => Date()..mergeFromMessage(this);
  Date copyWith(void Function(Date) updates) => super.copyWith((message) => updates(message as Date));
  $pb.BuilderInfo get info_ => _i;
  static Date create() => Date();
  Date createEmptyInstance() => create();
  static $pb.PbList<Date> createRepeated() => $pb.PbList<Date>();
  static Date getDefault() => _defaultInstance ??= create()..freeze();
  static Date _defaultInstance;

  Date_Utc whichUtc() => _Date_UtcByTag[$_whichOneof(0)];
  void clearUtc() => clearField($_whichOneof(0));

  Int64 get unix => $_getI64(0);
  set unix(Int64 v) { $_setInt64(0, v); }
  $core.bool hasUnix() => $_has(0);
  void clearUnix() => clearField(1);

  $core.String get rfc3339 => $_getS(1, '');
  set rfc3339($core.String v) { $_setString(1, v); }
  $core.bool hasRfc3339() => $_has(1);
  void clearRfc3339() => clearField(2);
}

class Form {
  static final $pb.Extension defaultField = $pb.Extension<Field>('google.protobuf.EnumOptions', 'defaultField', 919113, $pb.PbFieldType.OM, Field.getDefault, Field.create);
  static final $pb.Extension inputOption = $pb.Extension<InputOption>('google.protobuf.EnumValueOptions', 'inputOption', 919113, $pb.PbFieldType.OM, InputOption.getDefault, InputOption.create);
  static final $pb.Extension entry = $pb.Extension<InputOption>('google.protobuf.MessageOptions', 'entry', 919113, $pb.PbFieldType.OM, InputOption.getDefault, InputOption.create);
  static final $pb.Extension field = $pb.Extension<Field>('google.protobuf.FieldOptions', 'field', 919113, $pb.PbFieldType.OM, Field.getDefault, Field.create);
  static void registerAllExtensions($pb.ExtensionRegistry registry) {
    registry.add(defaultField);
    registry.add(inputOption);
    registry.add(entry);
    registry.add(field);
  }
}

