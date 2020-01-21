import 'package:protobuf/protobuf.dart';

abstract class FormData<E> {
  final FormData parent;

  // tag number from Message
  final int tag;

  // index from repeated list
  final int index;
  // insert mode for repeated list
  final bool insert;

  // key from map
  final dynamic key;

  final E extra;

  FormData(
      FormData parent, this.tag, this.index, this.insert, this.key, this.extra)
      : this.parent = parent,
        _root = parent?._root;

  FormData.root(
      FormData root, this.tag, this.index, this.insert, this.key, this.extra)
      : this.parent = null,
        _root = root;

  final FormData _root;

  FormData get root => _root ?? this;

  bool get isRoot => _root == null;
}

class FormMessageData<E> extends FormData<E> {
  final GeneratedMessage saved;

  GeneratedMessage _editing;
  GeneratedMessage get editing {
    if (_editing == null) _editing = saved.clone();
    return _editing;
  }

  FormMessageData(FormData parent, int tag, int index, bool insert, dynamic key,
      E extra, this.saved, GeneratedMessage editing)
      : this._editing = editing,
        super(parent, tag, index, insert, key, extra);

  FormData childMessage(int tag) => FormMessageData(this, tag, null, false,
      null, extra, saved.getField(tag), editing.getField(tag));

  FormData childMessageList(int tag) => FormMessageListData(this, tag, null,
      false, null, extra, saved.getField(tag), editing.getField(tag));

  FormData childMessageMap(int tag) => FormMessageMapData(this, tag, null,
      false, null, extra, saved.getField(tag), editing.getField(tag));

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

class FormMessageListData<E> extends FormData<E> {
  final List saved;
  final List editing;

  FormMessageListData(FormData parent, int tag, int index, bool insert,
      dynamic key, E extra, this.saved, this.editing)
      : super(parent, tag, index, false, key, extra);

  FormData childMessage(int idx, bool insert) => FormMessageData(
      this, null, idx, insert, null, extra, saved[idx], editing[idx]);
}

class FormMessageMapData<E> extends FormData<E> {
  final Map saved;
  final Map editing;

  FormMessageMapData(FormData parent, int tag, int index, bool insert,
      dynamic key, E extra, this.saved, this.editing)
      : super(parent, tag, index, false, key, extra);

  FormData childMessage(dynamic key) => FormMessageData(
      this, null, null, false, key, extra, saved[key], editing[key]);
}
