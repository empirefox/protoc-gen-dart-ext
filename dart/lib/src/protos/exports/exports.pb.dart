///
//  Generated code. Do not modify.
//  source: protos/exports/exports.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

class Field extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Field', package: const $pb.PackageName('pgde.exports'))
    ..aOS(1, 'ref')
    ..aOS(2, 'name')
    ..hasRequiredFields = false
  ;

  Field() : super();
  Field.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Field.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Field clone() => Field()..mergeFromMessage(this);
  Field copyWith(void Function(Field) updates) => super.copyWith((message) => updates(message as Field));
  $pb.BuilderInfo get info_ => _i;
  static Field create() => Field();
  Field createEmptyInstance() => create();
  static $pb.PbList<Field> createRepeated() => $pb.PbList<Field>();
  static Field getDefault() => _defaultInstance ??= create()..freeze();
  static Field _defaultInstance;

  $core.String get ref => $_getS(0, '');
  set ref($core.String v) { $_setString(0, v); }
  $core.bool hasRef() => $_has(0);
  void clearRef() => clearField(1);

  $core.String get name => $_getS(1, '');
  set name($core.String v) { $_setString(1, v); }
  $core.bool hasName() => $_has(1);
  void clearName() => clearField(2);
}

class Entity extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Entity', package: const $pb.PackageName('pgde.exports'))
    ..aOS(1, 'name')
    ..pc<Field>(2, 'fields', $pb.PbFieldType.PM,Field.create)
    ..hasRequiredFields = false
  ;

  Entity() : super();
  Entity.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Entity.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Entity clone() => Entity()..mergeFromMessage(this);
  Entity copyWith(void Function(Entity) updates) => super.copyWith((message) => updates(message as Entity));
  $pb.BuilderInfo get info_ => _i;
  static Entity create() => Entity();
  Entity createEmptyInstance() => create();
  static $pb.PbList<Entity> createRepeated() => $pb.PbList<Entity>();
  static Entity getDefault() => _defaultInstance ??= create()..freeze();
  static Entity _defaultInstance;

  $core.String get name => $_getS(0, '');
  set name($core.String v) { $_setString(0, v); }
  $core.bool hasName() => $_has(0);
  void clearName() => clearField(1);

  $core.List<Field> get fields => $_getList(1);
}

class Package extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Package', package: const $pb.PackageName('pgde.exports'))
    ..aOS(1, 'path')
    ..pc<Entity>(2, 'entities', $pb.PbFieldType.PM,Entity.create)
    ..hasRequiredFields = false
  ;

  Package() : super();
  Package.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Package.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Package clone() => Package()..mergeFromMessage(this);
  Package copyWith(void Function(Package) updates) => super.copyWith((message) => updates(message as Package));
  $pb.BuilderInfo get info_ => _i;
  static Package create() => Package();
  Package createEmptyInstance() => create();
  static $pb.PbList<Package> createRepeated() => $pb.PbList<Package>();
  static Package getDefault() => _defaultInstance ??= create()..freeze();
  static Package _defaultInstance;

  $core.String get path => $_getS(0, '');
  set path($core.String v) { $_setString(0, v); }
  $core.bool hasPath() => $_has(0);
  void clearPath() => clearField(1);

  $core.List<Entity> get entities => $_getList(1);
}

class Validator extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Validator', package: const $pb.PackageName('pgde.exports'))
    ..aOS(1, 'path')
    ..pPS(2, 'entities')
    ..hasRequiredFields = false
  ;

  Validator() : super();
  Validator.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Validator.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Validator clone() => Validator()..mergeFromMessage(this);
  Validator copyWith(void Function(Validator) updates) => super.copyWith((message) => updates(message as Validator));
  $pb.BuilderInfo get info_ => _i;
  static Validator create() => Validator();
  Validator createEmptyInstance() => create();
  static $pb.PbList<Validator> createRepeated() => $pb.PbList<Validator>();
  static Validator getDefault() => _defaultInstance ??= create()..freeze();
  static Validator _defaultInstance;

  $core.String get path => $_getS(0, '');
  set path($core.String v) { $_setString(0, v); }
  $core.bool hasPath() => $_has(0);
  void clearPath() => clearField(1);

  $core.List<$core.String> get entities => $_getList(1);
}

class Exports extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('Exports', package: const $pb.PackageName('pgde.exports'))
    ..pc<Package>(1, 'packages', $pb.PbFieldType.PM,Package.create)
    ..pc<Validator>(2, 'validators', $pb.PbFieldType.PM,Validator.create)
    ..hasRequiredFields = false
  ;

  Exports() : super();
  Exports.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  Exports.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  Exports clone() => Exports()..mergeFromMessage(this);
  Exports copyWith(void Function(Exports) updates) => super.copyWith((message) => updates(message as Exports));
  $pb.BuilderInfo get info_ => _i;
  static Exports create() => Exports();
  Exports createEmptyInstance() => create();
  static $pb.PbList<Exports> createRepeated() => $pb.PbList<Exports>();
  static Exports getDefault() => _defaultInstance ??= create()..freeze();
  static Exports _defaultInstance;

  $core.List<Package> get packages => $_getList(0);

  $core.List<Validator> get validators => $_getList(1);
}

