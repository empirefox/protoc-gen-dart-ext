import './units.l10n.dart';

import './atom.dart';
import './prefix.dart';
import './currency.dart';

class Unit {
  final String Function(dynamic) format;
  final List<Cell> dots;
  final List<Cell> per;
  const Unit({this.format, this.dots, this.per});

  String get symbol => l10n(null, false);

  String l10n(UnitsLocalization l, bool p) =>
      _join(dots, l, p) + (per == null ? '' : ('/' + _join(per, l, p)));

  String _join(List<Cell> cells, UnitsLocalization l, bool p) =>
      cells == null ? '' : cells.map((c) => c.l10n(l, p)).join('.');
}

class CurrencyUnit extends CellUnit {
  final CurrencyV1 _c;
  CurrencyUnit(this._c);
  @override
  String from([UnitsLocalization l, bool p]) => _c.l10n(l);
}

class Cell {
  final String exponent;
  final CellPrefix prefix;
  final CellUnit unit;
  const Cell({this.exponent = '', this.prefix, this.unit});
  String l10n(UnitsLocalization l, bool p) =>
      '${prefix.from(l)}${unit.from(l, p)}$exponent';
}

class CellPrefix {
  static const _symbol = const _PrefixSymbolValuer();
  static const _name = const _PrefixNameValuer();

  final PrefixV1 _p;
  final _PrefixValuer _v;

  const CellPrefix.symbol(this._p) : _v = _symbol;
  const CellPrefix.name(this._p) : _v = _name;

  String from([UnitsLocalization l]) => _v.from(_p, l);
}

abstract class CellUnit {
  const CellUnit();
  String from([UnitsLocalization l, bool p]);
}

// AtomUnit
class AtomUnit extends CellUnit {
  static const _symbol = const _AtomSymbolValuer();
  static const _singular = const _AtomSingularValuer();
  static const _plural = const _AtomPluralValuer();
  static const _parse = const _AtomParseValuer();

  final AtomV1 _c;
  final _AtomValuer _v;

  const AtomUnit.symbol(this._c) : _v = _symbol;
  const AtomUnit.singular(this._c) : _v = _singular;
  const AtomUnit.plural(this._c) : _v = _plural;
  const AtomUnit.parse(this._c) : _v = _parse;

  @override
  String from([UnitsLocalization l, bool p]) => _v.from(_c, l, p);
}

// SymbolUnit
class SymbolUnit extends CellUnit {
  final String _s;
  const SymbolUnit(this._s);
  @override
  String from([UnitsLocalization l, bool p]) => _s;
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
  String from(AtomV1 a, [UnitsLocalization l, bool p]);
}

class _AtomSymbolValuer implements _AtomValuer {
  const _AtomSymbolValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, bool p]) => a.symbol;
}

class _AtomSingularValuer implements _AtomValuer {
  const _AtomSingularValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, bool p]) => a.l10n(l, false);
}

class _AtomPluralValuer implements _AtomValuer {
  const _AtomPluralValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, bool p]) => a.l10n(l, true);
}

class _AtomParseValuer implements _AtomValuer {
  const _AtomParseValuer();
  @override
  String from(AtomV1 a, [UnitsLocalization l, bool p]) => a.l10n(l, p);
}
