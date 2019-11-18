///
//  Generated code. Do not modify.
//  source: hybrid/config.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

import '../google/protobuf/wrappers.pb.dart' as $0;

class ConfigTree extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('ConfigTree', package: const $pb.PackageName('hybrid'))
    ..aOS(1, 'version')
    ..aOS(2, 'rootName')
    ..aOS(3, 'rootPath')
    ..aOS(4, 'configName')
    ..aOS(5, 'configPath')
    ..aOS(6, 'ipfsName')
    ..aOS(7, 'ipfsPath')
    ..aOS(8, 'storeName')
    ..aOS(9, 'storePath')
    ..aOS(10, 'filesRootName')
    ..aOS(11, 'filesRootPath')
    ..aOS(12, 'rulesRootName')
    ..aOS(13, 'rulesRootPath')
    ..hasRequiredFields = false
  ;

  ConfigTree() : super();
  ConfigTree.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  ConfigTree.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  ConfigTree clone() => ConfigTree()..mergeFromMessage(this);
  ConfigTree copyWith(void Function(ConfigTree) updates) => super.copyWith((message) => updates(message as ConfigTree));
  $pb.BuilderInfo get info_ => _i;
  static ConfigTree create() => ConfigTree();
  ConfigTree createEmptyInstance() => create();
  static $pb.PbList<ConfigTree> createRepeated() => $pb.PbList<ConfigTree>();
  static ConfigTree getDefault() => _defaultInstance ??= create()..freeze();
  static ConfigTree _defaultInstance;

  $core.String get version => $_getS(0, '');
  set version($core.String v) { $_setString(0, v); }
  $core.bool hasVersion() => $_has(0);
  void clearVersion() => clearField(1);

  $core.String get rootName => $_getS(1, '');
  set rootName($core.String v) { $_setString(1, v); }
  $core.bool hasRootName() => $_has(1);
  void clearRootName() => clearField(2);

  $core.String get rootPath => $_getS(2, '');
  set rootPath($core.String v) { $_setString(2, v); }
  $core.bool hasRootPath() => $_has(2);
  void clearRootPath() => clearField(3);

  $core.String get configName => $_getS(3, '');
  set configName($core.String v) { $_setString(3, v); }
  $core.bool hasConfigName() => $_has(3);
  void clearConfigName() => clearField(4);

  $core.String get configPath => $_getS(4, '');
  set configPath($core.String v) { $_setString(4, v); }
  $core.bool hasConfigPath() => $_has(4);
  void clearConfigPath() => clearField(5);

  $core.String get ipfsName => $_getS(5, '');
  set ipfsName($core.String v) { $_setString(5, v); }
  $core.bool hasIpfsName() => $_has(5);
  void clearIpfsName() => clearField(6);

  $core.String get ipfsPath => $_getS(6, '');
  set ipfsPath($core.String v) { $_setString(6, v); }
  $core.bool hasIpfsPath() => $_has(6);
  void clearIpfsPath() => clearField(7);

  $core.String get storeName => $_getS(7, '');
  set storeName($core.String v) { $_setString(7, v); }
  $core.bool hasStoreName() => $_has(7);
  void clearStoreName() => clearField(8);

  $core.String get storePath => $_getS(8, '');
  set storePath($core.String v) { $_setString(8, v); }
  $core.bool hasStorePath() => $_has(8);
  void clearStorePath() => clearField(9);

  $core.String get filesRootName => $_getS(9, '');
  set filesRootName($core.String v) { $_setString(9, v); }
  $core.bool hasFilesRootName() => $_has(9);
  void clearFilesRootName() => clearField(10);

  $core.String get filesRootPath => $_getS(10, '');
  set filesRootPath($core.String v) { $_setString(10, v); }
  $core.bool hasFilesRootPath() => $_has(10);
  void clearFilesRootPath() => clearField(11);

  $core.String get rulesRootName => $_getS(11, '');
  set rulesRootName($core.String v) { $_setString(11, v); }
  $core.bool hasRulesRootName() => $_has(11);
  void clearRulesRootName() => clearField(12);

  $core.String get rulesRootPath => $_getS(12, '');
  set rulesRootPath($core.String v) { $_setString(12, v); }
  $core.bool hasRulesRootPath() => $_has(12);
  void clearRulesRootPath() => clearField(13);
}

