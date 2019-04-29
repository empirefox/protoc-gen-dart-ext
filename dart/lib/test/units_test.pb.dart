///
//  Generated code. Do not modify.
//  source: test/units_test.proto
///
// ignore_for_file: non_constant_identifier_names,library_prefixes,unused_import

// ignore: UNUSED_SHOWN_NAME
import 'dart:core' show int, bool, double, String, List, Map, override;

import 'package:protobuf/protobuf.dart' as $pb;

import 'units_test.pbenum.dart';

export 'units_test.pbenum.dart';

class MyProfile extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = new $pb.BuilderInfo('MyProfile', package: const $pb.PackageName('test'))
    ..a<int>(1, 'length', $pb.PbFieldType.OU3)
    ..a<int>(2, 'volume', $pb.PbFieldType.OU3)
    ..a<int>(3, 'price1', $pb.PbFieldType.O3)
    ..a<int>(4, 'price2', $pb.PbFieldType.O3)
    ..e<HowManyEnum>(5, 'howMany', $pb.PbFieldType.OE, HowManyEnum.ZERO, HowManyEnum.valueOf, HowManyEnum.values)
    ..a<List<int>>(6, 'sign', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  MyProfile() : super();
  MyProfile.fromBuffer(List<int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  MyProfile.fromJson(String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  MyProfile clone() => new MyProfile()..mergeFromMessage(this);
  MyProfile copyWith(void Function(MyProfile) updates) => super.copyWith((message) => updates(message as MyProfile));
  $pb.BuilderInfo get info_ => _i;
  static MyProfile create() => new MyProfile();
  MyProfile createEmptyInstance() => create();
  static $pb.PbList<MyProfile> createRepeated() => new $pb.PbList<MyProfile>();
  static MyProfile getDefault() => _defaultInstance ??= create()..freeze();
  static MyProfile _defaultInstance;
  static void $checkItem(MyProfile v) {
    if (v is! MyProfile) $pb.checkItemFailed(v, _i.qualifiedMessageName);
  }

  int get length => $_get(0, 0);
  set length(int v) { $_setUnsignedInt32(0, v); }
  bool hasLength() => $_has(0);
  void clearLength() => clearField(1);

  int get volume => $_get(1, 0);
  set volume(int v) { $_setUnsignedInt32(1, v); }
  bool hasVolume() => $_has(1);
  void clearVolume() => clearField(2);

  int get price1 => $_get(2, 0);
  set price1(int v) { $_setSignedInt32(2, v); }
  bool hasPrice1() => $_has(2);
  void clearPrice1() => clearField(3);

  int get price2 => $_get(3, 0);
  set price2(int v) { $_setSignedInt32(3, v); }
  bool hasPrice2() => $_has(3);
  void clearPrice2() => clearField(4);

  HowManyEnum get howMany => $_getN(4);
  set howMany(HowManyEnum v) { setField(5, v); }
  bool hasHowMany() => $_has(4);
  void clearHowMany() => clearField(5);

  List<int> get sign => $_getN(5);
  set sign(List<int> v) { $_setBytes(5, v); }
  bool hasSign() => $_has(5);
  void clearSign() => clearField(6);
}

