import 'package:intl/intl.dart';

// Only used to generate arb!
// make gen_validate_arb
// make rewrite_validate_arb
class Translator {
  const Translator();

  String fieldLength(String field) =>
      Intl.message('$field length', name: 'fieldLength', args: [field]);

  String fieldItems(String field) =>
      Intl.message('The quantity of $field', name: 'fieldItems', args: [field]);

  /// used as kConst by GtLtError
  String get now => Intl.message('now', name: 'now');

  String get eq => Intl.message('equal', name: 'eq');
  String get gt => Intl.message('be greater than', name: 'gt');
  String get gte => Intl.message('be greater than or equal to', name: 'gte');
  String get lt => Intl.message('less than', name: 'lt');
  String get lte => Intl.message('less than or equal to', name: 'lte');
  String get oneof => Intl.message('be one of', name: 'oneof');
  String get oneofx => Intl.message('not be one of', name: 'oneofx');
  String get range => Intl.message('be inside range', name: 'range');
  String get rangex => Intl.message('be outside range', name: 'rangex');
  String get pattern => Intl.message('match the pattern', name: 'pattern');
  String get prefix => Intl.message('have prefix', name: 'prefix');
  String get suffix => Intl.message('have suffix', name: 'suffix');
  String get contains => Intl.message('contain', name: 'contains');

  String mustConst(String field, String rule, kConstWithUnit) => Intl.message(
        "$field must $rule $kConstWithUnit",
        name: 'mustConst',
        args: [field, rule, kConstWithUnit],
      );

  String get unique => Intl.message('unique', name: 'unique');
  String get noSparse =>
      Intl.message('a non-sparse map, all pairs must be non-nil',
          name: 'noSparse');
  String get email => Intl.message('email', name: 'email');
  String get hostname => Intl.message('hostname', name: 'hostname');
  String get ip => Intl.message('ip', name: 'ip');
  String get ipv4 => Intl.message('ipv4', name: 'ipv4');
  String get ipv6 => Intl.message('ipv6', name: 'ipv6');
  String get uri => Intl.message('uri', name: 'uri');
  String get uriRef => Intl.message('uri ref', name: 'uriRef');

  String mustBe(String field, String something) => Intl.message(
        "$field must be $something",
        name: 'mustBe',
        args: [field, something],
      );

  String mustWithinNow(String field, String duration) => Intl.message(
        '$field must be within $duration of now',
        name: 'mustWithinNow',
        args: [field, duration],
      );

  String mustWithinGtNow(String field, String duration) => Intl.message(
        '$field must be greater than now within $duration',
        name: 'mustWithinGtNow',
        args: [field, duration],
      );

  String mustWithinLtNow(String field, String duration) => Intl.message(
        '$field must be less than now within $duration',
        name: 'mustWithinLtNow',
        args: [field, duration],
      );

  String required(String field) => Intl.message(
        '$field is required',
        name: 'required',
        args: [field],
      );
}
