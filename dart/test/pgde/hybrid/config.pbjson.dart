///
//  Generated code. Do not modify.
//  source: hybrid/config.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const ConfigTree$json = const {
  '1': 'ConfigTree',
  '2': const [
    const {'1': 'version', '3': 1, '4': 1, '5': 9, '10': 'version'},
    const {'1': 'rootName', '3': 2, '4': 1, '5': 9, '10': 'rootName'},
    const {'1': 'rootPath', '3': 3, '4': 1, '5': 9, '10': 'rootPath'},
    const {'1': 'configName', '3': 4, '4': 1, '5': 9, '10': 'configName'},
    const {'1': 'configPath', '3': 5, '4': 1, '5': 9, '10': 'configPath'},
    const {'1': 'ipfsName', '3': 6, '4': 1, '5': 9, '10': 'ipfsName'},
    const {'1': 'ipfsPath', '3': 7, '4': 1, '5': 9, '10': 'ipfsPath'},
    const {'1': 'storeName', '3': 8, '4': 1, '5': 9, '10': 'storeName'},
    const {'1': 'storePath', '3': 9, '4': 1, '5': 9, '10': 'storePath'},
    const {'1': 'filesRootName', '3': 10, '4': 1, '5': 9, '10': 'filesRootName'},
    const {'1': 'filesRootPath', '3': 11, '4': 1, '5': 9, '10': 'filesRootPath'},
    const {'1': 'rulesRootName', '3': 12, '4': 1, '5': 9, '10': 'rulesRootName'},
    const {'1': 'rulesRootPath', '3': 13, '4': 1, '5': 9, '10': 'rulesRootPath'},
  ],
  '7': const {},
};

const Basic$json = const {
  '1': 'Basic',
  '2': const [
    const {'1': 'version', '3': 2, '4': 1, '5': 9, '8': const {}, '10': 'version'},
    const {'1': 'dev', '3': 3, '4': 1, '5': 8, '10': 'dev'},
    const {'1': 'enableBind', '3': 4, '4': 1, '5': 8, '10': 'enableBind'},
    const {'1': 'bindIp', '3': 5, '4': 1, '5': 12, '8': const {}, '10': 'bindIp'},
    const {'1': 'bindPort', '3': 6, '4': 1, '5': 13, '8': const {}, '10': 'bindPort'},
    const {'1': 'flushIntervalMs', '3': 7, '4': 1, '5': 13, '8': const {}, '10': 'flushIntervalMs'},
    const {'1': 'token', '3': 8, '4': 1, '5': 12, '8': const {}, '10': 'token'},
  ],
  '7': const {},
};

const Ipfs$json = const {
  '1': 'Ipfs',
  '2': const [
    const {'1': 'fakeApiListenIp', '3': 2, '4': 1, '5': 12, '8': const {}, '10': 'fakeApiListenIp'},
    const {'1': 'fakeApiListenPort', '3': 3, '4': 1, '5': 13, '8': const {}, '10': 'fakeApiListenPort'},
    const {'1': 'enableGateway', '3': 4, '4': 1, '5': 8, '10': 'enableGateway'},
    const {'1': 'gatewayServerName', '3': 5, '4': 1, '5': 9, '8': const {}, '10': 'gatewayServerName'},
    const {'1': 'enableApi', '3': 6, '4': 1, '5': 8, '10': 'enableApi'},
    const {'1': 'apiServerName', '3': 7, '4': 1, '5': 9, '8': const {}, '10': 'apiServerName'},
    const {'1': 'profile', '3': 8, '4': 3, '5': 9, '8': const {}, '10': 'profile'},
    const {'1': 'autoMigrate', '3': 9, '4': 1, '5': 8, '10': 'autoMigrate'},
    const {'1': 'enableIpnsPubSub', '3': 10, '4': 1, '5': 8, '10': 'enableIpnsPubSub'},
    const {'1': 'enablePubSub', '3': 11, '4': 1, '5': 8, '10': 'enablePubSub'},
    const {'1': 'enableMultiplex', '3': 12, '4': 1, '5': 8, '10': 'enableMultiplex'},
    const {'1': 'token', '3': 13, '4': 1, '5': 9, '8': const {}, '10': 'token'},
  ],
  '7': const {},
};

