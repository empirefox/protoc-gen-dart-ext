// DO NOT EDIT. Generated by protoc-gen-dart-ext/tools.

import '../plural.dart';
import '../l10n.dart';

abstract class _Valuer {
  String of(PgdeLocalizations l, Form p);
}

class _NoAtom implements _Valuer {
  const _NoAtom();
  String of(PgdeLocalizations l, Form p) => '';
}

class _Meter implements _Valuer {
  const _Meter();
  String of(PgdeLocalizations l, Form p) => l.atomMeter(p);
}

class _Foot implements _Valuer {
  const _Foot();
  String of(PgdeLocalizations l, Form p) => l.atomFoot(p);
}

class _Inch implements _Valuer {
  const _Inch();
  String of(PgdeLocalizations l, Form p) => l.atomInch(p);
}

class _Yard implements _Valuer {
  const _Yard();
  String of(PgdeLocalizations l, Form p) => l.atomYard(p);
}

class _Mile implements _Valuer {
  const _Mile();
  String of(PgdeLocalizations l, Form p) => l.atomMile(p);
}

class _NauticalMile implements _Valuer {
  const _NauticalMile();
  String of(PgdeLocalizations l, Form p) => l.atomNauticalMile(p);
}

class _LightYear implements _Valuer {
  const _LightYear();
  String of(PgdeLocalizations l, Form p) => l.atomLightYear(p);
}

class _Hectare implements _Valuer {
  const _Hectare();
  String of(PgdeLocalizations l, Form p) => l.atomHectare(p);
}

class _Are implements _Valuer {
  const _Are();
  String of(PgdeLocalizations l, Form p) => l.atomAre(p);
}

class _Liter implements _Valuer {
  const _Liter();
  String of(PgdeLocalizations l, Form p) => l.atomLiter(p);
}

class _Gallon implements _Valuer {
  const _Gallon();
  String of(PgdeLocalizations l, Form p) => l.atomGallon(p);
}

class _Barrel implements _Valuer {
  const _Barrel();
  String of(PgdeLocalizations l, Form p) => l.atomBarrel(p);
}

class _Gram implements _Valuer {
  const _Gram();
  String of(PgdeLocalizations l, Form p) => l.atomGram(p);
}

class _Ton implements _Valuer {
  const _Ton();
  String of(PgdeLocalizations l, Form p) => l.atomTon(p);
}

class _Pound implements _Valuer {
  const _Pound();
  String of(PgdeLocalizations l, Form p) => l.atomPound(p);
}

class _Ounce implements _Valuer {
  const _Ounce();
  String of(PgdeLocalizations l, Form p) => l.atomOunce(p);
}

class _Second implements _Valuer {
  const _Second();
  String of(PgdeLocalizations l, Form p) => l.atomSecond(p);
}

class _Minute implements _Valuer {
  const _Minute();
  String of(PgdeLocalizations l, Form p) => l.atomMinute(p);
}

class _Hour implements _Valuer {
  const _Hour();
  String of(PgdeLocalizations l, Form p) => l.atomHour(p);
}

class _Day implements _Valuer {
  const _Day();
  String of(PgdeLocalizations l, Form p) => l.atomDay(p);
}

class _Week implements _Valuer {
  const _Week();
  String of(PgdeLocalizations l, Form p) => l.atomWeek(p);
}

class _Month implements _Valuer {
  const _Month();
  String of(PgdeLocalizations l, Form p) => l.atomMonth(p);
}

class _Year implements _Valuer {
  const _Year();
  String of(PgdeLocalizations l, Form p) => l.atomYear(p);
}

class _Century implements _Valuer {
  const _Century();
  String of(PgdeLocalizations l, Form p) => l.atomCentury(p);
}

class _SecondAngle implements _Valuer {
  const _SecondAngle();
  String of(PgdeLocalizations l, Form p) => l.atomSecondAngle(p);
}

class _MinuteAngle implements _Valuer {
  const _MinuteAngle();
  String of(PgdeLocalizations l, Form p) => l.atomMinuteAngle(p);
}

class _Degree implements _Valuer {
  const _Degree();
  String of(PgdeLocalizations l, Form p) => l.atomDegree(p);
}

class _Ampere implements _Valuer {
  const _Ampere();
  String of(PgdeLocalizations l, Form p) => l.atomAmpere(p);
}

class _Electronvolt implements _Valuer {
  const _Electronvolt();
  String of(PgdeLocalizations l, Form p) => l.atomElectronvolt(p);
}

class _Bel implements _Valuer {
  const _Bel();
  String of(PgdeLocalizations l, Form p) => l.atomBel(p);
}

class _Kelvin implements _Valuer {
  const _Kelvin();
  String of(PgdeLocalizations l, Form p) => l.atomKelvin(p);
}

