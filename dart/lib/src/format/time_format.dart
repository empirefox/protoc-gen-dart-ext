import 'package:fixnum/fixnum.dart' show Int64;
import 'package:flutter/material.dart' show MaterialLocalizations, TimeOfDay;

import 'formatter.dart';

abstract class TimeFormatter<F, T> implements Formatter<F> {
  final BeTime<F, T> bt;

  const TimeFormatter(this.bt) : assert(bt != null);

  String format(v, md, l, {form, show}) => formatTime(md, bt.toTime(v));
  String formatTime(MaterialLocalizations md, T v);
}

class OfDayFormatter<F> extends TimeFormatter<F, DateTime> {
  const OfDayFormatter(BeTime<F, DateTime> bt) : super(bt);

  @override
  String formatTime(md, v) => md.formatTimeOfDay(TimeOfDay.fromDateTime(v));
}

class OfDay24HFormatter<F> extends TimeFormatter<F, DateTime> {
  const OfDay24HFormatter(BeTime<F, DateTime> bt) : super(bt);

  @override
  String formatTime(md, v) => md.formatTimeOfDay(TimeOfDay.fromDateTime(v),
      alwaysUse24HourFormat: true);
}

class MediumDateFormatter<F> extends TimeFormatter<F, DateTime> {
  const MediumDateFormatter(BeTime<F, DateTime> bt) : super(bt);
  @override
  String formatTime(md, v) => md.formatMediumDate(v);
}

class FullDateFormatter<F> extends TimeFormatter<F, DateTime> {
  const FullDateFormatter(BeTime<F, DateTime> bt) : super(bt);
  @override
  String formatTime(md, v) => md.formatFullDate(v);
}

class MediumDatetimeFormatter<F> extends TimeFormatter<F, DateTime> {
  const MediumDatetimeFormatter(
    BeTime<F, DateTime> bt,
  ) : super(bt);
  @override
  String formatTime(md, v) =>
      '${md.formatMediumDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v))}';
}

class MediumDatetime24HFormatter<F> extends TimeFormatter<F, DateTime> {
  const MediumDatetime24HFormatter(
    BeTime<F, DateTime> bt,
  ) : super(bt);
  @override
  String formatTime(md, v) =>
      '${md.formatMediumDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v), alwaysUse24HourFormat: true)}';
}

class FullDatetimeFormatter<F> extends TimeFormatter<F, DateTime> {
  const FullDatetimeFormatter(
    BeTime<F, DateTime> bt,
  ) : super(bt);
  @override
  String formatTime(md, v) =>
      '${md.formatFullDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v))}';
}

class FullDatetime24HFormatter<F> extends TimeFormatter<F, DateTime> {
  const FullDatetime24HFormatter(
    BeTime<F, DateTime> bt,
  ) : super(bt);
  @override
  String formatTime(md, v) =>
      '${md.formatFullDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v), alwaysUse24HourFormat: true)}';
}

class DurationFormatter<F> extends TimeFormatter<F, Duration> {
  const DurationFormatter(BeTime<F, Duration> bt) : super(bt);

  // TODO, fix this: the duration format not implemented
  @override
  String formatTime(md, v) => v.toString();
}

abstract class BeTime<F, T> {
  // Time
  static const BeTime<DateTime, DateTime> time =
      const NoActionBeTime<DateTime>();

  static const BeTime<int, DateTime> nanosecondTime = const _NanosecondTime();
  static const BeTime<int, DateTime> microsecondTime = const _MicrosecondTime();
  static const BeTime<int, DateTime> millisecondTime = const _MillisecondTime();
  static const BeTime<int, DateTime> secondTime =
      const _MsPerTime(Duration.millisecondsPerSecond);
  static const BeTime<int, DateTime> minuteTime =
      const _MsPerTime(Duration.millisecondsPerMinute);
  static const BeTime<int, DateTime> hourTime =
      const _MsPerTime(Duration.millisecondsPerHour);