class Basic extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Basic', package: const $pb.PackageName('hybrid'))
    ..aOS(1, 'version')
    ..aOB(2, 'dev')
    ..aOB(3, 'enableBind')
    ..aOS(4, 'bindIp')
    ..a<$core.int>(5, 'bindPort', $pb.PbFieldType.OU3)
    ..a<$core.int>(6, 'flushIntervalMs', $pb.PbFieldType.OU3)
    ..aOS(7, 'token')
    ..hasRequiredFields = false
  ;

  Basic() : super();
  Basic.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Basic.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Basic clone() => Basic()..mergeFromMessage(this);
  Basic copyWith(void Function(Basic) updates) => super.copyWith((message) => updates(message as Basic));
  $pb.BuilderInfo get info_ => _i;
  static Basic create() => Basic();
  Basic createEmptyInstance() => create();
  static $pb.PbList<Basic> createRepeated() => $pb.PbList<Basic>();
  static Basic getDefault() => _defaultInstance ??= create()..freeze();
  static Basic _defaultInstance;

  $core.String get version => $_getS(0, '');
  set version($core.String v) { $_setString(0, v); }
  $core.bool hasVersion() => $_has(0);
  void clearVersion() => clearField(1);

  $core.bool get dev => $_get(1, false);
  set dev($core.bool v) { $_setBool(1, v); }
  $core.bool hasDev() => $_has(1);
  void clearDev() => clearField(2);

  $core.bool get enableBind => $_get(2, false);
  set enableBind($core.bool v) { $_setBool(2, v); }
  $core.bool hasEnableBind() => $_has(2);
  void clearEnableBind() => clearField(3);

  $core.String get bindIp => $_getS(3, '');
  set bindIp($core.String v) { $_setString(3, v); }
  $core.bool hasBindIp() => $_has(3);
  void clearBindIp() => clearField(4);

  $core.int get bindPort => $_get(4, 0);
  set bindPort($core.int v) { $_setUnsignedInt32(4, v); }
  $core.bool hasBindPort() => $_has(4);
  void clearBindPort() => clearField(5);

  $core.int get flushIntervalMs => $_get(5, 0);
  set flushIntervalMs($core.int v) { $_setUnsignedInt32(5, v); }
  $core.bool hasFlushIntervalMs() => $_has(5);
  void clearFlushIntervalMs() => clearField(6);

  $core.String get token => $_getS(6, '');
  set token($core.String v) { $_setString(6, v); }
  $core.bool hasToken() => $_has(6);
  void clearToken() => clearField(7);
}

class Log extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Log', package: const $pb.PackageName('hybrid'))
    ..aOB(1, 'dev')
    ..aOS(2, 'level')
    ..aOS(3, 'target')
    ..hasRequiredFields = false
  ;

  Log() : super();
  Log.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Log.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Log clone() => Log()..mergeFromMessage(this);
  Log copyWith(void Function(Log) updates) => super.copyWith((message) => updates(message as Log));
  $pb.BuilderInfo get info_ => _i;
  static Log create() => Log();
  Log createEmptyInstance() => create();
  static $pb.PbList<Log> createRepeated() => $pb.PbList<Log>();
  static Log getDefault() => _defaultInstance ??= create()..freeze();
  static Log _defaultInstance;

  $core.bool get dev => $_get(0, false);
  set dev($core.bool v) { $_setBool(0, v); }
  $core.bool hasDev() => $_has(0);
  void clearDev() => clearField(1);

  $core.String get level => $_getS(1, '');
  set level($core.String v) { $_setString(1, v); }
  $core.bool hasLevel() => $_has(1);
  void clearLevel() => clearField(2);

  $core.String get target => $_getS(2, '');
  set target($core.String v) { $_setString(2, v); }
  $core.bool hasTarget() => $_has(2);
  void clearTarget() => clearField(3);
}

