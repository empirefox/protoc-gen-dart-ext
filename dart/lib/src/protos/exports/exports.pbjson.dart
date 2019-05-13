///
//  Generated code. Do not modify.
//  source: protos/exports/exports.proto
///
// ignore_for_file: camel_case_types,non_constant_identifier_names,library_prefixes,unused_import,unused_shown_name

const Field$json = const {
  '1': 'Field',
  '2': const [
    const {'1': 'ref', '3': 1, '4': 1, '5': 9, '10': 'ref'},
    const {'1': 'name', '3': 2, '4': 1, '5': 9, '10': 'name'},
  ],
};

const Entity$json = const {
  '1': 'Entity',
  '2': const [
    const {'1': 'name', '3': 1, '4': 1, '5': 9, '10': 'name'},
    const {'1': 'fields', '3': 2, '4': 3, '5': 11, '6': '.exports.Field', '10': 'fields'},
  ],
};

const Package$json = const {
  '1': 'Package',
  '2': const [
    const {'1': 'entities', '3': 1, '4': 3, '5': 11, '6': '.exports.Entity', '10': 'entities'},
  ],
};

const Exports$json = const {
  '1': 'Exports',
  '2': const [
    const {'1': 'packages', '3': 1, '4': 3, '5': 11, '6': '.exports.Exports.PackagesEntry', '10': 'packages'},
  ],
  '3': const [Exports_PackagesEntry$json],
};

const Exports_PackagesEntry$json = const {
  '1': 'PackagesEntry',
  '2': const [
    const {'1': 'key', '3': 1, '4': 1, '5': 9, '10': 'key'},
    const {'1': 'value', '3': 2, '4': 1, '5': 11, '6': '.exports.Package', '10': 'value'},
  ],
  '7': const {'7': true},
};