const Log$json = const {
  '1': 'Log',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'dev', '3': 2, '4': 1, '5': 8, '10': 'dev'},
    const {'1': 'level', '3': 3, '4': 1, '5': 9, '8': const {}, '10': 'level'},
    const {'1': 'target', '3': 4, '4': 1, '5': 9, '10': 'target'},
  ],
  '3': const [Log_View$json],
  '7': const {},
};

const Log_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'target', '3': 4, '4': 1, '5': 9, '10': 'target'},
  ],
  '7': const {},
};

const IpfsServer$json = const {
  '1': 'IpfsServer',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'disabled', '3': 2, '4': 1, '5': 8, '10': 'disabled'},
    const {'1': 'name', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '8': const {}, '10': 'name'},
    const {'1': 'peer', '3': 4, '4': 1, '5': 9, '8': const {}, '10': 'peer'},
    const {'1': 'token', '3': 5, '4': 1, '5': 9, '8': const {}, '10': 'token'},
  ],
  '3': const [IpfsServer_View$json],
  '7': const {},
};

const IpfsServer_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
  '7': const {},
};

const FileServer$json = const {
  '1': 'FileServer',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'disabled', '3': 2, '4': 1, '5': 8, '10': 'disabled'},
    const {'1': 'name', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '8': const {}, '10': 'name'},
    const {'1': 'zip', '3': 4, '4': 1, '5': 9, '8': const {}, '10': 'zip'},
    const {'1': 'redirect', '3': 5, '4': 3, '5': 11, '6': '.hybrid.FileServer.RedirectEntry', '8': const {}, '10': 'redirect'},
    const {'1': 'dev', '3': 6, '4': 1, '5': 8, '10': 'dev'},
  ],
  '3': const [FileServer_View$json, FileServer_RedirectEntry$json],
  '7': const {},
};

const FileServer_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
  '7': const {},
};

const FileServer_RedirectEntry$json = const {
  '1': 'RedirectEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 9, '10': 'value'},
  ],
  '7': const {'7': true},
};

const HttpProxyServer$json = const {
  '1': 'HttpProxyServer',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'disabled', '3': 2, '4': 1, '5': 8, '10': 'disabled'},
    const {'1': 'name', '3': 3, '4': 1, '5': 11, '6': '.google.protobuf.StringValue', '8': const {}, '10': 'name'},
    const {'1': 'host', '3': 4, '4': 1, '5': 9, '8': const {}, '10': 'host'},
    const {'1': 'port', '3': 5, '4': 1, '5': 13, '8': const {}, '10': 'port'},
    const {'1': 'keepAlive', '3': 6, '4': 1, '5': 8, '10': 'keepAlive'},
  ],
  '3': const [HttpProxyServer_View$json],
  '7': const {},
};

const HttpProxyServer_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
  '7': const {},
};

const ServerView$json = const {
  '1': 'ServerView',
  '2': const [
    const {'1': 'ipfs', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IpfsServer.View', '9': 0, '10': 'ipfs'},
    const {'1': 'file', '3': 2, '4': 1, '5': 11, '6': '.hybrid.FileServer.View', '9': 0, '10': 'file'},
    const {'1': 'http', '3': 3, '4': 1, '5': 11, '6': '.hybrid.HttpProxyServer.View', '9': 0, '10': 'http'},
  ],
  '7': const {},
  '8': const [
    const {'1': 'group'},
  ],
};

const AdpRouter$json = const {
  '1': 'AdpRouter',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'disabled', '3': 2, '4': 1, '5': 8, '10': 'disabled'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '8': const {}, '10': 'name'},
    const {'1': 'rulesDirName', '3': 4, '4': 1, '5': 9, '8': const {}, '10': 'rulesDirName'},
    const {'1': 'blocked', '3': 5, '4': 1, '5': 11, '6': '.hybrid.ServerView', '10': 'blocked'},
    const {'1': 'unblocked', '3': 6, '4': 1, '5': 11, '6': '.hybrid.ServerView', '10': 'unblocked'},
    const {'1': 'etcHostsIpAsBlocked', '3': 7, '4': 1, '5': 8, '10': 'etcHostsIpAsBlocked'},
    const {'1': 'dev', '3': 8, '4': 1, '5': 8, '10': 'dev'},
  ],
  '3': const [AdpRouter_View$json],
  '7': const {},
};

