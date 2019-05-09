import 'dart:collection';
import 'package:intl/intl.dart';

import '../plural/plural.dart';

final HashMap<String, NumberFormat> _decimal = HashMap();
final HashMap<String, NumberFormat> _percent = HashMap();
final HashMap<String, NumberFormat> _scientific = HashMap();
final HashMap<String, NumberFormat> _currency = HashMap();
final HashMap<String, NumberFormat> _currencySimple = HashMap();
final HashMap<String, NumberFormat> _custom = HashMap();

abstract class NumberFormatterGetter {
  static const NumberFormatterGetter decimal = const _DecimalFormatterGetter();
  static const NumberFormatterGetter percent = const _PercentFormatterGetter();
  static const NumberFormatterGetter scientific =
      const _ScientificFormatterGetter();
  static const NumberFormatterGetter currency =
      const _CurrencyFormatterGetter();
  static const NumberFormatterGetter currencySimple =
      const _CurrencySimpleFormatterGetter();

  // must register before use
  static const NumberFormatterGetter custom = const _CustomFormatterGetter();

  static void registerCustom(String k, NumberFormat f) => _custom[k] = f;

  NumberFormat from(String key);
}

String pluralLocale(String locale) =>
    Intl.verifiedLocale(locale, rules.containsKey,
        onFailure: (locale) => 'root');

PluralFunc pluralRule(String locale) => rules[pluralLocale(locale)];

abstract class _LocaleFormatterGetter implements NumberFormatterGetter {
  HashMap<String, NumberFormat> get formatters;

  const _LocaleFormatterGetter();

  NumberFormat from(String locale) {
    locale = pluralLocale(locale);
    if (!formatters.containsKey(locale)) {
      formatters[locale] = create(locale);
    }
    return formatters[locale];
  }

  NumberFormat create(String locale);
}

class _DecimalFormatterGetter extends _LocaleFormatterGetter {
  HashMap<String, NumberFormat> get formatters => _decimal;
  const _DecimalFormatterGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.decimalPattern(locale);
}

class _PercentFormatterGetter extends _LocaleFormatterGetter {
  HashMap<String, NumberFormat> get formatters => _percent;
  const _PercentFormatterGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.percentPattern(locale);
}

class _ScientificFormatterGetter extends _LocaleFormatterGetter {
  HashMap<String, NumberFormat> get formatters => _scientific;
  const _ScientificFormatterGetter();
  @override
  NumberFormat create(String locale) => NumberFormat.scientificPattern(locale);
}

class _CurrencyFormatterGetter extends _LocaleFormatterGetter {
  HashMap<String, NumberFormat> get formatters => _currency;
  const _CurrencyFormatterGetter();
  @override
  NumberFormat create(String locale) =>
      NumberFormat.compactCurrency(locale: locale);
}

class _CurrencySimpleFormatterGetter extends _LocaleFormatterGetter {
  HashMap<String, NumberFormat> get formatters => _currencySimple;
  const _CurrencySimpleFormatterGetter();
  @override
  NumberFormat create(String locale) =>
      NumberFormat.compactSimpleCurrency(locale: locale);
}

class _CustomFormatterGetter implements NumberFormatterGetter {
  const _CustomFormatterGetter();
  @override
  NumberFormat from(String key) => _custom[key];
}
