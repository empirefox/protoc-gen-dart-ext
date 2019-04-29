import 'package:collection/collection.dart';

import './generated_validator.dart';

const maxUint8 = 1 << 8 - 1;

class Bytes {
  static final bool Function(List<int> list1, List<int> list2) equal =
      const ListEquality<int>().equals;

  static bool hasPrefix(List<int> s, List<int> prefix) {
    if (isEmpty(prefix)) return true;
    if (isEmpty(s)) return false;
    return s.length >= prefix.length &&
        equal(s.sublist(0, prefix.length), prefix);
  }

  static bool hasSuffix(List<int> s, List<int> suffix) {
    if (isEmpty(suffix)) return true;
    if (isEmpty(s)) return false;
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

  static int len(List<int> s) => s == null ? 0 : s.length;
  static bool isEmpty(List<int> s) => s == null || s.isEmpty;
  static bool isNotEmpty(List<int> s) => s != null && s.isNotEmpty;
}
