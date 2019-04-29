import 'package:protobuf/protobuf.dart' show GeneratedMessage;

import '../test/units_test.pb.dart';
import '../test/units_test.pbenum.dart';

T defaultOf_config<T extends GeneratedMessage>(Type type) {
  switch (type) {
    case MyProfile:
      return MyProfile.getDefault() as T;
    // default: throw ArgumentError.value(T);
    default:
      return null;
  }
}
