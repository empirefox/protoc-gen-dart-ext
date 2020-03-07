import 'dart:collection' show HashMap;
import 'package:intl/intl.dart' show NumberFormat;

final HashMap<String, NumberFormat> _formats = HashMap();

NumberFormat getCurrencyFormat(String code) {
  if (!_formats.containsKey(code)) {
    _formats[code] = NumberFormat.compactCurrency(name: code);
  }
  return _formats[code];
}

final HashMap<String, NumberFormat> _simpleFormats = HashMap();

NumberFormat getSimpleCurrencyFormat(String code) {
  if (!_simpleFormats.containsKey(code)) {
    _simpleFormats[code] = NumberFormat.compactSimpleCurrency(name: code);
  }
  return _simpleFormats[code];
}

final HashMap<String, NumberFormat> _numberFormats = HashMap();

NumberFormat getCurrencyNumberFormat(String code) {
  if (!_numberFormats.containsKey(code)) {
    _numberFormats[code] = NumberFormat.compactCurrency(name: code, symbol: '');
  }
  return _numberFormats[code];
}