const AdpRouter_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
  '7': const {},
};

const IPNetRouter$json = const {
  '1': 'IPNetRouter',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'disabled', '3': 2, '4': 1, '5': 8, '10': 'disabled'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '8': const {}, '10': 'name'},
    const {'1': 'ip', '3': 4, '4': 3, '5': 11, '6': '.google.protobuf.BytesValue', '8': const {}, '10': 'ip'},
    const {'1': 'netv4', '3': 5, '4': 3, '5': 11, '6': '.google.protobuf.BytesValue', '8': const {}, '10': 'netv4'},
    const {'1': 'netv6', '3': 6, '4': 3, '5': 11, '6': '.google.protobuf.BytesValue', '8': const {}, '10': 'netv6'},
    const {'1': 'matched', '3': 7, '4': 1, '5': 11, '6': '.hybrid.ServerView', '8': const {}, '10': 'matched'},
    const {'1': 'unmatched', '3': 8, '4': 1, '5': 11, '6': '.hybrid.ServerView', '8': const {}, '10': 'unmatched'},
    const {'1': 'fileTest', '3': 9, '4': 1, '5': 11, '6': '.hybrid.FileServer.View', '8': const {}, '10': 'fileTest'},
  ],
  '3': const [IPNetRouter_View$json],
  '7': const {},
};

const IPNetRouter_View$json = const {
  '1': 'View',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '8': const {}, '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
  ],
  '7': const {},
};

const RouterUseView$json = const {
  '1': 'RouterUseView',
  '2': const [
    const {'1': 'adp', '3': 1, '4': 1, '5': 11, '6': '.hybrid.AdpRouter.View', '9': 0, '10': 'adp'},
    const {'1': 'ipnet', '3': 2, '4': 1, '5': 11, '6': '.hybrid.IPNetRouter.View', '9': 0, '10': 'ipnet'},
  ],
  '3': const [RouterUseView_Id$json, RouterUseView_Element$json],
  '7': const {},
  '8': const [
    const {'1': 'group'},
  ],
};

const RouterUseView_Id$json = const {
  '1': 'Id',
  '2': const [
    const {'1': 'adp', '3': 1, '4': 1, '5': 13, '9': 0, '10': 'adp'},
    const {'1': 'ipnet', '3': 2, '4': 1, '5': 13, '9': 0, '10': 'ipnet'},
  ],
  '8': const [
    const {'1': 'is'},
  ],
};

const RouterUseView_Element$json = const {
  '1': 'Element',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 11, '6': '.hybrid.RouterUseView.Id', '10': 'id'},
    const {'1': 'name', '3': 3, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'rank', '3': 4, '4': 1, '5': 3, '10': 'rank'},
  ],
  '7': const {},
};

const Config$json = const {
  '1': 'Config',
  '2': const [
    const {'1': 'basic', '3': 1, '4': 1, '5': 11, '6': '.hybrid.Basic', '8': const {}, '10': 'basic'},
    const {'1': 'ipfs', '3': 2, '4': 1, '5': 11, '6': '.hybrid.Ipfs', '8': const {}, '10': 'ipfs'},
    const {'1': 'log', '3': 3, '4': 1, '5': 11, '6': '.hybrid.Log.View', '8': const {}, '10': 'log'},
    const {'1': 'logs', '3': 4, '4': 3, '5': 11, '6': '.hybrid.Log.View', '8': const {}, '10': 'logs'},
    const {'1': 'servers', '3': 5, '4': 3, '5': 11, '6': '.hybrid.ServerView', '8': const {}, '10': 'servers'},
    const {'1': 'adps', '3': 6, '4': 3, '5': 11, '6': '.hybrid.AdpRouter.View', '8': const {}, '10': 'adps'},
    const {'1': 'ipnets', '3': 7, '4': 3, '5': 11, '6': '.hybrid.IPNetRouter.View', '8': const {}, '10': 'ipnets'},
    const {'1': 'routers', '3': 8, '4': 3, '5': 11, '6': '.hybrid.RouterUseView.Element', '8': const {}, '10': 'routers'},
  ],
  '7': const {},
};

