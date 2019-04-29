import "package:test/test.dart";
import '../lib/plural/finvtw.dart';

void _testVars(num value, int fExpected, int iExpected, double nExpected,
    int vExpected, int tExpected, int wExpected) {
  test("finvtw callback with right symbol values", () {
    finvtw(value, (int f, int i, double n, int v, int t, int w) {
      expect(f, equals(fExpected), reason: '$value f');
      expect(i, equals(iExpected), reason: '$value i');
      expect(n, equals(nExpected), reason: '$value n');
      expect(v, equals(vExpected), reason: '$value v');
      expect(t, equals(tExpected), reason: '$value t');
      expect(w, equals(wExpected), reason: '$value w');
    });
  });
}

void main() {
  _testVars(-1, 0, 1, 1.0, 0, 0, 0);
  _testVars(1, 0, 1, 1.0, 0, 0, 0);
  _testVars(1.0, 0, 1, 1.0, 1, 0, 0);
  _testVars(10.20, 2, 10, 10.2, 1, 2, 1);
  _testVars(-123, 0, 123, 123.0, 0, 0, 0);
  _testVars(-123.990, 99, 123, 123.99, 2, 99, 2);
  _testVars(0.8, 8, 0, 0.8, 1, 8, 1);
  _testVars(123456.305, 305, 123456, 123456.305, 3, 305, 3);
  _testVars(123456.3057892, 3057892, 123456, 123456.3057892, 7, 3057892, 7);
  _testVars(123456.3057000, 3057, 123456, 123456.3057, 4, 3057, 4);
  _testVars(1000000000000, 0, 1000000000000, 1000000000000, 0, 0, 0);
  _testVars(0.33333, 33333, 0, 0.33333, 5, 33333, 5);
}
