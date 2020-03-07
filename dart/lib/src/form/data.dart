import 'package:path/path.dart' show posix;
import 'package:protobuf/protobuf.dart' show GeneratedMessage;

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

  bool deleted = false;

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

  final Map<int, FormMessageData> cacheByTag = {};
  final Map<int, FormMessageListData> listCacheByTag = {};
  final Map<int, FormMessageMapData> mapCacheByTag = {};

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

  void saveSelf(GeneratedMessage draft) {
    if (isRoot) return;
    final p = parent;
    if (p is FormMessageData)
      p.editing.setField(tag, draft);
    else if (p is FormMessageListData)
      p.saveChild(index, draft);
    else if (p is FormMessageMapData)
      p.saveChild(key, draft);
    else
      assert(false, 'save draft data, parent type error: ${p.runtimeType}');
  }
}

class FormMessageListData extends FormData {
  final List saved;
  final List editing;

  final List<FormMessageData> cache = [];

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

  void saveChild(int idx, GeneratedMessage draft) {
    if (insert)
      editing.insert(index, draft);
    else
      editing[idx] = draft.clone();
  }

  void exchangeChildMessage(int idx1, int idx2, int idTag) {
    final t1 = editing[idx1] as GeneratedMessage;
    final t2 = editing[idx2] as GeneratedMessage;
    final id1 = t1.getField(idTag);
    final id2 = t2.getField(idTag);
    editing[idx1] = t2..setField(idTag, id1);
    editing[2] = t1..setField(idTag, id2);
  }

  void exchangeChild(int idx1, int idx2) {
    final t1 = editing[idx1];
    final t2 = editing[idx2];
    editing[idx1] = t2;
    editing[2] = t1;
  }

  void removeChild(int idx) {
    editing.removeAt(idx);
  }
}

class FormMessageMapData extends FormData {
  final Map saved;
  final Map editing;

  final Map<dynamic, FormMessageData> cache = {};

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

  FormData childMessage(dynamic key) => cache.putIfAbsent(
      key,
      () => FormMessageData(
          route: key is String ? posix.join(route, 's', key) : childRoute(key),
          parent: this,
          tag: null,
          index: null,
          insert: false,
          key: key,
          extra: extra,
          saved: saved[key],
          editing: editing[key]));

  void changeKey(dynamic key, dynamic newKey) {
    editing[newKey] = editing.remove(key);
  }

  void saveChild(dynamic key, GeneratedMessage draft) {
    editing[key] = draft.clone();
  }

  void removeChild(dynamic key) {
    editing.remove(key);
    cache.remove(key);
  }
}
