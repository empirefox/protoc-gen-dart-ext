import 'dart:collection';
import 'package:intl/intl.dart';

final HashMap<String, NumberFormat> _formats = HashMap();

NumberFormat getCurrencyFormat(String code) {
  if (!_formats.containsKey(code)) {
    _formats[code] = NumberFormat.currency(name: code);
  }
  return _formats[code];
}

final HashMap<String, NumberFormat> _simpleFormats = HashMap();

NumberFormat getSimpleCurrencyFormat(String code) {
  if (!_simpleFormats.containsKey(code)) {
    _simpleFormats[code] = NumberFormat.currency(name: code);
  }
  return _simpleFormats[code];
}
