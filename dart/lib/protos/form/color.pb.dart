///
//  Generated code. Do not modify.
//  source: protos/form/color.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

import 'dart:core' as $core show bool, Deprecated, double, int, List, Map, override, String;

import 'package:protobuf/protobuf.dart' as $pb;

import 'color.pbenum.dart';

export 'color.pbenum.dart';

enum MaterialColor_Color {
  non, 
  primary, 
  accent, 
  notSet
}

class MaterialColor extends $pb.GeneratedMessage {
  static const $core.Map<$core.int, MaterialColor_Color> _MaterialColor_ColorByTag = {
    1 : MaterialColor_Color.non,
    2 : MaterialColor_Color.primary,
    3 : MaterialColor_Color.accent,
    0 : MaterialColor_Color.notSet
  };
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('MaterialColor', package: const $pb.PackageName('form'))
    ..a<NonColor>(1, 'non', $pb.PbFieldType.OM, NonColor.getDefault, NonColor.create)
    ..a<PrimaryColor>(2, 'primary', $pb.PbFieldType.OM, PrimaryColor.getDefault, PrimaryColor.create)
    ..a<AccentColor>(3, 'accent', $pb.PbFieldType.OM, AccentColor.getDefault, AccentColor.create)
    ..oo(0, [1, 2, 3])
    ..hasRequiredFields = false
  ;

  MaterialColor() : super();
  MaterialColor.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  MaterialColor.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  MaterialColor clone() => MaterialColor()..mergeFromMessage(this);
  MaterialColor copyWith(void Function(MaterialColor) updates) => super.copyWith((message) => updates(message as MaterialColor));
  $pb.BuilderInfo get info_ => _i;
  static MaterialColor create() => MaterialColor();
  MaterialColor createEmptyInstance() => create();
  static $pb.PbList<MaterialColor> createRepeated() => $pb.PbList<MaterialColor>();
  static MaterialColor getDefault() => _defaultInstance ??= create()..freeze();
  static MaterialColor _defaultInstance;

  MaterialColor_Color whichColor() => _MaterialColor_ColorByTag[$_whichOneof(0)];
  void clearColor() => clearField($_whichOneof(0));

  NonColor get non => $_getN(0);
  set non(NonColor v) { setField(1, v); }
  $core.bool hasNon() => $_has(0);
  void clearNon() => clearField(1);

  PrimaryColor get primary => $_getN(1);
  set primary(PrimaryColor v) { setField(2, v); }
  $core.bool hasPrimary() => $_has(1);
  void clearPrimary() => clearField(2);

  AccentColor get accent => $_getN(2);
  set accent(AccentColor v) { setField(3, v); }
  $core.bool hasAccent() => $_has(2);
  void clearAccent() => clearField(3);
}

class NonColor extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('NonColor', package: const $pb.PackageName('form'))
    ..e<NonColor_Color>(1, 'color', $pb.PbFieldType.OE, NonColor_Color.transparent, NonColor_Color.valueOf, NonColor_Color.values)
    ..hasRequiredFields = false
  ;

  NonColor() : super();
  NonColor.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  NonColor.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  NonColor clone() => NonColor()..mergeFromMessage(this);
  NonColor copyWith(void Function(NonColor) updates) => super.copyWith((message) => updates(message as NonColor));
  $pb.BuilderInfo get info_ => _i;
  static NonColor create() => NonColor();
  NonColor createEmptyInstance() => create();
  static $pb.PbList<NonColor> createRepeated() => $pb.PbList<NonColor>();
  static NonColor getDefault() => _defaultInstance ??= create()..freeze();
  static NonColor _defaultInstance;

  NonColor_Color get color => $_getN(0);
  set color(NonColor_Color v) { setField(1, v); }
  $core.bool hasColor() => $_has(0);
  void clearColor() => clearField(1);
}

class PrimaryColor extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('PrimaryColor', package: const $pb.PackageName('form'))
    ..e<PrimaryColor_Color>(1, 'color', $pb.PbFieldType.OE, PrimaryColor_Color.red, PrimaryColor_Color.valueOf, PrimaryColor_Color.values)
    ..e<PrimaryColor_Shade>(2, 'shade', $pb.PbFieldType.OE, PrimaryColor_Shade.s50, PrimaryColor_Shade.valueOf, PrimaryColor_Shade.values)
    ..hasRequiredFields = false
  ;

  PrimaryColor() : super();
  PrimaryColor.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  PrimaryColor.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  PrimaryColor clone() => PrimaryColor()..mergeFromMessage(this);
  PrimaryColor copyWith(void Function(PrimaryColor) updates) => super.copyWith((message) => updates(message as PrimaryColor));
  $pb.BuilderInfo get info_ => _i;
  static PrimaryColor create() => PrimaryColor();
  PrimaryColor createEmptyInstance() => create();
  static $pb.PbList<PrimaryColor> createRepeated() => $pb.PbList<PrimaryColor>();
  static PrimaryColor getDefault() => _defaultInstance ??= create()..freeze();
  static PrimaryColor _defaultInstance;

  PrimaryColor_Color get color => $_getN(0);
  set color(PrimaryColor_Color v) { setField(1, v); }
  $core.bool hasColor() => $_has(0);
  void clearColor() => clearField(1);

  PrimaryColor_Shade get shade => $_getN(1);
  set shade(PrimaryColor_Shade v) { setField(2, v); }
  $core.bool hasShade() => $_has(1);
  void clearShade() => clearField(2);
}

class AccentColor extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo('AccentColor', package: const $pb.PackageName('form'))
    ..e<AccentColor_Color>(1, 'color', $pb.PbFieldType.OE, AccentColor_Color.red, AccentColor_Color.valueOf, AccentColor_Color.values)
    ..e<AccentColor_Shade>(2, 'shade', $pb.PbFieldType.OE, AccentColor_Shade.s50, AccentColor_Shade.valueOf, AccentColor_Shade.values)
    ..hasRequiredFields = false
  ;

  AccentColor() : super();
  AccentColor.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromBuffer(i, r);
  AccentColor.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) : super.fromJson(i, r);
  AccentColor clone() => AccentColor()..mergeFromMessage(this);
  AccentColor copyWith(void Function(AccentColor) updates) => super.copyWith((message) => updates(message as AccentColor));
  $pb.BuilderInfo get info_ => _i;
  static AccentColor create() => AccentColor();
  AccentColor createEmptyInstance() => create();
  static $pb.PbList<AccentColor> createRepeated() => $pb.PbList<AccentColor>();
  static AccentColor getDefault() => _defaultInstance ??= create()..freeze();
  static AccentColor _defaultInstance;

  AccentColor_Color get color => $_getN(0);
  set color(AccentColor_Color v) { setField(1, v); }
  $core.bool hasColor() => $_has(0);
  void clearColor() => clearField(1);

  AccentColor_Shade get shade => $_getN(1);
  set shade(AccentColor_Shade v) { setField(2, v); }
  $core.bool hasShade() => $_has(1);
  void clearShade() => clearField(2);
}

