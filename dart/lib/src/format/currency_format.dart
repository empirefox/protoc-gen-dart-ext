import '../l10n/pgde.l10n.dart';

import 'currency.dart';
import 'formatter.dart';

class CurrencyFormatter implements Formatter<num> {
  static const _code = const _CurrencyCodeFormatter();
  static const _symbol = const _CurrencySymbolFormatter();
  static const _name = const _CurrencyNameFormatter();

  final Currency _c;
  final _CurrencyFormatterter _fmt;
  final bool _centMode;

  const CurrencyFormatter.code(this._c, [this._centMode = false])
      : assert(_c != null),
        _fmt = _code;

  const CurrencyFormatter.symbol(this._c, [this._centMode = false])
      : assert(_c != null),
        _fmt = _symbol;

  const CurrencyFormatter.name(this._c, [this._centMode = false])
      : assert(_c != null),
        _fmt = _name;

  bool get convert => _centMode && _c.mnrFactor != null && _c.mnrFactor != 1;

  double toDolar(num v) => convert ? v / _c.mnrFactor : v.toDouble();

  // make sure called when centMode
  int toCent(double v) => (convert ? v * _c.mnrFactor : v).toInt();

  String format(v, md, l, {form, show}) => _fmt.format(_c, toDolar(v), l);
}

// _CurrencyFormatterter
abstract class _CurrencyFormatterter {
  String format(Currency c, v, [PgdeLocalizations l]);
}

class _CurrencyCodeFormatter implements _CurrencyFormatterter {
  const _CurrencyCodeFormatter();
  @override
  String format(Currency c, v, [PgdeLocalizations l]) => c.format(v);
}

class _CurrencySymbolFormatter implements _CurrencyFormatterter {
  const _CurrencySymbolFormatter();
  @override
  String format(Currency c, v, [PgdeLocalizations l]) => c.formatSimple(v);
}

class _CurrencyNameFormatter implements _CurrencyFormatterter {
  const _CurrencyNameFormatter();
  @override
  String format(Currency c, v, [PgdeLocalizations l]) => c.formatName(v, l);
}
