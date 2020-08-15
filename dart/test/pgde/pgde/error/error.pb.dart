///
//  Generated code. Do not modify.
//  source: pgde/error/error.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

class OperateError extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('OperateError', package: const $pb.PackageName('pgde.error'), createEmptyInstance: create)
    ..a<$core.int>(1, 'code', $pb.PbFieldType.O3)
    ..aOS(2, 'message')
    ..hasRequiredFields = false
  ;

  OperateError._() : super();
  factory OperateError() => create();
  factory OperateError.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory OperateError.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  OperateError clone() => OperateError()..mergeFromMessage(this);
  OperateError copyWith(void Function(OperateError) updates) => super.copyWith((message) => updates(message as OperateError));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static OperateError create() => OperateError._();
  OperateError createEmptyInstance() => create();
  static $pb.PbList<OperateError> createRepeated() => $pb.PbList<OperateError>();
  @$core.pragma('dart2js:noInline')
  static OperateError getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<OperateError>(create);
  static OperateError _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get code => $_getIZ(0);
  @$pb.TagNumber(1)
  set code($core.int v) { $_setSignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasCode() => $_has(0);
  @$pb.TagNumber(1)
  void clearCode() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get message => $_getSZ(1);
  @$pb.TagNumber(2)
  set message($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasMessage() => $_has(1);
  @$pb.TagNumber(2)
  void clearMessage() => clearField(2);
}

class SubmitError extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('SubmitError', package: const $pb.PackageName('pgde.error'), createEmptyInstance: create)
    ..p<$core.int>(1, 'path', $pb.PbFieldType.P3)
    ..a<$core.int>(2, 'code', $pb.PbFieldType.O3)
    ..aOS(3, 'message')
    ..hasRequiredFields = false
  ;

  SubmitError._() : super();
  factory SubmitError() => create();
  factory SubmitError.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory SubmitError.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  SubmitError clone() => SubmitError()..mergeFromMessage(this);
  SubmitError copyWith(void Function(SubmitError) updates) => super.copyWith((message) => updates(message as SubmitError));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static SubmitError create() => SubmitError._();
  SubmitError createEmptyInstance() => create();
  static $pb.PbList<SubmitError> createRepeated() => $pb.PbList<SubmitError>();
  @$core.pragma('dart2js:noInline')
  static SubmitError getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<SubmitError>(create);
  static SubmitError _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$core.int> get path => $_getList(0);

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