class _Mole implements _Valuer {
  const _Mole();
  String of(PgdeLocalizations l, Form p) => l.atomMole(p);
}

class _Candela implements _Valuer {
  const _Candela();
  String of(PgdeLocalizations l, Form p) => l.atomCandela(p);
}

class _Percent implements _Valuer {
  const _Percent();
  String of(PgdeLocalizations l, Form p) => l.atomPercent(p);
}

class _PerThousand implements _Valuer {
  const _PerThousand();
  String of(PgdeLocalizations l, Form p) => l.atomPerThousand(p);
}

class _Bit implements _Valuer {
  const _Bit();
  String of(PgdeLocalizations l, Form p) => l.atomBit(p);
}

class _Byte implements _Valuer {
  const _Byte();
  String of(PgdeLocalizations l, Form p) => l.atomByte(p);
}

class _Character implements _Valuer {
  const _Character();
  String of(PgdeLocalizations l, Form p) => l.atomCharacter(p);
}

class _Word implements _Valuer {
  const _Word();
  String of(PgdeLocalizations l, Form p) => l.atomWord(p);
}

class _Radian implements _Valuer {
  const _Radian();
  String of(PgdeLocalizations l, Form p) => l.atomRadian(p);
}

class _Steradian implements _Valuer {
  const _Steradian();
  String of(PgdeLocalizations l, Form p) => l.atomSteradian(p);
}

class _Hertz implements _Valuer {
  const _Hertz();
  String of(PgdeLocalizations l, Form p) => l.atomHertz(p);
}

class _Newton implements _Valuer {
  const _Newton();
  String of(PgdeLocalizations l, Form p) => l.atomNewton(p);
}

class _Pascal implements _Valuer {
  const _Pascal();
  String of(PgdeLocalizations l, Form p) => l.atomPascal(p);
}

class _Joule implements _Valuer {
  const _Joule();
  String of(PgdeLocalizations l, Form p) => l.atomJoule(p);
}

class _Watt implements _Valuer {
  const _Watt();
  String of(PgdeLocalizations l, Form p) => l.atomWatt(p);
}

class _Coulomb implements _Valuer {
  const _Coulomb();
  String of(PgdeLocalizations l, Form p) => l.atomCoulomb(p);
}

class _Volt implements _Valuer {
  const _Volt();
  String of(PgdeLocalizations l, Form p) => l.atomVolt(p);
}

class _Farad implements _Valuer {
  const _Farad();
  String of(PgdeLocalizations l, Form p) => l.atomFarad(p);
}

class _Ohm implements _Valuer {
  const _Ohm();
  String of(PgdeLocalizations l, Form p) => l.atomOhm(p);
}

class _Siemens implements _Valuer {
  const _Siemens();
  String of(PgdeLocalizations l, Form p) => l.atomSiemens(p);
}

class _Weber implements _Valuer {
  const _Weber();
  String of(PgdeLocalizations l, Form p) => l.atomWeber(p);
}

class _Tesla implements _Valuer {
  const _Tesla();
  String of(PgdeLocalizations l, Form p) => l.atomTesla(p);
}

class _Henry implements _Valuer {
  const _Henry();
  String of(PgdeLocalizations l, Form p) => l.atomHenry(p);
}

class _DegreeCelsius implements _Valuer {
  const _DegreeCelsius();
  String of(PgdeLocalizations l, Form p) => l.atomDegreeCelsius(p);
}

class _Lumen implements _Valuer {
  const _Lumen();
  String of(PgdeLocalizations l, Form p) => l.atomLumen(p);
}

class _Lux implements _Valuer {
  const _Lux();
  String of(PgdeLocalizations l, Form p) => l.atomLux(p);
}

class _Becquerel implements _Valuer {
  const _Becquerel();
  String of(PgdeLocalizations l, Form p) => l.atomBecquerel(p);
}

class _Gray implements _Valuer {
  const _Gray();
  String of(PgdeLocalizations l, Form p) => l.atomGray(p);
}

class _Sievert implements _Valuer {
  const _Sievert();
  String of(PgdeLocalizations l, Form p) => l.atomSievert(p);
}

class _Katal implements _Valuer {
  const _Katal();
  String of(PgdeLocalizations l, Form p) => l.atomKatal(p);
}

class Atom {
  static const noAtom = const Atom._('', const _NoAtom());

  static const meter = const Atom._(r'm', const _Meter());

  static const foot = const Atom._(r'ft', const _Foot());

  static const inch = const Atom._(r'in', const _Inch());

  static const yard = const Atom._(r'yd', const _Yard());

  static const mile = const Atom._(r'mi', const _Mile());