class Ipfs extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Ipfs', package: const $pb.PackageName('hybrid'))
    ..aOS(1, 'fakeApiListenIp')
    ..a<$core.int>(2, 'fakeApiListenPort', $pb.PbFieldType.OU3)
    ..aOB(3, 'enableGateway')
    ..aOS(4, 'gatewayServerName')
    ..aOB(5, 'enableApi')
    ..aOS(6, 'apiServerName')
    ..pPS(7, 'profile')
    ..aOB(8, 'autoMigrate')
    ..aOB(9, 'enableIpnsPubSub')
    ..aOB(10, 'enablePubSub')
    ..aOB(11, 'enableMultiplex')
    ..aOS(12, 'token')
    ..hasRequiredFields = false
  ;

  Ipfs() : super();
  Ipfs.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Ipfs.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Ipfs clone() => Ipfs()..mergeFromMessage(this);
  Ipfs copyWith(void Function(Ipfs) updates) => super.copyWith((message) => updates(message as Ipfs));
  $pb.BuilderInfo get info_ => _i;
  static Ipfs create() => Ipfs();
  Ipfs createEmptyInstance() => create();
  static $pb.PbList<Ipfs> createRepeated() => $pb.PbList<Ipfs>();
  static Ipfs getDefault() => _defaultInstance ??= create()..freeze();
  static Ipfs _defaultInstance;

  $core.String get fakeApiListenIp => $_getS(0, '');
  set fakeApiListenIp($core.String v) { $_setString(0, v); }
  $core.bool hasFakeApiListenIp() => $_has(0);
  void clearFakeApiListenIp() => clearField(1);

  $core.int get fakeApiListenPort => $_get(1, 0);
  set fakeApiListenPort($core.int v) { $_setUnsignedInt32(1, v); }
  $core.bool hasFakeApiListenPort() => $_has(1);
  void clearFakeApiListenPort() => clearField(2);

  $core.bool get enableGateway => $_get(2, false);
  set enableGateway($core.bool v) { $_setBool(2, v); }
  $core.bool hasEnableGateway() => $_has(2);
  void clearEnableGateway() => clearField(3);

  $core.String get gatewayServerName => $_getS(3, '');
  set gatewayServerName($core.String v) { $_setString(3, v); }
  $core.bool hasGatewayServerName() => $_has(3);
  void clearGatewayServerName() => clearField(4);

  $core.bool get enableApi => $_get(4, false);
  set enableApi($core.bool v) { $_setBool(4, v); }
  $core.bool hasEnableApi() => $_has(4);
  void clearEnableApi() => clearField(5);

  $core.String get apiServerName => $_getS(5, '');
  set apiServerName($core.String v) { $_setString(5, v); }
  $core.bool hasApiServerName() => $_has(5);
  void clearApiServerName() => clearField(6);

  $core.List<$core.String> get profile => $_getList(6);

  $core.bool get autoMigrate => $_get(7, false);
  set autoMigrate($core.bool v) { $_setBool(7, v); }
  $core.bool hasAutoMigrate() => $_has(7);
  void clearAutoMigrate() => clearField(8);

  $core.bool get enableIpnsPubSub => $_get(8, false);
  set enableIpnsPubSub($core.bool v) { $_setBool(8, v); }
  $core.bool hasEnableIpnsPubSub() => $_has(8);
  void clearEnableIpnsPubSub() => clearField(9);

  $core.bool get enablePubSub => $_get(9, false);
  set enablePubSub($core.bool v) { $_setBool(9, v); }
  $core.bool hasEnablePubSub() => $_has(9);
  void clearEnablePubSub() => clearField(10);

  $core.bool get enableMultiplex => $_get(10, false);
  set enableMultiplex($core.bool v) { $_setBool(10, v); }
  $core.bool hasEnableMultiplex() => $_has(10);
  void clearEnableMultiplex() => clearField(11);

  $core.String get token => $_getS(11, '');
  set token($core.String v) { $_setString(11, v); }
  $core.bool hasToken() => $_has(11);
  void clearToken() => clearField(12);
}

class IpfsServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IpfsServer', package: const $pb.PackageName('hybrid'))
    ..aOB(1, 'disabled')
    ..a<$0.StringValue>(2, 'name', $pb.PbFieldType.OM, $0.StringValue.getDefault, $0.StringValue.create)
    ..aOS(3, 'peer')
    ..aOS(4, 'token')
    ..hasRequiredFields = false
  ;

  IpfsServer() : super();
  IpfsServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  IpfsServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  IpfsServer clone() => IpfsServer()..mergeFromMessage(this);
  IpfsServer copyWith(void Function(IpfsServer) updates) => super.copyWith((message) => updates(message as IpfsServer));
  $pb.BuilderInfo get info_ => _i;
  static IpfsServer create() => IpfsServer();
  IpfsServer createEmptyInstance() => create();
  static $pb.PbList<IpfsServer> createRepeated() => $pb.PbList<IpfsServer>();
  static IpfsServer getDefault() => _defaultInstance ??= create()..freeze();
  static IpfsServer _defaultInstance;

  $core.bool get disabled => $_get(0, false);
  set disabled($core.bool v) { $_setBool(0, v); }
  $core.bool hasDisabled() => $_has(0);
  void clearDisabled() => clearField(1);

  $0.StringValue get name => $_getN(1);
  set name($0.StringValue v) { setField(2, v); }
  $core.bool hasName() => $_has(1);
  void clearName() => clearField(2);

  $core.String get peer => $_getS(2, '');
  set peer($core.String v) { $_setString(2, v); }
  $core.bool hasPeer() => $_has(2);
  void clearPeer() => clearField(3);

  $core.String get token => $_getS(3, '');
  set token($core.String v) { $_setString(3, v); }
  $core.bool hasToken() => $_has(3);
  void clearToken() => clearField(4);
}

class FileServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('FileServer', package: const $pb.PackageName('hybrid'))
    ..aOB(1, 'disabled')
    ..a<$0.StringValue>(2, 'name', $pb.PbFieldType.OM, $0.StringValue.getDefault, $0.StringValue.create)
    ..aOS(3, 'zip')
    ..m<$core.String, $core.String>(4, 'redirect', 'FileServer.RedirectEntry',$pb.PbFieldType.OS, $pb.PbFieldType.OS, null, null, null , const $pb.PackageName('hybrid'))
    ..aOB(5, 'dev')
    ..hasRequiredFields = false
  ;

  FileServer() : super();
  FileServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  FileServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  FileServer clone() => FileServer()..mergeFromMessage(this);
  FileServer copyWith(void Function(FileServer) updates) => super.copyWith((message) => updates(message as FileServer));
  $pb.BuilderInfo get info_ => _i;
  static FileServer create() => FileServer();
  FileServer createEmptyInstance() => create();
  static $pb.PbList<FileServer> createRepeated() => $pb.PbList<FileServer>();
  static FileServer getDefault() => _defaultInstance ??= create()..freeze();
  static FileServer _defaultInstance;

  $core.bool get disabled => $_get(0, false);
  set disabled($core.bool v) { $_setBool(0, v); }
  $core.bool hasDisabled() => $_has(0);
  void clearDisabled() => clearField(1);

  $0.StringValue get name => $_getN(1);
  set name($0.StringValue v) { setField(2, v); }
  $core.bool hasName() => $_has(1);
  void clearName() => clearField(2);

  $core.String get zip => $_getS(2, '');
  set zip($core.String v) { $_setString(2, v); }
  $core.bool hasZip() => $_has(2);
  void clearZip() => clearField(3);

  $core.Map<$core.String, $core.String> get redirect => $_getMap(3);

  $core.bool get dev => $_get(4, false);
  set dev($core.bool v) { $_setBool(4, v); }
  $core.bool hasDev() => $_has(4);
  void clearDev() => clearField(5);
}

class HttpProxyServer extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('HttpProxyServer', package: const $pb.PackageName('hybrid'))
    ..aOB(1, 'disabled')
    ..a<$0.StringValue>(2, 'name', $pb.PbFieldType.OM, $0.StringValue.getDefault, $0.StringValue.create)
    ..aOS(3, 'host')
    ..a<$core.int>(4, 'port', $pb.PbFieldType.OU3)
    ..aOB(5, 'keepAlive')
    ..hasRequiredFields = false
  ;

  HttpProxyServer() : super();
  HttpProxyServer.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  HttpProxyServer.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  HttpProxyServer clone() => HttpProxyServer()..mergeFromMessage(this);
  HttpProxyServer copyWith(void Function(HttpProxyServer) updates) => super.copyWith((message) => updates(message as HttpProxyServer));
  $pb.BuilderInfo get info_ => _i;
  static HttpProxyServer create() => HttpProxyServer();
  HttpProxyServer createEmptyInstance() => create();
  static $pb.PbList<HttpProxyServer> createRepeated() => $pb.PbList<HttpProxyServer>();
  static HttpProxyServer getDefault() => _defaultInstance ??= create()..freeze();
  static HttpProxyServer _defaultInstance;

  $core.bool get disabled => $_get(0, false);
  set disabled($core.bool v) { $_setBool(0, v); }
  $core.bool hasDisabled() => $_has(0);
  void clearDisabled() => clearField(1);

  $0.StringValue get name => $_getN(1);
  set name($0.StringValue v) { setField(2, v); }
  $core.bool hasName() => $_has(1);
  void clearName() => clearField(2);

  $core.String get host => $_getS(2, '');
  set host($core.String v) { $_setString(2, v); }
  $core.bool hasHost() => $_has(2);
  void clearHost() => clearField(3);

  $core.int get port => $_get(3, 0);
  set port($core.int v) { $_setUnsignedInt32(3, v); }
  $core.bool hasPort() => $_has(3);
  void clearPort() => clearField(4);

  $core.bool get keepAlive => $_get(4, false);
  set keepAlive($core.bool v) { $_setBool(4, v); }
  $core.bool hasKeepAlive() => $_has(4);
  void clearKeepAlive() => clearField(5);
}

class AdpRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AdpRouter', package: const $pb.PackageName('hybrid'))
    ..aOS(1, 'rulesDirName')
    ..a<$0.StringValue>(2, 'blocked', $pb.PbFieldType.OM, $0.StringValue.getDefault, $0.StringValue.create)
    ..a<$0.StringValue>(3, 'unblocked', $pb.PbFieldType.OM, $0.StringValue.getDefault, $0.StringValue.create)
    ..aOB(4, 'etcHostsIpAsBlocked')
    ..aOB(5, 'dev')
    ..hasRequiredFields = false
  ;

  AdpRouter() : super();
  AdpRouter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  AdpRouter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  AdpRouter clone() => AdpRouter()..mergeFromMessage(this);
  AdpRouter copyWith(void Function(AdpRouter) updates) => super.copyWith((message) => updates(message as AdpRouter));
  $pb.BuilderInfo get info_ => _i;
  static AdpRouter create() => AdpRouter();
  AdpRouter createEmptyInstance() => create();
  static $pb.PbList<AdpRouter> createRepeated() => $pb.PbList<AdpRouter>();
  static AdpRouter getDefault() => _defaultInstance ??= create()..freeze();
  static AdpRouter _defaultInstance;

  $core.String get rulesDirName => $_getS(0, '');
  set rulesDirName($core.String v) { $_setString(0, v); }
  $core.bool hasRulesDirName() => $_has(0);
  void clearRulesDirName() => clearField(1);

  $0.StringValue get blocked => $_getN(1);
  set blocked($0.StringValue v) { setField(2, v); }
  $core.bool hasBlocked() => $_has(1);
  void clearBlocked() => clearField(2);

  $0.StringValue get unblocked => $_getN(2);
  set unblocked($0.StringValue v) { setField(3, v); }
  $core.bool hasUnblocked() => $_has(2);
  void clearUnblocked() => clearField(3);

  $core.bool get etcHostsIpAsBlocked => $_get(3, false);
  set etcHostsIpAsBlocked($core.bool v) { $_setBool(3, v); }
  $core.bool hasEtcHostsIpAsBlocked() => $_has(3);
  void clearEtcHostsIpAsBlocked() => clearField(4);

  $core.bool get dev => $_get(4, false);
  set dev($core.bool v) { $_setBool(4, v); }
  $core.bool hasDev() => $_has(4);
  void clearDev() => clearField(5);
}

