import 'package:flutter/material.dart' show MaterialLocalizations;
import 'package:pgde/src/format/number_format.dart';
import 'package:protobuf/protobuf.dart';

import '../format/formatter.dart';
import '../l10n/pgde.l10n.dart';
import '../plural/plural.dart';
import '../units/units.dart';

/// The base interface for all generated PGV validators.
/// Asserts validation rules on a protobuf object.
/// Throw [ValidationError] with the first validation error encountered.
abstract class GeneratedValidator<T extends GeneratedMessage> {
  void assertProto();
  void assertField(int tag);
  void assertOneof(int oneof);
}

class ValidateInfo<T extends GeneratedMessage> {
  final MaterialLocalizations md;
  final PgdeLocalization l10n;
  final Formatter fmt;
  final T proto;

  const ValidateInfo(this.md, this.l10n, this.fmt, this.proto);

  Type get pbType => proto.runtimeType;

  BuilderInfo get bi => proto.info_;
  String get messageName => bi.messageName;

  String atomForm(Atom a, Form p) => a.l10n(l10n, p);
  String unitForm(Form p) => numFmt.unit.l10n(l10n, p);

  bool get isNum => fmt is NumberFormatter;
  NumberFormatter get numFmt => fmt as NumberFormatter;

  Form plural(v, [bool ordinal]) => numFmt.plural(v, ordinal);

  ValidateInfo<C> clone<C extends GeneratedMessage>(C proto) =>
      ValidateInfo(md, l10n, fmt, proto);
}

abstract class ValidateError extends Error {
  final ValidateInfo info;
  final int tagNumber;
  final String fieldName;

  final Type pbType;
  final MaterialLocalizations md;
  final PgdeLocalization l10n;
  final Formatter fmt;

  BuilderInfo get bi => info.bi;
  get value => tagNumber != null ? info.proto.getFieldOrNull(tagNumber) : null;
  FieldInfo get fi => info.bi.fieldInfo[tagNumber];
  String get fieldAccessor => '${info.messageName}.${bi.fieldName(tagNumber)}';

  ValidateError(this.info, this.tagNumber, this.fieldName)
      : pbType = info.pbType,
        md = info.md,
        l10n = info.l10n,
        fmt = info.fmt;

  String format(v, [bool unitOff = false]) {
    if (fmt != null)
      return fmt.format(v, md, l10n, show: unitOff ? Show.disabled : null);
    return v is String ? v : '$v';
  }

  @override
  String toString() => throw UnimplementedError('$runtimeType.toString');
}

//
// required
//

class RequiredError extends ValidateError {
  RequiredError(ValidateInfo info, int tagNumber, String fieldName)
      : super(info, tagNumber, fieldName);
  @override
  String toString() => l10n.validateRequired(fieldName);
}

class OneofRequiredError extends RequiredError {
  // from l10n
  final String oneofName;
  OneofRequiredError(ValidateInfo info, this.oneofName)
      : super(info, null, null);
  @override
  String get fieldName => oneofName;
}

//
// be something
//

class BeSomethingError extends ValidateError {
  /// something: l10n.email, l10n.ipv4...
  final String something;
  BeSomethingError(
      ValidateInfo info, int tagNumber, String fieldName, this.something)
      : super(info, tagNumber, fieldName);
  @override
  String toString() => l10n.validateMustBe(fieldName, something);
}

//
// eq gt lte pattern contains const
//

class ConstError extends ValidateError {
  /// any mustConst rule: l10n.eq, l10n.gt, l10n.lte...
  final String rule;

  // kConst can be l10n.now
  // for bytes: 'unicode printable string' or '\x${byte_hex}\x${byte_hex}'
  final kConst;

  ConstError(ValidateInfo info, int tagNumber, String fieldName, this.rule,
      this.kConst)
      : super(info, tagNumber, fieldName);

  @override
  String toString() => l10n.validateMustConst(fieldName, rule, format(kConst));
}

class LenConstError extends ValidateError {
  /// any int or time mustConst rule: l10n.eq, l10n.gt, l10n.lte...
  final String rule;
  final int kConst;
  final Atom lenAtom;

  LenConstError.byte(ValidateInfo info, int tagNumber, String fieldName,
      this.rule, this.kConst)
      : lenAtom = Atom.byte,
        super(info, tagNumber, fieldName);

  LenConstError.character(ValidateInfo info, int tagNumber, String fieldName,
      this.rule, this.kConst)
      : lenAtom = Atom.character,
        super(info, tagNumber, fieldName);

  String get lenUnit => info.atomForm(lenAtom, info.plural(kConst, false));

  @override
  String toString() {
    return l10n.validateMustConst(
        l10n.validateFieldLength(fieldName), rule, '$kConst$lenUnit');
  }
}

