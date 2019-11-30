import 'package:protobuf/protobuf.dart';

abstract class ZeroAction<T extends GeneratedMessage> {
  void onZeroCreate(T proto);
  void onZeroLoad(T proto);
  void onZeroSave(T proto, [T orig]);
}
