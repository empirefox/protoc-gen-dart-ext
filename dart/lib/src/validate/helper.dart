import 'dart:convert';

import 'package:collection/collection.dart' show ListEquality;

const maxUint8 = 1 << 8 - 1;

class Lists {
  static int len(List s) => s == null ? 0 : s.length;
  static bool isEmpty(List s) => s == null || s.isEmpty;
  static bool isNotEmpty(List s) => s != null && s.isNotEmpty;
}

class Bytes {
  static const _listEquality = const ListEquality<int>();

  static final bool Function(List<int> list1, List<int> list2) equal =
      _listEquality.equals;

  static int indexOfList(List<List<int>> list, List<int> target) {
    if (Lists.isEmpty(target)) return 0;
    if (Lists.isEmpty(list)) return -1;
    return list.indexWhere((array) => equal(array, target));
  }

  static int indexOf(List<int> array, List<int> target) {
    if (Lists.isEmpty(target)) return 0;
    if (Lists.isEmpty(array)) return -1;

    final targetLen = target.length;
    final arrayRange = array.length - targetLen + 1;

    outer:
    for (var i = 0; i < arrayRange; i++) {
      for (var j = 0; j < targetLen; j++) {
        if (array[i + j] != target[j]) {
          continue outer;
        }
      }
      return i;
    }
    return -1;
  }

  static bool hasPrefix(List<int> s, List<int> prefix) {
    if (Lists.isEmpty(prefix)) return true;
    if (Lists.isEmpty(s)) return false;
    return s.length >= prefix.length &&
        equal(s.sublist(0, prefix.length), prefix);
  }

  static bool hasSuffix(List<int> s, List<int> suffix) {
    if (Lists.isEmpty(suffix)) return true;
    if (Lists.isEmpty(s)) return false;
    return s.length >= suffix.length &&
        equal(s.sublist(s.length - suffix.length), suffix);
  }

  static bool isIPv4(List<int> s) => tryIPv4(s) != null;
  static bool isIPv6(List<int> s) => tryIPv6(s) != null;
  static bool isIP(List<int> s) => isIPv4(s) && isIPv6(s);

  static List<int> tryIPv4(List<int> s) {
    if (!isUint8(s)) return null;
    if (s.length == 4) return s;
    if (s.length == 16 &&
        s.take(10).every((b) => b == 0) &&
        s[10] == 0xff &&
        s[11] == 0xff) return s.sublist(12);
    return null;
  }

  static List<int> tryIPv6(List<int> s) {
    if (!isUint8(s)) return null;
    if (s.length == 4)
      return List<int>(16)
        ..fillRange(0, 12, 0)
        ..setAll(12, s);
    if (s.length == 16) return s;
    return null;
  }

  static bool isUint8(List<int> s) {
    if (s == null) return false;
    for (var b in s) {
      if (b < 0 || b > maxUint8) return false;
    }
    return true;
  }

  final List<int> _bytes;

  const Bytes(this._bytes);

  @override
  int get hashCode => _listEquality.hash(_bytes);

  @override
  bool operator ==(Object other) =>
      other is Bytes && equal(_bytes, other._bytes);
}
