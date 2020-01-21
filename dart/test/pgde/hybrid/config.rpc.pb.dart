///
//  Generated code. Do not modify.
//  source: hybrid/config.rpc.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'config.pb.dart' as $2;
import '../pgde/error/error.pb.dart' as $3;

enum GetRouterItemResponse_Is {
  data, 
  error, 
  notSet
}

class GetRouterItemResponse extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, GetRouterItemResponse_Is> _GetRouterItemResponse_IsByTag = {
    1 : GetRouterItemResponse_Is.data,
    2 : GetRouterItemResponse_Is.error,
    0 : GetRouterItemResponse_Is.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('GetRouterItemResponse', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..aOM<$2.RouterItem>(1, 'data', subBuilder: $2.RouterItem.create)
    ..aOM<$3.BackendError>(2, 'error', subBuilder: $3.BackendError.create)
    ..hasRequiredFields = false
  ;

  GetRouterItemResponse._() : super();
  factory GetRouterItemResponse() => create();
  factory GetRouterItemResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory GetRouterItemResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  GetRouterItemResponse clone() => GetRouterItemResponse()..mergeFromMessage(this);
  GetRouterItemResponse copyWith(void Function(GetRouterItemResponse) updates) => super.copyWith((message) => updates(message as GetRouterItemResponse));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static GetRouterItemResponse create() => GetRouterItemResponse._();
  GetRouterItemResponse createEmptyInstance() => create();
  static $pb.PbList<GetRouterItemResponse> createRepeated() => $pb.PbList<GetRouterItemResponse>();
  @$core.pragma('dart2js:noInline')
  static GetRouterItemResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<GetRouterItemResponse>(create);
  static GetRouterItemResponse _defaultInstance;

  GetRouterItemResponse_Is whichIs() => _GetRouterItemResponse_IsByTag[$_whichOneof(0)];
  void clearIs() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $2.RouterItem get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($2.RouterItem v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $2.RouterItem ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $3.BackendError get error => $_getN(1);
  @$pb.TagNumber(2)
  set error($3.BackendError v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasError() => $_has(1);
  @$pb.TagNumber(2)
  void clearError() => clearField(2);
  @$pb.TagNumber(2)
  $3.BackendError ensureError() => $_ensure(1);
}

class ModifyRouterItemRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ModifyRouterItemRequest', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<$2.RouterItem>(1, 'data', subBuilder: $2.RouterItem.create)
    ..p<$core.int>(2, 'fields', $pb.PbFieldType.P3)
    ..hasRequiredFields = false
  ;

  ModifyRouterItemRequest._() : super();
  factory ModifyRouterItemRequest() => create();
  factory ModifyRouterItemRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ModifyRouterItemRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ModifyRouterItemRequest clone() => ModifyRouterItemRequest()..mergeFromMessage(this);
  ModifyRouterItemRequest copyWith(void Function(ModifyRouterItemRequest) updates) => super.copyWith((message) => updates(message as ModifyRouterItemRequest));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ModifyRouterItemRequest create() => ModifyRouterItemRequest._();
  ModifyRouterItemRequest createEmptyInstance() => create();
  static $pb.PbList<ModifyRouterItemRequest> createRepeated() => $pb.PbList<ModifyRouterItemRequest>();
  @$core.pragma('dart2js:noInline')
  static ModifyRouterItemRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ModifyRouterItemRequest>(create);
  static ModifyRouterItemRequest _defaultInstance;

  @$pb.TagNumber(1)
  $2.RouterItem get data => $_getN(0);
  @$pb.TagNumber(1)
  set data($2.RouterItem v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasData() => $_has(0);
  @$pb.TagNumber(1)
  void clearData() => clearField(1);
  @$pb.TagNumber(1)
  $2.RouterItem ensureData() => $_ensure(0);

  @$pb.TagNumber(2)
  $core.List<$core.int> get fields => $_getList(1);
}

