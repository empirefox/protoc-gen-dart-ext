///
//  Generated code. Do not modify.
//  source: pgde/error/error.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class BackendError extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('BackendError', package: const $pb.PackageName('pgde.error'), createEmptyInstance: create)
    ..a<$core.int>(1, 'fieldNumber', $pb.PbFieldType.O3, protoName: 'fieldNumber')
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOS(3, 'message')
    ..hasRequiredFields = false
  ;

  BackendError._() : super();
  factory BackendError() => create();
  factory BackendError.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BackendError.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  BackendError clone() => BackendError()..mergeFromMessage(this);
  BackendError copyWith(void Function(BackendError) updates) => super.copyWith((message) => updates(message as BackendError));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BackendError create() => BackendError._();
  BackendError createEmptyInstance() => create();
  static $pb.PbList<BackendError> createRepeated() => $pb.PbList<BackendError>();
  @$core.pragma('dart2js:noInline')
  static BackendError getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BackendError>(create);
  static BackendError _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get fieldNumber => $_getIZ(0);
  @$pb.TagNumber(1)
  set fieldNumber($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFieldNumber() => $_has(0);
  @$pb.TagNumber(1)
  void clearFieldNumber() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get code => $_getIZ(1);
  @$pb.TagNumber(2)
  set code($core.int v) { $_setSignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasCode() => $_has(1);
  @$pb.TagNumber(2)
  void clearCode() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get message => $_getSZ(2);
  @$pb.TagNumber(3)
  set message($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasMessage() => $_has(2);
  @$pb.TagNumber(3)
  void clearMessage() => clearField(3);
}

