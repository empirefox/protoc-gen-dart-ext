import 'dart:io';

import '../protos/exports/exports.pb.dart';

final _exports = Exports();
var _hasExports = false;

void registerPkgBytes(String importPath, List<int> pkg) {
  assert(!_exports.packages.containsKey(importPath));
  _exports.packages[importPath] = Package.fromBuffer(pkg);
  _hasExports = true;
}

void registerPkg(String importPath, Package pkg) {
  assert(!_exports.packages.containsKey(importPath));
  _exports.packages[importPath] = pkg;
  _hasExports = true;
}

void writeToFile(String path) {
  if (!_hasExports) return;
  final contents = _exports.writeToBuffer();
  File(path).writeAsBytesSync(contents, flush: true);
}