  static const BeTime<Int64, DateTime> nanosecondTime64 =
      const _Int64Time(nanosecondTime);
  static const BeTime<Int64, DateTime> microsecondTime64 =
      const _Int64Time(microsecondTime);
  static const BeTime<Int64, DateTime> millisecondTime64 =
      const _Int64Time(millisecondTime);
  static const BeTime<Int64, DateTime> secondTime64 =
      const _Int64Time(secondTime);
  static const BeTime<Int64, DateTime> minuteTime64 =
      const _Int64Time(minuteTime);
  static const BeTime<Int64, DateTime> hourTime64 = const _Int64Time(hourTime);

  // OfDay
  static const BeTime<DateTime, DateTime> ofDay = const _TimeOfDay();

  static const BeTime<int, DateTime> nanosecondOfDay =
      const _NanosecondTimeOfDay();
  static const BeTime<int, DateTime> microsecondOfDay =
      const _MicrosecondTimeOfDay();
  static const BeTime<int, DateTime> millisecondOfDay =
      const _MillisecondTimeOfDay();
  static const BeTime<int, DateTime> secondOfDay =
      const _MsPerTimeOfDay(Duration.millisecondsPerSecond);
  static const BeTime<int, DateTime> minuteOfDay =
      const _MsPerTime(Duration.millisecondsPerMinute);
  static const BeTime<int, DateTime> hourOfDay =
      const _MsPerTimeOfDay(Duration.millisecondsPerHour);

  static const BeTime<Int64, DateTime> nanosecondOfDay64 =
      const _Int64Time(nanosecondOfDay);
  static const BeTime<Int64, DateTime> microsecondOfDay64 =
      const _Int64Time(microsecondOfDay);
  static const BeTime<Int64, DateTime> millisecondOfDay64 =
      const _Int64Time(millisecondOfDay);
  static const BeTime<Int64, DateTime> secondOfDay64 =
      const _Int64Time(secondOfDay);
  static const BeTime<Int64, DateTime> minuteOfDay64 =
      const _Int64Time(minuteOfDay);
  static const BeTime<Int64, DateTime> hourOfDay64 =
      const _Int64Time(hourOfDay);

  // Dur
  static const BeTime<Duration, Duration> duration =
      const NoActionBeTime<Duration>();

  static const BeTime<int, Duration> nanosecondDur = const _NanosecondDur();
  static const BeTime<int, Duration> microsecondDur = const _MicrosecondDur();
  static const BeTime<int, Duration> millisecondDur = const _MillisecondDur();
  static const BeTime<int, Duration> secondDur = const _SecondDur();
  static const BeTime<int, Duration> minuteDur = const _MinuteDur();
  static const BeTime<int, Duration> hourDur = const _HourDur();

  static const BeTime<Int64, Duration> nanosecondDur64 =
      const _Int64Time(nanosecondDur);
  static const BeTime<Int64, Duration> microsecondDur64 =
      const _Int64Time(microsecondDur);
  static const BeTime<Int64, Duration> millisecondDur64 =
      const _Int64Time(millisecondDur);
  static const BeTime<Int64, Duration> secondDur64 =
      const _Int64Time(secondDur);
  static const BeTime<Int64, Duration> minuteDur64 =
      const _Int64Time(minuteDur);
  static const BeTime<Int64, Duration> hourDur64 = const _Int64Time(hourDur);

  const BeTime();
  T toTime(F v);
  F ofTime(T v);
}

mixin _NoActionMixin<T> {
  T toTime(T v) => v;
  T ofTime(T v) => v;
}

class NoActionBeTime<T> with _NoActionMixin<T> implements BeTime<T, T> {
  const NoActionBeTime();
}

class _Int64Time<T> implements BeTime<Int64, T> {
  final BeTime<int, T> bt;
  const _Int64Time(this.bt);
  @override
  T toTime(Int64 v) => bt.toTime(v.toInt());
  @override
  Int64 ofTime(T v) => Int64(bt.ofTime(v));
}

class AnyIntTime<T> implements BeTime<T, DateTime> {
  final BeTime<T, int> beInt;
  final BeTime<int, DateTime> beTime;
  const AnyIntTime(this.beInt, this.beTime);
  @override
  DateTime toTime(T v) => beTime.toTime(beInt.toTime(v));
  @override
  T ofTime(DateTime v) => beInt.ofTime(beTime.ofTime(v));
}

// DateTime

