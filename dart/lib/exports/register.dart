import 'dart:io';

import '../protos/exports/exports.pb.dart';

final _exports = Exports();

void registerPkgBytes(String importPath, List<int> pkg) {
  assert(!_exports.packages.containsKey(importPath));
  _exports.packages[importPath] = Package.fromBuffer(pkg);
}

void registerPkg(String importPath, Package pkg) {
  assert(!_exports.packages.containsKey(importPath));
  _exports.packages[importPath] = pkg;
}

void writeToFile(String path) {
  final contents = _exports.writeToBuffer();
  File(path).writeAsBytesSync(contents, flush: true);
}
