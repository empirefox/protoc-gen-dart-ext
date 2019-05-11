import 'package:flutter/material.dart' show MaterialLocalizations, TimeOfDay;

abstract class TimeFormatter<F, T> {
  final BeTime<F, T> bt;
  final bool alwaysUse24HourFormat;

  const TimeFormatter(this.bt, {this.alwaysUse24HourFormat});

  String format(MaterialLocalizations md, F v) =>
      formatTime(md, bt.toTime(v), alwaysUse24HourFormat);

  String formatTime(MaterialLocalizations md, T v, bool alwaysUse24HourFormat);
}

class OfDayFormatter<F> extends TimeFormatter<F, DateTime> {
  const OfDayFormatter(BeTime<F, DateTime> bt, {bool alwaysUse24HourFormat})
      : super(bt, alwaysUse24HourFormat: alwaysUse24HourFormat);

  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) =>
      md.formatTimeOfDay(TimeOfDay.fromDateTime(v),
          alwaysUse24HourFormat: alwaysUse24HourFormat);
}

class MediumDateFormatter<F> extends TimeFormatter<F, DateTime> {
  const MediumDateFormatter(BeTime<F, DateTime> bt) : super(bt);
  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) =>
      md.formatMediumDate(v);
}

class FullDateFormatter<F> extends TimeFormatter<F, DateTime> {
  const FullDateFormatter(BeTime<F, DateTime> bt) : super(bt);
  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) => md.formatFullDate(v);
}

class MediumDatetimeFormatter<F> extends TimeFormatter<F, DateTime> {
  const MediumDatetimeFormatter(BeTime<F, DateTime> bt,
      {bool alwaysUse24HourFormat})
      : super(bt, alwaysUse24HourFormat: alwaysUse24HourFormat);
  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) =>
      '${md.formatMediumDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v), alwaysUse24HourFormat: alwaysUse24HourFormat)}';
}

class FullDatetimeFormatter<F> extends TimeFormatter<F, DateTime> {
  const FullDatetimeFormatter(BeTime<F, DateTime> bt,
      {bool alwaysUse24HourFormat})
      : super(bt, alwaysUse24HourFormat: alwaysUse24HourFormat);
  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) =>
      '${md.formatFullDate(v)} ${md.formatTimeOfDay(TimeOfDay.fromDateTime(v), alwaysUse24HourFormat: alwaysUse24HourFormat)}';
}

class DurationFormatter<F> extends TimeFormatter<F, Duration> {
  const DurationFormatter(BeTime<F, Duration> bt) : super(bt);

  // TODO, fix this: the duration format not implemented
  @override
  String formatTime(md, v, bool alwaysUse24HourFormat) => v.toString();
}

abstract class BeTime<F, T> {
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
  static const BeTime<int, DateTime> dayTime =
      const _MsPerTime(Duration.millisecondsPerDay);
  static const BeTime<int, DateTime> weekTime =
      const _MsPerTime(millisecondsPerWeek);

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

  static const BeTime<Duration, Duration> duration =
      const NoActionBeTime<Duration>();

  static const BeTime<int, Duration> nanosecondDur = const _NanosecondDur();
  static const BeTime<int, Duration> microsecondDur = const _MicrosecondDur();
  static const BeTime<int, Duration> millisecondDur = const _MillisecondDur();
  static const BeTime<int, Duration> secondDur = const _SecondDur();
  static const BeTime<int, Duration> minuteDur = const _MinuteDur();
  static const BeTime<int, Duration> hourDur = const _HourDur();
  static const BeTime<int, Duration> dayDur = const _DayDur();
  static const BeTime<int, Duration> weekDur = const _WeekDur();

  static const int millisecondsPerWeek = Duration.millisecondsPerDay * 7;

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

class _DayDur implements BeTime<int, Duration> {
  const _DayDur();
  @override
  Duration toTime(int v) => Duration(days: v);
  @override
  int ofTime(Duration v) => v.inDays;
}

class _WeekDur implements BeTime<int, Duration> {
  const _WeekDur();
  @override
  Duration toTime(int v) => Duration(days: v * 7);
  @override
  int ofTime(Duration v) => v.inDays ~/ 7;
}
