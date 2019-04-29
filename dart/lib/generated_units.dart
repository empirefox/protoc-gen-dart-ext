import 'package:protobuf/protobuf.dart' as pb show FieldInfo, GeneratedMessage;

import './units/units.pb.dart';

class MessageInfo {
  Type type;
  pb.GeneratedMessage Function() getDefault;

  List<Unit> byIndex = [];
  Map<int, Unit> fieldUnit;
  Map<String, Unit> byTagAsString;
  Map<String, Unit> byName;

  void _addUnit(Unit fi) {
    byIndex.add(fi);
    assert(byIndex[fi.index] == fi);
    fieldInfo[fi.tagNumber] = fi;
    byTagAsString["${fi.tagNumber}"] = fi;
    byName[fi.name] = fi;
  }
}

class FieldInfo {
  pb.FieldInfo protobuf;
  Unit unit;
}
