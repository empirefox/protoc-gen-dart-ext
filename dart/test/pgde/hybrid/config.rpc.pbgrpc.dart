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
import 'config.pb.dart' as $2;
import '../pgde/error/error.pb.dart' as $3;
export 'config.rpc.pb.dart';

class RouterItemServiceClient extends $grpc.Client {
  static final _$getRouterItem =
      $grpc.ClientMethod<$0.Empty, $1.GetRouterItemResponse>(
          '/hybrid.RouterItemService/GetRouterItem',
          ($0.Empty value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $1.GetRouterItemResponse.fromBuffer(value));
  static final _$setRouterItem =
      $grpc.ClientMethod<$2.RouterItem, $3.BackendError>(
          '/hybrid.RouterItemService/SetRouterItem',
          ($2.RouterItem value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $3.BackendError.fromBuffer(value));
  static final _$modifyRouterItem =
      $grpc.ClientMethod<$1.ModifyRouterItemRequest, $3.BackendError>(
          '/hybrid.RouterItemService/ModifyRouterItem',
          ($1.ModifyRouterItemRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) => $3.BackendError.fromBuffer(value));
  static final _$delRouterItem = $grpc.ClientMethod<$0.Empty, $3.BackendError>(
      '/hybrid.RouterItemService/DelRouterItem',
      ($0.Empty value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $3.BackendError.fromBuffer(value));

  RouterItemServiceClient($grpc.ClientChannel channel,
      {$grpc.CallOptions options})
      : super(channel, options: options);

  $grpc.ResponseFuture<$1.GetRouterItemResponse> getRouterItem($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$getRouterItem, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$3.BackendError> setRouterItem($2.RouterItem request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$setRouterItem, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$3.BackendError> modifyRouterItem(
      $1.ModifyRouterItemRequest request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$modifyRouterItem, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }

  $grpc.ResponseFuture<$3.BackendError> delRouterItem($0.Empty request,
      {$grpc.CallOptions options}) {
    final call = $createCall(
        _$delRouterItem, $async.Stream.fromIterable([request]),
        options: options);
    return $grpc.ResponseFuture(call);
  }
}

abstract class RouterItemServiceBase extends $grpc.Service {
  $core.String get $name => 'hybrid.RouterItemService';

  RouterItemServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.Empty, $1.GetRouterItemResponse>(
        'GetRouterItem',
        getRouterItem_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($1.GetRouterItemResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$2.RouterItem, $3.BackendError>(
        'SetRouterItem',
        setRouterItem_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $2.RouterItem.fromBuffer(value),
        ($3.BackendError value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$1.ModifyRouterItemRequest, $3.BackendError>(
        'ModifyRouterItem',
        modifyRouterItem_Pre,
        false,
        false,
        ($core.List<$core.int> value) =>
            $1.ModifyRouterItemRequest.fromBuffer(value),
        ($3.BackendError value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.Empty, $3.BackendError>(
        'DelRouterItem',
        delRouterItem_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.Empty.fromBuffer(value),
        ($3.BackendError value) => value.writeToBuffer()));
  }

  $async.Future<$1.GetRouterItemResponse> getRouterItem_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return getRouterItem(call, await request);
  }

  $async.Future<$3.BackendError> setRouterItem_Pre(
      $grpc.ServiceCall call, $async.Future<$2.RouterItem> request) async {
    return setRouterItem(call, await request);
  }

  $async.Future<$3.BackendError> modifyRouterItem_Pre($grpc.ServiceCall call,
      $async.Future<$1.ModifyRouterItemRequest> request) async {
    return modifyRouterItem(call, await request);
  }

  $async.Future<$3.BackendError> delRouterItem_Pre(
      $grpc.ServiceCall call, $async.Future<$0.Empty> request) async {
    return delRouterItem(call, await request);
  }

  $async.Future<$1.GetRouterItemResponse> getRouterItem(
      $grpc.ServiceCall call, $0.Empty request);
  $async.Future<$3.BackendError> setRouterItem(
      $grpc.ServiceCall call, $2.RouterItem request);
  $async.Future<$3.BackendError> modifyRouterItem(
      $grpc.ServiceCall call, $1.ModifyRouterItemRequest request);
  $async.Future<$3.BackendError> delRouterItem(
      $grpc.ServiceCall call, $0.Empty request);
}
