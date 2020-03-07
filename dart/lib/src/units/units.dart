import '../plural.dart';
import '../l10n.dart';

import 'atom.dart';
import 'prefix.dart';

class Unit {
  static const showSymbol = Show();

  final Show show;
  final List<Cell> dots;
  final List<Cell> per;

  const Unit([this.show = showSymbol, this.dots, this.per]);

  String get symbol => l10n(null, Form.one, showSymbol);

  String l10n(PgdeLocalizations l, Form p, [Show show]) {
    show = show ?? this.show;
    if (show.off) return '';
    final buf = StringBuffer();
    _join(buf, null, dots, l, p, show);
    _join(buf, '/', per, l, p, show);
    return buf.toString();
  }

  void _join(StringBuffer buf, String prefix, List<Cell> cells,
      PgdeLocalizations l, Form p, Show show) {
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

  static const $byNumberPrefix = <_PrefixValuer>[
    symbolPrefix,
    namePrefix,
  ];

  static const symbol = _AtomSymbolValuer();
  static const one = _AtomOneValuer();
  static const other = _AtomOtherValuer();
  static const parse = _AtomParseValuer();

  static const $byNumber = <_AtomValuer>[
    symbol,
    one,
    other,
    parse,
  ];

  final bool off;
  final _PrefixValuer prefix;
  final _AtomValuer atom;

  const Show(
      [this.off = false, this.prefix = symbolPrefix, this.atom = symbol]);

  const Show._disabled()
      : off = true,
        prefix = null,
        atom = null;

  bool get willParse => !off && atom == parse;
}

class Cell {
  final Prefix prefix;
  final Atom atom;
  final String exponent;

  const Cell.int(this.prefix, this.atom, [int exponent = 1])
      : this.exponent = exponent == 1 ? '' : '$exponent';

  const Cell(this.prefix, this.atom, [this.exponent = ''])
      : assert(prefix != null),
        assert(atom != null);

  String l10n(_PrefixValuer pv, _AtomValuer av, PgdeLocalizations l, Form p) =>
      '${pv.from(prefix, l)}${av.from(atom, l, p)}$exponent';
}

// _PrefixValuer
abstract class _PrefixValuer {
  String from(Prefix p, [PgdeLocalizations l]);
}

class _PrefixSymbolValuer implements _PrefixValuer {
  const _PrefixSymbolValuer();
  @override
  String from(Prefix p, [PgdeLocalizations l]) => p.symbol;
}

class _PrefixNameValuer implements _PrefixValuer {
  const _PrefixNameValuer();
  @override
  String from(Prefix p, [PgdeLocalizations l]) => p.l10n(l);
}

// _AtomValuer
abstract class _AtomValuer {
  String from(Atom a, [PgdeLocalizations l, Form p]);
}

class _AtomSymbolValuer implements _AtomValuer {
  const _AtomSymbolValuer();
  @override
  String from(Atom a, [PgdeLocalizations l, Form p]) => a.symbol;
}

class _AtomOneValuer implements _AtomValuer {
  const _AtomOneValuer();
  @override
  String from(Atom a, [PgdeLocalizations l, Form p]) => a.l10n(l, Form.one);
}

class _AtomOtherValuer implements _AtomValuer {
  const _AtomOtherValuer();
  @override
  String from(Atom a, [PgdeLocalizations l, Form p]) => a.l10n(l, Form.other);
}

class _AtomParseValuer implements _AtomValuer {
  const _AtomParseValuer();
  @override
  String from(Atom a, [PgdeLocalizations l, Form p]) => a.l10n(l, p);
}
