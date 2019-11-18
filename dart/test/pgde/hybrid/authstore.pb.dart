///
//  Generated code. Do not modify.
//  source: hybrid/authstore.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:fixnum/fixnum.dart';
import 'package:protobuf/protobuf.dart' as $pb;

class AuthKey extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AuthKey', package: const $pb.PackageName('hybrid'))
    ..a<Int64>(1, 'id', $pb.PbFieldType.OU6, Int64.ZERO)
    ..a<$core.List<$core.int>>(2, 'key', $pb.PbFieldType.OY)
    ..pPS(3, 'tags')
    ..aOS(4, 'desc')
    ..aInt64(5, 'createdAt')
    ..aInt64(6, 'expiresAt')
    ..hasRequiredFields = false
  ;

  AuthKey() : super();
  AuthKey.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  AuthKey.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  AuthKey clone() => AuthKey()..mergeFromMessage(this);
  AuthKey copyWith(void Function(AuthKey) updates) => super.copyWith((message) => updates(message as AuthKey));
  $pb.BuilderInfo get info_ => _i;
  static AuthKey create() => AuthKey();
  AuthKey createEmptyInstance() => create();
  static $pb.PbList<AuthKey> createRepeated() => $pb.PbList<AuthKey>();
  static AuthKey getDefault() => _defaultInstance ??= create()..freeze();
  static AuthKey _defaultInstance;

  Int64 get id => $_getI64(0);
  set id(Int64 v) { $_setInt64(0, v); }
  $core.bool hasId() => $_has(0);
  void clearId() => clearField(1);

  $core.List<$core.int> get key => $_getN(1);
  set key($core.List<$core.int> v) { $_setBytes(1, v); }
  $core.bool hasKey() => $_has(1);
  void clearKey() => clearField(2);

  $core.List<$core.String> get tags => $_getList(2);

  $core.String get desc => $_getS(3, '');
  set desc($core.String v) { $_setString(3, v); }
  $core.bool hasDesc() => $_has(3);
  void clearDesc() => clearField(4);

  Int64 get createdAt => $_getI64(4);
  set createdAt(Int64 v) { $_setInt64(4, v); }
  $core.bool hasCreatedAt() => $_has(4);
  void clearCreatedAt() => clearField(5);

  Int64 get expiresAt => $_getI64(5);
  set expiresAt(Int64 v) { $_setInt64(5, v); }
  $core.bool hasExpiresAt() => $_has(5);
  void clearExpiresAt() => clearField(6);
}

