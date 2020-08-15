///
//  Generated code. Do not modify.
//  source: hybrid/config.rpc.proto
//
// @dart = 2.3
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name,return_of_invalid_type

const BasicPayload$json = const {
  '1': 'BasicPayload',
  '3': const [BasicPayload_GetResp$json],
};

const BasicPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.Basic', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const IpfsPayload$json = const {
  '1': 'IpfsPayload',
  '3': const [IpfsPayload_GetResp$json],
};

const IpfsPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.Ipfs', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const LogPayload$json = const {
  '1': 'LogPayload',
  '3': const [LogPayload_SrcId$json, LogPayload_GetResp$json, LogPayload_View$json],
};

const LogPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const LogPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.Log', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const LogPayload_View$json = const {
  '1': 'View',
  '3': const [LogPayload_View_SrcResp$json, LogPayload_View_AddResp$json],
};

const LogPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.Log.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const LogPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.Log.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const IpfsServerPayload$json = const {
  '1': 'IpfsServerPayload',
  '3': const [IpfsServerPayload_SrcId$json, IpfsServerPayload_GetResp$json, IpfsServerPayload_View$json],
};

const IpfsServerPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const IpfsServerPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IpfsServer', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const IpfsServerPayload_View$json = const {
  '1': 'View',
  '3': const [IpfsServerPayload_View_SrcResp$json, IpfsServerPayload_View_AddResp$json],
};

const IpfsServerPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.IpfsServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const IpfsServerPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IpfsServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const FileServerPayload$json = const {
  '1': 'FileServerPayload',
  '3': const [FileServerPayload_SrcId$json, FileServerPayload_SrcIds$json, FileServerPayload_GetResp$json, FileServerPayload_View$json],
};

const FileServerPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const FileServerPayload_SrcIds$json = const {
  '1': 'SrcIds',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 13, '10': 'ids'},
  ],
};

const FileServerPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.FileServer', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const FileServerPayload_View$json = const {
  '1': 'View',
  '3': const [FileServerPayload_View_SrcResp$json, FileServerPayload_View_AddResp$json],
};

const FileServerPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.FileServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const FileServerPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.FileServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const HttpProxyServerPayload$json = const {
  '1': 'HttpProxyServerPayload',
  '3': const [HttpProxyServerPayload_SrcId$json, HttpProxyServerPayload_GetResp$json, HttpProxyServerPayload_View$json],
};

const HttpProxyServerPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const HttpProxyServerPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.HttpProxyServer', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const HttpProxyServerPayload_View$json = const {
  '1': 'View',
  '3': const [HttpProxyServerPayload_View_SrcResp$json, HttpProxyServerPayload_View_AddResp$json],
};

const HttpProxyServerPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.HttpProxyServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const HttpProxyServerPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.HttpProxyServer.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const ServerViewPayload$json = const {
  '1': 'ServerViewPayload',
  '3': const [ServerViewPayload_SrcId$json, ServerViewPayload_SrcResp$json, ServerViewPayload_Group$json, ServerViewPayload_GroupResp$json, ServerViewPayload_SrcElement$json, ServerViewPayload_GetResp$json, ServerViewPayload_AddResp$json],
};

const ServerViewPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'ipfs', '3': 1, '4': 1, '5': 13, '10': 'ipfs'},
    const {'1': 'file', '3': 2, '4': 1, '5': 13, '10': 'file'},
    const {'1': 'http', '3': 3, '4': 1, '5': 13, '10': 'http'},
  ],
};

const ServerViewPayload_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.ServerView', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const ServerViewPayload_Group$json = const {
  '1': 'Group',
  '2': const [
    const {'1': 'ipfs', '3': 1, '4': 3, '5': 11, '6': '.hybrid.IpfsServer.View', '10': 'ipfs'},
    const {'1': 'file', '3': 2, '4': 3, '5': 11, '6': '.hybrid.FileServer.View', '10': 'file'},
    const {'1': 'http', '3': 3, '4': 3, '5': 11, '6': '.hybrid.HttpProxyServer.View', '10': 'http'},
  ],
};

