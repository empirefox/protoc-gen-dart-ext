import 'package:flutter/material.dart' show MaterialLocalizations;

import '../l10n.dart';
import '../plural.dart';
import '../units.dart';

abstract class Formatter<F> {
  String format(F v, MaterialLocalizations md, PgdeLocalizations l,
      {Form form, Show show});
}
