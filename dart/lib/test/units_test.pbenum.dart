///
//  Generated code. Do not modify.
//  source: test/units_test.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' show int, dynamic, String, List, Map;
import 'package:protobuf/protobuf.dart' as $pb;

class HowManyEnum extends $pb.ProtobufEnum {
  static const HowManyEnum ZERO = const HowManyEnum._(0, 'ZERO');
  static const HowManyEnum ONE = const HowManyEnum._(1, 'ONE');
  static const HowManyEnum TWO = const HowManyEnum._(2, 'TWO');

  static const List<HowManyEnum> values = const <HowManyEnum> [
    ZERO,
    ONE,
    TWO,
  ];

  static final Map<int, HowManyEnum> _byValue = $pb.ProtobufEnum.initByValue(values);
  static HowManyEnum valueOf(int value) => _byValue[value];
  static void $checkItem(HowManyEnum v) {
    if (v is! HowManyEnum) $pb.checkItemFailed(v, 'HowManyEnum');
  }

  const HowManyEnum._(int v, String n) : super(v, n);
}

class TimeFormatNameRegister extends $pb.ProtobufEnum {
  static const TimeFormatNameRegister ANSIC = const TimeFormatNameRegister._(0, 'ANSIC');
  static const TimeFormatNameRegister RFC822 = const TimeFormatNameRegister._(1, 'RFC822');

  static const List<TimeFormatNameRegister> values = const <TimeFormatNameRegister> [
    ANSIC,
    RFC822,
  ];

  static final Map<int, TimeFormatNameRegister> _byValue = $pb.ProtobufEnum.initByValue(values);
  static TimeFormatNameRegister valueOf(int value) => _byValue[value];
  static void $checkItem(TimeFormatNameRegister v) {
    if (v is! TimeFormatNameRegister) $pb.checkItemFailed(v, 'TimeFormatNameRegister');
  }

  const TimeFormatNameRegister._(int v, String n) : super(v, n);
}

