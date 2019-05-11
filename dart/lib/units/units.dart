library units;

import '../plural/plural.dart';

import './atom.dart';
import './currency.dart';
import './number_format.dart';
import './prefix.dart';
import './time_format.dart';
import './units.l10n.dart';

export './atom.dart';
export './currency_format.dart';
export './currency.dart';
export './number_format.dart';
export './prefix.dart';
export './time_format.dart';
export './units.l10n.dart';

class Unit<F, T> {
  static const symbolPrefix = const _PrefixSymbolValuer();
  static const namePrefix = const _PrefixNameValuer();

  static const symbolAtom = const _AtomSymbolValuer();
  static const singularAtom = const _AtomOneValuer();
  static const pluralAtom = const _AtomOtherValuer();
  static const parseAtom = const _AtomParseValuer();

  // just place them here
  final TimeFormatter<F, T> time;
  final CurrencyFormatter currency;
  final NumberFormatterGetter _numberFmtGetter;

  final _PrefixValuer _pv;
  final _AtomValuer _av;
  final String _fmtLocale;
  final bool _ordinal;
  final List<Cell> _dots;
  final List<Cell> _per;

  const Unit(this.time, this.currency, this._numberFmtGetter, this._pv,
      this._av, this._fmtLocale, this._ordinal, this._dots, this._per);

  String get symbol => l10n(null, Form.one);

  String formatNumber(v) =>
      _numberFmtGetter?.from(_fmtLocale)?.format(v) ?? '$v';

  Form plural(v, [bool ordinal]) =>
      pluralRule(_fmtLocale)(v, ordinal ?? _ordinal ?? false);

  String l10n(UnitsLocalization l, Form p) =>
      _join(_dots, l, p) + (_per == null ? '' : ('/' + _join(_per, l, p)));

  String _join(List<Cell> cells, UnitsLocalization l, Form p) =>
      cells == null ? '' : cells.map((c) => c.l10n(_pv, _av, l, p)).join('.');
}

class CurrencyFormatter {
  static const _code = const _CurrencyCodeFormatter();
  static const _symbol = const _CurrencySymbolFormatter();
  static const _name = const _CurrencyNameFormatter();

  final CurrencyV1 _c;
  final _CurrencyFormatterter _fmt;
  final bool _centMode;

  const CurrencyFormatter.code(this._c, this._centMode) : _fmt = _code;
  const CurrencyFormatter.symbol(this._c, this._centMode) : _fmt = _symbol;
  const CurrencyFormatter.name(this._c, this._centMode) : _fmt = _name;

  bool get convert => _centMode && _c.mnrFactor != null && _c.mnrFactor != 1;

  double toDolar(int v) => convert ? v / _c.mnrFactor : v.toDouble();

  int toCent(double v) => (convert ? v * _c.mnrFactor : v).toInt();

  String format(v, [UnitsLocalization l]) => _fmt.format(_c, toDolar(v), l);
}

class Cell {
  final String exponent;
  final PrefixV1 prefix;
  final AtomV1 atom;
  const Cell({this.exponent = '', this.prefix, this.atom});
  String l10n(_PrefixValuer pv, _AtomValuer av, UnitsLocalization l, Form p) =>
      '${pv.from(prefix, l)}${av.from(atom, l, p)}$exponent';
}

// _CurrencyFormatterter
abstract class _CurrencyFormatterter {
  String format(CurrencyV1 c, v, [UnitsLocalization l]);
}

class _CurrencyCodeFormatter implements _CurrencyFormatterter {
  const _CurrencyCodeFormatter();
  @override
  String format(CurrencyV1 c, v, [UnitsLocalization l]) => c.format(v);
}

class _CurrencySymbolFormatter implements _CurrencyFormatterter {
  const _CurrencySymbolFormatter();
  @override
  String format(CurrencyV1 c, v, [UnitsLocalization l]) => c.formatSimple(v);
}

class _CurrencyNameFormatter implements _CurrencyFormatterter {
  const _CurrencyNameFormatter();
  @override
  String format(CurrencyV1 c, v, [UnitsLocalization l]) => c.formatName(v, l);
}

// _PrefixValuer
abstract class _PrefixValuer {
  String from(PrefixV1 p, [UnitsLocalization l]);
}

class _PrefixSymbolValuer implements _PrefixValuer {
  const _PrefixSymbolValuer();
  @override
  String from(PrefixV1 p, [UnitsLocalization l]) => p.symbol;
}

class _PrefixNameValuer implements _PrefixValuer {
  const _PrefixNameValuer();
  @override
  String from(PrefixV1 p, [UnitsLocalization l]) => p.l10n(l);
}

// _AtomValuer
abstract class _AtomValuer {
  String from(AtomV1 a, [UnitsLocalization l, Form p]);
}

class _AtomSymbolValuer implements _AtomValuer {
  const _AtomSymbolValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, Form p]) => a.symbol;
}

class _AtomOneValuer implements _AtomValuer {
  const _AtomOneValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, Form p]) => a.l10n(l, Form.one);
}

class _AtomOtherValuer implements _AtomValuer {
  const _AtomOtherValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, Form p]) => a.l10n(l, Form.other);
}

class _AtomParseValuer implements _AtomValuer {
  const _AtomParseValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, Form p]) => a.l10n(l, p);
}
