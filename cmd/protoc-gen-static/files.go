package main

import (
	"github.com/empirefox/protoc-gen-dart-ext/pkg/genshared"
)

var files = map[string][]byte{
	"google/protobuf/timestamp.fmt.dart": []byte(genshared.NoEditCommentHead + `
import 'package:pgde/pgde.dart' show BeTime;

import 'timestamp.pb.dart';

const BeTime<Timestamp, DateTime> beTime = _BeTime._();

class _BeTime extends BeTime<Timestamp, DateTime> {
  const _BeTime._();
  @override
  Timestamp ofTime(DateTime v) => Timestamp.fromDateTime(v);
  @override
  DateTime toTime(Timestamp v) => v.toDateTime();
}
`),

	"google/protobuf/duration.fmt.dart": []byte(genshared.NoEditCommentHead + `
import 'package:fixnum/fixnum.dart' show Int64;
import 'package:pgde/pgde.dart' show BeTime;

import 'duration.pb.dart' as pb;

const BeTime<pb.Duration, Duration> beTime = _BeTime._();

class _BeTime extends BeTime<pb.Duration, Duration> {
  const _BeTime._();
  @override
  pb.Duration ofTime(Duration v) => pb.Duration()
    ..seconds = Int64(v.inSeconds)
    ..nanos = (v.inMicroseconds % Duration.microsecondsPerSecond) * 1000;

  @override
  Duration toTime(pb.Duration v) =>
      Duration(seconds: v.seconds.toInt(), microseconds: v.nanos ~/ 1000);
}
`),
}