class IPNetRouter extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('IPNetRouter', package: const $pb.PackageName('hybrid'))
    ..pPS(1, 'ip')
    ..pPS(2, 'net')
    ..aOS(3, 'matched')
    ..aOS(4, 'unmatched')
    ..aOS(5, 'fileTest')
    ..hasRequiredFields = false
  ;

  IPNetRouter() : super();
  IPNetRouter.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  IPNetRouter.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  IPNetRouter clone() => IPNetRouter()..mergeFromMessage(this);
  IPNetRouter copyWith(void Function(IPNetRouter) updates) => super.copyWith((message) => updates(message as IPNetRouter));
  $pb.BuilderInfo get info_ => _i;
  static IPNetRouter create() => IPNetRouter();
  IPNetRouter createEmptyInstance() => create();
  static $pb.PbList<IPNetRouter> createRepeated() => $pb.PbList<IPNetRouter>();
  static IPNetRouter getDefault() => _defaultInstance ??= create()..freeze();
  static IPNetRouter _defaultInstance;

  $core.List<$core.String> get ip => $_getList(0);

  $core.List<$core.String> get net => $_getList(1);

  $core.String get matched => $_getS(2, '');
  set matched($core.String v) { $_setString(2, v); }
  $core.bool hasMatched() => $_has(2);
  void clearMatched() => clearField(3);

  $core.String get unmatched => $_getS(3, '');
  set unmatched($core.String v) { $_setString(3, v); }
  $core.bool hasUnmatched() => $_has(3);
  void clearUnmatched() => clearField(4);

  $core.String get fileTest => $_getS(4, '');
  set fileTest($core.String v) { $_setString(4, v); }
  $core.bool hasFileTest() => $_has(4);
  void clearFileTest() => clearField(5);
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
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('RouterItem', package: const $pb.PackageName('hybrid'))
    ..aOB(1, 'disabled')
    ..aOS(2, 'name')
    ..a<AdpRouter>(3, 'adp', $pb.PbFieldType.OM, AdpRouter.getDefault, AdpRouter.create)
    ..a<IPNetRouter>(4, 'ipnet', $pb.PbFieldType.OM, IPNetRouter.getDefault, IPNetRouter.create)
    ..oo(0, [3, 4])
    ..hasRequiredFields = false
  ;

  RouterItem() : super();
  RouterItem.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  RouterItem.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  RouterItem clone() => RouterItem()..mergeFromMessage(this);
  RouterItem copyWith(void Function(RouterItem) updates) => super.copyWith((message) => updates(message as RouterItem));
  $pb.BuilderInfo get info_ => _i;
  static RouterItem create() => RouterItem();
  RouterItem createEmptyInstance() => create();
  static $pb.PbList<RouterItem> createRepeated() => $pb.PbList<RouterItem>();
  static RouterItem getDefault() => _defaultInstance ??= create()..freeze();
  static RouterItem _defaultInstance;

  RouterItem_Router whichRouter() => _RouterItem_RouterByTag[$_whichOneof(0)];
  void clearRouter() => clearField($_whichOneof(0));

  $core.bool get disabled => $_get(0, false);
  set disabled($core.bool v) { $_setBool(0, v); }
  $core.bool hasDisabled() => $_has(0);
  void clearDisabled() => clearField(1);

  $core.String get name => $_getS(1, '');
  set name($core.String v) { $_setString(1, v); }
  $core.bool hasName() => $_has(1);
  void clearName() => clearField(2);

  AdpRouter get adp => $_getN(2);
  set adp(AdpRouter v) { setField(3, v); }
  $core.bool hasAdp() => $_has(2);
  void clearAdp() => clearField(3);

  IPNetRouter get ipnet => $_getN(3);
  set ipnet(IPNetRouter v) { setField(4, v); }
  $core.bool hasIpnet() => $_has(3);
  void clearIpnet() => clearField(4);
}

class Config extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Config', package: const $pb.PackageName('hybrid'))
    ..a<Basic>(1, 'basic', $pb.PbFieldType.OM, Basic.getDefault, Basic.create)
    ..a<Log>(2, 'log', $pb.PbFieldType.OM, Log.getDefault, Log.create)
    ..a<Ipfs>(3, 'ipfs', $pb.PbFieldType.OM, Ipfs.getDefault, Ipfs.create)
    ..pc<IpfsServer>(4, 'ipfsServers', $pb.PbFieldType.PM,IpfsServer.create)
    ..pc<FileServer>(5, 'fileServers', $pb.PbFieldType.PM,FileServer.create)
    ..pc<HttpProxyServer>(6, 'httpProxyServers', $pb.PbFieldType.PM,HttpProxyServer.create)
    ..pc<RouterItem>(7, 'routers', $pb.PbFieldType.PM,RouterItem.create)
    ..hasRequiredFields = false
  ;

  Config() : super();
  Config.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Config.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Config clone() => Config()..mergeFromMessage(this);
  Config copyWith(void Function(Config) updates) => super.copyWith((message) => updates(message as Config));
  $pb.BuilderInfo get info_ => _i;
  static Config create() => Config();
  Config createEmptyInstance() => create();
  static $pb.PbList<Config> createRepeated() => $pb.PbList<Config>();
  static Config getDefault() => _defaultInstance ??= create()..freeze();
  static Config _defaultInstance;

  Basic get basic => $_getN(0);
  set basic(Basic v) { setField(1, v); }
  $core.bool hasBasic() => $_has(0);
  void clearBasic() => clearField(1);

  Log get log => $_getN(1);
  set log(Log v) { setField(2, v); }
  $core.bool hasLog() => $_has(1);
  void clearLog() => clearField(2);

  Ipfs get ipfs => $_getN(2);
  set ipfs(Ipfs v) { setField(3, v); }
  $core.bool hasIpfs() => $_has(2);
  void clearIpfs() => clearField(3);

  $core.List<IpfsServer> get ipfsServers => $_getList(3);

  $core.List<FileServer> get fileServers => $_getList(4);

  $core.List<HttpProxyServer> get httpProxyServers => $_getList(5);

  $core.List<RouterItem> get routers => $_getList(6);
}

