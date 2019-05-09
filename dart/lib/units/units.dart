library units;

import '../plural/plural.dart';

import './atom.dart';
import './currency.dart';
import './number_format.dart';
import './prefix.dart';
import './units.l10n.dart';

export './atom.dart';
export './currency_format.dart';
export './currency.dart';
export './number_format.dart';
export './prefix.dart';
export './units.l10n.dart';

class Unit {
  static const prefixSymbol = const _PrefixSymbolValuer();
  static const prefixName = const _PrefixNameValuer();

  static const atomSymbol = const _AtomSymbolValuer();
  static const atomSingular = const _AtomOneValuer();
  static const atomPlural = const _AtomOtherValuer();
  static const atomParse = const _AtomParseValuer();

  final CurrencyUnit currency;
  final int decimalLeft;
  final _PrefixValuer _pv;
  final _AtomValuer _av;
  final String fmtLocale;
  final NumberFormatterGetter numberFmtGetter;
  final bool ordinal;
  final List<Cell> _dots;
  final List<Cell> _per;

  const Unit(
      this.currency,
      this.decimalLeft,
      this._pv,
      this._av,
      this.fmtLocale,
      this.numberFmtGetter,
      this.ordinal,
      this._dots,
      this._per);

  String get symbol => l10n(null, Form.one);

  String l10n(UnitsLocalization l, Form p) =>
      _join(_dots, l, p) + (_per == null ? '' : ('/' + _join(_per, l, p)));

  String _join(List<Cell> cells, UnitsLocalization l, Form p) =>
      cells == null ? '' : cells.map((c) => c.l10n(_pv, _av, l, p)).join('.');
}

class CurrencyUnit {
  static const _code = const _CurrencyCodeFormatter();
  static const _symbol = const _CurrencySymbolFormatter();
  static const _name = const _CurrencyNameFormatter();

  final CurrencyV1 _c;
  final _CurrencyFormatter _fmt;

  const CurrencyUnit.code(this._c) : _fmt = _code;
  const CurrencyUnit.symbol(this._c) : _fmt = _symbol;
  const CurrencyUnit.name(this._c) : _fmt = _name;

  String format(v, [UnitsLocalization l]) => _fmt.format(_c, v, l);
}

class Cell {
  final String exponent;
  final PrefixV1 prefix;
  final AtomV1 atom;
  const Cell({this.exponent = '', this.prefix, this.atom});
  String l10n(_PrefixValuer pv, _AtomValuer av, UnitsLocalization l, Form p) =>
      '${pv.from(prefix, l)}${av.from(atom, l, p)}$exponent';
}

// _CurrencyFormatter
abstract class _CurrencyFormatter {
  String format(CurrencyV1 c, v, [UnitsLocalization l]);
}

class _CurrencyCodeFormatter implements _CurrencyFormatter {
  const _CurrencyCodeFormatter();
  @override
  String format(CurrencyV1 c, v, [UnitsLocalization l]) => c.format(v);
}

class _CurrencySymbolFormatter implements _CurrencyFormatter {
  const _CurrencySymbolFormatter();
  @override
  String format(CurrencyV1 c, v, [UnitsLocalization l]) => c.formatSimple(v);
}

class _CurrencyNameFormatter implements _CurrencyFormatter {
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
