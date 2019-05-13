/// @see http://unicode.org/reports/tr35/tr35-numbers.html#Operands
///
/// | Symbol | Value |
/// | ------ | ----- |
/// | [n] | absolute value of the source number (integer and decimals). |
/// | [i] | integer digits of n. |
/// | [v] | number of visible fraction digits in n, with trailing zeros. |
/// | [w] | number of visible fraction digits in n, without trailing zeros. |
/// | [f] | visible fractional digits in n, with trailing zeros. |
/// | [t] | visible fractional digits in n, without trailing zeros. |
/// 
/// For dart:
///   * double(1) double(1.00) will always be seen as 1.0
///   * double(0.100) will always be seen as 0.1
typedef FinvtwCallback<T> = T Function(
    int f, int i, double n, int v, int t, int w);

T finvtw<T>(num value, FinvtwCallback<T> cb) {
  if (value is int) {
    final i = value.abs();
    return cb(0, i, i.toDouble(), 0, 0, 0);
  }

  final n = value.abs();
  final i = n.toInt();
  final ns = n.toString();
  final il = i.toString().length;
  final v = ns.length - il - 1;
  if (v == -1) {
    return cb(0, i, n, 0, 0, 0);
  }

  final f = int.parse(ns.substring(il + 1));
  final fs = f.toString();
  final w = fs.codeUnits.lastIndexWhere((c) => c != 48) + 1;
  final t = w == 0 ? f : int.parse(fs.substring(0, w));
  return cb(f, i, n, v, t, w);
}
