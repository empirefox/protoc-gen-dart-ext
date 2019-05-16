import '../plural/plural.dart';
import '../l10n/pgde.l10n.dart';

import 'atom.dart';
import 'prefix.dart';

export 'atom.dart';
export 'prefix.dart';

class Unit {
  static const showSymbol = Show(false, Show.symbolPrefix, Show.symbol);

  final Show show;
  final List<Cell> _dots;
  final List<Cell> _per;

  const Unit(this.show, this._dots, this._per) : assert(show != null);

  String get symbol => l10n(null, Form.one, showSymbol);

  String l10n(PgdeLocalization l, Form p, [Show show]) {
    show = show ?? this.show;
    if (show.off) return '';
    final buf = StringBuffer();
    _join(buf, null, _dots, l, p, show);
    _join(buf, '/', _per, l, p, show);
    return buf.toString();
  }

  void _join(StringBuffer buf, String prefix, List<Cell> cells,
      PgdeLocalization l, Form p, Show show) {
    if (cells != null) {
      if (prefix != null) buf.write(prefix);
      buf.writeAll(cells.map((c) => c.l10n(show.prefix, show.atom, l, p)), '.');
    }
  }
}

class Show {
  static const disabled = Show._disabled();

  static const symbolPrefix = _PrefixSymbolValuer();
  static const namePrefix = _PrefixNameValuer();

  static const symbol = _AtomSymbolValuer();
  static const one = _AtomOneValuer();
  static const other = _AtomOtherValuer();
  static const parse = _AtomParseValuer();

  final bool off;
  final _PrefixValuer prefix;
  final _AtomValuer atom;

  const Show(this.off, this.prefix, this.atom)
      : assert(off != null),
        assert(prefix != null),
        assert(atom != null);

  const Show._disabled()
      : off = true,
        prefix = null,
        atom = null;

  bool get willParse => !off && atom == parse;
}

class Cell {
  final String exponent;
  final PrefixV1 prefix;
  final AtomV1 atom;

  const Cell({this.exponent = '', this.prefix, this.atom})
      : assert(prefix != null),
        assert(atom != null);

  String l10n(_PrefixValuer pv, _AtomValuer av, PgdeLocalization l, Form p) =>
      '${pv.from(prefix, l)}${av.from(atom, l, p)}$exponent';
}

// _PrefixValuer
abstract class _PrefixValuer {
  String from(PrefixV1 p, [PgdeLocalization l]);
}

class _PrefixSymbolValuer implements _PrefixValuer {
  const _PrefixSymbolValuer();
  @override
  String from(PrefixV1 p, [PgdeLocalization l]) => p.symbol;
}

class _PrefixNameValuer implements _PrefixValuer {
  const _PrefixNameValuer();
  @override
  String from(PrefixV1 p, [PgdeLocalization l]) => p.l10n(l);
}

// _AtomValuer
abstract class _AtomValuer {
  String from(AtomV1 a, [PgdeLocalization l, Form p]);
}

class _AtomSymbolValuer implements _AtomValuer {
  const _AtomSymbolValuer();
  @override
  String from(AtomV1 a, [PgdeLocalization l, Form p]) => a.symbol;
}

class _AtomOneValuer implements _AtomValuer {
  const _AtomOneValuer();
  @override
  String from(AtomV1 a, [PgdeLocalization l, Form p]) => a.l10n(l, Form.one);
}

class _AtomOtherValuer implements _AtomValuer {
  const _AtomOtherValuer();
  @override
  String from(AtomV1 a, [PgdeLocalization l, Form p]) => a.l10n(l, Form.other);
}

class _AtomParseValuer implements _AtomValuer {
  const _AtomParseValuer();
  @override
  String from(AtomV1 a, [PgdeLocalization l, Form p]) => a.l10n(l, p);
}
