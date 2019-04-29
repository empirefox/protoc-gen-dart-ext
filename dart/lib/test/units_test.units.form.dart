import 'dart:collection' show HashSet;
import 'package:flutter/foundation.dart' show SynchronousFuture;
import 'package:flutter/material.dart' show Locale, LocalizationsDelegate;

import './units_test.units.form.base.dart';

export './units_test.units.form.base.dart';

class UnitsLocalizationEn extends UnitsLocalization {
  const UnitsLocalizationEn();
  @override
  String get howManyEnum => 'How many enum';
  @override
  String get howManyEnum_ZERO => 'ONE';
  @override
  String get howManyEnum_ONE => 'TWO';
  @override
  String get howManyEnum_TWO => 'ZERO';
}

class UnitsLocalizationDelegate
    extends LocalizationsDelegate<UnitsLocalization> {
  const UnitsLocalizationDelegate._();
  static const instance = UnitsLocalizationDelegate._();

  @override
  bool isSupported(Locale locale) =>
      kSupportedLanguages.contains(locale.languageCode);

  @override
  Future<UnitsLocalization> load(Locale locale) =>
      SynchronousFuture<UnitsLocalization>(getTranslation(locale));

  @override
  bool shouldReload(UnitsLocalizationDelegate old) => false;

  static final Set<String> kSupportedLanguages =
      HashSet<String>.from(const <String>[
    'en',
    'zh',
  ]);

  static UnitsLocalization getTranslation(Locale locale) {
    switch (locale.languageCode) {
      case 'en':
        {
          switch (locale.countryCode) {
            case 'AU':
              return UnitsLocalizationEnAu();
            case 'CA':
              return UnitsLocalizationEnCa();
          }
          return UnitsLocalizationEn();
        }
    }
    assert(false, 'getTranslation() called for unsupported locale "$locale"');
    return null;
  }
}