  static const nauticalMile = const Atom._(r'NM', const _NauticalMile());

  static const lightYear = const Atom._(r'LY', const _LightYear());

  static const hectare = const Atom._(r'ha', const _Hectare());

  static const are = const Atom._(r'a', const _Are());

  static const liter = const Atom._(r'L', const _Liter());

  static const gallon = const Atom._(r'gal', const _Gallon());

  static const barrel = const Atom._(r'bbl', const _Barrel());

  static const gram = const Atom._(r'g', const _Gram());

  static const ton = const Atom._(r't', const _Ton());

  static const pound = const Atom._(r'lbs', const _Pound());

  static const ounce = const Atom._(r'oz', const _Ounce());

  static const second = const Atom._(r's', const _Second());

  static const minute = const Atom._(r'min', const _Minute());

  static const hour = const Atom._(r'h', const _Hour());

  static const day = const Atom._(r'd', const _Day());

  static const week = const Atom._(r'week', const _Week());

  static const month = const Atom._(r'month', const _Month());

  static const year = const Atom._(r'yr', const _Year());

  static const century = const Atom._(r'century', const _Century());

  static const secondAngle = const Atom._("'" "'", const _SecondAngle());

  static const minuteAngle = const Atom._("'", const _MinuteAngle());

  static const degree = const Atom._(r'°', const _Degree());

  static const ampere = const Atom._(r'A', const _Ampere());

  static const electronvolt = const Atom._(r'eV', const _Electronvolt());

  static const bel = const Atom._(r'B', const _Bel());

  static const kelvin = const Atom._(r'K', const _Kelvin());

  static const mole = const Atom._(r'mol', const _Mole());

  static const candela = const Atom._(r'cd', const _Candela());

  static const percent = const Atom._(r'%', const _Percent());

  static const perThousand = const Atom._(r'‰', const _PerThousand());

  static const bit = const Atom._(r'b', const _Bit());

  static const byte = const Atom._(r'B', const _Byte());

  static const character = const Atom._(r'character', const _Character());

  static const word = const Atom._(r'word', const _Word());

  static const radian = const Atom._(r'rad', const _Radian());

  static const steradian = const Atom._(r'sr', const _Steradian());

  static const hertz = const Atom._(r'Hz', const _Hertz());

  static const newton = const Atom._(r'N', const _Newton());

  static const pascal = const Atom._(r'Pa', const _Pascal());

  static const joule = const Atom._(r'J', const _Joule());

  static const watt = const Atom._(r'W', const _Watt());

  static const coulomb = const Atom._(r'C', const _Coulomb());

  static const volt = const Atom._(r'V', const _Volt());

  static const farad = const Atom._(r'F', const _Farad());

  static const ohm = const Atom._(r'Ω', const _Ohm());

  static const siemens = const Atom._(r'S', const _Siemens());

  static const weber = const Atom._(r'Wb', const _Weber());

  static const tesla = const Atom._(r'T', const _Tesla());

  static const henry = const Atom._(r'H', const _Henry());

  static const degreeCelsius = const Atom._(r'°C', const _DegreeCelsius());

  static const lumen = const Atom._(r'lm', const _Lumen());

  static const lux = const Atom._(r'lx', const _Lux());

  static const becquerel = const Atom._(r'Bq', const _Becquerel());

  static const gray = const Atom._(r'Gy', const _Gray());

  static const sievert = const Atom._(r'Sv', const _Sievert());

  static const katal = const Atom._(r'kat', const _Katal());

  static const $byNumber = <Atom>[
    noAtom,
    meter,
    foot,
    inch,
    yard,
    mile,
    nauticalMile,
    lightYear,
    hectare,
    are,
    liter,
    gallon,
    barrel,
    gram,
    ton,
    pound,
    ounce,
    second,
    minute,
    hour,
    day,
    week,
    month,
    year,
    century,
    secondAngle,
    minuteAngle,
    degree,
    ampere,
    electronvolt,
    bel,
    kelvin,
    mole,
    candela,
    percent,
    perThousand,
    bit,
    byte,
    character,
    word,
    radian,
    steradian,
    hertz,
    newton,
    pascal,
    joule,
    watt,
    coulomb,
    volt,
    farad,
    ohm,
    siemens,
    weber,
    tesla,
    henry,
    degreeCelsius,
    lumen,
    lux,
    becquerel,
    gray,
    sievert,
    katal,
  ];

  final String symbol;
  final _Valuer _v;
  const Atom._(this.symbol, this._v);
  const Atom.symbol(this.symbol) : _v = null;
  String l10n(PgdeLocalizations l10n, Form form) =>
      l10n == null ? symbol : _v?.of(l10n, form) ?? symbol;
}
