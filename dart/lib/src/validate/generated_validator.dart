import 'package:flutter/material.dart' show MaterialLocalizations;
import 'package:pgde/src/format/number_format.dart';
import 'package:protobuf/protobuf.dart';

import '../format/formatter.dart';
import '../l10n/pgde.l10n.dart';
import '../plural/plural.dart';
import '../units/units.dart';

import 'translate.dart';

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
  String oneofName(Type msgType, int oneof);
  String fieldName(Type msgType, int tagNumber);
  String boolFieldValue(Type msgType, int tagNumber);

  // fallback format
  String fallbackFormat(kConst) => kConst is String ? kConst : '$kConst';
}

class ValidateErrorInfo {
  final MaterialLocalizations md;
  final PgdeLocalization l10n;
  final Formatter fmt;
  final Translator tr;
  final TranslateContext ctx;
  final GeneratedMessage proto;
  final int tagNumber;

  const ValidateErrorInfo(this.md, this.l10n, this.fmt, this.tr, this.ctx,
      this.proto, this.tagNumber);

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

  ValidateErrorInfo notField() =>
      ValidateErrorInfo(md, l10n, fmt, tr, ctx, proto, null);
}

abstract class ValidateError extends Error {
  final ValidateErrorInfo info;

  final Type pbType;

  final dynamic value;

  final Formatter fmt;

  // for string_len string_len_bytes bytes_len
  // eq_const_len
  final AtomV1 lenAtom;

  final bool isItems;

  ValidateError(this.info, {this.lenAtom, this.isItems})
      : pbType = info.pbType,
        value = info.value,
        fmt = info.fmt,
        assert(lenAtom == null ||
            lenAtom == AtomV1.bit ||
            lenAtom == AtomV1.byte ||
            lenAtom == AtomV1.character),
        assert(lenAtom == null || !isItems);

  bool get isLen => lenAtom != null;
  bool get isLenOrItems => isLen || isItems;

  String get fieldLenOrItemsOrName => isLen
      ? tr.fieldLength(fieldName)
      : (isItems ? tr.fieldItems(fieldName) : fieldName);

  String get fieldName => info.fieldName;
  Translator get tr => info.tr;
  TranslateContext get ctx => info.ctx;

  // only aviliable when lenAtom exist
  String lenUnit(kConst) => info.atomForm(lenAtom, info.plural(kConst, false));
  String get lenOther => info.atomForm(lenAtom, Form.other);

  // only aviliable when isNum
  String get unitOtherOr => info.isNum ? info.unitForm(Form.other) : null;
  String unitL10nOr(v) => info.isNum ? info.unitForm(info.plural(v)) : null;

  String format(v) {
    if (lenAtom != null) return '$v';

    if (info.fmt != null) return info.fmt.format(v, info.md, info.l10n);

    // v is kConst
    if (v is ProtobufEnum) return ctx.enumValue(v);

    final pf = formatPreFallback(v);
    if (pf != null) return pf;

    return ctx.fallbackFormat(v);
  }

  String formatPreFallback(v) => null;

  @override
  String toString() => throw UnimplementedError('$runtimeType.toString');
}

//
// required
//

class RequiredError extends ValidateError {
  RequiredError(ValidateErrorInfo info) : super(info);
  @override
  String toString() => tr.required(fieldName);
}

class OneofRequiredError extends RequiredError {
  final int oneof;
  OneofRequiredError(ValidateErrorInfo info, this.oneof)
      : super(info.notField());
  @override
  String get fieldName => ctx.oneofName(pbType, oneof);
}

//
// be something
//

class BeSomethingError extends ValidateError {
  /// something: tr.email, tr.ipv4...
  final String something;
  BeSomethingError(ValidateErrorInfo info, this.something) : super(info);
  @override
  String toString() => tr.mustBe(fieldName, something);
}

//
// eq const
//

class ConstError extends ValidateError {
  final dynamic kConst;
  // for bytes
  final String kConstPrint;

  ConstError(ValidateErrorInfo info, this.kConst,
      {this.kConstPrint, final AtomV1 lenAtom})
      : super(info, lenAtom: lenAtom);

  String get fConst => format(kConst);

  @override
  String toString() {
    // string_len, string_len_bytes..., kConstPrint is empty
    if (isLen)
      return tr.mustConst(
          tr.fieldLength(fieldName), tr.eq, fConst, lenUnit(kConst));
    // bytes already converted to string
    if (kConstPrint != null) return tr.mustConst(fieldName, tr.eq, kConstPrint);
    // get from l10n
    if (kConst is bool) return ctx.boolFieldValue(pbType, info.tagNumber);
    // others like bytes ProtobufEnum Duration and Datetime or any other numbers
    return tr.mustConst(fieldName, tr.eq, fConst, unitL10nOr(kConst));
  }
}

//
// in or not
//

class InError extends ValidateError {
  final List kConst;
  final List<String> kConstPrint;
  final bool isNot;

  InError(ValidateErrorInfo info, this.kConst,
      {this.kConstPrint, this.isNot = false})
      : super(info);

  InError.not(ValidateErrorInfo info, List kConst, {List<String> kConstPrint})
      : this(info, kConst, kConstPrint: kConstPrint, isNot: true);

  @override
  String toString() {
    return tr.mustConst(
        fieldName,
        isNot ? tr.oneofx : tr.oneof,
        // number_in string_in bytes_in any_in
        kConstPrint ??
            // others like enum.defined_only duration.in
            kConst.map((v) => format(v)).toList(),
        unitOtherOr);
  }
}

//  Overflow of uint or int32?

//
// gt or lt
//

class GtLtError extends ValidateError {
  // kConst can be tr.now
  final dynamic kConst;

  /// rule: tr.gt, tr.lte...
  final String rule;

  GtLtError(ValidateErrorInfo info, this.rule, this.kConst,
      {AtomV1 lenAtom, bool isItems})
      : super(info, lenAtom: lenAtom, isItems: isItems);

  @override
  String toString() {
    return tr.mustConst(
        fieldLenOrItemsOrName, rule, kConst, unitL10nOr(kConst));
  }
}

//
// range in or out
//

class RangeError extends ValidateError {
  dynamic start;
  dynamic end;
  final bool isOut;
  final bool canEqStart;
  final bool canEqEnd;

  RangeError(ValidateErrorInfo info, this.start, this.end,
      {AtomV1 lenAtom,
      bool isItems,
      this.isOut = false,
      this.canEqStart = false,
      this.canEqEnd = false})
      : super(info, lenAtom: lenAtom, isItems: isItems);

  String get prefix => (isOut ? !canEqStart : canEqStart) ? '[' : '(';
  String get suffix => (isOut ? !canEqEnd : canEqEnd) ? ']' : ')';
  // example: [1, 30)
  String get range => "$prefix${format(start)}, ${format(end)}$suffix";

  @override
  String toString() {
    return tr.mustConst(fieldLenOrItemsOrName, isOut ? tr.rangex : tr.range,
        range, isLen ? lenOther : unitOtherOr);
  }
}

//
// within now, gt or lt
//

class WithinNowError extends ValidateError {
  final bool isGtNow;
  final Duration kConst;

  WithinNowError(ValidateErrorInfo info, this.isGtNow, this.kConst)
      : super(info);

  @override
  String toString() {
    return isGtNow == null
        ? tr.mustWithinNow
        : (isGtNow ? tr.mustWithinGtNow : tr.mustWithinLtNow)(
            fieldLenOrItemsOrName, format(kConst));
  }
}
