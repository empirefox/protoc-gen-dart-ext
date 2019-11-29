// DO NOT EDIT. Generated by protoc-gen-dart-ext/tools.

/// Code generated by protoc-gen-dart-ext. DO NOT EDIT.
/// source: hybrid/config.proto
// ignore_for_file: non_constant_identifier_names,unnecessary_brace_in_string_interps,unused_local_variable

import 'config.pb.dart' as $0
    show
        AdpRouter,
        Basic,
        Config,
        FileServer,
        HttpProxyServer,
        IPNetRouter,
        Ipfs,
        IpfsServer,
        Log,
        RouterItem,
        RouterItem_Router;
import 'package:pgde/pgde.dart' as $1
    show
        BeSomethingError,
        Bytes,
        ConstError,
        GeneratedValidator,
        InError,
        LenConstError,
        Lists,
        OneofRequiredError,
        RequiredError,
        ValidateInfo;
import 'package:flutter/material.dart' as $2 show BuildContext, Localizations;
import 'config.l10n.dart' as $3 show ConfigLocalizations;
import 'dart:collection' as $4 show HashSet;

/// Validates [$0.Basic] protobuf objects.
class BasicValidator implements $1.GeneratedValidator<$0.Basic> {
  @override
  void assertProto() {
    assertField_version();
    assertField_dev();
    assertField_enableBind();
    assertField_bindIp();
    assertField_bindPort();
    assertField_flushIntervalMs();
    assertField_token();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_version();
      case 2:
        return assertField_dev();
      case 3:
        return assertField_enableBind();
      case 4:
        return assertField_bindIp();
      case 5:
        return assertField_bindPort();
      case 6:
        return assertField_flushIntervalMs();
      case 7:
        return assertField_token();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_version() {
    // no validation rules for version
  }
  void assertField_dev() {
    // no validation rules for dev
  }
  void assertField_enableBind() {
    // no validation rules for enableBind
  }
  void assertField_bindIp() {
    final _v = _info.proto.bindIp;

    if (!$1.Bytes.isIP(_v))
      throw $1.BeSomethingError(
          _info, 4, _l10n.BasicBindIp, _info.l10n.validateIp);
  }

  void assertField_bindPort() {
    final _v = _info.proto.bindPort;

    if (_v >= 65535)
      throw $1.ConstError(
          _info, 5, _l10n.BasicBindPort, _info.l10n.validateLt, 65535);
  }

  void assertField_flushIntervalMs() {
    // no validation rules for flushIntervalMs
  }
  void assertField_token() {
    final _v = _info.proto.token;

    final _vrl = _v.runes.length;

    if (_vrl > 732)
      throw $1.LenConstError.character(
          _info, 7, _l10n.BasicToken, _info.l10n.validateLte, 732);
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.Basic> _info;

  $3.ConfigLocalizations _l10n;

  BasicValidator($2.BuildContext context, $1.ValidateInfo<$0.Basic> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.Log] protobuf objects.
class LogValidator implements $1.GeneratedValidator<$0.Log> {
  static const Map<String, Null> _levelIn = {
    r'': null,
    r'debug': null,
    r'info': null,
    r'warn': null,
    r'error': null,
    r'dpanic': null,
    r'panic': null,
    r'fatal': null,
  };

  @override
  void assertProto() {
    assertField_dev();
    assertField_level();
    assertField_target();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_dev();
      case 2:
        return assertField_level();
      case 3:
        return assertField_target();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_dev() {
    // no validation rules for dev
  }
  void assertField_level() {
    final _v = _info.proto.level;

    if (!_levelIn.containsKey(_v))
      throw $1.InError(_info, 2, _l10n.LogLevel, _levelIn.keys.toList());
  }

  void assertField_target() {
    // no validation rules for target
  }
  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.Log> _info;

  $3.ConfigLocalizations _l10n;

  LogValidator($2.BuildContext context, $1.ValidateInfo<$0.Log> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.Ipfs] protobuf objects.
class IpfsValidator implements $1.GeneratedValidator<$0.Ipfs> {
  @override
  void assertProto() {
    assertField_fakeApiListenIp();
    assertField_fakeApiListenPort();
    assertField_enableGateway();
    assertField_gatewayServerName();
    assertField_enableApi();
    assertField_apiServerName();
    assertField_profile();
    assertField_autoMigrate();
    assertField_enableIpnsPubSub();
    assertField_enablePubSub();
    assertField_enableMultiplex();
    assertField_token();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_fakeApiListenIp();
      case 2:
        return assertField_fakeApiListenPort();
      case 3:
        return assertField_enableGateway();
      case 4:
        return assertField_gatewayServerName();
      case 5:
        return assertField_enableApi();
      case 6:
        return assertField_apiServerName();
      case 7:
        return assertField_profile();
      case 8:
        return assertField_autoMigrate();
      case 9:
        return assertField_enableIpnsPubSub();
      case 10:
        return assertField_enablePubSub();
      case 11:
        return assertField_enableMultiplex();
      case 12:
        return assertField_token();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_fakeApiListenIp() {
    final _v = _info.proto.fakeApiListenIp;

    if (!$1.Bytes.isIP(_v))
      throw $1.BeSomethingError(
          _info, 1, _l10n.IpfsFakeApiListenIp, _info.l10n.validateIp);
  }

  void assertField_fakeApiListenPort() {
    final _v = _info.proto.fakeApiListenPort;

    if (_v >= 65535)
      throw $1.ConstError(
          _info, 2, _l10n.IpfsFakeApiListenPort, _info.l10n.validateLt, 65535);
  }

  void assertField_enableGateway() {
    // no validation rules for enableGateway
  }
  void assertField_gatewayServerName() {
    final _v = _info.proto.gatewayServerName;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 4, _l10n.IpfsGatewayServerName, _info.l10n.validateHostname);
    }
  }

  void assertField_enableApi() {
    // no validation rules for enableApi
  }
  void assertField_apiServerName() {
    final _v = _info.proto.apiServerName;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 6, _l10n.IpfsApiServerName, _info.l10n.validateHostname);
    }
  }

  void assertField_profile() {
    final _v = _info.proto.profile;

    final Set<String> _unique = $4.HashSet();

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      if (!_unique.add(_ritem))
        throw $1.BeSomethingError(
            _info, 7, _l10n.IpfsProfile, _info.l10n.validateUnique(_ridx + 1));

      // no validation rules for profile

    }
  }

  void assertField_autoMigrate() {
    // no validation rules for autoMigrate
  }
  void assertField_enableIpnsPubSub() {
    // no validation rules for enableIpnsPubSub
  }
  void assertField_enablePubSub() {
    // no validation rules for enablePubSub
  }
  void assertField_enableMultiplex() {
    // no validation rules for enableMultiplex
  }
  void assertField_token() {
    final _v = _info.proto.token;

    final _vrl = _v.runes.length;

    if (_vrl > 732)
      throw $1.LenConstError.character(
          _info, 12, _l10n.IpfsToken, _info.l10n.validateLte, 732);
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.Ipfs> _info;

  $3.ConfigLocalizations _l10n;

  IpfsValidator($2.BuildContext context, $1.ValidateInfo<$0.Ipfs> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.IpfsServer] protobuf objects.
class IpfsServerValidator implements $1.GeneratedValidator<$0.IpfsServer> {
  @override
  void assertProto() {
    assertField_disabled();
    assertField_name();
    assertField_peer();
    assertField_token();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_disabled();
      case 2:
        return assertField_name();
      case 3:
        return assertField_peer();
      case 4:
        return assertField_token();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_disabled() {
    // no validation rules for disabled
  }
  void assertField_name() {
    final _v = _info.proto.name;

    if (_info.proto.hasName()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 2, _l10n.IpfsServerName, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_peer() {
    final _v = _info.proto.peer;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 3, _l10n.IpfsServerPeer, _info.l10n.validateHostname);
    }
  }

  void assertField_token() {
    final _v = _info.proto.token;

    final _vrl = _v.runes.length;

    if (_vrl > 732)
      throw $1.LenConstError.character(
          _info, 4, _l10n.IpfsServerToken, _info.l10n.validateLte, 732);
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.IpfsServer> _info;

  $3.ConfigLocalizations _l10n;

  IpfsServerValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.IpfsServer> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.FileServer] protobuf objects.
class FileServerValidator implements $1.GeneratedValidator<$0.FileServer> {
  @override
  void assertProto() {
    assertField_disabled();
    assertField_name();
    assertField_zip();
    assertField_redirect();
    assertField_dev();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_disabled();
      case 2:
        return assertField_name();
      case 3:
        return assertField_zip();
      case 4:
        return assertField_redirect();
      case 5:
        return assertField_dev();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_disabled() {
    // no validation rules for disabled
  }
  void assertField_name() {
    final _v = _info.proto.name;

    if (_info.proto.hasName()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 2, _l10n.FileServerName, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_zip() {
    final _v = _info.proto.zip;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 3, _l10n.FileServerZip, _info.l10n.validateHostname);
    }
  }

  void assertField_redirect() {
    final _v = _info.proto.redirect;

    _v.forEach((_mk, _mv) {
      try {
        Uri.parse(_mk);
      } on FormatException {
        throw $1.BeSomethingError(_info, 4,
            '`${_l10n.FileServerRedirect}[${_mk}]`', _info.l10n.validateUri);
      }

      try {
        Uri.parse(_mv);
      } on FormatException {
        throw $1.BeSomethingError(_info, 4,
            '`${_l10n.FileServerRedirect}[${_mk}]`', _info.l10n.validateUri);
      }
    });
  }

  void assertField_dev() {
    // no validation rules for dev
  }
  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.FileServer> _info;

  $3.ConfigLocalizations _l10n;

  FileServerValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.FileServer> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.HttpProxyServer] protobuf objects.
class HttpProxyServerValidator
    implements $1.GeneratedValidator<$0.HttpProxyServer> {
  @override
  void assertProto() {
    assertField_disabled();
    assertField_name();
    assertField_host();
    assertField_port();
    assertField_keepAlive();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_disabled();
      case 2:
        return assertField_name();
      case 3:
        return assertField_host();
      case 4:
        return assertField_port();
      case 5:
        return assertField_keepAlive();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_disabled() {
    // no validation rules for disabled
  }
  void assertField_name() {
    final _v = _info.proto.name;

    if (_info.proto.hasName()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 2, _l10n.HttpProxyServerName, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_host() {
    final _v = _info.proto.host;

    try {
      Uri.parseIPv4Address(_v);
    } on FormatException {
      try {
        Uri.parseIPv6Address(_v);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 3, _l10n.HttpProxyServerHost, _info.l10n.validateIp);
      }
    }
  }

  void assertField_port() {
    final _v = _info.proto.port;

    if (_v >= 65535)
      throw $1.ConstError(
          _info, 4, _l10n.HttpProxyServerPort, _info.l10n.validateLt, 65535);
  }

  void assertField_keepAlive() {
    // no validation rules for keepAlive
  }
  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.HttpProxyServer> _info;

  $3.ConfigLocalizations _l10n;

  HttpProxyServerValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.HttpProxyServer> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.AdpRouter] protobuf objects.
class AdpRouterValidator implements $1.GeneratedValidator<$0.AdpRouter> {
  @override
  void assertProto() {
    assertField_rulesDirName();
    assertField_blocked();
    assertField_unblocked();
    assertField_etcHostsIpAsBlocked();
    assertField_dev();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_rulesDirName();
      case 2:
        return assertField_blocked();
      case 3:
        return assertField_unblocked();
      case 4:
        return assertField_etcHostsIpAsBlocked();
      case 5:
        return assertField_dev();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_rulesDirName() {
    final _v = _info.proto.rulesDirName;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 1, _l10n.AdpRouterRulesDirName, _info.l10n.validateHostname);
    }
  }

  void assertField_blocked() {
    final _v = _info.proto.blocked;

    if (_info.proto.hasBlocked()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 2, _l10n.AdpRouterBlocked, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_unblocked() {
    final _v = _info.proto.unblocked;

    if (_info.proto.hasUnblocked()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 3, _l10n.AdpRouterUnblocked, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_etcHostsIpAsBlocked() {
    // no validation rules for etcHostsIpAsBlocked
  }
  void assertField_dev() {
    // no validation rules for dev
  }
  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.AdpRouter> _info;

  $3.ConfigLocalizations _l10n;

  AdpRouterValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.AdpRouter> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.IPNetRouter] protobuf objects.
class IPNetRouterValidator implements $1.GeneratedValidator<$0.IPNetRouter> {
  @override
  void assertProto() {
    assertField_ip();
    assertField_netv4();
    assertField_netv6();
    assertField_matched();
    assertField_unmatched();
    assertField_fileTest();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_ip();
      case 2:
        return assertField_netv4();
      case 3:
        return assertField_netv6();
      case 4:
        return assertField_matched();
      case 5:
        return assertField_unmatched();
      case 6:
        return assertField_fileTest();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_ip() {
    final _v = _info.proto.ip;

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      if (!$1.Bytes.isIP(_ritem.value))
        throw $1.BeSomethingError(
            _info, 1, _l10n.IPNetRouterIp, _info.l10n.validateIp);
    }
  }

  void assertField_netv4() {
    final _v = _info.proto.netv4;

    final _bl = $1.Lists.len(_v);

    if (_bl != 5)
      throw $1.LenConstError.byte(
          _info, 2, _l10n.IPNetRouterNetv4, _info.l10n.validateEq, 5);
  }

  void assertField_netv6() {
    final _v = _info.proto.netv6;

    final _bl = $1.Lists.len(_v);

    if (_bl != 17)
      throw $1.LenConstError.byte(
          _info, 3, _l10n.IPNetRouterNetv6, _info.l10n.validateEq, 17);
  }

  void assertField_matched() {
    final _v = _info.proto.matched;

    if (_info.proto.hasMatched()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 4, _l10n.IPNetRouterMatched, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_unmatched() {
    final _v = _info.proto.unmatched;

    if (_info.proto.hasUnmatched()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 5, _l10n.IPNetRouterUnmatched, _info.l10n.validateHostname);
      }
    }
  }

  void assertField_fileTest() {
    final _v = _info.proto.fileTest;

    if (_info.proto.hasFileTest()) {
      try {
        Uri(host: _v.value);
      } on FormatException {
        throw $1.BeSomethingError(
            _info, 6, _l10n.IPNetRouterFileTest, _info.l10n.validateHostname);
      }
    }
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.IPNetRouter> _info;

  $3.ConfigLocalizations _l10n;

  IPNetRouterValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.IPNetRouter> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.RouterItem] protobuf objects.
class RouterItemValidator implements $1.GeneratedValidator<$0.RouterItem> {
  @override
  void assertProto() {
    assertField_disabled();
    assertField_name();

    assertOneof_Router();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_disabled();
      case 2:
        return assertField_name();
      case 3:
        return assertField_adp();
      case 4:
        return assertField_ipnet();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_disabled() {
    // no validation rules for disabled
  }
  void assertField_name() {
    final _v = _info.proto.name;

    try {
      Uri(host: _v);
    } on FormatException {
      throw $1.BeSomethingError(
          _info, 2, _l10n.RouterItemName, _info.l10n.validateHostname);
    }
  }

  void assertField_adp() {
    final _v = _info.proto.adp;

    AdpRouterValidator(_ctx, _info.clone(_v)).assertProto();
  }

  void assertField_ipnet() {
    final _v = _info.proto.ipnet;

    IPNetRouterValidator(_ctx, _info.clone(_v)).assertProto();
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      case $0.RouterItem_Router:
        return assertOneof_Router();

      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  void assertOneof_Router() {
    switch (_info.proto.whichRouter()) {
      case $0.RouterItem_Router.adp:
        assertField_adp();
        break;
      case $0.RouterItem_Router.ipnet:
        assertField_ipnet();
        break;

      default:
        throw $1.OneofRequiredError(_info, _l10n.RouterItemRouter);
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.RouterItem> _info;

  $3.ConfigLocalizations _l10n;

  RouterItemValidator(
      $2.BuildContext context, $1.ValidateInfo<$0.RouterItem> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}

/// Validates [$0.Config] protobuf objects.
class ConfigValidator implements $1.GeneratedValidator<$0.Config> {
  @override
  void assertProto() {
    assertField_basic();
    assertField_log();
    assertField_ipfs();
    assertField_ipfsServers();
    assertField_fileServers();
    assertField_httpProxyServers();
    assertField_routers();
  }

  @override
  void assertField(int tag) {
    switch (tag) {
      case 1:
        return assertField_basic();
      case 2:
        return assertField_log();
      case 3:
        return assertField_ipfs();
      case 4:
        return assertField_ipfsServers();
      case 5:
        return assertField_fileServers();
      case 6:
        return assertField_httpProxyServers();
      case 7:
        return assertField_routers();

      default:
        assert(false, 'tag number($tag) for ${_info.bi.messageName} not found');
    }
  }

  void assertField_basic() {
    final _v = _info.proto.basic;

    if (_v == null) throw $1.RequiredError(_info, 1, _l10n.ConfigBasic);

    BasicValidator(_ctx, _info.clone(_v)).assertProto();
  }

  void assertField_log() {
    final _v = _info.proto.log;

    LogValidator(_ctx, _info.clone(_v)).assertProto();
  }

  void assertField_ipfs() {
    final _v = _info.proto.ipfs;

    IpfsValidator(_ctx, _info.clone(_v)).assertProto();
  }

  void assertField_ipfsServers() {
    final _v = _info.proto.ipfsServers;

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      IpfsServerValidator(_ctx, _info.clone(_ritem)).assertProto();
    }
  }

  void assertField_fileServers() {
    final _v = _info.proto.fileServers;

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      FileServerValidator(_ctx, _info.clone(_ritem)).assertProto();
    }
  }

  void assertField_httpProxyServers() {
    final _v = _info.proto.httpProxyServers;

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      HttpProxyServerValidator(_ctx, _info.clone(_ritem)).assertProto();
    }
  }

  void assertField_routers() {
    final _v = _info.proto.routers;

    for (var _ridx = 0; _ridx < _v.length; _ridx++) {
      final _ritem = _v[_ridx];

      RouterItemValidator(_ctx, _info.clone(_ritem)).assertProto();
    }
  }

  @override
  void assertOneof(Type oneof) {
    switch (oneof) {
      default:
        assert(
            false, 'oneof type($oneof) for ${_info.bi.messageName} not found');
    }
  }

  $2.BuildContext _ctx;

  $1.ValidateInfo<$0.Config> _info;

  $3.ConfigLocalizations _l10n;

  ConfigValidator($2.BuildContext context, $1.ValidateInfo<$0.Config> info)
      : _ctx = context,
        _info = info,
        _l10n = $2.Localizations.of<$3.ConfigLocalizations>(
            context, $3.ConfigLocalizations);
}
