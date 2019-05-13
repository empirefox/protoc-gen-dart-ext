import 'package:flutter/material.dart' show MaterialLocalizations;

import '../l10n/pgde.l10n.dart';
import '../plural/plural.dart';
import '../units/units.dart';

abstract class Formatter<F> {
  String format(F v, MaterialLocalizations md, PgdeLocalization l,
      {Form form, Show show});
}
