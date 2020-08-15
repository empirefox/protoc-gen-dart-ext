///
//  Generated code. Do not modify.
//  source: hybrid/config.rpc.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import '../google/protobuf/empty.pb.dart' as $0;
import 'config.rpc.pb.dart' as $1;
import '../pgde/error/error.pb.dart' as $2;
import '../google/protobuf/wrappers.pb.dart' as $3;
import 'config.pb.dart' as $4;
export 'config.rpc.pb.dart';

class IPNetRouterMatchedServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$0.Empty, $1.ServerViewPayload_SrcResp>(
          '/hybrid.IPNetRouterMatchedService/Src',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.ServerViewPayload_SrcResp.fromBuffer(value));
  static final _$select =
      $grpc.ClientMethod<$1.ServerViewPayload_SrcId, $2.SubmitError>(
          '/hybrid.IPNetRouterMatchedService/Select',
          ($1.ServerViewPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));

  IPNetRouterMatchedServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.ServerViewPayload_SrcResp> src($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> select(
      $1.ServerViewPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$select, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class IPNetRouterMatchedServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.IPNetRouterMatchedService';

  IPNetRouterMatchedServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.ServerViewPayload_SrcResp>(
        'Src',
        src_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.ServerViewPayload_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServerViewPayload_SrcId, $2.SubmitError>(
        'Select',
        select_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ServerViewPayload_SrcId.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return src(call, await request);
  }

  $async.Future<$2.SubmitError> select_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcId> request) async {
    return select(call, await request);
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$2.SubmitError> select(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcId request);
}

class IPNetRouterUnmatchedServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$0.Empty, $1.ServerViewPayload_SrcResp>(
          '/hybrid.IPNetRouterUnmatchedService/Src',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.ServerViewPayload_SrcResp.fromBuffer(value));
  static final _$select =
      $grpc.ClientMethod<$1.ServerViewPayload_SrcId, $2.SubmitError>(
          '/hybrid.IPNetRouterUnmatchedService/Select',
          ($1.ServerViewPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));

  IPNetRouterUnmatchedServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.ServerViewPayload_SrcResp> src($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> select(
      $1.ServerViewPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$select, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class IPNetRouterUnmatchedServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.IPNetRouterUnmatchedService';

  IPNetRouterUnmatchedServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.ServerViewPayload_SrcResp>(
        'Src',
        src_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.ServerViewPayload_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServerViewPayload_SrcId, $2.SubmitError>(
        'Select',
        select_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ServerViewPayload_SrcId.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return src(call, await request);
  }

  $async.Future<$2.SubmitError> select_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcId> request) async {
    return select(call, await request);
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$2.SubmitError> select(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcId request);
}

class IPNetRouterFileTestServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$3.UInt32Value, $1.FileServerPayload_View_SrcResp>(
          '/hybrid.IPNetRouterFileTestService/Src',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.FileServerPayload_View_SrcResp.fromBuffer(value));
  static final _$select =
      $grpc.ClientMethod<$1.FileServerPayload_SrcId, $0.Empty>(
          '/hybrid.IPNetRouterFileTestService/Select',
          ($1.FileServerPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$selectMany =
      $grpc.ClientMethod<$1.FileServerPayload_SrcIds, $0.Empty>(
          '/hybrid.IPNetRouterFileTestService/SelectMany',
          ($1.FileServerPayload_SrcIds value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $0.Empty.fromBuffer(value));
  static final _$dst =
      $grpc.ClientMethod<$0.Empty, $1.FileServerPayload_View_SrcResp>(
          '/hybrid.IPNetRouterFileTestService/Dst',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.FileServerPayload_View_SrcResp.fromBuffer(value));

  IPNetRouterFileTestServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.FileServerPayload_View_SrcResp> src(
      $3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.Empty> select($1.FileServerPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$select, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$0.Empty> selectMany($1.FileServerPayload_SrcIds request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$selectMany, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.FileServerPayload_View_SrcResp> dst($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$dst, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class IPNetRouterFileTestServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.IPNetRouterFileTestService';

  IPNetRouterFileTestServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$3.UInt32Value, $1.FileServerPayload_View_SrcResp>(
            'Src',
            src_Pre,
            false,
            false,
            ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
            ($1.FileServerPayload_View_SrcResp value) =>
                value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.FileServerPayload_SrcId, $0.Empty>(
        'Select',
        select_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.FileServerPayload_SrcId.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.FileServerPayload_SrcIds, $0.Empty>(
        'SelectMany',
        selectMany_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.FileServerPayload_SrcIds.fromBuffer(value),
        ($0.Empty value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.FileServerPayload_View_SrcResp>(
        'Dst',
        dst_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.FileServerPayload_View_SrcResp value) => value.writeToBuffer()));
  }

  $async.Future<$1.FileServerPayload_View_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return src(call, await request);
  }

  $async.Future<$0.Empty> select_Pre($grpc.ServiceCall call,
      $async.Future<$1.FileServerPayload_SrcId> request) async {
    return select(call, await request);
  }

  $async.Future<$0.Empty> selectMany_Pre($grpc.ServiceCall call,
      $async.Future<$1.FileServerPayload_SrcIds> request) async {
    return selectMany(call, await request);
  }

  $async.Future<$1.FileServerPayload_View_SrcResp> dst_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return dst(call, await request);
  }

  $async.Future<$1.FileServerPayload_View_SrcResp> src(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$0.Empty> select(
      $grpc.ServiceCall call, $1.FileServerPayload_SrcId request);
  $async.Future<$0.Empty> selectMany(
      $grpc.ServiceCall call, $1.FileServerPayload_SrcIds request);
  $async.Future<$1.FileServerPayload_View_SrcResp> dst(
      $grpc.ServiceCall call, $0.Empty request);
}

class ConfigBasicServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.Empty, $1.BasicPayload_GetResp>(
      '/hybrid.ConfigBasicService/Get',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.BasicPayload_GetResp.fromBuffer(value));

  ConfigBasicServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.BasicPayload_GetResp> get($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigBasicServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigBasicService';

  ConfigBasicServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.BasicPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.BasicPayload_GetResp value) => value.writeToBuffer()));
  }

  $async.Future<$1.BasicPayload_GetResp> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return get(call, await request);
  }

  $async.Future<$1.BasicPayload_GetResp> get(
      $grpc.ServiceCall call, $0.Empty request);
}

class ConfigIpfsServiceClient extends $grpc.Client {
  static final _$get = $grpc.ClientMethod<$0.Empty, $1.IpfsPayload_GetResp>(
      '/hybrid.ConfigIpfsService/Get',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.IpfsPayload_GetResp.fromBuffer(value));

  ConfigIpfsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.IpfsPayload_GetResp> get($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigIpfsServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigIpfsService';

  ConfigIpfsServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.IpfsPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.IpfsPayload_GetResp value) => value.writeToBuffer()));
  }

  $async.Future<$1.IpfsPayload_GetResp> get_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return get(call, await request);
  }

  $async.Future<$1.IpfsPayload_GetResp> get(
      $grpc.ServiceCall call, $0.Empty request);
}

class ConfigLogServiceClient extends $grpc.Client {
  static final _$src = $grpc.ClientMethod<$0.Empty, $1.LogPayload_View_SrcResp>(
      '/hybrid.ConfigLogService/Src',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.LogPayload_View_SrcResp.fromBuffer(value));
  static final _$select =
      $grpc.ClientMethod<$1.LogPayload_SrcId, $2.SubmitError>(
          '/hybrid.ConfigLogService/Select',
          ($1.LogPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));

  ConfigLogServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.LogPayload_View_SrcResp> src($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> select($1.LogPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$select, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigLogServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigLogService';

  ConfigLogServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.LogPayload_View_SrcResp>(
        'Src',
        src_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.LogPayload_View_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.LogPayload_SrcId, $2.SubmitError>(
        'Select',
        select_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.LogPayload_SrcId.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
  }

  $async.Future<$1.LogPayload_View_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return src(call, await request);
  }

  $async.Future<$2.SubmitError> select_Pre($grpc.ServiceCall call,
      $async.Future<$1.LogPayload_SrcId> request) async {
    return select(call, await request);
  }

  $async.Future<$1.LogPayload_View_SrcResp> src(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$2.SubmitError> select(
      $grpc.ServiceCall call, $1.LogPayload_SrcId request);
}

class ConfigLogsServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$3.UInt32Value, $1.LogPayload_View_SrcResp>(
          '/hybrid.ConfigLogsService/Src',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.LogPayload_View_SrcResp.fromBuffer(value));
  static final _$get =
      $grpc.ClientMethod<$1.LogPayload_SrcId, $1.LogPayload_GetResp>(
          '/hybrid.ConfigLogsService/Get',
          ($1.LogPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.LogPayload_GetResp.fromBuffer(value));
  static final _$add =
      $grpc.ClientMethod<$4.Log_View, $1.LogPayload_View_AddResp>(
          '/hybrid.ConfigLogsService/Add',
          ($4.Log_View value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.LogPayload_View_AddResp.fromBuffer(value));
  static final _$save = $grpc.ClientMethod<$4.Log_View, $2.SubmitError>(
      '/hybrid.ConfigLogsService/Save',
      ($4.Log_View value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));
  static final _$remove =
      $grpc.ClientMethod<$1.LogPayload_SrcId, $2.OperateError>(
          '/hybrid.ConfigLogsService/Remove',
          ($1.LogPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.OperateError.fromBuffer(value));

  ConfigLogsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.LogPayload_View_SrcResp> src($3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.LogPayload_GetResp> get($1.LogPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.LogPayload_View_AddResp> add($4.Log_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> save($4.Log_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$save, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.OperateError> remove($1.LogPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigLogsServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigLogsService';

  ConfigLogsServiceBase() {
    $addMethod($grpc.ServiceMethod<$3.UInt32Value, $1.LogPayload_View_SrcResp>(
        'Src',
        src_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
        ($1.LogPayload_View_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.LogPayload_SrcId, $1.LogPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.LogPayload_SrcId.fromBuffer(value),
        ($1.LogPayload_GetResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Log_View, $1.LogPayload_View_AddResp>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Log_View.fromBuffer(value),
        ($1.LogPayload_View_AddResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.Log_View, $2.SubmitError>(
        'Save',
        save_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.Log_View.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.LogPayload_SrcId, $2.OperateError>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $1.LogPayload_SrcId.fromBuffer(value),
        ($2.OperateError value) => value.writeToBuffer()));
  }

  $async.Future<$1.LogPayload_View_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return src(call, await request);
  }

  $async.Future<$1.LogPayload_GetResp> get_Pre($grpc.ServiceCall call,
      $async.Future<$1.LogPayload_SrcId> request) async {
    return get(call, await request);
  }

  $async.Future<$1.LogPayload_View_AddResp> add_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Log_View> request) async {
    return add(call, await request);
  }

  $async.Future<$2.SubmitError> save_Pre(
      $grpc.ServiceCall call, $async.Future<$4.Log_View> request) async {
    return save(call, await request);
  }

  $async.Future<$2.OperateError> remove_Pre($grpc.ServiceCall call,
      $async.Future<$1.LogPayload_SrcId> request) async {
    return remove(call, await request);
  }

  $async.Future<$1.LogPayload_View_SrcResp> src(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$1.LogPayload_GetResp> get(
      $grpc.ServiceCall call, $1.LogPayload_SrcId request);
  $async.Future<$1.LogPayload_View_AddResp> add(
      $grpc.ServiceCall call, $4.Log_View request);
  $async.Future<$2.SubmitError> save(
      $grpc.ServiceCall call, $4.Log_View request);
  $async.Future<$2.OperateError> remove(
      $grpc.ServiceCall call, $1.LogPayload_SrcId request);
}

class ConfigServersServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$3.UInt32Value, $1.ServerViewPayload_SrcResp>(
          '/hybrid.ConfigServersService/Src',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.ServerViewPayload_SrcResp.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$1.ServerViewPayload_SrcId,
          $1.ServerViewPayload_GetResp>(
      '/hybrid.ConfigServersService/Get',
      ($1.ServerViewPayload_SrcId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.ServerViewPayload_GetResp.fromBuffer(value));
  static final _$add = $grpc.ClientMethod<$1.ServerViewPayload_SrcElement,
          $1.ServerViewPayload_AddResp>(
      '/hybrid.ConfigServersService/Add',
      ($1.ServerViewPayload_SrcElement value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.ServerViewPayload_AddResp.fromBuffer(value));
  static final _$save =
      $grpc.ClientMethod<$1.ServerViewPayload_SrcElement, $2.SubmitError>(
          '/hybrid.ConfigServersService/Save',
          ($1.ServerViewPayload_SrcElement value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));
  static final _$remove =
      $grpc.ClientMethod<$1.ServerViewPayload_SrcId, $2.OperateError>(
          '/hybrid.ConfigServersService/Remove',
          ($1.ServerViewPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.OperateError.fromBuffer(value));

  ConfigServersServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.ServerViewPayload_SrcResp> src($3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.ServerViewPayload_GetResp> get(
      $1.ServerViewPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.ServerViewPayload_AddResp> add(
      $1.ServerViewPayload_SrcElement request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> save(
      $1.ServerViewPayload_SrcElement request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$save, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.OperateError> remove(
      $1.ServerViewPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigServersServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigServersService';

  ConfigServersServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$3.UInt32Value, $1.ServerViewPayload_SrcResp>(
            'Src',
            src_Pre,
            false,
            false,
            ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
            ($1.ServerViewPayload_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServerViewPayload_SrcId,
            $1.ServerViewPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ServerViewPayload_SrcId.fromBuffer(value),
        ($1.ServerViewPayload_GetResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServerViewPayload_SrcElement,
            $1.ServerViewPayload_AddResp>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ServerViewPayload_SrcElement.fromBuffer(value),
        ($1.ServerViewPayload_AddResp value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.ServerViewPayload_SrcElement, $2.SubmitError>(
            'Save',
            save_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.ServerViewPayload_SrcElement.fromBuffer(value),
            ($2.SubmitError value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ServerViewPayload_SrcId, $2.OperateError>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ServerViewPayload_SrcId.fromBuffer(value),
        ($2.OperateError value) => value.writeToBuffer()));
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return src(call, await request);
  }

  $async.Future<$1.ServerViewPayload_GetResp> get_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcId> request) async {
    return get(call, await request);
  }

  $async.Future<$1.ServerViewPayload_AddResp> add_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcElement> request) async {
    return add(call, await request);
  }

  $async.Future<$2.SubmitError> save_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcElement> request) async {
    return save(call, await request);
  }

  $async.Future<$2.OperateError> remove_Pre($grpc.ServiceCall call,
      $async.Future<$1.ServerViewPayload_SrcId> request) async {
    return remove(call, await request);
  }

  $async.Future<$1.ServerViewPayload_SrcResp> src(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$1.ServerViewPayload_GetResp> get(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcId request);
  $async.Future<$1.ServerViewPayload_AddResp> add(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcElement request);
  $async.Future<$2.SubmitError> save(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcElement request);
  $async.Future<$2.OperateError> remove(
      $grpc.ServiceCall call, $1.ServerViewPayload_SrcId request);
}

class ConfigAdpsServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$3.UInt32Value, $1.AdpRouterPayload_View_SrcResp>(
          '/hybrid.ConfigAdpsService/Src',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.AdpRouterPayload_View_SrcResp.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$1.AdpRouterPayload_SrcId,
          $1.AdpRouterPayload_GetResp>(
      '/hybrid.ConfigAdpsService/Get',
      ($1.AdpRouterPayload_SrcId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.AdpRouterPayload_GetResp.fromBuffer(value));
  static final _$add =
      $grpc.ClientMethod<$4.AdpRouter_View, $1.AdpRouterPayload_View_AddResp>(
          '/hybrid.ConfigAdpsService/Add',
          ($4.AdpRouter_View value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.AdpRouterPayload_View_AddResp.fromBuffer(value));
  static final _$save = $grpc.ClientMethod<$4.AdpRouter_View, $2.SubmitError>(
      '/hybrid.ConfigAdpsService/Save',
      ($4.AdpRouter_View value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));
  static final _$remove =
      $grpc.ClientMethod<$1.AdpRouterPayload_SrcId, $2.OperateError>(
          '/hybrid.ConfigAdpsService/Remove',
          ($1.AdpRouterPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.OperateError.fromBuffer(value));

  ConfigAdpsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.AdpRouterPayload_View_SrcResp> src(
      $3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.AdpRouterPayload_GetResp> get(
      $1.AdpRouterPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.AdpRouterPayload_View_AddResp> add(
      $4.AdpRouter_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> save($4.AdpRouter_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$save, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.OperateError> remove(
      $1.AdpRouterPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigAdpsServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigAdpsService';

  ConfigAdpsServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$3.UInt32Value, $1.AdpRouterPayload_View_SrcResp>(
            'Src',
            src_Pre,
            false,
            false,
            ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
            ($1.AdpRouterPayload_View_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.AdpRouterPayload_SrcId,
            $1.AdpRouterPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.AdpRouterPayload_SrcId.fromBuffer(value),
        ($1.AdpRouterPayload_GetResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.AdpRouter_View,
            $1.AdpRouterPayload_View_AddResp>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.AdpRouter_View.fromBuffer(value),
        ($1.AdpRouterPayload_View_AddResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.AdpRouter_View, $2.SubmitError>(
        'Save',
        save_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.AdpRouter_View.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.AdpRouterPayload_SrcId, $2.OperateError>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.AdpRouterPayload_SrcId.fromBuffer(value),
        ($2.OperateError value) => value.writeToBuffer()));
  }

  $async.Future<$1.AdpRouterPayload_View_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return src(call, await request);
  }

  $async.Future<$1.AdpRouterPayload_GetResp> get_Pre($grpc.ServiceCall call,
      $async.Future<$1.AdpRouterPayload_SrcId> request) async {
    return get(call, await request);
  }

  $async.Future<$1.AdpRouterPayload_View_AddResp> add_Pre(
      $grpc.ServiceCall call, $async.Future<$4.AdpRouter_View> request) async {
    return add(call, await request);
  }

  $async.Future<$2.SubmitError> save_Pre(
      $grpc.ServiceCall call, $async.Future<$4.AdpRouter_View> request) async {
    return save(call, await request);
  }

  $async.Future<$2.OperateError> remove_Pre($grpc.ServiceCall call,
      $async.Future<$1.AdpRouterPayload_SrcId> request) async {
    return remove(call, await request);
  }

  $async.Future<$1.AdpRouterPayload_View_SrcResp> src(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$1.AdpRouterPayload_GetResp> get(
      $grpc.ServiceCall call, $1.AdpRouterPayload_SrcId request);
  $async.Future<$1.AdpRouterPayload_View_AddResp> add(
      $grpc.ServiceCall call, $4.AdpRouter_View request);
  $async.Future<$2.SubmitError> save(
      $grpc.ServiceCall call, $4.AdpRouter_View request);
  $async.Future<$2.OperateError> remove(
      $grpc.ServiceCall call, $1.AdpRouterPayload_SrcId request);
}

class ConfigIpnetsServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$3.UInt32Value, $1.IPNetRouterPayload_View_SrcResp>(
          '/hybrid.ConfigIpnetsService/Src',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.IPNetRouterPayload_View_SrcResp.fromBuffer(value));
  static final _$get = $grpc.ClientMethod<$1.IPNetRouterPayload_SrcId,
          $1.IPNetRouterPayload_GetResp>(
      '/hybrid.ConfigIpnetsService/Get',
      ($1.IPNetRouterPayload_SrcId value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.IPNetRouterPayload_GetResp.fromBuffer(value));
  static final _$add = $grpc.ClientMethod<$4.IPNetRouter_View,
          $1.IPNetRouterPayload_View_AddResp>(
      '/hybrid.ConfigIpnetsService/Add',
      ($4.IPNetRouter_View value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.IPNetRouterPayload_View_AddResp.fromBuffer(value));
  static final _$save = $grpc.ClientMethod<$4.IPNetRouter_View, $2.SubmitError>(
      '/hybrid.ConfigIpnetsService/Save',
      ($4.IPNetRouter_View value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $2.SubmitError.fromBuffer(value));
  static final _$remove =
      $grpc.ClientMethod<$1.IPNetRouterPayload_SrcId, $2.OperateError>(
          '/hybrid.ConfigIpnetsService/Remove',
          ($1.IPNetRouterPayload_SrcId value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.OperateError.fromBuffer(value));

  ConfigIpnetsServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.IPNetRouterPayload_View_SrcResp> src(
      $3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.IPNetRouterPayload_GetResp> get(
      $1.IPNetRouterPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$get, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.IPNetRouterPayload_View_AddResp> add(
      $4.IPNetRouter_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$add, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.SubmitError> save($4.IPNetRouter_View request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$save, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.OperateError> remove(
      $1.IPNetRouterPayload_SrcId request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigIpnetsServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigIpnetsService';

  ConfigIpnetsServiceBase() {
    $addMethod(
        $grpc.ServiceMethod<$3.UInt32Value, $1.IPNetRouterPayload_View_SrcResp>(
            'Src',
            src_Pre,
            false,
            false,
            ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
            ($1.IPNetRouterPayload_View_SrcResp value) =>
                value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.IPNetRouterPayload_SrcId,
            $1.IPNetRouterPayload_GetResp>(
        'Get',
        get_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.IPNetRouterPayload_SrcId.fromBuffer(value),
        ($1.IPNetRouterPayload_GetResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.IPNetRouter_View,
            $1.IPNetRouterPayload_View_AddResp>(
        'Add',
        add_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.IPNetRouter_View.fromBuffer(value),
        ($1.IPNetRouterPayload_View_AddResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.IPNetRouter_View, $2.SubmitError>(
        'Save',
        save_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.IPNetRouter_View.fromBuffer(value),
        ($2.SubmitError value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$1.IPNetRouterPayload_SrcId, $2.OperateError>(
            'Remove',
            remove_Pre,
            false,
            false,
            ($core.List<$core.int> value) =>
                $1.IPNetRouterPayload_SrcId.fromBuffer(value),
            ($2.OperateError value) => value.writeToBuffer()));
  }

  $async.Future<$1.IPNetRouterPayload_View_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return src(call, await request);
  }

  $async.Future<$1.IPNetRouterPayload_GetResp> get_Pre($grpc.ServiceCall call,
      $async.Future<$1.IPNetRouterPayload_SrcId> request) async {
    return get(call, await request);
  }

  $async.Future<$1.IPNetRouterPayload_View_AddResp> add_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.IPNetRouter_View> request) async {
    return add(call, await request);
  }

  $async.Future<$2.SubmitError> save_Pre($grpc.ServiceCall call,
      $async.Future<$4.IPNetRouter_View> request) async {
    return save(call, await request);
  }

  $async.Future<$2.OperateError> remove_Pre($grpc.ServiceCall call,
      $async.Future<$1.IPNetRouterPayload_SrcId> request) async {
    return remove(call, await request);
  }

  $async.Future<$1.IPNetRouterPayload_View_SrcResp> src(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$1.IPNetRouterPayload_GetResp> get(
      $grpc.ServiceCall call, $1.IPNetRouterPayload_SrcId request);
  $async.Future<$1.IPNetRouterPayload_View_AddResp> add(
      $grpc.ServiceCall call, $4.IPNetRouter_View request);
  $async.Future<$2.SubmitError> save(
      $grpc.ServiceCall call, $4.IPNetRouter_View request);
  $async.Future<$2.OperateError> remove(
      $grpc.ServiceCall call, $1.IPNetRouterPayload_SrcId request);
}

class ConfigRoutersServiceClient extends $grpc.Client {
  static final _$src =
      $grpc.ClientMethod<$0.Empty, $1.RouterUseViewPayload_SrcResp>(
          '/hybrid.ConfigRoutersService/Src',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.RouterUseViewPayload_SrcResp.fromBuffer(value));
  static final _$select = $grpc.ClientMethod<$4.RouterUseView_Id,
          $1.RouterUseViewPayload_SelectResp>(
      '/hybrid.ConfigRoutersService/Select',
      ($4.RouterUseView_Id value) => value.writeToBuffer(),
      ($core.List<$core.int> value) =>
          $1.RouterUseViewPayload_SelectResp.fromBuffer(value));
  static final _$dst =
      $grpc.ClientMethod<$3.UInt32Value, $1.RouterUseViewPayload_DstResp>(
          '/hybrid.ConfigRoutersService/Dst',
          ($3.UInt32Value value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.RouterUseViewPayload_DstResp.fromBuffer(value));
  static final _$remove =
      $grpc.ClientMethod<$4.RouterUseView_Id, $2.OperateError>(
          '/hybrid.ConfigRoutersService/Remove',
          ($4.RouterUseView_Id value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $2.OperateError.fromBuffer(value));

  ConfigRoutersServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.RouterUseViewPayload_SrcResp> src($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$src, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.RouterUseViewPayload_SelectResp> select(
      $4.RouterUseView_Id request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$select, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$1.RouterUseViewPayload_DstResp> dst(
      $3.UInt32Value request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$dst, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$2.OperateError> remove($4.RouterUseView_Id request,
      {$grpc.CallOptions options}) {
    final call = $createCall(_$remove, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class ConfigRoutersServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.ConfigRoutersService';

  ConfigRoutersServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.RouterUseViewPayload_SrcResp>(
        'Src',
        src_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.RouterUseViewPayload_SrcResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RouterUseView_Id,
            $1.RouterUseViewPayload_SelectResp>(
        'Select',
        select_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RouterUseView_Id.fromBuffer(value),
        ($1.RouterUseViewPayload_SelectResp value) => value.writeToBuffer()));
    $addMethod(
        $grpc.ServiceMethod<$3.UInt32Value, $1.RouterUseViewPayload_DstResp>(
            'Dst',
            dst_Pre,
            false,
            false,
            ($core.List<$core.int> value) => $3.UInt32Value.fromBuffer(value),
            ($1.RouterUseViewPayload_DstResp value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$4.RouterUseView_Id, $2.OperateError>(
        'Remove',
        remove_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $4.RouterUseView_Id.fromBuffer(value),
        ($2.OperateError value) => value.writeToBuffer()));
  }

  $async.Future<$1.RouterUseViewPayload_SrcResp> src_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return src(call, await request);
  }

  $async.Future<$1.RouterUseViewPayload_SelectResp> select_Pre(
      $grpc.ServiceCall call,
      $async.Future<$4.RouterUseView_Id> request) async {
    return select(call, await request);
  }

  $async.Future<$1.RouterUseViewPayload_DstResp> dst_Pre(
      $grpc.ServiceCall call, $async.Future<$3.UInt32Value> request) async {
    return dst(call, await request);
  }

  $async.Future<$2.OperateError> remove_Pre($grpc.ServiceCall call,
      $async.Future<$4.RouterUseView_Id> request) async {
    return remove(call, await request);
  }

  $async.Future<$1.RouterUseViewPayload_SrcResp> src(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$1.RouterUseViewPayload_SelectResp> select(
      $grpc.ServiceCall call, $4.RouterUseView_Id request);
  $async.Future<$1.RouterUseViewPayload_DstResp> dst(
      $grpc.ServiceCall call, $3.UInt32Value request);
  $async.Future<$2.OperateError> remove(
      $grpc.ServiceCall call, $4.RouterUseView_Id request);
}
