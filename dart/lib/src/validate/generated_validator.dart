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
  void assertProto(T proto);
  void assertField(T proto, int tag);
  void assertOneof(T proto, int oneof);
}

// TODO refactor?
abstract class TranslateContext {
  String enumName(Type pbEnumType);
  String enumValue(ProtobufEnum value);
  String protoName(Type msgType);
  String oneofName(Type msgType, Type oneofEnumType);
  String fieldName(Type msgType, int tagNumber);
  String boolFieldValue(Type msgType, int tagNumber);

  // fallback format
  String fallbackFormat(ErrorRequiredInfo info, kConst);
}

class ErrorRequiredInfo {
  final MaterialLocalizations md;
  final PgdeLocalization l10n;
  final Formatter fmt;
  final TranslateContext ctx;
  final GeneratedMessage proto;
  final int tagNumber;

  const ErrorRequiredInfo(
      this.md, this.l10n, this.fmt, this.ctx, this.proto, this.tagNumber);

  Type get pbType => proto.runtimeType;
  get value => tagNumber != null ? proto.getFieldOrNull(tagNumber) : null;

  BuilderInfo get bi => proto.info_;
  FieldInfo get fi => bi.fieldInfo[tagNumber];
  String get messageName => bi.messageName;
  String get fieldAccessor => '$messageName.${bi.fieldName(tagNumber)}';
  String get fieldName => ctx.fieldName(pbType, tagNumber);

  String atomForm(AtomV1 a, Form p) => a.l10n(l10n, p);
  String unitForm(Form p) => numFmt.unit.l10n(l10n, p);

  bool get isNum => fmt is NumberFormatter;
  NumberFormatter get numFmt => fmt as NumberFormatter;

  Form plural(v, [bool ordinal]) => numFmt.plural(v, ordinal);

  ErrorRequiredInfo notField() =>
      ErrorRequiredInfo(md, l10n, fmt, ctx, proto, null);
}

abstract class ValidateError extends Error {
  final ErrorRequiredInfo info;

  final Type pbType;
  final String fieldName;
  final dynamic value;

  final PgdeLocalization l10n;
  final Formatter fmt;
  final TranslateContext ctx;

  ValidateError(this.info)
      : pbType = info.pbType,
        fieldName = info.fieldName,
        value = info.value,
        l10n = info.l10n,
        fmt = info.fmt,
        ctx = info.ctx;

  String format(v, [bool unitOff = false]) {
    if (info.fmt != null)
      return info.fmt
          .format(v, info.md, info.l10n, show: unitOff ? Show.disabled : null);

    // v is kConst
    if (v is ProtobufEnum) return ctx.enumValue(v);

    return ctx.fallbackFormat(info, v) ?? (v is String ? v : '$v');
  }

  @override
  String toString() => throw UnimplementedError('$runtimeType.toString');
}

//
// required
//

class RequiredError extends ValidateError {
  RequiredError(ErrorRequiredInfo info) : super(info);
  @override
  String toString() => l10n.validateRequired(fieldName);
}

class OneofRequiredError extends RequiredError {
  // eg MaterialColor_Color
  final Type oneofEnumType;
  OneofRequiredError(ErrorRequiredInfo info, this.oneofEnumType)
      : super(info.notField());
  @override
  String get fieldName => ctx.oneofName(pbType, oneofEnumType);
}

//
// be something
//

class BeSomethingError extends ValidateError {
  /// something: l10n.email, l10n.ipv4...
  final String something;
  BeSomethingError(ErrorRequiredInfo info, this.something) : super(info);
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
  final kConst;

  // for bytes
  final String kConstPrint;

  ConstError(ErrorRequiredInfo info, this.rule, this.kConst, [this.kConstPrint])
      : super(info);

  @override
  String toString() =>
      l10n.validateMustConst(fieldName, rule, kConstPrint ?? format(kConst));
}