const ServerViewPayload_GroupResp$json = const {
  '1': 'GroupResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.ServerViewPayload.Group', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const ServerViewPayload_SrcElement$json = const {
  '1': 'SrcElement',
  '2': const [
    const {'1': 'ipfs', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IpfsServer', '9': 0, '10': 'ipfs'},
    const {'1': 'file', '3': 2, '4': 1, '5': 11, '6': '.hybrid.FileServer', '9': 0, '10': 'file'},
    const {'1': 'http', '3': 3, '4': 1, '5': 11, '6': '.hybrid.HttpProxyServer', '9': 0, '10': 'http'},
  ],
  '8': const [
    const {'1': 'is'},
  ],
};

const ServerViewPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.ServerViewPayload.SrcElement', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const ServerViewPayload_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.ServerView', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const AdpRouterPayload$json = const {
  '1': 'AdpRouterPayload',
  '3': const [AdpRouterPayload_SrcId$json, AdpRouterPayload_GetResp$json, AdpRouterPayload_View$json],
};

const AdpRouterPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const AdpRouterPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.AdpRouter', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const AdpRouterPayload_View$json = const {
  '1': 'View',
  '3': const [AdpRouterPayload_View_SrcResp$json, AdpRouterPayload_View_AddResp$json],
};

const AdpRouterPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.AdpRouter.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const AdpRouterPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.AdpRouter.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const IPNetRouterPayload$json = const {
  '1': 'IPNetRouterPayload',
  '3': const [IPNetRouterPayload_SrcId$json, IPNetRouterPayload_GetResp$json, IPNetRouterPayload_View$json],
};

const IPNetRouterPayload_SrcId$json = const {
  '1': 'SrcId',
  '2': const [
    const {'1': 'id', '3': 1, '4': 1, '5': 13, '10': 'id'},
  ],
};

const IPNetRouterPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IPNetRouter', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const IPNetRouterPayload_View$json = const {
  '1': 'View',
  '3': const [IPNetRouterPayload_View_SrcResp$json, IPNetRouterPayload_View_AddResp$json],
};

const IPNetRouterPayload_View_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.IPNetRouter.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const IPNetRouterPayload_View_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.IPNetRouter.View', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

const RouterUseViewPayload$json = const {
  '1': 'RouterUseViewPayload',
  '3': const [RouterUseViewPayload_SrcResp$json, RouterUseViewPayload_Group$json, RouterUseViewPayload_GroupResp$json, RouterUseViewPayload_SrcElement$json, RouterUseViewPayload_GetResp$json, RouterUseViewPayload_AddResp$json, RouterUseViewPayload_SrcIds$json, RouterUseViewPayload_DstResp$json, RouterUseViewPayload_SelectResp$json],
};

const RouterUseViewPayload_SrcResp$json = const {
  '1': 'SrcResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.RouterUseView', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const RouterUseViewPayload_Group$json = const {
  '1': 'Group',
  '2': const [
    const {'1': 'adp', '3': 1, '4': 3, '5': 11, '6': '.hybrid.AdpRouter.View', '10': 'adp'},
    const {'1': 'ipnet', '3': 2, '4': 3, '5': 11, '6': '.hybrid.IPNetRouter.View', '10': 'ipnet'},
  ],
};

const RouterUseViewPayload_GroupResp$json = const {
  '1': 'GroupResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.RouterUseViewPayload.Group', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const RouterUseViewPayload_SrcElement$json = const {
  '1': 'SrcElement',
  '2': const [
    const {'1': 'adp', '3': 1, '4': 1, '5': 11, '6': '.hybrid.AdpRouter', '9': 0, '10': 'adp'},
    const {'1': 'ipnet', '3': 2, '4': 1, '5': 11, '6': '.hybrid.IPNetRouter', '9': 0, '10': 'ipnet'},
  ],
  '8': const [
    const {'1': 'is'},
  ],
};

const RouterUseViewPayload_GetResp$json = const {
  '1': 'GetResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.RouterUseViewPayload.SrcElement', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const RouterUseViewPayload_AddResp$json = const {
  '1': 'AddResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.RouterUseView', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const RouterUseViewPayload_SrcIds$json = const {
  '1': 'SrcIds',
  '2': const [
    const {'1': 'ids', '3': 1, '4': 3, '5': 11, '6': '.hybrid.RouterUseView.Id', '10': 'ids'},
  ],
};

const RouterUseViewPayload_DstResp$json = const {
  '1': 'DstResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 3, '5': 11, '6': '.hybrid.RouterUseView.Element', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.OperateError', '10': 'error'},
  ],
};

const RouterUseViewPayload_SelectResp$json = const {
  '1': 'SelectResp',
  '2': const [
    const {'1': 'data', '3': 1, '4': 1, '5': 11, '6': '.hybrid.RouterUseView.Element', '10': 'data'},
    const {'1': 'error', '3': 2, '4': 1, '5': 11, '6': '.pgde.error.SubmitError', '10': 'error'},
  ],
};

