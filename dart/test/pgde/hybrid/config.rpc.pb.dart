///
//  Generated code. Do not modify.
//  source: hybrid/config.rpc.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'config.pb.dart' as $3;
import '../pgde/error/error.pb.dart' as $2;

class BasicPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('BasicPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.Basic>(1, 'data', subBuilder: $3.Basic.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  BasicPayload_GetResp._() : super();
  factory BasicPayload_GetResp() => create();
  factory BasicPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BasicPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  BasicPayload_GetResp clone() => BasicPayload_GetResp()..mergeFromMessage(this);
  BasicPayload_GetResp copyWith(void Function(BasicPayload_GetResp) updates) => super.copyWith((message) => updates(message as BasicPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BasicPayload_GetResp create() => BasicPayload_GetResp._();
  BasicPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<BasicPayload_GetResp> createRepeated() => $pb.PbList<BasicPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static BasicPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BasicPayload_GetResp>(create);
  static BasicPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.Basic get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.Basic v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.Basic ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class BasicPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('BasicPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  BasicPayload._() : super();
  factory BasicPayload() => create();
  factory BasicPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory BasicPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  BasicPayload clone() => BasicPayload()..mergeFromMessage(this);
  BasicPayload copyWith(void Function(BasicPayload) updates) => super.copyWith((message) => updates(message as BasicPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static BasicPayload create() => BasicPayload._();
  BasicPayload createEmptyInstance() => create();
  static $pb.PbList<BasicPayload> createRepeated() => $pb.PbList<BasicPayload>();
  @$core.pragma('dart2js:noInline')
  static BasicPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<BasicPayload>(create);
  static BasicPayload _defaultInstance;
}

class IpfsPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.Ipfs>(1, 'data', subBuilder: $3.Ipfs.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IpfsPayload_GetResp._() : super();
  factory IpfsPayload_GetResp() => create();
  factory IpfsPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsPayload_GetResp clone() => IpfsPayload_GetResp()..mergeFromMessage(this);
  IpfsPayload_GetResp copyWith(void Function(IpfsPayload_GetResp) updates) => super.copyWith((message) => updates(message as IpfsPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsPayload_GetResp create() => IpfsPayload_GetResp._();
  IpfsPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<IpfsPayload_GetResp> createRepeated() => $pb.PbList<IpfsPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static IpfsPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsPayload_GetResp>(create);
  static IpfsPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.Ipfs get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.Ipfs v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.Ipfs ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IpfsPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IpfsPayload._() : super();
  factory IpfsPayload() => create();
  factory IpfsPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsPayload clone() => IpfsPayload()..mergeFromMessage(this);
  IpfsPayload copyWith(void Function(IpfsPayload) updates) => super.copyWith((message) => updates(message as IpfsPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsPayload create() => IpfsPayload._();
  IpfsPayload createEmptyInstance() => create();
  static $pb.PbList<IpfsPayload> createRepeated() => $pb.PbList<IpfsPayload>();
  @$core.pragma('dart2js:noInline')
  static IpfsPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsPayload>(create);
  static IpfsPayload _defaultInstance;
}

class LogPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  LogPayload_SrcId._() : super();
  factory LogPayload_SrcId() => create();
  factory LogPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload_SrcId clone() => LogPayload_SrcId()..mergeFromMessage(this);
  LogPayload_SrcId copyWith(void Function(LogPayload_SrcId) updates) => super.copyWith((message) => updates(message as LogPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload_SrcId create() => LogPayload_SrcId._();
  LogPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<LogPayload_SrcId> createRepeated() => $pb.PbList<LogPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static LogPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload_SrcId>(create);
  static LogPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class LogPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.Log>(1, 'data', subBuilder: $3.Log.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  LogPayload_GetResp._() : super();
  factory LogPayload_GetResp() => create();
  factory LogPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload_GetResp clone() => LogPayload_GetResp()..mergeFromMessage(this);
  LogPayload_GetResp copyWith(void Function(LogPayload_GetResp) updates) => super.copyWith((message) => updates(message as LogPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload_GetResp create() => LogPayload_GetResp._();
  LogPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<LogPayload_GetResp> createRepeated() => $pb.PbList<LogPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static LogPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload_GetResp>(create);
  static LogPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.Log get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.Log v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.Log ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class LogPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.Log_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.Log_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  LogPayload_View_SrcResp._() : super();
  factory LogPayload_View_SrcResp() => create();
  factory LogPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload_View_SrcResp clone() => LogPayload_View_SrcResp()..mergeFromMessage(this);
  LogPayload_View_SrcResp copyWith(void Function(LogPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as LogPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload_View_SrcResp create() => LogPayload_View_SrcResp._();
  LogPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<LogPayload_View_SrcResp> createRepeated() => $pb.PbList<LogPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static LogPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload_View_SrcResp>(create);
  static LogPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.Log_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class LogPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.Log_View>(1, 'data', subBuilder: $3.Log_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  LogPayload_View_AddResp._() : super();
  factory LogPayload_View_AddResp() => create();
  factory LogPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload_View_AddResp clone() => LogPayload_View_AddResp()..mergeFromMessage(this);
  LogPayload_View_AddResp copyWith(void Function(LogPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as LogPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload_View_AddResp create() => LogPayload_View_AddResp._();
  LogPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<LogPayload_View_AddResp> createRepeated() => $pb.PbList<LogPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static LogPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload_View_AddResp>(create);
  static LogPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.Log_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.Log_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.Log_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class LogPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  LogPayload_View._() : super();
  factory LogPayload_View() => create();
  factory LogPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload_View clone() => LogPayload_View()..mergeFromMessage(this);
  LogPayload_View copyWith(void Function(LogPayload_View) updates) => super.copyWith((message) => updates(message as LogPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload_View create() => LogPayload_View._();
  LogPayload_View createEmptyInstance() => create();
  static $pb.PbList<LogPayload_View> createRepeated() => $pb.PbList<LogPayload_View>();
  @$core.pragma('dart2js:noInline')
  static LogPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload_View>(create);
  static LogPayload_View _defaultInstance;
}

class LogPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('LogPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  LogPayload._() : super();
  factory LogPayload() => create();
  factory LogPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory LogPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  LogPayload clone() => LogPayload()..mergeFromMessage(this);
  LogPayload copyWith(void Function(LogPayload) updates) => super.copyWith((message) => updates(message as LogPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static LogPayload create() => LogPayload._();
  LogPayload createEmptyInstance() => create();
  static $pb.PbList<LogPayload> createRepeated() => $pb.PbList<LogPayload>();
  @$core.pragma('dart2js:noInline')
  static LogPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<LogPayload>(create);
  static LogPayload _defaultInstance;
}

class IpfsServerPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload_SrcId._() : super();
  factory IpfsServerPayload_SrcId() => create();
  factory IpfsServerPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload_SrcId clone() => IpfsServerPayload_SrcId()..mergeFromMessage(this);
  IpfsServerPayload_SrcId copyWith(void Function(IpfsServerPayload_SrcId) updates) => super.copyWith((message) => updates(message as IpfsServerPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_SrcId create() => IpfsServerPayload_SrcId._();
  IpfsServerPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload_SrcId> createRepeated() => $pb.PbList<IpfsServerPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload_SrcId>(create);
  static IpfsServerPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IpfsServerPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.IpfsServer>(1, 'data', subBuilder: $3.IpfsServer.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload_GetResp._() : super();
  factory IpfsServerPayload_GetResp() => create();
  factory IpfsServerPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload_GetResp clone() => IpfsServerPayload_GetResp()..mergeFromMessage(this);
  IpfsServerPayload_GetResp copyWith(void Function(IpfsServerPayload_GetResp) updates) => super.copyWith((message) => updates(message as IpfsServerPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_GetResp create() => IpfsServerPayload_GetResp._();
  IpfsServerPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload_GetResp> createRepeated() => $pb.PbList<IpfsServerPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload_GetResp>(create);
  static IpfsServerPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.IpfsServer get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.IpfsServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.IpfsServer ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IpfsServerPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.IpfsServer_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.IpfsServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload_View_SrcResp._() : super();
  factory IpfsServerPayload_View_SrcResp() => create();
  factory IpfsServerPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload_View_SrcResp clone() => IpfsServerPayload_View_SrcResp()..mergeFromMessage(this);
  IpfsServerPayload_View_SrcResp copyWith(void Function(IpfsServerPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as IpfsServerPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View_SrcResp create() => IpfsServerPayload_View_SrcResp._();
  IpfsServerPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload_View_SrcResp> createRepeated() => $pb.PbList<IpfsServerPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload_View_SrcResp>(create);
  static IpfsServerPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.IpfsServer_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IpfsServerPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.IpfsServer_View>(1, 'data', subBuilder: $3.IpfsServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload_View_AddResp._() : super();
  factory IpfsServerPayload_View_AddResp() => create();
  factory IpfsServerPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload_View_AddResp clone() => IpfsServerPayload_View_AddResp()..mergeFromMessage(this);
  IpfsServerPayload_View_AddResp copyWith(void Function(IpfsServerPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as IpfsServerPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View_AddResp create() => IpfsServerPayload_View_AddResp._();
  IpfsServerPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload_View_AddResp> createRepeated() => $pb.PbList<IpfsServerPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload_View_AddResp>(create);
  static IpfsServerPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.IpfsServer_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.IpfsServer_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.IpfsServer_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IpfsServerPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload_View._() : super();
  factory IpfsServerPayload_View() => create();
  factory IpfsServerPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload_View clone() => IpfsServerPayload_View()..mergeFromMessage(this);
  IpfsServerPayload_View copyWith(void Function(IpfsServerPayload_View) updates) => super.copyWith((message) => updates(message as IpfsServerPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View create() => IpfsServerPayload_View._();
  IpfsServerPayload_View createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload_View> createRepeated() => $pb.PbList<IpfsServerPayload_View>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload_View>(create);
  static IpfsServerPayload_View _defaultInstance;
}

class IpfsServerPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServerPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IpfsServerPayload._() : super();
  factory IpfsServerPayload() => create();
  factory IpfsServerPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServerPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServerPayload clone() => IpfsServerPayload()..mergeFromMessage(this);
  IpfsServerPayload copyWith(void Function(IpfsServerPayload) updates) => super.copyWith((message) => updates(message as IpfsServerPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload create() => IpfsServerPayload._();
  IpfsServerPayload createEmptyInstance() => create();
  static $pb.PbList<IpfsServerPayload> createRepeated() => $pb.PbList<IpfsServerPayload>();
  @$core.pragma('dart2js:noInline')
  static IpfsServerPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServerPayload>(create);
  static IpfsServerPayload _defaultInstance;
}

class FileServerPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  FileServerPayload_SrcId._() : super();
  factory FileServerPayload_SrcId() => create();
  factory FileServerPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload_SrcId clone() => FileServerPayload_SrcId()..mergeFromMessage(this);
  FileServerPayload_SrcId copyWith(void Function(FileServerPayload_SrcId) updates) => super.copyWith((message) => updates(message as FileServerPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_SrcId create() => FileServerPayload_SrcId._();
  FileServerPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload_SrcId> createRepeated() => $pb.PbList<FileServerPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload_SrcId>(create);
  static FileServerPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class FileServerPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.FileServer>(1, 'data', subBuilder: $3.FileServer.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  FileServerPayload_GetResp._() : super();
  factory FileServerPayload_GetResp() => create();
  factory FileServerPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload_GetResp clone() => FileServerPayload_GetResp()..mergeFromMessage(this);
  FileServerPayload_GetResp copyWith(void Function(FileServerPayload_GetResp) updates) => super.copyWith((message) => updates(message as FileServerPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_GetResp create() => FileServerPayload_GetResp._();
  FileServerPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload_GetResp> createRepeated() => $pb.PbList<FileServerPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload_GetResp>(create);
  static FileServerPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.FileServer get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.FileServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.FileServer ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class FileServerPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.FileServer_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.FileServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  FileServerPayload_View_SrcResp._() : super();
  factory FileServerPayload_View_SrcResp() => create();
  factory FileServerPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload_View_SrcResp clone() => FileServerPayload_View_SrcResp()..mergeFromMessage(this);
  FileServerPayload_View_SrcResp copyWith(void Function(FileServerPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as FileServerPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View_SrcResp create() => FileServerPayload_View_SrcResp._();
  FileServerPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload_View_SrcResp> createRepeated() => $pb.PbList<FileServerPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload_View_SrcResp>(create);
  static FileServerPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.FileServer_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class FileServerPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.FileServer_View>(1, 'data', subBuilder: $3.FileServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  FileServerPayload_View_AddResp._() : super();
  factory FileServerPayload_View_AddResp() => create();
  factory FileServerPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload_View_AddResp clone() => FileServerPayload_View_AddResp()..mergeFromMessage(this);
  FileServerPayload_View_AddResp copyWith(void Function(FileServerPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as FileServerPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View_AddResp create() => FileServerPayload_View_AddResp._();
  FileServerPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload_View_AddResp> createRepeated() => $pb.PbList<FileServerPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload_View_AddResp>(create);
  static FileServerPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.FileServer_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.FileServer_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.FileServer_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class FileServerPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  FileServerPayload_View._() : super();
  factory FileServerPayload_View() => create();
  factory FileServerPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload_View clone() => FileServerPayload_View()..mergeFromMessage(this);
  FileServerPayload_View copyWith(void Function(FileServerPayload_View) updates) => super.copyWith((message) => updates(message as FileServerPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View create() => FileServerPayload_View._();
  FileServerPayload_View createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload_View> createRepeated() => $pb.PbList<FileServerPayload_View>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload_View>(create);
  static FileServerPayload_View _defaultInstance;
}

class FileServerPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServerPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  FileServerPayload._() : super();
  factory FileServerPayload() => create();
  factory FileServerPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServerPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServerPayload clone() => FileServerPayload()..mergeFromMessage(this);
  FileServerPayload copyWith(void Function(FileServerPayload) updates) => super.copyWith((message) => updates(message as FileServerPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServerPayload create() => FileServerPayload._();
  FileServerPayload createEmptyInstance() => create();
  static $pb.PbList<FileServerPayload> createRepeated() => $pb.PbList<FileServerPayload>();
  @$core.pragma('dart2js:noInline')
  static FileServerPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServerPayload>(create);
  static FileServerPayload _defaultInstance;
}

class HttpProxyServerPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload_SrcId._() : super();
  factory HttpProxyServerPayload_SrcId() => create();
  factory HttpProxyServerPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload_SrcId clone() => HttpProxyServerPayload_SrcId()..mergeFromMessage(this);
  HttpProxyServerPayload_SrcId copyWith(void Function(HttpProxyServerPayload_SrcId) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_SrcId create() => HttpProxyServerPayload_SrcId._();
  HttpProxyServerPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload_SrcId> createRepeated() => $pb.PbList<HttpProxyServerPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload_SrcId>(create);
  static HttpProxyServerPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class HttpProxyServerPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.HttpProxyServer>(1, 'data', subBuilder: $3.HttpProxyServer.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload_GetResp._() : super();
  factory HttpProxyServerPayload_GetResp() => create();
  factory HttpProxyServerPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload_GetResp clone() => HttpProxyServerPayload_GetResp()..mergeFromMessage(this);
  HttpProxyServerPayload_GetResp copyWith(void Function(HttpProxyServerPayload_GetResp) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_GetResp create() => HttpProxyServerPayload_GetResp._();
  HttpProxyServerPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload_GetResp> createRepeated() => $pb.PbList<HttpProxyServerPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload_GetResp>(create);
  static HttpProxyServerPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.HttpProxyServer get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.HttpProxyServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.HttpProxyServer ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class HttpProxyServerPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.HttpProxyServer_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.HttpProxyServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload_View_SrcResp._() : super();
  factory HttpProxyServerPayload_View_SrcResp() => create();
  factory HttpProxyServerPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload_View_SrcResp clone() => HttpProxyServerPayload_View_SrcResp()..mergeFromMessage(this);
  HttpProxyServerPayload_View_SrcResp copyWith(void Function(HttpProxyServerPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View_SrcResp create() => HttpProxyServerPayload_View_SrcResp._();
  HttpProxyServerPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload_View_SrcResp> createRepeated() => $pb.PbList<HttpProxyServerPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload_View_SrcResp>(create);
  static HttpProxyServerPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.HttpProxyServer_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class HttpProxyServerPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.HttpProxyServer_View>(1, 'data', subBuilder: $3.HttpProxyServer_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload_View_AddResp._() : super();
  factory HttpProxyServerPayload_View_AddResp() => create();
  factory HttpProxyServerPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload_View_AddResp clone() => HttpProxyServerPayload_View_AddResp()..mergeFromMessage(this);
  HttpProxyServerPayload_View_AddResp copyWith(void Function(HttpProxyServerPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View_AddResp create() => HttpProxyServerPayload_View_AddResp._();
  HttpProxyServerPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload_View_AddResp> createRepeated() => $pb.PbList<HttpProxyServerPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload_View_AddResp>(create);
  static HttpProxyServerPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.HttpProxyServer_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.HttpProxyServer_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.HttpProxyServer_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class HttpProxyServerPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload_View._() : super();
  factory HttpProxyServerPayload_View() => create();
  factory HttpProxyServerPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload_View clone() => HttpProxyServerPayload_View()..mergeFromMessage(this);
  HttpProxyServerPayload_View copyWith(void Function(HttpProxyServerPayload_View) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View create() => HttpProxyServerPayload_View._();
  HttpProxyServerPayload_View createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload_View> createRepeated() => $pb.PbList<HttpProxyServerPayload_View>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload_View>(create);
  static HttpProxyServerPayload_View _defaultInstance;
}

class HttpProxyServerPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServerPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  HttpProxyServerPayload._() : super();
  factory HttpProxyServerPayload() => create();
  factory HttpProxyServerPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServerPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServerPayload clone() => HttpProxyServerPayload()..mergeFromMessage(this);
  HttpProxyServerPayload copyWith(void Function(HttpProxyServerPayload) updates) => super.copyWith((message) => updates(message as HttpProxyServerPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload create() => HttpProxyServerPayload._();
  HttpProxyServerPayload createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServerPayload> createRepeated() => $pb.PbList<HttpProxyServerPayload>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServerPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServerPayload>(create);
  static HttpProxyServerPayload _defaultInstance;
}

class ServerViewPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'ipfs', $pb.PbFieldType.OU3)
    ..a<$core.int>(2, 'file', $pb.PbFieldType.OU3)
    ..a<$core.int>(3, 'http', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_SrcId._() : super();
  factory ServerViewPayload_SrcId() => create();
  factory ServerViewPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_SrcId clone() => ServerViewPayload_SrcId()..mergeFromMessage(this);
  ServerViewPayload_SrcId copyWith(void Function(ServerViewPayload_SrcId) updates) => super.copyWith((message) => updates(message as ServerViewPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcId create() => ServerViewPayload_SrcId._();
  ServerViewPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_SrcId> createRepeated() => $pb.PbList<ServerViewPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_SrcId>(create);
  static ServerViewPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get ipfs => $_getIZ(0);
  @$pb.TagNumber(1)
  set ipfs($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasIpfs() => $_has(0);
  @$pb.TagNumber(1)
  void clearIpfs() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get file => $_getIZ(1);
  @$pb.TagNumber(2)
  set file($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get http => $_getIZ(2);
  @$pb.TagNumber(3)
  set http($core.int v) { $_setUnsignedInt32(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasHttp() => $_has(2);
  @$pb.TagNumber(3)
  void clearHttp() => clearField(3);
}

class ServerViewPayload_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.ServerView>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.ServerView.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_SrcResp._() : super();
  factory ServerViewPayload_SrcResp() => create();
  factory ServerViewPayload_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_SrcResp clone() => ServerViewPayload_SrcResp()..mergeFromMessage(this);
  ServerViewPayload_SrcResp copyWith(void Function(ServerViewPayload_SrcResp) updates) => super.copyWith((message) => updates(message as ServerViewPayload_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcResp create() => ServerViewPayload_SrcResp._();
  ServerViewPayload_SrcResp createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_SrcResp> createRepeated() => $pb.PbList<ServerViewPayload_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_SrcResp>(create);
  static ServerViewPayload_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.ServerView> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class ServerViewPayload_Group extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.Group', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.IpfsServer_View>(1, 'ipfs', $pb.PbFieldType.PM, subBuilder: $3.IpfsServer_View.create)
    ..pc<$3.FileServer_View>(2, 'file', $pb.PbFieldType.PM, subBuilder: $3.FileServer_View.create)
    ..pc<$3.HttpProxyServer_View>(3, 'http', $pb.PbFieldType.PM, subBuilder: $3.HttpProxyServer_View.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_Group._() : super();
  factory ServerViewPayload_Group() => create();
  factory ServerViewPayload_Group.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_Group.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_Group clone() => ServerViewPayload_Group()..mergeFromMessage(this);
  ServerViewPayload_Group copyWith(void Function(ServerViewPayload_Group) updates) => super.copyWith((message) => updates(message as ServerViewPayload_Group));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_Group create() => ServerViewPayload_Group._();
  ServerViewPayload_Group createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_Group> createRepeated() => $pb.PbList<ServerViewPayload_Group>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_Group getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_Group>(create);
  static ServerViewPayload_Group _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.IpfsServer_View> get ipfs => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$3.FileServer_View> get file => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$3.HttpProxyServer_View> get http => $_getList(2);
}

class ServerViewPayload_GroupResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.GroupResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<ServerViewPayload_Group>(1, 'data', subBuilder: ServerViewPayload_Group.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_GroupResp._() : super();
  factory ServerViewPayload_GroupResp() => create();
  factory ServerViewPayload_GroupResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_GroupResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_GroupResp clone() => ServerViewPayload_GroupResp()..mergeFromMessage(this);
  ServerViewPayload_GroupResp copyWith(void Function(ServerViewPayload_GroupResp) updates) => super.copyWith((message) => updates(message as ServerViewPayload_GroupResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_GroupResp create() => ServerViewPayload_GroupResp._();
  ServerViewPayload_GroupResp createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_GroupResp> createRepeated() => $pb.PbList<ServerViewPayload_GroupResp>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_GroupResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_GroupResp>(create);
  static ServerViewPayload_GroupResp _defaultInstance;

  @$pb.TagNumber(1)
  ServerViewPayload_Group get data => $_getN(0);
  @$pb.TagNumber(1)
  set data(ServerViewPayload_Group v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  ServerViewPayload_Group ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

enum ServerViewPayload_SrcElement_Is {
  ipfs, 
  file, 
  http, 
  notSet
}

class ServerViewPayload_SrcElement extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ServerViewPayload_SrcElement_Is> _ServerViewPayload_SrcElement_IsByTag = {
    1 : ServerViewPayload_SrcElement_Is.ipfs,
    2 : ServerViewPayload_SrcElement_Is.file,
    3 : ServerViewPayload_SrcElement_Is.http,
    0 : ServerViewPayload_SrcElement_Is.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.SrcElement', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3])
    ..aOM<$3.IpfsServer>(1, 'ipfs', subBuilder: $3.IpfsServer.create)
    ..aOM<$3.FileServer>(2, 'file', subBuilder: $3.FileServer.create)
    ..aOM<$3.HttpProxyServer>(3, 'http', subBuilder: $3.HttpProxyServer.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_SrcElement._() : super();
  factory ServerViewPayload_SrcElement() => create();
  factory ServerViewPayload_SrcElement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_SrcElement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_SrcElement clone() => ServerViewPayload_SrcElement()..mergeFromMessage(this);
  ServerViewPayload_SrcElement copyWith(void Function(ServerViewPayload_SrcElement) updates) => super.copyWith((message) => updates(message as ServerViewPayload_SrcElement));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcElement create() => ServerViewPayload_SrcElement._();
  ServerViewPayload_SrcElement createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_SrcElement> createRepeated() => $pb.PbList<ServerViewPayload_SrcElement>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_SrcElement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_SrcElement>(create);
  static ServerViewPayload_SrcElement _defaultInstance;

  ServerViewPayload_SrcElement_Is whichIs() => _ServerViewPayload_SrcElement_IsByTag[$_whichOneof(0)];
  void clearIs() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $3.IpfsServer get ipfs => $_getN(0);
  @$pb.TagNumber(1)
  set ipfs($3.IpfsServer v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasIpfs() => $_has(0);
  @$pb.TagNumber(1)
  void clearIpfs() => clearField(1);
  @$pb.TagNumber(1)
  $3.IpfsServer ensureIpfs() => $_ensure(0);

  @$pb.TagNumber(2)
  $3.FileServer get file => $_getN(1);
  @$pb.TagNumber(2)
  set file($3.FileServer v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
  @$pb.TagNumber(2)
  $3.FileServer ensureFile() => $_ensure(1);

  @$pb.TagNumber(3)
  $3.HttpProxyServer get http => $_getN(2);
  @$pb.TagNumber(3)
  set http($3.HttpProxyServer v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasHttp() => $_has(2);
  @$pb.TagNumber(3)
  void clearHttp() => clearField(3);
  @$pb.TagNumber(3)
  $3.HttpProxyServer ensureHttp() => $_ensure(2);
}

class ServerViewPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<ServerViewPayload_SrcElement>(1, 'data', subBuilder: ServerViewPayload_SrcElement.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_GetResp._() : super();
  factory ServerViewPayload_GetResp() => create();
  factory ServerViewPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_GetResp clone() => ServerViewPayload_GetResp()..mergeFromMessage(this);
  ServerViewPayload_GetResp copyWith(void Function(ServerViewPayload_GetResp) updates) => super.copyWith((message) => updates(message as ServerViewPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_GetResp create() => ServerViewPayload_GetResp._();
  ServerViewPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_GetResp> createRepeated() => $pb.PbList<ServerViewPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_GetResp>(create);
  static ServerViewPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  ServerViewPayload_SrcElement get data => $_getN(0);
  @$pb.TagNumber(1)
  set data(ServerViewPayload_SrcElement v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  ServerViewPayload_SrcElement ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class ServerViewPayload_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.ServerView>(1, 'data', subBuilder: $3.ServerView.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload_AddResp._() : super();
  factory ServerViewPayload_AddResp() => create();
  factory ServerViewPayload_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload_AddResp clone() => ServerViewPayload_AddResp()..mergeFromMessage(this);
  ServerViewPayload_AddResp copyWith(void Function(ServerViewPayload_AddResp) updates) => super.copyWith((message) => updates(message as ServerViewPayload_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_AddResp create() => ServerViewPayload_AddResp._();
  ServerViewPayload_AddResp createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload_AddResp> createRepeated() => $pb.PbList<ServerViewPayload_AddResp>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload_AddResp>(create);
  static ServerViewPayload_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.ServerView get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.ServerView v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.ServerView ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class ServerViewPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerViewPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  ServerViewPayload._() : super();
  factory ServerViewPayload() => create();
  factory ServerViewPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerViewPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerViewPayload clone() => ServerViewPayload()..mergeFromMessage(this);
  ServerViewPayload copyWith(void Function(ServerViewPayload) updates) => super.copyWith((message) => updates(message as ServerViewPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload create() => ServerViewPayload._();
  ServerViewPayload createEmptyInstance() => create();
  static $pb.PbList<ServerViewPayload> createRepeated() => $pb.PbList<ServerViewPayload>();
  @$core.pragma('dart2js:noInline')
  static ServerViewPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerViewPayload>(create);
  static ServerViewPayload _defaultInstance;
}

class AdpRouterPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload_SrcId._() : super();
  factory AdpRouterPayload_SrcId() => create();
  factory AdpRouterPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload_SrcId clone() => AdpRouterPayload_SrcId()..mergeFromMessage(this);
  AdpRouterPayload_SrcId copyWith(void Function(AdpRouterPayload_SrcId) updates) => super.copyWith((message) => updates(message as AdpRouterPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_SrcId create() => AdpRouterPayload_SrcId._();
  AdpRouterPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload_SrcId> createRepeated() => $pb.PbList<AdpRouterPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload_SrcId>(create);
  static AdpRouterPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class AdpRouterPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.AdpRouter>(1, 'data', subBuilder: $3.AdpRouter.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload_GetResp._() : super();
  factory AdpRouterPayload_GetResp() => create();
  factory AdpRouterPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload_GetResp clone() => AdpRouterPayload_GetResp()..mergeFromMessage(this);
  AdpRouterPayload_GetResp copyWith(void Function(AdpRouterPayload_GetResp) updates) => super.copyWith((message) => updates(message as AdpRouterPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_GetResp create() => AdpRouterPayload_GetResp._();
  AdpRouterPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload_GetResp> createRepeated() => $pb.PbList<AdpRouterPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload_GetResp>(create);
  static AdpRouterPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.AdpRouter get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.AdpRouter v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.AdpRouter ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class AdpRouterPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.AdpRouter_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.AdpRouter_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload_View_SrcResp._() : super();
  factory AdpRouterPayload_View_SrcResp() => create();
  factory AdpRouterPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload_View_SrcResp clone() => AdpRouterPayload_View_SrcResp()..mergeFromMessage(this);
  AdpRouterPayload_View_SrcResp copyWith(void Function(AdpRouterPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as AdpRouterPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View_SrcResp create() => AdpRouterPayload_View_SrcResp._();
  AdpRouterPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload_View_SrcResp> createRepeated() => $pb.PbList<AdpRouterPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload_View_SrcResp>(create);
  static AdpRouterPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.AdpRouter_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class AdpRouterPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.AdpRouter_View>(1, 'data', subBuilder: $3.AdpRouter_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload_View_AddResp._() : super();
  factory AdpRouterPayload_View_AddResp() => create();
  factory AdpRouterPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload_View_AddResp clone() => AdpRouterPayload_View_AddResp()..mergeFromMessage(this);
  AdpRouterPayload_View_AddResp copyWith(void Function(AdpRouterPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as AdpRouterPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View_AddResp create() => AdpRouterPayload_View_AddResp._();
  AdpRouterPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload_View_AddResp> createRepeated() => $pb.PbList<AdpRouterPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload_View_AddResp>(create);
  static AdpRouterPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.AdpRouter_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.AdpRouter_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.AdpRouter_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class AdpRouterPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload_View._() : super();
  factory AdpRouterPayload_View() => create();
  factory AdpRouterPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload_View clone() => AdpRouterPayload_View()..mergeFromMessage(this);
  AdpRouterPayload_View copyWith(void Function(AdpRouterPayload_View) updates) => super.copyWith((message) => updates(message as AdpRouterPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View create() => AdpRouterPayload_View._();
  AdpRouterPayload_View createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload_View> createRepeated() => $pb.PbList<AdpRouterPayload_View>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload_View>(create);
  static AdpRouterPayload_View _defaultInstance;
}

class AdpRouterPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouterPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  AdpRouterPayload._() : super();
  factory AdpRouterPayload() => create();
  factory AdpRouterPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouterPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouterPayload clone() => AdpRouterPayload()..mergeFromMessage(this);
  AdpRouterPayload copyWith(void Function(AdpRouterPayload) updates) => super.copyWith((message) => updates(message as AdpRouterPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload create() => AdpRouterPayload._();
  AdpRouterPayload createEmptyInstance() => create();
  static $pb.PbList<AdpRouterPayload> createRepeated() => $pb.PbList<AdpRouterPayload>();
  @$core.pragma('dart2js:noInline')
  static AdpRouterPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouterPayload>(create);
  static AdpRouterPayload _defaultInstance;
}

class IPNetRouterPayload_SrcId extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload.SrcId', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload_SrcId._() : super();
  factory IPNetRouterPayload_SrcId() => create();
  factory IPNetRouterPayload_SrcId.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload_SrcId.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload_SrcId clone() => IPNetRouterPayload_SrcId()..mergeFromMessage(this);
  IPNetRouterPayload_SrcId copyWith(void Function(IPNetRouterPayload_SrcId) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload_SrcId));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_SrcId create() => IPNetRouterPayload_SrcId._();
  IPNetRouterPayload_SrcId createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload_SrcId> createRepeated() => $pb.PbList<IPNetRouterPayload_SrcId>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_SrcId getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload_SrcId>(create);
  static IPNetRouterPayload_SrcId _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
}

class IPNetRouterPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.IPNetRouter>(1, 'data', subBuilder: $3.IPNetRouter.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload_GetResp._() : super();
  factory IPNetRouterPayload_GetResp() => create();
  factory IPNetRouterPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload_GetResp clone() => IPNetRouterPayload_GetResp()..mergeFromMessage(this);
  IPNetRouterPayload_GetResp copyWith(void Function(IPNetRouterPayload_GetResp) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_GetResp create() => IPNetRouterPayload_GetResp._();
  IPNetRouterPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload_GetResp> createRepeated() => $pb.PbList<IPNetRouterPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload_GetResp>(create);
  static IPNetRouterPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.IPNetRouter get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.IPNetRouter v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.IPNetRouter ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IPNetRouterPayload_View_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload.View.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.IPNetRouter_View>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.IPNetRouter_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload_View_SrcResp._() : super();
  factory IPNetRouterPayload_View_SrcResp() => create();
  factory IPNetRouterPayload_View_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload_View_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload_View_SrcResp clone() => IPNetRouterPayload_View_SrcResp()..mergeFromMessage(this);
  IPNetRouterPayload_View_SrcResp copyWith(void Function(IPNetRouterPayload_View_SrcResp) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload_View_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View_SrcResp create() => IPNetRouterPayload_View_SrcResp._();
  IPNetRouterPayload_View_SrcResp createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload_View_SrcResp> createRepeated() => $pb.PbList<IPNetRouterPayload_View_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload_View_SrcResp>(create);
  static IPNetRouterPayload_View_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.IPNetRouter_View> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IPNetRouterPayload_View_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload.View.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.IPNetRouter_View>(1, 'data', subBuilder: $3.IPNetRouter_View.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload_View_AddResp._() : super();
  factory IPNetRouterPayload_View_AddResp() => create();
  factory IPNetRouterPayload_View_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload_View_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload_View_AddResp clone() => IPNetRouterPayload_View_AddResp()..mergeFromMessage(this);
  IPNetRouterPayload_View_AddResp copyWith(void Function(IPNetRouterPayload_View_AddResp) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload_View_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View_AddResp create() => IPNetRouterPayload_View_AddResp._();
  IPNetRouterPayload_View_AddResp createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload_View_AddResp> createRepeated() => $pb.PbList<IPNetRouterPayload_View_AddResp>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload_View_AddResp>(create);
  static IPNetRouterPayload_View_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.IPNetRouter_View get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.IPNetRouter_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.IPNetRouter_View ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class IPNetRouterPayload_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload_View._() : super();
  factory IPNetRouterPayload_View() => create();
  factory IPNetRouterPayload_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload_View clone() => IPNetRouterPayload_View()..mergeFromMessage(this);
  IPNetRouterPayload_View copyWith(void Function(IPNetRouterPayload_View) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View create() => IPNetRouterPayload_View._();
  IPNetRouterPayload_View createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload_View> createRepeated() => $pb.PbList<IPNetRouterPayload_View>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload_View>(create);
  static IPNetRouterPayload_View _defaultInstance;
}

class IPNetRouterPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouterPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  IPNetRouterPayload._() : super();
  factory IPNetRouterPayload() => create();
  factory IPNetRouterPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouterPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouterPayload clone() => IPNetRouterPayload()..mergeFromMessage(this);
  IPNetRouterPayload copyWith(void Function(IPNetRouterPayload) updates) => super.copyWith((message) => updates(message as IPNetRouterPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload create() => IPNetRouterPayload._();
  IPNetRouterPayload createEmptyInstance() => create();
  static $pb.PbList<IPNetRouterPayload> createRepeated() => $pb.PbList<IPNetRouterPayload>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouterPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouterPayload>(create);
  static IPNetRouterPayload _defaultInstance;
}

class RouterUseViewPayload_SrcResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.SrcResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.RouterUseView>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.RouterUseView.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_SrcResp._() : super();
  factory RouterUseViewPayload_SrcResp() => create();
  factory RouterUseViewPayload_SrcResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_SrcResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_SrcResp clone() => RouterUseViewPayload_SrcResp()..mergeFromMessage(this);
  RouterUseViewPayload_SrcResp copyWith(void Function(RouterUseViewPayload_SrcResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_SrcResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcResp create() => RouterUseViewPayload_SrcResp._();
  RouterUseViewPayload_SrcResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_SrcResp> createRepeated() => $pb.PbList<RouterUseViewPayload_SrcResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_SrcResp>(create);
  static RouterUseViewPayload_SrcResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.RouterUseView> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class RouterUseViewPayload_Group extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.Group', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.AdpRouter_View>(1, 'adp', $pb.PbFieldType.PM, subBuilder: $3.AdpRouter_View.create)
    ..pc<$3.IPNetRouter_View>(2, 'ipnet', $pb.PbFieldType.PM, subBuilder: $3.IPNetRouter_View.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_Group._() : super();
  factory RouterUseViewPayload_Group() => create();
  factory RouterUseViewPayload_Group.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_Group.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_Group clone() => RouterUseViewPayload_Group()..mergeFromMessage(this);
  RouterUseViewPayload_Group copyWith(void Function(RouterUseViewPayload_Group) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_Group));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_Group create() => RouterUseViewPayload_Group._();
  RouterUseViewPayload_Group createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_Group> createRepeated() => $pb.PbList<RouterUseViewPayload_Group>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_Group getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_Group>(create);
  static RouterUseViewPayload_Group _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.AdpRouter_View> get adp => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$3.IPNetRouter_View> get ipnet => $_getList(1);
}

class RouterUseViewPayload_GroupResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.GroupResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<RouterUseViewPayload_Group>(1, 'data', subBuilder: RouterUseViewPayload_Group.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_GroupResp._() : super();
  factory RouterUseViewPayload_GroupResp() => create();
  factory RouterUseViewPayload_GroupResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_GroupResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_GroupResp clone() => RouterUseViewPayload_GroupResp()..mergeFromMessage(this);
  RouterUseViewPayload_GroupResp copyWith(void Function(RouterUseViewPayload_GroupResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_GroupResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_GroupResp create() => RouterUseViewPayload_GroupResp._();
  RouterUseViewPayload_GroupResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_GroupResp> createRepeated() => $pb.PbList<RouterUseViewPayload_GroupResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_GroupResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_GroupResp>(create);
  static RouterUseViewPayload_GroupResp _defaultInstance;

  @$pb.TagNumber(1)
  RouterUseViewPayload_Group get data => $_getN(0);
  @$pb.TagNumber(1)
  set data(RouterUseViewPayload_Group v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  RouterUseViewPayload_Group ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

enum RouterUseViewPayload_SrcElement_Is {
  adp, 
  ipnet, 
  notSet
}

class RouterUseViewPayload_SrcElement extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RouterUseViewPayload_SrcElement_Is> _RouterUseViewPayload_SrcElement_IsByTag = {
    1 : RouterUseViewPayload_SrcElement_Is.adp,
    2 : RouterUseViewPayload_SrcElement_Is.ipnet,
    0 : RouterUseViewPayload_SrcElement_Is.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.SrcElement', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..aOM<$3.AdpRouter>(1, 'adp', subBuilder: $3.AdpRouter.create)
    ..aOM<$3.IPNetRouter>(2, 'ipnet', subBuilder: $3.IPNetRouter.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_SrcElement._() : super();
  factory RouterUseViewPayload_SrcElement() => create();
  factory RouterUseViewPayload_SrcElement.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_SrcElement.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_SrcElement clone() => RouterUseViewPayload_SrcElement()..mergeFromMessage(this);
  RouterUseViewPayload_SrcElement copyWith(void Function(RouterUseViewPayload_SrcElement) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_SrcElement));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcElement create() => RouterUseViewPayload_SrcElement._();
  RouterUseViewPayload_SrcElement createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_SrcElement> createRepeated() => $pb.PbList<RouterUseViewPayload_SrcElement>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcElement getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_SrcElement>(create);
  static RouterUseViewPayload_SrcElement _defaultInstance;

  RouterUseViewPayload_SrcElement_Is whichIs() => _RouterUseViewPayload_SrcElement_IsByTag[$_whichOneof(0)];
  void clearIs() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $3.AdpRouter get adp => $_getN(0);
  @$pb.TagNumber(1)
  set adp($3.AdpRouter v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasAdp() => $_has(0);
  @$pb.TagNumber(1)
  void clearAdp() => clearField(1);
  @$pb.TagNumber(1)
  $3.AdpRouter ensureAdp() => $_ensure(0);

  @$pb.TagNumber(2)
  $3.IPNetRouter get ipnet => $_getN(1);
  @$pb.TagNumber(2)
  set ipnet($3.IPNetRouter v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasIpnet() => $_has(1);
  @$pb.TagNumber(2)
  void clearIpnet() => clearField(2);
  @$pb.TagNumber(2)
  $3.IPNetRouter ensureIpnet() => $_ensure(1);
}

class RouterUseViewPayload_GetResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.GetResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<RouterUseViewPayload_SrcElement>(1, 'data', subBuilder: RouterUseViewPayload_SrcElement.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_GetResp._() : super();
  factory RouterUseViewPayload_GetResp() => create();
  factory RouterUseViewPayload_GetResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_GetResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_GetResp clone() => RouterUseViewPayload_GetResp()..mergeFromMessage(this);
  RouterUseViewPayload_GetResp copyWith(void Function(RouterUseViewPayload_GetResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_GetResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_GetResp create() => RouterUseViewPayload_GetResp._();
  RouterUseViewPayload_GetResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_GetResp> createRepeated() => $pb.PbList<RouterUseViewPayload_GetResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_GetResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_GetResp>(create);
  static RouterUseViewPayload_GetResp _defaultInstance;

  @$pb.TagNumber(1)
  RouterUseViewPayload_SrcElement get data => $_getN(0);
  @$pb.TagNumber(1)
  set data(RouterUseViewPayload_SrcElement v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  RouterUseViewPayload_SrcElement ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class RouterUseViewPayload_AddResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.AddResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.RouterUseView>(1, 'data', subBuilder: $3.RouterUseView.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_AddResp._() : super();
  factory RouterUseViewPayload_AddResp() => create();
  factory RouterUseViewPayload_AddResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_AddResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_AddResp clone() => RouterUseViewPayload_AddResp()..mergeFromMessage(this);
  RouterUseViewPayload_AddResp copyWith(void Function(RouterUseViewPayload_AddResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_AddResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_AddResp create() => RouterUseViewPayload_AddResp._();
  RouterUseViewPayload_AddResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_AddResp> createRepeated() => $pb.PbList<RouterUseViewPayload_AddResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_AddResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_AddResp>(create);
  static RouterUseViewPayload_AddResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.RouterUseView get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.RouterUseView v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.RouterUseView ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class RouterUseViewPayload_SrcIds extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.SrcIds', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.RouterUseView_Id>(1, 'ids', $pb.PbFieldType.PM, subBuilder: $3.RouterUseView_Id.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_SrcIds._() : super();
  factory RouterUseViewPayload_SrcIds() => create();
  factory RouterUseViewPayload_SrcIds.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_SrcIds.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_SrcIds clone() => RouterUseViewPayload_SrcIds()..mergeFromMessage(this);
  RouterUseViewPayload_SrcIds copyWith(void Function(RouterUseViewPayload_SrcIds) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_SrcIds));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcIds create() => RouterUseViewPayload_SrcIds._();
  RouterUseViewPayload_SrcIds createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_SrcIds> createRepeated() => $pb.PbList<RouterUseViewPayload_SrcIds>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SrcIds getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_SrcIds>(create);
  static RouterUseViewPayload_SrcIds _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.RouterUseView_Id> get ids => $_getList(0);
}

class RouterUseViewPayload_DstResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.DstResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$3.RouterUseView_Element>(1, 'data', $pb.PbFieldType.PM, subBuilder: $3.RouterUseView_Element.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_DstResp._() : super();
  factory RouterUseViewPayload_DstResp() => create();
  factory RouterUseViewPayload_DstResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_DstResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_DstResp clone() => RouterUseViewPayload_DstResp()..mergeFromMessage(this);
  RouterUseViewPayload_DstResp copyWith(void Function(RouterUseViewPayload_DstResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_DstResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_DstResp create() => RouterUseViewPayload_DstResp._();
  RouterUseViewPayload_DstResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_DstResp> createRepeated() => $pb.PbList<RouterUseViewPayload_DstResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_DstResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_DstResp>(create);
  static RouterUseViewPayload_DstResp _defaultInstance;

  @$pb.TagNumber(1)
  $core.List<$3.RouterUseView_Element> get data => $_getList(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class RouterUseViewPayload_SelectResp extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload.SelectResp', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$3.RouterUseView_Element>(1, 'data', subBuilder: $3.RouterUseView_Element.create)
    ..aOM<$2.BackendError>(2, 'error', subBuilder: $2.BackendError.create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload_SelectResp._() : super();
  factory RouterUseViewPayload_SelectResp() => create();
  factory RouterUseViewPayload_SelectResp.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload_SelectResp.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload_SelectResp clone() => RouterUseViewPayload_SelectResp()..mergeFromMessage(this);
  RouterUseViewPayload_SelectResp copyWith(void Function(RouterUseViewPayload_SelectResp) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload_SelectResp));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SelectResp create() => RouterUseViewPayload_SelectResp._();
  RouterUseViewPayload_SelectResp createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload_SelectResp> createRepeated() => $pb.PbList<RouterUseViewPayload_SelectResp>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload_SelectResp getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload_SelectResp>(create);
  static RouterUseViewPayload_SelectResp _defaultInstance;

  @$pb.TagNumber(1)
  $3.RouterUseView_Element get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($3.RouterUseView_Element v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $3.RouterUseView_Element ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $2.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($2.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $2.BackendError ensureError() => $_ensure(1);
}

class RouterUseViewPayload extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseViewPayload', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..hasRequiredFields = false
  ;

  RouterUseViewPayload._() : super();
  factory RouterUseViewPayload() => create();
  factory RouterUseViewPayload.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseViewPayload.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseViewPayload clone() => RouterUseViewPayload()..mergeFromMessage(this);
  RouterUseViewPayload copyWith(void Function(RouterUseViewPayload) updates) => super.copyWith((message) => updates(message as RouterUseViewPayload));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload create() => RouterUseViewPayload._();
  RouterUseViewPayload createEmptyInstance() => create();
  static $pb.PbList<RouterUseViewPayload> createRepeated() => $pb.PbList<RouterUseViewPayload>();
  @$core.pragma('dart2js:noInline')
  static RouterUseViewPayload getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseViewPayload>(create);
  static RouterUseViewPayload _defaultInstance;
}

