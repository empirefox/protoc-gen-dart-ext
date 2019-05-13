import 'dart:collection';
import 'package:intl/intl.dart';

import '../plural/plural.dart';
import '../units/units.dart';

import 'formatter.dart';

class NumberFormatter implements Formatter<num> {
  final NumberFormatGetter getter;
  final String locale;
  final bool ordinal;
  final Unit unit;

  const NumberFormatter(this.getter,
      {this.locale, this.ordinal = false, this.unit});

  String format(v, md, l, {form, show}) {
    var s = getter?.from(locale)?.format(v) ?? '$v';
    if (unit != null) {
      if (form == null && (show ?? unit.show).willParse) form = plural(v);
      s += unit.l10n(l, form, show);
    }
    return s;
  }

  Form plural(v, [bool ordinal]) =>
      pluralRule(locale)(v, ordinal ?? this.ordinal);
}

final HashMap<String, NumberFormat> _decimal = HashMap();
final HashMap<String, NumberFormat> _percent = HashMap();
final HashMap<String, NumberFormat> _scientific = HashMap();
final HashMap<String, NumberFormat> _currency = HashMap();
final HashMap<String, NumberFormat> _currencySimple = HashMap();
final HashMap<String, NumberFormat> _custom = HashMap();

abstract class NumberFormatGetter {
  static const NumberFormatGetter decimal = const _DecimalFormatGetter();
  static const NumberFormatGetter percent = const _PercentFormatGetter();
  static const NumberFormatGetter scientific = const _ScientificFormatGetter();
  static const NumberFormatGetter currency = const _CurrencyFormatGetter();
  static const NumberFormatGetter currencySimple =
      const _CurrencySimpleFormatGetter();

  // must register before use
  static const NumberFormatGetter custom = const _CustomFormatGetter();

  static void registerCustom(String k, NumberFormat f) => _custom[k] = f;

  NumberFormat from(String key);
}

String pluralLocale(String locale) =>
    Intl.verifiedLocale(locale, rules.containsKey,
        onFailure: (locale) => 'root');

PluralFunc pluralRule(String locale) => rules[pluralLocale(locale)];

abstract class _LocaleFormatGetter implements NumberFormatGetter {
  HashMap<String, NumberFormat> get formatters;

  const _LocaleFormatGetter();

  NumberFormat from(String locale) {
    locale = pluralLocale(locale);
    if (!formatters.containsKey(locale)) {
      formatters[locale] = create(locale);
    }
    return formatters[locale];
  }

  NumberFormat create(String locale);
}

class _DecimalFormatGetter extends _LocaleFormatGetter {
  HashMap<String, NumberFormat> get formatters => _decimal;
  const _DecimalFormatGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.decimalPattern(locale);
}

class _PercentFormatGetter extends _LocaleFormatGetter {
  HashMap<String, NumberFormat> get formatters => _percent;
  const _PercentFormatGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.percentPattern(locale);
}

class _ScientificFormatGetter extends _LocaleFormatGetter {
  HashMap<String, NumberFormat> get formatters => _scientific;
  const _ScientificFormatGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.scientificPattern(locale);
}

class _CurrencyFormatGetter extends _LocaleFormatGetter {
  HashMap<String, NumberFormat> get formatters => _currency;
  const _CurrencyFormatGetter();
  @override
  NumberFormat create(String locale) =>
      NumberFormat.compactCurrency(locale: locale);
}

class _CurrencySimpleFormatGetter extends _LocaleFormatGetter {
  HashMap<String, NumberFormat> get formatters => _currencySimple;
  const _CurrencySimpleFormatGetter();
  @override
  NumberFormat create(String locale) =>
      NumberFormat.compactSimpleCurrency(locale: locale);
}

class _CustomFormatGetter implements NumberFormatGetter {
  const _CustomFormatGetter();
  @override
  NumberFormat from(String key) => _custom[key];
}
