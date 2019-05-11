///
//  Generated code. Do not modify.
//  source: protos/form/form.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' as $core show int, dynamic, String, List, Map;
import 'package:protobuf/protobuf.dart' as $pb;

class UnitPlace extends $pb.ProtobufEnum {
  static const UnitPlace noUnitPlace = UnitPlace._(0, 'noUnitPlace');
  static const UnitPlace suffix = UnitPlace._(1, 'suffix');
  static const UnitPlace labelSuffix = UnitPlace._(2, 'labelSuffix');

  static const $core.List<UnitPlace> values = <UnitPlace> [
    noUnitPlace,
    suffix,
    labelSuffix,
  ];

  static final $core.Map<$core.int, UnitPlace> _byValue = $pb.ProtobufEnum.initByValue(values);
  static UnitPlace valueOf($core.int value) => _byValue[value];

  const UnitPlace._($core.int v, $core.String n) : super(v, n);
}

class TextInput_Type extends $pb.ProtobufEnum {
  static const TextInput_Type autoByType = TextInput_Type._(0, 'autoByType');
  static const TextInput_Type datetime = TextInput_Type._(2, 'datetime');
  static const TextInput_Type emailAddress = TextInput_Type._(3, 'emailAddress');
  static const TextInput_Type multiline = TextInput_Type._(4, 'multiline');
  static const TextInput_Type number = TextInput_Type._(5, 'number');
  static const TextInput_Type phone = TextInput_Type._(6, 'phone');
  static const TextInput_Type text = TextInput_Type._(7, 'text');
  static const TextInput_Type url = TextInput_Type._(8, 'url');

  static const $core.List<TextInput_Type> values = <TextInput_Type> [
    autoByType,
    datetime,
    emailAddress,
    multiline,
    number,
    phone,
    text,
    url,
  ];

  static final $core.Map<$core.int, TextInput_Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TextInput_Type valueOf($core.int value) => _byValue[value];

  const TextInput_Type._($core.int v, $core.String n) : super(v, n);
}

class RadioInput_Type extends $pb.ProtobufEnum {
  static const RadioInput_Type radio = RadioInput_Type._(0, 'radio');
  static const RadioInput_Type dropdown = RadioInput_Type._(1, 'dropdown');
  static const RadioInput_Type segmentedControl = RadioInput_Type._(2, 'segmentedControl');

  static const $core.List<RadioInput_Type> values = <RadioInput_Type> [
    radio,
    dropdown,
    segmentedControl,
  ];

  static final $core.Map<$core.int, RadioInput_Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static RadioInput_Type valueOf($core.int value) => _byValue[value];

  const RadioInput_Type._($core.int v, $core.String n) : super(v, n);
}

class ListInput_Type extends $pb.ProtobufEnum {
  static const ListInput_Type checkbox = ListInput_Type._(0, 'checkbox');
  static const ListInput_Type chips = ListInput_Type._(1, 'chips');

  static const $core.List<ListInput_Type> values = <ListInput_Type> [
    checkbox,
    chips,
  ];

  static final $core.Map<$core.int, ListInput_Type> _byValue = $pb.ProtobufEnum.initByValue(values);
  static ListInput_Type valueOf($core.int value) => _byValue[value];

  const ListInput_Type._($core.int v, $core.String n) : super(v, n);
}

class DatetimePickerInput_DatePickerMode extends $pb.ProtobufEnum {
  static const DatetimePickerInput_DatePickerMode day = DatetimePickerInput_DatePickerMode._(0, 'day');
  static const DatetimePickerInput_DatePickerMode year = DatetimePickerInput_DatePickerMode._(1, 'year');

  static const $core.List<DatetimePickerInput_DatePickerMode> values = <DatetimePickerInput_DatePickerMode> [
    day,
    year,
  ];

  static final $core.Map<$core.int, DatetimePickerInput_DatePickerMode> _byValue = $pb.ProtobufEnum.initByValue(values);
  static DatetimePickerInput_DatePickerMode valueOf($core.int value) => _byValue[value];

  const DatetimePickerInput_DatePickerMode._($core.int v, $core.String n) : super(v, n);
}

