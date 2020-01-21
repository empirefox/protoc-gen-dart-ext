///
//  Generated code. Do not modify.
//  source: hybrid/config.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

import 'dart:core' as $core;

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
    ..aOS(1, 'version')
    ..aOB(2, 'dev')
    ..aOB(3, 'enableBind', protoName: 'enableBind')
    ..a<$core.List<$core.int>>(4, 'bindIp', $pb.PbFieldType.OY, protoName: 'bindIp')
    ..a<$core.int>(5, 'bindPort', $pb.PbFieldType.OU3, protoName: 'bindPort')
    ..a<$core.int>(6, 'flushIntervalMs', $pb.PbFieldType.OU3, protoName: 'flushIntervalMs')
    ..aOS(7, 'token')
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

  @$pb.TagNumber(1)
  $core.String get version => $_getSZ(0);
  @$pb.TagNumber(1)
  set version($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasVersion() => $_has(0);
  @$pb.TagNumber(1)
  void clearVersion() => clearField(1);

  @$pb.TagNumber(2)
  $core.bool get dev => $_getBF(1);
  @$pb.TagNumber(2)
  set dev($core.bool v) { $_setBool(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDev() => $_has(1);
  @$pb.TagNumber(2)
  void clearDev() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get enableBind => $_getBF(2);
  @$pb.TagNumber(3)
  set enableBind($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEnableBind() => $_has(2);
  @$pb.TagNumber(3)
  void clearEnableBind() => clearField(3);

  @$pb.TagNumber(4)
  $core.List<$core.int> get bindIp => $_getN(3);
  @$pb.TagNumber(4)
  set bindIp($core.List<$core.int> v) { $_setBytes(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasBindIp() => $_has(3);
  @$pb.TagNumber(4)
  void clearBindIp() => clearField(4);

  @$pb.TagNumber(5)
  $core.int get bindPort => $_getIZ(4);
  @$pb.TagNumber(5)
  set bindPort($core.int v) { $_setUnsignedInt32(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasBindPort() => $_has(4);
  @$pb.TagNumber(5)
  void clearBindPort() => clearField(5);

  @$pb.TagNumber(6)
  $core.int get flushIntervalMs => $_getIZ(5);
  @$pb.TagNumber(6)
  set flushIntervalMs($core.int v) { $_setUnsignedInt32(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasFlushIntervalMs() => $_has(5);
  @$pb.TagNumber(6)
  void clearFlushIntervalMs() => clearField(6);

  @$pb.TagNumber(7)
  $core.String get token => $_getSZ(6);
  @$pb.TagNumber(7)
  set token($core.String v) { $_setString(6, v); }
  @$pb.TagNumber(7)
  $core.bool hasToken() => $_has(6);
  @$pb.TagNumber(7)
  void clearToken() => clearField(7);
}

class Log extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Log', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOB(1, 'dev')
    ..aOS(2, 'level')
    ..aOS(3, 'target')
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
  $core.bool get dev => $_getBF(0);
  @$pb.TagNumber(1)
  set dev($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDev() => $_has(0);
  @$pb.TagNumber(1)
  void clearDev() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get level => $_getSZ(1);
  @$pb.TagNumber(2)
  set level($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasLevel() => $_has(1);
  @$pb.TagNumber(2)
  void clearLevel() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get target => $_getSZ(2);
  @$pb.TagNumber(3)
  set target($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasTarget() => $_has(2);
  @$pb.TagNumber(3)
  void clearTarget() => clearField(3);
}

class Ipfs extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Ipfs', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..a<$core.List<$core.int>>(1, 'fakeApiListenIp', $pb.PbFieldType.OY, protoName: 'fakeApiListenIp')
    ..a<$core.int>(2, 'fakeApiListenPort', $pb.PbFieldType.OU3, protoName: 'fakeApiListenPort')
    ..aOB(3, 'enableGateway', protoName: 'enableGateway')
    ..aOS(4, 'gatewayServerName', protoName: 'gatewayServerName')
    ..aOB(5, 'enableApi', protoName: 'enableApi')
    ..aOS(6, 'apiServerName', protoName: 'apiServerName')
    ..pPS(7, 'profile')
    ..aOB(8, 'autoMigrate', protoName: 'autoMigrate')
    ..aOB(9, 'enableIpnsPubSub', protoName: 'enableIpnsPubSub')
    ..aOB(10, 'enablePubSub', protoName: 'enablePubSub')
    ..aOB(11, 'enableMultiplex', protoName: 'enableMultiplex')
    ..aOS(12, 'token')
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

  @$pb.TagNumber(1)
  $core.List<$core.int> get fakeApiListenIp => $_getN(0);
  @$pb.TagNumber(1)
  set fakeApiListenIp($core.List<$core.int> v) { $_setBytes(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasFakeApiListenIp() => $_has(0);
  @$pb.TagNumber(1)
  void clearFakeApiListenIp() => clearField(1);

  @$pb.TagNumber(2)
  $core.int get fakeApiListenPort => $_getIZ(1);
  @$pb.TagNumber(2)
  set fakeApiListenPort($core.int v) { $_setUnsignedInt32(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasFakeApiListenPort() => $_has(1);
  @$pb.TagNumber(2)
  void clearFakeApiListenPort() => clearField(2);

  @$pb.TagNumber(3)
  $core.bool get enableGateway => $_getBF(2);
  @$pb.TagNumber(3)
  set enableGateway($core.bool v) { $_setBool(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasEnableGateway() => $_has(2);
  @$pb.TagNumber(3)
  void clearEnableGateway() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get gatewayServerName => $_getSZ(3);
  @$pb.TagNumber(4)
  set gatewayServerName($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasGatewayServerName() => $_has(3);
  @$pb.TagNumber(4)
  void clearGatewayServerName() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get enableApi => $_getBF(4);
  @$pb.TagNumber(5)
  set enableApi($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasEnableApi() => $_has(4);
  @$pb.TagNumber(5)
  void clearEnableApi() => clearField(5);

  @$pb.TagNumber(6)
  $core.String get apiServerName => $_getSZ(5);
  @$pb.TagNumber(6)
  set apiServerName($core.String v) { $_setString(5, v); }
  @$pb.TagNumber(6)
  $core.bool hasApiServerName() => $_has(5);
  @$pb.TagNumber(6)
  void clearApiServerName() => clearField(6);

  @$pb.TagNumber(7)
  $core.List<$core.String> get profile => $_getList(6);

  @$pb.TagNumber(8)
  $core.bool get autoMigrate => $_getBF(7);
  @$pb.TagNumber(8)
  set autoMigrate($core.bool v) { $_setBool(7, v); }
  @$pb.TagNumber(8)
  $core.bool hasAutoMigrate() => $_has(7);
  @$pb.TagNumber(8)
  void clearAutoMigrate() => clearField(8);

  @$pb.TagNumber(9)
  $core.bool get enableIpnsPubSub => $_getBF(8);
  @$pb.TagNumber(9)
  set enableIpnsPubSub($core.bool v) { $_setBool(8, v); }
  @$pb.TagNumber(9)
  $core.bool hasEnableIpnsPubSub() => $_has(8);
  @$pb.TagNumber(9)
  void clearEnableIpnsPubSub() => clearField(9);

  @$pb.TagNumber(10)
  $core.bool get enablePubSub => $_getBF(9);
  @$pb.TagNumber(10)
  set enablePubSub($core.bool v) { $_setBool(9, v); }
  @$pb.TagNumber(10)
  $core.bool hasEnablePubSub() => $_has(9);
  @$pb.TagNumber(10)
  void clearEnablePubSub() => clearField(10);

  @$pb.TagNumber(11)
  $core.bool get enableMultiplex => $_getBF(10);
  @$pb.TagNumber(11)
  set enableMultiplex($core.bool v) { $_setBool(10, v); }
  @$pb.TagNumber(11)
  $core.bool hasEnableMultiplex() => $_has(10);
  @$pb.TagNumber(11)
  void clearEnableMultiplex() => clearField(11);

  @$pb.TagNumber(12)
  $core.String get token => $_getSZ(11);
  @$pb.TagNumber(12)
  set token($core.String v) { $_setString(11, v); }
  @$pb.TagNumber(12)
  $core.bool hasToken() => $_has(11);
  @$pb.TagNumber(12)
  void clearToken() => clearField(12);
}

class IpfsServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOB(1, 'disabled')
    ..aOM<$4.StringValue>(2, 'name', subBuilder: $4.StringValue.create)
    ..aOS(3, 'peer')
    ..aOS(4, 'token')
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
  $core.bool get disabled => $_getBF(0);
  @$pb.TagNumber(1)
  set disabled($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDisabled() => $_has(0);
  @$pb.TagNumber(1)
  void clearDisabled() => clearField(1);

  @$pb.TagNumber(2)
  $4.StringValue get name => $_getN(1);
  @$pb.TagNumber(2)
  set name($4.StringValue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);
  @$pb.TagNumber(2)
  $4.StringValue ensureName() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get peer => $_getSZ(2);
  @$pb.TagNumber(3)
  set peer($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasPeer() => $_has(2);
  @$pb.TagNumber(3)
  void clearPeer() => clearField(3);

  @$pb.TagNumber(4)
  $core.String get token => $_getSZ(3);
  @$pb.TagNumber(4)
  set token($core.String v) { $_setString(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasToken() => $_has(3);
  @$pb.TagNumber(4)
  void clearToken() => clearField(4);
}

class FileServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOB(1, 'disabled')
    ..aOM<$4.StringValue>(2, 'name', subBuilder: $4.StringValue.create)
    ..aOS(3, 'zip')
    ..m<$core.String, $core.String>(4, 'redirect', entryClassName: 'FileServer.RedirectEntry', keyFieldType: $pb.PbFieldType.OS, valueFieldType: $pb.PbFieldType.OS, packageName: const $pb.PackageName('hybrid'))
    ..aOB(5, 'dev')
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
  $core.bool get disabled => $_getBF(0);
  @$pb.TagNumber(1)
  set disabled($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDisabled() => $_has(0);
  @$pb.TagNumber(1)
  void clearDisabled() => clearField(1);

  @$pb.TagNumber(2)
  $4.StringValue get name => $_getN(1);
  @$pb.TagNumber(2)
  set name($4.StringValue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);
  @$pb.TagNumber(2)
  $4.StringValue ensureName() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get zip => $_getSZ(2);
  @$pb.TagNumber(3)
  set zip($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasZip() => $_has(2);
  @$pb.TagNumber(3)
  void clearZip() => clearField(3);

  @$pb.TagNumber(4)
  $core.Map<$core.String, $core.String> get redirect => $_getMap(3);

  @$pb.TagNumber(5)
  $core.bool get dev => $_getBF(4);
  @$pb.TagNumber(5)
  set dev($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDev() => $_has(4);
  @$pb.TagNumber(5)
  void clearDev() => clearField(5);
}

class HttpProxyServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServer', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOB(1, 'disabled')
    ..aOM<$4.StringValue>(2, 'name', subBuilder: $4.StringValue.create)
    ..aOS(3, 'host')
    ..a<$core.int>(4, 'port', $pb.PbFieldType.OU3)
    ..aOB(5, 'keepAlive', protoName: 'keepAlive')
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
  $core.bool get disabled => $_getBF(0);
  @$pb.TagNumber(1)
  set disabled($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDisabled() => $_has(0);
  @$pb.TagNumber(1)
  void clearDisabled() => clearField(1);

  @$pb.TagNumber(2)
  $4.StringValue get name => $_getN(1);
  @$pb.TagNumber(2)
  set name($4.StringValue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);
  @$pb.TagNumber(2)
  $4.StringValue ensureName() => $_ensure(1);

  @$pb.TagNumber(3)
  $core.String get host => $_getSZ(2);
  @$pb.TagNumber(3)
  set host($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasHost() => $_has(2);
  @$pb.TagNumber(3)
  void clearHost() => clearField(3);

  @$pb.TagNumber(4)
  $core.int get port => $_getIZ(3);
  @$pb.TagNumber(4)
  set port($core.int v) { $_setUnsignedInt32(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasPort() => $_has(3);
  @$pb.TagNumber(4)
  void clearPort() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get keepAlive => $_getBF(4);
  @$pb.TagNumber(5)
  set keepAlive($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasKeepAlive() => $_has(4);
  @$pb.TagNumber(5)
  void clearKeepAlive() => clearField(5);
}

class AdpRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouter', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOS(1, 'rulesDirName', protoName: 'rulesDirName')
    ..aOM<$4.StringValue>(2, 'blocked', subBuilder: $4.StringValue.create)
    ..aOM<$4.StringValue>(3, 'unblocked', subBuilder: $4.StringValue.create)
    ..aOB(4, 'etcHostsIpAsBlocked', protoName: 'etcHostsIpAsBlocked')
    ..aOB(5, 'dev')
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
  $core.String get rulesDirName => $_getSZ(0);
  @$pb.TagNumber(1)
  set rulesDirName($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasRulesDirName() => $_has(0);
  @$pb.TagNumber(1)
  void clearRulesDirName() => clearField(1);

  @$pb.TagNumber(2)
  $4.StringValue get blocked => $_getN(1);
  @$pb.TagNumber(2)
  set blocked($4.StringValue v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasBlocked() => $_has(1);
  @$pb.TagNumber(2)
  void clearBlocked() => clearField(2);
  @$pb.TagNumber(2)
  $4.StringValue ensureBlocked() => $_ensure(1);

  @$pb.TagNumber(3)
  $4.StringValue get unblocked => $_getN(2);
  @$pb.TagNumber(3)
  set unblocked($4.StringValue v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasUnblocked() => $_has(2);
  @$pb.TagNumber(3)
  void clearUnblocked() => clearField(3);
  @$pb.TagNumber(3)
  $4.StringValue ensureUnblocked() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.bool get etcHostsIpAsBlocked => $_getBF(3);
  @$pb.TagNumber(4)
  set etcHostsIpAsBlocked($core.bool v) { $_setBool(3, v); }
  @$pb.TagNumber(4)
  $core.bool hasEtcHostsIpAsBlocked() => $_has(3);
  @$pb.TagNumber(4)
  void clearEtcHostsIpAsBlocked() => clearField(4);

  @$pb.TagNumber(5)
  $core.bool get dev => $_getBF(4);
  @$pb.TagNumber(5)
  set dev($core.bool v) { $_setBool(4, v); }
  @$pb.TagNumber(5)
  $core.bool hasDev() => $_has(4);
  @$pb.TagNumber(5)
  void clearDev() => clearField(5);
}

class IPNetRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouter', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..pc<$4.BytesValue>(1, 'ip', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..pc<$4.BytesValue>(2, 'netv4', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..pc<$4.BytesValue>(3, 'netv6', $pb.PbFieldType.PM, subBuilder: $4.BytesValue.create)
    ..aOM<$4.StringValue>(4, 'matched', subBuilder: $4.StringValue.create)
    ..aOM<$4.StringValue>(5, 'unmatched', subBuilder: $4.StringValue.create)
    ..aOM<$4.StringValue>(6, 'fileTest', protoName: 'fileTest', subBuilder: $4.StringValue.create)
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
  $core.List<$4.BytesValue> get ip => $_getList(0);

  @$pb.TagNumber(2)
  $core.List<$4.BytesValue> get netv4 => $_getList(1);

  @$pb.TagNumber(3)
  $core.List<$4.BytesValue> get netv6 => $_getList(2);

  @$pb.TagNumber(4)
  $4.StringValue get matched => $_getN(3);
  @$pb.TagNumber(4)
  set matched($4.StringValue v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasMatched() => $_has(3);
  @$pb.TagNumber(4)
  void clearMatched() => clearField(4);
  @$pb.TagNumber(4)
  $4.StringValue ensureMatched() => $_ensure(3);

  @$pb.TagNumber(5)
  $4.StringValue get unmatched => $_getN(4);
  @$pb.TagNumber(5)
  set unmatched($4.StringValue v) { setField(5, v); }
  @$pb.TagNumber(5)
  $core.bool hasUnmatched() => $_has(4);
  @$pb.TagNumber(5)
  void clearUnmatched() => clearField(5);
  @$pb.TagNumber(5)
  $4.StringValue ensureUnmatched() => $_ensure(4);

  @$pb.TagNumber(6)
  $4.StringValue get fileTest => $_getN(5);
  @$pb.TagNumber(6)
  set fileTest($4.StringValue v) { setField(6, v); }
  @$pb.TagNumber(6)
  $core.bool hasFileTest() => $_has(5);
  @$pb.TagNumber(6)
  void clearFileTest() => clearField(6);
  @$pb.TagNumber(6)
  $4.StringValue ensureFileTest() => $_ensure(5);
}

enum RouterItem_Router {
  adp, 
  ipnet, 
  notSet
}

class RouterItem extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, RouterItem_Router> _RouterItem_RouterByTag = {
    3 : RouterItem_Router.adp,
    4 : RouterItem_Router.ipnet,
    0 : RouterItem_Router.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterItem', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..oo(0, [3, 4])
    ..aOB(1, 'disabled')
    ..aOS(2, 'name')
    ..aOM<AdpRouter>(3, 'adp', subBuilder: AdpRouter.create)
    ..aOM<IPNetRouter>(4, 'ipnet', subBuilder: IPNetRouter.create)
    ..hasRequiredFields = false
  ;

  RouterItem._() : super();
  factory RouterItem() => create();
  factory RouterItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory RouterItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  RouterItem clone() => RouterItem()..mergeFromMessage(this);
  RouterItem copyWith(void Function(RouterItem) updates) => super.copyWith((message) => updates(message as RouterItem));
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static RouterItem create() => RouterItem._();
  RouterItem createEmptyInstance() => create();
  static $pb.PbList<RouterItem> createRepeated() => $pb.PbList<RouterItem>();
  @$core.pragma('dart2js:noInline')
  static RouterItem getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<RouterItem>(create);
  static RouterItem _defaultInstance;

  RouterItem_Router whichRouter() => _RouterItem_RouterByTag[$_whichOneof(0)];
  void clearRouter() => clearField($_whichOneof(0));

  @$pb.TagNumber(1)
  $core.bool get disabled => $_getBF(0);
  @$pb.TagNumber(1)
  set disabled($core.bool v) { $_setBool(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasDisabled() => $_has(0);
  @$pb.TagNumber(1)
  void clearDisabled() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get name => $_getSZ(1);
  @$pb.TagNumber(2)
  set name($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasName() => $_has(1);
  @$pb.TagNumber(2)
  void clearName() => clearField(2);

  @$pb.TagNumber(3)
  AdpRouter get adp => $_getN(2);
  @$pb.TagNumber(3)
  set adp(AdpRouter v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasAdp() => $_has(2);
  @$pb.TagNumber(3)
  void clearAdp() => clearField(3);
  @$pb.TagNumber(3)
  AdpRouter ensureAdp() => $_ensure(2);

  @$pb.TagNumber(4)
  IPNetRouter get ipnet => $_getN(3);
  @$pb.TagNumber(4)
  set ipnet(IPNetRouter v) { setField(4, v); }
  @$pb.TagNumber(4)
  $core.bool hasIpnet() => $_has(3);
  @$pb.TagNumber(4)
  void clearIpnet() => clearField(4);
  @$pb.TagNumber(4)
  IPNetRouter ensureIpnet() => $_ensure(3);
}

class Config extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Config', package: const $pb.PackageName('hybrid'), createEmptyInstance: create)
    ..aOM<Basic>(1, 'basic', subBuilder: Basic.create)
    ..aOM<Log>(2, 'log', subBuilder: Log.create)
    ..aOM<Ipfs>(3, 'ipfs', subBuilder: Ipfs.create)
    ..pc<IpfsServer>(4, 'ipfsServers', $pb.PbFieldType.PM, protoName: 'ipfsServers', subBuilder: IpfsServer.create)
    ..pc<FileServer>(5, 'fileServers', $pb.PbFieldType.PM, protoName: 'fileServers', subBuilder: FileServer.create)
    ..pc<HttpProxyServer>(6, 'httpProxyServers', $pb.PbFieldType.PM, protoName: 'httpProxyServers', subBuilder: HttpProxyServer.create)
    ..pc<RouterItem>(7, 'routers', $pb.PbFieldType.PM, subBuilder: RouterItem.create)
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
  Log get log => $_getN(1);
  @$pb.TagNumber(2)
  set log(Log v) { setField(2, v); }
  @$pb.TagNumber(2)
  $core.bool hasLog() => $_has(1);
  @$pb.TagNumber(2)
  void clearLog() => clearField(2);
  @$pb.TagNumber(2)
  Log ensureLog() => $_ensure(1);

  @$pb.TagNumber(3)
  Ipfs get ipfs => $_getN(2);
  @$pb.TagNumber(3)
  set ipfs(Ipfs v) { setField(3, v); }
  @$pb.TagNumber(3)
  $core.bool hasIpfs() => $_has(2);
  @$pb.TagNumber(3)
  void clearIpfs() => clearField(3);
  @$pb.TagNumber(3)
  Ipfs ensureIpfs() => $_ensure(2);

  @$pb.TagNumber(4)
  $core.List<IpfsServer> get ipfsServers => $_getList(3);

  @$pb.TagNumber(5)
  $core.List<FileServer> get fileServers => $_getList(4);

  @$pb.TagNumber(6)
  $core.List<HttpProxyServer> get httpProxyServers => $_getList(5);

  @$pb.TagNumber(7)
  $core.List<RouterItem> get routers => $_getList(6);
}