class LenConstError extends ValidateError {
  /// any int or time mustConst rule: l10n.eq, l10n.gt, l10n.lte...
  final String rule;
  final int kConst;
  final AtomV1 lenAtom;

  LenConstError.byte(ErrorRequiredInfo info, this.rule, this.kConst)
      : lenAtom = AtomV1.byte,
        super(info);

  LenConstError.character(ErrorRequiredInfo info, this.rule, this.kConst)
      : lenAtom = AtomV1.character,
        super(info);

  String get lenUnit => info.atomForm(lenAtom, info.plural(kConst, false));

  @override
  String toString() {
    return l10n.validateMustConst(
        l10n.validateFieldLength(fieldName), rule, '$kConst$lenUnit');
  }
}

class ItemsLenConstError extends ValidateError {
  final int kConst;
  ItemsLenConstError(ErrorRequiredInfo info, this.kConst) : super(info);
  @override
  String toString() => l10n.validateMustConst(
      l10n.validateFieldItems(fieldName), l10n.validateEq, '$kConst');
}

class BoolError extends ValidateError {
  BoolError(ErrorRequiredInfo info) : super(info);
  @override
  String toString() => ctx.boolFieldValue(pbType, info.tagNumber);
}

//
// in or not
//

class InError extends ValidateError {
  final String rule;
  final List kConst;
  final List<String> kConstPrint;

  InError.inList(ErrorRequiredInfo info, this.kConst, {this.kConstPrint})
      : rule = info.l10n.validateOneof,
        super(info);

  InError.notInList(ErrorRequiredInfo info, this.kConst, {this.kConstPrint})
      : rule = info.l10n.validateOneofx,
        super(info);

  @override
  String toString() {
    final kConstWithUnit = kConstPrint ?? kConst.map((v) => format(v)).toList();
    return l10n.validateMustConst(fieldName, rule, '$kConstWithUnit');
  }
}

//
// range in or out
//

class ErrorRange {
  final ErrorRequiredInfo info;
  final start;
  final end;
  final bool isOut;
  final bool canEqStart;
  final bool canEqEnd;

  const ErrorRange(this.info, this.start, this.end,
      {this.isOut = false, this.canEqStart = false, this.canEqEnd = false});

  String get rule => isOut ? info.l10n.validateRangex : info.l10n.validateRange;

  String get prefix => (isOut ? !canEqStart : canEqStart) ? '[' : '(';
  String get suffix => (isOut ? !canEqEnd : canEqEnd) ? ']' : ')';
}

class RangeError extends ValidateError {
  final ErrorRange r;

  RangeError(this.r) : super(r.info);

  String get trField => fieldName;
  String get unit => info.isNum ? info.unitForm(Form.other) : '';

  // example: [1, 30)
  String get range =>
      "${r.prefix}${format(r.start, true)}, ${format(r.end, true)}${r.suffix}$unit";

  @override
  String toString() => l10n.validateMustConst(trField, r.rule, range);
}

class RangeLenError extends RangeError {
  final AtomV1 lenAtom;

  RangeLenError.byte(ErrorRange r)
      : lenAtom = AtomV1.byte,
        super(r);

  RangeLenError.character(ErrorRange r)
      : lenAtom = AtomV1.character,
        super(r);

  String get trField => l10n.validateFieldLength(fieldName);
  String get unit => info.atomForm(lenAtom, Form.other);
}

class RangeItemsLenError extends RangeError {
  RangeItemsLenError(ErrorRange r) : super(r);
  String get trField => l10n.validateFieldItems(fieldName);
  String get unit => '';
}

//
// within now, gt or lt
//

class WithinNowError extends ValidateError {
  final bool isGtNow;
  final Duration kConst;

  WithinNowError(ErrorRequiredInfo info, this.isGtNow, this.kConst)
      : super(info);

  @override
  String toString() {
    return isGtNow == null
        ? l10n.validateMustWithinNow
        : (isGtNow
            ? l10n.validateMustWithinGtNow
            : l10n.validateMustWithinLtNow)(fieldName, format(kConst));
  }
}
