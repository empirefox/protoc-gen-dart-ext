import 'dart:io';
import 'dart:mirrors'; // ignore: uri_does_not_exist

import 'package:path/path.dart' as pathlib;

import '../protos/exports/exports.pb.dart';

typedef RegisterFunc = void Function(String pkgUri);

final _exports = Exports();
var _hasExports = false;

void registerPkgBytes(String pkgUri, List<int> pkg) {
  final p = Package.fromBuffer(pkg);
  if (!p.path.startsWith('package:')) p.path = pathlib.join(pkgUri, p.path);
  assert(_exports.packages.every((item) => item.path != p.path));
  _exports.packages.add(p);
  _hasExports = true;
}

void registerValidatorBytes(String pkgUri, List<int> validator) {
  final v = Validator.fromBuffer(validator);
  if (!v.path.startsWith('package:')) v.path = pathlib.join(pkgUri, v.path);
  assert(_exports.validators.every((item) => item.path != v.path));
  _exports.validators.add(v);
  _hasExports = true;
}

void writeToFile(String path) {
  if (!_hasExports) return;
  final contents = _exports.writeToBuffer();
  File(path).writeAsBytesSync(contents, flush: true);
}

void invoke(RegisterFunc fn) {
  fn(pathlib.dirname(_packageUri(fn).toString()));
}

// fn must be a generated RegisterFunc in exports.go
Uri _packageUri(RegisterFunc fn) {
  final instance = reflect(fn); // ignore: undefined_function
  // ignore: type_test_with_undefined_name
  if (instance is ClosureMirror) {
    final owner = instance.function.owner;
    // ignore: type_test_with_undefined_name
    if (owner is LibraryMirror) {
      return owner.uri;
    }
  }
  throw ArgumentError.value(fn);
}
