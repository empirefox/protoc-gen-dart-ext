///
//  Generated code. Do not modify.
//  source: test/units_test.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore_for_file: UNDEFINED_SHOWN_NAME,UNUSED_SHOWN_NAME
import 'dart:core' show String, Type;
import 'package:protobuf/protobuf.dart' show GeneratedMessage, ProtobufEnum;
import 'package:flutter/material.dart' show IconData, Icons;

import './units_test.pbenum.dart';

class UnitsTestUnitsIcons {
  IconData fromEnum(ProtobufEnum v) {
    switch (v.runtimeType) {
      case HowManyEnum:
        return fromHowManyEnum(v);
    }
    assert(false, 'fromEnum() called for unsupported enum "$v"');
    return null;
  }

  IconData fromHowManyEnum(HowManyEnum v) {
    switch (v) {
      case HowManyEnum.ZERO:
        return Icons.star;
    }
    assert(false, 'fromHowManyEnum() called for unsupported enum "$v"');
    return null;
  }
}

abstract class UnitsTestUnitsLocalization {
  const UnitsTestUnitsLocalization();

  String get howManyEnum;
  String get howManyEnum_ZERO;
  String get howManyEnum_ONE;
  String get howManyEnum_TWO;

  String a()=>'aaa';

  String fromEnumType(Type type) {
    switch (type) {
      case HowManyEnum:
        return howManyEnum;
    }
    assert(false, 'from() called for unsupported enum "$type"');
    return null;
  }

  String fromEnum(ProtobufEnum v) {
    switch (v.runtimeType) {
      case HowManyEnum:
        return fromHowManyEnum(v);
    }
    assert(false, 'fromEnum() called for unsupported enum "$v"');
    return null;
  }

  String fromHowManyEnum(HowManyEnum v) {
    switch (v) {
      case HowManyEnum.ZERO:
        return howManyEnum_ZERO;
    }
    assert(false, 'fromHowManyEnum() called for unsupported enum "$v"');
    return null;
  }
}
