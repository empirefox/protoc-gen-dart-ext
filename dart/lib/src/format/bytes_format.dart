import 'package:flutter/material.dart' show MaterialLocalizations;
import 'package:ip/ip.dart';

import '../l10n/pgde.l10n.dart';
import '../plural/plural.dart';
import '../units/units.dart';

import 'formatter.dart';

abstract class BytesFormatter implements Formatter<List<int>> {
  static const BytesFormatter ipv4 = const _Ipv4BytesFormatter();
  static const BytesFormatter ipv6 = const _Ipv6BytesFormatter();
  static const BytesFormatter ip = const _IpBytesFormatter();

  static const BytesFormatter ipv4Port = const _Ipv4PortBytesFormatter();
  static const BytesFormatter ipv6Port = const _Ipv6PortBytesFormatter();
  static const BytesFormatter ipPort = const _IpPortBytesFormatter();

  static const BytesFormatter cidrv4 = const _Cidrv4BytesFormatter();
  static const BytesFormatter cidrv6 = const _Cidrv6BytesFormatter();
  static const BytesFormatter cidr = const _CidrBytesFormatter();

  const BytesFormatter();

  @override
  String format(List<int> v, MaterialLocalizations md, PgdeLocalizations l,
          {Form form, Show show}) =>
      v == null || v.length == 0 ? '' : formatBytes(v);

  String formatBytes(List<int> v);
}

class _Ipv4BytesFormatter extends BytesFormatter {
  const _Ipv4BytesFormatter();
  @override
  String formatBytes(List<int> v) => v.join('.');
}

class _Ipv6BytesFormatter extends BytesFormatter {
  const _Ipv6BytesFormatter();
  @override
  String formatBytes(List<int> v) => Ip6Address.fromBytes(v, 0).toString();
}

class _IpBytesFormatter extends BytesFormatter {
  const _IpBytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      (v.length == 16 ? BytesFormatter.ipv6 : BytesFormatter.ipv4)
          .formatBytes(v);
}

class _Ipv4PortBytesFormatter extends BytesFormatter {
  const _Ipv4PortBytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      '${BytesFormatter.ipv4.formatBytes(v.sublist(0, 4))}:${_toUint16(v.sublist(4))}';
}

class _Ipv6PortBytesFormatter extends BytesFormatter {
  const _Ipv6PortBytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      '${BytesFormatter.ipv6.formatBytes(v.sublist(0, 16))}:${_toUint16(v.sublist(16))}';
}

class _IpPortBytesFormatter extends BytesFormatter {
  const _IpPortBytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      (v.length == 18 ? BytesFormatter.ipv6Port : BytesFormatter.ipv4Port)
          .formatBytes(v);
}

class _Cidrv4BytesFormatter extends BytesFormatter {
  const _Cidrv4BytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      '${BytesFormatter.ipv4.formatBytes(v.sublist(0, 4))}/${v[4]}';
}

class _Cidrv6BytesFormatter extends BytesFormatter {
  const _Cidrv6BytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      '${BytesFormatter.ipv6.formatBytes(v.sublist(0, 16))}/${v[16]}';
}

class _CidrBytesFormatter extends BytesFormatter {
  const _CidrBytesFormatter();
  @override
  String formatBytes(List<int> v) =>
      (v.length == 17 ? BytesFormatter.cidrv6 : BytesFormatter.cidrv4)
          .formatBytes(v);
}

int _toUint16(List<int> v) => 0xFF & v[1] | (0xFF & v[0]) << 8;
