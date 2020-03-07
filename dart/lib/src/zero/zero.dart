import 'package:protobuf/protobuf.dart' show GeneratedMessage;

abstract class ZeroAction<T extends GeneratedMessage> {
  T create();
  T onZeroCreate(T proto);
  T onZeroLoad(T proto);
  T onZeroSave(T proto, [T orig]);
}