class _NanosecondTime implements BeTime<int, DateTime> {
  const _NanosecondTime();
  @override
  DateTime toTime(int v) =>
      DateTime.fromMicrosecondsSinceEpoch(v ~/ 1000, isUtc: true);
  @override
  int ofTime(DateTime v) => v.microsecondsSinceEpoch * 1000;
}

class _MicrosecondTime implements BeTime<int, DateTime> {
  const _MicrosecondTime();
  @override
  DateTime toTime(int v) => DateTime.fromMicrosecondsSinceEpoch(v, isUtc: true);
  @override
  int ofTime(DateTime v) => v.microsecondsSinceEpoch;
}

class _MillisecondTime implements BeTime<int, DateTime> {
  const _MillisecondTime();
  @override
  DateTime toTime(int v) => DateTime.fromMillisecondsSinceEpoch(v, isUtc: true);
  @override
  int ofTime(DateTime v) => v.millisecondsSinceEpoch;
}

class _MsPerTime implements BeTime<int, DateTime> {
  final int _msPer;
  const _MsPerTime(this._msPer);
  @override
  DateTime toTime(int v) =>
      DateTime.fromMillisecondsSinceEpoch(v * _msPer, isUtc: true);
  @override
  int ofTime(DateTime v) => v.millisecondsSinceEpoch ~/ _msPer;
}

// TimeOfDay

class _TimeOfDay implements BeTime<DateTime, DateTime> {
  const _TimeOfDay();
  @override
  DateTime toTime(DateTime v) => _truncate(v);
  @override
  DateTime ofTime(DateTime v) => _truncate(v);
  DateTime _truncate(DateTime v) => DateTime.fromMillisecondsSinceEpoch(
      v.millisecondsSinceEpoch % Duration.millisecondsPerDay);
}

class _NanosecondTimeOfDay extends _NanosecondTime {
  const _NanosecondTimeOfDay();
  @override
  int ofTime(DateTime v) =>
      v.microsecondsSinceEpoch % (Duration.microsecondsPerDay ~/ 1000);
}

class _MicrosecondTimeOfDay extends _MicrosecondTime {
  const _MicrosecondTimeOfDay();
  @override
  int ofTime(DateTime v) =>
      v.microsecondsSinceEpoch % Duration.microsecondsPerDay;
}

class _MillisecondTimeOfDay extends _MillisecondTime {
  const _MillisecondTimeOfDay();
  @override
  int ofTime(DateTime v) =>
      v.millisecondsSinceEpoch % Duration.millisecondsPerDay;
}

class _MsPerTimeOfDay implements BeTime<int, DateTime> {
  final int _msPer;
  const _MsPerTimeOfDay(this._msPer);
  @override
  DateTime toTime(int v) =>
      DateTime.fromMillisecondsSinceEpoch(v * _msPer, isUtc: true);
  @override
  int ofTime(DateTime v) =>
      v.millisecondsSinceEpoch % (Duration.millisecondsPerDay * _msPer);
}

// Duration

class _NanosecondDur implements BeTime<int, Duration> {
  const _NanosecondDur();
  @override
  Duration toTime(int v) => Duration(microseconds: v ~/ 1000);
  @override
  int ofTime(Duration v) => v.inMicroseconds * 1000;
}

class _MicrosecondDur implements BeTime<int, Duration> {
  const _MicrosecondDur();
  @override
  Duration toTime(int v) => Duration(microseconds: v);
  @override
  int ofTime(Duration v) => v.inMicroseconds;
}

class _MillisecondDur implements BeTime<int, Duration> {
  const _MillisecondDur();
  @override
  Duration toTime(int v) => Duration(milliseconds: v);
  @override
  int ofTime(Duration v) => v.inMilliseconds;
}

class _SecondDur implements BeTime<int, Duration> {
  const _SecondDur();
  @override
  Duration toTime(int v) => Duration(seconds: v);
  @override
  int ofTime(Duration v) => v.inSeconds;
}

class _MinuteDur implements BeTime<int, Duration> {
  const _MinuteDur();
  @override
  Duration toTime(int v) => Duration(minutes: v);
  @override
  int ofTime(Duration v) => v.inMinutes;
}

class _HourDur implements BeTime<int, Duration> {
  const _HourDur();
  @override
  Duration toTime(int v) => Duration(hours: v);
  @override
  int ofTime(Duration v) => v.inHours;
}
