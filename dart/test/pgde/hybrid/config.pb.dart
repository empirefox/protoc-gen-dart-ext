///
//  Generated code. Do not modify.
//  source: hybrid/config.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

import 'package:fixnum/fixnum.dart' as $fixnum;
import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/wrappers.pb.dart' as $4;

class ConfigTree extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ConfigTree', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOS(1, 'version')
    ..aOS(2, 'rootName', protoName: 'rootName')
    ..aOS(3, 'rootPath', protoName: 'rootPath')
    ..aOS(4, 'configName', protoName: 'configName')
    ..aOS(5, 'configPath', protoName: 'configPath')
    ..aOS(6, 'ipfsName', protoName: 'ipfsName')
    ..aOS(7, 'ipfsPath', protoName: 'ipfsPath')
    ..aOS(8, 'storeName', protoName: 'storeName')
    ..aOS(9, 'storePath', protoName: 'storePath')
    ..aOS(10, 'filesRootName', protoName: 'filesRootName')
    ..aOS(11, 'filesRootPath', protoName: 'filesRootPath')
    ..aOS(12, 'rulesRootName', protoName: 'rulesRootName')
    ..aOS(13, 'rulesRootPath', protoName: 'rulesRootPath')
    ..hasRequiredFields = false
  ;

  ConfigTree._() : super();
  factory ConfigTree() => create();
  factory ConfigTree.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ConfigTree.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ConfigTree clone() => ConfigTree()..mergeFromMessage(this);
  ConfigTree copyWith(void Function(ConfigTree) updates) => super.copyWith((message) => updates(message as ConfigTree));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ConfigTree create() => ConfigTree._();
  ConfigTree createEmptyInstance() => create();
  static $pb.PbList<ConfigTree> createRepeated() => $pb.PbList<ConfigTree>();
  @$core.pragma('dart2js:noInline')
  static ConfigTree getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ConfigTree>(create);
  static ConfigTree _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get version => $_getSZ(0);
  @$pb.TagNumber(1)
  set version($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasVersion() => $_has(0);
  @$pb.TagNumber(1)
  void clearVersion() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get rootName => $_getSZ(1);
  @$pb.TagNumber(2)
  set rootName($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasRootName() => $_has(1);
  @$pb.TagNumber(2)
  void clearRootName() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get rootPath => $_getSZ(2);
  @$pb.TagNumber(3)
  set rootPath($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasRootPath() => $_has(2);
  @$pb.TagNumber(3)
  void clearRootPath() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get configName => $_getSZ(3);
  @$pb.TagNumber(4)
  set configName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasConfigName() => $_has(3);
  @$pb.TagNumber(4)
  void clearConfigName() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get configPath => $_getSZ(4);
  @$pb.TagNumber(5)
  set configPath($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasConfigPath() => $_has(4);
  @$pb.TagNumber(5)
  void clearConfigPath() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get ipfsName => $_getSZ(5);
  @$pb.TagNumber(6)
  set ipfsName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasIpfsName() => $_has(5);
  @$pb.TagNumber(6)
  void clearIpfsName() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get ipfsPath => $_getSZ(6);
  @$pb.TagNumber(7)
  set ipfsPath($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasIpfsPath() => $_has(6);
  @$pb.TagNumber(7)
  void clearIpfsPath() => clearField(7);

  @$pb.TagNumber(8)
  $core.String get storeName => $_getSZ(7);
  @$pb.TagNumber(8)
  set storeName($core.String v) { $_setString(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasStoreName() => $_has(7);
  @$pb.TagNumber(8)
  void clearStoreName() => clearField(8);

  @$pb.TagNumber(9)
  $core.String get storePath => $_getSZ(8);
  @$pb.TagNumber(9)
  set storePath($core.String v) { $_setString(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasStorePath() => $_has(8);
  @$pb.TagNumber(9)
  void clearStorePath() => clearField(9);

  @$pb.TagNumber(10)
  $core.String get filesRootName => $_getSZ(9);
  @$pb.TagNumber(10)
  set filesRootName($core.String v) { $_setString(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasFilesRootName() => $_has(9);
  @$pb.TagNumber(10)
  void clearFilesRootName() => clearField(10);

  @$pb.TagNumber(11)
  $core.String get filesRootPath => $_getSZ(10);
  @$pb.TagNumber(11)
  set filesRootPath($core.String v) { $_setString(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasFilesRootPath() => $_has(10);
  @$pb.TagNumber(11)
  void clearFilesRootPath() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get rulesRootName => $_getSZ(11);
  @$pb.TagNumber(12)
  set rulesRootName($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasRulesRootName() => $_has(11);
  @$pb.TagNumber(12)
  void clearRulesRootName() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get rulesRootPath => $_getSZ(12);
  @$pb.TagNumber(13)
  set rulesRootPath($core.String v) { $_setString(12, v); }
  @$pb.TagNumber(13)
  $core.bool hasRulesRootPath() => $_has(12);
  @$pb.TagNumber(13)
  void clearRulesRootPath() => clearField(13);
}

class Basic extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Basic', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOS(2, 'version')
    ..aOB(3, 'dev')
    ..aOB(4, 'enableBind', protoName: 'enableBind')
    ..a<$core.List<$core.int>>(5, 'bindIp', $pb.PbFieldType.OY, protoName: 'bindIp')
    ..a<$core.int>(6, 'bindPort', $pb.PbFieldType.OU3, protoName: 'bindPort')
    ..a<$core.int>(7, 'flushIntervalMs', $pb.PbFieldType.OU3, protoName: 'flushIntervalMs')
    ..a<$core.List<$core.int>>(8, 'token', $pb.PbFieldType.OY)
    ..hasRequiredFields = false
  ;

  Basic._() : super();
  factory Basic() => create();
  factory Basic.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Basic.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Basic clone() => Basic()..mergeFromMessage(this);
  Basic copyWith(void Function(Basic) updates) => super.copyWith((message) => updates(message as Basic));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Basic create() => Basic._();
  Basic createEmptyInstance() => create();
  static $pb.PbList<Basic> createRepeated() => $pb.PbList<Basic>();
  @$core.pragma('dart2js:noInline')
  static Basic getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Basic>(create);
  static Basic _defaultInstance;

  @$pb.TagNumber(2)
  $core.String get version => $_getSZ(0);
  @$pb.TagNumber(2)
  set version($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasVersion() => $_has(0);
  @$pb.TagNumber(2)
  void clearVersion() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get dev => $_getBF(1);
  @$pb.TagNumber(3)
  set dev($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasDev() => $_has(1);
  @$pb.TagNumber(3)
  void clearDev() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get enableBind => $_getBF(2);
  @$pb.TagNumber(4)
  set enableBind($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(4)
  $core.bool hasEnableBind() => $_has(2);
  @$pb.TagNumber(4)
  void clearEnableBind() => clearField(4);

  @$pb.TagNumber(5)
  $core.List<$core.int> get bindIp => $_getN(3);
  @$pb.TagNumber(5)
  set bindIp($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(5)
  $core.bool hasBindIp() => $_has(3);
  @$pb.TagNumber(5)
  void clearBindIp() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get bindPort => $_getIZ(4);
  @$pb.TagNumber(6)
  set bindPort($core.int v) { $_setUnsignedInt32(4, v); }
  @$pb.TagNumber(6)
  $core.bool hasBindPort() => $_has(4);
  @$pb.TagNumber(6)
  void clearBindPort() => clearField(6);

  @$pb.TagNumber(7)
  $core.int get flushIntervalMs => $_getIZ(5);
  @$pb.TagNumber(7)
  set flushIntervalMs($core.int v) { $_setUnsignedInt32(5, v); }
  @$pb.TagNumber(7)
  $core.bool hasFlushIntervalMs() => $_has(5);
  @$pb.TagNumber(7)
  void clearFlushIntervalMs() => clearField(7);

  @$pb.TagNumber(8)
  $core.List<$core.int> get token => $_getN(6);
  @$pb.TagNumber(8)
  set token($core.List<$core.int> v) { $_setBytes(6, v); }
  @$pb.TagNumber(8)
  $core.bool hasToken() => $_has(6);
  @$pb.TagNumber(8)
  void clearToken() => clearField(8);
}

class Ipfs extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Ipfs', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(2, 'fakeApiListenIp', $pb.PbFieldType.OY, protoName: 'fakeApiListenIp')
    ..a<$core.int>(3, 'fakeApiListenPort', $pb.PbFieldType.OU3, protoName: 'fakeApiListenPort')
    ..aOB(4, 'enableGateway', protoName: 'enableGateway')
    ..aOS(5, 'gatewayServerName', protoName: 'gatewayServerName')
    ..aOB(6, 'enableApi', protoName: 'enableApi')
    ..aOS(7, 'apiServerName', protoName: 'apiServerName')
    ..pPS(8, 'profile')
    ..aOB(9, 'autoMigrate', protoName: 'autoMigrate')
    ..aOB(10, 'enableIpnsPubSub', protoName: 'enableIpnsPubSub')
    ..aOB(11, 'enablePubSub', protoName: 'enablePubSub')
    ..aOB(12, 'enableMultiplex', protoName: 'enableMultiplex')
    ..aOS(13, 'token')
    ..hasRequiredFields = false
  ;

  Ipfs._() : super();
  factory Ipfs() => create();
  factory Ipfs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Ipfs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Ipfs clone() => Ipfs()..mergeFromMessage(this);
  Ipfs copyWith(void Function(Ipfs) updates) => super.copyWith((message) => updates(message as Ipfs));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Ipfs create() => Ipfs._();
  Ipfs createEmptyInstance() => create();
  static $pb.PbList<Ipfs> createRepeated() => $pb.PbList<Ipfs>();
  @$core.pragma('dart2js:noInline')
  static Ipfs getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Ipfs>(create);
  static Ipfs _defaultInstance;

  @$pb.TagNumber(2)
  $core.List<$core.int> get fakeApiListenIp => $_getN(0);
  @$pb.TagNumber(2)
  set fakeApiListenIp($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(2)
  $core.bool hasFakeApiListenIp() => $_has(0);
  @$pb.TagNumber(2)
  void clearFakeApiListenIp() => clearField(2);

  @$pb.TagNumber(3)
  $core.int get fakeApiListenPort => $_getIZ(1);
  @$pb.TagNumber(3)
  set fakeApiListenPort($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasFakeApiListenPort() => $_has(1);
  @$pb.TagNumber(3)
  void clearFakeApiListenPort() => clearField(3);

  @$pb.TagNumber(4)
  $core.bool get enableGateway => $_getBF(2);
  @$pb.TagNumber(4)
  set enableGateway($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(4)
  $core.bool hasEnableGateway() => $_has(2);
  @$pb.TagNumber(4)
  void clearEnableGateway() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get gatewayServerName => $_getSZ(3);
  @$pb.TagNumber(5)
  set gatewayServerName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(5)
  $core.bool hasGatewayServerName() => $_has(3);
  @$pb.TagNumber(5)
  void clearGatewayServerName() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get enableApi => $_getBF(4);
  @$pb.TagNumber(6)
  set enableApi($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(6)
  $core.bool hasEnableApi() => $_has(4);
  @$pb.TagNumber(6)
  void clearEnableApi() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get apiServerName => $_getSZ(5);
  @$pb.TagNumber(7)
  set apiServerName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(7)
  $core.bool hasApiServerName() => $_has(5);
  @$pb.TagNumber(7)
  void clearApiServerName() => clearField(7);

  @$pb.TagNumber(8)
  $core.List<$core.String> get profile => $_getList(6);

  @$pb.TagNumber(9)
  $core.bool get autoMigrate => $_getBF(7);
  @$pb.TagNumber(9)
  set autoMigrate($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(9)
  $core.bool hasAutoMigrate() => $_has(7);
  @$pb.TagNumber(9)
  void clearAutoMigrate() => clearField(9);

  @$pb.TagNumber(10)
  $core.bool get enableIpnsPubSub => $_getBF(8);
  @$pb.TagNumber(10)
  set enableIpnsPubSub($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(10)
  $core.bool hasEnableIpnsPubSub() => $_has(8);
  @$pb.TagNumber(10)
  void clearEnableIpnsPubSub() => clearField(10);

  @$pb.TagNumber(11)
  $core.bool get enablePubSub => $_getBF(9);
  @$pb.TagNumber(11)
  set enablePubSub($core.bool v) { $_setBool(9, v); }
  @$pb.TagNumber(11)
  $core.bool hasEnablePubSub() => $_has(9);
  @$pb.TagNumber(11)
  void clearEnablePubSub() => clearField(11);

  @$pb.TagNumber(12)
  $core.bool get enableMultiplex => $_getBF(10);
  @$pb.TagNumber(12)
  set enableMultiplex($core.bool v) { $_setBool(10, v); }
  @$pb.TagNumber(12)
  $core.bool hasEnableMultiplex() => $_has(10);
  @$pb.TagNumber(12)
  void clearEnableMultiplex() => clearField(12);

  @$pb.TagNumber(13)
  $core.String get token => $_getSZ(11);
  @$pb.TagNumber(13)
  set token($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(13)
  $core.bool hasToken() => $_has(11);
  @$pb.TagNumber(13)
  void clearToken() => clearField(13);
}

class Log_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Log.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(4, 'target')
    ..hasRequiredFields = false
  ;

  Log_View._() : super();
  factory Log_View() => create();
  factory Log_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Log_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Log_View clone() => Log_View()..mergeFromMessage(this);
  Log_View copyWith(void Function(Log_View) updates) => super.copyWith((message) => updates(message as Log_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Log_View create() => Log_View._();
  Log_View createEmptyInstance() => create();
  static $pb.PbList<Log_View> createRepeated() => $pb.PbList<Log_View>();
  @$core.pragma('dart2js:noInline')
  static Log_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Log_View>(create);
  static Log_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(4)
  $core.String get target => $_getSZ(1);
  @$pb.TagNumber(4)
  set target($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(4)
  $core.bool hasTarget() => $_has(1);
  @$pb.TagNumber(4)
  void clearTarget() => clearField(4);
}

class Log extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Log', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'dev')
    ..aOS(3, 'level')
    ..aOS(4, 'target')
    ..hasRequiredFields = false
  ;

  Log._() : super();
  factory Log() => create();
  factory Log.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Log.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Log clone() => Log()..mergeFromMessage(this);
  Log copyWith(void Function(Log) updates) => super.copyWith((message) => updates(message as Log));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Log create() => Log._();
  Log createEmptyInstance() => create();
  static $pb.PbList<Log> createRepeated() => $pb.PbList<Log>();
  @$core.pragma('dart2js:noInline')
  static Log getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Log>(create);
  static Log _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get dev => $_getBF(1);
  @$pb.TagNumber(2)
  set dev($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDev() => $_has(1);
  @$pb.TagNumber(2)
  void clearDev() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get level => $_getSZ(2);
  @$pb.TagNumber(3)
  set level($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasLevel() => $_has(2);
  @$pb.TagNumber(3)
  void clearLevel() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get target => $_getSZ(3);
  @$pb.TagNumber(4)
  set target($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasTarget() => $_has(3);
  @$pb.TagNumber(4)
  void clearTarget() => clearField(4);
}

class IpfsServer_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServer.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(3, 'name')
    ..hasRequiredFields = false
  ;

  IpfsServer_View._() : super();
  factory IpfsServer_View() => create();
  factory IpfsServer_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServer_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServer_View clone() => IpfsServer_View()..mergeFromMessage(this);
  IpfsServer_View copyWith(void Function(IpfsServer_View) updates) => super.copyWith((message) => updates(message as IpfsServer_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServer_View create() => IpfsServer_View._();
  IpfsServer_View createEmptyInstance() => create();
  static $pb.PbList<IpfsServer_View> createRepeated() => $pb.PbList<IpfsServer_View>();
  @$core.pragma('dart2js:noInline')
  static IpfsServer_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServer_View>(create);
  static IpfsServer_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
}

class IpfsServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'disabled')
    ..aOM<$4.StringValue>(3, 'name', subBuilder: $4.StringValue.create)
    ..aOS(4, 'peer')
    ..aOS(5, 'token')
    ..hasRequiredFields = false
  ;

  IpfsServer._() : super();
  factory IpfsServer() => create();
  factory IpfsServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IpfsServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IpfsServer clone() => IpfsServer()..mergeFromMessage(this);
  IpfsServer copyWith(void Function(IpfsServer) updates) => super.copyWith((message) => updates(message as IpfsServer));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IpfsServer create() => IpfsServer._();
  IpfsServer createEmptyInstance() => create();
  static $pb.PbList<IpfsServer> createRepeated() => $pb.PbList<IpfsServer>();
  @$core.pragma('dart2js:noInline')
  static IpfsServer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IpfsServer>(create);
  static IpfsServer _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get disabled => $_getBF(1);
  @$pb.TagNumber(2)
  set disabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisabled() => clearField(2);

  @$pb.TagNumber(3)
  $4.StringValue get name => $_getN(2);
  @$pb.TagNumber(3)
  set name($4.StringValue v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
  @$pb.TagNumber(3)
  $4.StringValue ensureName() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.String get peer => $_getSZ(3);
  @$pb.TagNumber(4)
  set peer($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPeer() => $_has(3);
  @$pb.TagNumber(4)
  void clearPeer() => clearField(4);

  @$pb.TagNumber(5)
  $core.String get token => $_getSZ(4);
  @$pb.TagNumber(5)
  set token($core.String v) { $_setString(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasToken() => $_has(4);
  @$pb.TagNumber(5)
  void clearToken() => clearField(5);
}

class FileServer_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServer.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(3, 'name')
    ..hasRequiredFields = false
  ;

  FileServer_View._() : super();
  factory FileServer_View() => create();
  factory FileServer_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServer_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServer_View clone() => FileServer_View()..mergeFromMessage(this);
  FileServer_View copyWith(void Function(FileServer_View) updates) => super.copyWith((message) => updates(message as FileServer_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServer_View create() => FileServer_View._();
  FileServer_View createEmptyInstance() => create();
  static $pb.PbList<FileServer_View> createRepeated() => $pb.PbList<FileServer_View>();
  @$core.pragma('dart2js:noInline')
  static FileServer_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServer_View>(create);
  static FileServer_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
}

class FileServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'disabled')
    ..aOM<$4.StringValue>(3, 'name', subBuilder: $4.StringValue.create)
    ..aOS(4, 'zip')
    ..m<$core.String, $core.String>(5, 'redirect', entryClassName: 'FileServer.RedirectEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('hybrid'))
    ..aOB(6, 'dev')
    ..hasRequiredFields = false
  ;

  FileServer._() : super();
  factory FileServer() => create();
  factory FileServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory FileServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  FileServer clone() => FileServer()..mergeFromMessage(this);
  FileServer copyWith(void Function(FileServer) updates) => super.copyWith((message) => updates(message as FileServer));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static FileServer create() => FileServer._();
  FileServer createEmptyInstance() => create();
  static $pb.PbList<FileServer> createRepeated() => $pb.PbList<FileServer>();
  @$core.pragma('dart2js:noInline')
  static FileServer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<FileServer>(create);
  static FileServer _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get disabled => $_getBF(1);
  @$pb.TagNumber(2)
  set disabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisabled() => clearField(2);

  @$pb.TagNumber(3)
  $4.StringValue get name => $_getN(2);
  @$pb.TagNumber(3)
  set name($4.StringValue v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
  @$pb.TagNumber(3)
  $4.StringValue ensureName() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.String get zip => $_getSZ(3);
  @$pb.TagNumber(4)
  set zip($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasZip() => $_has(3);
  @$pb.TagNumber(4)
  void clearZip() => clearField(4);

  @$pb.TagNumber(5)
  $core.Map<$core.String, $core.String> get redirect => $_getMap(4);

  @$pb.TagNumber(6)
  $core.bool get dev => $_getBF(5);
  @$pb.TagNumber(6)
  set dev($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasDev() => $_has(5);
  @$pb.TagNumber(6)
  void clearDev() => clearField(6);
}

class HttpProxyServer_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServer.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(3, 'name')
    ..hasRequiredFields = false
  ;

  HttpProxyServer_View._() : super();
  factory HttpProxyServer_View() => create();
  factory HttpProxyServer_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServer_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServer_View clone() => HttpProxyServer_View()..mergeFromMessage(this);
  HttpProxyServer_View copyWith(void Function(HttpProxyServer_View) updates) => super.copyWith((message) => updates(message as HttpProxyServer_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServer_View create() => HttpProxyServer_View._();
  HttpProxyServer_View createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServer_View> createRepeated() => $pb.PbList<HttpProxyServer_View>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServer_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServer_View>(create);
  static HttpProxyServer_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
}

class HttpProxyServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'disabled')
    ..aOM<$4.StringValue>(3, 'name', subBuilder: $4.StringValue.create)
    ..aOS(4, 'host')
    ..a<$core.int>(5, 'port', $pb.PbFieldType.OU3)
    ..aOB(6, 'keepAlive', protoName: 'keepAlive')
    ..hasRequiredFields = false
  ;

  HttpProxyServer._() : super();
  factory HttpProxyServer() => create();
  factory HttpProxyServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory HttpProxyServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  HttpProxyServer clone() => HttpProxyServer()..mergeFromMessage(this);
  HttpProxyServer copyWith(void Function(HttpProxyServer) updates) => super.copyWith((message) => updates(message as HttpProxyServer));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static HttpProxyServer create() => HttpProxyServer._();
  HttpProxyServer createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServer> createRepeated() => $pb.PbList<HttpProxyServer>();
  @$core.pragma('dart2js:noInline')
  static HttpProxyServer getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<HttpProxyServer>(create);
  static HttpProxyServer _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get disabled => $_getBF(1);
  @$pb.TagNumber(2)
  set disabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisabled() => clearField(2);

  @$pb.TagNumber(3)
  $4.StringValue get name => $_getN(2);
  @$pb.TagNumber(3)
  set name($4.StringValue v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
  @$pb.TagNumber(3)
  $4.StringValue ensureName() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.String get host => $_getSZ(3);
  @$pb.TagNumber(4)
  set host($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasHost() => $_has(3);
  @$pb.TagNumber(4)
  void clearHost() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get port => $_getIZ(4);
  @$pb.TagNumber(5)
  set port($core.int v) { $_setUnsignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasPort() => $_has(4);
  @$pb.TagNumber(5)
  void clearPort() => clearField(5);

  @$pb.TagNumber(6)
  $core.bool get keepAlive => $_getBF(5);
  @$pb.TagNumber(6)
  set keepAlive($core.bool v) { $_setBool(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasKeepAlive() => $_has(5);
  @$pb.TagNumber(6)
  void clearKeepAlive() => clearField(6);
}

enum ServerView_Group {
  ipfs, 
  file, 
  http, 
  notSet
}

class ServerView extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, ServerView_Group> _ServerView_GroupByTag = {
    1 : ServerView_Group.ipfs,
    2 : ServerView_Group.file,
    3 : ServerView_Group.http,
    0 : ServerView_Group.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ServerView', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2, 3])
    ..aOM<IpfsServer_View>(1, 'ipfs', subBuilder: IpfsServer_View.create)
    ..aOM<FileServer_View>(2, 'file', subBuilder: FileServer_View.create)
    ..aOM<HttpProxyServer_View>(3, 'http', subBuilder: HttpProxyServer_View.create)
    ..hasRequiredFields = false
  ;

  ServerView._() : super();
  factory ServerView() => create();
  factory ServerView.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory ServerView.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  ServerView clone() => ServerView()..mergeFromMessage(this);
  ServerView copyWith(void Function(ServerView) updates) => super.copyWith((message) => updates(message as ServerView));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static ServerView create() => ServerView._();
  ServerView createEmptyInstance() => create();
  static $pb.PbList<ServerView> createRepeated() => $pb.PbList<ServerView>();
  @$core.pragma('dart2js:noInline')
  static ServerView getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<ServerView>(create);
  static ServerView _defaultInstance;

  ServerView_Group whichGroup() => _ServerView_GroupByTag[$_whichOneof(0)];
  void clearGroup() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  IpfsServer_View get ipfs => $_getN(0);
  @$pb.TagNumber(1)
  set ipfs(IpfsServer_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasIpfs() => $_has(0);
  @$pb.TagNumber(1)
  void clearIpfs() => clearField(1);
  @$pb.TagNumber(1)
  IpfsServer_View ensureIpfs() => $_ensure(0);

  @$pb.TagNumber(2)
  FileServer_View get file => $_getN(1);
  @$pb.TagNumber(2)
  set file(FileServer_View v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasFile() => $_has(1);
  @$pb.TagNumber(2)
  void clearFile() => clearField(2);
  @$pb.TagNumber(2)
  FileServer_View ensureFile() => $_ensure(1);

  @$pb.TagNumber(3)
  HttpProxyServer_View get http => $_getN(2);
  @$pb.TagNumber(3)
  set http(HttpProxyServer_View v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasHttp() => $_has(2);
  @$pb.TagNumber(3)
  void clearHttp() => clearField(3);
  @$pb.TagNumber(3)
  HttpProxyServer_View ensureHttp() => $_ensure(2);
}

class AdpRouter_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouter.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(3, 'name')
    ..hasRequiredFields = false
  ;

  AdpRouter_View._() : super();
  factory AdpRouter_View() => create();
  factory AdpRouter_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouter_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouter_View clone() => AdpRouter_View()..mergeFromMessage(this);
  AdpRouter_View copyWith(void Function(AdpRouter_View) updates) => super.copyWith((message) => updates(message as AdpRouter_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouter_View create() => AdpRouter_View._();
  AdpRouter_View createEmptyInstance() => create();
  static $pb.PbList<AdpRouter_View> createRepeated() => $pb.PbList<AdpRouter_View>();
  @$core.pragma('dart2js:noInline')
  static AdpRouter_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouter_View>(create);
  static AdpRouter_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
}

class AdpRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouter', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'disabled')
    ..aOS(3, 'name')
    ..aOS(4, 'rulesDirName', protoName: 'rulesDirName')
    ..aOM<ServerView>(5, 'blocked', subBuilder: ServerView.create)
    ..aOM<ServerView>(6, 'unblocked', subBuilder: ServerView.create)
    ..aOB(7, 'etcHostsIpAsBlocked', protoName: 'etcHostsIpAsBlocked')
    ..aOB(8, 'dev')
    ..hasRequiredFields = false
  ;

  AdpRouter._() : super();
  factory AdpRouter() => create();
  factory AdpRouter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory AdpRouter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  AdpRouter clone() => AdpRouter()..mergeFromMessage(this);
  AdpRouter copyWith(void Function(AdpRouter) updates) => super.copyWith((message) => updates(message as AdpRouter));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static AdpRouter create() => AdpRouter._();
  AdpRouter createEmptyInstance() => create();
  static $pb.PbList<AdpRouter> createRepeated() => $pb.PbList<AdpRouter>();
  @$core.pragma('dart2js:noInline')
  static AdpRouter getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<AdpRouter>(create);
  static AdpRouter _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get disabled => $_getBF(1);
  @$pb.TagNumber(2)
  set disabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisabled() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get rulesDirName => $_getSZ(3);
  @$pb.TagNumber(4)
  set rulesDirName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasRulesDirName() => $_has(3);
  @$pb.TagNumber(4)
  void clearRulesDirName() => clearField(4);

  @$pb.TagNumber(5)
  ServerView get blocked => $_getN(4);
  @$pb.TagNumber(5)
  set blocked(ServerView v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasBlocked() => $_has(4);
  @$pb.TagNumber(5)
  void clearBlocked() => clearField(5);
  @$pb.TagNumber(5)
  ServerView ensureBlocked() => $_ensure(4);

  @$pb.TagNumber(6)
  ServerView get unblocked => $_getN(5);
  @$pb.TagNumber(6)
  set unblocked(ServerView v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasUnblocked() => $_has(5);
  @$pb.TagNumber(6)
  void clearUnblocked() => clearField(6);
  @$pb.TagNumber(6)
  ServerView ensureUnblocked() => $_ensure(5);

  @$pb.TagNumber(7)
  $core.bool get etcHostsIpAsBlocked => $_getBF(6);
  @$pb.TagNumber(7)
  set etcHostsIpAsBlocked($core.bool v) { $_setBool(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasEtcHostsIpAsBlocked() => $_has(6);
  @$pb.TagNumber(7)
  void clearEtcHostsIpAsBlocked() => clearField(7);

  @$pb.TagNumber(8)
  $core.bool get dev => $_getBF(7);
  @$pb.TagNumber(8)
  set dev($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasDev() => $_has(7);
  @$pb.TagNumber(8)
  void clearDev() => clearField(8);
}

class IPNetRouter_View extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouter.View', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOS(3, 'name')
    ..hasRequiredFields = false
  ;

  IPNetRouter_View._() : super();
  factory IPNetRouter_View() => create();
  factory IPNetRouter_View.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouter_View.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouter_View clone() => IPNetRouter_View()..mergeFromMessage(this);
  IPNetRouter_View copyWith(void Function(IPNetRouter_View) updates) => super.copyWith((message) => updates(message as IPNetRouter_View));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouter_View create() => IPNetRouter_View._();
  IPNetRouter_View createEmptyInstance() => create();
  static $pb.PbList<IPNetRouter_View> createRepeated() => $pb.PbList<IPNetRouter_View>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouter_View getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouter_View>(create);
  static IPNetRouter_View _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);
}

class IPNetRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouter', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.int>(1, 'id', $pb.PbFieldType.OU3)
    ..aOB(2, 'disabled')
    ..aOS(3, 'name')
    ..pc<$4.BytesValue>(4, 'ip', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..pc<$4.BytesValue>(5, 'netv4', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..pc<$4.BytesValue>(6, 'netv6', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..aOM<ServerView>(7, 'matched', subBuilder: ServerView.create)
    ..aOM<ServerView>(8, 'unmatched', subBuilder: ServerView.create)
    ..aOM<FileServer_View>(9, 'fileTest', protoName: 'fileTest', subBuilder: FileServer_View.create)
    ..hasRequiredFields = false
  ;

  IPNetRouter._() : super();
  factory IPNetRouter() => create();
  factory IPNetRouter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory IPNetRouter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  IPNetRouter clone() => IPNetRouter()..mergeFromMessage(this);
  IPNetRouter copyWith(void Function(IPNetRouter) updates) => super.copyWith((message) => updates(message as IPNetRouter));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static IPNetRouter create() => IPNetRouter._();
  IPNetRouter createEmptyInstance() => create();
  static $pb.PbList<IPNetRouter> createRepeated() => $pb.PbList<IPNetRouter>();
  @$core.pragma('dart2js:noInline')
  static IPNetRouter getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<IPNetRouter>(create);
  static IPNetRouter _defaultInstance;

  @$pb.TagNumber(1)
  $core.int get id => $_getIZ(0);
  @$pb.TagNumber(1)
  set id($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get disabled => $_getBF(1);
  @$pb.TagNumber(2)
  set disabled($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDisabled() => $_has(1);
  @$pb.TagNumber(2)
  void clearDisabled() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(2);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(2);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$4.BytesValue> get ip => $_getList(3);

  @$pb.TagNumber(5)
  $core.List<$4.BytesValue> get netv4 => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<$4.BytesValue> get netv6 => $_getList(5);

  @$pb.TagNumber(7)
  ServerView get matched => $_getN(6);
  @$pb.TagNumber(7)
  set matched(ServerView v) { setField(7, v); }
  @$pb.TagNumber(7)
  $core.bool hasMatched() => $_has(6);
  @$pb.TagNumber(7)
  void clearMatched() => clearField(7);
  @$pb.TagNumber(7)
  ServerView ensureMatched() => $_ensure(6);

  @$pb.TagNumber(8)
  ServerView get unmatched => $_getN(7);
  @$pb.TagNumber(8)
  set unmatched(ServerView v) { setField(8, v); }
  @$pb.TagNumber(8)
  $core.bool hasUnmatched() => $_has(7);
  @$pb.TagNumber(8)
  void clearUnmatched() => clearField(8);
  @$pb.TagNumber(8)
  ServerView ensureUnmatched() => $_ensure(7);

  @$pb.TagNumber(9)
  FileServer_View get fileTest => $_getN(8);
  @$pb.TagNumber(9)
  set fileTest(FileServer_View v) { setField(9, v); }
  @$pb.TagNumber(9)
  $core.bool hasFileTest() => $_has(8);
  @$pb.TagNumber(9)
  void clearFileTest() => clearField(9);
  @$pb.TagNumber(9)
  FileServer_View ensureFileTest() => $_ensure(8);
}

enum RouterUseView_Id_Is {
  adp, 
  ipnet, 
  notSet
}

class RouterUseView_Id extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RouterUseView_Id_Is> _RouterUseView_Id_IsByTag = {
    1 : RouterUseView_Id_Is.adp,
    2 : RouterUseView_Id_Is.ipnet,
    0 : RouterUseView_Id_Is.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseView.Id', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..a<$core.int>(1, 'adp', $pb.PbFieldType.OU3)
    ..a<$core.int>(2, 'ipnet', $pb.PbFieldType.OU3)
    ..hasRequiredFields = false
  ;

  RouterUseView_Id._() : super();
  factory RouterUseView_Id() => create();
  factory RouterUseView_Id.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseView_Id.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseView_Id clone() => RouterUseView_Id()..mergeFromMessage(this);
  RouterUseView_Id copyWith(void Function(RouterUseView_Id) updates) => super.copyWith((message) => updates(message as RouterUseView_Id));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseView_Id create() => RouterUseView_Id._();
  RouterUseView_Id createEmptyInstance() => create();
  static $pb.PbList<RouterUseView_Id> createRepeated() => $pb.PbList<RouterUseView_Id>();
  @$core.pragma('dart2js:noInline')
  static RouterUseView_Id getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseView_Id>(create);
  static RouterUseView_Id _defaultInstance;

  RouterUseView_Id_Is whichIs() => _RouterUseView_Id_IsByTag[$_whichOneof(0)];
  void clearIs() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.int get adp => $_getIZ(0);
  @$pb.TagNumber(1)
  set adp($core.int v) { $_setUnsignedInt32(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasAdp() => $_has(0);
  @$pb.TagNumber(1)
  void clearAdp() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get ipnet => $_getIZ(1);
  @$pb.TagNumber(2)
  set ipnet($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasIpnet() => $_has(1);
  @$pb.TagNumber(2)
  void clearIpnet() => clearField(2);
}

class RouterUseView_Element extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseView.Element', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<RouterUseView_Id>(1, 'id', subBuilder: RouterUseView_Id.create)
    ..aOS(3, 'name')
    ..aInt64(4, 'rank')
    ..hasRequiredFields = false
  ;

  RouterUseView_Element._() : super();
  factory RouterUseView_Element() => create();
  factory RouterUseView_Element.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseView_Element.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseView_Element clone() => RouterUseView_Element()..mergeFromMessage(this);
  RouterUseView_Element copyWith(void Function(RouterUseView_Element) updates) => super.copyWith((message) => updates(message as RouterUseView_Element));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseView_Element create() => RouterUseView_Element._();
  RouterUseView_Element createEmptyInstance() => create();
  static $pb.PbList<RouterUseView_Element> createRepeated() => $pb.PbList<RouterUseView_Element>();
  @$core.pragma('dart2js:noInline')
  static RouterUseView_Element getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseView_Element>(create);
  static RouterUseView_Element _defaultInstance;

  @$pb.TagNumber(1)
  RouterUseView_Id get id => $_getN(0);
  @$pb.TagNumber(1)
  set id(RouterUseView_Id v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasId() => $_has(0);
  @$pb.TagNumber(1)
  void clearId() => clearField(1);
  @$pb.TagNumber(1)
  RouterUseView_Id ensureId() => $_ensure(0);

  @$pb.TagNumber(3)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(3)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(3)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(3)
  void clearName() => clearField(3);

  @$pb.TagNumber(4)
  $fixnum.Int64 get rank => $_getI64(2);
  @$pb.TagNumber(4)
  set rank($fixnum.Int64 v) { $_setInt64(2, v); }
  @$pb.TagNumber(4)
  $core.bool hasRank() => $_has(2);
  @$pb.TagNumber(4)
  void clearRank() => clearField(4);
}

enum RouterUseView_Group {
  adp, 
  ipnet, 
  notSet
}

class RouterUseView extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RouterUseView_Group> _RouterUseView_GroupByTag = {
    1 : RouterUseView_Group.adp,
    2 : RouterUseView_Group.ipnet,
    0 : RouterUseView_Group.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterUseView', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [1, 2])
    ..aOM<AdpRouter_View>(1, 'adp', subBuilder: AdpRouter_View.create)
    ..aOM<IPNetRouter_View>(2, 'ipnet', subBuilder: IPNetRouter_View.create)
    ..hasRequiredFields = false
  ;

  RouterUseView._() : super();
  factory RouterUseView() => create();
  factory RouterUseView.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterUseView.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterUseView clone() => RouterUseView()..mergeFromMessage(this);
  RouterUseView copyWith(void Function(RouterUseView) updates) => super.copyWith((message) => updates(message as RouterUseView));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterUseView create() => RouterUseView._();
  RouterUseView createEmptyInstance() => create();
  static $pb.PbList<RouterUseView> createRepeated() => $pb.PbList<RouterUseView>();
  @$core.pragma('dart2js:noInline')
  static RouterUseView getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterUseView>(create);
  static RouterUseView _defaultInstance;

  RouterUseView_Group whichGroup() => _RouterUseView_GroupByTag[$_whichOneof(0)];
  void clearGroup() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  AdpRouter_View get adp => $_getN(0);
  @$pb.TagNumber(1)
  set adp(AdpRouter_View v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasAdp() => $_has(0);
  @$pb.TagNumber(1)
  void clearAdp() => clearField(1);
  @$pb.TagNumber(1)
  AdpRouter_View ensureAdp() => $_ensure(0);

  @$pb.TagNumber(2)
  IPNetRouter_View get ipnet => $_getN(1);
  @$pb.TagNumber(2)
  set ipnet(IPNetRouter_View v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasIpnet() => $_has(1);
  @$pb.TagNumber(2)
  void clearIpnet() => clearField(2);
  @$pb.TagNumber(2)
  IPNetRouter_View ensureIpnet() => $_ensure(1);
}

class Config extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Config', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<Basic>(1, 'basic', subBuilder: Basic.create)
    ..aOM<Ipfs>(2, 'ipfs', subBuilder: Ipfs.create)
    ..aOM<Log_View>(3, 'log', subBuilder: Log_View.create)
    ..pc<Log_View>(4, 'logs', $pb.PbFieldType.PM, subBuilder: Log_View.create)
    ..pc<ServerView>(5, 'servers', $pb.PbFieldType.PM, subBuilder: ServerView.create)
    ..pc<AdpRouter_View>(6, 'adps', $pb.PbFieldType.PM, subBuilder: AdpRouter_View.create)
    ..pc<IPNetRouter_View>(7, 'ipnets', $pb.PbFieldType.PM, subBuilder: IPNetRouter_View.create)
    ..pc<RouterUseView_Element>(8, 'routers', $pb.PbFieldType.PM, subBuilder: RouterUseView_Element.create)
    ..hasRequiredFields = false
  ;

  Config._() : super();
  factory Config() => create();
  factory Config.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory Config.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  Config clone() => Config()..mergeFromMessage(this);
  Config copyWith(void Function(Config) updates) => super.copyWith((message) => updates(message as Config));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static Config create() => Config._();
  Config createEmptyInstance() => create();
  static $pb.PbList<Config> createRepeated() => $pb.PbList<Config>();
  @$core.pragma('dart2js:noInline')
  static Config getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<Config>(create);
  static Config _defaultInstance;

  @$pb.TagNumber(1)
  Basic get basic => $_getN(0);
  @$pb.TagNumber(1)
  set basic(Basic v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasBasic() => $_has(0);
  @$pb.TagNumber(1)
  void clearBasic() => clearField(1);
  @$pb.TagNumber(1)
  Basic ensureBasic() => $_ensure(0);

  @$pb.TagNumber(2)
  Ipfs get ipfs => $_getN(1);
  @$pb.TagNumber(2)
  set ipfs(Ipfs v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasIpfs() => $_has(1);
  @$pb.TagNumber(2)
  void clearIpfs() => clearField(2);
  @$pb.TagNumber(2)
  Ipfs ensureIpfs() => $_ensure(1);

  @$pb.TagNumber(3)
  Log_View get log => $_getN(2);
  @$pb.TagNumber(3)
  set log(Log_View v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasLog() => $_has(2);
  @$pb.TagNumber(3)
  void clearLog() => clearField(3);
  @$pb.TagNumber(3)
  Log_View ensureLog() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.List<Log_View> get logs => $_getList(3);

  @$pb.TagNumber(5)
  $core.List<ServerView> get servers => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<AdpRouter_View> get adps => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<IPNetRouter_View> get ipnets => $_getList(6);

  @$pb.TagNumber(8)
  $core.List<RouterUseView_Element> get routers => $_getList(7);
}

