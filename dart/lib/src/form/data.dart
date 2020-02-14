import 'package:path/path.dart' show posix;
import 'package:protobuf/protobuf.dart';

abstract class FormData {
  final FormData parent;

  // tag number from Message
  final int tag;

  // index from repeated list
  final int index;
  // insert mode for repeated list
  final bool insert;

  // key from map
  final dynamic key;

  final String route;

  final dynamic extra;

  FormData(
      {this.route,
      FormData parent,
      this.tag,
      this.index,
      this.insert,
      this.key,
      this.extra})
      : this.parent = parent,
        _root = parent?._root;

  final FormData _root;

  FormData get root => _root ?? this;

  bool get isRoot => _root == null;

  String childRoute(int i) => posix.join(route, i.toString());
}

class FormMessageData extends FormData {
  final GeneratedMessage saved;

  GeneratedMessage _editing;
  GeneratedMessage get editing {
    if (_editing == null) _editing = saved.clone();
    return _editing;
  }

  FormMessageData(
      {String route,
      FormData parent,
      int tag,
      int index,
      bool insert,
      dynamic key,
      dynamic extra,
      this.saved,
      GeneratedMessage editing})
      : this._editing = editing,
        super(
            route: route,
            parent: parent,
            tag: tag,
            index: index,
            insert: insert,
            key: key,
            extra: extra);

  FormMessageData.root(
      String route, this.saved, GeneratedMessage editing, dynamic extra)
      : this._editing = editing,
        super(
            route: route,
            parent: null,
            tag: null,
            index: null,
            insert: false,
            key: null,
            extra: extra);

  FormData childMessage(int tag) => FormMessageData(
      route: childRoute(tag),
      parent: this,
      tag: tag,
      index: null,
      insert: false,
      key: null,
      extra: extra,
      saved: saved.getField(tag),
      editing: editing.getField(tag));

  FormData childMessageList(int tag) => FormMessageListData(
      route: childRoute(tag),
      parent: this,
      tag: tag,
      index: null,
      insert: false,
      key: null,
      extra: extra,
      saved: saved.getField(tag),
      editing: editing.getField(tag));

  FormData childMessageMap(int tag) => FormMessageMapData(
      route: childRoute(tag),
      parent: this,
      tag: tag,
      index: null,
      insert: false,
      key: null,
      extra: extra,
      saved: saved.getField(tag),
      editing: editing.getField(tag));

  void save(GeneratedMessage draft) {
    if (isRoot) return;
    final p = parent;
    if (p is FormMessageData)
      p.editing.setField(tag, draft);
    else if (p is FormMessageListData) {
      if (p.insert)
        p.editing.insert(index, draft);
      else
        p.editing[index] = draft;
    } else if (p is FormMessageMapData)
      p.editing[key] = draft;
    else
      assert(false, 'save draft data, parent type error: ${p.runtimeType}');
  }
}

class FormMessageListData extends FormData {
  final List saved;
  final List editing;

  FormMessageListData(
      {String route,
      FormData parent,
      int tag,
      int index,
      bool insert,
      dynamic key,
      dynamic extra,
      this.saved,
      this.editing})
      : super(
            route: route,
            parent: parent,
            tag: tag,
            index: index,
            insert: false,
            key: key,
            extra: extra);

  FormData childMessage(int idx, bool insert) => FormMessageData(
      route: childRoute(idx),
      parent: this,
      tag: null,
      index: idx,
      insert: insert,
      key: null,
      extra: extra,
      saved: saved[idx],
      editing: editing[idx]);
}

class FormMessageMapData extends FormData {
  final Map saved;
  final Map editing;

  FormMessageMapData(
      {String route,
      FormData parent,
      int tag,
      int index,
      bool insert,
      dynamic key,
      dynamic extra,
      this.saved,
      this.editing})
      : super(
            route: route,
            parent: parent,
            tag: tag,
            index: index,
            insert: false,
            key: key,
            extra: extra);

  FormData childMessage(dynamic key) => FormMessageData(
      route: key is String ? posix.join(route, 's', key) : childRoute(key),
      parent: this,
      tag: null,
      index: null,
      insert: false,
      key: key,
      extra: extra,
      saved: saved[key],
      editing: editing[key]);
}