class ItemsLenConstError extends ValidateError {
  final String rule;
  final int kConst;
  ItemsLenConstError(ValidateInfo info, int tagNumber, String fieldName,
      this.rule, this.kConst)
      : super(info, tagNumber, fieldName);
  @override
  String toString() => l10n.validateMustConst(
      l10n.validateFieldItems(fieldName), rule, '$kConst');
}

class BoolError extends ValidateError {
  final String kConst;
  BoolError(ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : super(info, tagNumber, fieldName);
  @override
  String toString() => kConst;
}

//
// in or not
//

class InError extends ValidateError {
  final String rule;
  final List kConst;

  InError(ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : rule = info.l10n.validateOneof,
        super(info, tagNumber, fieldName);

  InError.not(ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : rule = info.l10n.validateOneofx,
        super(info, tagNumber, fieldName);

  String get unit => info.isNum ? info.unitForm(Form.other) : '';

  @override
  String toString() {
    final kConstPrint = kConst.map((v) => format(v)).toList();
    return l10n.validateMustConst(fieldName, rule, '$kConstPrint$unit');
  }
}

//
// range in or out
//

class ErrorRange {
  final start;
  final end;
  final bool isOut;
  final bool eqStart;
  final bool eqEnd;

  const ErrorRange.inEE(this.start, this.end)
      : isOut = false,
        eqStart = true,
        eqEnd = true;

  const ErrorRange.inNE(this.start, this.end)
      : isOut = false,
        eqStart = false,
        eqEnd = true;

  const ErrorRange.inEN(this.start, this.end)
      : isOut = false,
        eqStart = true,
        eqEnd = false;

  const ErrorRange.inNN(this.start, this.end)
      : isOut = false,
        eqStart = false,
        eqEnd = false;

  const ErrorRange.outEE(this.start, this.end)
      : isOut = true,
        eqStart = true,
        eqEnd = true;

  const ErrorRange.outNE(this.start, this.end)
      : isOut = true,
        eqStart = false,
        eqEnd = true;

  const ErrorRange.outEN(this.start, this.end)
      : isOut = true,
        eqStart = true,
        eqEnd = false;

  const ErrorRange.outNN(this.start, this.end)
      : isOut = true,
        eqStart = false,
        eqEnd = false;

  String get prefix => (isOut ? !eqStart : eqStart) ? '[' : '(';
  String get suffix => (isOut ? !eqEnd : eqEnd) ? ']' : ')';
}

class RangeError extends ValidateError {
  final ErrorRange r;

  RangeError(ValidateInfo info, int tagNumber, String fieldName, this.r)
      : super(info, tagNumber, fieldName);

  String get trField => fieldName;
  String get unit => info.isNum ? info.unitForm(Form.other) : '';

  String get rule => r.isOut ? l10n.validateRangex : l10n.validateRange;

  // example: [1, 30)
  String get range =>
      "${r.prefix}${format(r.start, true)}, ${format(r.end, true)}${r.suffix}$unit";

  @override
  String toString() => l10n.validateMustConst(trField, rule, range);
}

class RangeLenError extends RangeError {
  final Atom lenAtom;

  RangeLenError.byte(
      ValidateInfo info, int tagNumber, String fieldName, ErrorRange r)
      : lenAtom = Atom.byte,
        super(info, tagNumber, fieldName, r);

  RangeLenError.character(
      ValidateInfo info, int tagNumber, String fieldName, ErrorRange r)
      : lenAtom = Atom.character,
        super(info, tagNumber, fieldName, r);

  String get trField => l10n.validateFieldLength(fieldName);
  String get unit => info.atomForm(lenAtom, Form.other);
}

class RangeItemsLenError extends RangeError {
  RangeItemsLenError(
      ValidateInfo info, int tagNumber, String fieldName, ErrorRange r)
      : super(info, tagNumber, fieldName, r);
  String get trField => l10n.validateFieldItems(fieldName);
  String get unit => '';
}

//
// within now, gt or lt
//

class WithinError extends ValidateError {
  final String Function(String field, String duration) rule;
  final Duration kConst;

  WithinError.now(
      ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : rule = info.l10n.validateMustWithinNow,
        super(info, tagNumber, fieldName);

  WithinError.gtNow(
      ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : rule = info.l10n.validateMustWithinGtNow,
        super(info, tagNumber, fieldName);

  WithinError.ltNow(
      ValidateInfo info, int tagNumber, String fieldName, this.kConst)
      : rule = info.l10n.validateMustWithinLtNow,
        super(info, tagNumber, fieldName);

  @override
  String toString() => rule(fieldName, format(kConst));
}
