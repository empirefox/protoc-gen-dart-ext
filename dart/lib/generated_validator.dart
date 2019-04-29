import 'package:intl/intl.dart';
import 'package:protobuf/protobuf.dart';

import './rules.dart';
import './translate.dart';

/// The base interface for all generated PGV validators.
/// Asserts validation rules on a protobuf object.
/// Throw [ValidationError] with the first validation error encountered.
abstract class GeneratedValidator<T extends GeneratedMessage> {
  void assertProto(T proto);
  void assertField(T proto, int tag);
  void assertOneof(T proto, int oneof);
}

abstract class TranslateContext {
  String enumName(Type pbEnumType);
  String enumValue(ProtobufEnum value);
  String protoName(Type msgType);
  String oneofName(Type msgType, int oneof);
  String fieldName(Type msgType, int tagNumber);
  String boolFieldValue(Type msgType, int tagNumber);
}

abstract class Printer {
  String print(dynamic kConst);
}

class DefaultPrinter extends Printer {
  final DateFormat dateFormat;
  // TODO add number format per number type.
  final NumberFormat numberFormat;
  DefaultPrinter(this.dateFormat, this.numberFormat);
  String print(dynamic kConst) {
    // TODO use package:duration?
    if (kConst is Duration) return kConst.toString();
    if (kConst is DateTime) return dateFormat.format(kConst);
    return kConst is String ? kConst : '$kConst';
  }
}

Printer printer = DefaultPrinter(DateFormat("yyyy.MM.dd G ',' HH:mm:ss vvvv"));

abstract class ValidateError extends Error {
  final GeneratedMessage proto;
  final int tagNumber;
  final Rule rule;

  final dynamic value;
  ValidateError(this.proto, this.tagNumber, this.rule)
      : value = proto.getFieldOrNull(tagNumber);

  BuilderInfo get bi => proto.info_;
  FieldInfo get fi => bi.fieldInfo[tagNumber];
  Type get pbType => proto.runtimeType;
  String get messageName => bi.messageName;
  String get fieldAccessor => '${messageName}.${bi.fieldName(tagNumber)}';
  String fieldName(TranslateContext ctx) => ctx.fieldName(pbType, tagNumber);

  String translate(Translator tr, TranslateContext ctx);
}

//
// required
//

class RequiredError extends ValidateError {
  RequiredError(GeneratedMessage proto, int tagNumber)
      : super(proto, tagNumber, Rule.required);
  @override
  String toString() => '$fieldAccessor is required';
  @override
  String translate(Translator tr, TranslateContext ctx) =>
      tr.required(fieldName(ctx));
}

class OneofRequiredError extends RequiredError {
  final int oneof;
  OneofRequiredError(GeneratedMessage proto, this.oneof) : super(proto, null);
  @override
  String fieldName(TranslateContext ctx) => ctx.oneofName(pbType, oneof);
  @override
  String toString() => '${messageName}.oneof[${oneof}]  is required';
}

//
// eq_const
//

class ConstError extends ValidateError {
  final dynamic kConst;
  // for bytes
  final String kConstPrint;
  // for string_len string_len_bytes bytes_len
  final NumName numName;

  ConstError(GeneratedMessage proto, int tagNumber, this.kConst,
      {this.kConstPrint = '', this.numName = NumName.none})
      : super(proto, tagNumber, Rule.eq_const);

  @override
  String toString() =>
      '$fieldAccessor ${numName == NumName.none ? '' : (numNameString(numName) + ' ')}must equal $kConst';

  @override
  String translate(Translator tr, TranslateContext ctx) {
    if (kConstPrint.isNotEmpty) return tr.eq_const(fieldName(ctx), kConstPrint);
    if (kConst is bool) return ctx.boolFieldValue(pbType, tagNumber);
    if (kConst is ProtobufEnum)
      return tr.eq_const(fieldName(ctx), ctx.enumValue(kConst));
    return tr.eq_const(fieldName(ctx), printer.print(kConst));
  }
}

//
// in_or_not
//

class InError extends ValidateError {
  final List kConst;
  final List<String> kConstPrint;
  final bool isNot;

  InError(GeneratedMessage proto, int tagNumber, this.kConst,
      {this.kConstPrint = const [], this.isNot = false})
      : super(proto, tagNumber, Rule.in_or_not);

  InError.not(GeneratedMessage proto, int tagNumber, List kConst,
      {List<String> kConstPrint})
      : this(proto, tagNumber, kConst, kConstPrint: kConstPrint, isNot: true);

  @override
  String toString() =>
      '$fieldAccessor must be${isNot ? " not" : ""} one of $kConst';

  @override
  String translate(Translator tr, TranslateContext ctx) {
    var kConst = this.kConst;

    if (kConstPrint.isNotEmpty)
      kConst = kConstPrint;
    else if (kConst is List<ProtobufEnum>)
      kConst = kConst.map((pbe) => ctx.enumValue(pbe)).toList();
    else
      kConst = kConst.map((v) => printer.print(v)).toList();

    return (isNot ? tr.not_list : tr.in_list)(fieldName(ctx), kConst);
  }
}

//
// range Overflow of uint or int32?
//

class RangeError extends ValidateError {
  dynamic start;
  dynamic end;
  final bool isOut;
  final bool canEqStart;
  final bool canEqEnd;
  final NumName numName;
  RangeError(GeneratedMessage proto, int tagNumber, this.start, this.end,
      {this.isOut = false,
      this.canEqStart = false,
      this.canEqEnd = false,
      this.numName = NumName.none})
      : super(proto, tagNumber, Rule.range);

  String get prefix => (isOut ? !canEqStart : canEqStart) ? '[' : '(';
  String get suffix => (isOut ? !canEqEnd : canEqEnd) ? ']' : ')';
  // example: [1, 30)
  String get range =>
      "${prefix}${printer.print(start)}, ${printer.print(end)}${suffix}";

  @override
  String toString() =>
      "$fieldAccessor must be ${isOut ? 'out' : 'in'}side range $range";

  @override
  String translate(Translator tr, TranslateContext ctx) {
    final numName = tr.numNameOf(this.numName);
    if (numName.isEmpty)
      return (isOut ? tr.range_out : tr.range_in)(fieldName(ctx), range);
    return (isOut ? tr.range_out_numName : tr.range_in_numName)(
        fieldName(ctx), range, numName);
  }
}
